<template>
  <div>
    <!-- 两句诗 -->
    <div class="my-animation-slide-top">
      <twoPoem :isShehui="true">
        <div style="padding: 40px 0 10px" class="my-animation-slide-bottom">
          <h1 style="text-align: center; color: #39c5bb">
            ♪(^∇^*) 欢迎来到{{
              $store.state.webInfo.webName
            }}音乐平台，请尽情欣赏~~
          </h1>
          <h3 style="text-align: center; color: #39c5bb">
            还想听更多音乐吗？请移步<a
              class="my-music"
              style="color: var(--red)"
              href="https://www.zjh2002.icu/"
              >我的云音乐</a
            >
          </h3>
          <!-- 音乐 -->
          <div class="funny-wrap" v-if="!$common.isEmpty(funnys)">
            <div v-for="(item, index) in funnys" :key="index">
              <div style="display: flex">
                <span class="iconRotate">
                  <img src="../assets/svg/windmill.svg" />
                </span>
                <span class="funny-title"
                  >{{ index + 1 }}号厅：{{ item.classify }}</span
                >
              </div>
              <div class="process-wrap">
                <el-collapse
                  v-model="activeName"
                  accordion
                  @change="changeFunny(item.classify)"
                >
                  <el-collapse-item
                    :title="`🎻 要来${index + 1}号厅吗，为您播放<${
                      item.classify
                    }>歌单`"
                    :name="index"
                  >
                    <div
                      class="my-animation-slide-bottom"
                      style="display: flex; flex-flow: wrap; margin-left: 20px"
                      v-if="!$common.isEmpty(item.data)"
                    >
                      <div
                        style="width: 150px"
                        v-for="(funny, i) in item.data"
                        :key="i"
                      >
                        <el-image
                          class="funny-avatar myCenter"
                          lazy
                          :size="110"
                          style="margin: 20px"
                          @click.native="playSound(funny.url, item.data, i)"
                          :src="funny.cover"
                        >
                        </el-image>
                        <div class="funny-item-title">{{ funny.title }}</div>
                      </div>
                    </div>
                  </el-collapse-item>
                </el-collapse>
                <hr />
              </div>
            </div>
          </div>
        </div>
        <!-- 页脚 -->
        <div>
          <myFooter></myFooter>
        </div>
      </twoPoem>
    </div>
  </div>
</template>

<script>
const myFooter = () => import("./common/myFooter");
const twoPoem = () => import("./common/twoPoem");
export default {
  components: {
    twoPoem,
    myFooter,
  },
  data() {
    return {
      pagination: {
        current: 1,
        size: 9999,
        order: "title",
        desc: false,
        resourceType: "funny",
        classify: "",
      },
      activeName: 0,
      audio: null,
      playList: null,
      index: null,
      funnys: [
        {
          classify: "",
          count: null,
          data: [
            {
              classify: "",
              cover: "",
              url: "",
              title: "",
            },
          ],
        },
      ],
      funny: {
        classify: "",
        title: "",
        cover: "",
        url: "",
      },
    };
  },
  created() {
    this.getFunny();
  },
  beforeDestroy() {
    if (this.audio != null && !this.audio.paused) {
      this.audio.pause();
    }
  },
  methods: {
    getFunny() {
      this.$http
        .get(this.$constant.baseURL + "/webInfo/getClassifyList/", {
          type: "funny",
        })
        .then((res) => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.funnys = res.result[0].data;
            this.changeFunny(this.funnys[0].classify);
          }
        })
        .catch((error) => {
          this.$message({
            message: error.message,
            type: "error",
          });
        });
    },
    listFunny() {
      this.$http
        .post(
          this.$constant.baseURL + "/webInfo/clistResourcePath/",
          this.pagination
        )
        .then((res) => {
          if (
            !this.$common.isEmpty(res.result[0]) &&
            !this.$common.isEmpty(res.result[0].records)
          ) {
            this.funnys.forEach((funny) => {
              if (funny.classify === this.pagination.classify) {
                funny.data = res.result[0].records;
                this.$forceUpdate();
              }
            });
          }
          this.pagination.classify = "";
        })
        .catch((error) => {
          this.$message({
            message: error.message,
            type: "error",
          });
        });
    },
    changeFunny(classify) {
      this.funnys.forEach((funny) => {
        if (funny.classify === classify && this.$common.isEmpty(funny.data)) {
          this.pagination.classify = classify;
          this.listFunny();
        }
      });
    },
    playSound(src, playList, index) {
      // this.audio.volume = 0.1
      this.playList = playList;
      this.index = index;
      if (this.audio != null) {
        if (this.audio.src === src) {
          if (this.audio.paused) {
            this.audio.play();
          } else {
            this.audio.pause();
          }
        } else {
          this.audio.pause();
          this.audio.src = src;
          this.audio.load();
          this.audio.play();
        }
      } else {
        this.audio = new Audio(src);
        this.audio.play();
        this.audio.onended = () => {
          this.index = this.index + 1;
          if (this.index < this.playList.length) {
            this.audio.src = this.playList[this.index].url;
            this.audio.load();
            setTimeout(() => {
              this.audio.play();
            }, 3000);
          }
        };
      }
    },
  },
};
</script>

<style scoped>
::v-deep .background-image-changeBg {
  background-attachment: inherit !important;
  width: unset !important;
}
::v-deep .poem-container {
  display: block;
  overflow: scroll;
}
::v-deep .poem-wrap {
  margin: 0 auto;
}
.my-music {
  background: var(--gradientAnimation);
  background-size: 0px 3px;
  transition: background-size 800ms;
}
.my-music:hover {
  background-position-x: left;
  background-size: 100% 3px;
}
.funny-wrap {
  background-color: var(--background);
  width: 99%;
  border-radius: 10px;
  max-width: 1600px;
  margin: 0 auto;
  padding: 40px 20px 80px;
}

.funny-title {
  color: wheat;
  font-size: 28px;
  font-weight: 400;
  margin-left: 12px;
}

.process-wrap {
  margin: 20px 0 40px;
}

.process-wrap hr {
  position: relative;
  margin: 30px auto 100px;
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
  font-size: 16px;
  font-weight: 400;
  background-color: rgba(255, 233, 236, 0.4);
  color: white;
  padding: 40px;
  line-height: 20px;
}

.process-wrap >>> .el-collapse-item__wrap {
  background-color: rgba(255, 233, 236, 0.7);
}

.process-wrap .el-collapse {
  border-top: unset;
  border-bottom: unset;
  border-radius: 10px;
  overflow: hidden;
}

.process-wrap >>> .el-collapse-item__wrap {
  border-bottom: unset;
}

.funny-item-title {
  text-align: center;
  margin: 0 10px;
  font-size: 18px;
  font-weight: 400;
  white-space: nowrap;
  text-overflow: ellipsis;
  overflow: hidden;
  color: var(--bigRed);
}

.funny-item-title:hover {
  text-overflow: unset;
  overflow: unset;
}

.funny-avatar {
  transition: all 0.5s;
  user-select: none;
  border-radius: 50%;
}

.funny-avatar:hover {
  transform: rotate(360deg);
}
::v-deep .poem-container {
  padding-bottom: 0 !important;
}
</style>
