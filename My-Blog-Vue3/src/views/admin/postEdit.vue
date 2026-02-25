<template>
  <div class="post-edit-container">
    <el-tag effect="dark" class="admin-title-tag">
      <img class="admin-title-tag-img" src="../../assets/svg/document.svg" />
      文章信息
    </el-tag>
    <el-form
      :model="article"
      :rules="rules"
      ref="ruleForm"
      label-width="110px"
      class="demo-ruleForm"
    >
      <el-form-item label="文章标题" prop="articleTitle">
        <el-input placeholder="请输入文章标题" maxlength="30" v-model="article.articleTitle" />
      </el-form-item>
      <div class="post-edit-default-author">默认文章作者：{{ constant.siteName }}</div>
      <el-form-item label="文章作者" prop="articleAuthor">
        <el-input maxlength="30" v-model="article.articleAuthor" />
      </el-form-item>
      <el-form-item label="文章内容" prop="articleContent">
        <MdEditor
          ref="md"
          v-model="article.articleContent"
          @onUploadImg="onUploadImg"
          style="height: 500px"
        />
      </el-form-item>
      <el-form-item label="是否启用评论" prop="commentStatus">
        <el-tag :type="article.commentStatus === false ? 'danger' : 'success'" disable-transitions>
          {{ article.commentStatus === false ? "否" : "是" }}
        </el-tag>
        <el-switch class="admin-switch" v-model="article.commentStatus" />
      </el-form-item>
      <el-form-item label="是否推荐" prop="recommendStatus" v-hasPermi="['user:visit:read']">
        <el-tag
          :type="article.recommendStatus === false ? 'danger' : 'success'"
          disable-transitions
        >
          {{ article.recommendStatus === false ? "否" : "是" }}
        </el-tag>
        <el-switch class="admin-switch" v-model="article.recommendStatus" />
      </el-form-item>
      <el-form-item label="是否可见" prop="viewStatus" v-hasPermi="['user:visit:read']">
        <el-tag :type="article.viewStatus === false ? 'danger' : 'success'" disable-transitions>
          {{ article.viewStatus === false ? "否" : "是" }}
        </el-tag>
        <el-switch class="admin-switch" v-model="article.viewStatus" />
      </el-form-item>
      <!-- <el-form-item
        v-if="article.viewStatus === false"
        label="访问密码"
        prop="password"
      >
        <el-input maxlength="30" v-model="article.password"></el-input>
      </el-form-item> -->
      <el-form-item label="封面" prop="articleCover">
        <div class="post-edit-article-cover">
          <el-input v-model="article.articleCover" />
          <el-image
            v-if="article.articleCover"
            class="admin-image-thumb"
            lazy
            :preview-src-list="[article.articleCover]"
            :src="article.articleCover"
            fit="cover"
          />
        </div>
        <uploadPicture
          class="post-edit-upload-picture"
          :ResourceType="'articleCover'"
          @addPicture="addArticleCover"
          :maxSize="3"
          :maxNumber="1"
        />
      </el-form-item>
      <el-form-item label="分类" prop="sortId">
        <el-select class="post-edit-select" v-model="article.sortId" placeholder="请选择分类">
          <el-option v-for="item in sorts" :key="item.id" :label="item.sortName" :value="item.id" />
        </el-select>
      </el-form-item>
      <el-form-item label="标签" prop="labelId">
        <el-select class="post-edit-select" v-model="article.labelId" placeholder="请选择标签">
          <el-option
            v-for="item in labelsTemp"
            :key="item.id"
            :label="item.labelName"
            :value="item.id"
          />
        </el-select>
      </el-form-item>
    </el-form>
    <div class="post-edit-button-container myCenter">
      <el-button type="primary" @click="submitForm()">保存</el-button>
      <el-button type="danger" @click="resetForm()">重置所有修改</el-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { notifySuccess, notifyError, notifyWarning } from "@/utils/notify";
import api from "@/api";
import http from "@/utils/request";
import { ref, watch, onMounted, defineAsyncComponent, computed } from "vue";
import { useStore } from "@/stores";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
// 导入 md-editor-v3
import { MdEditor } from "md-editor-v3";
import "md-editor-v3/lib/style.css";
import type { SaveArticleParams, UpdateArticleParams } from "@/types/article";
import type { SortInfo, LabelInfo } from "@/types/sort";
import type { FormInstance } from "element-plus";
import type { SaveResourceParams } from "@/types/resource";
import type { CurrentAdmin } from "@/types/user";
import { useRouter } from "vue-router";
import constant from "@/utils/constant";

const store = useStore();
const { $route, $common, $constant, $confirm } = useGlobalProperties();
const router = useRouter();

interface ArticleForm {
  id?: number;
  articleTitle: string;
  articleAuthor: string;
  articleContent: string;
  commentStatus: boolean;
  recommendStatus: boolean;
  viewStatus: boolean;
  password: string;
  articleCover: string;
  sortId: number | null;
  labelId: number | null;
}

interface FormRules {
  articleTitle: Array<{ required: boolean; message: string; trigger: string }>;
  articleAuthor: Array<{ required: boolean; message: string; trigger: string }>;
  articleContent: Array<{ required: boolean; message: string; trigger: string }>;
  commentStatus: Array<{ required: boolean; message: string; trigger: string }>;
  recommendStatus: Array<{ required: boolean; message: string; trigger: string }>;
  viewStatus: Array<{ required: boolean; message: string; trigger: string }>;
  articleCover: Array<{ required: boolean; message: string; trigger: string }>;
  sortId: Array<{ required: boolean; message: string; trigger: string }>;
  labelId: Array<{ required: boolean; message: string; trigger: string }>;
}

interface UploadResult {
  url: string;
  alt: string;
  title: string;
}

const currentAdmin = computed(() => store.currentAdmin ?? ({} as CurrentAdmin));

const uploadPicture = defineAsyncComponent(() => import("../common/uploadPicture.vue"));

const ruleForm = ref<FormInstance | null>(null);
const md = ref<InstanceType<typeof MdEditor> | null>(null);

const id = ref<number | null>($route.query.id ? parseInt($route.query.id as string) : null);
const article = ref<ArticleForm>({
  articleTitle: "",
  articleAuthor: constant.siteName,
  articleContent: "",
  commentStatus: true,
  recommendStatus: false,
  viewStatus: false,
  password: "",
  articleCover: "",
  sortId: null,
  labelId: null,
});
const sorts = ref<SortInfo[]>([]);
const labels = ref<LabelInfo[]>([]);
const labelsTemp = ref<LabelInfo[]>([]);
const rules = ref<FormRules>({
  articleTitle: [{ required: true, message: "请输入标题", trigger: "change" }],
  articleAuthor: [{ required: true, message: "请输入作者名称", trigger: "change" }],
  articleContent: [{ required: true, message: "请输入内容", trigger: "change" }],
  commentStatus: [{ required: true, message: "是否启用评论", trigger: "change" }],
  recommendStatus: [{ required: true, message: "是否推荐", trigger: "change" }],
  viewStatus: [{ required: true, message: "是否可见", trigger: "change" }],
  articleCover: [{ required: true, message: "封面", trigger: "change" }],
  sortId: [{ required: true, message: "分类", trigger: "change" }],
  labelId: [{ required: true, message: "标签", trigger: "blur" }],
});

watch(
  () => article.value.sortId,
  (newVal: number | null, oldVal: number | null) => {
    if (oldVal !== null) {
      article.value.labelId = null;
    }
    if (!$common.isEmpty(newVal) && !$common.isEmpty(labels.value)) {
      labelsTemp.value = labels.value.filter((l) => l.sortId === newVal);
    }
  }
);

onMounted(() => {
  article.value.viewStatus = currentAdmin.value.userType !== 3;
  getSortAndLabel();
});

// md-editor-v3 图片上传处理
const onUploadImg = async (files: File[], callback: (urls: string[]) => void): Promise<void> => {
  const uploadResults = await Promise.all(
    files.map((file: File): Promise<UploadResult> => {
      return new Promise((resolve) => {
        const fd = new FormData();
        fd.append("image", file);
        fd.append("userId", String(currentAdmin.value.id));
        fd.append("folder", "articlePicture");
        // 上传图片到七牛云
        http
          .uploadQiniu($constant.qiniuUploadImages, fd)
          .then(async (res: unknown) => {
            const uploadRes = res as { url?: string };
            if (!uploadRes.url) {
              notifyWarning("上传出错！");
              resolve({ url: "", alt: file.name, title: file.name });
              return;
            }
            notifySuccess("上传图片成功！");
            // 保存图片地址到数据库
            try {
              const resource: SaveResourceParams = {
                type: file.type,
                path: uploadRes.url,
                size: file.size,
                mimeType: "articlePicture",
              };
              await api.saveResource(resource);
            } catch {
              /* ignored */
            }
            resolve({
              url: uploadRes.url,
              alt: file.name,
              title: file.name,
            });
          })
          .catch(() => {
            notifyWarning("上传出错！");
            resolve({ url: "", alt: file.name, title: file.name });
          });
      });
    })
  );
  // 回调插入图片
  callback(uploadResults.map((item) => item.url));
};

const addArticleCover = (res: string): void => {
  article.value.articleCover = res;
};

const getSortAndLabel = async (): Promise<void> => {
  try {
    const res = await api.getSortAndLabelList();
    const response = res as unknown as { sorts?: SortInfo[]; labels?: LabelInfo[] };
    sorts.value = Array.isArray(response?.sorts) ? response.sorts : [];
    labels.value = Array.isArray(response?.labels) ? response.labels : [];
    if (!$common.isEmpty(id.value)) {
      await editArticle();
    }
  } catch {
    /* ignored */
  }
};

const editArticle = async (): Promise<void> => {
  try {
    if (id.value === null) return;
    const res = await api.getAdminArticleById(id.value);
    if (res) {
      article.value = {
        ...res,
        articleCover: res.articleCover || "",
        articleAuthor: res.articleAuthor || constant.siteName,
        password: "",
      };
    }
  } catch {
    /* ignored */
  }
};

const submitForm = (): void => {
  // if (
  //   article.value.viewStatus === false &&
  //   $common.isEmpty(article.value.password)
  // ) {
  //   notifyError("文章不可见时必须输入密码！");
  //   return;
  // }
  if (ruleForm.value) {
    ruleForm.value.validate((valid: boolean) => {
      if (valid) {
        if ($common.isEmpty(id.value)) {
          saveArticle(article.value);
        } else {
          article.value.id = id.value ?? undefined;
          saveArticle(article.value);
        }
      } else {
        notifyError("请完善必填项！");
      }
    });
  }
};

const resetForm = (): void => {
  if (ruleForm.value) {
    ruleForm.value.resetFields();
  }
  if (!$common.isEmpty(id.value)) {
    editArticle();
  }
};

const saveArticle = (value: ArticleForm): void => {
  $confirm("确认保存？", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
    center: true,
  })
    .then(async () => {
      try {
        if ($common.isEmpty(id.value)) {
          const articleParams: SaveArticleParams = {
            ...value,
            sortId: value.sortId!,
            labelId: value.labelId!,
          };
          await api.saveArticle(articleParams);
        } else {
          const articleParams: UpdateArticleParams = {
            ...value,
            id: id.value!,
            sortId: value.sortId!,
            labelId: value.labelId!,
          };
          await api.updateArticle(articleParams);
        }
        router.push({ path: "/postList" });
        if (currentAdmin.value.userType !== 0) {
          try {
            await api.sendCodeComment({
              email: constant.bossEmail,
              comment: `${currentAdmin.value.username ?? ""}${
                $common.isEmpty(id.value) ? "新增" : "修改"
              }文章：${value.articleTitle}，请尽快审核！`,
              name: currentAdmin.value.username ?? "",
              type: "approve",
            });
            notifySuccess("发布成功，已发送邮件通知博主审核！");
          } catch {
            /* ignored */
          }
        }
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
.post-edit-container {
  .admin-title-tag {
    .admin-title-tag-img {
      vertical-align: -3px;
    }
  }

  .demo-ruleForm {
    .el-form-item {
      margin-bottom: 40px;

      .post-edit-select {
        width: 200px;
      }
    }
  }

  .post-edit-default-author {
    color: var(--lightRed);
    padding-left: 34px;
    margin-bottom: 5px;
  }

  .post-edit-article-cover {
    display: flex;

    .admin-image-thumb {
      margin-left: 10px;
    }
  }

  .post-edit-upload-picture {
    margin-top: 10px;
  }

  .post-edit-button-container {
    margin-bottom: 22px;
  }

  .markdown-body {
    height: 400px;
  }
}
</style>
