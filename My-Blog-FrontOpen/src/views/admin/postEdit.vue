<template>
  <div>
    <el-tag effect="dark" class="my-tag">
      <img style="vertical-align: -3px" src="../../assets/svg/document.svg" />
      文章信息
    </el-tag>
    <el-form
      :model="article"
      :rules="rules"
      ref="ruleForm"
      label-width="110px"
      class="demo-ruleForm"
    >
      <el-form-item label="文章标题" prop="articleTitle">
        <el-input maxlength="30" v-model="article.articleTitle"></el-input>
      </el-form-item>
      <div style="color: #ff4b2b; padding-left: 34px; margin-bottom: 5px">
        默认文章作者：Monkey-PaPa
      </div>
      <el-form-item label="文章作者" prop="articleAuthor">
        <el-input maxlength="30" v-model="article.articleAuthor"></el-input>
      </el-form-item>

      <el-form-item label="文章内容" prop="articleContent">
        <mavon-editor
          ref="md"
          @imgAdd="imgAdd"
          @imgDel="handleEditorImgDel"
          v-model="article.articleContent"
        >
        </mavon-editor>
      </el-form-item>

      <el-form-item label="是否启用评论" prop="commentStatus">
        <el-tag
          :type="article.commentStatus === false ? 'danger' : 'success'"
          disable-transitions
        >
          {{ article.commentStatus === false ? "否" : "是" }}
        </el-tag>
        <el-switch v-model="article.commentStatus"></el-switch>
      </el-form-item>

      <el-form-item label="是否推荐" prop="recommendStatus">
        <el-tag
          :type="article.recommendStatus === false ? 'danger' : 'success'"
          disable-transitions
        >
          {{ article.recommendStatus === false ? "否" : "是" }}
        </el-tag>
        <el-switch v-model="article.recommendStatus"></el-switch>
      </el-form-item>

      <el-form-item label="是否可见" prop="viewStatus">
        <el-tag
          :type="article.viewStatus === false ? 'danger' : 'success'"
          disable-transitions
        >
          {{ article.viewStatus === false ? "否" : "是" }}
        </el-tag>
        <el-switch v-model="article.viewStatus"></el-switch>
      </el-form-item>

      <el-form-item
        v-if="article.viewStatus === false"
        label="访问密码"
        prop="password"
      >
        <el-input maxlength="30" v-model="article.password"></el-input>
      </el-form-item>

      <el-form-item label="封面" prop="articleCover">
        <div style="display: flex">
          <el-input v-model="article.articleCover"></el-input>
          <el-image
            class="table-td-thumb"
            lazy
            style="margin-left: 10px"
            :preview-src-list="[article.articleCover]"
            :src="article.articleCover"
            fit="cover"
          ></el-image>
        </div>
        <uploadPicture
          :isAdmin="true"
          :ResourceType="'articleCover'"
          style="margin-top: 10px"
          @addPicture="addArticleCover"
          :maxSize="3"
          :maxNumber="1"
        ></uploadPicture>
      </el-form-item>
      <el-form-item label="分类" prop="sortId">
        <el-select v-model="article.sortId" placeholder="请选择分类">
          <el-option
            v-for="item in sorts"
            :key="item.id"
            :label="item.sortName"
            :value="item.id"
          >
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="标签" prop="labelId">
        <el-select v-model="article.labelId" placeholder="请选择标签">
          <el-option
            v-for="item in labelsTemp"
            :key="item.id"
            :label="item.labelName"
            :value="item.id"
          >
          </el-option>
        </el-select>
      </el-form-item>
    </el-form>
    <div class="myCenter" style="margin-bottom: 22px">
      <el-button type="primary" @click="submitForm('ruleForm')">保存</el-button>
      <el-button type="danger" @click="resetForm('ruleForm')"
        >重置所有修改</el-button
      >
    </div>
  </div>
</template>

<script>
const uploadPicture = () => import("../common/uploadPicture");

export default {
  components: {
    uploadPicture,
  },
  data() {
    return {
      id: this.$route.query.id,
      article: {
        articleTitle: "",
        articleAuthor: "Monkey-PaPa",
        articleContent: "",
        commentStatus: true,
        recommendStatus: false,
        viewStatus: true,
        password: "",
        articleCover: "",
        sortId: null,
        labelId: null,
        userId: null,
      },
      sorts: [],
      labels: [],
      labelsTemp: [],
      rules: {
        articleTitle: [
          { required: true, message: "请输入标题", trigger: "change" },
        ],
        articleAuthor: [
          { required: true, message: "请输入作者名称", trigger: "change" },
        ],
        articleContent: [
          { required: true, message: "请输入内容", trigger: "change" },
        ],
        commentStatus: [
          { required: true, message: "是否启用评论", trigger: "change" },
        ],
        recommendStatus: [
          { required: true, message: "是否推荐", trigger: "change" },
        ],
        viewStatus: [
          { required: true, message: "是否可见", trigger: "change" },
        ],
        articleCover: [{ required: true, message: "封面", trigger: "change" }],
        sortId: [{ required: true, message: "分类", trigger: "change" }],
        labelId: [{ required: true, message: "标签", trigger: "blur" }],
      },
      url: "",
    };
  },
  watch: {
    "article.sortId"(newVal, oldVal) {
      if (oldVal !== null) {
        this.article.labelId = null;
      }
      if (!this.$common.isEmpty(newVal) && !this.$common.isEmpty(this.labels)) {
        this.labelsTemp = this.labels.filter((l) => l.sortId === newVal);
      }
    },
  },
  created() {
    this.getSortAndLabel();
  },
  methods: {
    imgAdd(pos, file) {
      let fd = new FormData();
      fd.append("image", file);
      //上传md文档的图片到七牛云
      this.$http
        .uploadQiniu(this.$constant.qiniuUploadImages, fd)
        .then((res) => {
          if (!res.url) {
            this.$message({
              message: "上传出错！",
              type: "warning",
            });
            return;
          } else {
            this.url = res.url;
            this.$message({
              message: "上传图片成功！",
              type: "success",
            });
          }
          //将md中文件名为pos的这个图片路径替换为服务端返回后的链接
          this.$refs.md.$img2Url(pos, this.url);
          if (!this.$common.isEmpty(res.url)) {
            //将图片地址保存
            let url = res.url;
            let id = this.$store.state.currentAdmin.id;
            //存取资源接口
            this.$common.saveResource(
              this,
              "articlePicture",
              url,
              file.size,
              file.type,
              id,
              true
            );
          }
        })
        .catch((error) => {
          this.$message({
            message: error.message,
            type: "error",
          });
        });
    },
    handleEditorImgDel() {
      this.$http
        .post(
          this.$constant.baseURL + "/delArticleImage/",
          { url: this.url },
          true
        )
        .then((res) => {
          this.$message({
            message: "删除图片成功！",
            type: "success",
          });
        })
        .catch((error) => {
          this.$message({
            message: error.message,
            type: "error",
          });
        });
    },
    addArticleCover(res) {
      this.article.articleCover = res;
    },
    getSortAndLabel() {
      this.$http
        .get(this.$constant.baseURL + "/webInfo/listSortAndLabel/")
        .then((res) => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.sorts = res.result[0].data[0].sorts;
            this.labels = res.result[0].data[0].labels;
            if (!this.$common.isEmpty(this.id)) {
              this.editArticle();
            }
          }
        })
        .catch((error) => {
          this.$message({
            message: error.message,
            type: "error",
          });
        });
    },
    editArticle() {
      this.$http
        .get(
          this.$constant.baseURL + "/admin/article/getArticleById/",
          { id: parseInt(this.id) },
          true
        )
        .then((res) => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.article = res.result[0].data[0];
          }
        })
        .catch((error) => {
          this.$message({
            message: error.message,
            type: "error",
          });
        });
    },
    submitForm(formName) {
      if (
        this.article.viewStatus === false &&
        this.$common.isEmpty(this.article.password)
      ) {
        this.$message({
          message: "文章不可见时必须输入密码！",
          type: "error",
        });
        return;
      }
      this.$refs[formName].validate((valid) => {
        if (valid) {
          if (this.$common.isEmpty(this.id)) {
            this.saveArticle(this.article, "/article/saveArticle/");
          } else {
            this.article.id = this.id;
            this.saveArticle(this.article, "/article/updateArticle/");
          }
        } else {
          this.$message({
            message: "请完善必填项！",
            type: "error",
          });
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
      if (!this.$common.isEmpty(this.id)) {
        this.editArticle();
      }
    },
    saveArticle(value, url) {
      this.article.userId = this.$store.state.currentAdmin.id;
      this.$confirm("确认保存？", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
        center: true,
      })
        .then(() => {
          this.$http
            .post(this.$constant.baseURL + url, value, true)
            .then((res) => {
              this.$message({
                message: "保存成功！",
                type: "success",
              });
              this.$router.push({ path: "/postList" });
            })
            .catch((error) => {
              this.$message({
                message: error.message,
                type: "error",
              });
            });
        })
        .catch(() => {
          this.$message({
            type: "success",
            message: "已取消保存!",
          });
        });
    },
  },
};
</script>

<style scoped>
.my-tag {
  margin-bottom: 20px;
  width: 100%;
  text-align: left;
  background: var(--lightBlue);
  border: none;
  height: 40px;
  line-height: 40px;
  font-size: 16px;
  color: black;
}

.table-td-thumb {
  border-radius: 2px;
  width: 40px;
  height: 40px;
}

.el-switch {
  margin-left: 10px;
}

.el-form-item {
  margin-bottom: 40px;
}

.markdown-body {
  height: 400px;
}
</style>
