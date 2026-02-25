import { defineStore } from "pinia";
import constant from "@/utils/constant";
import type {
  SortInfo,
  LabelInfo,
  CurrentUser,
  CurrentAdmin,
  WebInfo,
  LocationInfo,
  PageView,
  NewArticle,
  PublicUserListItem,
  ToolbarState,
  WeatherInfo,
} from "@/types";

// Store State 接口
interface MainState {
  toolbar: ToolbarState;
  sortInfo: SortInfo[];
  /** null 表示未登录 */
  currentUser: CurrentUser | null;
  /** null 表示未登录 */
  currentAdmin: CurrentAdmin | null;
  webInfo: WebInfo;
  locationInfo: LocationInfo;
  changeBg: string;
  isShowLoading: boolean;
  top: number;
  pageView: PageView;
  newArticles: NewArticle[];
  articleTotal: number;
  userList: PublicUserListItem[];
}

// 位置信息输入参数
interface LocationInfoInput {
  city?: string;
  province?: string;
  weather?: WeatherInfo[];
  tip?: string;
}

// WebInfo 输入参数（允许 webTitle 和其他字段是字符串）
interface WebInfoInput {
  id?: number;
  webName?: string;
  webTitle?: string[] | string;
  notices?: string[] | string;
  footer?: string;
  backgroundImage?: string;
  avatar?: string;
  randomCover?: string[] | string;
  waifuJson?: string;
  waifu?: string;
  status?: boolean;
}

export const useStore = defineStore("main", {
  state: (): MainState => ({
    toolbar: {
      visible: false,
      enter: true,
    },
    sortInfo: [],
    currentUser: null,
    currentAdmin: null,
    webInfo: {
      id: 0,
      webName: "",
      webTitle: [],
      notices: [],
      footer: "",
      backgroundImage: "",
      avatar: "",
      randomCover: [],
      status: true,
    },
    // 位置和天气信息（路由初始化时获取）
    locationInfo: {
      city: "",
      province: "", // 省份
      weather: [], // 多天天气数据
      tip: "",
      isLoading: false,
      isLoaded: false, // 是否已加载过
    },
    changeBg: `url(${constant.defaultBackground})`,
    isShowLoading: true, // 初始化为 true，避免首次加载时闪烁
    top: 0,
    pageView: {},
    newArticles: [],
    articleTotal: 0,
    userList: [],
  }),

  getters: {
    // 访客禁用按钮权限
    permissions: (state): string[] => {
      return state.currentAdmin?.userType === 3 ? [] : ["user:visit:read"];
    },

    // 文章总数
    totalArticles: (state): number => {
      if (state.sortInfo !== null && state.sortInfo.length !== 0) {
        if (state.sortInfo.length === 1) {
          return state.sortInfo[0]?.countOfSort ?? 0;
        } else {
          return state.sortInfo.reduce((prev: number, curr: SortInfo) => {
            return prev + (curr.countOfSort ?? 0);
          }, 0);
        }
      } else {
        return 0;
      }
    },

    // 导航栏
    navigationBar: (state): SortInfo[] => {
      if (state.sortInfo !== null && state.sortInfo.length !== 0) {
        return state.sortInfo.filter((f) => f.sortType === 0);
      } else {
        return [];
      }
    },

    // 标签信息
    labelInfo: (state): LabelInfo[] => {
      if (!state.sortInfo || !Array.isArray(state.sortInfo)) {
        return [];
      }
      let labels: LabelInfo[] = [];
      for (let index = 0; index < state.sortInfo.length; index++) {
        const sortLabels = state.sortInfo[index]?.labels;
        if (Array.isArray(sortLabels)) {
          labels = labels.concat(sortLabels.filter((label): label is LabelInfo => label != null));
        }
      }
      return labels;
    },
  },

  actions: {
    // 设置用户列表
    setUserList(userList: PublicUserListItem[]): void {
      this.userList = userList;
    },

    // 设置文章总数
    setArticleTotal(articleTotal: number): void {
      this.articleTotal = articleTotal;
    },

    // 设置最新文章
    setNewArticles(newArticles: NewArticle[]): void {
      this.newArticles = newArticles;
    },

    // 设置页面浏览量
    setPageView(pageView: PageView): void {
      this.pageView = pageView;
    },

    // 改变工具栏状态
    changeToolbarStatus(toolbarState: ToolbarState): void {
      this.toolbar = toolbarState;
    },

    // 加载分类信息
    loadSortInfo(sortInfo: SortInfo[]): void {
      if (sortInfo !== null && sortInfo.length !== 0) {
        this.sortInfo = sortInfo.sort((s1, s2) => (s1.priority ?? 0) - (s2.priority ?? 0));
      }
    },

    // 加载当前用户（传 null 表示登出）
    loadCurrentUser(user: CurrentUser | null): void {
      this.currentUser = user;
    },

    // 加载当前管理员（传 null 表示登出）
    loadCurrentAdmin(user: CurrentAdmin | null): void {
      this.currentAdmin = user;
    },

    // 加载网站信息
    loadWebInfo(webInfo: WebInfoInput): void {
      if (!webInfo || typeof webInfo !== "object") {
        console.warn("loadWebInfo: 无效的网站信息数据");
        return;
      }

      const processedWebInfo: WebInfo = {
        ...webInfo,
        id: webInfo.id ?? 0,
        webName: webInfo.webName ?? "",
        webTitle: [],
        notices: [],
        footer: webInfo.footer ?? "",
        backgroundImage: webInfo.backgroundImage ?? "",
        avatar: webInfo.avatar ?? "",
        randomCover: [],
        status: webInfo.status ?? true,
      };

      // 安全解析 webTitle
      if (webInfo.webTitle && typeof webInfo.webTitle === "string") {
        processedWebInfo.webTitle = webInfo.webTitle.split("");
      } else if (Array.isArray(webInfo.webTitle)) {
        processedWebInfo.webTitle = webInfo.webTitle;
      }

      // 安全解析 notices
      if (webInfo.notices && typeof webInfo.notices === "string") {
        try {
          processedWebInfo.notices = JSON.parse(webInfo.notices);
        } catch {
          processedWebInfo.notices = [];
        }
      } else if (Array.isArray(webInfo.notices)) {
        processedWebInfo.notices = webInfo.notices;
      }

      // 安全解析 randomCover
      if (webInfo.randomCover && typeof webInfo.randomCover === "string") {
        try {
          processedWebInfo.randomCover = JSON.parse(webInfo.randomCover);
        } catch {
          processedWebInfo.randomCover = [];
        }
      } else if (Array.isArray(webInfo.randomCover)) {
        processedWebInfo.randomCover = webInfo.randomCover;
      }

      this.webInfo = processedWebInfo;
    },

    // 改变背景
    changeBgBox(changeBg: string): void {
      this.changeBg = changeBg;
    },

    // 设置 loading 状态
    setShowLoading(data: boolean): void {
      this.isShowLoading = data;
    },

    // 设置滚动百分比
    setTopPercentage(top: number): void {
      this.top = top;
    },

    // 设置位置信息加载状态
    setLocationLoading(isLoading: boolean): void {
      this.locationInfo.isLoading = isLoading;
    },

    // 加载位置和天气信息
    loadLocationInfo(info: LocationInfoInput): void {
      this.locationInfo = {
        city: info.city || "",
        province: info.province || "",
        weather: info.weather || [],
        tip: info.tip || "",
        isLoading: false,
        isLoaded: true,
      };
    },

    // 重置位置信息（用于强制刷新）
    resetLocationInfo(): void {
      this.locationInfo = {
        city: "",
        province: "",
        weather: [],
        tip: "",
        isLoading: false,
        isLoaded: false,
      };
    },
  },

  // 持久化配置
  persist: {
    key: "main",
    storage: localStorage,
    // 只持久化以下字段（其他字段每次刷新都会重置）
    pick: [
      "toolbar",
      "sortInfo",
      "currentUser",
      "currentAdmin",
      "webInfo",
      "changeBg",
      "top",
      "pageView",
      "newArticles",
      "articleTotal",
      "userList",
    ],
  },
});

// 导出 Store 类型
export type MainStoreType = ReturnType<typeof useStore>;
