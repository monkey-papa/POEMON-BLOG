<template>
  <div class="message-container">
    <div class="message-content">
      <div
        style="animation: header-effect 2s; top: 0"
        :style="{ background: `${changeBgState}` }"
        class="background-image background-image-changeBg"
      ></div>
      <!-- 输入框 -->
      <div class="message-in">
        <h2 class="message-title">树洞</h2>
        <h4 class="message-title">想说的·想问的·吐槽·交流</h4>
        <div>
          <input
            ref="focus"
            class="message-input"
            type="text"
            placeholder="留下点什么啦~"
            v-model="messageContent"
            @click="show = true"
            maxlength="60"
          />
          <button v-show="show" @click="submitMessage" class="message-input message-input-button">
            发射
          </button>
        </div>
      </div>
      <!-- 弹幕 -->
      <div class="barrage-container">
        <vue-danmaku
          :danmus="barrageList"
          :loop="true"
          :isSuspend="true"
          :top="20"
          :speeds="100"
          useSlot
          style="height: 100%; width: 100%"
        >
          <template v-slot:dm="{ danmu }">
            <div class="danmaku-name">
              <span class="bullet-item" :style="{ color: getRandomColor() }">
                <img :src="danmu.avatar" alt="用户头像" />
                {{ danmu.msg }}</span
              >
            </div>
          </template>
        </vue-danmaku>
      </div>
    </div>
    <div class="comment-wrap">
      <div class="comment-content">
        <comment :source="0" :type="'message'" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { notifySuccess } from "@/utils/notify";
import api from "@/api";
import { ref, onMounted, computed, type Ref, type ComputedRef } from "vue";
import vueDanmaku from "vue3-danmaku";
// eslint-disable-next-line @typescript-eslint/no-unused-expressions
vueDanmaku; // 注册弹幕组件
import { defineAsyncComponent } from "vue";
import { useStore } from "@/stores";
import { useFormValidation } from "@/composables/useFormValidation";
import constant from "@/utils/constant";
import type { TreeHole } from "@/types";

interface BarrageItem {
  id: number;
  avatar: string;
  msg: string;
  time: number;
}

interface TreeHoleForm {
  message: string;
  avatar: string;
  username?: string;
}

const store = useStore();
const { validateRequiredWarn } = useFormValidation();

const changeBgState: ComputedRef<string> = computed(() => store.changeBg);
const currentUser = computed(() => store.currentUser);

const comment = defineAsyncComponent(() => import("./common/comment.vue"));

const show: Ref<boolean> = ref(false);
const messageContent: Ref<string> = ref("");
const barrageList: Ref<BarrageItem[]> = ref([]);

const colorList: string[] = [
  "rgb(204,255,255)",
  "white",
  "rgb(204,255,204)",
  "white",
  "rgb(0,255,255)",
  "white",
  "rgb(255,204,255)",
  "pink",
];

const focus: Ref<HTMLInputElement | null> = ref(null);

onMounted(() => {
  getTreeHole();
  focus.value?.focus();
});

const getRandomColor = (): string => {
  const color = colorList[Math.floor(Math.random() * 8)];
  return color;
};

const getTreeHole = async (): Promise<void> => {
  try {
    const res = await api.getBossTreeHoleList({ noPagination: true });
    (res.list || []).forEach((m: TreeHole) => {
      barrageList.value.push({
        id: m.id,
        avatar: m.avatar || "",
        msg: m.message,
        time: Math.floor(Math.random() * 10 + 5),
      });
    });
  } catch {
    /* ignored */
  }
};

const submitMessage = async (): Promise<void> => {
  if (!validateRequiredWarn({ [messageContent.value]: "写" })) return;

  const treeHole: TreeHoleForm = {
    message: messageContent.value.trim(),
    avatar: currentUser.value?.avatar || constant.defaultAvatar,
  };
  treeHole.username = currentUser.value?.username;

  try {
    const res = await api.saveTreeHole(treeHole);
    // 通知博主已由后端自动处理
    notifySuccess("发射成功！");
    barrageList.value.push({
      id: res.id,
      avatar: res.avatar || "",
      msg: res.message,
      time: Math.floor(Math.random() * 10 + 5),
    });
  } catch {
    /* ignored */
  }
  messageContent.value = "";
  show.value = false;
};
</script>

<style lang="scss" scoped>
.message-container {
  .message-content {
    .message-in {
      position: absolute;
      left: 50%;
      top: 40%;
      transform: translate(-50%, -50%);
      color: var(--green2);
      animation: hideToShow 2.5s;
      width: 360px;
      z-index: 10;
      text-align: center;

      .message-title {
        user-select: none;
        text-align: center;
        font-size: 24px;
        font-weight: 400;
      }

      .message-input {
        border-radius: 1.2rem;
        border: var(--lightRed) 1px solid;
        color: var(--lightRed);
        background: var(--transparent);
        padding: 10px 10px;
        outline: none;
        width: 70%;

        &::-webkit-input-placeholder {
          color: var(--lightRed);
        }

        &.message-input-button {
          margin-left: 12px;
          width: 20%;
        }
      }
    }

    .barrage-container {
      position: absolute;
      top: 50px;
      left: 0;
      right: 0;
      bottom: 0;
      height: calc(100% - 50px);
      width: 100%;
      user-select: none;
      overflow: hidden;

      .danmaku-name {
        .bullet-item {
          white-space: nowrap;
          background-color: rgba(0, 0, 0, 0.7);
          border-radius: 30px;
          height: 30px;
          font-size: 16px;
          color: #ffffff;
          line-height: 30px;
          padding-right: 20px;
          display: flex;

          img {
            width: 30px;
            height: 30px;
            border-radius: 50%;
            margin-right: 10px;
          }
        }
      }
    }
  }

  .comment-wrap {
    margin-top: 100vh;
    background: var(--background);
    width: 100%;

    .comment-content {
      max-width: 900px;
      margin: 0 auto;
      padding: 40px 20px;
    }
  }
}
</style>
