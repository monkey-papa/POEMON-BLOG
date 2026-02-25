import { createRouter, createWebHistory, type RouteRecordRaw } from "vue-router";
import { startRouteNavigation, endRouteNavigation, initLoading } from "../utils/request";
import { useStore } from "../stores";
import type { MainStoreType } from "../stores";
import type { IpCityResult } from "@/types";

// 延迟获取 store 实例，避免在 Pinia 初始化前访问
let storeInstance: MainStoreType | null = null;
const getStore = (): MainStoreType => {
  if (!storeInstance) {
    storeInstance = useStore();
  }
  return storeInstance;
};

// 位置信息获取（只在首次加载时调用）
let locationFetched = false;
const fetchLocationInfo = async (): Promise<void> => {
  const mainStore = getStore();

  // 如果已经加载过或正在加载，跳过
  if (mainStore.locationInfo.isLoaded || mainStore.locationInfo.isLoading || locationFetched) {
    return;
  }

  locationFetched = true;
  mainStore.setLocationLoading(true);

  try {
    // 动态导入避免循环依赖
    const { default: $common } = await import("../utils/common");
    const result: IpCityResult = await $common.getIpAndCity();

    if (result && (result.city || result.province)) {
      mainStore.loadLocationInfo(result);
    } else {
      mainStore.setLocationLoading(false);
    }
  } catch {
    mainStore.setLocationLoading(false);
  }
};

const routes: RouteRecordRaw[] = [
  {
    path: "/",
    component: () => import("../views/home.vue"),
    children: [
      {
        path: "/",
        name: "index",
        component: () => import("../views/index.vue"),
      },
      {
        path: "/sort",
        name: "sort",
        component: () => import("../views/sort.vue"),
      },
      {
        path: "/tags",
        name: "tags",
        component: () => import("../views/tags.vue"),
      },
      {
        path: "/map",
        name: "map",
        component: () => import("../views/map.vue"),
      },
      {
        path: "/article",
        name: "article",
        component: () => import("../views/article.vue"),
      },
      {
        path: "/weiYan",
        name: "weiYan",
        component: () => import("../views/weiYan.vue"),
      },
      {
        path: "/love",
        name: "love",
        component: () => import("../views/love.vue"),
      },
      {
        path: "/travel",
        name: "travel",
        component: () => import("../views/travel.vue"),
      },
      {
        path: "/tools",
        name: "tools",
        component: () => import("../views/tools.vue"),
      },
      {
        path: "/message",
        name: "message",
        component: () => import("../views/message.vue"),
      },
      {
        path: "/friend",
        name: "friend",
        component: () => import("../views/friend.vue"),
      },
      {
        path: "/funny",
        name: "funny",
        component: () => import("../views/funny.vue"),
      },
      {
        path: "/about",
        name: "about",
        component: () => import("../views/about.vue"),
      },
      {
        path: "/user",
        name: "user",
        component: () => import("../views/user.vue"),
      },
    ],
  },
  {
    path: "/backend",
    redirect: "/backendMain",
    meta: { requiresAuth: true },
    component: () => import("../views/admin/admin.vue"),
    children: [
      {
        path: "/backendMain",
        name: "main",
        component: () => import("../views/admin/main.vue"),
      },
      {
        path: "/webEdit",
        name: "webEdit",
        component: () => import("../views/admin/webEdit.vue"),
      },
      {
        path: "/userList",
        name: "userList",
        component: () => import("../views/admin/userList.vue"),
      },
      {
        path: "/postList",
        name: "postList",
        component: () => import("../views/admin/postList.vue"),
      },
      {
        path: "/postEdit",
        name: "postEdit",
        component: () => import("../views/admin/postEdit.vue"),
      },
      {
        path: "/sortList",
        name: "sortList",
        component: () => import("../views/admin/sortList.vue"),
      },
      {
        path: "/weiYanList",
        name: "weiYanList",
        component: () => import("../views/admin/weiYanList.vue"),
      },
      {
        path: "/progressList",
        name: "progressList",
        component: () => import("../views/admin/progressList.vue"),
      },
      {
        path: "/commentList",
        name: "commentList",
        component: () => import("../views/admin/commentList.vue"),
      },
      {
        path: "/treeHoleList",
        name: "treeHoleList",
        component: () => import("../views/admin/treeHoleList.vue"),
      },
      {
        path: "/resourceList",
        name: "resourceList",
        component: () => import("../views/admin/resourceList.vue"),
      },
      {
        path: "/loveList",
        name: "loveList",
        component: () => import("../views/admin/loveList.vue"),
      },
      {
        path: "/resourcePathList",
        name: "resourcePathList",
        component: () => import("../views/admin/resourcePathList.vue"),
      },
      {
        path: "/prohibitedWordsList",
        name: "prohibitedWordsList",
        component: () => import("../views/admin/prohibitedWordsList.vue"),
      },
    ],
  },
  {
    path: "/verifyLogin",
    name: "verify",
    component: () => import("../views/admin/verify.vue"),
  },
];

// 如果是访客剔除后台用户列表路由（延迟检查）
const filterVisitorRoutes = (): void => {
  try {
    const mainStore = getStore();
    if (mainStore.currentAdmin?.userType === 3) {
      const backendRoute = routes[1];
      if (backendRoute && backendRoute.children) {
        backendRoute.children = backendRoute.children.filter((item) => item.name !== "userList");
      }
    }
  } catch {
    // store 尚未初始化，忽略
  }
};

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior() {
    return { left: 0, top: 0 };
  },
});

router.beforeEach((to, from, next) => {
  // 每次切换页面时，调用进度条
  NProgress.start();

  // 首次路由时过滤访客路由
  filterVisitorRoutes();

  // 如果是首次加载（刷新页面），初始化 loading 状态
  const isInitialLoad = from.name === null || from.name === undefined;
  if (isInitialLoad) {
    initLoading();
    // 首次加载时获取位置信息（异步，不阻塞路由）
    fetchLocationInfo();
  } else {
    // 路由跳转时，开始路由导航，打开 loading
    startRouteNavigation();
  }

  const mainStore = getStore();

  // 如果已登录用户访问后台登录页面，自动跳转到后台首页
  if (to.path === "/verifyLogin") {
    if (mainStore.currentAdmin?.accessToken) {
      next({ path: "/backendMain" });
      return;
    }
  }

  // 后台页面跳转的判断
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    const admin = mainStore.currentAdmin;
    // 必须有 Token 且用户类型不是普通用户（userType !== 2）
    // Boss(0)、Admin(1)、Visitor(3) 可以访问后台
    const hasValidToken = !!admin?.accessToken;
    const isAllowedUser = admin?.userType !== undefined && admin.userType !== 2;

    if (!hasValidToken || !isAllowedUser) {
      // 清除无效的管理员信息
      if (hasValidToken && !isAllowedUser) {
        mainStore.loadCurrentAdmin(null);
      }
      next({
        path: "/verifyLogin",
        query: { redirect: to.fullPath },
      });
    } else {
      next();
    }
  } else {
    // 对于非需要认证的页面
    // 如果未登录用户从其他页面跳转到 /user，保存来源页面用于登录后返回
    const isNavigatingToUser = to.path === "/user";
    const isNotLoggedIn = !mainStore.currentUser?.accessToken;
    const hasNoRedirect = !to.query.redirect;
    const isFromValidPage = from.path !== "/" && from.path !== "/user";
    const isNotInitialLoad = from.name !== null && from.name !== undefined;

    if (
      isNavigatingToUser &&
      isNotLoggedIn &&
      hasNoRedirect &&
      isFromValidPage &&
      isNotInitialLoad
    ) {
      // 使用 replace 替换当前导航，添加 redirect 参数，不产生额外的历史记录
      next({
        path: "/user",
        query: { redirect: from.fullPath },
        replace: true,
      });
    } else {
      next();
    }
  }
});

router.afterEach(() => {
  // 完成进度条
  NProgress.done();

  // 延迟调用 endRouteNavigation，等待组件 setup 执行并发起请求
  // 增加延迟时间到 200ms，确保组件有足够时间发起请求
  setTimeout(() => {
    endRouteNavigation();
  }, 200);
});

export default router;
