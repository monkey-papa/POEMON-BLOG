<template>
  <div class="resource-path-list-container">
    <div class="resource-path-list-header">
      <div class="header-search">
        <el-select
          @clear="clear"
          clearable
          v-model="pagination.resourceType"
          placeholder="娱乐类型"
          class="mr10"
        >
          <el-option
            v-for="(item, i) in resourceTypes"
            :key="i"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
        <el-select
          @clear="clear"
          clearable
          v-model="pagination.status"
          placeholder="状态"
          class="mr10"
        >
          <el-option key="1" label="启用" :value="true" />
          <el-option key="2" label="禁用" :value="false" />
        </el-select>
        <el-button type="primary" @click="search()"
          ><el-icon color="white"><Search /></el-icon>搜索</el-button
        >
        <el-button
          v-hasPermi="['user:visit:read']"
          type="primary"
          @click="addResourcePathDialog = true"
          >新增娱乐资源</el-button
        >
      </div>
      <el-table
        :data="resourcePaths"
        border
        class="admin-table"
        header-cell-class-name="table-header"
      >
        <el-table-column prop="id" label="ID" width="55" align="center" />
        <el-table-column prop="title" label="标题" align="center" />
        <el-table-column width="130" prop="classify" label="分类" align="center" />
        <el-table-column width="130" prop="type" label="资源类型" align="center" />
        <el-table-column width="200" prop="introduction" label="简介" align="center" />
        <el-table-column label="封面" align="center">
          <template v-slot="scope">
            <el-image
              lazy
              :preview-src-list="[scope.row.cover]"
              class="admin-image-thumb"
              :src="scope.row.cover"
              fit="cover"
            />
          </template>
        </el-table-column>
        <el-table-column width="300" prop="url" label="链接" align="center" />
        <el-table-column label="状态" align="center">
          <template v-slot="scope">
            <el-tag :type="scope.row.status === false ? 'danger' : 'success'" disable-transitions>
              {{ scope.row.status === false ? "禁用" : "启用" }}
            </el-tag>
            <el-switch
              class="admin-switch"
              v-hasPermi="['user:visit:read']"
              @click="changeStatus(scope.row)"
              v-model="scope.row.status"
            />
          </template>
        </el-table-column>
        <el-table-column label="友人头像" align="center">
          <template v-slot="scope">
            <el-image
              v-if="scope.row.type === 'friendUrl'"
              lazy
              :preview-src-list="[scope.row.friendAvatar]"
              class="admin-image-thumb"
              :src="scope.row.friendAvatar"
              fit="cover"
            />
          </template>
        </el-table-column>
        <el-table-column
          :formatter="$common.formatter"
          prop="createTime"
          label="创建时间"
          align="center"
        />
        <el-table-column label="操作" width="180" align="center">
          <template v-slot="scope">
            <el-button v-hasPermi="['user:visit:read']" type="text" @click="handleEdit(scope.row)"
              ><el-icon color="#409eff"><EditPen /></el-icon>编辑</el-button
            >
            <el-button
              style="color: var(--red)"
              v-hasPermi="['user:visit:read']"
              type="text"
              @click="handleDelete(scope.row)"
            >
              <el-icon color="var(--red)"><Delete /></el-icon>
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="admin-pagination">
        <el-pagination
          background
          layout="total, sizes, prev, pager, next"
          :current-page="pagination.current"
          :page-size="pagination.size"
          :page-sizes="[10, 20, 50, 100]"
          :total="pagination.total"
          @size-change="handleSizeChange"
          @current-change="handlePageChange"
        />
      </div>
    </div>

    <!-- 统一上传弹窗 -->
    <el-dialog
      :title="uploadDialogTitle"
      v-model="uploadDialogVisible"
      width="25%"
      :append-to-body="true"
      destroy-on-close
      center
    >
      <div>
        <uploadPicture
          :ResourceType="uploadResourceType"
          @addPicture="handleUploadSuccess"
          :maxSize="10"
          :maxNumber="1"
          :listType="uploadListType"
          :accept="uploadAccept"
        />
      </div>
    </el-dialog>

    <el-dialog
      class="resource-path-dialog"
      title="娱乐资源"
      v-model="addResourcePathDialog"
      width="50%"
      :before-close="clearDialog"
      :append-to-body="true"
      center
    >
      <div class="content">
        <div class="form">
          <div class="form-title">标题：</div>
          <el-input maxlength="60" v-model="resourcePath.title" />
          <div class="form-classify">分类：</div>
          <el-input
            :disabled="!['lovePhoto', 'funny', 'favorites'].includes(resourcePath.type)"
            maxlength="30"
            v-model="resourcePath.classify"
          />
          <div class="form-introduction">简介：</div>
          <el-input
            :disabled="!['friendUrl', 'favorites'].includes(resourcePath.type)"
            maxlength="1000"
            v-model="resourcePath.introduction"
          />
          <div class="form-cover">封面：</div>
          <div class="form-cover-input">
            <el-input v-model="resourcePath.cover" />
            <div class="form-cover-button-wrap">
              <proButton
                :info="'上传封面'"
                @click="addResourcePathCover()"
                :before="$constant.before_color_1"
                :after="$constant.after_color_1"
              />
            </div>
          </div>
          <div class="form-url">链接：</div>
          <div class="form-url-input">
            <el-input
              :disabled="!['friendUrl', 'funny', 'favorites'].includes(resourcePath.type)"
              v-model="resourcePath.url"
            />
            <div
              v-if="!['friendUrl', 'favorites'].includes(resourcePath.type)"
              class="form-url-button-wrap"
            >
              <proButton
                :info="'上传文件'"
                @click="addResourcePathUrl()"
                :before="$constant.before_color_1"
                :after="$constant.after_color_1"
              />
            </div>
          </div>
          <div class="form-type">资源类型：</div>
          <el-select clearable v-model="resourcePath.type" placeholder="资源路径类型" class="mr10">
            <el-option
              v-for="(item, i) in resourceTypes"
              :key="i"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
          <div class="form-friend-avatar">友人头像：</div>
          <div class="form-friend-avatar-input">
            <el-input
              :disabled="!['friendUrl'].includes(resourcePath.type)"
              v-model="resourcePath.friendAvatar"
            />
            <div class="form-friend-avatar-button-wrap">
              <proButton
                :info="'上传头像'"
                @click="addFriendAvatar()"
                :before="$constant.before_color_1"
                :after="$constant.after_color_1"
              />
            </div>
          </div>
        </div>
        <div class="button-wrap myCenter">
          <proButton
            :info="'提交'"
            @click="addResourcePath()"
            :before="$constant.before_color_1"
            :after="$constant.after_color_1"
          />
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, defineAsyncComponent, type Ref } from "vue";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
import { notifySuccess, notifyError } from "@/utils/notify";
import api from "@/api";
import type { ResourcePath } from "@/types/resource";
import type { PaginationResponse } from "@/types/api";

const uploadPicture = defineAsyncComponent(() => import("../common/uploadPicture.vue"));
const proButton = defineAsyncComponent(() => import("../common/proButton.vue"));

const { $common, $constant, $confirm } = useGlobalProperties();

interface ResourceTypeOption {
  label: string;
  value: string;
}

interface PaginationState {
  current: number;
  size: number;
  total: number;
  resourceType: string;
  status: boolean | null;
}

interface ResourcePathForm {
  id?: number;
  title: string;
  classify: string;
  introduction: string;
  cover: string;
  url: string;
  type: string;
  friendAvatar: string;
}

type UploadType = "cover" | "friendAvatar" | "file";
type UploadListType = "picture" | "text";

const resourceTypes: Ref<ResourceTypeOption[]> = ref([
  { label: "友链", value: "friendUrl" },
  { label: "相册", value: "lovePhoto" },
  { label: "音乐", value: "funny" },
  { label: "收藏夹", value: "favorites" },
]);
const pagination: Ref<PaginationState> = ref({
  current: 1,
  size: 10,
  total: 0,
  resourceType: "",
  status: null,
});
const resourcePaths: Ref<ResourcePath[]> = ref([]);
const uploadDialogVisible: Ref<boolean> = ref(false);
const uploadDialogTitle: Ref<string> = ref("");
const uploadType: Ref<UploadType | ""> = ref("");
const uploadResourceType: Ref<string> = ref(""); //'friendUrlUrl', 'friendUrl', 'friendAvatar', 'funny', 'funnyUrl', 'favorites', 'favoritesUrl', 'lovePhoto'
const uploadListType: Ref<UploadListType> = ref("picture");
const uploadAccept: Ref<string> = ref("image/*");
const addResourcePathDialog: Ref<boolean> = ref(false);
const isUpdate: Ref<boolean> = ref(false);
const resourcePath: Ref<ResourcePathForm> = ref({
  title: "",
  classify: "",
  introduction: "",
  cover: "",
  url: "",
  type: "",
  friendAvatar: "",
});

onMounted(() => {
  getResourcePaths();
});

// 统一处理上传成功
const handleUploadSuccess = (res: string): void => {
  if (uploadType.value === "friendAvatar") {
    resourcePath.value.friendAvatar = res;
  } else if (uploadType.value === "cover") {
    resourcePath.value.cover = res;
  } else if (uploadType.value === "file") {
    resourcePath.value.url = res;
  }
  uploadDialogVisible.value = false;
  uploadType.value = "";
};

// 上传文件（音乐等）
const addResourcePathUrl = (): void => {
  if (addResourcePathDialog.value === false) {
    return;
  }
  if (!["funny", "favorites", "friendUrl"].includes(resourcePath.value.type)) {
    notifyError("请选择资源类型！");
    return;
  }
  uploadType.value = "file";
  uploadDialogTitle.value = "上传文件";
  uploadResourceType.value = resourcePath.value.type + "Url";
  uploadListType.value = "text";
  uploadAccept.value = "image/*, video/*, audio/*";
  uploadDialogVisible.value = true;
};

// 上传封面
const addResourcePathCover = (): void => {
  if (addResourcePathDialog.value === false) {
    return;
  }
  if ($common.isEmpty(resourcePath.value.type)) {
    notifyError("请选择资源类型！");
    return;
  }
  uploadType.value = "cover";
  uploadDialogTitle.value = "上传封面";
  uploadResourceType.value = resourcePath.value.type;
  uploadListType.value = "picture";
  uploadAccept.value = "image/*";
  uploadDialogVisible.value = true;
};

// 上传友人头像
const addFriendAvatar = (): void => {
  if (addResourcePathDialog.value === false) {
    return;
  }
  if (!["friendUrl"].includes(resourcePath.value.type)) {
    notifyError("请选择资源类型！");
    return;
  }
  uploadType.value = "friendAvatar";
  uploadDialogTitle.value = "上传友人头像";
  uploadResourceType.value = "friendAvatar";
  uploadListType.value = "picture";
  uploadAccept.value = "image/*";
  uploadDialogVisible.value = true;
};

const addResourcePath = async (): Promise<void> => {
  if ($common.isEmpty(resourcePath.value.title) || $common.isEmpty(resourcePath.value.type)) {
    notifyError("标题和资源类型不能为空！");
    return;
  }

  try {
    if (isUpdate.value) {
      await api.updateResourcePath({ ...resourcePath.value, id: resourcePath.value.id! });
    } else {
      await api.saveResourcePath(resourcePath.value);
    }
    notifySuccess("保存成功！");
    addResourcePathDialog.value = false;
    clearDialog();
    search(pagination.value.current);
  } catch {
    /* ignored */
  }
};

const search = (cur?: number): void => {
  pagination.value.total = 0;
  pagination.value.current = cur || 1;
  getResourcePaths();
};

const getResourcePaths = async (): Promise<void> => {
  try {
    const res: PaginationResponse<ResourcePath> = await api.getResourcePathList(pagination.value);
    resourcePaths.value = res.list || [];
    pagination.value.total = res.totalCount || 0;
  } catch {
    /* ignored */
  }
};

const changeStatus = async (item: ResourcePath): Promise<void> => {
  try {
    await api.updateResourcePath({ ...item, id: item.id });
    notifySuccess("修改成功！");
  } catch {
    /* ignored */
  }
};

const handlePageChange = (val: number): void => {
  pagination.value.current = val;
  getResourcePaths();
};

const handleSizeChange = (val: number): void => {
  pagination.value.size = val;
  getResourcePaths();
};

const handleDelete = async (item: ResourcePath): Promise<void> => {
  try {
    await $confirm("确认删除？", "提示", {
      confirmButtonText: "确定",
      cancelButtonText: "取消",
      type: "warning",
      center: true,
    });

    await api.deleteResourcePath(item.id);
    search(pagination.value.current);
    notifySuccess("删除成功！");
  } catch (error) {
    if (error === "cancel") {
      notifySuccess("已取消删除！");
    }
  }
};

const handleEdit = (item: ResourcePath): void => {
  resourcePath.value = JSON.parse(JSON.stringify(item));
  addResourcePathDialog.value = true;
  isUpdate.value = true;
};

const clearDialog = (): void => {
  isUpdate.value = false;
  addResourcePathDialog.value = false;
  resourcePath.value = {
    title: "",
    classify: "",
    introduction: "",
    cover: "",
    url: "",
    type: "",
    friendAvatar: "",
  };
};

const clear = (): void => {
  pagination.value.status = null;
  getResourcePaths();
};
</script>

<style lang="scss">
.resource-path-dialog {
  .content {
    .form {
      .form-title {
        margin-bottom: 5px;
      }

      .form-classify,
      .form-introduction,
      .form-cover,
      .form-url,
      .form-type,
      .form-friend-avatar {
        margin-top: 10px;
        margin-bottom: 5px;
      }

      .form-cover-input,
      .form-url-input,
      .form-friend-avatar-input {
        display: flex;

        .form-cover-button-wrap,
        .form-url-button-wrap,
        .form-friend-avatar-button-wrap {
          width: 66px;
          margin: 3.5px 0 0 10px;
        }
      }
    }

    .button-wrap {
      display: flex;
      margin-top: 30px;
    }
  }
}
</style>
