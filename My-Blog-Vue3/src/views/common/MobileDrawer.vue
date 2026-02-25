<template>
  <!-- 移动端抽屉 -->
  <el-drawer
    v-model="drawerVisible"
    :show-close="false"
    :close-on-click-modal="true"
    size="300px"
    class="toolbarDrawer"
    direction="ltr"
  >
    <div class="backdrop-color">
      <!-- 信息 -->
      <div class="sidebar is-center">
        <div class="avatar-img">
          <img :src="webInfo.avatar" />
        </div>
        <div class="author-info_name">『{{ webInfo.webName }}』</div>
        <div class="author-info__description">让我再享受一下</div>
      </div>
      <!-- 分类 -->
      <div class="site-data is-center">
        <div class="blog-info-box mySpaceAround">
          <span>文章</span>
          <span class="blog-info-num">{{ articleTotal }}</span>
        </div>
        <div class="blog-info-box mySpaceAround">
          <span>标签</span>
          <span class="blog-info-num">{{ labelInfo.length }}</span>
        </div>
        <div class="blog-info-box mySpaceAround">
          <span>分类</span>
          <span class="blog-info-num">{{ sortInfo.length }}</span>
        </div>
      </div>
      <hr />
      <!-- 菜单 -->
      <div>
        <div class="small-menu">
          <!-- 文章 -->
          <div>
            <span class="menu-item-link">
              <img src="../../assets/svg/document.svg" />
              <span>文章</span>
            </span>
            <div class="menu_item">
              <!-- 分类 -->
              <span @click="handleMenuClick({ path: '/sort' })">
                <div class="my-menu">
                  <img src="../../assets/svg/sort.svg" />
                  分类
                </div>
              </span>
              <!-- 标签 -->
              <span @click="handleMenuClick({ path: '/tags', query: { labelId: 25 } })">
                <div class="my-menu">
                  <img src="../../assets/svg/tag.svg" />
                  标签
                </div>
              </span>
            </div>
          </div>
          <!-- 空间 -->
          <div>
            <span class="menu-item-link">
              <img src="../../assets/svg/space.svg" />
              <span>空间</span>
            </span>
            <div class="menu_item">
              <!-- 音乐 -->
              <span @click="handleMenuClick({ path: '/funny' })">
                <div class="my-menu">
                  <img src="../../assets/svg/music.svg" />
                  幻音坊
                </div>
              </span>
              <!-- 藏宝阁 -->
              <span @click="handleMenuClick({ path: '/tools' })">
                <div class="my-menu">
                  <img src="../../assets/svg/treasure.svg" />
                  藏宝阁
                </div>
              </span>
            </div>
          </div>
          <!-- 我的 -->
          <div>
            <span class="menu-item-link">
              <img src="../../assets/svg/home.svg" />
              <span>我的</span>
            </span>
            <div class="menu_item">
              <!-- 相册 -->
              <span @click="handleMenuClick({ path: '/travel' })">
                <div class="my-menu">📸 <span>相册</span></div>
              </span>
              <!-- 爱链 -->
              <span @click="handleMenuClick({ path: '/love' })">
                <div class="my-menu">
                  <img src="../../assets/svg/love.svg" />
                  爱链
                </div>
              </span>
              <span @click="handleOpenPcGame">
                <div class="my-menu">🎮 <span>小游戏</span></div>
              </span>
              <span @click="handleEdit">
                <div class="my-menu">
                  <img src="../../assets/svg/pencil.svg" />
                  编辑怪
                </div>
              </span>
              <!-- 关于 -->
              <span @click="handleMenuClick({ path: '/about' })">
                <div class="my-menu">🐷 <span>关于我</span></div>
              </span>
            </div>
          </div>
          <!-- 社交 -->
          <div>
            <span class="menu-item-link">
              <img src="../../assets/svg/socialContact.svg" />
              <span>社交</span>
            </span>
            <div class="menu_item">
              <!-- 留言厅 -->
              <span @click="handleMenuClick({ path: '/message' })">
                <div class="my-menu">✍🏻 <span>留言厅</span></div>
              </span>
              <!-- 友链 -->
              <span @click="handleMenuClick({ path: '/friend' })">
                <div class="my-menu">🎀 <span>友链</span></div>
              </span>
            </div>
          </div>
          <template v-if="!currentUser">
            <li @click="handleMenuClick({ path: '/user' })">
              <div>
                <i class="fa fa-sign-in" aria-hidden="true"></i>
                <span>&nbsp;登录</span>
              </div>
            </li>
          </template>
          <template v-else>
            <li @click="handleMenuClick({ path: '/user' })">
              <div>
                <i class="fa fa-user-circle" aria-hidden="true"></i>
                <span>&nbsp;个人</span>
              </div>
            </li>
            <li @click="handleLogout">
              <div>
                <i class="fa fa-sign-out" aria-hidden="true"></i>
                <span>&nbsp;退出</span>
              </div>
            </li>
          </template>
          <!-- 随机文章 -->
          <li class="menu-tool-item" @click="handleOpenRandomArticle">
            <div class="my-menu"><i class="fa fa-grav"></i></div>
          </li>
          <!-- 切换背景 -->
          <li class="menu-tool-item" @click="handleOpenChangeBg">
            <div class="my-menu"><i class="fa fa-image"></i></div>
          </li>
          <!-- 关灯 -->
          <li class="menu-tool-item" @click="handleChangeColor">
            <div class="my-menu"><i class="fa-adjust fa"></i></div>
          </li>
          <!-- 关闭樱花 -->
          <li class="menu-tool-item" @click="handleSakura">
            <div class="my-menu"><i class="fa fa-pagelines"></i></div>
          </li>
        </div>
      </div>
    </div>
  </el-drawer>
</template>

<script setup lang="ts">
import { computed } from "vue";
import { useStore } from "@/stores";
import type { WebInfo, SortInfo, LabelInfo } from "@/types";
import type { RouteLocationRaw } from "vue-router";
import { useRouter } from "vue-router";

const store = useStore();
const router = useRouter();

interface Props {
  modelValue?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  modelValue: false,
});

interface Emits {
  "update:modelValue": [value: boolean];
  menuClick: [data: RouteLocationRaw];
  openPcGame: [];
  edit: [];
  openRandomArticle: [];
  openChangeBg: [];
  changeColor: [];
  handleSakura: [];
  logout: [];
}

const emit = defineEmits<Emits>();

const webInfo = computed<WebInfo>(() => store.webInfo);
const sortInfo = computed<SortInfo[]>(() => store.sortInfo);
const currentUser = computed(() => store.currentUser);
const articleTotal = computed<number>(() => store.totalArticles);
const labelInfo = computed<LabelInfo[]>(() => store.labelInfo);

const drawerVisible = computed<boolean>({
  get(): boolean {
    return props.modelValue;
  },
  set(val: boolean): void {
    emit("update:modelValue", val);
  },
});

const handleMenuClick = (data: RouteLocationRaw): void => {
  router.push(data);
  emit("update:modelValue", false);
  emit("menuClick", data);
};

const handleOpenPcGame = (): void => {
  emit("openPcGame");
  emit("update:modelValue", false);
};

const handleEdit = (): void => {
  emit("edit");
  emit("update:modelValue", false);
};

const handleOpenRandomArticle = (): void => {
  emit("openRandomArticle");
};

const handleOpenChangeBg = (): void => {
  emit("openChangeBg");
  emit("update:modelValue", false);
};

const handleChangeColor = (): void => {
  emit("changeColor");
};

const handleSakura = (): void => {
  emit("handleSakura");
};

const handleLogout = (): void => {
  emit("logout");
  emit("update:modelValue", false);
};
</script>

<style lang="scss">
.toolbarDrawer {
  position: relative;
  letter-spacing: 3px;

  .el-drawer__body {
    padding: 0;
    background: linear-gradient(60deg, var(--pink1) 20%, var(--yellow2) 93%);
    z-index: 999;
  }

  .el-drawer__header {
    padding: 0;
    margin: 0;
  }

  .sidebar {
    background-image: var(--toolbar);
    background-position: top;
    background-size: 120%;
    background-repeat: no-repeat;

    .avatar-img {
      overflow: hidden;
      margin: 0 auto;
      width: 110px;
      height: 110px;
      border-radius: 25px;
      box-shadow: 2.2px 2.2px 2.2px var(--toolbarBackground);

      img {
        width: 100%;
        height: 100%;
        object-fit: cover;
      }
    }

    .author-info_name {
      margin-top: 10px;
      color: var(--blue);
      font-weight: 800;
      font-size: 25.2px;
    }

    .author-info__description {
      color: var(--blue);
      margin-top: 20px;
      font-size: 18px;
      padding-bottom: 8px;
    }
  }

  .is-center {
    text-align: center;
  }

  .site-data {
    margin-top: 20px;
    color: var(--blue);
    width: 100%;
    display: flex;
    flex-direction: row;
    justify-content: space-around;

    .blog-info-box {
      flex-direction: column;

      .blog-info-num {
        margin-top: 12px;
      }
    }
  }

  hr {
    position: relative;
    margin: 20px auto 10px;
    border: 2px dashed var(--blue);
    overflow: visible;

    &:before {
      position: absolute;
      top: -21px;
      left: 5%;
      color: var(--red);
      content: "\e673";
      font-size: 40px;
      line-height: 1;
      transition: all 1s ease-in-out;
      font-family: iconfont;
    }

    &:hover:before {
      left: calc(95% - 20px);
    }
  }

  .small-menu {
    padding: 15px;
    margin: 0;
    color: var(--blue);
    font-size: 20px;
    user-select: none;
    position: relative;

    > li {
      padding-top: 5px;
      padding-bottom: 5px;
      list-style: none;
      height: 45px;
      line-height: 35px;

      &.menu-tool-item {
        color: var(--black);
        font-size: 20px;
      }

      &:hover {
        border-radius: 8px;
        padding-left: 10px;
        background: var(--red);
      }
    }

    .menu-item-link {
      display: flex;
      align-items: center;
      font-size: 14px;

      img {
        width: 14px;
        height: 14px;
        vertical-align: -3px;
      }

      span {
        color: var(--red);
        margin-left: 6px;
      }
    }

    .menu_item {
      flex-wrap: wrap;
      margin: 0;
      width: 100%;
      display: flex;
      justify-content: space-between;
      flex-direction: row;
      margin-top: 8px;

      .my-menu {
        font-size: 15px;
        line-height: 27px;

        img {
          width: 15px;
          height: 15px;
          vertical-align: -3px;
        }
      }

      > span {
        width: calc(50% - 6px);
        margin: 0;
        border: 1px solid var(--blue);
        border-radius: 8px;
        padding: 4px 15px;
        margin-bottom: 8px;
        transition: all 0.3s;

        &:hover {
          border: 1px solid var(--red);
          background: var(--red);
          color: var(--white);
        }
      }
    }
  }
}
</style>
