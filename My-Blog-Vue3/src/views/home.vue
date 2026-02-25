<template>
  <div class="home-container">
    <!-- el过渡动画 -->
    <transition name="el-fade-in-linear">
      <!-- 导航栏 -->
      <div
        v-show="toolbar.visible"
        @mouseenter="hoverEnter = true"
        @mouseleave="hoverEnter = false"
        :class="[
          { enter: toolbar.enter },
          {
            hoverEnter: hoverEnter && !toolbar.enter,
          },
        ]"
        class="toolbar-content myBetween"
      >
        <!-- 网站名称 -->
        <div @click="router.push({ path: '/' })" class="toolbar-title myCenter">
          <h2>
            {{ webInfo.webName }}
          </h2>
        </div>
        <!-- 手机抽屉按钮 -->
        <div
          v-if="$common.mobile() || mobile"
          class="toolbar-mobile-menu"
          @click="toolbarDrawer = !toolbarDrawer"
          :class="{ enter: toolbar.enter }"
        >
          <el-icon :size="30" color="var(--red)"><Operation /></el-icon>
        </div>
        <!-- 常规头部导航列表 -->
        <NavMenu
          v-else
          :isMobile="false"
          @openPcGame="openPcGame"
          @edit="edit"
          @openRandomArticle="openRandomArticle"
          @openChangeBg="openChangeBg"
          @handleSakura="handleSakura"
          @changeColor="changeColor"
          @logout="logout"
        />
      </div>
    </transition>
    <div id="main-container">
      <router-view />
    </div>
    <!-- 猫 -->
    <div href="#" class="cd-top" v-if="!$common.mobile()" @click="toTop()"></div>
    <!-- 右下角按钮 -->
    <ToolButton
      :isMobile="$common.mobile()"
      :showBackTop="toolButton"
      :topPercentage="topPercentage"
      :topPercentageType="topPercentageType"
      :mouseAnimation="mouseAnimation"
      @toTop="toTop"
      @changeMouseAnimation="changeMouseAnimation"
    />
    <!-- 点击动画 -->
    <canvas
      v-if="mouseAnimation"
      id="mousedown"
      style="position: fixed; left: 0; top: 0; pointer-events: none; z-index: 1000"
    >
    </canvas>
    <!-- 移动端抽屉 -->
    <MobileDrawer
      v-model="toolbarDrawer"
      @openPcGame="openPcGame"
      @edit="edit"
      @openRandomArticle="openRandomArticle"
      @openChangeBg="openChangeBg"
      @changeColor="changeColor"
      @handleSakura="handleSakura"
      @logout="logout"
    />
    <!-- 小游戏 -->
    <MiniGame v-model="disGame" :gameUrl="game" />
    <!-- 修改背景 -->
    <ChangeBackground
      ref="changeBgRef"
      v-model="changeBgBox"
      :themeMap="themeMap"
      @defaultBtn="defaultBtn"
      @httpInputBtn="httpInputBtn"
      @setColor="setColor"
      @changeBg="changeBg"
      @handleChangeBg="handleChangeBg"
    />
  </div>
</template>
<script setup lang="ts">
import { notifySuccess, notifyWarning } from "@/utils/notify";
import api from "@/api";
import { ref, computed, watch, onMounted, onBeforeUnmount, nextTick } from "vue";
import type { ToolbarState, CurrentUser, WebInfo, SortInfo, ThemeMapConfigItem } from "@/types";
import mousedown from "../utils/mousedown";
import NavMenu from "./common/NavMenu.vue";
import MobileDrawer from "./common/MobileDrawer.vue";
import ChangeBackground from "./common/ChangeBackground.vue";
import MiniGame from "./common/MiniGame.vue";
import ToolButton from "./common/ToolButton.vue";
import { useStore } from "@/stores";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
import { useRouter } from "vue-router";

const store = useStore();
const router = useRouter();
const { $route, $common, $constant } = useGlobalProperties();

const disGame = ref<boolean>(false);
const pcGameList = ref<number[]>([1, 2, 3, 4, 5, 6, 7, 8]);
const game = ref<string>($constant.gameURL);
const toolButton = ref<boolean>(false);
const hoverEnter = ref<boolean>(false);
const mouseAnimation = ref<boolean>(false);
const isLight = ref<boolean>(false); // false是黑夜,显示月亮按钮
const scrollTop = ref<number>(0);
const toolbarDrawer = ref<boolean>(false);
const mobile = ref<boolean>(false);
const changeBgBox = ref<boolean>(false);
const editFlag = ref<boolean>(false);
const sakuraFlag = ref<boolean>(false);
const topPercentageType = ref<boolean>(false);
const themeMap = ref<ThemeMapConfigItem[]>($constant.themeMapConfig);
const themeActive = ref<string>(localStorage.getItem("themeColor") || "#425aef");
const changeBgRef = ref<InstanceType<typeof ChangeBackground> | null>(null);

const toolbar = computed(() => store.toolbar);
const topPercentage = computed(() => store.top);
const webInfo = computed(() => store.webInfo);
const changeBgState = computed(() => store.changeBg);
const articleTotal = computed(() => store.articleTotal);

const changeToolbarStatus = (value: ToolbarState): void => store.changeToolbarStatus(value);
const loadCurrentUser = (value: CurrentUser | null): void => store.loadCurrentUser(value);
const loadWebInfo = (value: WebInfo): void => store.loadWebInfo(value);
const loadSortInfo = (value: SortInfo[]): void => store.loadSortInfo(value);
const changeBgBoxMutation = (value: string): void => store.changeBgBox(value);

watch(
  () => $route.path,
  (to: string): void => {
    if (to === "/funny") {
      topPercentageType.value = true;
    } else {
      topPercentageType.value = false;
    }
  },
  { immediate: true }
);

watch(scrollTop, (newScrollTop: number, oldScrollTop: number): void => {
  // 如果滑动距离超过屏幕高度三分之一视为进入页面，背景改为蓝色
  const enter = newScrollTop > window.innerHeight / 2;
  const top = newScrollTop - oldScrollTop < 0;
  const isShow = newScrollTop - window.innerHeight > 30;
  toolButton.value = isShow;
  if (isShow && !$common.mobile()) {
    if (window.innerHeight > 950) {
      window.$(".cd-top").css("top", "0");
    } else {
      window.$(".cd-top").css("top", window.innerHeight - 950 + "px");
    }
  } else if (!isShow && !$common.mobile()) {
    window.$(".cd-top").css("top", "-900px");
  }
  // 导航栏显示与颜色
  const toolbarStatus = {
    enter,
    visible: top,
  };
  changeToolbarStatus(toolbarStatus);
});

onMounted((): void => {
  // 默认关闭樱花
  handleSakura();

  if (localStorage.getItem("theme") == null) {
    localStorage.setItem("theme", "true");
  }
  if (mouseAnimation.value) {
    mousedown();
  }
  window.addEventListener("scroll", onScrollPage, { passive: true });

  // 关灯
  if (localStorage.getItem("theme") == "false") {
    isLight.value = false;
  } else {
    isLight.value = true;
  }
  checkSwitchBtnRef();
  changeBg();

  const toolbarStatus = {
    enter: false,
    visible: true,
  };
  // 设置默认主题颜色
  setColor(localStorage.getItem("themeColor") || "#425aef");
  changeToolbarStatus(toolbarStatus);
  getWebInfo();
  getSortInfo();
  mobile.value = document.body.clientWidth < 600;
  window.addEventListener("resize", onWindowResize);
});

// 定时器和重试次数
let switchBtnCheckTimer: ReturnType<typeof setTimeout> | null = null;
const MAX_RETRY_COUNT = 50; // 最大重试次数（5秒）

const onWindowResize = (): void => {
  const docWidth = document.body.clientWidth;
  if (docWidth < 600) {
    mobile.value = true;
  } else {
    mobile.value = false;
  }
};

onBeforeUnmount((): void => {
  window.removeEventListener("scroll", onScrollPage);
  window.removeEventListener("resize", onWindowResize);
  // 清理定时器
  if (switchBtnCheckTimer) {
    clearTimeout(switchBtnCheckTimer);
    switchBtnCheckTimer = null;
  }
});

// 轮询检查 switchBtnRef 是否存在
const checkSwitchBtnRef = (retryCount: number = 0): void => {
  const switchBtnRef = document.querySelector("#switchBtnRef") as HTMLInputElement | null;

  if (switchBtnRef) {
    // 找到元素，设置状态并清理定时器
    switchBtnRef.checked = !isLight.value;
    if (switchBtnCheckTimer) {
      clearTimeout(switchBtnCheckTimer);
      switchBtnCheckTimer = null;
    }
  } else if (retryCount < MAX_RETRY_COUNT) {
    // 未找到且未超过最大重试次数，继续尝试
    switchBtnCheckTimer = setTimeout(() => {
      checkSwitchBtnRef(retryCount + 1);
    }, 100);
  } else {
    // 超过最大重试次数，停止尝试
    if (switchBtnCheckTimer) {
      clearTimeout(switchBtnCheckTimer);
      switchBtnCheckTimer = null;
    }
  }
};

const setColor = (color: string): void => {
  themeActive.value = color;
  const root = document.querySelector(":root") as HTMLElement | null;
  if (root) {
    root.style.setProperty("--themeColor", color);
    $common.getThemeRgb();
  }
  document.documentElement.dataset.theme =
    localStorage.getItem("theme") === "false" ? "theme2-dark" : "theme2";
  localStorage.setItem("themeColor", color);
};

const openRandomArticle = (): void => {
  // 我的数据库是从id为12的文章开始的,所以随机数 + 12
  const total = articleTotal.value;
  const random = Math.floor(Math.random() * total) + 12;
  router.push({ path: "/article", query: { id: random } });
};

const httpInputBtn = async (httpInput: string): Promise<void> => {
  if (httpInput.length === 0) {
    notifyWarning("请输入有效的图片链接！", "链接不对🤣");
    return;
  }
  const status = await $common.isValidHttpUrl(httpInput);
  if (status) {
    changeBg(httpInput);
    notifySuccess("切换自定义背景成功");
  } else {
    notifyWarning("请输入有效的图片链接！", "链接不对🤣");
  }
};

const handleSakura = (): void => {
  if (sakuraFlag.value) {
    if (window.startSakura) {
      window.startSakura();
    }
    sakuraFlag.value = false;
  } else {
    const dom = document.querySelector("#canvas_sakura");
    sakuraFlag.value = true;
    if (dom) {
      dom.remove();
    }
  }
};

const logout = (): void => {
  notifySuccess("退出登录成功！");
  loadCurrentUser(null);
  router.push({ path: "/" });
};

const getWebInfo = async (): Promise<void> => {
  try {
    const res = await api.getWebInfo();
    loadWebInfo(res);
  } catch {
    /* ignored */
  }
};

const getSortInfo = async (): Promise<void> => {
  try {
    const res = await api.getSortInfo();
    loadSortInfo(res);
  } catch {
    /* ignored */
  }
};

// 关灯/开灯
const changeColor = (): void => {
  isLight.value = !isLight.value;
  localStorage.setItem("theme", String(isLight.value));
  // false是黑夜，显示月亮图标
  if (!isLight.value) {
    document.documentElement.dataset.theme = localStorage.getItem("themeColor")
      ? "theme2-dark"
      : "dark";
    if (document.documentElement.dataset.theme === "theme2-dark") {
      const root = document.querySelector(":root") as HTMLElement | null;
      if (root) {
        root.style.setProperty("--themeColor", "var(--sYellow)");
        localStorage.setItem("themeColor", "#ffc848");
        if (changeBgRef.value) {
          changeBgRef.value.handleSetColor("#ffc848");
        }
        $common.getThemeRgb();
      }
    }
    notifySuccess("是要关灯睡觉了吗~~~");
  } else {
    document.documentElement.dataset.theme = localStorage.getItem("themeColor")
      ? "theme2"
      : "light";
    if (document.documentElement.dataset.theme !== "theme2-dark") {
      const root = document.querySelector(":root") as HTMLElement | null;
      if (root) {
        root.style.setProperty("--themeColor", "var(--sBlue)");
        localStorage.setItem("themeColor", "#425aef");
        if (changeBgRef.value) {
          changeBgRef.value.handleSetColor("#425aef");
        }
        $common.getThemeRgb();
      }
    }
  }
};

const toTop = (): void => {
  window.scrollTo({
    top: 0,
    behavior: "smooth",
  });
};

const onScrollPage = (): void => {
  scrollTop.value = document.documentElement.scrollTop || document.body.scrollTop;
};

const changeMouseAnimation = (): void => {
  mouseAnimation.value = !mouseAnimation.value;
  if (mouseAnimation.value) {
    nextTick(() => {
      mousedown();
    });
  }
};

const openPcGame = (): void => {
  disGame.value = true;
  // 随机获取一个游戏
  let index = Math.floor(Math.random() * pcGameList.value.length);
  if (index === 0 || index === 1) {
    index = 6;
  }
  game.value = $constant.gameURL + index++;
};

const openChangeBg = (): void => {
  changeBgBox.value = true;
};

const changeBg = (item?: string): void => {
  // 刷新时触发并且没有本地缓存的背景，也没有点击切换背景
  const mainStorage = localStorage.getItem("main");
  if (!item && !(mainStorage ? JSON.parse(mainStorage)?.changeBg : false)) {
    const dom = document.querySelector(".background-image-changeBg");
    const storeInfo = changeBgState.value;
    const string = storeInfo.split("#");
    if (dom) {
      if (string[0] === "") {
        dom.setAttribute("style", `background-color: ${changeBgState.value}`);
      } else if (string.length === 1) {
        dom.setAttribute(
          "style",
          `background-image: ${changeBgState.value};background-size: cover;background-attachment: local;background-position: center;width: 100%;height: 100%;`
        );
      } else {
        dom.setAttribute("style", `background-image: ${changeBgState.value}`);
      }
    }
  } else if (item) {
    // 只有点击切换时触发
    const dom = document.querySelector(".background-image-changeBg");
    const string = item.split("#");
    // 纯色
    if (dom) {
      if (string[0] === "") {
        changeBgBoxMutation(item);
        dom.setAttribute("style", `background-color: ${changeBgState.value}`);
      } else if (string.length === 1) {
        // 图片
        const changeBgUrl = "url(" + item + ")";
        changeBgBoxMutation(changeBgUrl);
        dom.setAttribute(
          "style",
          `background-image: ${changeBgState.value};background-size: cover;background-attachment: local;background-position: center;width: 100%;height: 100%;`
        );
      } else {
        // 渐变
        const changeBgUrl = "linear-gradient(" + item + ")";
        changeBgBoxMutation(changeBgUrl);
        dom.setAttribute("style", `background-image: ${changeBgState.value}`);
      }
    }
  }
};

const defaultBtn = (): void => {
  themeActive.value = "#425aef";
  isLight.value = true;
  localStorage.setItem("theme", String(isLight.value));
  const switchBtnRef = document.querySelector("#switchBtnRef") as HTMLInputElement | null;
  if (switchBtnRef) {
    switchBtnRef.checked = false;
  }
  const root = document.querySelector(":root") as HTMLElement | null;
  if (root) {
    root.style.setProperty("--themeColor", "var(--sBlue)");
  }
  document.documentElement.dataset.theme = "theme2";
  localStorage.setItem("themeColor", "#425aef");
  const dom = document.querySelector(".background-image-changeBg");
  if (dom) {
    dom.setAttribute(
      "style",
      `background-image: url(${$constant.defaultBackground});background-size: cover;background-attachment: local;background-position: center;width: 100%;height: 100%;`
    );
  }
  changeBgBoxMutation(`url(${$constant.defaultBackground})`);
};

const handleChangeBg = async (val: string, i: number): Promise<void> => {
  if (val === "pc" && themeMap.value[i].dataList.length === 0) {
    getAllBg(i);
  }
  if (val === "mobile" && themeMap.value[i].dataList.length === 0) {
    try {
      const res = await api.getClientResourceList({
        resourceType: "mobilePhoto",
        noPagination: true,
      });
      (res.list || []).forEach((item) => {
        themeMap.value[i].dataList.push(item.path);
      });
    } catch {
      /* ignored */
    }
  }
  if (val === "gradient") {
    themeMap.value[i].dataList = $constant.gradient;
  }
  if (val === "solid") {
    themeMap.value[i].dataList = $constant.SolidColor;
  }
};

const getAllBg = async (i: number): Promise<void> => {
  try {
    const res = await api.getClientResourceList({
      resourceType: "webBackgroundImage",
      noPagination: true,
    });
    (res.list || []).forEach((item) => {
      themeMap.value[i].dataList.push(item.path);
    });
  } catch {
    /* ignored */
  }
};

const edit = (): void => {
  if (editFlag.value) {
    notifySuccess("已关闭编辑！！");
    document.body.contentEditable = "false";
    editFlag.value = false;
    return;
  }
  notifySuccess("可以随意修改本站的文字喔~~~不要干坏事！！");
  document.body.contentEditable = "true";
  editFlag.value = true;
};
</script>
<style lang="scss">
#main-container {
  min-width: 350px;
}

/* 导航栏 */
</style>
<style lang="scss" scoped>
.home-container {
  // 导航栏
  .toolbar-content {
    padding: 0 40px;
    top: 0;
    width: 100%;
    height: 50px;
    color: var(--red);
    position: fixed;
    z-index: 99999;
    user-select: none;
    transition: all 0.3s ease;

    &.enter {
      background: var(--background);
      color: var(--red);
      box-shadow: 0 1px 3px 0 var(--miniMask);
    }

    &.hoverEnter {
      background: var(--background);
      box-shadow: 0 1px 3px 0 var(--miniMask);
    }

    .toolbar-title {
      transition: 0.3s;
      height: 50px;

      &:hover {
        background-color: var(--red);
        color: var(--red);
        border-radius: 8px;

        &:after {
          opacity: 1;
          transform: translateY(0) scale(0.3);
          transition-timing-function: ease-in;
        }
      }

      &:after {
        background: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAMAAACdt4HsAAAABGdBTUEAALGPC/xhBQAAAAFzUkdCAK7OHOkAAABgUExURUxpcf///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////05o8T8AAAAfdFJOUwCqxd+3iU1VQL+Uz7IrG+mO5TN4EVrvRditZp8S9oBdLCSzAAABAUlEQVRYw+3WyRKCMAyA4bKjgLjv2vd/S5FBIF2T9OKBXMP/DTidsUI4J4/rXARMIbspAvsAYejZwtgzhVnPEkDPEJSeLGg9UTD0JMHYEwRLjxasPVJw9CjB2SMET+8VvL1HQPROAdU7BGRvFdC9RSD0RoHUGwRirwnkXhEYPRBY/Uxg9qPA7gchoO+FXMowQQaOiMP6WFx2IX1Sdr9i+R1622fTWWJ8PJwFWIAF+F+g2pbjHFI6kMH9gwpE6r9vRgT2KtAQgaN2gUhogH4DWS2AD4jgttWBGj6RKus7XK91QDlJJ3UPT/tTBzbukyrE+/pbntPGdJXd3NrhgVc1veEHPLUmePVAutQAAAAASUVORK5CYII=)
          no-repeat 50% 50%;
        opacity: 0;
        position: absolute;
        z-index: 1;
        width: 75px;
        height: 58px;
        content: "";
        transition: 0.3s;
        transition-timing-function: ease-in;
        transform: scale(0.4);
      }
    }

    .toolbar-mobile-menu {
      font-size: 30px;
      margin-right: 15px;
      height: 30px;
    }
  }

  // 返回顶部
  .cd-top {
    background: var(--toTop) no-repeat center;
    position: fixed;
    right: 5vh;
    top: -900px;
    z-index: 99;
    width: 70px;
    height: 900px;
    background-size: contain;
    transition: all 0.5s ease-in-out;
  }
}
</style>
