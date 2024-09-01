<template>
  <div>
    <div
      style="animation: header-effect 2s"
      :style="{ background: `${$store.state.changeBg}` }"
      class="background-image background-image-changeBg"
    ></div>
    <!-- 两句诗 -->
    <div class="my-animation-slide-top">
      <twoPoem :isHitokoto="false"></twoPoem>
    </div>
    <div style="background: var(--background); animation: hideToShow 2.5s">
      <div>
        <treeHole
          :treeHoleList="treeHoleList"
          :avatar="
            !$common.isEmpty($store.state.currentUser)
              ? $store.state.currentUser.avatar
              : $store.state.webInfo.avatar
          "
          @launch="launch"
          @deleteTreeHole="deleteTreeHole"
        >
        </treeHole>
        <proPage
          :current="pagination.current"
          :size="pagination.size"
          :total="pagination.total"
          :buttonSize="3"
          :color="$constant.commentPageColor"
          @toPage="toPage"
        >
        </proPage>
      </div>
    </div>
    <el-dialog
      title="微言"
      :visible.sync="weiYanDialogVisible"
      width="40%"
      :before-close="handleClose"
      :append-to-body="true"
      destroy-on-close
      center
      custom-class="dialog"
    >
      <div>
        <div class="myCenter" style="padding-bottom: 20px">
          <el-radio-group v-model="isPublic">
            <el-radio-button :label="true">公开</el-radio-button>
            <el-radio-button :label="false">私密</el-radio-button>
          </el-radio-group>
        </div>
        <commentBox :disableGraffiti="true" @submitComment="submitWeiYan">
        </commentBox>
      </div>
    </el-dialog>
  </div>
</template>

<script>
const twoPoem = () => import("./common/twoPoem");
const treeHole = () => import("./common/treeHole");
const proPage = () => import("./common/proPage");
const commentBox = () => import("./common/commentBox");
export default {
  components: {
    twoPoem,
    treeHole,
    proPage,
    commentBox,
  },
  data() {
    return {
      treeHoleList: [],
      pagination: {
        current: 1,
        size: 10,
        total: 0,
        userId: this.$store.state.currentUser.id,
      },
      weiYanDialogVisible: false,
      isPublic: true,
      showFooter: false,
    };
  },
  created() {
    this.getWeiYan();
  },
  methods: {
    toPage(page) {
      this.pagination.current = page;
      window.scrollTo({
        top: 240,
        behavior: "smooth",
      });
      this.getWeiYan();
    },
    launch() {
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
      this.weiYanDialogVisible = true;
    },
    handleClose() {
      this.weiYanDialogVisible = false;
    },
    submitWeiYan(content) {
      let weiYan = {
        content: content,
        isPublic: this.isPublic,
        userId: this.$store.state.currentUser.id,
      };
      this.$http
        .post(
          this.$constant.baseURL + "/weiYan/saveWeiYan/",
          weiYan,
          false,
          true,
          true
        )
        .then((res) => {
          this.getWeiYan();
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
                title: "可以啦🍨",
                message: "删除成功！",
                type: "success",
                offset: 50,
                position: "top-left",
              });
              this.pagination.current = 1;
              this.getWeiYan();
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
            title: "可以啦🍨",
            message: "已取消删除！",
            type: "success",
            offset: 50,
            position: "top-left",
          });
        });
    },
    getWeiYan() {
      this.$http
        .post(this.$constant.baseURL + "/weiYan/listWeiYan/", this.pagination)
        .then((res) => {
          this.showFooter = false;
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
            this.pagination.total = res.result[0].total;
          }
          this.$nextTick(() => {
            this.showFooter = true;
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
    },
  },
};
</script>

<style lang="scss" scoped>
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
