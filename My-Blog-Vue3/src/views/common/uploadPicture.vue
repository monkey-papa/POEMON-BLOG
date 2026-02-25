<template>
  <div class="upload-picture-container">
    <el-upload
      class="upload-picture"
      ref="upload"
      multiple
      drag
      name="image"
      :action="$constant.qiniuUploadImages"
      :headers="uploadHeaders"
      :on-change="handleChange"
      :on-success="handleSuccess"
      :on-error="handleError"
      :list-type="listType"
      :accept="accept"
      :limit="maxNumber"
      :data="{
        userId: $route?.meta?.requiresAuth ? currentAdmin?.id : currentUser?.id,
        folder: props.ResourceType || 'images',
      }"
      :auto-upload="false"
    >
      <div class="el-upload__text">
        <img src="../../assets/svg/upload.svg" />
        <div>拖拽上传 / 点击上传</div>
      </div>
      <template v-if="listType === 'picture'" #tip>
        <div class="el-upload__tip">
          一次最多上传{{ maxNumber }}张图片，且每张图片不超过{{ maxSize }}M！
        </div>
      </template>
      <template v-else #tip>
        <div class="el-upload__tip">
          一次最多上传{{ maxNumber }}个文件，且每个文件不超过{{ maxSize }}M！
        </div>
      </template>
    </el-upload>
    <div class="upload-picture-button">
      <el-button type="success" @click="submitUpload"> 上传 </el-button>
    </div>
  </div>
</template>
<script setup lang="ts">
import { notifySuccess, notifyWarning } from "@/utils/notify";
import api from "@/api";
import { ref, computed } from "vue";
import { useStore } from "@/stores";
import { useGlobalProperties } from "@/composables/useGlobalProperties";

import type { UploadInstance, UploadFile, UploadFiles } from "element-plus";
import type { SaveResourceParams } from "@/types/resource";

const store = useStore();
const { $route } = useGlobalProperties();

interface Props {
  listType?: string;
  accept?: string;
  maxSize?: number;
  maxNumber?: number;
  ResourceType?: string;
}

const props = withDefaults(defineProps<Props>(), {
  listType: "picture",
  accept: "image/*",
  maxSize: 5,
  maxNumber: 5,
  ResourceType: "",
});

interface Emits {
  (e: "addPicture", url: string): void;
}

const emit = defineEmits<Emits>();

const currentUser = computed(() => store.currentUser);
const currentAdmin = computed(() => store.currentAdmin);

interface UploadHeaders {
  Authorization: string;
}

// 上传请求头（需要认证 token）
const uploadHeaders = computed<UploadHeaders>(() => {
  const token = $route?.meta?.requiresAuth
    ? currentAdmin.value?.accessToken
    : currentUser.value?.accessToken;
  return {
    Authorization: token ? "Token " + token : "",
  };
});

const upload = ref<UploadInstance | null>(null);

// 上传点击事件
const submitUpload = (): void => {
  upload.value?.submit();
};

interface UploadResponse {
  data?: {
    url?: string;
  };
  url?: string;
}

// 文件上传成功时的钩子
const handleSuccess = async (response: UploadResponse, file: UploadFile): Promise<void> => {
  const url = response.data?.url || response.url;

  if (!url) {
    notifyWarning("上传失败，未获取到图片地址！");
    return;
  }

  // 存取资源接口
  const resource: SaveResourceParams = {
    type: file.raw?.type || "",
    path: url,
    size: file.size ?? 0,
    mimeType: props.ResourceType,
  };

  try {
    await api.saveResource(resource);
  } catch {
    /* ignored */
  }

  emit("addPicture", url);
  notifySuccess("上传成功！");
};

// 文件上传失败时的钩子
const handleError = (): void => {
  notifyWarning("上传出错！");
};

// 添加文件、上传成功和上传失败时都会被调用
const handleChange = (file: UploadFile, fileList: UploadFiles): void => {
  let flag = false;
  if (file.size && file.size > props.maxSize * 1024 * 1024) {
    notifyWarning("图片最大为" + props.maxSize + "M！");
    flag = true;
  }
  if (flag) {
    fileList.splice(fileList.length - 1, 1);
  }
};
</script>

<style lang="scss" scoped>
.upload-picture-container {
  .upload-picture {
    .el-upload__text {
      img {
        margin-top: 10px;
      }
    }
  }

  .upload-picture-button {
    text-align: center;
    margin-top: 10px;
  }
}
</style>
