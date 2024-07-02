<template>
  <div>
    <!-- 框 -->
    <textarea class="comment-textarea" v-model="commentContent" placeholder="善语结善缘，恶言伤人心..." maxlength="1000" />
    <!-- 按钮 -->
    <div class="myBetween" style="margin-bottom: 10px">
      <div style="display: flex">
        <div :class="{ 'emoji-active': showEmoji }" @click="showEmoji = !showEmoji">
          <i class="el-icon-orange myEmoji" style="color: var(--darkBlue);"></i>
        </div>
        <div @click="openPicture()">
          <i class="el-icon-picture myPicture" style="color: var(--darkBlue);"></i>
        </div>
      </div>

      <div style="display: flex">
        <!-- <proButton
          :info="'涂鸦'"
          v-show="!$common.mobile() && !disableGraffiti"
          @click.native="showGraffiti()"
          :before="$constant.before_color_1"
          :after="$constant.after_color_1"
          style="margin-right: 6px"
        >
        </proButton> -->
        <proButton :info="'发射'" @click.native="submitComment()" :before="$constant.before_color_2" :after="$constant.after_color_2">
        </proButton>
      </div>
    </div>
    <!-- 表情 -->
    <emoji @addEmoji="addEmoji" :showEmoji="showEmoji"></emoji>

    <el-dialog title="图片" :visible.sync="showPicture" width="25%" :append-to-body="true" destroy-on-close center custom-class="dialog">
      <div>
        <uploadPicture :ResourceType="'commentPicture'" @addPicture="addPicture" :maxSize="5" :maxNumber="1">
        </uploadPicture>
      </div>
    </el-dialog>
  </div>
</template>

<script>
const emoji = () => import('./emoji')
const proButton = () => import('./proButton')
const uploadPicture = () => import('./uploadPicture')

export default {
  components: {
    emoji,
    proButton,
    uploadPicture
  },
  props: {
    disableGraffiti: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      commentContent: '',
      showEmoji: false,
      showPicture: false,
      picture: {
        name: this.$store.state.currentUser.username,
        url: ''
      }
    }
  },
  methods: {
    openPicture() {
      if (this.$common.isEmpty(this.$store.state.currentUser)) {
        this.$message({
          message: '请先登录！',
          type: 'error'
        })
        return
      }
      this.showPicture = true
    },
    addPicture(res) {
      this.picture.url = res
      this.savePicture()
    },
    savePicture() {
      let img = '<' + this.picture.name + ',' + this.picture.url + '>'
      this.commentContent += img
      this.picture.url = ''
      this.showPicture = false
    },
    addEmoji(key) {
      this.commentContent += key
    },
    showGraffiti() {
      if (this.$common.isEmpty(this.$store.state.currentUser)) {
        this.$message({
          message: '请先登录！',
          type: 'error'
        })
        return
      }
      this.commentContent = ''
      this.$emit('showGraffiti')
    },
    submitComment() {
      if (this.$common.isEmpty(this.$store.state.currentUser)) {
        this.$message({
          message: '请先登录！',
          type: 'error'
        })
        return
      }
      if (this.commentContent.trim() === '') {
        this.$message({
          message: '你还没写呢~',
          type: 'warning'
        })
        return
      }
      this.$emit('submitComment', this.commentContent.trim())
      this.commentContent = ''
    }
  }
}
</script>

<style scoped>
.comment-textarea {
  border: 2px solid var(--lightGray);
  width: 100%;
  font-size: 14px;
  padding: 15px;
  min-height: 180px;
  /* 不改变大小 */
  resize: none;
  /* 不改变边框 */
  outline: none;
  border-radius: 4px;
  background-image: var(--comment1URL), var(--comment2URL);
  background-position: left, right;
  background-repeat: no-repeat, no-repeat;
  background-size: 25%, 30%;
  margin-bottom: 10px;
}

.comment-textarea:focus {
  border-color: orange;
}
.comment-textarea::placeholder {
  color: var(--darkBlue);
}
.myEmoji {
  font-size: 18px;
  transition: all 0.5s;
  margin-right: 12px;
}

.myEmoji:hover {
  transform: rotate(360deg);
  font-size: 22px;
}

.myPicture {
  font-size: 18px;
}

.emoji-active {
  color: var(--red);
}
::v-deep .dialog {
  border-radius: 14px;
  overflow: scroll;
  box-shadow: 0 14px 28px rgba(0, 0, 0, 0.25), 0 10px 10px rgba(0, 0, 0, 0.22);
  height: 300px;
}
::v-deep .dialog::-webkit-scrollbar {
  width: 0px;
}
</style>
