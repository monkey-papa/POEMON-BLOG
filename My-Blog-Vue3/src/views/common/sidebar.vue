<template>
  <div class="sidebar">
    <el-menu
      class="sidebar-el-menu"
      :default-active="route.path"
      background-color="var(--favoriteBg)"
      text-color="var(--black5)"
      active-text-color="var(--blue1)"
      unique-opened
      router
    >
      <el-menu-item v-for="item in items" :index="item.index" :key="item.index">
        <el-icon>
          <component :is="item.icon" />
        </el-icon>
        {{ item.title }}
      </el-menu-item>
    </el-menu>
  </div>
</template>
<script setup lang="ts">
import { ref, onMounted, computed } from "vue";
import { useRoute } from "vue-router";
import { useStore } from "@/stores";
import type { CurrentAdmin } from "@/types/user";

const route = useRoute();
const store = useStore();

const currentAdmin = computed(() => store.currentAdmin ?? ({} as CurrentAdmin));

interface MenuItem {
  icon: string;
  index: string;
  title: string;
}

const items = ref<MenuItem[]>([
  {
    icon: "HomeFilled",
    index: "/backendMain",
    title: "系统首页",
  },
  {
    icon: "Tools",
    index: "/webEdit",
    title: "网站设置",
  },
  {
    icon: "UserFilled",
    index: "/userList",
    title: "用户管理",
  },
  {
    icon: "Postcard",
    index: "/postList",
    title: "文章管理",
  },
  {
    icon: "Notebook",
    index: "/sortList",
    title: "分类管理",
  },
  {
    icon: "ChatRound",
    index: "/weiYanList",
    title: "微言管理",
  },
  {
    icon: "CircleCheck",
    index: "/progressList",
    title: "进展管理",
  },
  {
    icon: "Edit",
    index: "/commentList",
    title: "评论管理",
  },
  {
    icon: "Comment",
    index: "/treeHoleList",
    title: "留言管理",
  },
  {
    icon: "Paperclip",
    index: "/resourceList",
    title: "七牛云资源管理",
  },
  {
    icon: "PictureRounded",
    index: "/resourcePathList",
    title: "娱乐管理",
  },
  {
    icon: "Sugar",
    index: "/loveList",
    title: "表白墙管理",
  },
  {
    icon: "Failed",
    index: "/prohibitedWordsList",
    title: "违禁词管理",
  },
]);

onMounted(() => {
  if (currentAdmin.value.userType === 3) {
    items.value = items.value.filter((item) => item.index !== "/userList");
  }
});
</script>
<style lang="scss" scoped>
.sidebar {
  display: block;
  position: absolute;
  left: 0;
  top: 70px;
  bottom: 0;
  overflow-y: scroll;

  &::-webkit-scrollbar {
    width: 0;
  }

  &-el-menu {
    width: 200px;
  }

  > ul {
    height: 100%;
  }
}
</style>
