<template>
  <div>
    <div style="animation: header-effect 2s;" :style="{ background: `${$store.state.changeBg}` }" class="background-image background-image-changeBg"></div>
    <!-- 两句诗 -->
    <div class="my-animation-slide-top">
      <twoPoem></twoPoem>
    </div>
    <div style="background: var(--background); padding-top: 40px" class="my-animation-slide-bottom">
      <!-- 标签 -->
      <div class="sort-warp shadow-box" v-if="!$common.isEmpty(sort) && !$common.isEmpty(sort.labels)">
        <div v-for="(label, index) in sort.labels" :key="index" :class="{
          isActive:
            !$common.isEmpty(labelId) && parseInt(labelId) === label.id,
        }" @click="toggleLabel(label)">
          <proTag :info="label.labelName + ' ' + label.countOfLabel" :color="$constant.tree_hole_color[Math.floor(Math.random() * 6)]" style="margin: 12px">
          </proTag>
        </div>
      </div>

      <!-- 文章 -->
      <div class="article-wrap">
        <articleList :articleList="articles"></articleList>
        <div class="pagination-wrap">
          <div @click="pageArticles()" class="pagination" v-if="pagination.total !== articles.length">
            下一页
          </div>
          <div v-else style="user-select: none;color: var(--red);">~~( •̀ ω •́ )y 到底啦~~</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
const twoPoem = () => import('./common/twoPoem')
const proTag = () => import('./common/proTag')
const articleList = () => import('./articleList')

export default {
  components: {
    twoPoem,
    proTag,
    articleList
  },
  data() {
    return {
      sortId: this.$route.query.sortId,
      labelId: this.$route.query.labelId,
      sort: null,
      pagination: {
        current: 1,
        size: 10,
        total: 0,
        searchKey: '',
        sortId: this.$route.query.sortId,
        labelId: this.$route.query.labelId
      },
      articles: []
    }
  },
  created() {
    this.getSort()
    this.getArticles()
  },
  methods: {
    pageArticles() {
      this.pagination.current = this.pagination.current + 1
      this.getArticles()
    },
    getSort() {
      let sortInfo = this.$store.state.sortInfo
      if (!this.$common.isEmpty(sortInfo)) {
        let sortArray = sortInfo.filter(f => {
          return f.id === parseInt(this.sortId)
        })
        if (!this.$common.isEmpty(sortArray)) {
          this.sort = sortArray[0]
        }
      }
    },
    toggleLabel(label) {
      if (parseInt(this.labelId) === label.id) {
        return
      }
      this.articles = []
      this.$router.push({
        path: '/sort',
        query: { sortId: this.sortId, labelId: label.id }
      })
      this.labelId = label.id
      this.pagination = {
        current: 1,
        size: 10,
        total: 0,
        searchKey: '',
        sortId: this.$route.query.sortId,
        labelId: label.id
      }
      this.$nextTick(() => {
        this.getArticles()
      })
    },
    getArticles() {
      this.$http
        .post(this.$constant.baseURL + '/article/listArticle', this.pagination)
        .then(res => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.articles = this.articles.concat(res.result[0].records)
            this.pagination.total = res.result[0].total
          }
        })
        .catch(error => {
          this.$notify({
            type: 'error',
            title: '可恶🤬',
            message: error.message,
            position: 'top-left',
            offset: 50
          })
        })
    }
  }
}
</script>

<style lang="scss" scoped>
.sort-warp {
  width: 70%;
  max-width: 780px;
  margin: 0 auto;
  padding: 20px;
  border-radius: 10px;
  display: flex;
  flex-wrap: wrap;
}
.article-wrap {
  width: 70%;
  margin: 40px auto 0;
  min-height: 600px;
}
.isActive {
  animation: scale 1.5s ease-in-out infinite;
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

@media screen and (max-width: 900px) {
  .sort-warp {
    width: 90%;
  }
  .article-wrap {
    width: 90%;
  }
}
</style>
