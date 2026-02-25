<template>
  <div class="index-container">
    <!-- 首页图片 -->
    <div
      style="animation: header-effect 2s"
      :style="{ background: `${changeBg}` }"
      class="background-image background-image-changeBg"
    ></div>
    <!-- 首页文字 -->
    <div class="signature-wall myCenter my-animation-hideToShow">
      <h1 class="playful">
        <span v-for="(a, index) in webInfo.webTitle" :key="index">{{ a }}</span>
      </h1>
      <div class="printer" @click="getGuShi()">
        <printer :printerInfo="printerInfo">
          <template v-slot:paper="scope">
            <h3>{{ scope.content }}<span class="cursor">|</span></h3>
          </template>
        </printer>
      </div>
      <section class="main-hero-waves-area waves-area">
        <svg
          class="waves-svg"
          xmlns="http://www.w3.org/2000/svg"
          xlink="http://www.w3.org/1999/xlink"
          viewBox="0 24 150 28"
          preserveAspectRatio="none"
          shape-rendering="auto"
        >
          <defs>
            <path
              id="gentle-wave"
              d="M -160 44 c 30 0 58 -18 88 -18 s 58 18 88 18 s 58 -18 88 -18 s 58 18 88 18 v 44 h -352 Z"
            />
          </defs>
          <g class="parallax">
            <use href="#gentle-wave" x="48" y="0" />
            <use href="#gentle-wave" x="48" y="3" />
            <use href="#gentle-wave" x="48" y="5" />
            <use href="#gentle-wave" x="48" y="7" />
          </g>
        </svg>
      </section>
      <el-icon :size="40" class="icon-arrow-down" @click="navigation('.page-container-wrap')"
        ><ArrowDownBold
      /></el-icon>
    </div>
    <!-- 首页内容 -->
    <div class="page-container-wrap">
      <div class="page-container">
        <!-- 内容页面 -->
        <div class="recent-posts">
          <!-- 提示框 -->
          <div class="announcement">
            <el-icon><Opportunity /></el-icon>
            <div class="notice-wrap">
              <div class="notice-item" v-for="(notice, index) in webInfo.notices" :key="index">
                {{ notice }}
              </div>
            </div>
          </div>
          <!-- 滚动栏 -->
          <div>
            <div class="content">
              <div class="category_group">
                <div class="category_item item">
                  <a :href="$constant.siteURL" class="category_button category_buttons">
                    <span class="category_button_text">云音乐</span>
                  </a>
                </div>
                <div @click="router.push({ path: '/map' })" class="category_item">
                  <a href="javascript:;" class="category_button">
                    <span class="category_button_text">四海之内</span>
                  </a>
                </div>
              </div>
              <div
                v-for="(article, index) in recommendArticles"
                :key="index"
                @click="router.push({ path: '/article', query: { id: article.id } })"
                class="page"
              >
                <div class="post_cover">
                  <a href="javascript:;">
                    <img :src="article.articleCover || ''" />
                  </a>
                </div>
                <div class="post_info">
                  <span class="article-title">{{ article.articleTitle }}</span>
                </div>
              </div>
            </div>
          </div>
          <!-- 侧边栏 -->
          <div class="aside-content" v-if="$common.mobile()">
            <myAside @selectSort="selectSort" />
          </div>
          <!-- 文章列表 -->
          <articleList :articleList="articles" :parentLoadingMark="parentLoadingMark" />
          <!-- 底部 -->
          <div class="pagination-wrap">
            <div
              @click="pageArticles()"
              class="pagination"
              v-if="pagination.total !== articles.length"
            >
              下一页
            </div>
            <div v-else class="pagination-text">~~( •̀ ω •́ )y 到底啦~~</div>
          </div>
        </div>
        <!-- 侧边栏 -->
        <div class="aside-content" v-if="!$common.mobile()">
          <myAside @selectSort="selectSort" />
        </div>
      </div>
      <!-- live2d -->
      <!-- <live2d
        class="live2d"
        :style="style"
        :model="['Potion-Maker/Pio', 'school-2017-costume-yellow']"
        :direction="direction"
        :size="$common.mobile() ? 140 : size"
      /> -->
    </div>
  </div>
</template>
<script setup lang="ts">
// 在组件中引入
// import live2d from "vue-live2d";
import api from "@/api";
import { ref, onMounted, onBeforeUnmount, nextTick, defineAsyncComponent, computed } from "vue";
import { onBeforeRouteLeave, useRouter } from "vue-router";
import { useStore } from "@/stores";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
import type { ArticleListItem, NewArticle } from "@/types/article";

const printer = defineAsyncComponent(() => import("./common/printer.vue"));
const articleList = defineAsyncComponent(() => import("./articleList.vue"));
const myAside = defineAsyncComponent(() => import("./myAside.vue"));

const store = useStore();
const { $common, $constant } = useGlobalProperties();
const router = useRouter();

const changeBg = computed(() => store.changeBg);
const webInfo = computed(() => store.webInfo);

const setArticleTotal = (value: number): void => store.setArticleTotal(value);
const setNewArticles = (value: NewArticle[]): void => store.setNewArticles(value);

interface Pagination {
  current: number;
  size: number;
  total: number;
  searchKey: string;
  sortId: number | undefined;
  deltaY?: number;
}

interface GuShi {
  content: string;
  origin: string;
  author: string;
  category: string;
}

const printerInfo = ref<string>("咦，这个地方是用来干什么的？");
const pagination = ref<Pagination>({
  current: 1,
  size: 10,
  total: 0,
  searchKey: "",
  sortId: undefined,
});
const guShi = ref<GuShi>({
  content: "",
  origin: "",
  author: "",
  category: "",
});
const articles = ref<ArticleListItem[]>([]);
const recommendArticles = ref<ArticleListItem[]>([]);
// const direction = ref<string>("right");
// const style = ref<string>("");
// const size = ref<number>(210);
const parentLoadingMark = ref<boolean>(false);

onBeforeRouteLeave((to): void => {
  if (to.path === "/") {
    parentLoadingMark.value = true;
  }
});

onMounted((): void => {
  getGuShi();
  getArticles();
  scrollTo();
  getRecommendArticles();
});

const selectSort = async (sort: { id: number }): Promise<void> => {
  pagination.value = {
    current: 1,
    size: 10,
    total: 0,
    searchKey: "",
    sortId: sort.id,
    deltaY: 0,
  };
  articles.value = [];
  await getArticles();
  nextTick(() => {
    const recentPosts = document.querySelector(".recent-posts");
    if (recentPosts) {
      recentPosts.scrollIntoView({
        behavior: "smooth",
        block: "start",
        inline: "nearest",
      });
    }
  });
};

const pageArticles = (): void => {
  pagination.value.current = pagination.value.current + 1;
  getArticles();
};

const getArticles = async (): Promise<void> => {
  try {
    const res = await api.getArticleList(pagination.value);
    // 使用扩展运算符创建新数组，避免响应式警告
    articles.value = [...articles.value, ...(res.list || [])];
    pagination.value.total = res.totalCount || 0;
    setArticleTotal(res.totalCount || 0);
    setNewArticles(articles.value.slice(0, 4));
  } catch {
    /* ignored */
  }
};

const getRecommendArticles = async (): Promise<void> => {
  const recommendPagination = {
    current: 1,
    size: 10,
    recommendStatus: true,
  };

  try {
    const res = await api.getArticleList(recommendPagination);
    const formatRecommendArticles = res.list || [];
    const myBlog = formatRecommendArticles.filter((item: ArticleListItem) => item.id === 20);
    const notMyBlog = formatRecommendArticles.filter((item: ArticleListItem) => item.id !== 20);
    const recommendArticlesList = myBlog.concat(notMyBlog);
    recommendArticles.value = recommendArticlesList;
  } catch {
    /* ignored */
  }
};

const navigation = (selector: string): void => {
  const pageId = document.querySelector(selector);
  if (pageId) {
    window.scrollTo({
      top: (pageId as HTMLElement).offsetTop,
      behavior: "smooth",
    });
  }
};

const getGuShi = (): void => {
  const xhr = new XMLHttpRequest();
  xhr.open("get", $constant.jinrishici);
  xhr.onreadystatechange = function () {
    if (xhr.readyState === 4) {
      guShi.value = JSON.parse(xhr.responseText);
      printerInfo.value = guShi.value.content;
    }
  };
  xhr.send();
};

const scrollTo = (): void => {
  const content = document.querySelector(".content");
  if (content) {
    // 需要阻止默认行为，所以使用 passive: false
    content.addEventListener(
      "wheel",
      ((event: WheelEvent): void => {
        event.preventDefault();
        (content as HTMLElement).scrollLeft += event.deltaY;
      }) as EventListener,
      { passive: false }
    );
  }
};
</script>

<style lang="scss">
.vue-live2d-tool {
  cursor: url(../../public/normal.png) 1 1, auto !important;
}
.vue-live2d-toggle {
  text-align: center;
  height: 64px;
  bottom: 90px !important;
  position: absolute;
  left: 0;
  bottom: 31px;
  border-radius: 0 0.5rem 0.5rem 0 !important;
  cursor: url(../../public/normal.png) 1 1, auto !important;
  &:hover {
    background: var(--blue24);
    width: 1.5rem !important;
  }
}
</style>
<style lang="scss" scoped>
.index-container {
  // 签名墙
  .signature-wall {
    display: flex;
    flex-direction: column;
    position: relative;
    user-select: none;
    height: 100vh;
    overflow: hidden;

    .playful {
      color: var(--red);
      font-size: 40px;
    }

    .printer {
      color: var(--darkBlue);
      background: var(--background);
      border-radius: 10px;
      padding-left: 10px;
      padding-right: 10px;

      .cursor {
        margin-left: 1px;
        animation: hideToShow 0.7s infinite;
        font-weight: 200;
      }
    }

    .main-hero-waves-area {
      width: 100%;
      position: absolute;
      left: 0;
      bottom: -11px;
      z-index: 5;

      .waves-svg {
        width: 100%;
        height: 60px;

        .parallax > use:nth-child(1) {
          animation-delay: -2s;
          animation-duration: 7s;
          fill: var(--background);
          opacity: 0.5;
        }

        .parallax > use:nth-child(2) {
          animation-delay: -3s;
          animation-duration: 10s;
          fill: var(--background);
          opacity: 0.6;
        }

        .parallax > use:nth-child(3) {
          animation-delay: -4s;
          animation-duration: 13s;
          fill: var(--background);
          opacity: 0.7;
        }

        .parallax > use:nth-child(4) {
          animation-delay: -5s;
          animation-duration: 20s;
          fill: var(--background);
        }

        .parallax > use {
          animation: move-forever 30s cubic-bezier(0.55, 0.5, 0.45, 0.5) infinite;
        }
      }
    }

    .icon-arrow-down {
      color: var(--white);
      position: absolute;
      bottom: 60px;
      animation: my-shake 1.5s ease-out infinite;
      z-index: 15;
    }
  }

  // 页面容器
  .page-container-wrap {
    background: var(--background);
    position: relative;

    .page-container {
      display: flex;
      justify-content: center;
      width: 100%;
      padding: 0 130px 40px 130px;
      flex-direction: row;

      .recent-posts {
        width: calc(100% - 310px);

        .announcement {
          padding: 22px;
          border: 1px dashed var(--red);
          color: var(--bigRed1);
          border-radius: 10px;
          display: flex;
          max-width: 1080px;
          margin: 40px auto 40px;

          i {
            color: var(--red);
            font-size: 22px;
            margin: auto 0;
            animation: scale 0.8s ease-in-out infinite;
          }

          .notice-wrap {
            .notice-item {
              width: 100%;
              margin-left: 20px;
              line-height: 30px;
            }
          }
        }

        .content {
          border-radius: 11px;
          position: relative;
          max-width: 1080px;
          margin: 20px auto 40px;
          display: flex;
          overflow-x: auto;
          overflow-y: hidden;

          .category_group {
            border-radius: 12px;
            flex-shrink: 0;
            display: flex;
            flex-direction: column;
            justify-content: space-between;
            width: 200px;
            height: 186px;

            .category_item {
              background: linear-gradient(to right, var(--blue24), var(--blue11));
              overflow: hidden;
              height: 48%;
              border-radius: 12px;
              transition: all 1s;

              &:hover {
                transform: scale(0.95);
                transition-duration: 0.4s;
              }

              &.item {
                background: linear-gradient(to right, var(--blue22), var(--blue12));
              }

              .category_button {
                text-decoration: none;
                height: 100%;
                width: 100%;
                border-radius: 12px;
                display: inline-block;
                text-align: left;
                line-height: 4em;
                font-weight: 800;
                font-size: 20px;
                color: var(--white);
                overflow: hidden;

                &_text {
                  padding: 20px;
                  background: linear-gradient(to right, var(--red), var(--green1)) no-repeat right
                    bottom;
                  background-size: 0px 3px;
                  transition: background-size 1300ms;

                  &:hover {
                    background-position-x: left;
                    background-size: 100% 3px;
                  }
                }

                &.category_buttons::before {
                  line-height: 60px;
                  content: "推荐";
                  position: absolute;
                  z-index: 2;
                  color: var(--white);
                  top: -15px;
                  letter-spacing: 3px;
                  left: 140px;
                  font-size: 15px;
                  width: 50px;
                  height: 50px;
                  display: flex;
                  justify-content: center;
                  background: linear-gradient(90deg, var(--orange), var(--orange6));
                  border-radius: 0px 0px 12px 12px;
                }
              }
            }
          }

          .page {
            border-radius: 12px;
            margin-left: 8px;
            width: 200px;
            height: 186px;
            flex-shrink: 0;
            border: 2px solid var(--myAsideBorderColor);
            transition: all 1s;

            &:hover {
              transform: scale(0.95);
              background-color: var(--blue);
              transition-duration: 0.4s;
              border-color: var(--myAsideBorderColor1);

              span {
                color: var(--black1);
              }
            }

            .post_cover {
              width: 100%;
              height: 130px;
              position: relative;

              > a {
                text-decoration: none;
                display: block;

                img {
                  border-radius: 12px;
                  object-fit: cover;
                  width: 100%;
                  height: 130px;
                }
              }
            }

            .post_info {
              padding: 6px 9px;

              .article-title {
                color: var(--fontColor);
                -webkit-line-clamp: 2;
                line-clamp: 2;
                overflow: hidden;
                display: -webkit-box;
                -webkit-box-orient: vertical;
                line-height: 1.4;
                font-size: 15px;
                padding: 0;
              }
            }
          }
        }

        .aside-content {
          user-select: none;
          margin-top: 40px;
          margin-left: 10px;
          width: 300px;
          float: left;
        }

        .pagination-wrap {
          display: flex;
          justify-content: center;
          margin-top: 40px;

          .pagination {
            padding: 13px 15px;
            border: 1px solid var(--red2);
            border-radius: 3rem;
            color: var(--red1);
            width: 100px;
            user-select: none;
            text-align: center;

            &:hover {
              border: 1px solid var(--blue2);
              color: var(--orange2);
              box-shadow: 0 0 5px var(--blue2);
            }
          }

          .pagination-text {
            user-select: none;
            color: var(--red);
          }
        }
      }

      .aside-content {
        user-select: none;
        margin-top: 40px;
        margin-left: 10px;
        width: 300px;
        float: left;
      }
    }

    .live2d {
      position: fixed;
      z-index: 100;
      bottom: -0.8rem;
    }
  }
}

@keyframes move-forever {
  0% {
    transform: translate3d(-90px, 0, 0);
  }
  100% {
    transform: translate3d(85px, 0, 0);
  }
}

@media screen and (max-width: 1278px) {
  .index-container {
    .page-container-wrap {
      .page-container {
        padding: 0 40px 40px;
      }
    }
  }
}

@media screen and (max-width: 1100px) {
  .index-container {
    .page-container-wrap {
      .page-container {
        padding: 0 40px 40px;
      }
    }
  }
}

@media screen and (max-width: 1000px) {
  .index-container {
    .page-container-wrap {
      .page-container {
        /* 文章栏与侧标栏垂直排列 */
        flex-direction: column;

        .recent-posts {
          width: 100%;
        }

        .aside-content {
          width: 100%;
          max-width: unset;
          float: unset;
          margin: 40px auto 0;
        }
      }
    }
  }
}

@media screen and (max-width: 768px) {
  .index-container {
    .signature-wall {
      .playful {
        font-size: 30px;
      }
    }
  }
}
</style>
