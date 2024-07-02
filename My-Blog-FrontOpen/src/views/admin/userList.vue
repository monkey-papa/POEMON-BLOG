<template>
  <div>
    <div>
      <div class="handle-box">
        <el-select @blur="closeOptions('userType')" ref="fuzzyUserType" v-model="pagination.userType" placeholder="用户类型" class="handle-select mrb10">
          <el-option key="1" label="Boss" :value="0"></el-option>
          <el-option key="2" label="管理员" :value="1"></el-option>
          <el-option key="3" label="普通用户" :value="2"></el-option>
        </el-select>
        <el-select @blur="closeOptions('userStatus')" ref="fuzzyUserStatus" v-model="pagination.userStatus" placeholder="用户状态" class="handle-select mrb10">
          <el-option key="1" label="启用" :value="true"></el-option>
          <el-option key="2" label="禁用" :value="false"></el-option>
        </el-select>
        <el-input v-model="pagination.searchKey" placeholder="完整用户名/完整手机号" class="handle-input mrb10"></el-input>
        <el-button type="primary" icon="el-icon-search" @click="searchUser()">搜索</el-button>
        <el-button type="danger" @click="clearSearch()">清除参数</el-button>
      </div>
      <div style="color: var(--red)">
        警告：Boss用户只能有1个，其他用户可以任意个
      </div>
      <el-table :data="users" border class="table" header-cell-class-name="table-header">
        <el-table-column prop="id" label="ID" width="55" align="center"></el-table-column>
        <el-table-column prop="username" label="用户名" align="center"></el-table-column>
        <el-table-column prop="phoneNumber" label="手机号" align="center"></el-table-column>
        <el-table-column prop="email" label="邮箱" align="center"></el-table-column>
        <el-table-column label="赞赏" width="100" align="center">
          <template slot-scope="scope">
            <el-input size="medium" maxlength="30" v-model="scope.row.admire" @blur="changeUserAdmire(scope.row)"></el-input>
          </template>
        </el-table-column>
        <el-table-column label="用户状态" align="center">
          <template slot-scope="scope">
            <el-tag :type="scope.row.userStatus === false ? 'danger' : 'success'" disable-transitions>
              {{ scope.row.userStatus === false ? "禁用" : "启用" }}
            </el-tag>
            <el-switch @click.native="changeUserStatus(scope.row)" v-model="scope.row.userStatus"></el-switch>
          </template>
        </el-table-column>
        <el-table-column label="头像" align="center">
          <template slot-scope="scope">
            <el-image :preview-src-list="[scope.row.avatar]" lazy class="table-td-thumb" :src="scope.row.avatar || $store.state.webInfo.avatar" fit="cover"></el-image>
          </template>
        </el-table-column>
        <el-table-column label="性别" align="center">
          <template slot-scope="scope">
            <el-tag type="success" v-if="scope.row.gender === 1" :class="{ boy: scope.row.gender === 1 }" disable-transitions>
              男生
            </el-tag>
            <el-tag type="success" v-else-if="scope.row.gender === 2" :class="{ girl: scope.row.gender === 2 }" disable-transitions>
              女生
            </el-tag>
            <el-tag type="success" v-else disable-transitions> 保密 </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="introduction" label="简介" align="center"></el-table-column>
        <el-table-column label="用户类型" width="100" align="center">
          <template slot-scope="scope">
            <el-tag type="success" v-if="scope.row.userType === 0" :class="{ boss: scope.row.userType === 0 }" @click.native="editUser(scope.row)" disable-transitions>
              Boss
            </el-tag>
            <el-tag type="success" v-else-if="scope.row.userType === 1" :class="{ manager: scope.row.userType === 1 }" @click.native="editUser(scope.row)" disable-transitions>
              管理员
            </el-tag>
            <el-tag type="success" v-else @click.native="editUser(scope.row)" disable-transitions>
              普通用户
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column :formatter="$common.formatter" prop="createTime" label="注册时间" align="center"></el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination background layout="total, prev, pager, next" :current-page="pagination.current" :page-size="pagination.size" :total="pagination.total" @current-change="handlePageChange">
        </el-pagination>
      </div>
    </div>

    <!-- 编辑弹出框 -->
    <el-dialog title="修改用户类型" :visible.sync="editVisible" width="30%" :before-close="handleClose" :append-to-body="true" destroy-on-close center>
      <div class="myCenter">
        <el-radio-group v-model="changeUser.userType">
          <el-radio-button :label="0">Boss</el-radio-button>
          <el-radio-button :label="1">管理员</el-radio-button>
          <el-radio-button :label="2">普通用户</el-radio-button>
        </el-radio-group>
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="handleClose()">取 消</el-button>
        <el-button type="primary" @click="saveEdit()">确 定</el-button>
      </span>
    </el-dialog>
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
        searchKey: '',
        userStatus: null,
        userType: null
      },
      users: [],
      changeUser: {
        id: null,
        userType: null
      },
      editVisible: false
    }
  },
  created() {
    this.getUsers()
  },
  methods: {
    clearSearch() {
      this.pagination = {
        current: 1,
        size: 10,
        total: 0,
        searchKey: '',
        userStatus: null,
        userType: null
      }
      this.getUsers()
    },
    getUsers() {
      this.$http
        .post(this.$constant.baseURL + '/admin/user/list/', this.pagination, true, false)
        .then(res => {
          if (!this.$common.isEmpty(res.result[0].data)) {
            this.users = res.result[0].data
            this.pagination.total = res.result[0].total
          } else {
            this.users = []
          }
        })
        .catch(error => {
          this.$message({
            message: error.message,
            type: 'error'
          })
        })
    },
    changeUserStatus(user) {
      this.$http
        .post(
          this.$constant.baseURL + '/admin/user/changeUserStatus/',
          {
            userId: user.id,
            flag: user.userStatus
          },
          true,
          false
        )
        .then(res => {
          this.$message({
            message: '修改成功！',
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
    changeUserAdmire(user) {
      if (!this.$common.isEmpty(user.admire)) {
        this.$confirm('确认保存？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'success',
          center: true
        })
          .then(() => {
            this.$http
              .post(
                this.$constant.baseURL + '/admin/user/changeUserAdmire/',
                {
                  userId: user.id,
                  admire: user.admire
                },
                true,
                false
              )
              .then(res => {
                this.$message({
                  message: '修改成功！',
                  type: 'success'
                })
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
              message: '已取消保存!'
            })
          })
      }
    },
    editUser(user) {
      this.changeUser.id = user.id
      this.changeUser.userType = user.userType
      if (this.changeUser.userType === 0) {
        this.$message({
          message: '禁止修改Boss用户',
          type: 'error'
        })
        return
      }
      this.editVisible = true
    },
    handlePageChange(val) {
      if (this.pagination.userType == 0 || this.pagination.userType == 1 || this.pagination.userType == 2 || this.pagination.userStatus == true || this.pagination.userStatus == false || this.pagination.searchKey) {
        this.$message({
          type: 'error',
          message: '请清除查询参数再进行操作!'
        })
        return
      }
      this.pagination.current = val
      this.getUsers()
    },
    searchUser() {
      this.pagination.total = 0
      this.pagination.current = 1
      this.getUsers()
    },
    handleClose() {
      this.changeUser = {
        id: null,
        userType: null
      }
      this.editVisible = false
    },
    saveEdit() {
      this.$http
        .post(
          this.$constant.baseURL + '/admin/user/changeUserType/',
          {
            userId: this.changeUser.id,
            userType: this.changeUser.userType
          },
          true,
          false
        )
        .then(res => {
          this.handleClose()
          this.getUsers()
          this.$message({
            message: '修改成功！',
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
    closeOptions(item) {
      if (item === 'userType') {
        this.$refs.fuzzyUserType.blur()
      } else {
        this.$refs.fuzzyUserStatus.blur()
      }
    }
  }
}
</script>

<style scoped>
.handle-box {
  margin-bottom: 20px;
}

.handle-select {
  width: 120px;
}

.handle-input {
  width: 180px;
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

.boy,
.manager {
  color: #23a6d5;
}

.girl,
.boss {
  color: var(--red);
}
</style>
