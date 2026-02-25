<template>
  <div class="web-edit-container">
    <!-- 基础信息 -->
    <div class="web-edit-container-item">
      <el-tag effect="dark" class="admin-title-tag">
        <img class="admin-title-tag-img" src="../../assets/svg/edit.svg" />
        基础信息
      </el-tag>
      <el-form
        :model="webInfo"
        :rules="rules"
        ref="ruleForm"
        label-width="100px"
        class="demo-ruleForm"
      >
        <el-form-item label="网站名称" prop="webName">
          <el-input v-model="webInfo.webName" />
        </el-form-item>
        <el-form-item label="网站标题" prop="webTitle">
          <el-input v-model="webInfo.webTitle" />
        </el-form-item>
        <el-form-item label="页脚" prop="footer">
          <el-input v-model="webInfo.footer" />
        </el-form-item>
        <el-form-item v-hasPermi="['user:visit:read']" label="状态" prop="status">
          <el-switch @click="changeWebStatus(webInfo)" v-model="webInfo.status" />
        </el-form-item>
        <el-form-item label="背景" prop="backgroundImage">
          <div class="web-edit">
            <el-input v-model="webInfo.backgroundImage" />
            <el-image
              v-if="webInfo.avatar"
              lazy
              class="admin-image-thumb"
              :preview-src-list="[webInfo.backgroundImage]"
              :src="webInfo.backgroundImage"
              fit="cover"
            />
          </div>
          <uploadPicture
            v-hasPermi="['user:visit:read']"
            :ResourceType="'webBackgroundImage'"
            class="web-edit-container-item-upload-picture"
            @addPicture="addBackgroundImage"
            :maxSize="5"
            :maxNumber="1"
          />
        </el-form-item>
        <el-form-item label="默认头像" prop="avatar">
          <div class="web-edit">
            <el-input v-model="webInfo.avatar" />
            <el-image
              v-if="webInfo.avatar"
              lazy
              class="admin-image-thumb"
              :preview-src-list="[webInfo.avatar]"
              :src="webInfo.avatar"
              fit="cover"
            />
          </div>
          <uploadPicture
            v-hasPermi="['user:visit:read']"
            :ResourceType="'userAvatar'"
            class="web-edit-container-item-upload-picture"
            @addPicture="addAvatar"
            :maxSize="5"
            :maxNumber="1"
          />
        </el-form-item>
        <el-form-item label="提示" prop="waifuJson">
          <div class="web-edit">
            <el-input :disabled="disabled" :rows="6" type="textarea" v-model="webInfo.waifuJson" />
            <el-icon class="edit-icon" @click="disabled = !disabled"><EditPen /></el-icon>
          </div>
        </el-form-item>
      </el-form>
      <div class="web-edit-container-item-button-wrap myCenter">
        <el-button v-hasPermi="['user:visit:read']" type="primary" @click="submitForm('ruleForm')"
          >保存</el-button
        >
      </div>
    </div>
    <!-- 公告 -->
    <div class="web-edit-container-item">
      <el-tag effect="dark" class="admin-title-tag">
        <img class="admin-title-tag-img" src="../../assets/svg/edit.svg" />
        公告
      </el-tag>
      <el-tag
        :key="i"
        v-for="(notice, i) in notices"
        closable
        :disable-transitions="false"
        @close="handleClose(notices, notice)"
      >
        {{ notice }}
      </el-tag>
      <el-input
        class="input-new-tag"
        v-if="inputNoticeVisible"
        v-model="inputNoticeValue"
        ref="saveNoticeInput"
        size="small"
        @keyup.enter="handleInputNoticeConfirm"
        @blur="handleInputNoticeConfirm"
      />
      <el-button v-else class="button-new-tag" size="small" @click="showNoticeInput()"
        >+ 公告</el-button
      >
      <div class="web-edit-container-item-button-wrap myCenter">
        <el-button v-hasPermi="['user:visit:read']" type="primary" @click="saveNotice()"
          >保存</el-button
        >
      </div>
    </div>
    <!-- 随机名称 -->
    <div class="web-edit-container-item">
      <el-tag effect="dark" class="admin-title-tag">
        <img class="admin-title-tag-img" src="../../assets/svg/edit.svg" />
        随机名称
      </el-tag>
      <el-tag
        :key="i"
        effect="dark"
        v-for="(name, i) in randomName"
        closable
        :disable-transitions="false"
        :type="types[Math.floor(Math.random() * 5)]"
        @close="handleClose(randomName, name)"
      >
        {{ name }}
      </el-tag>
      <el-input
        class="input-new-tag"
        v-if="inputRandomNameVisible"
        v-model="inputRandomNameValue"
        ref="saveRandomNameInput"
        size="small"
        @keyup.enter="handleInputRandomNameConfirm"
        @blur="handleInputRandomNameConfirm"
      />
      <el-button v-else class="button-new-tag" size="small" @click="showRandomNameInput"
        >+ 随机名称</el-button
      >
      <div class="web-edit-container-item-button-wrap myCenter">
        <el-button v-hasPermi="['user:visit:read']" type="primary" @click="saveRandomName()"
          >保存</el-button
        >
      </div>
    </div>
    <!-- 随机头像 -->
    <div class="web-edit-container-item">
      <el-tag effect="dark" class="admin-title-tag">
        <img class="admin-title-tag-img" src="../../assets/svg/edit.svg" />
        随机头像
      </el-tag>
      <div :key="i" class="web-edit" v-for="(avatar, i) in randomAvatar">
        <el-tag
          class="web-edit-item"
          closable
          :disable-transitions="false"
          @close="handleClose(randomAvatar, avatar)"
        >
          {{ avatar }}
        </el-tag>
        <div class="web-edit-item-image">
          <el-image
            lazy
            class="admin-image-thumb"
            :preview-src-list="[avatar]"
            :src="avatar"
            fit="cover"
          />
        </div>
      </div>
      <el-input
        class="input-new-tag"
        v-if="inputRandomAvatarVisible"
        v-model="inputRandomAvatarValue"
        ref="saveRandomAvatarInput"
        size="small"
        @keyup.enter="handleInputRandomAvatarConfirm"
        @blur="handleInputRandomAvatarConfirm"
      />
      <el-button v-else class="button-new-tag" size="small" @click="showRandomAvatarInput"
        >+ 随机头像</el-button
      >
      <uploadPicture
        v-hasPermi="['user:visit:read']"
        :ResourceType="'randomAvatar'"
        class="web-edit-container-item-upload-picture"
        @addPicture="addRandomAvatar"
        :maxSize="1"
        :maxNumber="5"
      />
      <div class="web-edit-container-item-button-wrap myCenter">
        <el-button v-hasPermi="['user:visit:read']" type="primary" @click="saveRandomAvatar()"
          >保存</el-button
        >
      </div>
    </div>
    <!-- 其他图片 -->
    <div class="web-edit-container-item">
      <el-tag effect="dark" class="admin-title-tag">
        <img class="admin-title-tag-img" src="../../assets/svg/edit.svg" />
        其他图片（随机封面等）
      </el-tag>
      <div class="web-edit web-edit-cover" :key="i" v-for="(cover, i) in otherPhoto">
        <el-tag
          class="web-edit-item"
          closable
          :disable-transitions="false"
          @close="handleClose(otherPhoto, cover)"
        >
          {{ cover }}
        </el-tag>
        <div class="web-edit-item-image">
          <el-image
            lazy
            class="admin-image-thumb"
            :preview-src-list="[cover]"
            :src="cover"
            fit="cover"
          />
        </div>
      </div>
      <el-input
        class="input-new-tag"
        v-if="inputOtherPhotoVisible"
        v-model="inputOtherPhotoValue"
        ref="saveOtherPhotoInput"
        size="small"
        @keyup.enter="handleInputOtherPhotoConfirm"
        @blur="handleInputOtherPhotoConfirm"
      />
      <el-button v-else class="button-new-tag" size="small" @click="showOtherPhotoInput"
        >+ 其他图片</el-button
      >
      <uploadPicture
        v-hasPermi="['user:visit:read']"
        :ResourceType="'otherPhoto'"
        class="web-edit-container-item-upload-picture"
        @addPicture="addOtherPhoto"
        :maxSize="5"
        :maxNumber="5"
      />
      <div
        class="web-edit-container-item-button-wrap web-edit-container-item-button-wrap-last myCenter"
      >
        <el-button v-hasPermi="['user:visit:read']" type="primary" @click="saveOtherPhoto()"
          >保存</el-button
        >
      </div>
    </div>
    <!-- 重置所有修改 -->
    <div>
      <el-button v-hasPermi="['user:visit:read']" type="danger" @click="resetForm('ruleForm')"
        >重置所有修改</el-button
      >
    </div>
  </div>
</template>
<script setup lang="ts">
import { notifySuccess, notifyError } from "@/utils/notify";
import api from "@/api";
import { ref, onMounted, nextTick, defineAsyncComponent, type Ref } from "vue";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
import type { FormInstance, FormRules } from "element-plus";
import type { WebInfo } from "@/types/webInfo";

const uploadPicture = defineAsyncComponent(() => import("../common/uploadPicture.vue"));

const { $confirm } = useGlobalProperties();

interface WebInfoForm {
  id?: number;
  webName: string;
  webTitle: string;
  footer: string;
  backgroundImage: string;
  avatar: string;
  waifuJson: string;
  status: boolean;
}

interface UpdateWebInfoParams {
  id?: number;
  notices?: string;
  randomName?: string;
  randomAvatar?: string;
  randomCover?: string;
  webName?: string;
  webTitle?: string;
  footer?: string;
  backgroundImage?: string;
  avatar?: string;
  waifuJson?: string;
  status?: boolean;
}

const ruleForm: Ref<FormInstance | null> = ref(null);
const saveNoticeInput: Ref<HTMLInputElement | null> = ref(null);
const saveRandomNameInput: Ref<HTMLInputElement | null> = ref(null);
const saveRandomAvatarInput: Ref<HTMLInputElement | null> = ref(null);
const saveOtherPhotoInput: Ref<HTMLInputElement | null> = ref(null);

// 控制提示框是否能被选中
const disabled: Ref<boolean> = ref(true);
// 控制随机名称的按钮背景色
const types: Ref<string[]> = ref(["", "success", "info", "danger", "warning"]);
// 控制+公告按钮的出现或者隐藏
const inputNoticeVisible: Ref<boolean> = ref(false);
const inputNoticeValue: Ref<string> = ref("");
// 控制+随机名称按钮的出现或者隐藏
const inputRandomNameVisible: Ref<boolean> = ref(false);
const inputRandomNameValue: Ref<string> = ref("");
// 控制+随机头像按钮的出现或者隐藏
const inputRandomAvatarVisible: Ref<boolean> = ref(false);
const inputRandomAvatarValue: Ref<string> = ref("");
// 控制+其他图片按钮的出现或者隐藏
const inputOtherPhotoVisible: Ref<boolean> = ref(false);
const inputOtherPhotoValue: Ref<string> = ref("");
const webInfo: Ref<WebInfoForm> = ref({
  id: undefined,
  webName: "",
  webTitle: "",
  footer: "",
  backgroundImage: "",
  avatar: "",
  // 提示
  waifuJson: "",
  status: false,
});
// 公告
const notices: Ref<string[]> = ref([]);
const randomAvatar: Ref<string[]> = ref([]);
const randomName: Ref<string[]> = ref([]);
const otherPhoto: Ref<string[]> = ref([]);
const rules: Ref<FormRules<WebInfoForm>> = ref({
  webName: [
    { required: true, message: "请输入网站名称", trigger: "blur" },
    {
      min: 1,
      max: 12,
      message: "长度在 1 到 12 个字符",
      trigger: "change",
    },
  ],
  webTitle: [{ required: true, message: "请输入网站标题", trigger: "blur" }],
  footer: [{ required: true, message: "请输入页脚", trigger: "blur" }],
  backgroundImage: [{ required: true, message: "请输入背景", trigger: "change" }],
  status: [{ required: true, message: "请设置网站状态", trigger: "change" }],
  avatar: [{ required: true, message: "请上传头像", trigger: "change" }],
});

onMounted(() => {
  getWebInfo();
});

// 接收传过来的url
const addBackgroundImage = (res: string): void => {
  webInfo.value.backgroundImage = res;
};

const addAvatar = (res: string): void => {
  webInfo.value.avatar = res;
};

const addRandomAvatar = (res: string): void => {
  randomAvatar.value.push(res);
};

const addOtherPhoto = (res: string): void => {
  otherPhoto.value.push(res);
};

const changeWebStatus = async (webInfoItem: WebInfoForm): Promise<void> => {
  try {
    await api.updateAdminWebInfo({
      id: webInfoItem.id,
      status: webInfoItem.status,
    });
    notifySuccess("保存成功！");
  } catch {
    /* ignored */
  }
};

const getWebInfo = async (): Promise<void> => {
  try {
    const res: WebInfo = await api.getAdminWebInfo();
    webInfo.value.id = res.id;
    webInfo.value.webName = res.webName;
    webInfo.value.webTitle = Array.isArray(res.webTitle) ? res.webTitle.join("") : res.webTitle;
    webInfo.value.footer = res.footer;
    webInfo.value.backgroundImage = res.backgroundImage;
    webInfo.value.avatar = res.avatar;
    webInfo.value.waifuJson = res.waifuJson || "";
    webInfo.value.status = res.status || false;
    // 安全解析 JSON 字段
    try {
      notices.value = res.notices
        ? Array.isArray(res.notices)
          ? res.notices
          : JSON.parse(res.notices)
        : [];
    } catch {
      notices.value = [];
    }
    try {
      randomAvatar.value = res.randomAvatar
        ? Array.isArray(res.randomAvatar)
          ? res.randomAvatar
          : JSON.parse(res.randomAvatar)
        : [];
    } catch {
      randomAvatar.value = [];
    }
    try {
      randomName.value = res.randomName
        ? Array.isArray(res.randomName)
          ? res.randomName
          : JSON.parse(res.randomName)
        : [];
    } catch {
      randomName.value = [];
    }
    try {
      otherPhoto.value = res.randomCover
        ? Array.isArray(res.randomCover)
          ? res.randomCover
          : JSON.parse(res.randomCover)
        : [];
    } catch {
      otherPhoto.value = [];
    }
  } catch {
    /* ignored */
  }
};

const submitForm = (_formName: string): void => {
  if (ruleForm.value) {
    ruleForm.value.validate((valid: boolean) => {
      if (valid) {
        updateWebInfo(webInfo.value);
      } else {
        notifyError("请完善必填项！");
      }
    });
  }
};

const resetForm = (_formName: string): void => {
  if (ruleForm.value) {
    ruleForm.value.resetFields();
  }
  getWebInfo();
};

const handleClose = (array: string[], item: string): void => {
  const index = array.indexOf(item);
  if (index > -1) {
    array.splice(index, 1);
  }
};

const handleInputNoticeConfirm = (): void => {
  if (inputNoticeValue.value) {
    notices.value.push(inputNoticeValue.value);
  }
  inputNoticeVisible.value = false;
  inputNoticeValue.value = "";
};

const showNoticeInput = (): void => {
  inputNoticeVisible.value = true;
  nextTick(() => {
    if (saveNoticeInput.value) {
      saveNoticeInput.value.focus();
    }
  });
};

const saveNotice = (): void => {
  const param: UpdateWebInfoParams = {
    id: webInfo.value.id,
    notices: JSON.stringify(notices.value),
  };
  updateWebInfo(param);
};

const handleInputRandomNameConfirm = (): void => {
  if (inputRandomNameValue.value) {
    randomName.value.push(inputRandomNameValue.value);
  }
  inputRandomNameVisible.value = false;
  inputRandomNameValue.value = "";
};

const showRandomNameInput = (): void => {
  inputRandomNameVisible.value = true;
  nextTick(() => {
    if (saveRandomNameInput.value) {
      saveRandomNameInput.value.focus();
    }
  });
};

const saveRandomName = (): void => {
  const param: UpdateWebInfoParams = {
    id: webInfo.value.id,
    randomName: JSON.stringify(randomName.value),
  };
  updateWebInfo(param);
};

const handleInputRandomAvatarConfirm = (): void => {
  if (inputRandomAvatarValue.value) {
    randomAvatar.value.push(inputRandomAvatarValue.value);
  }
  inputRandomAvatarVisible.value = false;
  inputRandomAvatarValue.value = "";
};

const showRandomAvatarInput = (): void => {
  inputRandomAvatarVisible.value = true;
  nextTick(() => {
    if (saveRandomAvatarInput.value) {
      saveRandomAvatarInput.value.focus();
    }
  });
};

const saveRandomAvatar = (): void => {
  const param: UpdateWebInfoParams = {
    id: webInfo.value.id,
    randomAvatar: JSON.stringify(randomAvatar.value),
  };
  updateWebInfo(param);
};

const handleInputOtherPhotoConfirm = (): void => {
  if (inputOtherPhotoValue.value) {
    otherPhoto.value.push(inputOtherPhotoValue.value);
  }
  inputOtherPhotoVisible.value = false;
  inputOtherPhotoValue.value = "";
};

const showOtherPhotoInput = (): void => {
  inputOtherPhotoVisible.value = true;
  nextTick(() => {
    if (saveOtherPhotoInput.value) {
      saveOtherPhotoInput.value.focus();
    }
  });
};

const saveOtherPhoto = (): void => {
  const param: UpdateWebInfoParams = {
    id: webInfo.value.id,
    randomCover: JSON.stringify(otherPhoto.value),
  };
  updateWebInfo(param);
};

const updateWebInfo = (value: UpdateWebInfoParams | WebInfoForm): void => {
  $confirm("确认保存？", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
    center: true,
  })
    .then(async () => {
      try {
        await api.updateAdminWebInfo(value);
        notifySuccess("保存成功！");
      } catch {
        /* ignored */
      }
    })
    .catch(() => {
      notifySuccess("已取消保存！");
    });
};
</script>
<style lang="scss" scoped>
.web-edit-container {
  .admin-title-tag {
    .admin-title-tag-img {
      vertical-align: -3px;
    }
  }

  .web-edit-container-item {
    .web-edit {
      display: flex;

      &.web-edit-cover {
        margin-bottom: 10px;
      }

      .admin-image-thumb {
        margin-left: 10px;
      }

      .edit-icon {
        margin-left: 10px;
        font-size: 18px;
        font-weight: bold;
        color: var(--maxLightRed);
      }

      .web-edit-item {
        white-space: normal;
        height: unset;
      }

      .web-edit-item-image {
        .admin-image-thumb {
          margin: 10px;
        }
      }
    }

    .web-edit-container-item-upload-picture {
      margin-top: 10px;
    }

    .web-edit-container-item-button-wrap {
      margin-top: 10px;

      &.web-edit-container-item-button-wrap-last {
        margin-bottom: 40px;
      }
    }

    .button-new-tag {
      height: 32px;
      line-height: 32px;
      padding-top: 0;
      padding-bottom: 0;
      margin: 0 10px;

      &:hover {
        background-color: var(--green2);
        border-color: var(--green2);
        color: var(--white);
      }
    }

    .input-new-tag {
      width: 200px;
      margin: 0 10px;
      height: 32px;

      :deep(.el-input__wrapper.is-focus) {
        box-shadow: 0 0 0 1px var(--green2) inset;
      }
    }
  }
}
</style>
