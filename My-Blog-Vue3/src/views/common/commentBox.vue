<template>
  <div class="comment-box-container">
    <!-- 框 -->
    <textarea
      class="comment-textarea"
      v-model="commentContent"
      placeholder="善语结善缘，恶言伤人心..."
      maxlength="100"
    />
    <!-- 按钮 -->
    <div class="myBetween">
      <div class="button-box-left">
        <div :class="{ 'emoji-active': showEmoji }" @click="showEmoji = !showEmoji">
          <el-icon :size="18" class="el-icon-orange myEmoji"><Orange /></el-icon>
        </div>
        <div @click="openPicture()">
          <el-icon :size="18" class="el-icon-picture myPicture"><PictureFilled /></el-icon>
        </div>
      </div>

      <div class="button-box-right">
        <proButton
          :info="'发射'"
          @click="submitComment()"
          :before="$constant?.before_color_1"
          :after="$constant?.after_color_1"
        />
      </div>
    </div>
    <!-- 表情 -->
    <emoji @addEmoji="addEmoji" :showEmoji="showEmoji" />

    <el-dialog
      title="图片"
      v-model="showPicture"
      width="25%"
      :append-to-body="true"
      destroy-on-close
      center
      class="custom-my-dialog"
    >
      <div>
        <uploadPicture
          :ResourceType="'commentPicture'"
          @addPicture="addPicture"
          :maxSize="5"
          :maxNumber="1"
        />
      </div>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from "vue";
import { defineAsyncComponent } from "vue";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
import { useFormValidation } from "@/composables/useFormValidation";
import { useStore } from "@/stores";

const emoji = defineAsyncComponent(() => import("./emoji.vue"));
const proButton = defineAsyncComponent(() => import("./proButton.vue"));
const uploadPicture = defineAsyncComponent(() => import("./uploadPicture.vue"));

interface Props {
  disableGraffiti?: boolean;
}

withDefaults(defineProps<Props>(), {
  disableGraffiti: false,
});

interface Emits {
  submitComment: [content: string];
}

const emit = defineEmits<Emits>();

const store = useStore();
const { $constant } = useGlobalProperties();
const { validateRequiredWarn } = useFormValidation();

const currentUser = computed(() => store.currentUser);

interface Picture {
  name: string;
  url: string;
}

const commentContent = ref<string>("");
const showEmoji = ref<boolean>(false);
const showPicture = ref<boolean>(false);
const picture = ref<Picture>({
  name: currentUser.value?.username || "",
  url: "",
});

const openPicture = (): void => {
  showPicture.value = true;
};

const addPicture = (res: string): void => {
  picture.value.url = res;
  savePicture();
};

const savePicture = (): void => {
  const img = "<" + picture.value.name + "," + picture.value.url + ">";
  commentContent.value += img;
  picture.value.url = "";
  showPicture.value = false;
};

const addEmoji = (key: string): void => {
  commentContent.value += key;
};

const submitComment = (): void => {
  if (!validateRequiredWarn({ [commentContent.value]: "写" })) return;
  emit("submitComment", commentContent.value.trim());
  commentContent.value = "";
};
</script>

<style lang="scss" scoped>
.comment-box-container {
  .comment-textarea {
    border: 2px solid var(--gray15);
    width: 100%;
    font-size: 14px;
    padding: 15px;
    min-height: 180px;
    resize: none;
    outline: none;
    border-radius: 4px;
    background-image: var(--comment1URL), var(--comment2URL);
    background-position: left, right;
    background-repeat: no-repeat, no-repeat;
    background-size: 25%, 30%;
    margin-bottom: 10px;

    &:focus {
      border-color: var(--orange2);
    }

    &::placeholder {
      color: var(--darkBlue);
    }
  }

  .myBetween {
    margin-bottom: 10px;

    .button-box-left {
      display: flex;

      .myEmoji {
        transition: all 0.3s ease;
        margin-right: 12px;
        color: var(--darkBlue);

        &:hover {
          transform: rotate(360deg);
          font-size: 22px !important;
          color: var(--black8);
        }
      }

      .myPicture {
        color: var(--darkBlue);
        transition: all 0.3s ease;

        &:hover {
          color: var(--black8);
        }
      }

      .emoji-active {
        color: var(--red);
      }
    }

    .button-box-right {
      display: flex;
    }
  }
}
</style>
