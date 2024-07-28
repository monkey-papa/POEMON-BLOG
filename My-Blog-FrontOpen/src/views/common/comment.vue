<template>
  <div>
    <!-- 评论框 -->
    <div style="margin-bottom: 40px">
      <div class="comment-head">
        <i
          class="el-icon-edit-outline"
          style="font-weight: bold; font-size: 22px"
        ></i>
        留言
      </div>
      <div>
        <!-- 文字评论 -->
        <div v-show="!isGraffiti">
          <commentBox
            @showGraffiti="isGraffiti = !isGraffiti"
            @submitComment="submitComment"
          >
          </commentBox>
        </div>
        <!-- 画笔 -->
        <!-- <div v-show="isGraffiti">
          <graffiti @showComment="isGraffiti = !isGraffiti" @addGraffitiComment="addGraffitiComment">
          </graffiti>
        </div> -->
      </div>
    </div>
    <!-- 评论内容 -->
    <div v-if="comments.length > 0">
      <!-- 评论数量 -->
      <div class="commentInfo-title">
        <span style="font-size: 1.15rem">Comments | </span>
        <span>{{ total }} 条留言</span>
      </div>
      <!-- 评论详情 -->
      <div
        id="comment-content"
        class="commentInfo-detail"
        v-for="(item, index) in comments"
        :key="index"
      >
        <!-- 头像 -->
        <el-avatar
          shape="square"
          class="commentInfo-avatar"
          :size="35"
          :src="item.avatar || $store.state.webInfo.avatar"
        ></el-avatar>
        <div style="flex: 1; padding-left: 12px">
          <!-- 评论人信息 -->
          <div style="display: flex; justify-content: space-between">
            <div class="commentInfo-left">
              <span class="commentInfo-username">{{ item.username }}</span>
              <span class="commentInfo-master" v-if="item.userType === 0"
                >店长</span
              >
              <span class="commentInfo-other">{{
                $common.getDateDiff(item.createTime)
              }}</span>
            </div>
            <div class="commentInfo-reply" @click="replyDialog(item, item)">
              <span v-if="item.childComments[0].total > 0"
                >{{ item.childComments[0].total }} </span
              ><span><i class="iconfont icon-chat--fill"></i></span>
            </div>
          </div>
          <!-- 评论内容 -->
          <div class="commentInfo-content">
            <span v-html="item.commentContent"></span>
          </div>
          <!-- 回复模块 -->
          <div
            v-if="
              !$common.isEmpty(item.childComments[0]) &&
              !$common.isEmpty(item.childComments[0].records)
            "
          >
            <div
              class="commentInfo-detail"
              v-for="(childItem, i) in item.childComments[0].records"
              :key="i"
            >
              <!-- 头像 -->
              <el-avatar
                shape="square"
                class="commentInfo-avatar"
                :size="30"
                :src="childItem.avatar || $store.state.webInfo.avatar"
              ></el-avatar>
              <div style="flex: 1; padding-left: 12px">
                <!-- 评论人信息 -->
                <div style="display: flex; justify-content: space-between">
                  <div class="commentReply-left">
                    <span class="commentInfo-username-small">{{
                      childItem.username
                    }}</span>
                    <span
                      class="commentInfo-master"
                      v-if="childItem.userType === 0"
                      >店长</span
                    >
                    <span class="commentInfo-other">{{
                      $common.getDateDiff(childItem.createTime)
                    }}</span>
                  </div>
                  <div>
                    <span
                      class="commentInfo-reply"
                      @click="replyDialog(childItem, item)"
                      ><i class="iconfont icon-chat--fill"></i
                    ></span>
                  </div>
                </div>
                <!-- 评论内容 -->
                <div class="commentInfo-content">
                  <template
                    v-if="
                      childItem.parentCommentId !== item.id &&
                      childItem.parentUserId !== childItem.userId
                    "
                  >
                    <span style="color: var(--blue2)"
                      >@{{ childItem.parentUsername }} </span
                    >:
                  </template>
                  <span v-html="childItem.commentContent"></span>
                </div>
              </div>
            </div>
            <!-- 分页 -->
            <div
              class="pagination-wrap"
              v-if="
                item.childComments[0].records.length <
                item.childComments[0].total
              "
            >
              <div class="pagination" @click="toChildPage(item)">展开</div>
            </div>
          </div>
        </div>
      </div>
      <!-- 分页 -->
      <proPage
        :current="pagination.current"
        :size="pagination.size"
        :total="pagination.total"
        :buttonSize="6"
        :color="$constant.commentPageColor"
        @toPage="toPage"
      >
      </proPage>
    </div>
    <div v-else class="myCenter" style="color: var(--red)">
      <i>来发第一个留言啦~</i>
    </div>
    <el-dialog
      title="留言"
      :visible.sync="replyDialogVisible"
      width="30%"
      :before-close="handleClose"
      :append-to-body="true"
      destroy-on-close
      center
      custom-class="dialog"
    >
      <div>
        <commentBox :disableGraffiti="true" @submitComment="submitReply">
        </commentBox>
      </div>
    </el-dialog>
  </div>
</template>
<script>
// const graffiti = () => import("./graffiti");
const commentBox = () => import("./commentBox");
const proPage = () => import("./proPage");
export default {
  components: {
    // graffiti,
    commentBox,
    proPage,
  },
  props: {
    source: {
      type: Number, // 0 树洞评论 1 恋爱留言 其他 文章评论
    },
    type: {
      type: String,
    },
  },
  data() {
    return {
      isGraffiti: false,
      total: 0,
      replyDialogVisible: false,
      floorComment: {},
      replyComment: {},
      comments: [],
      pagination: {
        current: 1,
        size: 5,
        total: 0,
        source: this.source,
        commentType: this.type,
        floorCommentId: null,
        csize: 5,
      },
      prohibitedWordsList: [],
      userPagination: {
        current: 1,
        size: 9999,
        total: 0,
        searchKey: "",
        userStatus: null,
        userType: null,
      },
    };
  },
  created() {
    this.getProhibitedWordsList();
    this.getComments(this.pagination);
    if (this.$store.state.userList.length === 0) {
      this.getUsers();
    }
  },
  methods: {
    getUsers() {
      this.$http
        .post(
          this.$constant.baseURL + "/admin/user/list/",
          this.userPagination,
          true,
          false
        )
        .then((res) => {
          if (!this.$common.isEmpty(res.result[0].data)) {
            this.$store.commit("userList", res.result[0].data);
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
    //跳转到指定页数
    toPage(page) {
      this.pagination.current = page;
      window.scrollTo({
        top: document.getElementById("comment-content").offsetTop,
      });
      this.getComments(this.pagination);
    },
    //展开
    toChildPage(floorComment) {
      //floorComment 传过来的一个一级评论对象
      floorComment.childComments[0].current += 1;
      let pagination = {
        current: floorComment.childComments[0].current,
        size: 5,
        total: 0,
        source: this.source,
        commentType: this.type,
        floorCommentId: floorComment.id,
        csize: 5,
      };
      this.getComments(pagination, floorComment, true);
    },
    emoji(comments, flag) {
      //传过来的true对应着comments是this.comments，false对应着comments是floorComment.childComments[0].records
      comments.forEach((c) => {
        c.commentContent = c.commentContent.replace(/\n/g, "<br/>");
        // 表情包转换
        c.commentContent = this.$common.faceReg(c.commentContent);
        // 图片转换
        c.commentContent = this.$common.pictureReg(c.commentContent);
        //传过来的true对应着comments是this.comments，false是floorComment.childComments[0].records
        if (flag) {
          if (
            !this.$common.isEmpty(c.childComments) &&
            !this.$common.isEmpty(c.childComments[0].records)
          ) {
            c.childComments[0].records.forEach((cc) => {
              c.commentContent = c.commentContent.replace(/\n/g, "<br/>");
              cc.commentContent = this.$common.faceReg(cc.commentContent);
              cc.commentContent = this.$common.pictureReg(cc.commentContent);
            });
          }
        }
      });
    },
    getComments(pagination, floorComment = {}, isToPage = false) {
      this.$http
        .post(this.$constant.baseURL + "/comment/listComment/", pagination)
        .then((res) => {
          if (
            !this.$common.isEmpty(res.result[0]) &&
            !this.$common.isEmpty(res.result[0].data)
          ) {
            if (this.$common.isEmpty(floorComment)) {
              this.comments = res.result[0].data;
              // pagination.total用来判断分页
              pagination.total =
                res.result[0].total - res.result[0].parenttotal;
              this.total = res.result[0].total;
              this.emoji(this.comments, true);
            } else {
              //floorComment 有值的情况是回复或者展开的情况
              //floorComment 传过来的都是一个一级评论对象
              //点击回复的时候
              if (isToPage === false) {
                //将新返回的二级列表返回给一级目录中，通过数据改变视图
                floorComment.childComments[0].records = res.result[0].data;
                floorComment.childComments[0].total = res.result[0].parenttotal;
              } else {
                //点击展开评论的时候
                floorComment.childComments[0].total = res.result[0].parenttotal;
                floorComment.childComments[0].records =
                  floorComment.childComments[0].records.concat(
                    res.result[0].data
                  );
              }
              this.emoji(floorComment.childComments[0].records, false);
            }
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
    addGraffitiComment(graffitiComment) {
      this.submitComment(graffitiComment);
    },
    submitComment(commentContent) {
      let comment = {
        source: this.source,
        type: this.type,
        commentContent: commentContent,
        parentCommentId: 0,
        userId: this.$store.state.currentUser.id,
        floorCommentId: null,
        parentUserId: null,
        likeCount: 0,
        commentInfo: "",
      };
      let newArr = [];
      let list = [];
      let word = comment.commentContent.split("");
      this.prohibitedWordsList.forEach((e) => list.push(e.message));
      for (let i = 0; i < list.length; i++) {
        for (let j = 0; j < word.length; j++) {
          if (word[j] === list[i]) {
            newArr.push(word[j]);
          }
        }
      }
      if (newArr.length > 0) {
        this.$notify({
          type: "error",
          title: "可恶🤬",
          message: "你发的评论带有违禁词！请发一条友好的评论~~~",
          position: "top-left",
          offset: 50,
        });
        return;
      }
      this.$http
        .post(
          this.$constant.baseURL + "/admin/comment/boss/addComment/",
          comment
        )
        .then((res) => {
          // 评论博客主人
          this.$http
            .post(this.$constant.baseURL + "/codeComment/", {
              email: "1816298537@qq.com",
              comment: commentContent,
              name: this.$store.state.currentUser.username,
            })
            .then((res) => {
              this.$notify({
                type: "success",
                title: "可以啦🍨",
                message: "评论成功！同时也给当事人发送了一封邮件~",
                position: "top-left",
                offset: 50,
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
          this.pagination = {
            current: 1,
            size: 5,
            total: 0,
            source: this.source,
            commentType: this.type,
            floorCommentId: null,
            csize: 5,
          };
          this.getComments(this.pagination);
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
    //commentContent是回复弹窗输入的值
    submitReply(commentContent) {
      // replyComment 上面的回复是item，下面的是childItem
      // floorComment 是一个一级评论对象
      let comment = {
        source: this.source,
        type: this.type,
        commentContent: commentContent,
        parentCommentId: this.replyComment.id,
        userId: this.$store.state.currentUser.id,
        likeCount: 0,
        floorCommentId: this.floorComment.id,
        parentUserId: this.replyComment.userId,
        commentInfo: "",
      };
      const name = this.$store.state.userList.find(
        (e) => e.id === comment.parentUserId
      ).username;
      const email = this.$store.state.userList.find(
        (e) => e.id === comment.parentUserId
      ).email;
      //floorComment 一级评论对象
      let floorComment = this.floorComment;
      floorComment.childComments[0].current = 1;
      let newArr = [];
      let list = [];
      let word = comment.commentContent.split("");
      this.prohibitedWordsList.forEach((e) => list.push(e.message));
      for (let i = 0; i < list.length; i++) {
        for (let j = 0; j < word.length; j++) {
          if (word[j] === list[i]) {
            newArr.push(word[j]);
          }
        }
      }
      if (newArr.length > 0) {
        this.$notify({
          type: "error",
          title: "可恶🤬",
          message: "你发的评论带有违禁词！请发一条友好的评论~~~",
          position: "top-left",
          offset: 50,
        });
        return;
      }
      this.$http
        .post(
          this.$constant.baseURL + "/admin/comment/boss/addComment/",
          comment
        )
        .then((res) => {
          // 回复时发邮件
          this.$http
            .post(this.$constant.baseURL + "/codeComment/", {
              email,
              comment: commentContent,
              name,
            })
            .then((res) => {
              this.$notify({
                type: "success",
                title: "可以啦🍨",
                message: "评论成功！同时也给当事人发送了一封邮件~",
                position: "top-left",
                offset: 50,
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
          let pagination = {
            current: 1,
            size: 5,
            total: 0,
            source: this.source,
            commentType: this.type,
            floorCommentId: floorComment.id,
          };
          this.getComments(pagination, floorComment);
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
      this.handleClose();
    },
    //floorComment 传过来的一个一级评论对象，comment上面的回复是item，下面的是childItem
    replyDialog(comment, floorComment) {
      this.replyComment = comment;
      this.floorComment = floorComment;
      this.replyDialogVisible = true;
    },
    handleClose() {
      this.replyDialogVisible = false;
      this.floorComment = {};
      this.replyComment = {};
    },
    getProhibitedWordsList() {
      this.$http
        .post(
          this.$constant.baseURL + "/prohibitedWords/list/",
          {
            current: 1,
            size: 10,
            total: 0,
            searchKey: "",
          },
          true
        )
        .then((res) => {
          if (res.result[0]) {
            this.prohibitedWordsList = res.result[0].data;
            this.$emit("getProhibitedWordsList", this.prohibitedWordsList);
          } else {
            this.$notify({
              type: "error",
              title: "可恶🤬",
              message: res.data.msg,
              position: "top-left",
              offset: 50,
            });
          }
        });
    },
  },
};
</script>
<style lang="scss" scoped>
.comment-head {
  display: flex;
  align-items: center;
  font-size: 20px;
  font-weight: bold;
  margin: 40px 0 20px 0;
  user-select: none;
  color: var(--red);
}
.commentInfo-title {
  margin-bottom: 20px;
  color: var(--bigRed);
  user-select: none;
}
.commentInfo-detail {
  display: flex;
}
.commentInfo-avatar {
  border-radius: 5px;
}
.commentInfo-username {
  color: var(--red);
  font-size: 18px;
  font-weight: 400;
  margin-right: 5px;
  &-small {
    color: var(--red);
    font-size: 16px;
    font-weight: 400;
    margin-right: 5px;
  }
}
.commentInfo-master {
  color: var(--bigRed);
  border: 1px solid var(--bigRed);
  border-radius: 0.2rem;
  font-size: 14px;
  padding: 2px 4px;
  margin-right: 5px;
}
.commentInfo-other {
  font-size: 14px;
  color: var(--darkBlue);
  user-select: none;
}
.commentInfo-reply {
  white-space: nowrap;
  max-height: 22px;
  font-size: 12px;
  user-select: none;
  color: var(--white1);
  background: var(--red3);
  border-radius: 0.2rem;
  padding: 3px 6px;
  transition: all 0.3s ease;
  border: 1px solid var(--gray1);
  &:hover {
    color: var(--white);
    background: var(--red);
    border-color: var(--gray4);
  }
}
.commentInfo-content {
  margin: 15px 0 25px;
  padding: 18px 20px;
  background: var(--favoriteBg);
  border-radius: 12px;
  color: var(--fontColor);
  word-break: break-word;
  transition: all 0.3s ease;
  border: 1px solid var(--gray1);
  &:hover {
    border-color: var(--gray4);
  }
}
.pagination-wrap {
  display: flex;
  justify-content: center;
  margin-bottom: 10px;
}
.pagination {
  padding: 6px 20px;
  border: 1px solid var(--red1);
  border-radius: 3rem;
  color: var(--red1);
  user-select: none;
  text-align: center;
  font-size: 12px;
  &:hover {
    border: 1px solid var(--blue2);
    color: var(--orange2);
    box-shadow: 0 0 5px var(--blue2);
  }
}
::v-deep .dialog {
  border-radius: 14px;
  overflow: scroll;
  box-shadow: 0 14px 28px var(--mask), 0 10px 10px var(--miniMask);
  height: 400px;
  &::-webkit-scrollbar {
    width: 0px;
  }
}

@media screen and (max-width: 515px) {
  .commentInfo-left {
    min-width: 196px;
  }
  .commentReply-left {
    min-width: 210px;
  }
}
</style>
