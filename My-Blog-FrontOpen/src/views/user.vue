<template>
  <div>
    <!-- 登录和注册 -->
    <div
      v-if="$common.isEmpty(currentUser)"
      class="myCenter in-up-container my-animation-hideToShow"
    >
      <!-- 背景图片 -->
      <div
        style="animation: header-effect 2s"
        :style="{ background: `${$store.state.changeBg}` }"
        class="background-image background-image-changeBg"
      ></div>
      <div
        class="in-up"
        id="loginAndRegister"
        :class="{ 'right-panel-active': active }"
      >
        <div class="form-container sign-up-container">
          <div class="myCenter">
            <h1>注册</h1>
            <input
              v-model="username"
              type="text"
              maxlength="30"
              placeholder="用户名"
            />
            <input
              v-model="password"
              type="password"
              maxlength="30"
              placeholder="密码"
            />
            <input v-model="email" type="email" placeholder="邮箱" />
            <input v-model="code" type="text" placeholder="验证码" disabled />
            <a style="margin: 0" href="#" @click="changeDialog('邮箱验证码')"
              >获取验证码</a
            >
            <button @click="register()">注册</button>
          </div>
        </div>
        <div class="form-container sign-in-container">
          <div class="myCenter">
            <h1>登录</h1>
            <input v-model="account" type="text" placeholder="用户名/邮箱" />
            <input
              v-model="password"
              @keyup.enter="login()"
              type="password"
              placeholder="密码"
            />
            <a href="javascript:;" @click="changeDialog('找回密码')"
              >忘记密码？</a
            >
            <button @click="login()">登录</button>
          </div>
        </div>
        <div class="overlay-container">
          <div class="overlay">
            <div class="overlay-panel myCenter overlay-left">
              <h1>已有帐号？</h1>
              <p>请登录🚀</p>
              <button class="ghost" @click="signIn()">登录</button>
            </div>
            <div class="overlay-panel myCenter overlay-right">
              <h1>没有帐号？</h1>
              <p>立即注册吧😃</p>
              <button class="ghost" @click="signUp()">注册</button>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- 已登录 用户信息 -->
    <div v-else class="user-container myCenter my-animation-hideToShow">
      <!-- 背景图片 -->
      <div
        style="animation: header-effect 2s"
        :style="{ background: `${$store.state.changeBg}` }"
        class="background-image background-image-changeBg"
      ></div>
      <div class="shadow-box-mini user-info" style="display: flex">
        <div class="user-left">
          <div style="display: flex; align-items: center">
            <p style="margin-right: 10px; color: var(--red)">撮我替换头像 👉</p>
            <el-avatar
              class="user-avatar"
              @click.native="changeDialog('修改头像')"
              :size="60"
              :src="currentUser.avatar || $store.state.webInfo.avatar"
            ></el-avatar>
          </div>
          <div class="myCenter" style="margin-top: 12px">
            <div class="user-title">
              <div>用户名：</div>
              <div>手机号：</div>
              <div>邮箱：</div>
              <div>性别：</div>
              <div>简介：</div>
              <div v-if="$store.state.currentUser.userType !== 2">
                七牛云域名：
              </div>
              <div v-if="$store.state.currentUser.userType !== 2">
                七牛云访问密钥：
              </div>
              <div v-if="$store.state.currentUser.userType !== 2">
                七牛云秘密秘钥：
              </div>
              <div v-if="$store.state.currentUser.userType !== 2">
                七牛云文件夹：
              </div>
            </div>
            <div class="user-content">
              <div>
                <el-input
                  maxlength="30"
                  v-model="currentUser.username"
                ></el-input>
              </div>
              <div>
                <div style="width: 100%">
                  <el-input
                    maxlength="11"
                    v-model="currentUser.phoneNumber"
                  ></el-input>
                </div>
                <div v-if="$common.isEmpty(currentUser.phoneNumber)">
                  <span class="changeInfo">输入手机号</span>
                </div>
              </div>
              <div>
                <div v-if="!$common.isEmpty(currentUser.email)">
                  <span style="color: var(--black5)">{{
                    currentUser.email
                  }}</span>
                  <span class="changeInfo" @click="changeDialog('修改邮箱')"
                    >修改</span
                  >
                </div>
              </div>
              <div>
                <el-radio-group
                  class="mk-radio__group"
                  v-model="currentUser.gender"
                >
                  <el-radio :label="0" style="margin-right: 10px"
                    >薛定谔的猫</el-radio
                  >
                  <el-radio :label="1" style="margin-right: 10px">男</el-radio>
                  <el-radio :label="2">女</el-radio>
                </el-radio-group>
              </div>
              <div>
                <el-input
                  v-model="currentUser.introduction"
                  maxlength="60"
                  type="textarea"
                  show-word-limit
                ></el-input>
              </div>
              <div v-if="$store.state.currentUser.userType !== 2">
                <el-input
                  v-model="currentUser.qiniuDomain"
                  maxlength="128"
                ></el-input>
              </div>
              <div v-if="$store.state.currentUser.userType !== 2">
                <el-input
                  v-model="currentUser.qiniuAccessKey"
                  maxlength="128"
                ></el-input>
              </div>
              <div v-if="$store.state.currentUser.userType !== 2">
                <el-input
                  v-model="currentUser.qiniuSecretKey"
                  maxlength="128"
                ></el-input>
              </div>
              <div v-if="$store.state.currentUser.userType !== 2">
                <el-input
                  v-model="currentUser.qiniuBucketName"
                  maxlength="128"
                ></el-input>
              </div>
            </div>
          </div>
          <div style="margin-top: 20px">
            <proButton
              :info="'提交'"
              @click.native="submitUserInfo()"
              :before="$constant.before_color_1"
              :after="$constant.after_color_1"
            >
            </proButton>
          </div>
        </div>
        <div class="user-right"></div>
      </div>
    </div>

    <el-dialog
      :title="dialogTitle"
      :visible.sync="showDialog"
      width="30%"
      :before-close="clearDialog"
      :append-to-body="true"
      :close-on-click-modal="false"
      center
      custom-class="dialog"
    >
      <div class="myCenter" style="flex-direction: column">
        <div>
          <div v-if="dialogTitle === '修改邮箱'">
            <div style="margin-bottom: 5px">邮箱：</div>
            <el-input v-model="email"></el-input>
            <div style="margin-top: 10px; margin-bottom: 5px">验证码：</div>
            <el-input v-model="code"></el-input>
          </div>
          <div v-else-if="dialogTitle === '修改头像'">
            <uploadPicture
              :ResourceType="'userAvatar'"
              @addPicture="addPicture"
              :maxSize="1"
              :maxNumber="1"
            >
            </uploadPicture>
          </div>
          <div v-else-if="dialogTitle === '找回密码'">
            <div>
              <div style="margin-bottom: 5px">用户名：</div>
              <el-input v-model="username"></el-input>
              <div style="margin-bottom: 5px">邮箱：</div>
              <el-input v-model="email"></el-input>
              <div style="margin-top: 10px; margin-bottom: 5px">验证码：</div>
              <el-input v-model="code"></el-input>
              <div style="margin-top: 10px; margin-bottom: 5px">新密码：</div>
              <el-input maxlength="30" v-model="password"></el-input>
            </div>
          </div>
          <div v-else-if="dialogTitle === '邮箱验证码'">
            <div>
              <div style="margin-bottom: 5px">邮箱：</div>
              <el-input v-model="email"></el-input>
              <div style="margin-top: 10px; margin-bottom: 5px">验证码：</div>
              <el-input v-model="code"></el-input>
              <span
                style="
                  font-size: 14px;
                  color: var(--bigRed);
                  margin-top: 6px;
                  display: inline-block;
                "
                >验证码不区分大小写</span
              >
            </div>
          </div>
        </div>
        <div
          style="display: flex; margin-top: 30px"
          v-show="dialogTitle !== '修改头像'"
        >
          <proButton
            :info="codeString"
            v-show="
              dialogTitle === '修改邮箱' ||
              dialogTitle === '找回密码' ||
              dialogTitle === '邮箱验证码'
            "
            @click.native="getCode()"
            :before="$constant.before_color_1"
            :after="$constant.after_color_1"
            style="margin-right: 20px"
          >
          </proButton>
          <proButton
            :info="'提交'"
            @click.native="submitDialog()"
            :before="$constant.before_color_1"
            :after="$constant.after_color_1"
          >
          </proButton>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
const proButton = () => import("./common/proButton");
const uploadPicture = () => import("./common/uploadPicture");

export default {
  components: {
    proButton,
    uploadPicture,
  },
  data() {
    return {
      username: "",
      account: "",
      password: "",
      email: "",
      avatar: "",
      showDialog: false,
      code: "",
      dialogTitle: "",
      codeString: "验证码",
      //定时器
      intervalCode: null,
      currentUser: {},
      active: false,
      province: "",
    };
  },
  mounted() {
    this.postProvinceAndCity();
    this.currentUser = this.$store.state.currentUser;
  },
  methods: {
    addPicture(res) {
      this.avatar = res;
      this.submitDialog();
    },
    signUp() {
      this.active = true;
    },
    signIn() {
      this.active = false;
    },
    async postProvinceAndCity() {
      this.province = "";
      const res = await this.$common.getIpAndCity(this);
      this.province = res.address;
    },
    login() {
      if (
        this.$common.isEmpty(this.account) ||
        this.$common.isEmpty(this.password)
      ) {
        this.$notify({
          type: "error",
          title: "可恶🤬",
          message: "请输入账号或密码！",
          position: "top-left",
          offset: 50,
        });
        return;
      }
      let user = {
        province: this.province || "",
        account: this.account.trim(),
        password: this.$common.encrypt(this.password.trim()),
      };
      this.$http
        .post(this.$constant.baseURL + "/appone/login/", user, false, false)
        .then((res) => {
          if (res.result === "The account is disabled or deleted") {
            this.$notify({
              type: "error",
              title: "可恶🤬",
              message: "您的账号已被管理员禁用！",
              position: "top-left",
              offset: 50,
            });
            return;
          }
          if (!this.$common.isEmpty(res.result[0])) {
            this.$notify({
              title: "可以啦🍨",
              message: "登录成功！欢迎光临小舍~~~🥰🥰🥰",
              type: "success",
              offset: 50,
              position: "top-left",
            });
            this.$store.commit("loadCurrentUser", res.result[0].data[0]);
            localStorage.setItem(
              "userToken",
              res.result[0].data[0].accessToken
            );
            this.account = "";
            this.password = "";
            this.$router.push({ path: "/" });
          }
        })
        .catch((error) => {
          if (
            error.response.data.error === "Username or password is incorrect."
          ) {
            this.$notify({
              type: "error",
              title: "可恶🤬",
              message: "密码错误！请输入正确密码~~",
              position: "top-left",
              offset: 50,
            });
          } else {
            this.$notify({
              type: "error",
              title: "可恶🤬",
              message: error.message,
              position: "top-left",
              offset: 50,
            });
          }
        });
    },
    register() {
      if (
        this.$common.isEmpty(this.username) ||
        this.$common.isEmpty(this.password)
      ) {
        this.$notify({
          type: "error",
          title: "可恶🤬",
          message: "请输入用户名或密码！",
          position: "top-left",
          offset: 50,
        });
        return;
      }
      if (
        this.dialogTitle === "邮箱验证码" &&
        this.$common.isEmpty(this.email)
      ) {
        this.$notify({
          type: "error",
          title: "可恶🤬",
          message: "请输入邮箱！",
          position: "top-left",
          offset: 50,
        });
        return false;
      }
      if (this.$common.isEmpty(this.code)) {
        this.$notify({
          type: "error",
          title: "可恶🤬",
          message: "请输入验证码！",
          position: "top-left",
          offset: 50,
        });
        return;
      }
      if (
        this.username.indexOf(" ") !== -1 ||
        this.password.indexOf(" ") !== -1
      ) {
        this.$notify({
          type: "error",
          title: "可恶🤬",
          message: "用户名或密码不能包含空格！",
          position: "top-left",
          offset: 50,
        });
        return;
      }
      if (this.username === "monkey-papa") {
        this.$notify({
          type: "error",
          title: "可恶🤬",
          message: "用户名不可以跟店长昵称一样喔~~~😊",
          position: "top-left",
          offset: 50,
        });
        return;
      }
      let user = {
        province: this.province || "",
        username: this.username.trim(),
        code: this.code.trim().toUpperCase(),
        password: this.$common.encrypt(this.password.trim()),
      };
      if (this.dialogTitle === "邮箱验证码") {
        user.email = this.email;
      }
      this.$http
        .post(this.$constant.baseURL + "/appone/registration/", user)
        .then((res) => {
          if (res.result === "验证码错误") {
            this.$notify({
              type: "error",
              title: "可恶🤬",
              message: "验证码错误！",
              position: "top-left",
              offset: 50,
            });
            return;
          }
          if (res.result === "验证码已过期") {
            this.$notify({
              type: "error",
              title: "可恶🤬",
              message: "验证码已过期！",
              position: "top-left",
              offset: 50,
            });
            return;
          }
          if (!this.$common.isEmpty(res.result[0].data)) {
            this.$notify({
              title: "可以啦🍨",
              message: "注册成功！",
              type: "success",
              offset: 50,
              position: "top-left",
            });
            this.$store.commit("loadCurrentUser", res.result[0].data[0]);
            localStorage.setItem(
              "userToken",
              res.result[0].data[0].accessToken
            );
            this.username = "";
            this.password = "";
            this.$router.push({ path: "/" });
          }
        })
        .catch((error) => {
          if (error.response.data.error === "Username already exists.") {
            this.$notify({
              type: "error",
              title: "可恶🤬",
              message: "账号已被注册，请重新注册~~",
              position: "top-left",
              offset: 50,
            });
          } else {
            this.$notify({
              type: "error",
              title: "可恶🤬",
              message: "用户名或者邮箱已存在！",
              position: "top-left",
              offset: 50,
            });
          }
        });
    },
    // 资料信息修改
    submitUserInfo() {
      if (!this.checkParameters()) {
        return;
      }
      let user = {
        gender: this.currentUser.gender,
        phoneNumber: this.currentUser.phoneNumber,
        email: this.currentUser.email,
        introduction: this.currentUser.introduction,
        userId: this.$store.state.currentUser.id,
      };
      if (!this.$common.isEmpty(this.currentUser.introduction)) {
        user.introduction = this.currentUser.introduction.trim();
      }
      const allQiniuFieldsEmpty =
        this.$common.isEmpty(this.currentUser.qiniuDomain) &&
        this.$common.isEmpty(this.currentUser.qiniuAccessKey) &&
        this.$common.isEmpty(this.currentUser.qiniuSecretKey) &&
        this.$common.isEmpty(this.currentUser.qiniuBucketName);
      const allQiniuFieldsFilled =
        !this.$common.isEmpty(this.currentUser.qiniuDomain) &&
        !this.$common.isEmpty(this.currentUser.qiniuAccessKey) &&
        !this.$common.isEmpty(this.currentUser.qiniuSecretKey) &&
        !this.$common.isEmpty(this.currentUser.qiniuBucketName);
      if (allQiniuFieldsEmpty || allQiniuFieldsFilled) {
        if (allQiniuFieldsFilled) {
          user.qiniuDomain = this.currentUser.qiniuDomain.trim();
          user.qiniuAccessKey = this.currentUser.qiniuAccessKey.trim();
          user.qiniuSecretKey = this.currentUser.qiniuSecretKey.trim();
          user.qiniuBucketName = this.currentUser.qiniuBucketName.trim();
        }
      } else {
        this.$notify({
          title: "淘气👻",
          message: "请将四项七牛云信息填写完整，或者四项全部留空！",
          type: "warning",
          offset: 50,
          position: "top-left",
        });
        return;
      }
      this.$confirm("确认保存？", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
        center: true,
      })
        .then(() => {
          this.$http
            .post(
              this.$constant.baseURL + "/user/updateUserInfo/",
              user,
              false,
              true,
              true
            )
            .then((res) => {
              if (!this.$common.isEmpty(res.result[0])) {
                this.$store.commit("loadCurrentUser", res.result[0].data[0]);
                this.currentUser = this.$store.state.currentUser;
                this.$notify({
                  title: "可以啦🍨",
                  message: "修改成功！",
                  type: "success",
                  offset: 50,
                  position: "top-left",
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
    checkParams(params) {
      if (
        this.dialogTitle === "修改邮箱" ||
        this.dialogTitle === "邮箱验证码" ||
        this.dialogTitle === "找回密码"
      ) {
        if (this.$common.isEmpty(this.email)) {
          this.$notify({
            type: "error",
            title: "可恶🤬",
            message: "请输入邮箱！",
            position: "top-left",
            offset: 50,
          });
          return false;
        }
        if (!/^\w+@[a-zA-Z0-9]{2,10}(?:\.[a-z]{2,4}){1,3}$/.test(this.email)) {
          this.$notify({
            type: "error",
            title: "可恶🤬",
            message: "邮箱格式有误！",
            position: "top-left",
            offset: 50,
          });
          return false;
        }
        params.email = this.email;
        return true;
      }
      return false;
    },
    checkParameters() {
      if (this.$common.isEmpty(this.currentUser.username)) {
        this.$notify({
          type: "error",
          title: "可恶🤬",
          message: "请输入用户名！",
          position: "top-left",
          offset: 50,
        });
        return false;
      }
      if (
        !/^1[345789]\d{9}$/.test(this.currentUser.phoneNumber) &&
        this.currentUser.phoneNumber
      ) {
        this.$notify({
          type: "error",
          title: "可恶🤬",
          message: "手机号格式有误！",
          position: "top-left",
          offset: 50,
        });
        return false;
      }
      if (this.currentUser.username.indexOf(" ") !== -1) {
        this.$notify({
          type: "error",
          title: "可恶🤬",
          message: "用户名不能包含空格！",
          position: "top-left",
          offset: 50,
        });
        return false;
      }
      return true;
    },
    changeDialog(value) {
      if (value === "邮箱验证码") {
        if (this.$common.isEmpty(this.email)) {
          this.$notify({
            type: "error",
            title: "可恶🤬",
            message: "请输入邮箱！",
            position: "top-left",
            offset: 50,
          });
          return false;
        }
        if (!/^\w+@[a-zA-Z0-9]{2,10}(?:\.[a-z]{2,4}){1,3}$/.test(this.email)) {
          this.$notify({
            type: "error",
            title: "可恶🤬",
            message: "邮箱格式有误！",
            position: "top-left",
            offset: 50,
          });
          return false;
        }
        // 一进入弹窗就获取验证码，前提得赋值弹窗标题
        this.dialogTitle = value;
        this.getCode();
      }
      this.dialogTitle = value;
      this.showDialog = true;
    },
    // 头像修改
    submitDialog() {
      if (this.dialogTitle === "修改头像") {
        if (this.$common.isEmpty(this.avatar)) {
          this.$notify({
            type: "error",
            title: "可恶🤬",
            message: "请上传头像！",
            position: "top-left",
            offset: 50,
          });
        } else {
          let user = {
            avatar: this.avatar.trim(),
            userId: this.$store.state.currentUser.id,
          };
          this.$http
            .post(
              this.$constant.baseURL + "/user/updateUserInfo/",
              user,
              false,
              true,
              true
            )
            .then((res) => {
              if (!this.$common.isEmpty(res.result[0])) {
                this.$store.commit("loadCurrentUser", res.result[0].data[0]);
                this.currentUser = this.$store.state.currentUser;
                this.clearDialog();
                this.$notify({
                  title: "可以啦🍨",
                  message: "修改成功！",
                  type: "success",
                  offset: 50,
                  position: "top-left",
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
        }
      } else if (this.dialogTitle === "修改邮箱") {
        this.updateSecretInfo();
      } else if (this.dialogTitle === "找回密码") {
        this.updateSecretInfo();
      } else if (this.dialogTitle === "邮箱验证码") {
        this.showDialog = false;
      }
    },
    updateSecretInfo() {
      if (this.$common.isEmpty(this.code)) {
        this.$notify({
          type: "error",
          title: "可恶🤬",
          message: "请输入验证码！",
          position: "top-left",
          offset: 50,
        });
        return;
      }
      if (
        this.$common.isEmpty(this.password) &&
        this.dialogTitle !== "修改邮箱"
      ) {
        this.$notify({
          type: "error",
          title: "可恶🤬",
          message: "请输入密码！",
          position: "top-left",
          offset: 50,
        });
        return;
      }
      let params = {
        email: this.email,
        code: this.code.trim(),
        password: this.$common.encrypt(this.password.trim()),
        username: this.username,
        userId: this.$store.state.currentUser.id,
      };
      if (!this.checkParams(params)) {
        return;
      }
      if (this.dialogTitle === "找回密码") {
        this.$http
          .post(
            this.$constant.baseURL + "/user/updateForForgetPassword/",
            params,
            false,
            false
          )
          .then((res) => {
            this.clearDialog();
            if (res.result === "验证码错误") {
              this.$notify({
                type: "error",
                title: "可恶🤬",
                message: "验证码错误！",
                position: "top-left",
                offset: 50,
              });
              return;
            }
            if (res.result === "验证码已过期") {
              this.$notify({
                type: "error",
                title: "可恶🤬",
                message: "验证码已过期！",
                position: "top-left",
                offset: 50,
              });
              return;
            }
            this.$notify({
              title: "可以啦🍨",
              message: "修改成功，请重新登录！",
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
        this.$http
          .post(
            this.$constant.baseURL + "/user/updateUserInfo/",
            params,
            false,
            false,
            true
          )
          .then((res) => {
            if (!this.$common.isEmpty(res.result[0])) {
              this.$store.commit("loadCurrentUser", res.result[0].data[0]);
              this.currentUser = this.$store.state.currentUser;
              this.clearDialog();
              this.$notify({
                title: "可以啦🍨",
                message: "修改成功！",
                type: "success",
                offset: 50,
                position: "top-left",
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
      }
    },
    getCode() {
      if (this.codeString === "验证码") {
        // 获取验证码
        let params = {};
        if (!this.checkParams(params)) {
          return;
        }
        this.$http
          .post(this.$constant.baseURL + "/code/", params)
          .then((res) => {
            this.$notify({
              title: "可以啦🍨",
              message: "验证码已发送，请注意查收！",
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
        this.codeString = "30";
        this.intervalCode = setInterval(() => {
          if (this.codeString === "0") {
            clearInterval(this.intervalCode);
            this.codeString = "验证码";
          } else {
            this.codeString = parseInt(this.codeString) - 1 + "";
          }
        }, 1000);
      } else {
        this.$notify({
          type: "warning",
          title: "淘气👻",
          message: "请稍后再试！",
          position: "top-left",
          offset: 50,
        });
      }
    },
    clearDialog() {
      this.password = "";
      this.email = "";
      this.avatar = "";
      this.showDialog = false;
      this.code = "";
      this.dialogTitle = "";
    },
  },
};
</script>

<style lang="scss" scoped>
.in-up-container {
  height: 100vh;
  position: relative;
}
.in-up {
  opacity: 0.9;
  border-radius: 10px;
  box-shadow: 0 15px 30px var(--miniMask), 0 10px 10px var(--miniMask);
  position: relative;
  overflow: hidden;
  width: 750px;
  max-width: 100%;
  min-height: 450px;
  margin: 10px;
  p {
    font-size: 14px;
    letter-spacing: 1px;
    margin: 20px 0 30px 0;
  }
  a {
    color: var(--black);
    font-size: 14px;
    text-decoration: none;
    margin: 15px 0;
  }
  button {
    border-radius: 2rem;
    border: none;
    background: var(--lightRed);
    color: var(--white);
    font-size: 16px;
    font-weight: bold;
    padding: 12px 45px;
    letter-spacing: 2px;
    &:hover {
      animation: scale 0.8s ease-in-out;
    }
    &.ghost {
      background: transparent;
      border: 1px solid var(--white);
    }
  }
  &.right-panel-active .sign-in-container {
    transform: translateY(100%);
  }
  &.right-panel-active .overlay-container {
    transform: translateX(-100%);
  }
  &.right-panel-active .sign-up-container {
    transform: translateX(100%);
    opacity: 1;
  }
  &.right-panel-active .overlay {
    transform: translateX(50%);
    &-left {
      transform: translateY(0);
    }
    &-right {
      transform: translateY(20%);
    }
  }
}
.form-container {
  position: absolute;
  height: 100%;
  transition: all 0.5s ease-in-out;
  div {
    background: var(--white);
    flex-direction: column;
    padding: 0 20px;
    height: 100%;
  }
  input {
    background: var(--white);
    border-radius: 2px;
    border: none;
    padding: 12px 15px;
    margin: 10px 0;
    width: 100%;
    outline: none;
  }
}
.sign-in-container {
  left: 0;
  width: 50%;
}
.sign-up-container {
  left: 0;
  width: 50%;
  opacity: 0;
  button {
    margin-top: 20px;
  }
}
.overlay-container {
  position: absolute;
  left: 50%;
  width: 50%;
  height: 100%;
  overflow: hidden;
  transition: all 0.5s ease-in-out;
}
.overlay {
  background: var(--gradualRed);
  color: var(--white);
  position: relative;
  left: -100%;
  height: 100%;
  width: 200%;
  &-panel {
    position: absolute;
    top: 0;
    flex-direction: column;
    height: 100%;
    width: 50%;
    transition: all 0.5s ease-in-out;
  }
  &-right {
    right: 0;
    transform: translateY(0);
  }
  &-left {
    transform: translateY(-20%);
  }
}
.user-container {
  width: 100vw;
  height: 100vh;
  position: relative;
}
.user-info {
  width: 80%;
  z-index: 10;
  margin-top: 70px;
  height: calc(100vh - 90px);
  margin-bottom: 20px;
  border-radius: 10px;
  overflow: hidden;
}
.user-left {
  width: 50%;
  background: var(--maxMaxWhiteMask);
  display: flex;
  flex-direction: column;
  align-items: center;
  overflow-y: auto;
  padding: 20px;
  .mk-radio__group {
    ::v-deep {
      .el-radio__label {
        color: var(--red);
      }
      .el-radio__input.is-checked .el-radio__inner {
        border-color: var(--red);
        background: var(--red);
      }
    }
  }
}
.user-right {
  width: 50%;
  background: var(--whiteMask);
  padding: 20px;
}
.user-title {
  text-align: right;
  user-select: none;
  div {
    text-align: right;
    color: var(--black5);
    height: 55px;
    line-height: 55px;
  }
}
.user-content {
  text-align: left;
  > div {
    height: 55px;
    display: flex;
    align-items: center;
  }
  ::v-deep .el-input__inner,
  ::v-deep .el-textarea__inner {
    border: none;
    background: var(--whiteMask);
  }
  ::v-deep .el-input__count {
    background: var(--transparent);
    user-select: none;
  }
}
.changeInfo {
  color: var(--white);
  font-size: 0.75rem;
  white-space: nowrap;
  background: var(--gradientAnimation);
  padding: 3px;
  border-radius: 0.2rem;
  user-select: none;
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

@media screen and (max-width: 940px) {
  .user-info {
    width: 90%;
  }
  .user-left {
    width: 100%;
  }
  .user-right {
    display: none;
  }
}
</style>
