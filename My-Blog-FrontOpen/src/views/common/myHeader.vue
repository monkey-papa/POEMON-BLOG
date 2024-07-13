<template>
  <div class="my-header myBetween">
    <div class="logo">博客后台管理</div>
    <div class="header-right">
      <div class="admin-index" @click="$router.push({ path: '/' })">
        <img style="vertical-align: -6px; margin-right: 2px" src="../../assets/svg/backend.svg" />
        <span>&nbsp;去往前台</span>
      </div>
      <div class="header-user-con">
        <el-dropdown placement="bottom">
          <el-avatar class="user-avatar" :size="40" :src="$store.state.currentAdmin.avatar">
          </el-avatar>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item @click.native="logout()">退出</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  methods: {
    logout() {
      this.$notify({
        title: '可以啦🍨',
        message: '退出登录成功！',
        type: 'success',
        offset: 50,
        position: 'top-left'
      })
      this.$store.commit('loadCurrentAdmin', {})
      localStorage.removeItem('adminToken')
      this.$router.push({ path: '/' })
    }
  }
}
</script>
<style lang="scss" scoped>
.my-header {
  position: relative;
  width: 100%;
  height: 70px;
  color: var(--black);
  background-color: var(--favoriteBg);
  .logo {
    line-height: 70px;
    margin-left: 70px;
    font-size: 22px;
  }
}
.header-right {
  display: flex;
  justify-content: flex-end;
  margin-right: 40px;
}
.admin-index {
  height: 70px;
  line-height: 70px;
  font-size: 15px;
  margin-right: 20px;
}
.header-user-con {
  display: flex;
  align-items: center;
}
</style>
