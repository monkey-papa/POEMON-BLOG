<template>
  <!-- 常规头部导航列表 -->
  <template v-if="!isMobile">
    <div class="center-toolbar screen-none">
      <ul class="scroll-menu">
        <!-- 文章 -->
        <li>
          <el-popover
            transition="el-zoom-in-center"
            :hide-after="300"
            popper-class="mk-popper"
            placement="bottom"
            trigger="hover"
          >
            <div class="mk-popover_item myBetween">
              <!-- 分类 -->
              <li @click="router.push({ path: '/sort' })">
                <img src="../../assets/svg/sort.svg" />
                分类
              </li>
              <!-- 标签 -->
              <li @click="router.push({ path: '/tags', query: { labelId: 25 } })">
                <img src="../../assets/svg/tag.svg" />
                标签
              </li>
            </div>
            <template #reference>
              <span class="menu-item-link">
                <img src="../../assets/svg/document.svg" />
                文章
              </span>
            </template>
          </el-popover>
        </li>
        <!-- 空间 -->
        <li>
          <el-popover
            transition="el-zoom-in-center"
            :hide-after="300"
            popper-class="mk-popper"
            placement="bottom"
            trigger="hover"
          >
            <div class="mk-popover_item myBetween">
              <!-- 音乐 -->
              <li @click="router.push({ path: '/funny' })">
                <img src="../../assets/svg/music.svg" />
                幻音坊
              </li>
              <!-- 藏宝阁 -->
              <li @click="router.push({ path: '/tools' })">
                <img src="../../assets/svg/treasure.svg" />
                藏宝阁
              </li>
            </div>
            <template #reference>
              <span class="menu-item-link">
                <img src="../../assets/svg/space.svg" />
                空间
              </span>
            </template>
          </el-popover>
        </li>
        <!-- 我的 -->
        <li>
          <el-popover
            transition="el-zoom-in-center"
            :hide-after="300"
            popper-class="mk-popper"
            placement="bottom"
            trigger="hover"
          >
            <div class="mk-popover_item myBetween">
              <!-- 相册 -->
              <li @click="router.push({ path: '/travel' })">📸 <span>相册</span></li>
              <!-- 爱链 -->
              <li @click="router.push({ path: '/love' })">
                <img src="../../assets/svg/love.svg" />
                爱链
              </li>
              <!-- 小游戏 -->
              <li @click="emit('openPcGame')">🎮 <span>小游戏</span></li>
              <!-- 编辑怪 -->
              <li @click="emit('edit')">
                <img src="../../assets/svg/pencil.svg" />
                编辑怪
              </li>
              <!-- 关于 -->
              <li @click="router.push({ path: '/about' })">🐷 <span>关于我</span></li>
            </div>
            <template #reference>
              <span class="menu-item-link">
                <img src="../../assets/svg/home.svg" />
                我的
              </span>
            </template>
          </el-popover>
        </li>
        <!-- 社交 -->
        <li>
          <el-popover
            transition="el-zoom-in-center"
            :hide-after="300"
            popper-class="mk-popper"
            placement="bottom"
            trigger="hover"
          >
            <div class="mk-popover_item myBetween">
              <!-- 留言厅 -->
              <li @click="router.push({ path: '/message' })">✍🏻 <span>留言厅</span></li>
              <!-- 友链 -->
              <li @click="router.push({ path: '/friend' })">🎀 <span>友链</span></li>
            </div>
            <template #reference>
              <span class="menu-item-link">
                <img src="../../assets/svg/socialContact.svg" />
                社交
              </span>
            </template>
          </el-popover>
        </li>
      </ul>
    </div>
    <div class="center-toolbar">
      <ul class="scroll-menu">
        <!-- 随机文章 -->
        <li
          class="menu-tool-item"
          title="随便看看-随机访问一篇本站文章"
          @click="emit('openRandomArticle')"
        >
          <div class="icon-container"><i class="fa fa-grav"></i></div>
        </li>
        <!-- 切换背景 -->
        <li
          class="menu-tool-item"
          title="切换背景-换一种背景，换一种感觉。"
          id="changeThemeRef"
          @click="emit('openChangeBg')"
        >
          <div class="icon-container"><i class="fa fa-image"></i></div>
        </li>
        <!-- 关闭樱花 -->
        <li class="menu-tool-item" @click="emit('handleSakura')">
          <div class="icon-container">
            <i class="fa fa-pagelines"></i>
          </div>
        </li>
        <!-- 关灯 -->
        <li class="menu-tool-item icon-container change-color-icon">
          <switchBtn id="changeColorRef" @click="emit('changeColor')" />
        </li>
        <!-- 个人中心 -->
        <li class="user-center">
          <el-dropdown placement="bottom" popper-class="user-center-dropdown">
            <el-avatar
              class="user-avatar"
              :size="36"
              :src="currentUser?.avatar ? currentUser.avatar : webInfo.avatar"
            />
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item
                  @click="router.push({ path: '/user' })"
                  v-if="currentUser"
                  class="user"
                >
                  <i class="fa fa-user-circle" aria-hidden="true"></i>
                  <span>个人</span>
                </el-dropdown-item>
                <el-dropdown-item @click="emit('logout')" v-if="currentUser" class="logout">
                  <i class="fa fa-sign-out" aria-hidden="true"></i>
                  <span>退出</span>
                </el-dropdown-item>
                <el-dropdown-item
                  @click="router.push({ path: '/user' })"
                  v-if="!currentUser"
                  class="login"
                >
                  <i class="fa fa-sign-in" aria-hidden="true"></i>
                  <span>登录</span>
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </li>
      </ul>
    </div>
  </template>
</template>

<script setup lang="ts">
import { defineAsyncComponent, computed } from "vue";
import { useRouter } from "vue-router";
import { useStore } from "@/stores";
import type { WebInfo } from "@/types";

const router = useRouter();
const store = useStore();

const currentUser = computed(() => store.currentUser);
const webInfo = computed<WebInfo>(() => store.webInfo);

const switchBtn = defineAsyncComponent(() => import("./switchBtn.vue"));

interface Props {
  isMobile?: boolean;
}

withDefaults(defineProps<Props>(), {
  isMobile: false,
});

interface Emits {
  openPcGame: [];
  edit: [];
  openRandomArticle: [];
  openChangeBg: [];
  handleSakura: [];
  changeColor: [];
  logout: [];
}

const emit = defineEmits<Emits>();
</script>

<style lang="scss" scoped>
.center-toolbar {
  .scroll-menu {
    color: var(--red1);
    display: flex;
    justify-content: flex-end;
    padding: 0;

    > li {
      list-style: none;
      margin-right: 24px;
      font-size: 17px;
      height: 50px;
      line-height: 50px;
      position: relative;
      text-decoration: none;
      background: var(--gradientAnimation);
      background-size: 0px 3px;
      transition: background-size 800ms;

      .menu-item-link {
        display: flex;
        align-items: center;

        img {
          vertical-align: -3px;
        }
      }

      &.menu-tool-item {
        color: var(--fontColor);
        font-size: 20px;

        &.change-color-icon {
          transform: scale(0.16);
          width: 70px;
          height: 0;
        }
      }

      &:hover {
        color: var(--orange2);
        background-position-x: left;
        background-size: 100% 3px;
      }

      &:last-child {
        margin-right: 0;
        display: flex;
        align-items: center;
      }
    }
  }
}

// 媒体查询
@media screen and (max-width: 950px) {
  .screen-none {
    display: none;
  }
}
</style>

<style lang="scss">
// 顶部工具栏下拉内容
.mk-popper {
  .mk-popover_item {
    li {
      list-style: none;
      padding: 8px;
      font-size: 16px;
      color: var(--red);

      img {
        vertical-align: -3px;
      }

      &:hover {
        background: var(--red);
        color: var(--white);
        padding: 8px;
        border-radius: 100px;
        animation: slide-top 0.5s ease-in-out both;
      }
    }
  }
}

.user-center-dropdown {
  font-size: unset !important;
  color: unset !important;
  background: unset !important;
  padding: unset !important;

  .el-dropdown-menu {
    align-items: center;
    border-radius: 13px;
    padding: 6px;
    border: 0;
    background: var(--blue);

    .el-dropdown-menu__item {
      font-size: unset;
      line-height: 26px;
      padding: 10px;
      color: var(--red);

      &:hover {
        border-radius: 8px;
        animation: slide-top 0.5s ease-in-out both;
        background-color: var(--red);
        color: var(--white);
      }

      &.user,
      &.login,
      &.logout {
        &:hover {
          color: var(--white);
        }
      }
    }

    &.active {
      display: flex;
    }
  }
}
</style>
