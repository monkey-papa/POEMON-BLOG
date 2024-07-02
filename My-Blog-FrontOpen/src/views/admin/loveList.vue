<template>
  <div>
    <div>
      <div class="handle-box">
        <el-select @blur="closeOptions" ref="status" clearable v-model="pagination.status" placeholder="状态" class="handle-select mrb10">
          <el-option key="1" label="启用" :value="true"></el-option>
          <el-option key="2" label="禁用" :value="false"></el-option>
        </el-select>
        <el-button type="primary" icon="el-icon-search" @click="search()">搜索</el-button>
      </div>
      <el-table :data="loves" border class="table" header-cell-class-name="table-header">
        <el-table-column prop="id" label="ID" width="55" align="center"></el-table-column>
        <el-table-column prop="userId" label="用户ID" align="center"></el-table-column>

        <el-table-column prop="manName" label="男生昵称" align="center"></el-table-column>
        <el-table-column prop="womanName" label="女生昵称" align="center"></el-table-column>

        <el-table-column label="背景封面" align="center">
          <template slot-scope="scope">
            <el-image lazy :preview-src-list="[scope.row.bgCover]" class="table-td-thumb" :src="scope.row.bgCover"
              fit="cover"></el-image>
          </template>
        </el-table-column>
        <el-table-column label="男生头像" align="center">
          <template slot-scope="scope">
            <el-image lazy :preview-src-list="[scope.row.manCover]" class="table-td-thumb" :src="scope.row.manCover"
              fit="cover"></el-image>
          </template>
        </el-table-column>
        <el-table-column label="女生头像" align="center">
          <template slot-scope="scope">
            <el-image lazy :preview-src-list="[scope.row.womanCover]" class="table-td-thumb" :src="scope.row.womanCover"
              fit="cover"></el-image>
          </template>
        </el-table-column>

        <el-table-column label="状态" align="center">
          <template slot-scope="scope">
            <el-tag :type="scope.row.status === false ? 'danger' : 'success'" disable-transitions>
              {{ scope.row.status === false ? '禁用' : '启用' }}
            </el-tag>
            <el-switch @click.native="changeStatus(scope.row)" v-model="scope.row.status"></el-switch>
          </template>
        </el-table-column>

        <el-table-column prop="timing" label="计时" align="center"></el-table-column>
        <el-table-column prop="countdownTitle" label="倒计时标题" align="center"></el-table-column>
        <el-table-column prop="countdownTime" label="倒计时时间" align="center"></el-table-column>
        <el-table-column prop="familyInfo" label="额外信息" align="center" show-overflow-tooltip></el-table-column>
        <el-table-column :formatter="$common.formatter" prop="createTime" label="创建时间" align="center"></el-table-column>
        <el-table-column label="操作" width="180" align="center">
          <template slot-scope="scope">
            <el-button type="text" icon="el-icon-delete" style="color: var(--orangeRed)" @click="handleDelete(scope.row)">
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination background layout="total, prev, pager, next" :current-page="pagination.current"
          :page-size="pagination.size" :total="pagination.total" @current-change="handlePageChange">
        </el-pagination>
      </div>
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
        status: null
      },
      loves: []
    }
  },
  created() {
    this.getLoves();
  },
  methods: {
    handleDelete(item) {
      this.$confirm('确认删除这对情侣吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        center: true
      }).then(() => {
        this.$http.get(this.$constant.baseURL + "/family/deleteFamily/", { id: item.id }, true)
          .then((res) => {
            this.pagination.current = 1;
            this.getLoves();
            this.$message({
              message: "删除成功！",
              type: "success"
            });
          })
          .catch((error) => {
            this.$message({
              message: error.message,
              type: "error"
            });
          });
      }).catch(() => {
        this.$message({
          type: 'success',
          message: '已取消删除!'
        });
      });
    },
    search() {
      this.pagination.total = 0;
      this.pagination.current = 1;
      this.getLoves();
    },
    getLoves() {
      this.$http.post(this.$constant.baseURL + "/family/listFamily/", this.pagination, true)
        .then((res) => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.loves = res.result[0].records;
            this.pagination.total = res.result[0].total;
          }
        })
        .catch((error) => {
          this.$message({
            message: error.message,
            type: "error"
          });
        });
    },
    changeStatus(item) {
      this.$http.get(this.$constant.baseURL + "/family/changeLoveStatus/", {
        id: item.id,
        flag: item.status
      }, true)
        .then((res) => {
          this.$message({
            message: "修改成功！",
            type: "success"
          });
        })
        .catch((error) => {
          this.$message({
            message: error.message,
            type: "error"
          });
        });
    },
    handlePageChange(val) {
      this.pagination.current = val;
      this.getLoves();
    },
    closeOptions() {
      this.$refs.status.blur()
    },
  }
}
</script>

<style scoped>
.handle-box {
  margin-bottom: 20px;
}

.handle-select {
  width: 200px;
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
</style>
