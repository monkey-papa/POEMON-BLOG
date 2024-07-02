<template>
  <div>
    <div class="travel-container">
      <!-- 背景图片 -->
      <div style="animation: header-effect 2s;" :style="{ background: `${$store.state.changeBg}` }" class="background-image background-image-changeBg"></div>
      <div class="travel-top">
        <div class="travel-header my-animation-slide-top">
          <!-- 顶部视频 -->
          <video class="index-video" autoplay muted loop playsinline webkit-playsinline :src="$constant.favoriteVideo"></video>
          <div style="position: absolute; left: 20px; top: 20px">
            <!-- 标题 -->
            <div>
              <div style="text-indent: 2em;">旅拍集</div>
              <div class="travel-introduce" style="
                  font-size: 36px;
                  font-weight: 400;
                  line-height: 1.5;
                ">
                这里是相册分类集
              </div>
            </div>
          </div>
          <div style="position: absolute; left: 20px; bottom: 40px; margin: 10px">
            生活中的小确幸。
          </div>
        </div>
        <div v-if="!$common.mobile() && !mobile" class="travel-header-right" style="
            background-image: url(https://www.qiniuyun.gif);
          ">
          <div class="card-content">
            <div class="author-content-item-tips">┗|｀O′|┛ 嗷～～</div>
            <span class="author-content-item-title">cat</span>
            <div class="banner-button-group">
              <a href="https://www.zjh2002.icu" class="banner-button">
                <span class="banner-button-text">🎻 秘密基地</span>
              </a>
            </div>
          </div>
        </div>
      </div>
      <div class="travel-content my-animation-slide-bottom">
        <!-- 标签 -->
        <div class="photo-title-warp" v-if="!$common.isEmpty(photoTitleList)">
          <div v-for="(item, index) in photoTitleList" :key="index" :class="{ isActive: photoPagination.classify === item.classify }" @click="changePhotoTitle(item.classify)">
            <proTag :info="item.classify + ' ' + item.count" :color="$constant.before_color_list[Math.floor(Math.random() * 6)]
              " style="margin: 12px">
            </proTag>
          </div>
        </div>
        <div class="photo-title">
          {{ photoPagination.classify }}
        </div>
        <photo :resourcePathList="photoList"></photo>
        <div class="pagination-wrap">
          <div @click="pagePhotos()" class="pagination" v-if="photoPagination.total !== photoList.length">
            下一页
          </div>
          <div v-else style="user-select: none;color: var(--red);">~~到底啦~~</div>
        </div>
      </div>
    </div>
    <!-- 页脚 -->
    <myFooter></myFooter>
  </div>
</template>
<script>
const myFooter = () => import('./common/myFooter')
const photo = () => import('./common/photo')
const proTag = () => import('./common/proTag')
export default {
  components: {
    photo,
    proTag,
    myFooter
  },
  data() {
    return {
      photoPagination: {
        current: 1,
        size: 10,
        total: 0,
        resourceType: 'lovePhoto',
        classify: ''
      },
      photoTitleList: [],
      photoList: [],
      mobile: false
    }
  },
  created() {
    this.getPhotoTitles()
    this.mobile = document.body.clientWidth < 780
    window.addEventListener('resize', () => {
      let docWidth = document.body.clientWidth
      if (docWidth < 780) {
        this.mobile = true
      } else {
        this.mobile = false
      }
    })
  },
  methods: {
    getPhotoTitles() {
      this.$http
        .get(this.$constant.baseURL + '/webInfo/getClassifyList/', { type: 'lovePhoto' })
        .then(res => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.photoTitleList = res.result[0].data
            this.photoPagination = {
              current: 1,
              size: 10,
              total: 0,
              resourceType: 'lovePhoto',
              classify: this.photoTitleList[0].classify
            }
            this.changePhoto()
          }
        })
        .catch(error => {
          this.$message({
            message: error.message,
            type: 'error'
          })
        })
    },
    changePhotoTitle(classify) {
      if (classify !== this.photoPagination.classify) {
        this.photoPagination = {
          current: 1,
          size: 10,
          total: 0,
          resourceType: 'lovePhoto',
          classify: classify
        }
        this.photoList = []
        this.changePhoto()
      }
    },
    pagePhotos() {
      this.photoPagination.current = this.photoPagination.current + 1
      this.changePhoto()
    },
    changePhoto() {
      this.$http
        .post(this.$constant.baseURL + '/webInfo/clistResourcePath/', this.photoPagination)
        .then(res => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.photoList = this.photoList.concat(res.result[0].records)
            this.photoPagination.total = res.result[0].total
          }
        })
        .catch(error => {
          this.$message({
            message: error.message,
            type: 'error'
          })
        })
    }
  }
}
</script>
<style scoped>
.travel-header {
  margin: 60px 10px 30px;
  height: 300px;
  overflow: hidden;
  border-radius: 20px;
  width: 1130px;
  color: var(--red);
  user-select: none;
}
.index-video {
  width: 100%;
  height: 180%;
  object-fit: cover;
}
.travel-content {
  margin: 0 auto;
  max-width: 1550px;
}
.photo-title-warp {
  max-width: 1250px;
  margin: 15px auto;
  padding: 20px;
  border-radius: 10px;
  display: flex;
  flex-wrap: wrap;
}
.isActive {
  animation: scale 2.5s ease-in-out infinite;
}
.photo-title {
  color: var(--red);
  text-align: center;
  font-size: 30px;
  font-weight: 700;
  line-height: 80px;
  letter-spacing: 2px;
}
.pagination-wrap {
  display: flex;
  justify-content: center;
  margin-top: 40px;
}
.pagination {
  padding: 13px 15px;
  border: 1px solid var(--lightGray);
  border-radius: 3rem;
  color: var(--greyFont);
  width: 100px;
  user-select: none;
  text-align: center;
}
@media screen and (max-width: 1150px) {
  .photo-title-warp {
    max-width: 780px;
  }
}
.travel-top {
  display: flex;
  justify-content: center;
}
.travel-header-right {
  right: 0;
  margin: 60px 10px 30px;
  height: 300px;
  overflow: hidden;
  border-radius: 20px;
  width: 270px;
  color: var(--red);
  user-select: none;
  position: relative;
  background-color: brown;
  background-repeat: no-repeat;
  background-position: center;
  background-size: contain;
  color: var(--red);
}
.card-content {
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  z-index: 2;
  display: flex;
  flex-direction: column;
  padding: 1rem 2rem;
}
.author-content-item-tips {
  font-size: 1.2em;
  margin-bottom: 0.5rem;
}
.banner-button-group {
  position: absolute;
  bottom: 0.3rem;
  right: 0.5rem;
}
.banner-button {
  color: var(--red);
  height: 2rem;
  width: 6rem;
  border-radius: 1.2rem;
  background: rgba(255, 255, 255, 0.8);
  display: flex;
  align-items: center;
  justify-content: center;
  text-decoration: none;
  outline: none;
}
</style>