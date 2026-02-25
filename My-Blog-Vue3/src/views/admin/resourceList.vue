<template>
  <div>
    <div>
      <div class="header-search">
        <el-select
          @clear="handleClear('resourceType')"
          clearable
          v-model="pagination.resourceType"
          placeholder="资源类型"
          class="mr10"
        >
          <el-option
            v-for="item in resourceTypes"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
        <el-button type="primary" @click="search()"
          ><el-icon color="white"><Search /></el-icon>搜索</el-button
        >
        <el-button v-hasPermi="['user:visit:read']" type="primary" @click="addResources()"
          >新增资源</el-button
        >
        <el-button type="danger" @click="clearSearch()">清除参数</el-button>
      </div>
      <el-table :data="resources" border class="admin-table" header-cell-class-name="table-header">
        <el-table-column prop="id" label="ID" width="55" align="center" />
        <el-table-column prop="userId" label="用户ID" align="center" />
        <el-table-column prop="username" label="用户名" align="center" width="120" />
        <el-table-column prop="mimeType" label="资源类型" align="center" />
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
        <el-table-column label="路径" align="center">
          <template v-slot="scope">
            <template v-if="!$common?.isEmpty(scope.row.mimeType)">
              <el-image
                lazy
                :preview-src-list="[scope.row.path]"
                class="admin-image-thumb"
                :src="scope.row.path"
                fit="cover"
              />
            </template>
            <template v-else>
              {{ scope.row.path }}
            </template>
          </template>
        </el-table-column>
        <el-table-column label="大小(KB)" align="center">
          <template v-slot="scope">
            {{ Math.round(scope.row.size / 1024) }}
          </template>
        </el-table-column>
        <el-table-column prop="type" label="类型" align="center" />
        <el-table-column
          :formatter="$common?.formatter"
          prop="createTime"
          label="创建时间"
          align="center"
        />
        <el-table-column label="操作" width="180" align="center">
          <template v-slot="scope">
            <el-button
              v-hasPermi="['user:visit:read']"
              type="text"
              style="color: var(--red)"
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
          @current-change="handlePageChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog
      title="图片"
      v-model="resourceDialog"
      width="50%"
      :append-to-body="true"
      @close="close"
      destroy-on-close
      center
    >
      <uploadPicture
        :ResourceType="pagination.resourceType"
        @addPicture="addPicture"
        :maxSize="5"
        :maxNumber="2"
      />
    </el-dialog>
  </div>
</template>
<script setup lang="ts">
import { notifySuccess, notifyError } from "@/utils/notify";
import api from "@/api";
import { ref, onMounted, type Ref } from "vue";
import { defineAsyncComponent } from "vue";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
import type { Resource } from "@/types/resource";
import type { PaginationResponse } from "@/types/api";

const uploadPicture = defineAsyncComponent(() => import("../common/uploadPicture.vue"));

interface PaginationState {
  current: number;
  size: number;
  total: number;
  resourceType: string;
}

interface ResourceTypeOption {
  label: string;
  value: string;
}

const { $common, $confirm } = useGlobalProperties();

const pagination: Ref<PaginationState> = ref({
  current: 1,
  size: 10,
  total: 0,
  resourceType: "",
});

const resources: Ref<Resource[]> = ref([]);
const resourceDialog: Ref<boolean> = ref(false);
const resourceTypes: Ref<ResourceTypeOption[]> = ref([
  { label: "用户头像", value: "userAvatar" },
  { label: "文章封面", value: "articleCover" },
  { label: "文章图片", value: "articlePicture" },
  { label: "PC背景图片", value: "webBackgroundImage" },
  { label: "随机头像", value: "randomAvatar" },
  { label: "其他图片", value: "otherPhoto" },
  { label: "评论图片", value: "commentPicture" },
  { label: "Love.Cover", value: "love/bgCover" },
  { label: "Love.Man", value: "love/manCover" },
  { label: "Love.Woman", value: "love/womanCover" },
  { label: "Mobile背景图片", value: "mobilePhoto" },
  { label: "友链头像", value: "friendAvatar" },
  { label: "友链封面", value: "friendUrl" },
  { label: "音乐封面", value: "funny" },
  { label: "音乐图片", value: "funnyUrl" },
  { label: "收藏夹封面", value: "favorites" },
  { label: "相册", value: "lovePhoto" },
]);

onMounted(() => {
  getResources();
});

const handleClear = (field: "resourceType"): void => {
  pagination.value[field] = "";
  getResources();
};

const close = (): void => {
  getResources();
};

const clearSearch = (): void => {
  pagination.value = {
    current: 1,
    size: 10,
    total: 0,
    resourceType: "",
  };
  getResources();
};

const handleDelete = (item: Resource): void => {
  $confirm("确认删除资源？", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
    center: true,
  })
    .then(async () => {
      try {
        await api.deleteArticleImage({
          url: item.path,
          id: item.id,
        });
        pagination.value.current = 1;
        await getResources();
        notifySuccess("删除成功！");
      } catch {
        /* ignored */
      }
    })
    .catch(() => {
      notifySuccess("已取消删除！");
    });
};

const addPicture = (_res: unknown): void => {
  resourceDialog.value = false;
};

const addResources = (): void => {
  if ($common.isEmpty(pagination.value.resourceType)) {
    notifyError("请选择资源类型！");
    return;
  }
  resourceDialog.value = true;
};

const search = (): void => {
  pagination.value.total = 0;
  pagination.value.current = 1;
  getResources();
};

const getResources = async (): Promise<void> => {
  try {
    const res = (await api.getResourceList(pagination.value)) as PaginationResponse<Resource>;
    resources.value = res.list || [];
    pagination.value.total = res.totalCount || 0;
  } catch {
    /* ignored */
  }
};

const changeStatus = async (item: Resource): Promise<void> => {
  try {
    await api.changeResourceStatus({
      id: item.id,
      flag: item.status,
    });
    notifySuccess("修改成功！");
  } catch {
    /* ignored */
  }
};

const handlePageChange = (val: number): void => {
  pagination.value.current = val;
  getResources();
};

const handleSizeChange = (val: number): void => {
  pagination.value.size = val;
  getResources();
};
</script>
