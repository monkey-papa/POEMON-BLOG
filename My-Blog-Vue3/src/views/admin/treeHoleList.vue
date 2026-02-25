<template>
  <div class="tree-hole-list-container">
    <el-tag effect="dark" class="admin-title-tag">
      <img src="../../assets/svg/message.svg" />
      留言列表
    </el-tag>
    <el-table :data="treeHoles" border class="admin-table" header-cell-class-name="table-header">
      <el-table-column prop="id" label="ID" width="55" align="center" />
      <el-table-column prop="username" label="用户" align="center" />
      <el-table-column label="头像" align="center">
        <template v-slot="scope">
          <el-image
            :preview-src-list="[scope.row.avatar]"
            lazy
            class="admin-image-thumb"
            :src="scope.row.avatar || webInfo.avatar"
            fit="cover"
          />
        </template>
      </el-table-column>
      <el-table-column prop="message" label="留言内容" align="center" />
      <el-table-column
        :formatter="$common.formatter"
        prop="createTime"
        label="创建时间"
        align="center"
      />
      <el-table-column label="操作" width="180" align="center">
        <template v-slot="scope">
          <el-button
            v-hasPermi="['user:visit:read']"
            type="text"
            style="color: green"
            @click="addProWords(scope.row)"
          >
            <el-icon color="green"><CirclePlusFilled /></el-icon>
            添加违禁词
          </el-button>
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
        layout="total, prev, pager, next"
        :current-page="pagination.current"
        :page-size="pagination.size"
        :total="pagination.total"
        @current-change="handlePageChange"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { notifySuccess } from "@/utils/notify";
import api from "@/api";
import { ref, onMounted, computed, type Ref, type ComputedRef } from "vue";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
import { useStore } from "@/stores";
import type { TreeHole } from "@/types/webInfo";
import type { PaginationResponse } from "@/types/api";
import type { WebInfo } from "@/types/webInfo";

interface PaginationState {
  current: number;
  size: number;
  total: number;
}

type TreeHoleWithUsername = TreeHole;

const store = useStore();
const { $common, $confirm } = useGlobalProperties();

const webInfo: ComputedRef<WebInfo> = computed(() => store.webInfo);

const pagination: Ref<PaginationState> = ref({
  current: 1,
  size: 10,
  total: 0,
});

const treeHoles: Ref<TreeHoleWithUsername[]> = ref([]);

onMounted(() => {
  getTreeHoles();
});

const getTreeHoles = async (): Promise<void> => {
  try {
    const res = (await api.getBossTreeHoleList(
      pagination.value
    )) as PaginationResponse<TreeHoleWithUsername>;
    treeHoles.value = res.list || [];
    pagination.value.total = res.totalCount || 0;
  } catch {
    /* ignored */
  }
};

const handlePageChange = (val: number): void => {
  pagination.value.current = val;
  getTreeHoles();
};

const handleDelete = (item: TreeHoleWithUsername): void => {
  $confirm("确定要删除这条留言吗？", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
    center: true,
  })
    .then(() => {
      doDelete(item);
    })
    .catch(() => {
      notifySuccess("已取消删除！");
    });
};

const doDelete = async (item: TreeHoleWithUsername): Promise<void> => {
  try {
    await api.deleteTreeHole(item.id);
    pagination.value.current = 1;
    await getTreeHoles();
    notifySuccess("删除成功！");
  } catch {
    /* ignored */
  }
};

const addProWords = (item: TreeHoleWithUsername): void => {
  $confirm("确定要添加为违禁词吗？", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
    center: true,
  }).then(async () => {
    try {
      await api.addProhibitedWord({
        message: item.message,
        username: item.username,
        avatar: item.avatar,
      });
      // 添加成功后删除树洞
      await doDelete(item);
    } catch {
      /* ignored */
    }
  });
};
</script>

<style lang="scss" scoped>
.tree-hole-list-container {
  .admin-title-tag {
    img {
      vertical-align: -6px;
    }
  }
}
</style>
