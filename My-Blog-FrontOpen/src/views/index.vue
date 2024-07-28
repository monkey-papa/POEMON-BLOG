<template>
  <div>
    <!-- 首页图片 -->
    <div
      style="animation: header-effect 2s"
      :style="{ background: `${$store.state.changeBg}` }"
      class="background-image background-image-changeBg"
    ></div>
    <!-- 首页文字 -->
    <div class="signature-wall myCenter my-animation-hideToShow">
      <h1 class="playful">
        <span
          v-for="(a, index) in $store.state.webInfo.webTitle"
          :key="index"
          >{{ a }}</span
        >
      </h1>
      <div class="printer" @click="getGuShi()">
        <printer :printerInfo="printerInfo">
          <template slot="paper" slot-scope="scope">
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
            ></path>
          </defs>
          <g class="parallax">
            <use href="#gentle-wave" x="48" y="0"></use>
            <use href="#gentle-wave" x="48" y="3"></use>
            <use href="#gentle-wave" x="48" y="5"></use>
            <use href="#gentle-wave" x="48" y="7"></use>
          </g>
        </svg>
      </section>
      <i
        class="el-icon-arrow-down"
        @click="navigation('.page-container-wrap')"
      ></i>
    </div>
    <!-- 首页内容 -->
    <div class="page-container-wrap">
      <div class="page-container">
        <!-- 内容页面 -->
        <div class="recent-posts">
          <!-- 提示框 -->
          <div class="announcement">
            <i class="el-icon-s-opportunity" aria-hidden="true"></i>
            <div>
              <div
                v-for="(notice, index) in $store.state.webInfo.notices"
                :key="index"
              >
                {{ notice }}
              </div>
            </div>
          </div>
          <!-- 滚动栏 -->
          <div>
            <div class="content">
              <div class="category_group">
                <div class="category_item item">
                  <a
                    href="https://www.zjh2002.icu/"
                    class="category_button category_buttons"
                  >
                    <span class="category_button_text">云音乐</span>
                  </a>
                </div>
                <div
                  @click="$router.push({ path: '/map' })"
                  class="category_item"
                >
                  <a href="javascript:;" class="category_button">
                    <span class="category_button_text">四海之内</span>
                  </a>
                </div>
              </div>
              <div
                v-for="(article, index) in recommendArticles"
                :key="index"
                @click="
                  $router.push({ path: '/article', query: { id: article.id } })
                "
                class="page"
              >
                <div class="post_cover">
                  <a href="javascript:;">
                    <img :src="article.articleCover" />
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
            <myAside @selectSort="selectSort"></myAside>
          </div>
          <!-- 文章列表 -->
          <articleList
            :articleList="articles"
            :parentLoadingMark="parentLoadingMark"
          ></articleList>
          <!-- 底部 -->
          <div class="pagination-wrap">
            <div
              @click="pageArticles()"
              class="pagination"
              v-if="pagination.total !== articles.length"
            >
              下一页
            </div>
            <div v-else style="user-select: none; color: var(--red)">
              ~~( •̀ ω •́ )y 到底啦~~
            </div>
          </div>
        </div>
        <!-- 侧边栏 -->
        <div class="aside-content" v-if="!$common.mobile()">
          <myAside @selectSort="selectSort"></myAside>
        </div>
      </div>
      <!-- live2d -->
      <live2d
        class="live2d"
        :style="style"
        :model="['Potion-Maker/Pio', 'school-2017-costume-yellow']"
        :direction="direction"
        :size="this.$common.mobile() ? 140 : size"
      ></live2d>
    </div>
  </div>
</template>
<script>
// 在组件中引入
import live2d from "vue-live2d";
const printer = () => import("./common/printer");
const articleList = () => import("./articleList");
const myAside = () => import("./myAside");
export default {
  components: {
    printer,
    articleList,
    myAside,
    live2d,
  },
  data() {
    return {
      loading: false,
      printerInfo: "咦，这个地方是用来干什么的？",
      pagination: {
        current: 1,
        size: 10,
        total: 0,
        searchKey: "",
        sortId: null,
      },
      guShi: {
        content: "",
        origin: "",
        author: "",
        category: "",
      },
      articles: [],
      recommendArticles: [],
      direction: "right",
      style: "",
      size: 210,
      tips: {
        mouseover: [
          {
            selector: ".vue-live2d",
            texts: [],
          },
        ],
      },
      parentLoadingMark: false,
    };
  },
  beforeRouteEnter(to, from, next) {
    if (from.path === "/") {
      next((vm) => {
        vm.parentLoadingMark = true;
      });
    } else {
      next();
    }
  },
  created() {
    this.getGuShi();
    this.getArticles();
  },
  mounted() {
    this.scrollTo();
    this.getRecommendArticles();
  },
  methods: {
    async selectSort(sort) {
      this.pagination = {
        current: 1,
        size: 10,
        total: 0,
        searchKey: "",
        sortId: sort.id,
        deltaY: 0,
      };
      this.articles = [];
      await this.getArticles();
      this.$nextTick(() => {
        document.querySelector(".recent-posts").scrollIntoView({
          behavior: "smooth",
          block: "start",
          inline: "nearest",
        });
      });
    },
    pageArticles() {
      this.pagination.current = this.pagination.current + 1;
      this.getArticles();
    },
    getArticles() {
      this.$http
        .post(this.$constant.baseURL + "/article/listArticle", this.pagination)
        .then((res) => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.articles = this.articles.concat(res.result[0].records);
            this.pagination.total = res.result[0].total;
            this.$store.commit("articleTotal", res.result[0].total);
            this.$store.commit("newArticles", this.articles.slice(0, 4));
          }
        })
        .catch((error) => {
          this.$notify({
            type: "error",
            title: "可恶🤬",
            message: error.message,
            position: "top-left",
            offset: 50,
          });
        });
    },
    getRecommendArticles() {
      const pagination = {
        current: 1,
        size: 10,
        recommendStatus: true,
      };
      this.$http
        .post(this.$constant.baseURL + "/article/listArticle", pagination)
        .then((res) => {
          if (!this.$common.isEmpty(res.result[0])) {
            let formatRecommendArticles = res.result[0].records;
            const myBlog = formatRecommendArticles.filter((item) => {
              return item.id === 20;
            });
            const notMyBlog = formatRecommendArticles.filter((item) => {
              return item.id !== 20;
            });
            const recommendArticles = myBlog.concat(notMyBlog);
            this.recommendArticles = recommendArticles;
          }
        })
        .catch((error) => {
          this.$notify({
            type: "error",
            title: "可恶🤬",
            message: error.message,
            position: "top-left",
            offset: 50,
          });
        });
    },
    navigation(selector) {
      let pageId = document.querySelector(selector);
      window.scrollTo({
        top: pageId.offsetTop,
        behavior: "smooth",
      });
    },
    getGuShi() {
      let that = this;
      let xhr = new XMLHttpRequest();
      xhr.open("get", this.$constant.jinrishici);
      xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
          that.guShi = JSON.parse(xhr.responseText);
          that.printerInfo = that.guShi.content;
        }
      };
      xhr.send();
    },
    scrollTo() {
      let content = document.querySelector(".content");
      content.addEventListener("wheel", (event) => {
        event.preventDefault();
        content.scrollLeft += event.deltaY;
      });
    },
  },
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
.live2d {
  position: fixed;
  z-index: 100;
  bottom: -0.8rem;
}
.signature-wall {
  display: flex;
  flex-direction: column;
  position: relative;
  user-select: none;
  height: 100vh;
  overflow: hidden;
}
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
}
.cursor {
  margin-left: 1px;
  animation: hideToShow 0.7s infinite;
  font-weight: 200;
}
.el-icon-arrow-down {
  font-size: 40px;
  font-weight: bold;
  color: var(--white);
  position: absolute;
  bottom: 60px;
  animation: my-shake 1.5s ease-out infinite;
  z-index: 15;
}
.page-container-wrap {
  background: var(--background);
  position: relative;
}
.page-container {
  display: flex;
  justify-content: center;
  width: 100%;
  padding: 0 130px 40px 130px;
  flex-direction: row;
}
.recent-posts {
  width: calc(100% - 310px);
}
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
  div div {
    width: 100%;
    margin-left: 20px;
    line-height: 30px;
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
}
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
.content {
  border-radius: 11px;
  position: relative;
  max-width: 1080px;
  margin: 20px auto 40px;
  display: flex;
  overflow-x: auto;
  overflow-y: hidden;
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
}
.category_group {
  border-radius: 12px;
  flex-shrink: 0;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  width: 200px;
  height: 186px;
}
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
    background: linear-gradient(to right, var(--red), var(--green1)) no-repeat
      right bottom;
    background-size: 0px 3px;
    transition: background-size 1300ms;
    &:hover {
      background-position-x: left;
      background-size: 100% 3px;
    }
  }
}
.category_buttons::before {
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
.post_cover {
  width: 100%;
  height: 130px;
  position: relative;
  > a {
    text-decoration: none;
    display: block;
  }
  img {
    border-radius: 12px;
    object-fit: cover;
    width: 100%;
    height: 130px;
  }
}
.post_info {
  padding: 6px 9px;
}
.article-title {
  color: var(--fontColor);
  -webkit-line-clamp: 2;
  overflow: hidden;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  line-height: 1.4;
  font-size: 15px;
  padding: 0;
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
  .page-container {
    padding: 0 40px 40px;
  }
}
@media screen and (max-width: 1100px) {
  .page-container {
    padding: 0 40px 40px;
  }
}
@media screen and (max-width: 1000px) {
  .page-container {
    /* 文章栏与侧标栏垂直排列 */
    flex-direction: column;
  }
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
@media screen and (max-width: 768px) {
  .playful {
    font-size: 30px;
  }
}
</style>
