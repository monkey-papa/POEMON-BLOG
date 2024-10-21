<template>
  <div v-if="!$common.isEmpty(article)">
    <!-- 首页图片 -->
    <div
      style="animation: header-effect 2s"
      :style="{ background: `url(${article.articleCover})` }"
      class="background-image background-image-changeBg blur-filter"
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
      <div style="display: flex; justify-content: space-between">
        <div class="article-container my-animation-slide-bottom shadow-box">
          <div class="post-ai">
            <div class="ai-title">
              <a class="ai-title-left" data-pjax-state="">
                <div @click="getSummary" class="ai-title-text">
                  <a class="icon"
                    ><i class="iconfont icon-jiqirenjiankong"></i></a
                  ><span class="text">{{
                    article.summary || summary ? "文章摘要" : "点我生成摘要"
                  }}</span
                  ><i class="el-icon el-icon-arrow-right"></i></div
              ></a>
              <div class="ai-tag" id="ai-tag">续写</div>
            </div>
            <div v-if="article.summary || summary" class="ai-explanation">
              <p class="text">{{ article.summary || summary }}</p>
              <p class="cover">
                <span class="text cover-text">
                  {{ article.summary || summary }}</span
                >
              </p>
            </div>
            <el-skeleton :rows="3" animated v-if="loading" />
            <div class="ai-bottom">
              <div class="ai-tips">
                此内容根据文章生成，并经过人工审核，仅用于文章内容的解释与总结
              </div>
            </div>
          </div>
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
            <div style="color: var(--blue2)">作者：{{ article.username }}</div>
            <div style="color: var(--blue2)">版权声明：转载请注明文章出处</div>
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
            <comment :type="'article'" :source="article.id"></comment>
          </div>
        </div>
        <!-- 侧边栏 -->
        <div class="aside-content" v-if="!$common.mobile()">
          <myAside @selectSort="selectSort"></myAside>
        </div>
      </div>
      <div
        v-show="!this.$common.mobile() && !mobile"
        id="toc"
        class="toc"
      ></div>
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
const comment = () => import("./common/comment");
const process = () => import("./common/process");
const commentBox = () => import("./common/commentBox");
const myAside = () => import("./myAside");
import MarkdownIt from "markdown-it";
import ColorThief from "colorthief";
export default {
  components: {
    comment,
    commentBox,
    process,
    myAside,
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
      summary: "",
      loading: false,
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
  beforeRouteLeave(to, from, next) {
    const root = document.querySelector(":root");
    root.style.setProperty("--themeColor", localStorage.getItem("themeColor"));
    this.$common.getThemeRgb();
    next();
  },
  methods: {
    getSummary() {
      if (this.article.summary || this.summary) {
        return;
      }
      this.loading = true;
      const message = this.article.articleContent;
      this.$http
        .post(
          this.$constant.baseURL + "/summary",
          { message, articleId: this.article.id },
          false,
          true,
          false
        )
        .then((res) => {
          this.summary = res.summary;
          this.loading = false;
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
        this.$notify({
          type: "error",
          title: "可恶🤬",
          message: "请先登录！",
          position: "top-left",
          offset: 50,
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
            .get(
              this.$constant.baseURL + "/weiYan/deleteWeiYan/",
              { id: id },
              false,
              true
            )
            .then((res) => {
              this.$notify({
                type: "success",
                title: "可以啦🍨",
                message: "删除成功!",
                position: "top-left",
                offset: 50,
              });
              this.getNews();
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
        })
        .catch(() => {
          this.$notify({
            type: "success",
            title: "可以啦🍨",
            message: "已取消删除!",
            position: "top-left",
            offset: 50,
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
        .post(
          this.$constant.baseURL + "/weiYan/saveNews/",
          weiYan,
          false,
          true,
          true
        )
        .then((res) => {
          this.$notify({
            title: "可以啦🍨",
            message: "添加进展成功!",
            type: "success",
            offset: 50,
            position: "top-left",
          });
          this.weiYanDialogVisible = false;
          this.newsTime = "";
          this.getNews();
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
          this.$notify({
            type: "error",
            title: "可恶🤬",
            message: error.message,
            position: "top-left",
            offset: 50,
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
    getColorFromImage(articleCover) {
      // 创建一个新的img元素
      const img = document.createElement("img");
      img.src = articleCover;
      img.setAttribute("crossOrigin", "");
      // 创建一个ColorThief实例
      const colorThief = new ColorThief();
      // 当图片加载完成后，提取颜色
      img.onload = () => {
        // 提取主色
        const dominantColor = colorThief.getColor(img);
        const root = document.querySelector(":root");
        const rgbToHex = (rgb) => {
          const [r, g, b] = rgb.map((num) => parseInt(num, 10));
          const toHex = (c) => {
            const hex = c.toString(16);
            return hex.length == 1 ? "0" + hex : hex;
          };
          return `#${toHex(r)}${toHex(g)}${toHex(b)}`;
        };
        const color = rgbToHex(dominantColor);
        root.style.setProperty("--themeColor", color);
        this.$common.getThemeRgb();
        // 销毁img元素
        img.remove();
      };
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
            this.getColorFromImage(this.article.articleCover);
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
          this.$notify({
            type: "error",
            title: "可恶🤬",
            message: error.message,
            position: "top-left",
            offset: 50,
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
        this.$notify({
          type: "error",
          title: "可恶🤬",
          message: "请先登录！",
          position: "top-left",
          offset: 50,
        });
        return;
      }
      if (!this.article.articleLikeStatus) {
        this.$http
          .post(
            this.$constant.baseURL + "/article/articleLike/",
            {
              userId: this.$store.state.currentUser.id,
              articleLikeStatus: true,
              id: this.id,
            },
            false,
            true,
            true
          )
          .then((res) => {
            this.article.articleLikeStatus = 1;
            this.$notify({
              title: "可以啦🍨",
              message: "感谢您的点赞！",
              type: "success",
              offset: 50,
              position: "top-left",
            });
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
      } else {
        this.$notify({
          type: "warning",
          title: "淘气👻",
          message: "你已经点过赞啦！",
          position: "top-left",
          offset: 50,
        });
      }
    },
  },
};
</script>

<style lang="scss" scoped>
@keyframes move-show {
  to {
    --p1: 100%;
  }
}
.blur-filter {
  filter: blur(30px);
}
#toc-button {
  position: fixed;
  right: 4vh;
  bottom: 14vh;
  z-index: 100;
  font-size: 23px;
  width: 23px;
  height: 23px;
  line-height: 23px;
  color: var(--black);
  text-align: center;
  &:hover {
    color: var(--green2);
  }
}
.aside-content {
  width: 300px;
}
.article-background {
  position: relative;
  padding: 10px 130px 0;
  &::before {
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    backdrop-filter: blur(6px);
    z-index: -1;
  }
}
.article-head {
  height: 40vh;
  position: relative;
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
  color: var(--bigRed5);
  span:not(:last-child) {
    margin-right: 5px;
  }
  &-news {
    right: 20px;
  }
}
.article-container {
  padding: 40px 20px;
  border-radius: 8px;
  border: 2px dashed var(--gray1);
  width: calc(100% - 310px);
  transition: all 0.3s ease;
  .post-ai {
    display: flex;
    flex-direction: column;
    border-radius: 8px;
    border: 1px dashed var(--red);
    background: var(--background);
    padding: 12px;
    margin-bottom: 12px;
    .ai-title {
      display: flex;
      justify-content: space-between;
      align-items: center;
      .ai-title-left {
        display: flex;
        align-items: center;
        .ai-title-icon {
          font-size: 20px;
          color: var(--red);
        }
        .ai-title-text {
          font-size: 14px;
          color: var(--red);
          display: flex;
          align-items: center;
          .el-icon-arrow-right {
            font-size: 15px;
          }
          .icon {
            display: flex;
            justify-content: center;
            align-items: center;
            width: 24px;
            height: 24px;
            border-radius: 50%;
            background: var(--red);
            margin-right: 8px;
            .icon-jiqirenjiankong {
              font-size: 14px;
              color: var(--favoriteBg);
              transition: all 0.3s ease;
              &:hover {
                opacity: 0.8;
              }
            }
          }
          .text {
            margin-right: 4px;
            transition: all 0.3s ease;
            &:hover {
              opacity: 0.8;
            }
          }
        }
        .ai-link {
          font-size: 20px;
          color: var(--red);
        }
      }
      .ai-tag {
        padding: 6px 8px;
        background-color: var(--red);
        color: var(--white);
        border-radius: 12px;
        font-size: 12px;
        transition: all 0.3s ease;
        &:hover {
          opacity: 0.8;
        }
      }
    }
    .ai-explanation {
      padding: 8px 12px;
      font-size: 15px;
      margin-top: 12px;
      border-radius: 8px;
      border: 1px solid var(--myAsideBorderColor);
      color: var(--fontColor);
      background: var(--white3);
      font-size: 14px;
      line-height: 1.4;
      position: relative;
      .text {
        margin: 0;
      }
      .cover {
        position: absolute;
        top: 8px;
        left: 12px;
        padding-right: 12px;
        margin: 0;
      }
      .cover-text {
        --p1: 0%;
        background: linear-gradient(
          to right,
          var(--black2) var(--p1),
          var(--white3) calc(var(--p1) + 20px)
        );
        color: transparent;
        animation: move-show 8s linear forwards;
      }
    }
    .ai-bottom {
      padding: 0 12px;
      font-size: 12px;
      margin-top: 12px;
    }
  }
  &:hover {
    border-color: var(--red);
  }
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
  background-color: var(--favoriteBg);
  border-radius: 4px;
  margin: 0 0 40px 0;
  user-select: none;
  color: var(--black);
}
.article-sort {
  display: flex;
  justify-content: flex-end;
  margin-bottom: 20px;
  span {
    padding: 3px 10px;
    background-color: var(--blue);
    border-radius: 5px;
    font-size: 14px;
    color: var(--white1);
    margin-right: 25px;
    user-select: none;
    &:hover {
      color: var(--white2);
      background: var(--gradualRed);
    }
  }
}
.article-like {
  color: var(--red) !important;
  &-icon {
    font-size: 60px;
    color: var(--fontColor);
    transition: all 0.5s ease;
    border-radius: 50%;
    margin-bottom: 20px;
    &:hover {
      transform: rotate(360deg);
    }
  }
}
::v-deep .process-wrap .el-collapse-item__header,
::v-deep .process-wrap .el-collapse-item__wrap {
  color: var(--bigRed);
  border-radius: 8px;
  padding: 0 8px;
}
::v-deep .timeline-item-time {
  display: flex;
  color: var(--bigRed1);
  align-items: center;
}
.process-wrap {
  margin: 0 0 40px;
  hr {
    position: relative;
    margin: 20px auto 60px;
    border: 2px dashed var(--blue2);
    overflow: visible;
    &:before {
      position: absolute;
      top: -20px;
      left: 5%;
      color: var(--red);
      content: "\e673";
      font-size: 40px;
      line-height: 1;
      transition: all 1s ease;
      font-family: iconfont;
    }
    &:hover:before {
      left: calc(95% - 20px);
    }
  }
  ::v-deep {
    .el-collapse-item__header {
      border-bottom: unset;
      font-size: 20px;
      background-color: var(--background);
      color: var(--green2);
    }
    .el-collapse-item__wrap {
      background-color: var(--background);
    }
  }
  .el-collapse {
    border-top: unset;
    border-bottom: unset;
  }
  ::v-deep .el-collapse-item__wrap {
    border-bottom: unset;
  }
}
@media screen and (max-width: 500px) {
  #toc-button {
    right: 36px;
    bottom: 13vh;
    color: var(--red);
  }
}
@media screen and (max-width: 1278px) {
  .article-background {
    padding: 10px 40px 0;
  }
}
@media screen and (max-width: 1200px) {
  .aside-content {
    display: none;
  }
  .article-container {
    width: 100%;
  }
  .article-background {
    padding: 10px 20px 0;
  }
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
  box-shadow: 0 14px 28px var(--mask), 0 10px 10px var(--miniMask);
  height: 450px;
  &::-webkit-scrollbar {
    width: 0px;
  }
}
</style>
