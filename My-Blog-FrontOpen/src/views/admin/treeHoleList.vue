<template>
  <div>
    <el-tag effect="dark" class="my-tag">
      <img style="vertical-align: -6px" src="../../assets/svg/message.svg" />
      留言列表
    </el-tag>
    <el-table :data="treeHoles" border class="table" header-cell-class-name="table-header">
      <el-table-column prop="id" label="ID" width="55" align="center"></el-table-column>
      <el-table-column prop="username" label="用户" align="center"></el-table-column>
      <el-table-column label="头像" align="center">
        <template slot-scope="scope">
          <el-image :preview-src-list="[scope.row.avatar]" lazy class="table-td-thumb" :src="scope.row.avatar || $store.state.webInfo.avatar" fit="cover"></el-image>
        </template>
      </el-table-column>
      <el-table-column prop="message" label="留言内容" align="center"></el-table-column>
      <el-table-column :formatter="$common.formatter" prop="createTime" label="创建时间" align="center"></el-table-column>
      <el-table-column label="操作" width="180" align="center">
        <template slot-scope="scope">
          <el-button type="text" icon="el-icon-circle-plus" style="color: green" @click="addProWords(scope.row)">
            添加违禁词
          </el-button>
          <el-button type="text" icon="el-icon-delete" style="color: var(--orangeRed)" @click="handleDelete(scope.row)">
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <div class="pagination">
      <el-pagination background layout="total, prev, pager, next" :current-page="pagination.current" :page-size="pagination.size" :total="pagination.total" @current-change="handlePageChange">
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
        total: 0
      },
      treeHoles: []
    }
  },
  created() {
    this.getTreeHoles()
  },
  methods: {
    getTreeHoles() {
      this.$http
        .post(this.$constant.baseURL + '/admin/treeHole/boss/list/', this.pagination, true)
        .then(res => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.treeHoles = res.result[0].records
            this.pagination.total = res.result[0].total
          }
        })
        .catch(error => {
          this.$message({
            message: error.message,
            type: 'error'
          })
        })
    },
    handlePageChange(val) {
      this.pagination.current = val
      this.getTreeHoles()
    },
    handleDelete(item) {
      this.$confirm('确定要删除这条留言吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        center: true
      })
        .then(() => {
          this.doDelete(item)
        })
        .catch(() => {
          this.$message({
            type: 'success',
            message: '已取消删除!'
          })
        })
    },
    doDelete(item) {
      this.$http
        .get(this.$constant.baseURL + '/webInfo/deleteTreeHole/', { id: item.id }, true)
        .then(res => {
          this.pagination.current = 1
          this.getTreeHoles()
          this.$message({
            message: '删除成功！',
            type: 'success'
          })
        })
        .catch(error => {
          this.$message({
            message: error.message,
            type: 'error'
          })
        })
    },
    addProWords(item) {
      this.$confirm('确定要添加为违禁词吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        center: true
      }).then(() => {
        this.$http
          .post(
            this.$constant.baseURL + '/prohibitedWords/add/',
            {
              message: item.message,
              username: item.username,
              avatar: item.avatar
            },
            true
          )
          .then(res => {
            if (res.failure !== 'exists null or error') {
              this.doDelete(item)
            } else {
              this.$message({
                message: '违禁词已存在！添加失败,请删除！',
                type: 'error'
              })
            }
          })
          .catch(error => {
            this.$message({
              message: '违禁词已存在！添加失败!',
              type: 'error'
            })
          })
      })
    }
  }
}
</script>

<style scoped>
.my-tag {
  margin-bottom: 40px;
  width: 100%;
  text-align: left;
  background: var(--lightBlue);
  border: none;
  height: 40px;
  line-height: 40px;
  font-size: 16px;
  color: black;
}

.pagination {
  margin: 20px 0;
  text-align: right;
}
.table-td-thumb {
  display: block;
  margin: auto;
  width: 50px;
  height: 50px;
}
</style>
