<template>
  <div v-if="!$common.isEmpty(article)">
    <!-- 首页图片 -->
    <div
      style="animation: header-effect 2s"
      :style="{ background: `${$store.state.changeBg}` }"
      class="background-image background-image-changeBg"
    ></div>
    <!-- 顶部 -->
    <div class="article-head my-animation-slide-top">
      <!-- 文章信息 -->
      <div class="article-info-container">
        <div class="article-title">{{ article.articleTitle }}</div>
        <div class="article-info">
          <img style="vertical-align: -2px" src="../assets/svg/redPeople.svg" />
          <span>&nbsp;{{ article.username }}</span>
          <span>·</span>
          <img style="vertical-align: -2px" src="../assets/svg/calendar.svg" />
          <span>&nbsp;{{ article.createTime | formatter }}</span>
          <span>·</span>
          <img style="vertical-align: -2px" src="../assets/svg/fire.svg" />
          <span>&nbsp;{{ article.viewCount }}</span>
          <span>·</span>
          <img style="vertical-align: -2px" src="../assets/svg/comment.svg" />
          <span>&nbsp;{{ article.commentCount }}</span>
          <span>·</span>
          <img style="vertical-align: -2px" src="../assets/svg/star.svg" />
          <span>&nbsp;{{ article.likeCount }}</span>
        </div>
      </div>
      <div
        class="article-info-news"
        @click="weiYanDialogVisible = true"
        v-if="
          !$common.isEmpty($store.state.currentUser) &&
          $store.state.currentUser.id === article.userId
        "
      >
        <img src="../assets/svg/plus.svg" />
      </div>
    </div>
    <!-- 文章 -->
    <div style="background: var(--background)" class="article-background">
      <div class="article-container my-animation-slide-bottom">
        <!-- 最新进展 -->
        <div v-if="!$common.isEmpty(treeHoleList)" class="process-wrap">
          <el-collapse accordion value="1">
            <el-collapse-item title="最新进展" name="1">
              <process
                :treeHoleList="treeHoleList"
                @deleteTreeHole="deleteTreeHole"
              ></process>
            </el-collapse-item>
          </el-collapse>
          <hr />
        </div>
        <!-- 文章内容 -->
        <div v-html="articleContentHtml" class="entry-content"></div>
        <!-- 最后更新时间 -->
        <div class="article-update-time">
          <span>文章最后更新于 {{ article.updateTime | formatter }}</span>
        </div>
        <!-- 分类 -->
        <div class="article-sort">
          <span
            @click="
              $router.push({
                path: '/sort',
                query: { sortId: article.sortId, labelId: article.labelId },
              })
            "
            >{{
              article.sort[0].sortName + " ▶ " + article.label[0].labelName
            }}</span
          >
        </div>
        <!-- 作者信息 -->
        <blockquote>
          <div style="color: var(--blue)">作者：{{ article.username }}</div>
          <div style="color: var(--blue)">版权声明：转载请注明文章出处</div>
        </blockquote>
        <!-- 点赞 -->
        <div class="myCenter" id="article-like" style="color: var(--bigRed)">
          <i
            @click="articleLike"
            class="el-icon-thumb article-like-icon"
            :class="{ 'article-like': article.articleLikeStatus === 1 }"
          ></i>
          点个赞再走叭~~
        </div>
        <!-- 评论 -->
        <div v-if="article.commentStatus === true">
          <comment
            :type="'article'"
            :source="article.id"
          ></comment>
        </div>
      </div>
      <div
        v-show="!this.$common.mobile() && !mobile"
        id="toc"
        class="toc"
      ></div>
    </div>
    <div style="background: var(--background)">
      <myFooter></myFooter>
    </div>
    <div id="toc-button" @click="clickTocButton()">
      <i class="fa fa-align-justify" aria-hidden="true"></i>
    </div>
    <el-dialog
      title="最新进展"
      :visible.sync="weiYanDialogVisible"
      width="40%"
      :append-to-body="true"
      destroy-on-close
      custom-class="dialog"
      center
    >
      <div>
        <div class="myCenter" style="margin-bottom: 20px">
          <el-date-picker
            v-model="newsTime"
            value-format="yyyy-MM-dd HH:mm:ss"
            type="datetime"
            align="center"
            placeholder="选择日期时间"
          >
          </el-date-picker>
        </div>
        <commentBox :disableGraffiti="true" @submitComment="submitWeiYan">
        </commentBox>
      </div>
    </el-dialog>
  </div>
</template>

<script>
const myFooter = () => import("./common/myFooter");
const comment = () => import("./common/comment");
const process = () => import("./common/process");
const commentBox = () => import("./common/commentBox");
import MarkdownIt from "markdown-it";
export default {
  components: {
    myFooter,
    comment,
    commentBox,
    process,
  },
  data() {
    return {
      id: this.$route.query.id,
      article: {},
      articleContentHtml: "",
      treeHoleList: [],
      weiYanDialogVisible: false,
      newsTime: "",
      mobile: false,
      tocbotDom: null,
    };
  },
  created() {
    this.getArticle();
    this.mobile = document.body.clientWidth < 500;
    window.addEventListener("resize", () => {
      let docWidth = document.body.clientWidth;
      if (docWidth < 500) {
        this.mobile = true;
      } else {
        this.mobile = false;
      }
    });
  },
  mounted() {
    if (!this.mobile) {
      window.addEventListener("scroll", this.onScrollPage);
    }
  },
  beforeDestroy() {
    window.removeEventListener("scroll", this.onScrollPage);
  },
  filters: {
    formatter(row) {
      const day = row.split(".")[0].split("T")[0];
      const time = row.split(".")[0].split("T")[1];
      return `${day} 日 ${time}`;
    },
  },
  methods: {
    clickTocButton() {
      let display = $(".toc");
      if ("none" === display.css("display") || !display.length) {
        const articleDom = $(".article-background");
        articleDom.append(this.tocbotDom);
      } else {
        this.tocbotDom = display;
        display.remove();
      }
    },
    deleteTreeHole(id) {
      if (this.$common.isEmpty(this.$store.state.currentUser)) {
        this.$message({
          message: "请先登录！",
          type: "error",
        });
        return;
      }
      this.$confirm("确认删除？", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
        center: true,
      })
        .then(() => {
          this.$http
            .get(this.$constant.baseURL + "/weiYan/deleteWeiYan/", { id: id })
            .then((res) => {
              this.$message({
                type: "success",
                message: "删除成功!",
              });
              this.getNews();
            })
            .catch((error) => {
              this.$message({
                message: error.message,
                type: "error",
              });
            });
        })
        .catch(() => {
          this.$message({
            type: "success",
            message: "已取消删除!",
          });
        });
    },
    submitWeiYan(content) {
      let weiYan = {
        content: content,
        createTime: this.newsTime,
        source: this.article.id,
        userId: this.$store.state.currentUser.id,
      };
      this.$http
        .post(this.$constant.baseURL + "/weiYan/saveNews/", weiYan)
        .then((res) => {
          this.$message({
            type: "success",
            message: "添加进展成功!",
          });
          this.weiYanDialogVisible = false;
          this.newsTime = "";
          this.getNews();
        })
        .catch((error) => {
          this.$message({
            message: error.message,
            type: "error",
          });
        });
    },
    getNews() {
      this.$http
        .post(this.$constant.baseURL + "/weiYan/listNews/", {
          current: 1,
          size: 9999,
          source: this.article.id,
        })
        .then((res) => {
          if (!this.$common.isEmpty(res.result[0])) {
            res.result[0].records.forEach((c) => {
              c.content = c.content.replace(
                /\n{2,}/g,
                '<div style="height: 12px"></div>'
              );
              c.content = c.content.replace(/\n/g, "<br/>");
              c.content = this.$common.faceReg(c.content);
              c.content = this.$common.pictureReg(c.content);
            });
            this.treeHoleList = res.result[0].records;
          }
        })
        .catch((error) => {
          this.$message({
            message: error.message,
            type: "error",
          });
        });
    },
    onScrollPage() {
      let scrollTop =
        document.documentElement.scrollTop || document.body.scrollTop;
      if (scrollTop < window.innerHeight / 4) {
        $(".toc").css("top", window.innerHeight / 2);
        $(".toc").css("display", "unset");
      } else if (
        scrollTop > window.innerHeight / 4 &&
        scrollTop < $("#article-like").offset().top - window.innerHeight
      ) {
        $(".toc").css("top", "100px");
        $(".toc").css("display", "unset");
      } else {
        $(".toc").css("display", "none");
      }
    },
    getTocbot() {
      let script = document.createElement("script");
      script.type = "text/javascript";
      script.src = this.$constant.tocbot;
      document.getElementsByTagName("head")[0].appendChild(script);
      // 引入成功
      script.onload = function () {
        tocbot.init({
          tocSelector: "#toc",
          contentSelector: ".entry-content",
          headingSelector: "h1, h2, h3, h4, h5, h6",
          scrollSmooth: true,
          fixedSidebarOffset: "auto",
          scrollSmoothOffset: -100,
          hasInnerContainers: true,
        });
      };
    },
    addId() {
      let headings = $(".entry-content").find("h1, h2, h3, h4, h5, h6");
      headings.attr("id", (i, id) => id || "toc-" + i);
    },
    getArticle() {
      this.$http
        .get(this.$constant.baseURL + "/article/getArticleById/", {
          id: this.id,
          flag: true,
          userId: this.$store.state.currentUser.id,
        })
        .then((res) => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.article = res.result[0].data[0];
            this.getNews();
            const md = new MarkdownIt({ breaks: true });
            this.articleContentHtml = md.render(this.article.articleContent);
            this.$nextTick(() => {
              this.highlight();
              this.addId();
              this.getTocbot();
            });
          }
        })
        .catch((error) => {
          this.$message({
            message: error.message,
            type: "error",
          });
        });
    },
    highlight() {
      let attributes = {
        autocomplete: "off",
        autocorrect: "off",
        autocapitalize: "off",
        spellcheck: "false",
        contenteditable: "false",
      };
      $("pre").each(function (i, item) {
        let preCode = $(item).children("code");
        let classNameStr = preCode[0].className;
        let classNameArr = classNameStr.split(" ");
        let lang = "";
        classNameArr.some(function (className) {
          if (className.indexOf("language-") > -1) {
            lang = className.substring(
              className.indexOf("-") + 1,
              className.length
            );
            return true;
          }
        });
        // 检测语言是否存在，不存在则自动检测
        let language = hljs.getLanguage(lang.toLowerCase());
        if (language === undefined) {
          // 启用自动检测
          let autoLanguage = hljs.highlightAuto(preCode.text());
          preCode.removeClass("language-" + lang);
          lang = autoLanguage.language;
          if (lang === undefined) {
            lang = "javascript";
          }
          preCode.addClass("language-" + lang);
        } else {
          lang = language.name;
        }
        $(item).addClass("highlight-wrap");
        $(item).attr(attributes);
        preCode
          .attr("data-rel", lang.toUpperCase())
          .addClass(lang.toLowerCase());
        // 启用代码高亮
        hljs.highlightBlock(preCode[0]);
        // 启用代码行号
        hljs.lineNumbersBlock(preCode[0]);
      });
      $("pre code").each(function (i, block) {
        $(block).attr({
          id: "hljs-" + i,
        });
        $(block).after(
          '<a class="copy-code" href="javascript:" data-clipboard-target="#hljs-' +
            i +
            '"><i class="fa fa-clipboard" aria-hidden="true"></i></a>'
        );
        new ClipboardJS(".copy-code");
      });
      if ($(".entry-content").children("table").length > 0) {
        $(".entry-content")
          .children("table")
          .wrap("<div class='table-wrapper'></div>");
      }
    },
    articleLike() {
      if (this.$common.isEmpty(this.$store.state.currentUser)) {
        this.$message({
          message: "请先登录！",
          type: "error",
        });
        return;
      }
      if (!this.article.articleLikeStatus) {
        this.$http
          .post(this.$constant.baseURL + "/article/articleLike/", {
            userId: this.$store.state.currentUser.id,
            articleLikeStatus: true,
            id: this.id,
          })
          .then((res) => {
            this.article.articleLikeStatus = 1;
            this.$message({
              message: "感谢您的点赞！",
              type: "success",
            });
          })
          .catch((error) => {
            this.$message({
              message: error.message,
              type: "error",
            });
          });
      } else {
        this.$message({
          message: "你已经点过赞啦！",
          type: "warning",
        });
      }
    },
  },
};
</script>

<style scoped>
#toc-button {
  position: fixed;
  right: 3vh;
  bottom: 16vh;
  z-index: 100;
  font-size: 23px;
  width: 30px;
  color: black;
}
#toc-button:hover {
  color: #39c5bb;
}
@media screen and (max-width: 500px) {
  #toc-button {
    right: 0vh;
    bottom: 22vh;
    color: var(--red);
  }
}
.article-background {
  position: relative;
}

.article-background::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  backdrop-filter: blur(6px);
  z-index: -1;
}

.article-head {
  height: 40vh;
  position: relative;
}

.article-image::before {
  position: absolute;
  width: 100%;
  height: 100%;
  background-color: var(--miniMask);
  content: "";
}

.article-info-container {
  position: absolute;
  bottom: 15px;
  left: 20%;
  color: var(--bigRed);
}

.article-info-news {
  position: absolute;
  bottom: 10px;
  right: 20%;
  animation: scale 1s ease-in-out infinite;
}

.article-title {
  font-size: 28px;
  margin-bottom: 15px;
}

.article-info {
  font-size: 14px;
  user-select: none;
}

.article-info i {
  margin-right: 6px;
}

.article-info span:not(:last-child) {
  margin-right: 5px;
}

.article-container {
  max-width: 900px;
  margin: 0 auto;
  padding: 40px 20px;
}

.article-update-time {
  color: var(--red);
  font-size: 14px;
  margin: 20px 0;
  user-select: none;
}

blockquote {
  line-height: 2;
  border-left: 0.2rem solid var(--blue);
  padding: 10px 1rem;
  background-color: var(--azure);
  border-radius: 4px;
  margin: 0 0 40px 0;
  user-select: none;
  color: black;
}

.article-sort {
  display: flex;
  justify-content: flex-end;
  margin-bottom: 20px;
}

.article-sort span {
  padding: 3px 10px;
  background-color: var(--blue);
  border-radius: 5px;
  font-size: 14px;
  color: white;
  margin-right: 25px;
  user-select: none;
}

.article-sort span:hover {
  background: var(--gradualRed);
}

.article-like {
  color: var(--red) !important;
}

.article-like-icon {
  font-size: 60px;
  color: var(--maxGreyFont);
  transition: all 0.5s;
  border-radius: 50%;
  margin-bottom: 20px;
}

.article-like-icon:hover {
  transform: rotate(360deg);
}

::v-deep .process-wrap .el-collapse-item__header,
::v-deep .process-wrap .el-collapse-item__wrap {
  color: var(--bigRed) !important;
  border-radius: 8px;
  padding: 0 8px;
}

::v-deep .timeline-item-time {
  color: var(--bigRed);
}

.process-wrap {
  margin: 0 0 40px;
}

.process-wrap hr {
  position: relative;
  margin: 10px auto 60px;
  border: 2px dashed var(--blue);
  overflow: visible;
}

.process-wrap hr:before {
  position: absolute;
  top: -20px;
  left: 5%;
  color: var(--red);
  content: "\e673";
  font-size: 40px;
  line-height: 1;
  transition: all 1s ease-in-out;
  font-family: iconfont;
}

.process-wrap hr:hover:before {
  left: calc(95% - 20px);
}

.process-wrap >>> .el-collapse-item__header {
  border-bottom: unset;
  font-size: 20px;
  background-color: var(--background);
  color: #39c5bb;
}

.process-wrap >>> .el-collapse-item__wrap {
  background-color: var(--background);
}

.process-wrap .el-collapse {
  border-top: unset;
  border-bottom: unset;
}

.process-wrap >>> .el-collapse-item__wrap {
  border-bottom: unset;
}

@media screen and (max-width: 700px) {
  .article-info-container {
    left: 20px;
    max-width: 320px;
  }

  .article-info-news {
    right: 20px;
  }
}

::v-deep .dialog {
  border-radius: 14px;
  overflow: scroll;
  box-shadow: 0 14px 28px rgba(0, 0, 0, 0.25), 0 10px 10px rgba(0, 0, 0, 0.22);
  height: 450px;
}

::v-deep .dialog::-webkit-scrollbar {
  width: 0px;
}
</style>
