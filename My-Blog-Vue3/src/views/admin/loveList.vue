<template>
  <div>
    <div>
      <div class="header-search">
        <el-select
          @blur="closeOptions"
          @clear="handleClear('status')"
          ref="status"
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
      </div>
      <el-table :data="loves" border class="admin-table" header-cell-class-name="table-header">
        <el-table-column prop="id" label="ID" width="55" align="center" />
        <el-table-column prop="userId" label="用户ID" align="center" />
        <el-table-column prop="manName" label="男生昵称" align="center" />
        <el-table-column prop="womanName" label="女生昵称" align="center" />
        <el-table-column label="背景封面" align="center">
          <template v-slot="scope">
            <el-image
              lazy
              :preview-src-list="[scope.row.bgCover]"
              class="admin-image-thumb"
              :src="scope.row.bgCover"
              fit="cover"
            />
          </template>
        </el-table-column>
        <el-table-column label="男生头像" align="center">
          <template v-slot="scope">
            <el-image
              lazy
              :preview-src-list="[scope.row.manCover]"
              class="admin-image-thumb"
              :src="scope.row.manCover"
              fit="cover"
            />
          </template>
        </el-table-column>
        <el-table-column label="女生头像" align="center">
          <template v-slot="scope">
            <el-image
              lazy
              :preview-src-list="[scope.row.womanCover]"
              class="admin-image-thumb"
              :src="scope.row.womanCover"
              fit="cover"
            />
          </template>
        </el-table-column>
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
        <el-table-column prop="timing" label="计时" align="center" />
        <el-table-column prop="countdownTitle" label="倒计时标题" align="center" />
        <el-table-column prop="countdownTime" label="倒计时时间" align="center" />
        <el-table-column prop="familyInfo" label="额外信息" align="center" show-overflow-tooltip />
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
          layout="total, prev, pager, next"
          :current-page="pagination.current"
          :page-size="pagination.size"
          :total="pagination.total"
          @current-change="handlePageChange"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { notifySuccess } from "@/utils/notify";
import api from "@/api";
import { ref, onMounted, type Ref } from "vue";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
import type { Family } from "@/types/family";
import type { PaginationResponse } from "@/types/api";
import type { ElSelect } from "element-plus";

interface PaginationState {
  current: number;
  size: number;
  total: number;
  status: boolean | null;
}

const { $common, $confirm } = useGlobalProperties();

const pagination: Ref<PaginationState> = ref({
  current: 1,
  size: 10,
  total: 0,
  status: null,
});

const loves: Ref<Family[]> = ref([]);
const status: Ref<InstanceType<typeof ElSelect> | null> = ref(null);

onMounted(() => {
  getLoves();
});

const handleClear = (field: "status"): void => {
  pagination.value[field] = null;
  getLoves();
};

const handleDelete = async (item: Family): Promise<void> => {
  try {
    await $confirm("确认删除这对情侣吗？", "提示", {
      confirmButtonText: "确定",
      cancelButtonText: "取消",
      type: "warning",
      center: true,
    });

    await api.deleteFamily(item.id);
    pagination.value.current = 1;
    getLoves();
    notifySuccess("删除成功！");
  } catch (error) {
    if (error === "cancel") {
      notifySuccess("已取消删除！");
    }
  }
};

const search = (): void => {
  pagination.value.total = 0;
  pagination.value.current = 1;
  getLoves();
};

const getLoves = async (): Promise<void> => {
  try {
    const res = (await api.listFamily(pagination.value)) as PaginationResponse<Family>;
    loves.value = res.list || [];
    pagination.value.total = res.totalCount || 0;
  } catch {
    /* ignored */
  }
};

const changeStatus = async (item: Family): Promise<void> => {
  try {
    await api.changeLoveStatus({
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
  getLoves();
};

const closeOptions = (): void => {
  status.value?.blur();
};
</script>
