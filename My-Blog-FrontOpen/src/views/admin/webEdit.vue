<template>
  <div>
    <!-- 基础信息 -->
    <div>
      <el-tag effect="dark" class="my-tag">
        <img style="vertical-align: -3px" src="../../assets/svg/edit.svg" />
        基础信息
      </el-tag>
      <el-form
        :model="webInfo"
        :rules="rules"
        ref="ruleForm"
        label-width="100px"
        class="demo-ruleForm"
      >
        <el-form-item label="网站名称" prop="webName">
          <el-input v-model="webInfo.webName"></el-input>
        </el-form-item>
        <el-form-item label="网站标题" prop="webTitle">
          <el-input v-model="webInfo.webTitle"></el-input>
        </el-form-item>
        <el-form-item label="页脚" prop="footer">
          <el-input v-model="webInfo.footer"></el-input>
        </el-form-item>
        <el-form-item
          v-hasPermi="['user:visit:read']"
          label="状态"
          prop="status"
        >
          <el-switch
            @click.native="changeWebStatus(webInfo)"
            v-model="webInfo.status"
          ></el-switch>
        </el-form-item>
        <el-form-item label="背景" prop="backgroundImage">
          <div style="display: flex">
            <el-input v-model="webInfo.backgroundImage"></el-input>
            <el-image
              v-if="webInfo.avatar"
              lazy
              class="table-td-thumb"
              style="margin-left: 10px"
              :preview-src-list="[webInfo.backgroundImage]"
              :src="webInfo.backgroundImage"
              fit="cover"
            ></el-image>
          </div>
          <uploadPicture
            v-hasPermi="['user:visit:read']"
            :ResourceType="'webBackgroundImage'"
            :isAdmin="true"
            style="margin-top: 15px"
            @addPicture="addBackgroundImage"
            :maxSize="5"
            :maxNumber="1"
          ></uploadPicture>
        </el-form-item>
        <el-form-item label="默认头像" prop="avatar">
          <div style="display: flex">
            <el-input v-model="webInfo.avatar"></el-input>
            <el-image
              v-if="webInfo.avatar"
              lazy
              class="table-td-thumb"
              style="margin-left: 10px"
              :preview-src-list="[webInfo.avatar]"
              :src="webInfo.avatar"
              fit="cover"
            ></el-image>
          </div>
          <uploadPicture
            v-hasPermi="['user:visit:read']"
            :ResourceType="'userAvatar'"
            :isAdmin="true"
            style="margin-top: 15px"
            @addPicture="addAvatar"
            :maxSize="5"
            :maxNumber="1"
          ></uploadPicture>
        </el-form-item>
        <el-form-item label="提示" prop="waifuJson">
          <div style="display: flex">
            <el-input
              :disabled="disabled"
              :rows="6"
              type="textarea"
              v-model="webInfo.waifuJson"
            ></el-input>
            <i class="el-icon-edit my-icon" @click="disabled = !disabled"></i>
          </div>
        </el-form-item>
      </el-form>
      <div class="myCenter" style="margin-bottom: 22px">
        <el-button
          v-hasPermi="['user:visit:read']"
          type="primary"
          @click="submitForm('ruleForm')"
          >保存</el-button
        >
      </div>
    </div>
    <!-- 公告 -->
    <div>
      <el-tag effect="dark" class="my-tag">
        <img style="vertical-align: -3px" src="../../assets/svg/edit.svg" />
        公告
      </el-tag>
      <el-tag
        :key="i"
        v-for="(notice, i) in notices"
        closable
        :disable-transitions="false"
        @close="handleClose(notices, notice)"
      >
        {{ notice }}
      </el-tag>
      <el-input
        class="input-new-tag"
        v-if="inputNoticeVisible"
        v-model="inputNoticeValue"
        ref="saveNoticeInput"
        size="small"
        @keyup.enter.native="handleInputNoticeConfirm"
        @blur="handleInputNoticeConfirm"
      >
      </el-input>
      <el-button
        v-else
        class="button-new-tag"
        size="small"
        @click="showNoticeInput()"
        >+ 公告</el-button
      >
      <div class="myCenter" style="margin-bottom: 22px">
        <el-button
          v-hasPermi="['user:visit:read']"
          type="primary"
          @click="saveNotice()"
          >保存</el-button
        >
      </div>
    </div>
    <!-- 随机名称 -->
    <div>
      <el-tag effect="dark" class="my-tag">
        <img style="vertical-align: -3px" src="../../assets/svg/edit.svg" />
        随机名称
      </el-tag>
      <el-tag
        :key="i"
        effect="dark"
        v-for="(name, i) in randomName"
        closable
        :disable-transitions="false"
        :type="types[Math.floor(Math.random() * 5)]"
        @close="handleClose(randomName, name)"
      >
        {{ name }}
      </el-tag>
      <el-input
        class="input-new-tag"
        v-if="inputRandomNameVisible"
        v-model="inputRandomNameValue"
        ref="saveRandomNameInput"
        size="small"
        @keyup.enter.native="handleInputRandomNameConfirm"
        @blur="handleInputRandomNameConfirm"
      >
      </el-input>
      <el-button
        v-else
        class="button-new-tag"
        size="small"
        @click="showRandomNameInput"
        >+ 随机名称</el-button
      >
      <div class="myCenter" style="margin-bottom: 22px">
        <el-button
          v-hasPermi="['user:visit:read']"
          type="primary"
          @click="saveRandomName()"
          >保存</el-button
        >
      </div>
    </div>
    <!-- 随机头像 -->
    <div>
      <el-tag effect="dark" class="my-tag">
        <img style="vertical-align: -3px" src="../../assets/svg/edit.svg" />
        随机头像
      </el-tag>
      <div :key="i" style="display: flex" v-for="(avatar, i) in randomAvatar">
        <el-tag
          style="white-space: normal; height: unset"
          closable
          :disable-transitions="false"
          @close="handleClose(randomAvatar, avatar)"
        >
          {{ avatar }}
        </el-tag>
        <div>
          <el-image
            lazy
            class="table-td-thumb"
            style="margin: 10px"
            :preview-src-list="[avatar]"
            :src="avatar"
            fit="cover"
          ></el-image>
        </div>
      </div>
      <el-input
        class="input-new-tag"
        v-if="inputRandomAvatarVisible"
        v-model="inputRandomAvatarValue"
        ref="saveRandomAvatarInput"
        size="small"
        @keyup.enter.native="handleInputRandomAvatarConfirm"
        @blur="handleInputRandomAvatarConfirm"
      >
      </el-input>
      <el-button
        v-else
        class="button-new-tag"
        size="small"
        @click="showRandomAvatarInput"
        >+ 随机头像</el-button
      >
      <uploadPicture
        v-hasPermi="['user:visit:read']"
        :ResourceType="'randomAvatar'"
        :isAdmin="true"
        style="margin: 10px"
        @addPicture="addRandomAvatar"
        :maxSize="1"
        :maxNumber="5"
      ></uploadPicture>
      <div class="myCenter" style="margin-bottom: 22px">
        <el-button
          v-hasPermi="['user:visit:read']"
          type="primary"
          @click="saveRandomAvatar()"
          >保存</el-button
        >
      </div>
    </div>
    <!-- 其他图片 -->
    <div>
      <el-tag effect="dark" class="my-tag">
        <img style="vertical-align: -3px" src="../../assets/svg/edit.svg" />
        其他图片
      </el-tag>
      <div :key="i" style="display: flex" v-for="(cover, i) in randomCover">
        <el-tag
          style="white-space: normal; height: unset"
          closable
          :disable-transitions="false"
          @close="handleClose(randomCover, cover)"
        >
          {{ cover }}
        </el-tag>
        <div>
          <el-image
            lazy
            class="table-td-thumb"
            style="margin: 10px"
            :preview-src-list="[cover]"
            :src="cover"
            fit="cover"
          ></el-image>
        </div>
      </div>
      <el-input
        class="input-new-tag"
        v-if="inputRandomCoverVisible"
        v-model="inputRandomCoverValue"
        ref="saveRandomCoverInput"
        size="small"
        @keyup.enter.native="handleInputRandomCoverConfirm"
        @blur="handleInputRandomCoverConfirm"
      >
      </el-input>
      <el-button
        v-else
        class="button-new-tag"
        size="small"
        @click="showRandomCoverInput"
        >+ 其他图片</el-button
      >
      <uploadPicture
        v-hasPermi="['user:visit:read']"
        :ResourceType="'randomCover'"
        :isAdmin="true"
        style="margin: 10px"
        @addPicture="addRandomCover"
        :maxSize="5"
        :maxNumber="5"
      ></uploadPicture>
      <div class="myCenter" style="margin-bottom: 40px">
        <el-button
          v-hasPermi="['user:visit:read']"
          type="primary"
          @click="saveRandomCover()"
          >保存</el-button
        >
      </div>
    </div>
    <!-- 重置所有修改 -->
    <div>
      <el-button
        v-hasPermi="['user:visit:read']"
        type="danger"
        @click="resetForm('ruleForm')"
        >重置所有修改</el-button
      >
    </div>
  </div>
</template>
<script>
const uploadPicture = () => import("../common/uploadPicture");
export default {
  components: {
    uploadPicture,
  },
  data() {
    return {
      //控制提示框是否能被选中
      disabled: true,
      //控制随机名称的按钮背景色
      types: ["", "success", "info", "danger", "warning"],
      //控制+公告按钮的出现或者隐藏
      inputNoticeVisible: false,
      inputNoticeValue: "",
      //控制+随机名称按钮的出现或者隐藏
      inputRandomNameVisible: false,
      inputRandomNameValue: "",
      //控制+随机头像按钮的出现或者隐藏
      inputRandomAvatarVisible: false,
      inputRandomAvatarValue: "",
      //控制+其他图片按钮的出现或者隐藏
      inputRandomCoverVisible: false,
      inputRandomCoverValue: "",
      webInfo: {
        id: null,
        webName: "",
        webTitle: "",
        footer: "",
        backgroundImage: "",
        avatar: "",
        //提示
        waifuJson: "",
        status: false,
      },
      //公告
      notices: [],
      randomAvatar: [],
      randomName: [],
      randomCover: [],
      rules: {
        webName: [
          { required: true, message: "请输入网站名称", trigger: "blur" },
          {
            min: 1,
            max: 12,
            message: "长度在 1 到 12 个字符",
            trigger: "change",
          },
        ],
        webTitle: [
          { required: true, message: "请输入网站标题", trigger: "blur" },
        ],
        footer: [{ required: true, message: "请输入页脚", trigger: "blur" }],
        backgroundImage: [
          { required: true, message: "请输入背景", trigger: "change" },
        ],
        status: [
          { required: true, message: "请设置网站状态", trigger: "change" },
        ],
        avatar: [{ required: true, message: "请上传头像", trigger: "change" }],
      },
    };
  },
  created() {
    this.getWebInfo();
  },
  methods: {
    //接收传过来的url
    addBackgroundImage(res) {
      this.webInfo.backgroundImage = res;
    },
    addAvatar(res) {
      this.webInfo.avatar = res;
    },
    addRandomAvatar(res) {
      this.randomAvatar.push(res);
    },
    addRandomCover(res) {
      this.randomCover.push(res);
    },
    changeWebStatus(webInfo) {
      this.$http
        .post(
          this.$constant.baseURL + "/admin/webInfo/updateAdminWebInfo/",
          {
            id: webInfo.id,
            status: webInfo.status,
          },
          true,
          false,
          true
        )
        .then((res) => {
          this.$notify({
            title: "可以啦🍨",
            message: "保存成功！",
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
    },
    getWebInfo() {
      this.$http
        .get(
          this.$constant.baseURL + "/admin/webInfo/getAdminWebInfo/",
          {},
          true,
          true
        )
        .then((res) => {
          if (!this.$common.isEmpty(res.result[0].data[0])) {
            this.webInfo.id = res.result[0].data[0].id;
            this.webInfo.webName = res.result[0].data[0].webName;
            this.webInfo.webTitle = res.result[0].data[0].webTitle;
            this.webInfo.footer = res.result[0].data[0].footer;
            this.webInfo.backgroundImage =
              res.result[0].data[0].backgroundImage;
            this.webInfo.avatar = res.result[0].data[0].avatar;
            this.webInfo.waifuJson = res.result[0].data[0].waifuJson;
            this.webInfo.status = res.result[0].data[0].status;
            this.notices = JSON.parse(res.result[0].data[0].notices);
            if (!res.result[0].data[0].randomAvatar) {
              this.randomAvatar = "";
            } else {
              this.randomAvatar = JSON.parse(
                res.result[0].data[0].randomAvatar
              );
            }
            if (!res.result[0].data[0].randomName) {
              this.randomName = "";
            } else {
              this.randomName = JSON.parse(res.result[0].data[0].randomName);
            }
            if (!res.result[0].data[0].randomCover) {
              this.randomCover = "";
            } else {
              this.randomCover = JSON.parse(res.result[0].data[0].randomCover);
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
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.updateWebInfo(this.webInfo);
        } else {
          this.$notify({
            type: "error",
            title: "可恶🤬",
            message: "请完善必填项！",
            position: "top-left",
            offset: 50,
          });
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
      this.getWebInfo();
    },
    handleClose(array, item) {
      array.splice(array.indexOf(item), 1);
    },
    handleInputNoticeConfirm() {
      if (this.inputNoticeValue) {
        this.notices.push(this.inputNoticeValue);
      }
      this.inputNoticeVisible = false;
      this.inputNoticeValue = "";
    },
    showNoticeInput() {
      this.inputNoticeVisible = true;
      this.$nextTick(() => {
        this.$refs.saveNoticeInput.$refs.input.focus();
      });
    },
    saveNotice() {
      let param = {
        id: this.webInfo.id,
        notices: JSON.stringify(this.notices),
      };
      this.updateWebInfo(param);
    },
    handleInputRandomNameConfirm() {
      if (this.inputRandomNameValue) {
        this.randomName.push(this.inputRandomNameValue);
      }
      this.inputRandomNameVisible = false;
      this.inputRandomNameValue = "";
    },
    showRandomNameInput() {
      this.inputRandomNameVisible = true;
      this.$nextTick(() => {
        this.$refs.saveRandomNameInput.$refs.input.focus();
      });
    },
    saveRandomName() {
      let param = {
        id: this.webInfo.id,
        randomName: JSON.stringify(this.randomName),
      };
      this.updateWebInfo(param);
    },
    handleInputRandomAvatarConfirm() {
      if (this.inputRandomAvatarValue) {
        this.randomAvatar.push(this.inputRandomAvatarValue);
      }
      this.inputRandomAvatarVisible = false;
      this.inputRandomAvatarValue = "";
    },
    showRandomAvatarInput() {
      this.inputRandomAvatarVisible = true;
      this.$nextTick(() => {
        this.$refs.saveRandomAvatarInput.$refs.input.focus();
      });
    },
    saveRandomAvatar() {
      let param = {
        id: this.webInfo.id,
        randomAvatar: JSON.stringify(this.randomAvatar),
      };
      this.updateWebInfo(param);
    },
    handleInputRandomCoverConfirm() {
      if (this.inputRandomCoverValue) {
        this.randomCover.push(this.inputRandomCoverValue);
      }
      this.inputRandomCoverVisible = false;
      this.inputRandomCoverValue = "";
    },
    showRandomCoverInput() {
      this.inputRandomCoverVisible = true;
      this.$nextTick(() => {
        this.$refs.saveRandomCoverInput.$refs.input.focus();
      });
    },
    saveRandomCover() {
      let param = {
        id: this.webInfo.id,
        randomCover: JSON.stringify(this.randomCover),
      };
      this.updateWebInfo(param);
    },
    updateWebInfo(value) {
      this.$confirm("确认保存？", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "success",
        center: true,
      })
        .then(() => {
          this.$http
            .post(
              this.$constant.baseURL + "/admin/webInfo/updateAdminWebInfo/",
              value,
              true,
              false,
              true
            )
            .then((res) => {
              this.$notify({
                title: "可以啦🍨",
                message: "保存成功！",
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
        })
        .catch(() => {
          this.$notify({
            title: "可以啦🍨",
            message: "已取消保存！",
            type: "success",
            offset: 50,
            position: "top-left",
          });
        });
    },
  },
};
</script>
<style lang="scss" scoped>
.my-tag {
  margin-bottom: 20px !important;
  width: 100%;
  text-align: left;
  background: var(--green2);
  border: none;
  height: 40px;
  line-height: 40px;
  font-size: 16px;
  color: var(--favoriteBg);
}
.el-tag {
  margin: 10px;
}
.button-new-tag {
  margin: 10px;
  height: 32px;
  line-height: 32px;
  padding-top: 0;
  padding-bottom: 0;
}
.input-new-tag {
  width: 200px;
  margin: 10px;
}
.my-icon {
  margin-left: 10px;
  font-size: 18px;
  font-weight: bold;
  color: var(--maxLightRed);
}
.table-td-thumb {
  border-radius: 2px;
  width: 40px;
  height: 40px;
}
</style>
