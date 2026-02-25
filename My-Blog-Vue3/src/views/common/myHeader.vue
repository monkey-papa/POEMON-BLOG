<template>
  <div class="my-header myBetween">
    <div class="logo">博客后台管理</div>
    <div class="header-right">
      <div class="admin-index" @click="router.push({ path: '/' })">
        <img class="admin-index-icon" src="../../assets/svg/backend.svg" />
        <span>&nbsp;去往前台</span>
      </div>
      <div class="header-user-avatar">
        <el-dropdown popper-class="user-avatar-dropdown" placement="bottom">
          <el-avatar class="user-avatar" :size="40" :src="currentAdmin.avatar" />
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item @click="logout()">退出</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { notifySuccess } from "@/utils/notify";
import { useStore } from "@/stores";
import { computed } from "vue";
import { useRouter } from "vue-router";
import type { CurrentAdmin } from "@/types/user";

const store = useStore();
const router = useRouter();

const currentAdmin = computed(() => store.currentAdmin ?? ({} as CurrentAdmin));

const logout = (): void => {
  notifySuccess("退出登录成功！");
  store.loadCurrentAdmin(null);
  router.push({ path: "/" });
};
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

  .header-right {
    display: flex;
    justify-content: flex-end;
    margin-right: 40px;

    .admin-index {
      height: 70px;
      line-height: 70px;
      font-size: 15px;
      margin-right: 20px;

      .admin-index-icon {
        vertical-align: -6px;
        margin-right: 2px;
      }
    }

    .header-user-avatar {
      display: flex;
      align-items: center;
    }
  }
}
</style>
