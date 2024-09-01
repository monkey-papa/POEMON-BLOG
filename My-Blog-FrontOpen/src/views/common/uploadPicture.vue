<template>
  <div>
    <el-upload
      class="upload-demo"
      ref="upload"
      multiple
      drag
      name="image"
      :action="this.$constant.qiniuUploadImages"
      :on-change="handleChange"
      :on-success="handleSuccess"
      :on-error="handleError"
      :list-type="listType"
      :accept="accept"
      :limit="maxNumber"
      :data="{
        userId:
          this.$store.state.currentAdmin.id || this.$store.state.currentUser.id,
      }"
      :auto-upload="false"
    >
      <div class="el-upload__text">
        <img style="margin-top: 10px" src="../../assets/svg/upload.svg" />
        <div>拖拽上传 / 点击上传</div>
      </div>
      <template v-if="listType === 'picture'">
        <div slot="tip" class="el-upload__tip">
          一次最多上传{{ maxNumber }}张图片，且每张图片不超过{{ maxSize }}M！
        </div>
      </template>
      <template v-else>
        <div slot="tip" class="el-upload__tip">
          一次最多上传{{ maxNumber }}个文件，且每个文件不超过{{ maxSize }}M！
        </div>
      </template>
    </el-upload>
    <div style="text-align: center; margin-top: 10px">
      <el-button type="success" style="font-size: 12px" @click="submitUpload">
        上传
      </el-button>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    isAdmin: {
      type: Boolean,
      default: false,
    },
    listType: {
      type: String,
      default: "picture",
    },
    //接受上传的文件类型
    accept: {
      type: String,
      default: "image/*",
    },
    maxSize: {
      type: Number,
      default: 5,
    },
    maxNumber: {
      type: Number,
      default: 5,
    },
    ResourceType: {
      type: String,
      default: "",
    },
  },
  methods: {
    //上传点击事件
    submitUpload() {
      this.$refs.upload.submit();
    },
    // 文件上传成功时的钩子
    handleSuccess(response, file) {
      let id =
        this.$store.state.currentUser.id || this.$store.state.currentAdmin.id;
      //存取资源接口
      this.$common.saveResource(
        this,
        file.raw.type,
        response.url,
        file.size,
        this.ResourceType,
        id,
        true
      );
      if (response.url) {
        this.$emit("addPicture", response.url);
      }
      this.$notify({
        title: "可以啦🍨",
        message: "上传成功！",
        type: "success",
        offset: 50,
        position: "top-left",
      });
    },
    //文件上传失败时的钩子
    handleError() {
      this.$notify({
        type: "warning",
        title: "淘气👻",
        message: "上传出错！",
        position: "top-left",
        offset: 50,
      });
    },
    // 添加文件、上传成功和上传失败时都会被调用
    handleChange(file, fileList) {
      let flag = false;
      if (file.size > this.maxSize * 1024 * 1024) {
        this.$notify({
          type: "warning",
          title: "淘气👻",
          message: "图片最大为" + this.maxSize + "M！",
          position: "top-left",
          offset: 50,
        });
        flag = true;
      }
      if (flag) {
        fileList.splice(fileList.size - 1, 1);
      }
    },
  },
};
</script>
