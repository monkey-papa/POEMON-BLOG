<template>
  <div>
    <myHeader />
    <sidebar />
    <div class="content-box">
      <div class="content">
        <router-view />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import myHeader from "../common/myHeader.vue";
import sidebar from "../common/sidebar.vue";
</script>

<style lang="scss" scoped>
.content-box {
  position: absolute;
  left: 200px;
  right: 0;
  top: 70px;
  bottom: 0;
  transition: left 0.3s ease-in-out;

  .content {
    width: auto;
    height: 100%;
    padding: 30px;
    overflow-y: scroll;
  }
}
</style>
