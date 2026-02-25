<template>
  <div class="weiYan-container">
    <div
      style="animation: header-effect 2s"
      :style="{ background: `${changeBgState}` }"
      class="background-image background-image-changeBg"
    ></div>
    <!-- 两句诗 -->
    <div class="my-animation-slide-top">
      <twoPoem :isHitokoto="false" />
    </div>
    <div class="weiYan-content">
      <div>
        <treeHole
          :treeHoleList="treeHoleList"
          :avatar="currentUser ? currentUser.avatar : webInfo.avatar"
          @launch="launch"
          @deleteTreeHole="deleteTreeHole"
        />
        <proPage
          :current="pagination.current"
          :size="pagination.size"
          :total="pagination.total"
          :buttonSize="3"
          :color="$constant?.commentPageColor"
          @toPage="toPage"
        />
      </div>
    </div>
    <el-dialog
      title="微言"
      v-model="weiYanDialogVisible"
      width="40%"
      :before-close="handleClose"
      :append-to-body="true"
      destroy-on-close
      center
      class="weiYan-dialog custom-my-dialog"
    >
      <div>
        <div class="weiYan-dialog-content myCenter">
          <el-radio-group v-model="isPublic">
            <el-radio-button :label="true">公开</el-radio-button>
            <el-radio-button :label="false">私密</el-radio-button>
          </el-radio-group>
        </div>
        <commentBox :disableGraffiti="true" @submitComment="submitWeiYan" />
      </div>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { notifySuccess } from "@/utils/notify";
import api from "@/api";
import { ref, onMounted, nextTick, computed, type Ref, type ComputedRef } from "vue";
import { defineAsyncComponent } from "vue";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
import type { SaveWeiYanParams, WebInfo, PaginationState } from "@/types";

const twoPoem = defineAsyncComponent(() => import("./common/twoPoem.vue"));
const treeHole = defineAsyncComponent(() => import("./common/treeHole.vue"));
const proPage = defineAsyncComponent(() => import("./common/proPage.vue"));
const commentBox = defineAsyncComponent(() => import("./common/commentBox.vue"));

import { useStore } from "@/stores";

/** 微言列表展示项（复用树洞结构，扩展微言特有字段） */
interface WeiYanDisplayItem {
  id: number;
  message: string;
  avatar: string;
  createTime: string;
  userId: number;
  content: string;
  isPublic?: boolean;
  username?: string;
}

const store = useStore();
const { $constant, $common, $confirm } = useGlobalProperties();
const currentUser = computed(() => store.currentUser);
const changeBgState: ComputedRef<string> = computed(() => store.changeBg);
const webInfo: ComputedRef<WebInfo> = computed(() => store.webInfo);

const treeHoleList: Ref<WeiYanDisplayItem[]> = ref([]);
const pagination: Ref<PaginationState> = ref({
  current: 1,
  size: 10,
  total: 0,
});
const weiYanDialogVisible: Ref<boolean> = ref(false);
const isPublic: Ref<boolean> = ref(true);
const showFooter: Ref<boolean> = ref(false);

onMounted(() => {
  getWeiYan();
});

const toPage = (page: number): void => {
  pagination.value.current = page;
  window.scrollTo({
    top: 240,
    behavior: "smooth",
  });
  getWeiYan();
};

const launch = (): void => {
  weiYanDialogVisible.value = true;
};

const handleClose = (): void => {
  weiYanDialogVisible.value = false;
};

const submitWeiYan = async (content: string): Promise<void> => {
  const weiYan: SaveWeiYanParams = {
    content,
    isPublic: isPublic.value,
  };

  try {
    await api.saveWeiYan(weiYan);
    getWeiYan();
  } catch {
    /* ignored */
  }
  handleClose();
};

const deleteTreeHole = async (id: number): Promise<void> => {
  try {
    await $confirm("确认删除？", "提示", {
      confirmButtonText: "确定",
      cancelButtonText: "取消",
      type: "warning",
      center: true,
    });

    await api.deleteWeiYan(id);
    notifySuccess("删除成功！");
    pagination.value.current = 1;
    getWeiYan();
  } catch (error: unknown) {
    if (error === "cancel") {
      notifySuccess("已取消删除！");
    }
  }
};

const getWeiYan = async (): Promise<void> => {
  try {
    const res = await api.getWeiYanList(pagination.value);
    showFooter.value = false;
    const list: WeiYanDisplayItem[] = (res.list || []).map((weiYan): WeiYanDisplayItem => {
      const content = weiYan.content || "";
      const processedContent = $common.processContent(content);

      return {
        id: weiYan.id,
        message: weiYan.content,
        avatar: weiYan.avatar || "",
        createTime: weiYan.createTime,
        userId: weiYan.userId,
        content: processedContent,
        isPublic: weiYan.isPublic,
      };
    });
    treeHoleList.value = list;
    pagination.value.total = res.totalCount || 0;
    nextTick(() => {
      showFooter.value = true;
    });
  } catch {
    /* ignored */
  }
};
</script>

<style lang="scss">
.weiYan-dialog {
  .weiYan-dialog-content {
    padding-bottom: 20px;
  }
}
</style>
<style lang="scss" scoped>
.weiYan-container {
  .weiYan-content {
    background: var(--background);
    animation: hideToShow 2.5s;
  }
}
</style>
