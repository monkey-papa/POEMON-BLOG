import Vue from "vue";
import VueRouter from "vue-router";
import store from "../store";

const originalPush = VueRouter.prototype.push;
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch((err) => err);
};

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    component: () => import("../views/home"),
    children: [
      {
        path: "/",
        name: "index",
        component: () => import("../views/index"),
      },
      {
        path: "/sort",
        name: "sort",
        component: () => import("../views/sort"),
      },
      {
        path: "/tags",
        name: "tags",
        component: () => import("../views/tags"),
      },
      {
        path: "/map",
        name: "map",
        component: () => import("../views/map"),
      },
      {
        path: "/article",
        name: "article",
        component: () => import("../views/article"),
      },
      {
        path: "/weiYan",
        name: "weiYan",
        component: () => import("../views/weiYan"),
      },
      {
        path: "/love",
        name: "love",
        component: () => import("../views/love"),
      },
      {
        path: "/travel",
        name: "travel",
        component: () => import("../views/travel"),
      },
      {
        path: "/tools",
        name: "tools",
        component: () => import("../views/tools"),
      },
      {
        path: "/message",
        name: "message",
        component: () => import("../views/message"),
      },
      {
        path: "/friend",
        name: "friend",
        component: () => import("../views/friend"),
      },
      {
        path: "/funny",
        name: "funny",
        component: () => import("../views/funny"),
      },
      {
        path: "/about",
        name: "about",
        component: () => import("../views/about"),
      },
      {
        path: "/user",
        name: "user",
        component: () => import("../views/user"),
      },
    ],
  },
  {
    path: "/backend",
    redirect: "/backendMain",
    meta: { requiresAuth: true },
    component: () => import("../views/admin/admin"),
    children: [
      {
        path: "/backendMain",
        name: "main",
        component: () => import("../views/admin/main"),
      },
      {
        path: "/webEdit",
        name: "webEdit",
        component: () => import("../views/admin/webEdit"),
      },
      {
        path: "/userList",
        name: "userList",
        component: () => import("../views/admin/userList"),
      },
      {
        path: "/postList",
        name: "postList",
        component: () => import("../views/admin/postList"),
      },
      {
        path: "/postEdit",
        name: "postEdit",
        component: () => import("../views/admin/postEdit"),
      },
      {
        path: "/sortList",
        name: "sortList",
        component: () => import("../views/admin/sortList"),
      },
      {
        path: "/commentList",
        name: "commentList",
        component: () => import("../views/admin/commentList"),
      },
      {
        path: "/treeHoleList",
        name: "treeHoleList",
        component: () => import("../views/admin/treeHoleList"),
      },
      {
        path: "/resourceList",
        name: "resourceList",
        component: () => import("../views/admin/resourceList"),
      },
      {
        path: "/loveList",
        name: "loveList",
        component: () => import("../views/admin/loveList"),
      },
      {
        path: "/resourcePathList",
        name: "resourcePathList",
        component: () => import("../views/admin/resourcePathList"),
      },
      {
        path: "/prohibitedWordsList",
        name: "prohibitedWordsList",
        component: () => import("../views/admin/prohibitedWordsList"),
      },
    ],
  },
  {
    path: "/verifyLogin",
    name: "verify",
    component: () => import("../views/admin/verify"),
  },
];

// 如果是访客剔除后台用户列表路由
if (
  store.state.currentAdmin.userType === 3 ||
  store.state.currentUser.userType === 3
) {
  routes[1].children = routes[1].children.filter(
    (item) => item.name != "userList"
  );
}

const router = new VueRouter({
  mode: "hash",
  routes: routes,
  scrollBehavior(to, from, savedPosition) {
    return { x: 0, y: 0 };
  },
});

router.beforeEach((to, from, next) => {
  // 每次切换页面时，调用进度条
  NProgress.start();
  store.commit("SET_SHOWLOADING", true);
  // 后台页面跳转的判断
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    if (
      !Boolean(localStorage.getItem("adminToken")) ||
      !JSON.parse(localStorage.getItem("vuex"))?.currentAdmin.id
    ) {
      localStorage.removeItem("adminToken");
      next({
        path: "/verifyLogin",
        query: { redirect: to.fullPath },
      });
    } else {
      next();
    }
  } else {
    next();
  }
});
router.afterEach((to) => {
  const title = to.fullPath;
  if (
    to.fullPath == "/about" ||
    to.fullPath == "/user" ||
    title.indexOf("/verifyLogin?redirect") != -1
  ) {
    setTimeout(() => {
      store.commit("SET_SHOWLOADING", false);
    }, 500);
    return;
  }
  setTimeout(() => {
    store.commit("SET_SHOWLOADING", false);
  }, 2500);
});
export default router;
