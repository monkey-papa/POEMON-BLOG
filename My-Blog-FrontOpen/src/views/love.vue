<template>
  <div>
    <!-- 顶部封面 -->
    <div class="bg-wrap my-animation-slide-top">
      <!-- 背景图片 -->
      <el-image class="love-image my-el-image" lazy :src="love.bgCover" fit="cover">
        <div slot="error" class="image-slot"></div>
      </el-image>
      <!-- 对象 -->
      <div class="love-wrap transformCenter">
        <div>
          <el-avatar class="love-avatar" :src="love.manCover"></el-avatar>
          <div class="love-title">
            {{ love.manName }}
          </div>
        </div>
        <div>
          <img class="love-img" :src="$store.state.webInfo.randomCover[7]" alt="心心" />
        </div>
        <div>
          <el-avatar class="love-avatar" :src="love.womanCover"></el-avatar>
          <div class="love-title">
            {{ love.womanName }}
          </div>
        </div>
      </div>
    </div>
    <!-- 内容 -->
    <div class="love-container">
      <div class="myCenter love-content">
        <!-- 时间 -->
        <div>
          <!-- 计时 -->
          <div>
            <div class="love-time-title1">这是我们一起走过的</div>
            <div class="love-time1">
              第
              <span class="love-time1-item">{{ timing.year }}</span>
              年
              <span class="love-time1-item">{{ timing.month }}</span>
              月
              <span class="love-time1-item">{{ timing.day }}</span>
              日
              <span class="love-time1-item">{{ timing.hour }}</span>
              时
              <span class="love-time1-item">{{ timing.minute }}</span>
              分
              <span class="love-time1-item">{{ timing.second }}</span>
              秒
            </div>
          </div>
          <!-- 倒计时 -->
          <div class="love-time-title2" v-if="
              !$common.isEmpty(love.countdownTitle) ||
              !$common.isEmpty(love.countdownTime)
            ">
            下个{{ love.countdownTitle }}: 还有{{ countdownChange }}
          </div>
        </div>
      </div>
      <div style="padding: 0 20px">
        <div class="family-button shadow-box-mini" @click="changeCard(4)">
          <span class="family-button-title">{{
            card === 4 ? "回到主人家" : "开往表白墙"
          }}</span>
          <span class="family-button-car">
            <img src="../assets/svg/car.svg" />
          </span>
        </div>
      </div>
      <div style="padding: 0 20px">
        <!-- 卡片 -->
        <div class="card-wrap" v-show="card !== 4">
          <div :class="{ Active: card === 2 }" class="card-content shadow-box-mini" @click="changeCard(2)">
            <div>
              <el-avatar :size="100" :src="$store.state.webInfo.randomCover[6]"> </el-avatar>
            </div>
            <div class="card-right">
              <div class="card-title">时光相册</div>
              <div class="card-desc">📸记录美好瞬间</div>
            </div>
          </div>
          <div :class="{ Active: card === 1 }" class="card-content shadow-box-mini" @click="changeCard(1)">
            <div>
              <el-avatar :size="100" :src="$store.state.webInfo.randomCover[4]"> </el-avatar>
            </div>
            <div class="card-right">
              <div class="card-title">点点滴滴</div>
              <div class="card-desc">☀️今朝有酒今朝醉</div>
            </div>
          </div>
          <div :class="{ Active: card === 3 }" class="card-content shadow-box-mini" @click="changeCard(3)">
            <div>
              <el-avatar :size="100" :src="$store.state.webInfo.randomCover[5]"> </el-avatar>
            </div>
            <div class="card-right">
              <div class="card-title">祝福板</div>
              <div class="card-desc">📋写下对我们的祝福</div>
            </div>
          </div>
        </div>
        <div class="card-container">
          <div v-show="card === 1 && !$common.isEmpty(treeHoleList)">
            <treeHole :treeHoleList="treeHoleList" :avatar="$store.state.webInfo.avatar" @launch="launch" @deleteTreeHole="deleteTreeHole">
            </treeHole>
          </div>
          <div v-show="card === 2 && !$common.isEmpty(photoTitleList)">
            <!-- 标签 -->
            <div class="photo-title-warp" v-if="!$common.isEmpty(photoTitleList)">
              <div v-for="(item, index) in photoTitleList" :key="index" :class="{
                  isActive: photoPagination.classify === item.classify,
                }" @click="changePhotoTitle(item.classify)">
                <proTag :info="item.classify + ' ' + item.count" :color="
                    $constant.before_color_list[Math.floor(Math.random() * 6)]
                  " style="margin: 12px">
                </proTag>
              </div>
              <div class="article-info-news" @click="addPhoto" v-if="
          !$common.isEmpty($store.state.currentUser) &&
          ($store.state.currentUser.userType === 0 || $store.state.currentUser.userType === 1)
        ">
                <img src="../assets/svg/plus.svg" />
              </div>
            </div>
            <div class="photo-title">
              {{ photoPagination.classify }}
            </div>
            <photo :resourcePathList="photoList"></photo>
            <div class="pagination-wrap">
              <div @click="pagePhotos()" class="pagination" v-if="photoPagination.total > photoList.length">
                下一页
              </div>
              <div v-else style="user-select: none; color: var(--red)">
                ~~到底啦~~
              </div>
            </div>
          </div>
          <div v-show="card === 3" class="comment-content">
            <comment :source="1" :type="'love'"></comment>
          </div>
        </div>
        <div v-show="card === 4">
          <div class="family-container" v-show="!$common.isEmpty(randomFamily)">
            <div v-for="(item, index) in randomFamily" :key="index" @click="changeFamily(item)">
              <div class="family-wrap" :style="{
                  background:
                    'url(' + item.bgCover + ') center center / cover no-repeat',
                }">
                <div style="width: 90px">
                  <el-avatar class="family-avatar" :src="item.manCover"></el-avatar>
                  <div class="family-title">
                    {{ item.manName }}
                  </div>
                </div>
                <div>
                  <img class="family-img" :src="$store.state.webInfo.randomCover[7]" alt="心心" />
                </div>
                <div style="width: 90px">
                  <el-avatar class="family-avatar" :src="item.womanCover"></el-avatar>
                  <div class="family-title">
                    {{ item.womanName }}
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="family-bottom-wrap">
            <div v-show="!$common.isEmpty(randomFamily) && randomFamily.length > 5" class="family-bottom" style="background-color: var(--maxLightRed)" @click="getRandomFamily()">
              <span style="line-height: 50px"> 换一换 </span>
              <span style="vertical-align: middle">
                <img src="../assets/svg/change.svg" />
              </span>
            </div>
            <div class="family-bottom" style="background-color: var(--blue)" @click="loveDialog()">
              <span style="line-height: 50px"> 申请入住 </span>
              <span style="vertical-align: middle">
                <img src="../assets/svg/loves.svg" />
              </span>
            </div>
          </div>
        </div>
      </div>
      <!-- 入住 -->
      <el-dialog title="入住表白墙" :visible.sync="loveDialogVisible" width="50%" :append-to-body="true" destroy-on-close center>
        <div>
          <div class="form-main">
            <img :src="$store.state.webInfo.randomCover[11]" style="width: 100%" />
            <div>
              <div>
                <div class="myCenter form-friend">
                  <div class="user-content">
                    <div>
                      <div class="form-title">背景封面&nbsp;</div>
                      <div style="display: flex">
                        <el-input maxlength="120" v-model="userLove.bgCover"></el-input>
                        <div style="margin: 3px 0 0 10px">
                          <proButton :info="'上传背景'" @click.native="openPicture('bgCover')" :before="$constant.before_color_1" :after="$constant.after_color_1">
                          </proButton>
                        </div>
                      </div>
                    </div>
                    <div>
                      <div class="form-title">男生头像&nbsp;</div>
                      <div style="display: flex">
                        <el-input maxlength="120" v-model="userLove.manCover"></el-input>
                        <div style="margin: 3px 0 0 10px">
                          <proButton :info="'上传头像'" @click.native="openPicture('manCover')" :before="$constant.before_color_1" :after="$constant.after_color_1">
                          </proButton>
                        </div>
                      </div>
                    </div>
                    <div>
                      <div class="form-title">女生头像&nbsp;</div>
                      <div style="display: flex">
                        <el-input maxlength="120" v-model="userLove.womanCover"></el-input>
                        <div style="margin: 3px 0 0 10px">
                          <proButton :info="'上传头像'" @click.native="openPicture('womanCover')" :before="$constant.before_color_1" :after="$constant.after_color_1">
                          </proButton>
                        </div>
                      </div>
                    </div>
                    <div>
                      <div class="form-title">男生昵称&nbsp;</div>
                      <div>
                        <el-input maxlength="10" v-model="userLove.manName"></el-input>
                      </div>
                    </div>
                    <div>
                      <div class="form-title">女生昵称&nbsp;</div>
                      <div>
                        <el-input maxlength="10" v-model="userLove.womanName"></el-input>
                      </div>
                    </div>
                    <div>
                      <div class="form-title">计时时间&nbsp;</div>
                      <div>
                        <el-date-picker v-model="userLove.timing" value-format="yyyy-MM-dd HH:mm:ss" type="datetime" align="center" placeholder="选择计时时间">
                        </el-date-picker>
                      </div>
                    </div>
                    <div>
                      <div class="form-title">倒计时标题</div>
                      <div>
                        <el-input maxlength="20" v-model="userLove.countdownTitle"></el-input>
                      </div>
                    </div>
                    <div>
                      <div class="form-title">倒计时时间</div>
                      <div>
                        <el-date-picker v-model="userLove.countdownTime" value-format="yyyy-MM-dd HH:mm:ss" type="datetime" align="center" placeholder="选择倒计时时间">
                        </el-date-picker>
                      </div>
                    </div>
                    <div>
                      <div class="form-title">告白信&nbsp;&nbsp;</div>
                      <div>
                        <el-input type="textarea" show-word-limit maxlength="1000" v-model="userLove.familyInfo"></el-input>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="myCenter" style="margin-top: 20px">
                  <proButton :info="'提交'" @click.native.stop="submitLove()" :before="$constant.before_color_2" :after="$constant.after_color_2">
                  </proButton>
                </div>
              </div>
              <div>
                <img :src="$store.state.webInfo.randomCover[10]" style="width: 100%; margin: 5px auto" />
              </div>
              <p style="font-size: 12px; text-align: center; color: #999">
                欢迎入住表白墙
              </p>
            </div>
          </div>
        </div>
      </el-dialog>
      <el-dialog title="图片" :visible.sync="addPictureDialog" width="25%" :append-to-body="true" destroy-on-close center>
        <div>
          <uploadPicture :ResourceType="pictureType === 'lovePhoto' ? resourcePath.type:''" :isAdmin="pictureType === 'lovePhoto'" @addPicture="addPicture" :maxSize="10" :maxNumber="1">
          </uploadPicture>
        </div>
      </el-dialog>
      <el-dialog title="添加图片" :visible.sync="addResourcePathDialog" width="50%" :before-close="clearDialog" :append-to-body="true" center>
        <div>
          <div>
            <div style="margin-bottom: 5px">标题：</div>
            <el-input maxlength="60" v-model="resourcePath.title"></el-input>
            <div style="margin-top: 10px; margin-bottom: 5px">分类：</div>
            <el-select @blur="closeOptions" ref="fuzzyCommentType" v-model="resourcePath.classify" filterable allow-create default-first-option placeholder="请选择分类">
              <el-option v-for="item in photoTitleList" :key="item.classify" :label="item.classify" :value="item.classify">
              </el-option>
            </el-select>
            <div style="margin-top: 10px; margin-bottom: 5px">封面：</div>
            <div style="display: flex">
              <el-input v-model="resourcePath.cover"></el-input>
              <div style="width: 66px; margin: 3.5px 0 0 10px">
                <proButton :info="'上传封面'" @click.native="addResourcePathCover()" :before="$constant.before_color_1" :after="$constant.after_color_1">
                </proButton>
              </div>
            </div>
          </div>
          <div style="display: flex; margin-top: 30px" class="myCenter">
            <proButton :info="'提交'" @click.native="addResourcePath()" :before="$constant.before_color_2" :after="$constant.after_color_2">
            </proButton>
          </div>
        </div>
      </el-dialog>
      <div>
        <!-- 页脚 -->
        <myFooter></myFooter>
      </div>
    </div>
  </div>
</template>
<script>
const treeHole = () => import('./common/treeHole')
const comment = () => import('./common/comment')
const myFooter = () => import('./common/myFooter')
const photo = () => import('./common/photo')
const proTag = () => import('./common/proTag')
const proButton = () => import('./common/proButton')
const uploadPicture = () => import('./common/uploadPicture')
export default {
  components: {
    comment,
    photo,
    treeHole,
    myFooter,
    proTag,
    proButton,
    uploadPicture
  },
  data() {
    return {
      userLove: {
        bgCover: '',
        manCover: '',
        womanCover: '',
        manName: '',
        womanName: '',
        countdownTitle: '',
        countdownTime: '',
        timing: '',
        familyInfo: '',
        userId: this.$store.state.currentUser.id
      },
      loveDialogVisible: false,
      addPictureDialog: false,
      pictureType: '',
      adminLove: {},
      love: {
        bgCover: '',
        manCover: '',
        womanCover: '',
        manName: '',
        womanName: '',
        countdownTitle: '',
        countdownTime: '',
        timing: ''
      },
      weiYanPagination: {
        current: 1,
        size: 10,
        total: 0,
        userId: null
      },
      photoPagination: {
        current: 1,
        size: 10,
        total: 0,
        resourceType: 'lovePhoto',
        classify: ''
      },
      treeHoleList: [],
      photoTitleList: [],
      photoList: [],
      randomFamily: [],
      card: null,
      countdownChange: '',
      timing: {
        year: 0,
        month: 0,
        day: 0,
        hour: 0,
        minute: 0,
        second: 0
      },
      addResourcePathDialog: false,
      resourcePath: {
        title: '',
        classify: '',
        cover: '',
        type: 'lovePhoto' // 固定值
      }
    }
  },
  created() {
    this.getAdminFamily()
    this.card = 2
    this.getPhotoTitles()
  },
  methods: {
    closeOptions() {
      this.$refs.fuzzyCommentType.blur()
    },
    addResourcePath() {
      if (this.$common.isEmpty(this.resourcePath.title) || this.$common.isEmpty(this.resourcePath.classify) || this.$common.isEmpty(this.resourcePath.cover)) {
        this.$message({
          message: '标题、分类、封面不能为空！',
          type: 'error'
        })
        return
      }
      this.$http
        .post(this.$constant.baseURL + '/webInfo/saveResourcePath/', this.resourcePath, true)
        .then(res => {
          this.$message({
            message: '保存成功！',
            type: 'success'
          })
          this.clearDialog()
          this.getPhotoTitles('1')
        })
        .catch(error => {
          this.$message({
            message: error.message,
            type: 'error'
          })
        })
    },
    addResourcePathCover() {
      this.addPictureDialog = true
    },
    clearDialog() {
      this.resourcePath = {
        title: '',
        classify: '',
        cover: '',
        type: 'lovePhoto' // 固定值
      }
      this.pictureType = ''
      this.addResourcePathDialog = false
    },
    addPhoto() {
      this.pictureType = 'lovePhoto'
      this.addResourcePathDialog = true
    },
    openPicture(type) {
      this.pictureType = type
      this.addPictureDialog = true
    },
    addPicture(res) {
      if (this.pictureType === 'bgCover') {
        this.userLove.bgCover = res
      } else if (this.pictureType === 'manCover') {
        this.userLove.manCover = res
      } else if (this.pictureType === 'womanCover') {
        this.userLove.womanCover = res
      } else if (this.pictureType === 'lovePhoto') {
        this.resourcePath.cover = res
      }
      this.pictureType = ''
      this.addPictureDialog = false
    },
    submitLove() {
      if (this.userLove.bgCover.trim() === '') {
        this.$message({
          message: '你还没设置背景封面呢~',
          type: 'warning'
        })
        return
      }
      if (this.userLove.manCover.trim() === '') {
        this.$message({
          message: '你还没设置男生头像呢~',
          type: 'warning'
        })
        return
      }
      if (this.userLove.womanCover.trim() === '') {
        this.$message({
          message: '你还没设置女生头像呢~',
          type: 'warning'
        })
        return
      }
      if (this.userLove.manName.trim() === '') {
        this.$message({
          message: '你还没写男生昵称呢~',
          type: 'warning'
        })
        return
      }
      if (this.userLove.womanName.trim() === '') {
        this.$message({
          message: '你还没写女生昵称呢~',
          type: 'warning'
        })
        return
      }
      if (this.userLove.timing.trim() === '') {
        this.$message({
          message: '你还没设置计时时间呢~',
          type: 'warning'
        })
        return
      }
      this.$http
        .post(this.$constant.baseURL + '/family/addFamily/', this.userLove)
        .then(res => {
          this.$message({
            type: 'success',
            message: '提交成功，待管理员审核！'
          })
          this.userLove = {}
          this.loveDialogVisible = false
        })
        .catch(error => {
          this.$message({
            message: error.message,
            type: 'error'
          })
        })
    },
    loveDialog() {
      if (this.$common.isEmpty(this.$store.state.currentUser)) {
        this.$message({
          message: '请先登录！',
          type: 'error'
        })
        return
      }
      this.loveDialogVisible = true
    },
    changeFamily(family) {
      this.love = family
    },
    getPhotoTitles(val) {
      this.$http
        .get(this.$constant.baseURL + '/webInfo/getClassifyList/', {
          type: 'lovePhoto'
        })
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
            this.changePhoto(val)
          }
        })
        .catch(error => {
          this.$message({
            message: error.message,
            type: 'error'
          })
        })
    },
    getAdminFamily() {
      this.$http
        .get(this.$constant.baseURL + '/family/getAdminFamily/')
        .then(res => {
          if (!this.$common.isEmpty(res.result[0].data[0])) {
            this.love = res.result[0].data[0]
            this.adminLove = res.result[0].data[0]
            this.getLove()
            this.countdown()
            setInterval(() => {
              this.getLove()
              this.countdown()
            }, 1000)
          }
        })
        .catch(error => {
          this.$message({
            message: error.message,
            type: 'error'
          })
        })
    },
    getRandomFamily() {
      this.$http
        .get(this.$constant.baseURL + '/family/listRandomFamily/')
        .then(res => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.randomFamily = res.result[0].data
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
    changePhoto(val) {
      this.$http
        .post(this.$constant.baseURL + '/webInfo/clistResourcePath/', this.photoPagination)
        .then(res => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.photoList = val === '1' ? res.result[0].records : this.photoList.concat(res.result[0].records)
            this.photoPagination.total = res.result[0].total
          }
        })
        .catch(error => {
          this.$message({
            message: error.message,
            type: 'error'
          })
        })
    },
    changeCard(card) {
      if (card !== 4 || this.card !== card) {
        this.card = card
      } else {
        card = 1
        this.card = 1
        this.love = this.adminLove
      }
      if (card === 1) {
        if (this.$common.isEmpty(this.treeHoleList)) {
          this.getWeiYan()
        }
      } else if (card === 2) {
        if (this.$common.isEmpty(this.photoTitleList)) {
          this.getPhotoTitles()
        }
      } else if (card === 4) {
        if (this.$common.isEmpty(this.randomFamily)) {
          this.getRandomFamily()
        }
      }
    },
    getLove() {
      if (this.$common.isEmpty(this.love.timing)) {
        return
      }
      let diff = this.$common.timeDiff(this.love.timing)
      this.timing.year = diff.diffYear
      this.timing.month = diff.diffMonth
      this.timing.day = diff.diffDay
      this.timing.hour = diff.diffHour
      this.timing.minute = diff.diffMinute
      this.timing.second = diff.diffSecond
    },
    //倒计时
    countdown() {
      if (this.$common.isEmpty(this.love.countdownTime)) {
        return
      }
      let countdown = this.$common.countdown(this.love.countdownTime)
      this.countdownChange = countdown.d + '天' + countdown.h + '时' + countdown.m + '分' + countdown.s + '秒'
    },
    launch() {
      if (this.weiYanPagination.total !== this.treeHoleList.length) {
        this.weiYanPagination.current = this.weiYanPagination.current + 1
        this.getWeiYan()
      } else {
        this.$message({
          message: '~~到底啦~~',
          type: 'warning'
        })
      }
    },
    getWeiYan() {
      this.$http
        .post(this.$constant.baseURL + '/weiYan/listWeiYan/', this.weiYanPagination)
        .then(res => {
          if (!this.$common.isEmpty(res.result[0])) {
            res.result[0].records.forEach(c => {
              c.content = c.content.replace(/\n{2,}/g, '<div style="height: 12px"></div>')
              c.content = c.content.replace(/\n/g, '<br/>')
              c.content = this.$common.faceReg(c.content)
              c.content = this.$common.pictureReg(c.content)
            })
            this.treeHoleList = res.result[0].records
            this.weiYanPagination.total = res.result[0].total
          }
        })
        .catch(error => {
          this.$message({
            message: error.message,
            type: 'error'
          })
        })
    },
    deleteTreeHole(id) {
      if (this.$common.isEmpty(this.$store.state.currentUser)) {
        this.$message({
          message: '请先登录！',
          type: 'error'
        })
        return
      }
      this.$confirm('确认删除？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        center: true
      })
        .then(() => {
          this.$http
            .get(this.$constant.baseURL + '/weiYan/deleteWeiYan/', { id: id })
            .then(res => {
              this.$message({
                type: 'success',
                message: '删除成功!'
              })
              this.weiYanPagination.current = 1
              this.getWeiYan()
            })
            .catch(error => {
              this.$message({
                message: error.message,
                type: 'error'
              })
            })
        })
        .catch(() => {
          this.$message({
            type: 'success',
            message: '已取消删除!'
          })
        })
    }
  }
}
</script>
<style scoped>
.article-info-news {
  height: 30px;
  animation: scale 1s ease-in-out infinite;
}
.love-container {
  background-image: linear-gradient(to right, rgba(37, 82, 110, 0.1) 1px, var(--background) 1px), linear-gradient(to bottom, rgba(37, 82, 110, 0.1) 1px, var(--background) 1px);
  background-size: 3rem 3rem;
}
.bg-wrap {
  height: 55vh;
  position: relative;
  overflow: hidden;
}
.love-image::before {
  content: '';
  position: absolute;
  width: 100%;
  height: 100%;
  background-color: var(--miniMask);
}
.love-wrap {
  width: 90%;
  backdrop-filter: blur(10px);
  background: rgba(255, 255, 255, 0.1);
  max-width: 950px;
  border-radius: 3em;
  display: flex;
  align-items: center;
  justify-content: space-around;
  padding: 50px 70px 30px 70px;
}
.love-avatar {
  border: rgba(255, 255, 255, 0.2) 4px solid;
  width: 180px;
  height: 180px;
}
.love-title {
  margin-top: 15px;
  text-align: center;
  font-size: 25px;
  font-weight: 700;
  color: var(--red);
}
.love-img {
  animation: imgScale 1.5s linear infinite;
  width: 120px;
  height: 120px;
}
.love-content {
  max-width: 1200px;
  overflow: hidden;
  margin: 0 auto;
  user-select: none;
}
.love-time-title1 {
  font-size: 2rem;
  font-weight: 600;
  letter-spacing: 0.2rem;
  line-height: 4rem;
  text-align: center;
  background-image: linear-gradient(270deg, #ff4500, #ffa500, #ec695c, #67c23a, #74bdf0, #1e90ff, #9370db, #ff69b4, #ff4500);
  -webkit-background-clip: text;
  animation: jianBian 60s linear infinite;
  width: 3000px;
  color: rgba(0, 0, 0, 0);
}
.love-time-title2 {
  color: var(--bigRed);
  text-align: center;
  font-size: 1.5rem;
  line-height: 4rem;
  font-weight: 400;
  letter-spacing: 2px;
}
.love-time1 {
  color: var(--red);
  text-align: center;
  font-size: 2rem;
  font-weight: 700;
}
.love-time1-item {
  font-size: 4rem;
  font-weight: 700;
}
.card-wrap {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  justify-content: center;
  padding: 20px 0;
}
.card-content {
  display: flex;
  padding: 25px;
  margin: 25px auto;
  border-radius: 20px;
  max-width: 780px;
  transition: all 0.3s;
  background: var(--background);
}
.card-content:hover,
.family-button:hover,
.family-wrap:hover {
  transform: translateY(-6px);
}
.card-right {
  margin-left: 20px;
}
.card-title {
  font-size: 1.6rem;
  letter-spacing: 0.2rem;
  line-height: 3.5rem;
  font-weight: 700;
}
.card-desc {
  font-size: 1.1rem;
  letter-spacing: 0.2rem;
  color: #777777;
}
.card-container {
  max-width: 1500px;
  margin: 20px auto 40px;
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
.pagination:hover {
  border: 1px solid var(--gradientAnimation);
  color: var(--gradientAnimation);
  box-shadow: 0 0 5px var(--gradientAnimation);
}
.comment-content {
  max-width: 1000px;
  margin: 0 auto;
}
.photo-title {
  text-align: center;
  font-size: 30px;
  font-weight: 700;
  line-height: 80px;
  letter-spacing: 2px;
  color: var(--red);
}
.photo-title-warp {
  max-width: 1150px;
  margin: 50px auto;
  padding: 20px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  flex-wrap: wrap;
}
.isActive {
  animation: scale 2.5s ease-in-out infinite;
}
.Active {
  background-color: var(--pink);
  animation: scale 6s ease-in-out infinite;
}
.family-button {
  position: relative;
  overflow: hidden;
  height: 150px;
  margin: 50px auto 15px;
  border-radius: 20px;
  max-width: 350px;
  transition: all 0.3s;
  background: var(--love) center center / cover no-repeat;
  user-select: none;
}
.family-button::before {
  content: '';
  position: absolute;
  width: 100%;
  height: 100%;
  background-color: var(--miniMask);
}
.family-button-title {
  position: absolute;
  line-height: 150px;
  margin-left: 20px;
  font-size: 25px;
  font-weight: 700;
  color: var(--red);
}
.family-button-car {
  position: absolute;
  margin-left: 220px;
  margin-top: 55px;
  animation: passing 4s linear infinite;
}
.family-container {
  display: flex;
  justify-content: space-around;
  flex-wrap: wrap;
  margin-bottom: 40px;
}
.family-wrap {
  width: 350px;
  border-radius: 20px;
  display: flex;
  align-items: center;
  justify-content: space-around;
  padding: 15px 25px 5px 25px;
  margin: 20px;
  transition: all 0.3s;
  user-select: none;
}
.family-avatar {
  border: var(--pink) 4px solid;
  width: 90px;
  height: 90px;
}
.family-title {
  margin-top: 15px;
  text-align: center;
  font-size: 15px;
  font-weight: 400;
  color: var(--red);
  display: -webkit-box;
  /*弹性伸缩盒子*/
  -webkit-box-orient: vertical;
  /*子元素垂直排列*/
  white-space: wrap;
  /*溢出换行*/
  -webkit-line-clamp: 2;
  /*显示的行数*/
  overflow: hidden;
  /*溢出部分隐藏*/
  text-overflow: ellipsis;
  /*溢出部分用省略号显示*/
}
.family-img {
  animation: imgScale 2s linear infinite;
  width: 60px;
  height: 60px;
}
.family-bottom-wrap {
  display: flex;
  justify-content: space-around;
  margin: 0 0 40px;
}
.family-bottom {
  color: white;
  border-radius: 3rem;
  width: 150px;
  text-align: center;
  height: 50px;
  user-select: none;
}
.form-main {
  animation: hideToShow 2s;
  border-radius: 12px;
  overflow: hidden;
}
.user-content > div {
  height: 65px;
  display: flex;
  align-items: center;
}
.user-content >>> .el-input__inner {
  border: none;
  height: 40px;
  width: 250px;
  background: var(--whiteMask);
}
.user-content >>> .el-textarea__inner {
  border: none;
  width: 250px;
  background: var(--whiteMask);
}
.user-content >>> .el-input__count {
  background: var(--transparent);
  user-select: none;
}
.form-friend {
  background-color: #ebf1f6;
  padding: 20px 0;
}
.form-title {
  margin: 10px;
  text-align: center;
}
@media screen and (max-width: 1200px) {
  .user-content > div {
    display: unset;
    align-items: unset;
  }
}
@media screen and (max-width: 800px) {
  .love-wrap {
    border-radius: 1.5em;
    padding: 40px 30px 10px 30px;
  }
  .love-avatar {
    width: 120px;
    height: 120px;
  }
  .love-img {
    width: 100px;
    height: 100px;
  }
  .love-time1 {
    font-size: 1.4rem;
  }
  .love-time1-item {
    font-size: 3rem;
  }
}
@media screen and (max-width: 600px) {
  .love-title {
    font-size: 16px;
    white-space: wrap;
    display: -webkit-box;
    /*弹性伸缩盒子*/
    -webkit-box-orient: vertical;
    /*溢出换行*/
    -webkit-line-clamp: 1;
    /*显示的行数*/
    overflow: hidden;
    /*溢出部分隐藏*/
    text-overflow: ellipsis;
    /*溢出部分用省略号显示*/
  }
  .love-wrap {
    padding: 30px 20px 10px 20px;
  }
  .love-avatar {
    width: 80px;
    height: 80px;
  }
  .love-img {
    width: 70px;
    height: 70px;
  }
  .love-time1-item {
    font-size: 1rem;
  }
  .love-time-title2 {
    font-size: 0.8rem;
  }
  .user-content >>> .el-textarea__inner {
    width: 100%;
  }
  .user-content >>> .el-input__inner {
    width: 190px;
  }
  .card-container .tree-hole-container {
    padding: 0;
  }
}
@media screen and (max-width: 1150px) {
  .card-wrap {
    display: unset;
    justify-content: unset;
  }
  .photo-title-warp {
    max-width: 780px;
  }
  .family-button {
    max-width: 780px;
  }
}
::v-deep .el-dialog {
  border-radius: 18px;
}
::v-deep .form-title {
  color: var(--fontColor);
}
::v-deep .el-input__inner::placeholder {
  color: var(--fontColor);
}
::v-deep .el-input__icon.el-icon-time {
  color: var(--fontColor);
}
</style>
