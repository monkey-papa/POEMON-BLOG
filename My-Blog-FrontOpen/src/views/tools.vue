<template>
  <div>
    <div class="tools-container" style="background: var(--background);">
      <div style="animation: header-effect 2s" :style="{ background: `${$store.state.changeBg}` }" class="background-image background-image-changeBg"></div>
      <!-- 封面 -->
      <div class="tools-header my-animation-slide-top">
        <!-- 顶部视频 -->
        <video class="index-video" autoplay="autoplay" muted="muted" loop="loop" src="https://www.wulihub.com.cn/gc/Q5ND2N/1920.js"></video>
        <div style="position: absolute; left: 0; top: 0; padding: 100px 20px 0px">
          <!-- 标题 -->
          <div style="color: var(--green1); margin: 10px">
            <div style="margin-left: 10px">记录</div>
            <div style="font-size: 36px; font-weight: bold; line-height: 2">
              百宝箱
            </div>
          </div>
          <div class="card-container">
            <!-- 收藏夹 -->
            <div @click="changeFavorite(1)" class="card-item">
              <div class="tools-image"></div>
              <div style="
                  position: absolute;
                  left: 0;
                  top: 0;
                  padding: 20px 25px 15px;
                ">
                <div class="card-name">生产力</div>
                <div class="card-desc">希望这些工具对你有帮助喔ヾ(≧▽≦*)o</div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <!-- 内容 -->
      <div class="tools-content my-animation-slide-bottom">
        <!-- 收藏夹 -->
        <div v-show="card === 1 && !$common.isEmpty(collects)">
          <div v-for="(value, key) in collects" :key="key" style="margin-top: 20px">
            <div class="collect-classify">
              {{ key }}
            </div>
            <div class="tools-item-wrap">
              <div v-for="(item, index) in value" :key="index" @click="toUrl(item.url)" class="tools-item">
                <div>
                  <el-avatar class="tools-item-image" style="background: var(--transparent)" :size="60" :src="item.cover">
                  </el-avatar>
                </div>
                <div style="width: calc(100% - 80px)">
                  <div class="tools-item-title">
                    {{ item.title }}
                  </div>
                  <div class="tools-item-introduction">
                    {{ item.introduction }}
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      card: null,
      collects: {}
    }
  },
  created() {
    this.card = 1
    this.getCollect()
  },
  methods: {
    toUrl(url) {
      window.open(url)
    },
    changeFavorite(card) {
      this.card = card
      if (card === 1) {
        if (this.$common.isEmpty(this.collects)) {
          this.getCollect()
        }
      }
    },
    getCollect() {
      this.$http
        .get(this.$constant.baseURL + '/webInfo/listCollect/')
        .then(res => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.collects = res.result[0].records
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
    }
  }
}
</script>
<style lang="scss" scoped>
.tools-image::before {
  content: '';
  position: absolute;
  width: 100%;
  height: 100%;
  background-color: var(--pink);
}
.card-container {
  display: flex;
  flex-wrap: wrap;
  margin-top: 60px;
}
.card-item {
  transition: all 0.3s;
  position: relative;
  width: 250px;
  height: 120px;
  border-radius: 20px;
  animation: hideToShow 1s ease-in-out;
  overflow: hidden;
  margin: 10px;
  color: var(--darkBlue);
  &:hover {
    transform: translateY(-6px);
  }
}
.card-name {
  font-weight: 400;
  font-size: 25px;
  &:after {
    top: 50px;
    width: 22px;
    left: 26px;
    height: 2px;
    background: var(--white);
    content: '';
    border-radius: 1px;
    position: absolute;
  }
}
.card-desc {
  font-weight: bold;
  margin-top: 15px;
}
.tools-content {
  margin: 0 auto;
  max-width: 1200px;
}
.collect-classify {
  color: var(--bigRed);
  font-size: 28px;
  margin-bottom: 10px;
}
.tools-item-wrap {
  justify-content: center;
  display: flex;
  flex-wrap: wrap;
  margin-left: -10px;
}
.tools-item {
  color: var(--bigRed);
  transition: all 0.3s;
  border-radius: 12px;
  box-shadow: 0 8px 16px -4px var(--mask);
  background: var(--pink);
  display: flex;
  width: calc(100% / 4 - 20px);
  max-width: 320px;
  height: 90px;
  overflow: hidden;
  padding: 15px;
  margin: 10px;
  &:hover {
    background: var(--blue);
    color: var(--white);
    .tools-item-image {
      transition: all 0.6s;
      width: 0 !important;
      height: 0 !important;
      opacity: 0;
      margin-right: 0;
    }
    div:nth-child(2) {
      width: 100% !important;
    }
  }
  &-image {
    margin-right: 20px;
    transition: all 0.3s;
  }
  &-title {
    font-size: 19px;
    font-weight: bold;
    white-space: nowrap;
    text-overflow: ellipsis;
    overflow: hidden;
    margin-bottom: 5px;
  }
  &-introduction {
    opacity: 0.7;
    font-weight: bold;
    letter-spacing: 1px;
    font-size: 14px;
    line-height: 1.2;
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
  }
}
@media screen and (max-width: 1100px) {
  .index-video {
    width: 1100px;
  }
}
@media screen and (max-width: 900px) {
  .tools-item {
    width: calc(100% / 3 - 20px);
  }
}
@media screen and (max-width: 700px) {
  .tools-item {
    width: calc(100% / 2 - 20px);
  }
}
@media screen and (max-width: 620px) {
  .card-container {
    margin-top: 0;
  }
  .index-video {
    width: 620px;
  }
}
@media screen and (max-width: 400px) {
  .tools-item {
    width: calc(100% - 20px);
  }
  .index-video {
    width: 400px;
  }
  .card-container {
    display: none;
  }
}
@media screen and (max-width: 308px) {
  .card-item {
    display: none;
  }
}
</style>
