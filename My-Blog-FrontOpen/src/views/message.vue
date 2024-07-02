<template>
  <div>
    <div>
      <div
        style="animation: header-effect 2s"
        :style="{ background: `${$store.state.changeBg}` }"
        class="background-image background-image-changeBg"
      ></div>
      <!-- 输入框 -->
      <div class="message-in" style="text-align: center">
        <h2 class="message-title">树洞</h2>
        <h4 class="message-title">想说的·想问的·吐槽·交流</h4>
        <div>
          <input
            ref="focus"
            class="message-input"
            type="text"
            style="outline: none; width: 70%"
            placeholder="留下点什么啦~"
            v-model="messageContent"
            @click="show = true"
            maxlength="60"
          />
          <button
            v-show="show"
            style="margin-left: 12px; width: 20%"
            @click="submitMessage"
            class="message-input"
          >
            发射
          </button>
        </div>
      </div>
      <!-- 弹幕 -->
      <div class="barrage-container">
        <vue-baberrage :barrageList="barrageList" :loop="true"></vue-baberrage>
      </div>
    </div>
    <div class="comment-wrap">
      <div class="comment-content">
        <comment
          :source="0"
          :type="'message'"
          @getProhibitedWordsList="getProhibitedWordsList"
        ></comment>
      </div>
      <myFooter></myFooter>
    </div>
  </div>
</template>

<script>
const comment = () => import("./common/comment");
const myFooter = () => import("./common/myFooter");

export default {
  components: {
    comment,
    myFooter,
  },
  data() {
    return {
      show: false,
      messageContent: "",
      barrageList: [],
      prohibitedWordsList: [],
      pagination: {
        current: 1,
        size: 10,
        total: 0,
        searchKey: "",
      },
    };
  },
  created() {
    this.getProhibitedWordsList();
    this.getTreeHole();
  },
  mounted() {
    this.$refs.focus.focus();
  },
  methods: {
    getProhibitedWordsList(prohibitedWordsList) {
      this.prohibitedWordsList = prohibitedWordsList;
    },
    getTreeHole() {
      this.$http
        .post(
          this.$constant.baseURL + "/admin/treeHole/boss/list/",
          {
            size: 0,
            current: 1,
          },
          true
        )
        .then((res) => {
          if (!this.$common.isEmpty(res.result[0])) {
            res.result[0].records.forEach((m) => {
              this.barrageList.push({
                id: m.id,
                avatar: m.avatar,
                msg: m.message,
                time: Math.floor(Math.random() * 10 + 5),
              });
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
    submitMessage() {
      if (this.messageContent.trim() === "") {
        this.$message({
          message: "你还没写呢~",
          type: "warning",
        });
        return;
      }
      let treeHole = {
        message: this.messageContent.trim(),
      };
      let newArr = [];
      let list = [];
      let word = treeHole.message.split("");
      this.prohibitedWordsList.forEach((e) => list.push(e.message));
      for (let i = 0; i < list.length; i++) {
        for (let j = 0; j < word.length; j++) {
          if (word[j] === list[i]) {
            newArr.push(word[j]);
          }
        }
      }
      if (newArr.length > 0) {
        this.$message({
          message: "你发的留言带有违禁词！请发一条友好的留言~~~",
          type: "error",
        });
        return;
      }
      if (this.$common.isEmpty(this.$store.state.currentUser)) {
        this.$message({
          message: "你还没有登录呢~~~",
          type: "error",
        });
        return;
      }
      if (
        !this.$common.isEmpty(this.$store.state.currentUser) &&
        !this.$common.isEmpty(this.$store.state.currentUser.avatar) &&
        !this.$common.isEmpty(this.$store.state.currentUser.username)
      ) {
        treeHole.avatar = this.$store.state.currentUser.avatar;
        treeHole.username = this.$store.state.currentUser.username;
      }
      this.$http
        .post(this.$constant.baseURL + "/webInfo/saveTreeHole/", treeHole)
        .then((res) => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.$message({
              message: "发射成功！",
              type: "success",
            });
            this.barrageList.push({
              id: res.result[0].records[0].id,
              avatar: res.result[0].records[0].avatar,
              msg: res.result[0].records[0].message,
              time: Math.floor(Math.random() * 10 + 5),
            });
          }
        })
        .catch((error) => {
          this.$message({
            message: error.message,
            type: "error",
          });
        });
      this.messageContent = "";
      this.show = false;
    },
  },
};
</script>

<style scoped>
.message-in {
  position: absolute;
  left: 50%;
  top: 40%;
  transform: translate(-50%, -50%);
  color: #39c5bb;
  animation: hideToShow 2.5s;
  width: 360px;
  z-index: 10;
}

.message-title {
  user-select: none;
  text-align: center;
  font-size: 24px;
  font-weight: 400;
}

.message-input {
  border-radius: 1.2rem;
  border: var(--lightRed) 1px solid;
  color: var(--lightRed);
  background: var(--transparent);
  padding: 10px 10px;
}

.message-input::-webkit-input-placeholder {
  color: var(--lightRed);
}

.barrage-container {
  position: absolute;
  top: 50px;
  left: 0;
  right: 0;
  bottom: 0;
  height: calc(100% - 50px);
  width: 100%;
  user-select: none;
  overflow: hidden;
}

.comment-wrap {
  background: var(--background);
  position: absolute;
  top: 100vh;
  width: 100%;
}

.comment-content {
  max-width: 900px;
  margin: 0 auto;
  padding: 40px 20px;
}
</style>
