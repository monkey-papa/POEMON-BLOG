<template>
  <div>
    <!-- 背景图片 -->
    <div
      style="animation: header-effect 2s"
      :style="{ background: `${$store.state.changeBg}` }"
      class="background-image background-image-changeBg"
    ></div>
    <!-- 封面 -->
    <div class="friend-head myCenter">
      <h1
        style="
          color: var(--blue2);
          z-index: 10;
          font-weight: 400;
          font-size: 45px;
        "
      >
        通讯录
      </h1>
    </div>
    <div class="friend-wrap">
      <div class="friend-main">
        <!-- 添加友链 -->
        <div @click="clickLetter()" class="form-wrap">
          <!-- 信封上面 -->
          <img
            class="before-img"
            :src="$store.state.webInfo.randomCover[8]"
            style="width: 100%"
          />
          <!-- 信 -->
          <div class="envelope" style="animation: hideToShow 2s">
            <div class="form-main shadow-box">
              <img
                :src="$store.state.webInfo.randomCover[11]"
                style="width: 100%"
              />
              <div>
                <h3 style="text-align: center">有朋自远方来</h3>
                <div>
                  <div class="myCenter form-friend">
                    <div style="color: var(--white)" class="user-title">
                      <div>名称：</div>
                      <div>简介：</div>
                      <div>封面：</div>
                      <div>头像：</div>
                      <div>网址：</div>
                    </div>
                    <div class="user-content">
                      <div>
                        <el-input
                          maxlength="30"
                          v-model="friend.title"
                        ></el-input>
                      </div>
                      <div>
                        <el-input
                          maxlength="120"
                          v-model="friend.introduction"
                        ></el-input>
                      </div>
                      <div>
                        <el-input
                          maxlength="200"
                          v-model="friend.cover"
                        ></el-input>
                      </div>
                      <div>
                        <el-input
                          maxlength="200"
                          v-model="friend.friendAvatar"
                        ></el-input>
                      </div>
                      <div>
                        <el-input
                          maxlength="200"
                          v-model="friend.url"
                        ></el-input>
                      </div>
                      <span style="color: var(--bigRed1); font-size: 14px"
                        >请填写以http或https开头的有效地址</span
                      >
                    </div>
                  </div>
                  <div class="myCenter" style="margin-top: 5px">
                    <proButton
                      :info="'提交'"
                      @click.native.stop="submitFriend()"
                      :before="$constant.before_color_1"
                      :after="$constant.after_color_1"
                    >
                    </proButton>
                  </div>
                </div>
                <div>
                  <img
                    :src="$store.state.webInfo.randomCover[10]"
                    style="width: 100%; margin: 5px auto"
                  />
                </div>
                <p
                  style="
                    font-size: 12px;
                    text-align: center;
                    color: var(--bigRed1);
                  "
                >
                  欢迎交换友链
                </p>
              </div>
            </div>
          </div>
          <img
            class="after-img"
            :src="$store.state.webInfo.randomCover[9]"
            style="width: 100%"
          />
        </div>
        <hr />
        <h2 style="color: var(--red)">
          <i
            class="fa fa-address-card"
            style="color: var(--green7); font-size: 24px; margin-right: 5px"
          ></i
          >特别鸣谢
        </h2>
        <card
          class="recommendFriend"
          :resourcePathList="thanksFriendList"
          @clickResourcePath="clickFriend"
        ></card>
        <h2 style="color: var(--darkBlue)">
          <i
            class="fa fa-chain"
            style="color: var(--green7); font-size: 30px; margin-right: 5px"
          ></i
          >友情链接
        </h2>
        <card
          :resourcePathList="friendList"
          @clickResourcePath="clickFriend"
        ></card>
      </div>
    </div>
  </div>
</template>
<script>
const card = () => import("./common/card");
const proButton = () => import("./common/proButton");
export default {
  components: {
    card,
    proButton,
  },
  data() {
    return {
      pagination: {
        current: 1,
        size: 9999,
        desc: false,
        resourceType: "friendUrl",
      },
      friendList: [],
      thanksFriendList: [
        {
          createTime: "2023-08-13T16:14:40.021946",
          introduction: " 这是一个 Vue2 Vue3 与 SpringBoot 结合的产物～ ",
          cover:
            "https://www.qiniuyun.monke-papa.icu/images/1d3561c80ada1e007ff1325c81ba96e.png",
          friendAvatar: "https://s1.ax1x.com/2022/11/10/z9E7X4.jpg",
          title: "生活倒影",
          url: "https://poetize.cn/",
        },
        {
          createTime: "2023-08-13T16:16:50.021946",
          introduction: " 宁静致远,倾尘轻笑 ",
          cover:
            "https://www.qiniuyun.monke-papa.icu/images/9869f45ff9ae65e6371c2d89bd09ce3.png",
          friendAvatar: "https://cdn.chuckle.top/img/head.webp",
          title: " 轻笑Chuckle ",
          url: "https://www.qcqx.cn/",
        },
      ],
      friend: {
        title: "",
        introduction: "",
        cover: "",
        friendAvatar: "",
        url: "",
        type: "friendUrl",
      },
    };
  },
  created() {
    this.getFriends();
  },
  methods: {
    clickLetter() {
      if (document.body.clientWidth < 700) {
        $(".form-wrap").css({ height: "1000px", top: "-200px" });
      } else {
        $(".form-wrap").css({ height: "1150px", top: "-200px" });
      }
    },
    submitFriend() {
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
      if (this.friend.title.trim() === "") {
        this.$notify({
          type: "warning",
          title: "淘气👻",
          message: "你还没写名称呢~",
          position: "top-left",
          offset: 50,
        });
        return;
      }
      if (this.friend.introduction.trim() === "") {
        this.$notify({
          type: "warning",
          title: "淘气👻",
          message: "你还没写简介呢~",
          position: "top-left",
          offset: 50,
        });
        return;
      }
      if (this.friend.cover.trim() === "") {
        this.$notify({
          type: "warning",
          title: "淘气👻",
          message: "你还没设置封面呢~",
          position: "top-left",
          offset: 50,
        });
        return;
      }
      if (this.friend.friendAvatar.trim() === "") {
        this.$notify({
          type: "warning",
          title: "淘气👻",
          message: "你还没设置头像呢~",
          position: "top-left",
          offset: 50,
        });
        return;
      }
      if (this.friend.url.trim() === "") {
        this.$notify({
          type: "warning",
          title: "淘气👻",
          message: "你还没写网址呢~",
          position: "top-left",
          offset: 50,
        });
        return;
      }
      if (!this.friend.url.includes("http")) {
        this.$notify({
          type: "warning",
          title: "淘气👻",
          message: "请填写完整的有效网址，例如：http://****.com",
          position: "top-left",
          offset: 50,
        });
        return;
      }
      this.$http
        .post(
          this.$constant.baseURL + "/webInfo/saveResourcePath/",
          this.friend
        )
        .then((res) => {
          $(".form-wrap").css({ height: "447px", top: "0" });
          this.$notify({
            title: "可以啦🍨",
            message: "提交成功，待管理员审核！",
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
    clickFriend(path) {
      if (path.includes("http")) {
        window.open(path);
      }
    },
    getFriends() {
      this.$http
        .post(
          this.$constant.baseURL + "/webInfo/clistResourcePath/",
          this.pagination
        )
        .then((res) => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.friendList = res.result[0].records;
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
  },
};
</script>
<style lang="scss" scoped>
.friend-head {
  height: 300px;
  position: relative;
  &::before {
    position: absolute;
    width: 100%;
    height: 100%;
    content: "";
  }
}
.friend-main {
  margin: 40px auto 0;
  padding: 40px;
  background: var(--background);
  h2 {
    font-size: 20px;
  }
}
hr {
  position: relative;
  margin: 40px auto;
  border: 2px dashed var(--blue2);
  overflow: visible;
  &:before {
    position: absolute;
    top: -21px;
    left: 5%;
    color: var(--red);
    content: "\e673";
    font-size: 40px;
    line-height: 1;
    transition: all 1s ease-in-out;
    font-family: iconfont;
  }
  &:hover:before {
    left: calc(95% - 20px);
  }
}
.form-wrap {
  margin: 0 auto;
  overflow: hidden;
  width: 530px;
  height: 447px;
  position: relative;
  top: 0;
  transition: all 1s ease-in-out 0.3s;
  z-index: 0;
}
.before-img {
  position: absolute;
  bottom: 126px;
  left: 0;
  background-repeat: no-repeat;
  width: 530px;
  height: 317px;
  z-index: -100;
}
.after-img {
  position: absolute;
  bottom: -2px;
  left: 0;
  background-repeat: no-repeat;
  width: 530px;
  height: 259px;
  z-index: 100;
}
.envelope {
  position: relative;
  margin: 0 auto;
  transition: all 1s ease-in-out 0.3s;
  padding: 200px 20px 20px;
}
.form-main {
  background: var(--pink2);
  margin: 0 auto;
  border-radius: 10px;
  overflow: hidden;
  transition: all 0.3s ease;
  border: 1px solid var(--gray1);
  &:hover {
    border-color: var(--gray4);
  }
}
.user-title {
  text-align: right;
  user-select: none;
  div {
    height: 55px;
    line-height: 55px;
    text-align: center;
  }
}
.user-content {
  text-align: left;
  > div {
    height: 55px;
    display: flex;
    align-items: center;
  }
  ::v-deep .el-input__inner {
    border: none;
    height: 35px;
    background: var(--whiteMask);
  }
}
.form-friend {
  margin-top: 12px;
  padding: 20px 0;
}
::v-deep .recommendFriend .card-item::before {
  line-height: 62px;
  content: "推荐";
  position: absolute;
  z-index: 10;
  color: var(--white);
  top: -15px;
  letter-spacing: 3px;
  left: 15px;
  font-size: 15px;
  width: 44px;
  height: 50px;
  display: flex;
  justify-content: center;
  background: linear-gradient(90deg, var(--orange), var(--orange6));
  border-radius: 0px 0px 12px 12px;
}
@media screen and (max-width: 700px) {
  .form-wrap {
    width: auto;
  }
  .before-img {
    width: auto;
  }
  .after-img {
    width: auto;
  }
}
@media screen and (max-width: 500px) {
  .friend-main {
    padding: 40px 15px;
  }
  .friend-wrap {
    padding: 0 10px;
  }
}
</style>
