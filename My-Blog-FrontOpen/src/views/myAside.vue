<template>
  <div>
    <!-- 网站信息 -->
    <div class="card-content1 shadow-box">
      <el-avatar style="margin-top: 20px" class="user-avatar" :size="120" :src="$store.state.webInfo.randomCover[12]"></el-avatar>
      <div class="web-name">{{ webInfo.webName }}</div>
      <div class="web-info">
        <div class="blog-info-box">
          <span>文章</span>
          <span class="blog-info-num">{{ $store.getters.articleTotal }}</span>
        </div>
        <div class="blog-info-box">
          <span>分类</span>
          <span class="blog-info-num">{{ sortInfo.length }}</span>
        </div>
        <div class="blog-info-box">
          <span>访问量</span>
          <span class="blog-info-num">{{ total_sum }}</span>
        </div>
      </div>
      <a class="collection-btn" @click="showTip()">
        <i class="el-icon-loading" style="
            margin-right: 5px;
            color: var(--maxLightRed);
            font-size: larger;
          "></i>微言圈
      </a>
      <div class="my-link a-svg">
        <a href="https://juejin.cn/user/3204412407287917/posts" target="_blank">
          <img src="../assets/svg/juejin.svg" />
        </a>
        <a href="https://blog.csdn.net/qq_53461589?spm=1038.2274.3001.5343" target="_blank">
          <img src="../assets/svg/csdn.svg" />
        </a>
        <a href="https://github.com/monkey-papa/POEMON-BLOG" target="_blank">
          <img src="../assets/svg/github.svg" />
        </a>
        <a href="https://www.zjh2002.icu">
          <img src="../assets/svg/cloudMusic.svg" />
        </a>
      </div>
    </div>
    <!-- 天气时间 -->
    <div v-if="$route.path !== '/article'" class="wow card-widget" style="
        margin-top: 20px;
        position: relative;
        padding: 25px 10px 5px 17px;
        border-radius: 18px;
        animation: hideToShow 1s ease-in-out;
      ">
      <div class="card-content2-title">
        <i class="fa fa-thermometer-3 card-content2-icon"></i>
        <span>小窝天气</span>
      </div>
      <div style="line-height: 25px; word-break: break-all; color: var(--darkBlue)">
        <div>
          欢迎来自<span style="color: var(--bigRed)">{{
            city || "五湖四海"
          }}</span>的小伙伴，让我们嗨起来叭！
        </div>
        <div>{{ city }}</div>
        <div>
          <span style="color: var(--bigRed)">今天（周{{ weather.week }}）: {{ weather.nighttemp }}℃~{{
              weather.daytemp
            }}℃ {{ weather.dayweather }} 🎈</span>
        </div>
        <div>{{ city }}</div>
        <div>
          <span style="color: var(--bigRed)">明天（周{{ tomWeather.week }}）: {{ tomWeather.nighttemp }}℃~{{
              tomWeather.daytemp
            }}℃ {{ tomWeather.dayweather }} 🎈</span>
        </div>
      </div>
    </div>
    <!-- 最近更新 -->
    <div v-if="!$common.isEmpty(latelyArticles)" style="
        position: relative;
        padding: 25px;
        border-radius: 18px;
        margin-top: 20px;
        animation: hideToShow 1s ease-in-out;
      " class="card-widget wow">
      <div class="card-content2-title">
        <i class="fa fa-history card-content2-icon"></i>
        <span>最近更新</span>
      </div>
      <div v-for="(article, index) in latelyArticles" :key="index" @click="$router.push({ path: '/article', query: { id: article.id } })">
        <div class="aside-post-detail">
          <div class="aside-post-image">
            <el-image lazy class="my-el-image" :src="article.articleCover" fit="cover">
              <!-- 懒加载颜色 -->
              <div slot="placeholder">
                <div class="dot" :class="{
                    leftImage: index % 2 !== 0,
                    rightImage: index % 2 === 0,
                  }"></div>
              </div>
              <div slot="error" class="image-slot">
                <div class="error-aside-image">
                  {{ article.username }}
                </div>
              </div>
            </el-image>
          </div>
          <div class="aside-post-title">
            <el-tooltip placement="top" effect="light" :content="article.articleTitle">
              <div>{{ article.articleTitle }}</div>
            </el-tooltip>
          </div>
        </div>
        <div class="aside-post-date">
          <i class="el-icon-date" style="color: var(--red)"></i>{{ article.createTime | formatter }}
        </div>
      </div>
    </div>
    <!-- 流浪者 -->
    <div class="card-widget wow" v-if="!$common.isEmpty(users) && $route.path !== '/article'" style="
        margin-top: 20px;
        position: relative;
        padding: 25px 10px 5px 17px;
        border-radius: 18px;
        animation: hideToShow 1s ease-in-out;
      ">
      <div style="font-weight: bold; margin-bottom: 20px">🎄流浪者</div>
      <div>
        <vue-seamless-scroll :data="users" style="height: 200px; overflow: hidden" :class-option="defaultOption">
          <div v-for="(item, i) in users" style="display: flex; justify-content: space-between" :key="i">
            <div style="display: flex">
              <el-avatar style="margin-bottom: 10px" :size="36" :src="item.avatar ? item.avatar : $store.state.webInfo.avatar"></el-avatar>
              <div style="
                  margin-left: 10px;
                  height: 36px;
                  line-height: 36px;
                  overflow: hidden;
                  max-width: 110px;
                ">
                {{ item.username }}
              </div>
            </div>
            <div style="height: 36px; line-height: 36px">
              {{ item.gender === 0 ? "保密" : item.gender === 1 ? "男" : "女" }}
            </div>
          </div>
        </vue-seamless-scroll>
      </div>
      <div class="admire-btn" @click="showTissue()">加入我们</div>
    </div>
    <!-- 分类 -->
    <div class="wow card-widget" style="
        margin-top: 20px;
        position: relative;
        padding: 25px 10px 5px 17px;
        border-radius: 18px;
        animation: hideToShow 1s ease-in-out;
      ">
      <div class="card-content2-title">
        <i class="el-icon-menu card-content2-icon"></i>
        <span>文章分类</span>
      </div>
      <div v-for="(sort, index) in sortInfo" :key="index" class="post-sort" @click="$router.push({ path: '/sort', query: { sortId: sort.id } })">
        <div>
          <span v-for="(s, i) in sort.sortName.split('')" :key="i">
            {{ s }}</span>
        </div>
      </div>
    </div>
    <!-- 书虫 -->
    <div v-if="!$common.isEmpty($store.getters.labelInfo)" style="
        position: relative;
        padding: 10px 10px 10px 15px;
        border-radius: 18px;
        margin-top: 20px;
        animation: hideToShow 1s ease-in-out;
      " class="wow card-widget">
      <div class="card-content2-title">
        <img class="card-content2-icon" style="vertical-align: -4px; margin-right: 2px" src="../assets/svg/insect.svg" />
        <span>书虫</span>
      </div>
      <div class="card-tag-cloud">
        <a class="item" v-for="(label, index) in $store.getters.labelInfo" :key="index" @click="$router.push({ path: '/tags', query: { labelId: label.id } })">
          {{ label.labelName }}
          <sup>{{ label.countOfLabel }}</sup>
        </a>
      </div>
    </div>
    <!-- 速览 -->
    <template v-if="$route.path !== '/article'">
      <div v-for="(sort, index) in fastSeeInfo" @click="selectSort(sort)" :key="index" class="wow card-widget" style="
        background-size: 100% 100%;
        position: relative;
        padding: 20px 40px 40px;
        animation: hideToShow 1s ease-in-out;
        margin-top: 20px;
        font-size: 18px;
        color: var(--fontColor);
      ">
        <div>
          <i class="el-icon-s-data card-content2-icon" style="margin-right: 5px; color: var(--red)"></i>速览
        </div>
        <div class="sort-name">
          {{ sort.sortName }}
        </div>
        <div style="
          font-weight: 400;
          font-size: 14px;
          margin-top: 20px;
          white-space: nowrap;
          text-overflow: ellipsis;
          overflow: hidden;
        ">
          {{ sort.sortDescription }}
        </div>
      </div>
    </template>
    <!-- 加入我们 -->
    <el-dialog title="加入组织" :visible.sync="showTissueDialog" width="25%" :append-to-body="true" destroy-on-close center>
      <div>
        <div class="admire-image"></div>
        <div>
          <div class="admire-content">1. 欢迎大家一起交流学习</div>
          <div class="admire-content">2. 申请通过后会加入博客交流群</div>
        </div>
      </div>
    </el-dialog>
  </div>
</template>
<script>
import vueSeamlessScroll from 'vue-seamless-scroll'
export default {
  components: {
    vueSeamlessScroll
  },
  data() {
    return {
      pagination: {
        current: 1,
        size: 4
      },
      latelyArticles: [],
      showTissueDialog: false,
      city: '',
      weather: {},
      tomWeather: {},
      week: '',
      total_sum: 0,
      userPagination: {
        current: 1,
        size: 9999,
        total: 0,
        searchKey: '',
        userStatus: null,
        userType: null
      },
      users: [],
      weekDay: ''
    }
  },
  computed: {
    webInfo() {
      return this.$store.state.webInfo
    },
    sortInfo() {
      return this.$store.state.sortInfo
    },
    fastSeeInfo() {
      return this.$store.getters.navigationBar
    },
    defaultOption() {
      return {
        step: 1, // 数值越大速度滚动越快
        limitMoveNum: this.users.length // 开始无缝滚动的数据量
      }
    }
  },
  mounted() {
    this.latelyArticles = this.$store.state.newArticles
    this.randomColor()
    this.postProvinceAndCity()
    this.getHistoryInfo()
    this.getUsers()
    if (!this.weather.week) {
      this.getWeek()
    }
  },
  filters: {
    formatter(row) {
      const day = row.split('.')[0].split('T')[0]
      const time = row.split('.')[0].split('T')[1]
      return `${day} 日 ${time}`
    }
  },
  methods: {
    getWeek() {
      const week = ['日', '一', '二', '三', '四', '五', '六']
      this.weekDay = new Date().getDay()
      this.weather.week = week[this.weekDay]
      this.tomWeather.week = week[this.weekDay + 1]
      if (this.weekDay == 6) {
        this.tomWeather.week = '日'
      }
    },
    selectSort(sort) {
      this.$emit('selectSort', sort)
    },
    showTip() {
      this.$router.push({ path: '/weiYan' })
    },
    randomColor() {
      function getRandomColor() {
        return `rgb(${Math.random() * 255}, ${Math.random() * 255}, ${Math.random() * 255})`
      }
      const itemEls = document.querySelectorAll('.item')
      for (const item of itemEls) {
        item.style.color = getRandomColor()
      }
    },
    async postProvinceAndCity() {
      const res = await this.$common.getIpAndCity(this)
      this.city = res.city
      this.weather = res.weather.casts[0]
      this.tomWeather = res.weather.casts[1]
      this.$http
        .post(this.$constant.baseURL + '/submit/', {
          province: res.weather.province,
          city: res.city,
          userId: this.$store.state.currentUser.id
        })
        .then(res => {})
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
    getHistoryInfo() {
      this.$http
        .get(this.$constant.baseURL + '/list/ip/')
        .then(res => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.total_sum = res.result[0].total_sum
            this.$store.commit('pageView', res.result[0])
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
    getUsers() {
      this.$http
        .post(this.$constant.baseURL + '/admin/user/list/', this.userPagination, true, false)
        .then(res => {
          if (!this.$common.isEmpty(res.result[0].data)) {
            this.users = res.result[0].data
          } else {
            this.users = []
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
    showTissue() {
      this.showTissueDialog = true
    }
  }
}
</script>
<style lang="scss" scoped>
.a-svg a {
  transition: all 0.5s;
  &:hover {
    transform: scale(1.3);
  }
}
.card-content1 {
  background: linear-gradient(-45deg, var(--wheat), var(--red6), var(--blue7), var(--red), var(--pink1));
  background-size: 400% 400%;
  animation: gradientBG 16s ease infinite;
  display: flex;
  flex-direction: column;
  align-items: center;
  border-radius: 18px;
  position: relative;
  overflow: hidden;
  :not(:first-child) {
    z-index: 10;
  }
}
.web-name {
  font-size: 30px;
  font-weight: bold;
  margin: 20px 0;
}
.web-info {
  margin: 15px 0;
  width: 80%;
  display: flex;
  flex-direction: row;
  justify-content: space-around;
}
.blog-info-box {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-around;
}
.blog-info-num {
  margin-top: 12px;
}
.collection-btn {
  position: relative;
  margin-top: 12px;
  background: var(--purple2);
  width: 65%;
  height: 35px;
  border-radius: 1rem;
  text-align: center;
  line-height: 35px;
  color: var(--bigRed);
  overflow: hidden;
  z-index: 1;
  margin-bottom: 10px;
  &::before {
    background: var(--gradualRed);
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    content: '';
    transform: scaleX(0);
    transform-origin: 0;
    transition: transform 0.5s ease-out;
    transition-timing-function: cubic-bezier(0.45, 1.64, 0.47, 0.66);
    border-radius: 1rem;
    z-index: -1;
  }
  &:hover::before {
    transform: scaleX(1);
  }
}
.card-content2-title {
  font-size: 18px;
  margin-bottom: 15px;
}
.card-content2-icon {
  color: var(--red);
  margin-right: 5px;
  animation: scale 1s ease-in-out infinite;
}
.aside-post-detail {
  display: flex;
}
.aside-post-image {
  width: 60%;
  border-radius: 0.2rem;
  margin-right: 8px;
  overflow: hidden;
  transition: all 1s;
  &:hover {
    transform: scale(1.1);
  }
}
.aside-post-title {
  height: 40px;
  line-height: 20px;
  width: 50%;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
.error-aside-image {
  background: var(--gradientAnimation);
  color: var(--white);
  padding: 10px;
  text-align: center;
  width: 100%;
  height: 100%;
}
.aside-post-date {
  margin-top: 8px;
  margin-bottom: 20px;
  color: var(--red);
  font-size: 12px;
}
.post-sort {
  border-radius: 1rem;
  margin-bottom: 15px;
  line-height: 30px;
  transition: all 0.3s;
  &:hover {
    background: var(--gradientAnimation);
    padding: 0px 15px;
    color: var(--white);
  }
}
.sort-name {
  font-weight: bold;
  font-size: 22px;
  margin-top: 30px;
  white-space: nowrap;
  text-overflow: ellipsis;
  overflow: hidden;
  &:after {
    top: 104px;
    width: 70px;
    left: 40px;
    height: 3px;
    background: var(--gradientAnimation);
    content: '';
    border-radius: 1px;
    position: absolute;
  }
}
.dot {
  position: absolute;
  height: 100%;
  width: 100%;
  background: var(--gradientAnimation);
}
.card-tag-cloud {
  width: inherit;
  display: flex;
  flex-wrap: wrap;
  justify-content: flex-start;
  a {
    height: 30px;
    line-height: 15px;
    padding: 4px;
    margin: 2px 2px 2px 0;
    border-radius: 3px;
    font-size: 18px;
    &:hover {
      color: var(--white) !important;
      background: var(--red);
      animation: hideToShow 0.3s ease-in-out;
    }
  }
  sup {
    font-size: 10px;
    opacity: 0.6;
  }
}
.my-link {
  margin: 8px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 195px;
}
.card-widget {
  clip-path: polygon(
    100% 0,
    100% 100%,
    0 100%,
    0 calc(100% - 52.5px),
    12.5px calc(100% - 40px),
    12.5px calc(100% - 50px),
    0 calc(100% - 62.5px),
    0 calc(100% - 82.5px),
    12.5px calc(100% - 70px),
    12.5px calc(100% - 80px),
    0 calc(100% - 92.5px),
    0 calc(100% - 112.5px),
    12.5px calc(100% - 100px),
    12.5px calc(100% - 112.5px),
    12.5px calc(100% - 110px),
    0 calc(100% - 122.5px),
    0 calc(100% - 142.5px),
    12.5px calc(100% - 130px),
    12.5px calc(100% - 141.5px),
    0 calc(100% - 154px),
    0 0
  );
  border-left: none;
  font-size: 105%;
  border-radius: 18px;
  border: 2px solid var(--myAsideBorderColor);
  background: var(--myAsideColor);
  transition: all 1s;
  &::before {
    content: '';
    width: 12.5px;
    background: linear-gradient(to top, transparent, var(--red));
    display: block;
    position: absolute;
    left: 0;
    height: 113px;
    bottom: 27px;
  }
  &:hover {
    transform: scale(0.95);
    background-color: var(--blue);
    transition-duration: 0.4s;
  }
}
.admire-box {
  background: var(--springBg) center center / cover no-repeat;
  padding: 25px;
  border-radius: 10px;
  animation: hideToShow 1s ease-in-out;
  margin-top: 40px;
}
.admire-btn {
  padding: 13px 15px;
  background: var(--maxLightRed);
  border-radius: 3rem;
  color: var(--white);
  width: 100px;
  user-select: none;
  text-align: center;
  margin: 20px auto 0;
  transition: all 1s;
  &:hover {
    transform: scale(1.2);
  }
}
.admire-image {
  margin: 0 auto 10px;
  border-radius: 10px;
  height: 150px;
  width: 150px;
  background: var(--admireImage) center center / cover no-repeat;
}
.admire-content {
  font-size: 14px;
  color: var(--bigRed);
  line-height: 1.5;
  margin: 5px;
}
::v-deep .el-dialog {
  background: var(--pink);
  border-radius: 14px;
  overflow: scroll;
  box-shadow: 0 14px 28px var(--mask), 0 10px 10px var(--miniMask);
  height: 320px;
  &::-webkit-scrollbar {
    width: 0px;
  }
}
@media screen and (max-width: 1100px) {
  .aside-post-image {
    width: 30%;
  }
}
</style>
