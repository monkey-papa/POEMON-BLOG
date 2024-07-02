<template>
  <div>
    <!-- 图片 -->
    <div
      style="animation: header-effect 2s"
      :style="{ background: `${$store.state.changeBg}` }"
      class="background-image background-image-changeBg"
    ></div>
    <!-- 标签区域 -->
    <div class="main">
      <div class="layout hide-aside">
        <div id="tag">
          <!-- 标签栏 -->
          <div id="catalog-bar">
            <i class="fa fa-tags" style="color: purple; margin-right: 4px"></i>
            <div id="catalog-list">
              <div
                v-for="(label, i) in $store.getters.labelInfo"
                :key="i"
                @click="toggleLabel(label)"
                class="catalog-list-item"
                :class="{
                  selected:
                    !$common.isEmpty(labelId) && parseInt(labelId) === label.id,
                }"
              >
                <a href="javascript:;"
                  >{{ label.labelName }}<sup>{{ label.countOfLabel }}</sup></a
                >
              </div>
            </div>
          </div>
          <h1 class="article-sort-title">标签 - {{ curLabelName }}</h1>
          <!-- 内容 -->
          <div v-if="!$common.isEmpty(articles)" class="article-sort">
            <div
              class="article-sort-item"
              v-for="(item, i) in articles"
              :key="i"
            >
              <a
                @click="
                  $router.push({
                    path: '/article',
                    query: { id: item.label[0].articleId },
                  })
                "
                href="javascript:;"
                class="article-sort-item-img"
              >
                <el-image class="img" v-once lazy :src="item.articleCover">
                  <!-- 懒加载图片 -->
                  <div slot="placeholder">
                    <div>
                      <img
                        :src="$store.state.webInfo.randomCover[1]"
                        style="object-fit: cover; width: 100%; height: 100%"
                      />
                    </div>
                  </div>
                  <!-- 错误图片 -->
                  <div slot="error" class="image-slot myCenter">
                    <img
                      class="error-img"
                      :src="$store.state.webInfo.randomCover[0]"
                      style="object-fit: cover; width: 100%; height: 100%"
                    />
                    <div class="error-text">
                      <div style="color: wheat">ya，图片跑丢了┮﹏┭</div>
                    </div>
                  </div>
                </el-image>
              </a>
              <div class="article-sort-item-info">
                <div class="article-sort-item-time">
                  <time class="post-meta-date-created">{{
                    item.createTime | formatter
                  }}</time>
                </div>
                <a
                  href="javascript:;"
                  @click="
                    $router.push({
                      path: '/article',
                      query: { id: item.label[0].articleId },
                    })
                  "
                  class="article-sort-item-title"
                  >{{ item.articleTitle
                  }}<sup class="article-sort-item-index">{{ i + 1 }}</sup></a
                >
                <div class="article-meta-wrap">
                  <span
                    @click.stop="
                      $router.push({
                        path: '/sort',
                        query: { sortId: item.sortId },
                      })
                    "
                    class="article-sort-item-categories"
                  >
                    <img
                      style="vertical-align: -3px"
                      src="../assets/svg/sort2.svg"
                    />
                    {{ item.sort[0].sortName }}</span
                  >
                  <span
                    @click.stop="
                      $router.push({
                        path: '/sort',
                        query: { sortId: item.sortId, labelId: labelId },
                      })
                    "
                    class="article-sort-item-tags"
                  >
                    <img
                      style="vertical-align: -3px"
                      src="../assets/svg/tag2.svg"
                    />
                    {{ item.label[0].labelName }}</span
                  >
                </div>
              </div>
            </div>
          </div>
          <div v-else class="myCenter" style="color: var(--darkBlue)">
            <h3>主人偷懒了，这个标签还没有文章噢~</h3>
          </div>
          <!-- 分页器 -->
          <div class="pagination">
            <el-pagination
              background
              layout="prev, pager, next"
              :current-page="pagination.current"
              :page-size="pagination.size"
              :total="pagination.total"
              @current-change="handlePageChange"
            >
            </el-pagination>
          </div>
        </div>
      </div>
    </div>
    <!-- 页脚 -->
    <div>
      <myFooter></myFooter>
    </div>
  </div>
</template>
<script>
const myFooter = () => import("./common/myFooter");
export default {
  data() {
    return {
      labelId: this.$route.query.labelId,
      pagination: {
        current: 1,
        size: 10,
        total: 0,
        searchKey: "",
        sortId: null,
        labelId: this.$route.query.labelId,
      },
      articles: [],
    };
  },
  components: {
    myFooter,
  },
  filters: {
    formatter(row) {
      const day = row.split(".")[0].split("T")[0];
      const time = row.split(".")[0].split("T")[1];
      return `${day} 日 ${time}`;
    },
  },
  mounted() {
    this.scrollTo();
    this.getArticles();
  },
  computed: {
    curLabelName() {
      const labelList = this.$store.getters.labelInfo;
      const curLabel = labelList.filter((item) => {
        return item.id === parseInt(this.labelId);
      });
      return curLabel[0].labelName;
    },
  },
  methods: {
    handlePageChange(val) {
      this.pagination.current = val;
      this.getArticles();
    },
    scrollTo() {
      let content = document.querySelector("#catalog-list");
      content.addEventListener("wheel", (event) => {
        event.preventDefault();
        content.scrollLeft += event.deltaY;
      });
    },
    toggleLabel(label) {
      if (parseInt(this.labelId) === label.id) {
        return;
      }
      this.articles = [];
      this.$router.push({ path: "/tags", query: { labelId: label.id } });
      this.labelId = label.id;
      this.pagination = {
        current: 1,
        size: 10,
        total: 0,
        searchKey: "",
        sortId: null,
        labelId: label.id,
      };
      this.$nextTick(() => {
        this.getArticles();
      });
    },
    getArticles() {
      this.$http
        .post(this.$constant.baseURL + "/article/listArticle", this.pagination)
        .then((res) => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.articles = res.result[0].records;
            this.pagination.total = res.result[0].total;
          }
        })
        .catch((error) => {
          this.$message({
            message: error.message,
            type: "error",
          });
        });
    },
  },
};
</script>
<style scoped>
.main {
  margin: 0 auto;
  padding: 45px 5px 10px;
  min-height: 100vh;
}
.layout {
  display: flex;
}
.layout .hide-aside {
  max-width: 1300px;
}
div #tag {
  height: 100%;
  width: 100%;
  border: 2px dashed rgba(0, 255, 255, 0.6);
  font-weight: 700;
  background: var(--background);
}
.layout > div {
  border-radius: 18px;
  padding: 10px 35px;
}
#catalog-bar {
  padding: 0.4rem 0.8rem;
  border-radius: 0.5rem;
  display: flex;
  border: 2px dashed var(--blue);
  justify-content: space-between;
}
#catalog-list {
  display: flex;
  white-space: nowrap;
  overflow-x: scroll;
  overflow-y: hidden;
}
.catalog-list-item a {
  font-size: 16px;
  margin: 0 0.3em;
  padding: 0.5em 0.4em 0.5em;
  font-weight: 400;
  border-radius: 0.5rem;
  color: var(--red);
  line-height: 1.5rem;
}
.catalog-list-item:hover a {
  transition: all 0.3s;
  border-radius: 0.5rem;
  background-color: rgba(29, 191, 255, 0.945);
  color: white;
}
.catalog-list-item.selected a {
  background: #00c4b6f1;
  color: #fff;
}
.article-sort-title {
  text-indent: 1em;
  font-weight: 400;
  color: var(--darkBlue);
}
.article-sort {
  padding-left: 5px;
  margin-top: 7.5px;
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: space-between;
}
.article-sort-item {
  position: relative;
  display: flex;
  align-items: center;
  margin: 12px 0;
  border-radius: 12px;
  padding: 8px;
  width: calc(50% - 7.5px);
  background: rgba(236, 105, 92, 0.3);
}
.article-sort-item-img {
  overflow: hidden;
  width: 160px;
  height: 90px;
  border-radius: 7px;
  border: 1px solid var(--blue);
}
.article-sort-item-img .img {
  transition: all 0.6s;
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 7px;
}
.article-sort-item-img .img:hover {
  transform: scale(1.2);
}
.article-sort-item-info {
  -webkit-box-flex: 1;
  -moz-box-flex: 1;
  -o-box-flex: 1;
  box-flex: 1;
  -webkit-flex: 1;
  -ms-flex: 1;
  flex: 1;
  padding-left: 10px;
}
.post-meta-date-created {
  bottom: 5px;
  right: 10px;
  position: absolute;
  font-size: 14px;
  line-height: 14px;
  color: var(--darkBlue);
}
.article-sort-item-title {
  width: 90%;
  color: var(--bigRed);
  font-size: 17px;
  transition: all 0.3s;
  font-weight: 500;
  line-height: 1.4em;
  display: -webkit-box;
  /*弹性伸缩盒子*/
  -webkit-box-orient: vertical;
  /*子元素垂直排列*/
  white-space: wrap;
  /*溢出换行*/
  -webkit-line-clamp: 2;
  /*显示的行数*/
  overflow: hidden;
  /*溢出部分隐藏*/
  text-overflow: ellipsis;
  /*溢出部分用省略号显示*/
}
.article-sort-item-title:hover {
  color: #00c4b6;
  padding: 0px 15px;
}
.article-sort-item-index {
  opacity: 0.8;
  position: absolute;
  top: 0.2rem;
  right: 0.5rem;
  font-style: italic;
  font-size: 2rem;
  line-height: 1.5rem;
  color: var(--red);
}
.article-meta-wrap {
  margin: 10px;
}
.article-sort-item-categories,
.article-sort-item-tags {
  margin: 5px;
  color: rgba(226, 36, 220, 0.632);
  font-size: 14px;
  font-weight: 400;
  display: -webkit-box;
  /*弹性伸缩盒子*/
  -webkit-box-orient: vertical;
  /*子元素垂直排列*/
  white-space: wrap;
  /*溢出换行*/
  -webkit-line-clamp: 1;
  /*显示的行数*/
  overflow: hidden;
  /*溢出部分隐藏*/
  text-overflow: ellipsis;
  /*溢出部分用省略号显示*/
}
.article-sort-item-categories:hover,
.article-sort-item-tags:hover {
  transition: all 0.4s;
  color: var(--darkBlue);
}
@media screen and (max-width: 1286px) {
  .main {
    padding: 60px 3px;
  }
}
@media screen and (max-width: 974px) {
  .article-sort-item {
    width: 100%;
  }
}
@media screen and (max-width: 523px) {
  #tag {
    padding: 10px 8px;
  }
  .article-sort-item {
    margin: 6px 0;
  }
  .article-sort-item-img {
    width: 80px;
    height: 76px;
  }
  .article-sort-item-title {
    width: 218px;
    font-size: 16px;
  }
}
.pagination {
  margin: 20px 0;
  text-align: center;
}
::v-deep .el-pagination.is-background .el-pager li:hover {
  color: white;
  background: rgba(29, 191, 255, 0.945);
}
::v-deep .el-pagination.is-background .el-pager li:not(.disabled).active {
  background: rgba(0, 196, 182, 0.945);
  color: white;
}
.error-img {
  position: relative;
  z-index: 1;
}
.error-text {
  z-index: 2;
  position: absolute;
  font-size: 16px;
  line-height: 1.8;
  letter-spacing: 8px;
  color: rgb(193, 27, 126);
}
</style>
