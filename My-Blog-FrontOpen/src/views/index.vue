<template>
  <div>
    <!-- 首页图片 -->
    <div style="animation: header-effect 2s" :style="{ background: `${$store.state.changeBg}` }" class="background-image background-image-changeBg"></div>
    <!-- 首页文字 -->
    <div class="signature-wall myCenter my-animation-hideToShow">
      <h1 class="playful">
        <span v-for="(a, index) in $store.state.webInfo.webTitle" :key="index">{{ a }}</span>
      </h1>
      <div class="printer" @click="getGuShi()">
        <printer :printerInfo="printerInfo">
          <template slot="paper" slot-scope="scope">
            <h3>{{ scope.content }}<span class="cursor">|</span></h3>
          </template>
        </printer>
      </div>
      <i class="el-icon-arrow-down" @click="navigation('.page-container-wrap')"></i>
    </div>
    <!-- 首页内容 -->
    <div class="page-container-wrap">
      <div class="page-container">
        <!-- 内容页面 -->
        <div class="recent-posts">
          <!-- 视频介绍框 -->
          <div class="announcement">
            <i class="el-icon-s-opportunity" aria-hidden="true"></i>
            <div>
              <div v-for="(notice, index) in $store.state.webInfo.notices" :key="index">
                {{ notice }}
              </div>
            </div>
          </div>
          <!-- 滚动栏 -->
          <div>
            <div class="content">
              <div class="category_group">
                <div class="category_item" style="
                    background: linear-gradient(to right, #00868bdb, #3fc1c9db);
                  ">
                  <a href="https://www.zjh2002.icu/" class="category_button category_buttons">
                    <span class="category_button_text">云音乐</span>
                  </a>
                </div>
                <div @click="$router.push({ path: '/map' })" class="category_item" style="
                    background: linear-gradient(to right, #0a5abedb, #2fcbffdb);
                  ">
                  <a href="javascript:;" class="category_button">
                    <span class="category_button_text">四海之内</span>
                  </a>
                </div>
              </div>
              <div v-for="(article, index) in recommendArticles" :key="index" @click="
                  $router.push({ path: '/article', query: { id: article.id } })
                " class="page">
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
          <articleList :articleList="articles" :parentLoadingMark="parentLoadingMark"></articleList>
          <!-- 底部 -->
          <div class="pagination-wrap">
            <div @click="pageArticles()" class="pagination" v-if="pagination.total !== articles.length">
              下一页
            </div>
            <div v-else style="user-select: none; color: var(--red)">
              ~~到底啦~~
            </div>
          </div>
        </div>
        <!-- 侧边栏 -->
        <div class="aside-content" v-if="!$common.mobile()">
          <myAside @selectSort="selectSort"></myAside>
        </div>
      </div>
      <!-- live2d -->
      <live2d class="live2d" :style="style" :model="['Potion-Maker/Pio', 'school-2017-costume-yellow']" :direction="direction" :size="this.$common.mobile() ? 140 : size"></live2d>
    </div>
    <!-- 页脚 -->
    <div style="background: var(--background)">
      <myFooter></myFooter>
    </div>
  </div>
</template>
<script>
// 在组件中引入
import live2d from 'vue-live2d'
const printer = () => import('./common/printer')
const articleList = () => import('./articleList')
const myFooter = () => import('./common/myFooter')
const myAside = () => import('./myAside')
export default {
  components: {
    printer,
    articleList,
    myFooter,
    myAside,
    live2d
  },
  data() {
    return {
      loading: false,
      printerInfo: '咦，这个地方是用来干什么的？',
      pagination: {
        current: 1,
        size: 10,
        total: 0,
        searchKey: '',
        sortId: null
      },
      guShi: {
        content: '',
        origin: '',
        author: '',
        category: ''
      },
      articles: [],
      recommendArticles: [],
      direction: 'right',
      style: '',
      size: 210,
      tips: {
        mouseover: [
          {
            selector: '.vue-live2d',
            texts: []
          }
        ]
      },
      parentLoadingMark: false
    }
  },
  beforeRouteEnter(to, from, next) {
    if (from.path === '/') {
      next(vm => {
        vm.parentLoadingMark = true
      })
    } else {
      next()
    }
  },
  created() {
    this.getGuShi()
    this.getArticles()
  },
  mounted() {
    this.scrollTo()
    this.getRecommendArticles()
  },
  methods: {
    async selectSort(sort) {
      this.pagination = {
        current: 1,
        size: 10,
        total: 0,
        searchKey: '',
        sortId: sort.id,
        deltaY: 0
      }
      this.articles = []
      await this.getArticles()
      this.$nextTick(() => {
        document.querySelector('.recent-posts').scrollIntoView({
          behavior: 'smooth',
          block: 'start',
          inline: 'nearest'
        })
      })
    },
    pageArticles() {
      this.pagination.current = this.pagination.current + 1
      this.getArticles()
    },
    getArticles() {
      this.$http
        .post(this.$constant.baseURL + '/article/listArticle', this.pagination)
        .then(res => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.articles = this.articles.concat(res.result[0].records)
            this.pagination.total = res.result[0].total
            this.$store.commit('articleTotal', res.result[0].total)
            this.$store.commit('newArticles', this.articles.slice(0, 4))
          }
        })
        .catch(error => {
          this.$message({
            message: error.message,
            type: 'error'
          })
        })
    },
    getRecommendArticles() {
      const pagination = {
        current: 1,
        size: 10,
        recommendStatus: true
      }
      this.$http
        .post(this.$constant.baseURL + '/article/listArticle', pagination)
        .then(res => {
          if (!this.$common.isEmpty(res.result[0])) {
            let formatRecommendArticles = res.result[0].records
            const myBlog = formatRecommendArticles.filter(item => {
              return item.id === 20
            })
            const notMyBlog = formatRecommendArticles.filter(item => {
              return item.id !== 20
            })
            const recommendArticles = myBlog.concat(notMyBlog)
            this.recommendArticles = recommendArticles
          }
        })
        .catch(error => {
          this.$message({
            message: error.message,
            type: 'error'
          })
        })
    },
    navigation(selector) {
      let pageId = document.querySelector(selector)
      window.scrollTo({
        top: pageId.offsetTop,
        behavior: 'smooth'
      })
    },
    getGuShi() {
      let that = this
      let xhr = new XMLHttpRequest()
      xhr.open('get', this.$constant.jinrishici)
      xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
          that.guShi = JSON.parse(xhr.responseText)
          that.printerInfo = that.guShi.content
        }
      }
      xhr.send()
    },
    scrollTo() {
      let content = document.querySelector('.content')
      content.addEventListener('wheel', event => {
        event.preventDefault()
        content.scrollLeft += event.deltaY
      })
    }
  }
}
</script>

<style scoped>
.live2d {
  position: fixed;
  z-index: 100;
  bottom: -0.8rem;
}

.signature-wall {
  /* 向下排列 */
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
  background: var(--translucent);
  border-radius: 10px;
  padding-left: 10px;
  padding-right: 10px;
}

/* 光标 */
.cursor {
  margin-left: 1px;
  animation: hideToShow 0.7s infinite;
  font-weight: 200;
}

.el-icon-arrow-down {
  font-size: 40px;
  font-weight: bold;
  color: white;
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
  width: 98%;
  padding: 0 0px 40px 0px;
  margin: 0 auto;
  flex-direction: row;
}

.recent-posts {
  width: 90%;
}

.announcement {
  padding: 22px;
  border: 1px dashed var(--red);
  color: var(--bigRed);
  border-radius: 10px;
  display: flex;
  max-width: 1080px;
  margin: 40px auto 40px;
}

.announcement i {
  color: var(--red);
  font-size: 22px;
  margin: auto 0;
  animation: scale 0.8s ease-in-out infinite;
}

.announcement div div {
  width: 100%;
  margin-left: 20px;
  line-height: 30px;
}

.aside-content {
  width: calc(30% - 40px);
  user-select: none;
  margin-top: 40px;
  margin-left: 20px;
  max-width: 300px;
  float: left;
}

.pagination-wrap {
  display: flex;
  justify-content: center;
  margin-top: 40px;
}

.pagination {
  padding: 13px 15px;
  border: 1px solid var(--red);
  border-radius: 3rem;
  color: var(--red);
  width: 100px;
  user-select: none;
  text-align: center;
}

.pagination:hover {
  border: 1px solid var(--blue);
  color: orange;
  box-shadow: 0 0 5px var(--blue);
}

@media screen and (max-width: 1100px) {
  .recent-posts {
    width: 100%;
  }

  .page-container {
    width: 94%;
  }
}

@media screen and (max-width: 1000px) {
  .page-container {
    /* 文章栏与侧标栏垂直排列 */
    flex-direction: column;
  }

  .aside-content {
    width: 90%;
    max-width: unset;
    float: unset;
    margin: 40px auto 0;
  }
}

@media screen and (max-width: 768px) {
  .playful {
    font-size: 36px;
  }
  h1 {
    font-size: 35px;
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
}

.page::before {
  content: '';
  border-radius: 12px;
  position: absolute;
  width: 198px;
  height: 185px;
  background-color: rgba(255, 255, 255, 0.4);
}

.page:hover {
  background-color: var(--blue);
  transition-duration: 0.4s;
}

.page:hover span {
  color: black;
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
  overflow: hidden;
  height: 48%;
  border-radius: 12px;
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
  color: white;
  overflow: hidden;
}
.category_buttons::before {
  line-height: 60px;
  content: '推荐';
  position: absolute;
  z-index: 2;
  color: rgb(255, 255, 255);
  top: -15px;
  letter-spacing: 3px;
  left: 140px;
  font-size: 15px;
  width: 50px;
  height: 50px;
  display: flex;
  justify-content: center;
  background: linear-gradient(90deg, rgb(255, 184, 99), rgb(255, 117, 0));
  border-radius: 0px 0px 12px 12px;
}
.category_button_text {
  padding: 20px;
  background: linear-gradient(to right, #ec695c, #61c454) no-repeat right bottom;
  background-size: 0px 3px;
  transition: background-size 1300ms;
}

.category_button_text:hover {
  background-position-x: left;
  background-size: 100% 3px;
}

.post_cover {
  width: 100%;
  height: 130px;
  position: relative;
}

.post_cover > a {
  text-decoration: none;
  display: block;
}

.post_cover img {
  border-radius: 12px;
  object-fit: cover;
  width: 100%;
  height: 130px;
}

.post_info {
  padding: 6px 9px;
}

.article-title {
  color: var(--maxLightRed);
  -webkit-line-clamp: 2;
  overflow: hidden;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  line-height: 1.4;
  font-size: 15px;
  padding: 0;
}
</style>
