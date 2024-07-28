<template>
  <div>
    <div style="animation: header-effect 2s" :style="{ background: `${$store.state.changeBg}` }" class="background-image background-image-changeBg"></div>
    <!-- 两句诗 -->
    <div class="my-animation-slide-top">
      <twoPoem :isShehui="true"> </twoPoem>
    </div>
    <div style="padding-top: 40px; background: var(--background)" class="my-animation-slide-bottom">
      <h1 style="text-align: center; color: var(--green3)">
        ♪(^∇^*) 欢迎来到{{ $store.state.webInfo.webName }}音乐平台，请尽情欣赏~~
      </h1>
      <h3 style="text-align: center; color: var(--green3)">
        还想听更多音乐吗？请移步<a class="my-music" style="color: var(--red)" href="https://www.zjh2002.icu/">我的云音乐</a>
      </h3>
      <!-- 音乐 -->
      <div class="funny-wrap" v-if="!$common.isEmpty(funnys)">
        <div v-for="(item, index) in funnys" :key="index">
          <div style="display: flex;align-items: center;">
            <span style="height: 16px;
    width: 16px;" class="iconRotate">
              <i style="display:flex;height: 16px;
    width: 16px;color:var(--red)" class="iconfont icon-fengche"></i>
            </span>
            <span class="funny-title">{{ index + 1 }}号厅：{{ item.classify }}</span>
          </div>
          <div class="process-wrap">
            <el-collapse class="shadow-box" v-model="activeName" accordion @change="changeFunny(item.classify)">
              <el-collapse-item :title="`🎻 要来${index + 1}号厅吗，为您播放<${
                  item.classify
                }>歌单`" :name="index">
                <div class="my-animation-slide-bottom" style="display: flex; flex-flow: wrap; margin-left: 20px" v-if="!$common.isEmpty(item.data)">
                  <div style="width: 150px" v-for="(funny, i) in item.data" :key="i">
                    <el-image class="funny-avatar myCenter" lazy :size="110" style="margin: 20px" @click.native="playSound(funny.url, item.data, i)" :src="funny.cover">
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
  </div>
</template>

<script>
const twoPoem = () => import('./common/twoPoem')
export default {
  components: {
    twoPoem
  },
  data() {
    return {
      pagination: {
        current: 1,
        size: 9999,
        order: 'title',
        desc: false,
        resourceType: 'funny',
        classify: ''
      },
      activeName: 0,
      audio: null,
      playList: null,
      index: null,
      funnys: [
        {
          classify: '',
          count: null,
          data: [
            {
              classify: '',
              cover: '',
              url: '',
              title: ''
            }
          ]
        }
      ],
      funny: {
        classify: '',
        title: '',
        cover: '',
        url: ''
      }
    }
  },
  created() {
    this.getFunny()
  },
  beforeDestroy() {
    if (this.audio != null && !this.audio.paused) {
      this.audio.pause()
    }
  },
  methods: {
    getFunny() {
      this.$http
        .get(this.$constant.baseURL + '/webInfo/getClassifyList/', {
          type: 'funny'
        })
        .then(res => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.funnys = res.result[0].data
            this.changeFunny(this.funnys[0].classify)
          }
        })
        .catch(error => {
          this.$notify({
            type: 'error',
            title: '可恶🤬',
            message: error.message,
            position: 'top-left',
            offset: 50
          })
        })
    },
    listFunny() {
      this.$http
        .post(this.$constant.baseURL + '/webInfo/clistResourcePath/', this.pagination)
        .then(res => {
          if (!this.$common.isEmpty(res.result[0]) && !this.$common.isEmpty(res.result[0].records)) {
            this.funnys.forEach(funny => {
              if (funny.classify === this.pagination.classify) {
                funny.data = res.result[0].records
                this.$forceUpdate()
              }
            })
          }
          this.pagination.classify = ''
        })
        .catch(error => {
          this.$notify({
            type: 'error',
            title: '可恶🤬',
            message: error.message,
            position: 'top-left',
            offset: 50
          })
        })
    },
    changeFunny(classify) {
      this.funnys.forEach(funny => {
        if (funny.classify === classify && this.$common.isEmpty(funny.data)) {
          this.pagination.classify = classify
          this.listFunny()
        }
      })
    },
    playSound(src, playList, index) {
      // this.audio.volume = 0.1
      this.playList = playList
      this.index = index
      if (this.audio != null) {
        if (this.audio.src === src) {
          if (this.audio.paused) {
            this.audio.play()
          } else {
            this.audio.pause()
          }
        } else {
          this.audio.pause()
          this.audio.src = src
          this.audio.load()
          this.audio.play()
        }
      } else {
        this.audio = new Audio(src)
        this.audio.play()
        this.audio.onended = () => {
          this.index = this.index + 1
          if (this.index < this.playList.length) {
            this.audio.src = this.playList[this.index].url
            this.audio.load()
            setTimeout(() => {
              this.audio.play()
            }, 3000)
          }
        }
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.icon-fengche:before {
  height: 16px;
  width: 16px;
}
.my-music {
  background: var(--gradientAnimation);
  background-size: 0px 3px;
  transition: background-size 800ms;
  &:hover {
    background-position-x: left;
    background-size: 100% 3px;
  }
}
.funny-wrap {
  width: 99%;
  border-radius: 10px;
  max-width: 1600px;
  margin: 0 auto;
  padding: 40px 20px 80px;
}
.funny-title {
  color: var(--yellow3);
  font-size: 28px;
  font-weight: 400;
  margin-left: 12px;
  &:hover {
    color: var(--yellow6);
  }
}
.process-wrap {
  margin: 20px 0 40px;
  hr {
    position: relative;
    margin: 30px auto 100px;
    border: 2px dashed var(--blue2);
    overflow: visible;
    &:before {
      position: absolute;
      top: -20px;
      left: 5%;
      color: var(--red);
      content: '\e673';
      font-size: 40px;
      line-height: 1;
      transition: all 1s ease-in-out;
      font-family: iconfont;
    }
    &:hover:before {
      left: calc(95% - 20px);
    }
  }
  ::v-deep .el-collapse-item__header {
    border-bottom: unset;
    font-size: 16px;
    font-weight: 400;
    background-color: var(--gray5);
    color: var(--fontColor);
    padding: 40px;
    line-height: 20px;
    &:hover {
      color: var(--black8);
    }
  }
  ::v-deep .el-collapse-item__wrap {
    background-color: var(--maxMaxWhiteMask);
  }
  .el-collapse {
    border: 1px solid var(--gray1);
    border-radius: 10px;
    overflow: hidden;
    transition: all 0.3s ease;
    &:hover {
      border-color: var(--gray4);
    }
  }
  ::v-deep .el-collapse-item__wrap {
    border-bottom: unset;
  }
}
.funny-item-title {
  text-align: center;
  margin: 0 10px;
  font-size: 18px;
  font-weight: 400;
  white-space: nowrap;
  text-overflow: ellipsis;
  overflow: hidden;
  color: var(--bigRed1);
  &:hover {
    text-overflow: unset;
    overflow: unset;
    color: var(--bigRed);
  }
}
.funny-avatar {
  transition: all 0.5s;
  user-select: none;
  border-radius: 50%;
  &:hover {
    transform: rotate(360deg);
  }
}
</style>
