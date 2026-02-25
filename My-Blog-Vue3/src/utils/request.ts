import axios, { type AxiosResponse, type InternalAxiosRequestConfig } from "axios";
import constant from "./constant";
import type { Router } from "vue-router";
import type { CurrentUser, CurrentAdmin } from "@/types";

// Store 类型定义
interface MainStore {
  isShowLoading: boolean;
  currentUser: CurrentUser | null;
  currentAdmin: CurrentAdmin | null;
  setShowLoading(data: boolean): void;
  loadCurrentUser(user: CurrentUser | null): void;
  loadCurrentAdmin(admin: CurrentAdmin | null): void;
}

type UseStoreFn = () => MainStore;

// 延迟获取 store 实例，避免在 Pinia 初始化前访问
let storeInstance: MainStore | null = null;
let useStoreFn: UseStoreFn | null = null;

// 路由实例引用（用于编程式导航，避免使用 window.location.href）
let routerInstance: Router | null = null;

// 初始化 store 引用（在 main.ts 中调用）
export const setStoreGetter = (fn: UseStoreFn): void => {
  useStoreFn = fn;
};

// 初始化 router 引用（在 main.ts 中调用）
export const setRouter = (router: Router): void => {
  routerInstance = router;
};

const getStore = (): MainStore => {
  if (!storeInstance && useStoreFn) {
    storeInstance = useStoreFn();
  }
  return storeInstance!;
};

// 使用 Vue Router 进行导航，避免页面刷新
// 增加防重复跳转逻辑
let isNavigating = false;
const navigateTo = (path: string, query: Record<string, string> = {}): void => {
  if (routerInstance) {
    // 检查是否已经在目标路径，避免重复跳转
    const currentPath = routerInstance.currentRoute?.value?.path;
    if (currentPath === path) {
      return;
    }
    // 防止重复导航
    if (isNavigating) {
      return;
    }
    isNavigating = true;
    routerInstance.push({ path, query }).finally(() => {
      isNavigating = false;
    });
  } else {
    // 降级方案：如果 router 还未初始化，使用 location
    window.location.href = path;
  }
};

// 请求计数器，用于管理 loading 状态
let loadingCount = 0;
// 最小显示时间（毫秒），避免 loading 闪烁
const MIN_LOADING_TIME = 300;

let loadingStartTime = 0;
let loadingTimer: ReturnType<typeof setTimeout> | null = null;

// 路由跳转标志
let isRouteNavigation = false;
let routeNavigationId = 0;
// 记录是否有请求开始（用于判断请求是否已发起）
let hasRequestStarted = false;

const startLoading = (): void => {
  if (!isRouteNavigation) return;
  loadingCount++;
  hasRequestStarted = true; // 标记有请求开始
  if (loadingCount === 1) {
    const mainStore = getStore();
    if (!mainStore.isShowLoading) {
      mainStore.setShowLoading(true);
    }
    if (loadingStartTime === 0) {
      loadingStartTime = Date.now();
    }
  }
};

const endLoading = (): void => {
  if (!isRouteNavigation) return;
  loadingCount = Math.max(loadingCount - 1, 0);
};

export const startRouteNavigation = (): void => {
  if (loadingTimer) {
    clearTimeout(loadingTimer);
    loadingTimer = null;
  }
  routeNavigationId++;
  loadingCount = 0;
  hasRequestStarted = false; // 重置请求开始标记
  isRouteNavigation = true;
  loadingStartTime = Date.now();
  getStore().setShowLoading(true);
};

export const endRouteNavigation = (): void => {
  if (!isRouteNavigation) return;
  const currentNavigationId = routeNavigationId;
  if (loadingTimer) {
    clearTimeout(loadingTimer);
    loadingTimer = null;
  }

  const checkAndClose = (): void => {
    if (currentNavigationId !== routeNavigationId || !isRouteNavigation) return;

    // 如果有请求在进行中，继续等待
    if (loadingCount > 0) {
      const maxWaitTime = 15000; // 最大等待 15 秒
      const elapsed = Date.now() - loadingStartTime;
      if (elapsed < maxWaitTime) {
        setTimeout(checkAndClose, 100);
      } else {
        // 超时强制关闭
        closeLoading(currentNavigationId);
      }
      return;
    }

    // loadingCount === 0 时，可能是请求还没开始，或者请求已完成
    // 如果曾经有请求开始过，说明请求已完成
    if (hasRequestStarted) {
      ensureMinLoadingTime(currentNavigationId);
      return;
    }

    // 如果没有请求开始过，等待一段时间看是否有请求发起
    const elapsed = Date.now() - loadingStartTime;
    const waitForRequest = 500; // 等待 500ms 让组件发起请求

    if (elapsed < waitForRequest) {
      setTimeout(checkAndClose, 50);
    } else {
      // 超过等待时间仍没有请求，直接关闭 loading
      ensureMinLoadingTime(currentNavigationId);
    }
  };

  const ensureMinLoadingTime = (navId: number): void => {
    const elapsed = Date.now() - loadingStartTime;
    const remaining = Math.max(MIN_LOADING_TIME - elapsed, 0);

    if (remaining > 0) {
      loadingTimer = setTimeout(() => {
        if (navId === routeNavigationId && isRouteNavigation && loadingCount === 0) {
          closeLoading(navId);
        }
      }, remaining);
    } else {
      closeLoading(navId);
    }
  };

  const closeLoading = (navId: number): void => {
    if (navId === routeNavigationId) {
      getStore().setShowLoading(false);
      loadingTimer = null;
      loadingStartTime = 0;
      isRouteNavigation = false;
      hasRequestStarted = false;
    }
  };

  // 开始检查
  setTimeout(checkAndClose, 50);
};

export const initLoading = (): void => {
  loadingCount = 0;
  hasRequestStarted = false; // 重置请求开始标记
  if (loadingTimer) {
    clearTimeout(loadingTimer);
    loadingTimer = null;
  }
  isRouteNavigation = true;
  loadingStartTime = Date.now();
  getStore().setShowLoading(true);
};

axios.defaults.baseURL = constant.baseURL;
axios.defaults.timeout = 30000;

// 请求拦截器
axios.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    startLoading();
    return config;
  },
  (error: unknown) => {
    endLoading();
    return Promise.reject(error);
  }
);

// API 响应数据结构
interface ApiResponseData<T = unknown> {
  code: number;
  msg: string;
  data: T;
  success?: boolean;
}

// 响应拦截器
axios.interceptors.response.use(
  (response: AxiosResponse<ApiResponseData>) => {
    endLoading();
    const { code, msg } = response.data || {};

    // 处理 Token 自动刷新
    const newToken = response.headers["x-new-token"] as string | undefined;
    if (newToken) {
      const mainStore = getStore();
      const isAdmin = isAdminPage();
      if (isAdmin && mainStore.currentAdmin?.accessToken) {
        mainStore.loadCurrentAdmin({
          ...mainStore.currentAdmin,
          accessToken: newToken,
        });
        console.log("🔄 Token 已自动刷新（管理员）");
      } else if (mainStore.currentUser?.accessToken) {
        mainStore.loadCurrentUser({
          ...mainStore.currentUser,
          accessToken: newToken,
        });
        console.log("🔄 Token 已自动刷新（用户）");
      }
    }

    // code === 0 表示成功
    if (code !== undefined && code !== 0) {
      // 需要重新登录
      if (code === 300 || code === 401) {
        const mainStore = getStore();
        mainStore.loadCurrentUser(null);
        mainStore.loadCurrentAdmin(null);
        navigateTo("/user");
      }
      return Promise.reject(new Error(msg || "请求失败"));
    }
    return response;
  },
  (error: { response?: AxiosResponse<{ msg?: string }> }) => {
    endLoading();
    // 处理网络错误
    if (!error.response) {
      return Promise.reject(new Error("网络连接失败，请检查网络"));
    }
    // 处理 HTTP 错误状态码
    const status = error.response.status;
    const msg = error.response.data?.msg;

    switch (status) {
      case 401: {
        const mainStore = getStore();
        mainStore.loadCurrentUser(null);
        mainStore.loadCurrentAdmin(null);
        navigateTo("/user");
        return Promise.reject(new Error(msg || "请先登录"));
      }
      case 403:
        return Promise.reject(new Error(msg || "没有操作权限"));
      case 429:
        return Promise.reject(new Error(msg || "请求过于频繁，请稍后再试"));
      case 500:
        return Promise.reject(new Error(msg || "服务器错误，请稍后再试"));
      default:
        return Promise.reject(new Error(msg || `请求失败 (${status})`));
    }
  }
);

/**
 * 判断当前是否在后台管理页面
 */
const isAdminPage = (): boolean => {
  const pathname = window.location.pathname;
  const adminPaths = [
    "/backendMain",
    "/webEdit",
    "/userList",
    "/postList",
    "/postEdit",
    "/sortList",
    "/weiYanList",
    "/progressList",
    "/commentList",
    "/treeHoleList",
    "/resourceList",
    "/loveList",
    "/resourcePathList",
    "/prohibitedWordsList",
  ];
  return adminPaths.some((path) => pathname.startsWith(path));
};

/**
 * 获取认证请求头
 */
const getAuthConfig = (): { headers: Record<string, string> } => {
  const mainStore = getStore();
  const isAdmin = isAdminPage();

  // 安全获取 token，防止 store 未初始化或未登录时出错
  let token = "";
  if (mainStore) {
    token =
      (isAdmin ? mainStore.currentAdmin?.accessToken : mainStore.currentUser?.accessToken) || "";
  }

  return {
    headers: {
      Authorization: token ? "Token " + token : "",
    },
  };
};

export interface HttpClient {
  post<T = unknown>(url: string, params?: object, json?: boolean): Promise<T>;
  uploadQiniu<T = unknown>(url: string, param: FormData): Promise<T>;
}

const http: HttpClient = {
  /**
   * POST 请求
   * @param url - 请求地址
   * @param params - 请求参数
   * @param json - 是否使用 JSON 格式（默认 true）
   * @returns 直接返回 data 字段内容
   */
  post<T = unknown>(url: string, params: object = {}, json = true): Promise<T> {
    const config = getAuthConfig();
    config.headers["Content-Type"] = json
      ? "application/json"
      : "application/x-www-form-urlencoded";

    const data = json ? params : new URLSearchParams(params as Record<string, string>).toString();

    return new Promise((resolve, reject) => {
      axios
        .post<ApiResponseData<T>>(url, data, config)
        .then((res) => {
          resolve(res.data.data);
        })
        .catch((err: Error) => {
          reject(err);
        });
    });
  },

  /**
   * 上传图片
   */
  uploadQiniu<T = unknown>(url: string, param: FormData): Promise<T> {
    const authConfig = getAuthConfig();
    const config = {
      headers: {
        ...authConfig.headers,
        "Content-Type": "multipart/form-data",
      },
      timeout: 60000,
    };
    return new Promise((resolve, reject) => {
      axios
        .post<ApiResponseData<T>>(url, param, config)
        .then((res) => {
          resolve(res.data.data);
        })
        .catch((err: Error) => {
          reject(err);
        });
    });
  },
};

export default http;
