<template>
  <div>
    <div style="margin-bottom: 20px">
      <el-select
        @blur="closeOptions"
        ref="fuzzyCommentType"
        v-model="pagination.commentType"
        placeholder="评论来源类型"
        style="margin-right: 10px"
      >
        <el-option key="1" label="文章评论" value="article"></el-option>
        <el-option key="2" label="树洞评论" value="message"></el-option>
        <el-option key="3" label="恋爱留言" value="love"></el-option>
      </el-select>
      <el-input
        class="my-input"
        type="number"
        style="width: 140px; margin-right: 10px"
        v-model="pagination.source"
        placeholder="评论来源标识"
      ></el-input>
      <el-button type="primary" icon="el-icon-search" @click="searchComments()"
        >搜索</el-button
      >
      <el-button type="danger" @click="clearSearch()">清除参数</el-button>
    </div>
    <el-table
      :data="comments"
      border
      class="table"
      header-cell-class-name="table-header"
    >
      <el-table-column
        prop="id"
        label="ID"
        width="55"
        align="center"
      ></el-table-column>
      <el-table-column
        prop="username"
        label="用户名称"
        align="center"
      ></el-table-column>
      <el-table-column label="头像" align="center">
        <template slot-scope="scope">
          <el-image
            lazy
            class="table-td-thumb"
            :src="scope.row.avatar || $store.state.webInfo.avatar"
            fit="cover"
          ></el-image>
        </template>
      </el-table-column>
      <el-table-column
        prop="source"
        label="评论来源标识"
        align="center"
      ></el-table-column>
      <el-table-column
        prop="type"
        label="评论来源类型"
        align="center"
      ></el-table-column>
      <el-table-column
        prop="likeCount"
        label="点赞数"
        align="center"
      ></el-table-column>
      <el-table-column
        prop="commentContent"
        label="评论内容"
        align="center"
      ></el-table-column>
      <el-table-column
        prop="commentInfo"
        label="评论额外信息"
        align="center"
      ></el-table-column>
      <el-table-column
        :formatter="$common.formatter"
        prop="createTime"
        label="创建时间"
        align="center"
      ></el-table-column>
      <el-table-column label="操作" width="180" align="center">
        <template slot-scope="scope">
          <el-button
            type="text"
            icon="el-icon-circle-plus"
            style="color: green"
            @click="addProWords(scope.row)"
          >
            添加违禁词
          </el-button>
          <el-button
            type="text"
            icon="el-icon-delete"
            style="color: var(--orangeRed)"
            @click="handleDelete(scope.row)"
          >
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <div class="pagination">
      <el-pagination
        background
        layout="total, prev, pager, next"
        :current-page="pagination.current"
        :page-size="pagination.size"
        :total="pagination.total"
        @current-change="handlePageChange"
      >
      </el-pagination>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      pagination: {
        current: 1,
        size: 10,
        total: 0,
        source: null,
        commentType: "",
      },
      comments: [],
    };
  },
  created() {
    this.getComments();
  },
  methods: {
    clearSearch() {
      this.pagination = {
        current: 1,
        size: 10,
        total: 0,
        source: null,
        commentType: "",
      };
      this.getComments();
    },
    getComments() {
      this.$http
        .post(
          this.$constant.baseURL + "/admin/comment/boss/list/",
          this.pagination,
          true
        )
        .then((res) => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.comments = res.result[0].data;
            this.pagination.total = res.result[0].total;
          }
        })
        .catch((error) => {
          this.$message({
            message: error.message,
            type: "error",
          });
        });
    },
    handlePageChange(val) {
      if (
        this.pagination.commentType ||
        String(this.pagination.source) != "null"
      ) {
        this.$message({
          type: "error",
          message: "请清除查询参数再进行操作!",
        });
        return;
      }
      this.pagination.current = val;
      this.getComments();
    },
    searchComments() {
      this.pagination.total = 0;
      this.pagination.current = 1;
      this.getComments();
    },
    handleDelete(item) {
      this.$confirm(
        "删除评论后，所有该评论的回复均不可见。确认删除？",
        "提示",
        {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning",
          center: true,
        }
      )
        .then(() => {
          this.doDelete(item);
        })
        .catch(() => {
          this.$message({
            type: "success",
            message: "已取消删除!",
          });
        });
    },
    doDelete(item) {
      this.$http
        .get(
          this.$constant.baseURL + "/admin/comment/boss/deleteComment/",
          { id: item.id },
          true
        )
        .then((res) => {
          this.pagination.current = 1;
          this.getComments();
          this.$message({
            message: "删除成功！",
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
    addProWords(item) {
      this.$confirm("确定要添加为违禁词吗？", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
        center: true,
      }).then(() => {
        this.$http
          .post(
            this.$constant.baseURL + "/prohibitedWords/add/",
            {
              message: item.commentContent,
              username: item.username,
              avatar: item.avatar,
            },
            true
          )
          .then((res) => {
            if (res.failure !== "exists null or error") {
              this.doDelete(item);
            } else {
              this.$message({
                message: "违禁词已存在！添加失败,请删除！",
                type: "error",
              });
            }
          })
          .catch((error) => {
            this.$message({
              message: "违禁词已存在！添加失败!",
              type: "error",
            });
          });
      });
    },
    closeOptions() {
      this.$refs.fuzzyCommentType.blur();
    },
  },
};
</script>

<style scoped>
.pagination {
  margin: 20px 0;
  text-align: right;
}

.my-input >>> input::-webkit-inner-spin-button {
  appearance: none;
}
</style>
