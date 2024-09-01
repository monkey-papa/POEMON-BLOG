<template>
  <div>
    <div class="handle-box">
      <el-input
        v-model="pagination.searchKey"
        placeholder="违禁词"
        class="handle-input mrb10"
        clearable
      ></el-input>
      <el-button
        type="primary"
        icon="el-icon-search"
        @click="searchForbiddenWord()"
        >搜索</el-button
      >
      <el-button
        v-hasPermi="['user:visit:read']"
        type="success"
        @click="openDialog(null, 0)"
        >添加</el-button
      >
      <el-button type="danger" @click="clearSearch()">清除参数</el-button>
    </div>
    <el-table
      :data="prohibitedWordsList"
      border
      class="table"
      header-cell-class-name="table-header"
    >
      <el-table-column prop="id" label="ID" align="center"></el-table-column>
      <el-table-column
        prop="username"
        label="用户"
        align="center"
      ></el-table-column>
      <el-table-column label="头像" width="100" align="center">
        <template slot-scope="scope">
          <el-image
            :preview-src-list="[scope.row.avatar]"
            lazy
            class="table-td-thumb"
            :src="scope.row.avatar"
            fit="cover"
          ></el-image>
        </template>
      </el-table-column>
      <el-table-column
        prop="message"
        label="违禁词"
        align="center"
      ></el-table-column>
      <el-table-column header-align="center" align="center" label="操作">
        <template slot-scope="scope">
          <el-button
            v-hasPermi="['user:visit:read']"
            type="text"
            icon="el-icon-edit"
            @click="openDialog(scope.row, 1)"
            >编辑</el-button
          >
          <el-button
            v-hasPermi="['user:visit:read']"
            type="text"
            icon="el-icon-delete"
            @click="deleteProWords(scope.row.id)"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>
    <el-dialog
      :before-close="handleClose"
      :title="type === 0 ? '添加违禁词' : '修改违禁词'"
      :visible.sync="dialogVisible"
      width="40%"
    >
      <input
        v-model="row.message"
        type="text"
        placeholder="请输入违禁词"
        style="width: 100%; height: 40px"
      />
      <div slot="footer" class="dialog-footer">
        <el-button @click="handleClose()">取 消</el-button>
        <el-button type="primary" @click="addOrUpdateProWords(type, row)"
          >确 定</el-button
        >
      </div>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "prohibitedWordsList",
  data() {
    return {
      row: {
        id: null,
        message: "",
        username: this.$store.state.currentUser.username,
        avatar: this.$store.state.currentUser.avatar,
      },
      type: 0,
      dialogVisible: false,
      pagination: {
        current: 1,
        size: 100,
        total: 0,
        searchKey: "",
      },
      prohibitedWordsList: [],
    };
  },
  created() {
    this.getProhibitedWordsList();
  },
  methods: {
    handleClose() {
      this.row = {
        id: null,
        message: "",
        username: this.$store.state.currentUser.username,
        avatar: this.$store.state.currentUser.avatar,
      };
      this.dialogVisible = false;
    },
    clearSearch() {
      this.pagination = {
        current: 1,
        size: 100,
        total: 0,
        searchKey: "",
      };
      this.getProhibitedWordsList();
    },
    openDialog(row, type) {
      this.dialogVisible = true;
      this.type = type;
      if (row !== null) {
        this.row = row;
      }
    },
    addOrUpdateProWords(type) {
      // 添加违禁词
      if (type === 0) {
        this.$http
          .post(
            this.$constant.baseURL + "/prohibitedWords/add/",
            this.row,
            true,
            true,
            true
          )
          .then((res) => {
            if (res.failure !== "exists null or error") {
              this.handleClose();
              this.row.message = "";
              this.$notify({
                title: "可以啦🍨",
                message: "添加成功！",
                type: "success",
                offset: 50,
                position: "top-left",
              });
            } else {
              this.$notify({
                type: "error",
                title: "可恶🤬",
                message: "违禁词已存在！添加失败,请删除！",
                position: "top-left",
                offset: 50,
              });
            }
            this.getProhibitedWordsList();
          });
      }
      // 编辑违禁词
      if (type === 1) {
        this.$http
          .post(
            this.$constant.baseURL + "/prohibitedWords/update/",
            this.row,
            true,
            true,
            true
          )
          .then((res) => {
            if (res.result === "success") {
              this.$notify({
                title: "可以啦🍨",
                message: "编辑成功！",
                type: "success",
                offset: 50,
                position: "top-left",
              });
            } else if (res.failure === "exists or null") {
              this.$notify({
                type: "error",
                title: "可恶🤬",
                message: "编辑失败！这个违禁词已存在",
                position: "top-left",
                offset: 50,
              });
            } else {
              this.$notify({
                type: "error",
                title: "可恶🤬",
                message: "编辑失败！",
                position: "top-left",
                offset: 50,
              });
            }
          });
        this.handleClose();
        this.getProhibitedWordsList();
      }
    },
    getProhibitedWordsList() {
      this.$http
        .post(
          this.$constant.baseURL + "/prohibitedWords/list/",
          this.pagination,
          true,
          true,
          true
        )
        .then((res) => {
          if (res.result[0]) {
            this.prohibitedWordsList = res.result[0].data;
            this.pagination.total = res.result[0].total;
          } else {
            this.$notify({
              type: "error",
              title: "可恶🤬",
              message: res.data.msg,
              position: "top-left",
              offset: 50,
            });
          }
        });
    },
    searchForbiddenWord() {
      this.getProhibitedWordsList();
    },
    deleteProWords(id) {
      this.$confirm("确定要删除这个违禁词吗？", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
        center: true,
      }).then(() => {
        this.$http
          .get(
            this.$constant.baseURL + "/prohibitedWords/delete/",
            { id: id },
            true,
            true
          )
          .then((res) => {
            if (res.result === "success") {
              this.$notify({
                title: "可以啦🍨",
                message: "删除成功！",
                type: "success",
                offset: 50,
                position: "top-left",
              });
              this.getProhibitedWordsList();
            } else {
              this.$notify({
                type: "error",
                title: "可恶🤬",
                message: res.data.msg,
                position: "top-left",
                offset: 50,
              });
            }
          });
      });
    },
  },
};
</script>

<style lang="scss" scoped>
.handle-box {
  margin-bottom: 20px;
}
.handle-select {
  width: 120px;
}
.handle-input {
  width: 160px;
  display: inline-block;
}
.table {
  width: 100%;
  font-size: 14px;
}
.mrb10 {
  margin-right: 10px;
  margin-bottom: 10px;
}
.table-td-thumb {
  display: block;
  margin: auto;
  width: 40px;
  height: 40px;
}
.pagination {
  margin: 20px 0;
  text-align: right;
}
.el-switch {
  margin: 5px;
}
.avatar {
  width: 40px;
  height: 40px;
  display: block;
}
.el-input-img {
  /*比其他输入框宽度少50px*/
  width: calc(100% - 50px);
  height: 40px;
  display: block;
}
</style>
