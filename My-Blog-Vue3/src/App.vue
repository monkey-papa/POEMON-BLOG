<template>
  <div id="app" @contextmenu.prevent="openMenu($event)">
    <router-view />
    <!-- 页脚 -->
    <div
      class="sf-my__footer"
      v-if="!['/user', '/love', '/verifyLogin'].includes($route.path) && !$route.meta.requiresAuth"
    >
      <myFooter />
    </div>
    <aplayer :class="{ 'aplayer-hidden': !!$route.meta.requiresAuth }" />
    <transition name="loading-fade">
      <div v-if="$store.isShowLoading" class="loading myCenter">
        <div class="author-box myCenter">
          <span></span>
          <div class="author-img">
            <img :src="avatarImg" alt="站点头像" />
          </div>
        </div>
        <div class="image-dot"></div>
      </div>
    </transition>
    <!-- 右键菜单部分 -->
    <ul v-show="visible" :style="{ left: left + 'px', top: top + 'px' }" class="contextmenu">
      <div class="rightMenu-group">
        <div class="rightMenu-item">
          <i class="fa fa-arrow-left" @click="backAndForward('2')"></i>
        </div>
        <div class="rightMenu-item">
          <i class="fa fa-arrow-right" @click="backAndForward('1')"></i>
        </div>
        <div class="rightMenu-item">
          <i class="fa fa-repeat" @click="refresh"></i>
        </div>
        <div class="rightMenu-item">
          <i class="fa fa-music" @click="musicHandle"></i>
        </div>
      </div>
      <div class="rightMenu-group rightMenu-line">
        <p @click="router.push('/')" class="rightMenu-item">
          <i class="fa fa-home"></i>
          <span>博客首页</span>
        </p>
        <p @click="dayAndNight" class="rightMenu-item">
          <i class="fa fa-adjust"></i>
          <span>昼夜切换</span>
        </p>
        <p @click="onCopy" class="rightMenu-item">
          <i class="fa fa-code-fork"></i>
          <span>加入我们</span>
        </p>
        <p @click="changeTheme" class="rightMenu-item">
          <i class="fa fa-image"></i>
          <span>美化设置</span>
        </p>
        <p @click="router.push('/tags?labelId=15')" class="rightMenu-item">
          <i class="fa fa-bookmark"></i>
          <span>标签</span>
        </p>
        <p @click="router.push('/sort?sortId=1')" class="rightMenu-item">
          <i class="fa fa-folder-open"></i>
          <span>分类</span>
        </p>
      </div>
    </ul>
  </div>
</template>
<script setup lang="ts">
import {
  ref,
  watch,
  onMounted,
  onBeforeUnmount,
  nextTick,
  defineAsyncComponent,
  computed,
} from "vue";
import { useRouter } from "vue-router";
import avatarImg from "./assets/file/avatar.jpg";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
import { notifySuccess } from "@/utils/notify";
import constant from "@/utils/constant";
import api from "@/api";

const myFooter = defineAsyncComponent(() => import("./views/common/footer.vue"));
const aplayer = defineAsyncComponent(() => import("./views/common/aplayer.vue"));

const { $store, $route, $common } = useGlobalProperties();
const router = useRouter();

// 获取用户和位置信息
const currentUser = computed(() => $store.currentUser);
const currentAdmin = computed(() => $store.currentAdmin);
const locationInfo = computed(() => $store.locationInfo);

// IP 记录逻辑（会话级防重复，用户ID由服务端从Token自动识别）
const SESSION_KEY = "ip_recorded";

const submitIPRecord = async (): Promise<void> => {
  if (!locationInfo.value.isLoaded) return;
  if (!locationInfo.value.city && !locationInfo.value.province) return;

  // 获取当前用户标识，用于区分是否需要重新提交
  const userId: number = currentUser.value?.id ?? currentAdmin.value?.id ?? 0;
  const recordKey = String(userId);

  // 如果已以当前身份提交过，不重复提交
  if (sessionStorage.getItem(SESSION_KEY) === recordKey) return;

  try {
    await api.submitLocation({
      province: locationInfo.value.province,
      city: locationInfo.value.city,
    });
    sessionStorage.setItem(SESSION_KEY, recordKey);
  } catch {
    /* ignored */
  }
};

// 后台页面隐藏看板娘（外部 live2d-widget 创建的 #waifu 元素）
watch(
  () => $route.meta.requiresAuth,
  (isAdmin) => {
    const waifu = document.getElementById("waifu");
    if (waifu) {
      waifu.style.display = isAdmin ? "none" : "";
    }
  },
  { immediate: true }
);

// 监听位置信息加载完成
watch(
  () => locationInfo.value.isLoaded,
  (isLoaded) => {
    if (isLoaded) {
      submitIPRecord();
    }
  },
  { immediate: true }
);

// 监听用户/管理员登录状态变化（登录后重新提交，服务端自动关联用户）
watch(
  () => currentUser.value?.id ?? currentAdmin.value?.id,
  (newId) => {
    if (newId && locationInfo.value.isLoaded) {
      submitIPRecord();
    }
  },
  { immediate: true }
);

const scrollProgress = ref<number>(0);
const visible = ref<boolean>(false); // 是否展示右键菜单
const top = ref<number>(0);
const left = ref<number>(0);
const copyContent = ref<string>(constant.wechatId);

// 初始化主题
if (window.CSS && window.CSS.registerProperty) {
  window.CSS.registerProperty({
    name: "--p1",
    syntax: "<percentage>",
    inherits: false,
    initialValue: "0%",
  });
}

// 黑夜
if (localStorage.getItem("theme") == "false") {
  document.documentElement.dataset.theme = localStorage.getItem("themeColor")
    ? "theme2-dark"
    : "dark";
  if (document.documentElement.dataset.theme === "theme2-dark") {
    const root = document.querySelector(":root") as HTMLElement | null;
    if (root) {
      root.style.setProperty("--themeColor", localStorage.getItem("themeColor"));
      $common.getThemeRgb();
    }
  }
} else {
  if (localStorage.getItem("themeColor")) {
    const root = document.querySelector(":root") as HTMLElement | null;
    if (root) {
      root.style.setProperty("--themeColor", localStorage.getItem("themeColor"));
      $common.getThemeRgb();
    }
    document.documentElement.dataset.theme = "theme2";
  } else {
    document.documentElement.dataset.theme = "light";
  }
}

watch(
  () => $store.isShowLoading,
  (value) => {
    if (value) {
      document.body.style.overflow = "hidden";
    } else {
      document.body.style.overflow = "auto";
    }
  }
);

// 监听 visible，来触发关闭右键菜单，调用关闭菜单的方法
watch(visible, (value) => {
  if (value) {
    document.body.addEventListener("click", closeMenu);
  } else {
    document.body.removeEventListener("click", closeMenu);
  }
});

onMounted(() => {
  window.addEventListener("scroll", handleScroll, { passive: true }); // 监听滚动条事件
  window.addEventListener("beforeunload", handleBeforeUnload);
  const styleTitle1 = `
font-size: 20px;
font-weight: 600;
color: rgb(244,167,89);
`;
  const styleTitle2 = `
font-size:12px;
color: #425AEF;
`;
  const styleContent = `
color: rgb(30,152,255);
`;
  const title1 = "ZJHの主页 被我发现了吧，既然你已经破解了，转发、拿东西记得标明出处喔~~";
  const title2 = `
    く__,.ヘヽ.        /  ,ー､ 〉
           ＼ ', !-─‐-i  /  /´
           ／｀ｰ'       L/／｀ヽ､
         /   ／,   /|   ,   ,       ',
       ｲ   / /-‐/  ｉ  L_ ﾊ ヽ!   i
        ﾚ ﾍ 7ｲ｀ﾄ   ﾚ'ｧ-ﾄ､!ハ|   |
          !,/7 '0'     ´0iソ|    |
          |.从"    _     ,,,, / |./    |
          ﾚ'| i＞.､,,__  _,.イ /   .i   |
            ﾚ'| | / k_７_/ﾚ'ヽ,  ﾊ.  |
              | |/i 〈|/   i  ,.ﾍ |  i  |
             .|/ /  ｉ：    ﾍ!    ＼  |
              kヽ>､ﾊ    _,.ﾍ､    /､!
              !'〈//｀Ｔ´', ＼ ｀'7'ｰr'
              ﾚ'ヽL__|___i,___,ンﾚ|ノ
                  ﾄ-,/  |___./
                  'ｰ'    !_,.:
██████╗ ██╗   ██╗║██████╗ 
██╔══██╗╚██╗ ██╔╝║██╔═══╗
██████╔╝ ╚████╔╝ ║██████║
██╔══██╗  ╚██╔╝  ║██╔═══╗
██████╔╝   ██║   ║██████║
╚═════╝    ╚═╝   ╚══════╝(wx:${constant.wechatId})OVO
`;
  const content = `
誰もが信じ崇めてる
まさに最強で無敵のアイドル
弱点なんて見当たらない
一番星を宿している
弱いとこなんて見せちゃダメダメ
知りたくないとこは見せずに
唯一無二じゃなくちゃイヤイヤ
それこそ本物のアイ
  ⌜IDOL⌟
`;
  console.log(
    `%c${title1} %c${title2}
%c${content}`,
    styleTitle1,
    styleTitle2,
    styleContent
  );
  nextTick(() => {
    // 禁止右键
    document.oncontextmenu = (): boolean => false;
  });
});

onBeforeUnmount(() => {
  window.removeEventListener("beforeunload", handleBeforeUnload);
  window.removeEventListener("scroll", handleScroll);
});

// 打开右键菜单
const openMenu = (e: MouseEvent): void => {
  visible.value = true;
  top.value = e.pageY;
  left.value = e.pageX;
};

// 关闭右键菜单
const closeMenu = (): void => {
  visible.value = false;
};

// 音乐跳转
const musicHandle = (): void => {
  window.open(constant.siteURL);
};

// 路由跳转
const backAndForward = (val: string): void => {
  if (val === "1") {
    window.history.forward();
  } else {
    window.history.back();
  }
};

// 刷新
const refresh = (): void => {
  location.reload();
};

// 昼夜切换
const dayAndNight = (): void => {
  const changeColorRef = document.getElementById("changeColorRef");
  const switchBtnRef = document.getElementById("switchBtnRef") as HTMLInputElement | null;
  if (changeColorRef) {
    changeColorRef.click();
  }
  if (switchBtnRef) {
    switchBtnRef.checked = !switchBtnRef.checked;
  }
};

// 加入我们（使用现代 Clipboard API，兼容旧浏览器）
const onCopy = async (): Promise<void> => {
  try {
    await navigator.clipboard.writeText(copyContent.value);
  } catch {
    const input = document.createElement("input");
    input.value = copyContent.value;
    document.body.appendChild(input);
    input.select();
    document.execCommand("Copy");
    document.body.removeChild(input);
  }
  notifySuccess("本博主的微信已经到你的剪贴板啦，快加入我们吧~~🎉");
};

// 美化设置
const changeTheme = (): void => {
  const changeThemeRef = document.getElementById("changeThemeRef");
  if (changeThemeRef) {
    changeThemeRef.click();
  }
};

const handleScroll = (): void => {
  // 屏幕剩余的高度
  const surplus = document.documentElement.scrollHeight - document.documentElement.clientHeight;
  // 当前滑动高度
  let scrollY = document.documentElement.scrollTop;
  if (scrollY < 0) {
    scrollY = 0;
  }
  if (scrollY > 0) {
    visible.value = false;
  }
  // 当前位置百分比小数
  scrollProgress.value = scrollY / surplus;
  $store.setTopPercentage(Math.floor(scrollProgress.value * 100));
  // 设置导航栏，这里使用NProgress.set() 动态更改进度条
  if (window.NProgress) {
    window.NProgress.set(scrollProgress.value);
  }
};

const handleBeforeUnload = (): void => {
  $store.setTopPercentage(0);
};
</script>
<style lang="scss">
#nprogress {
  .bar {
    background: linear-gradient(to right, var(--green1), var(--red)) no-repeat !important;
    height: 5px !important;
  }
  .peg {
    box-shadow: 0 0 10px var(--transparent), 0 0 5px var(--transparent) !important;
  }
}

// 后台隐藏音乐播放器（不使用 display:none，避免 APlayer 固定定位模式状态丢失）
.aplayer-hidden {
  position: fixed !important;
  bottom: -200px !important;
  pointer-events: none !important;
  opacity: 0 !important;
  z-index: -1 !important;
}
</style>
<style lang="scss" scoped>
#app {
  .contextmenu {
    margin: 0;
    background: var(--white);
    z-index: 3000;
    position: absolute;
    width: 9rem;
    height: fit-content;
    border-radius: 12px;
    border: 1px solid var(--favoriteBg);
    font-size: 12px;
    font-weight: 700;
    color: var(--black5);
    transition: 0.3s ease;
    padding: 0 0.25rem;

    .rightMenu-group {
      padding: 0.35rem 0.3rem;
      display: flex;
      justify-content: space-between;

      &:not(:nth-last-child(1)) {
        border-bottom: 1px dashed var(--miniMask);
      }

      &.rightMenu-line {
        display: block;

        .rightMenu-item {
          margin: 0.25rem 0;
          padding: 0.25rem 0;
          display: flex;
          font-size: 15px;

          i {
            margin: 0 0.25rem;
          }

          span {
            line-height: 1.5rem;
            font-weight: 500;
          }
        }
      }

      .rightMenu-item {
        border-radius: 8px;
        transition: 0.3s ease;

        i {
          font-size: 15px;
          display: inline-block;
          text-align: center;
          line-height: 1.5rem;
          font-weight: 900;
          width: 1.5rem;
          padding: 0 0.25rem;
        }

        &:hover {
          background-color: var(--blue25);
          color: var(--white);
          box-shadow: 0 8px 12px -3px var(--miniMask);
        }
      }
    }
  }

  .loading {
    width: 100vw;
    height: 100vh;
    background: linear-gradient(55deg, var(--blue1) 20%, var(--green6) 100%);
    position: fixed; // 改为 fixed，确保覆盖整个视口
    top: 0px;
    left: 0px;
    z-index: 1000000; // 提高 z-index，确保在 HTML loading-box (999999) 之上
    font-size: 30px;

    .author-box {
      position: relative;
      width: 159px;
      height: 159px;
      border-radius: 50%;
      overflow: hidden;

      &::before {
        content: "";
        position: absolute;
        width: 500px;
        height: 500px;
        background-image: conic-gradient(transparent, transparent, transparent, var(--purple1));
        animation: rotate 2s linear infinite;
        animation-delay: -1s;
      }

      &::after {
        content: "";
        position: absolute;
        width: 500px;
        height: 500px;
        background-image: conic-gradient(transparent, transparent, transparent, var(--blue2));
        animation: rotate 2s linear infinite;
      }

      span {
        position: absolute;
        inset: 5px;
        border-radius: 50%;
        background: var(--favoriteBg);
        z-index: 1;
      }

      .author-img {
        margin: auto;
        border-radius: 50%;
        overflow: hidden;
        width: 150px;
        height: 150px;
        z-index: 10;
        background: var(--maxMaxWhiteMask);

        img {
          border-radius: 11px;
          margin-right: 4px;
          display: block;
          margin: 0 auto 20px;
          max-width: 100%;
          animation: breath 700ms ease-in-out infinite;
        }
      }
    }

    .image-dot {
      width: 25px;
      height: 25px;
      background: var(--green1);
      position: absolute;
      border-radius: 50%;
      border: 4px solid var(--favoriteBg);
      z-index: 20;
      transform: translate(51px, 54px);
    }
  }
}

@keyframes breath {
  from {
    opacity: 0.7;
  }
  25% {
    opacity: 0.9;
  }
  50% {
    opacity: 1;
  }
  75% {
    opacity: 0.9;
  }
  to {
    opacity: 0.7;
  }
}

// Loading 过渡动画 - 优化为无过渡，避免闪烁
// 因为 HTML loading-box 和 Vue loading 需要无缝衔接
.loading-fade-enter-active {
  transition: none; // 进入时不使用过渡，立即显示
}

.loading-fade-leave-active {
  transition: opacity 0.3s ease; // 离开时使用过渡，平滑消失
}

.loading-fade-enter-from {
  opacity: 1; // 进入时完全不透明，立即显示
}

.loading-fade-leave-to {
  opacity: 0;
}
</style>
