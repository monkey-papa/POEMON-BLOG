<template>
  <div>
    <div class="header-search">
      <el-input v-model="pagination.searchKey" placeholder="违禁词" class="mr10" clearable />
      <el-button type="primary" @click="searchForbiddenWord()">
        <el-icon color="white"><Search /></el-icon>搜索</el-button
      >
      <el-button v-hasPermi="['user:visit:read']" type="success" @click="openDialog(null, 0)"
        >添加</el-button
      >
      <el-button type="danger" @click="clearSearch()">清除参数</el-button>
    </div>
    <el-table
      :data="prohibitedWordsList"
      border
      class="admin-table"
      header-cell-class-name="table-header"
    >
      <el-table-column prop="id" label="ID" align="center" />
      <el-table-column prop="username" label="用户" align="center" />
      <el-table-column label="头像" width="100" align="center">
        <template v-slot="scope">
          <el-image
            :preview-src-list="[scope.row.avatar]"
            lazy
            class="admin-image-thumb"
            :src="scope.row.avatar"
            fit="cover"
          />
        </template>
      </el-table-column>
      <el-table-column prop="message" label="违禁词" align="center" />
      <el-table-column header-align="center" align="center" label="操作">
        <template v-slot="scope">
          <el-button v-hasPermi="['user:visit:read']" type="text" @click="openDialog(scope.row, 1)"
            ><el-icon color="#409eff"><EditPen /></el-icon>编辑</el-button
          >
          <el-button
            v-hasPermi="['user:visit:read']"
            type="text"
            @click="deleteProWords(scope.row.id)"
            ><el-icon color="#409eff"><Delete /></el-icon>删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>
    <el-dialog
      class="custom-my-dialog"
      :before-close="handleClose"
      :title="type === 0 ? '添加违禁词' : '修改违禁词'"
      v-model="dialogVisible"
      width="40%"
    >
      <el-input class="prohibited-words-input" v-model="row.message" placeholder="请输入违禁词">
        <template v-slot:prepend>违禁词</template>
      </el-input>
      <template v-slot:footer>
        <div class="dialog-footer">
          <el-button @click="handleClose()">取 消</el-button>
          <el-button type="primary" @click="addOrUpdateProWords(type)">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { notifySuccess } from "@/utils/notify";
import api from "@/api";
import { ref, onMounted, computed, type Ref } from "vue";
import { useStore } from "@/stores";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
import type { ProhibitedWord } from "@/types/prohibitedWord";
import type { PaginationResponse } from "@/types/api";
import type { CurrentAdmin } from "@/types/user";

const store = useStore();
const { $confirm } = useGlobalProperties();

interface ProhibitedWordForm {
  id?: number;
  message: string;
  username?: string;
  avatar?: string | null;
}

interface PaginationState {
  current: number;
  size: number;
  total: number;
  searchKey: string;
}

const currentAdmin = computed(() => store.currentAdmin ?? ({} as CurrentAdmin));

const row: Ref<ProhibitedWordForm> = ref({
  id: undefined,
  message: "",
  username: currentAdmin.value.username || "",
  avatar: currentAdmin.value.avatar || "",
});

const type: Ref<number> = ref(0);
const dialogVisible: Ref<boolean> = ref(false);
const pagination: Ref<PaginationState> = ref({
  current: 1,
  size: 100,
  total: 0,
  searchKey: "",
});

const prohibitedWordsList: Ref<ProhibitedWord[]> = ref([]);

onMounted(() => {
  getProhibitedWordsList();
});

const handleClose = (): void => {
  row.value = {
    id: undefined,
    message: "",
    username: currentAdmin.value.username || "",
    avatar: currentAdmin.value.avatar || "",
  };
  dialogVisible.value = false;
};

const clearSearch = (): void => {
  pagination.value = {
    current: 1,
    size: 100,
    total: 0,
    searchKey: "",
  };
  getProhibitedWordsList();
};

const openDialog = (rowData: ProhibitedWord | null, typeValue: number): void => {
  dialogVisible.value = true;
  type.value = typeValue;
  if (rowData !== null) {
    row.value = rowData;
  }
};

const addOrUpdateProWords = async (typeValue: number): Promise<void> => {
  // 添加违禁词
  if (typeValue === 0) {
    try {
      await api.addProhibitedWord(row.value);
      handleClose();
      row.value.message = "";
      notifySuccess("添加成功！");
      await getProhibitedWordsList();
    } catch {
      /* ignored */
    }
  }
  // 编辑违禁词
  if (typeValue === 1) {
    try {
      await api.updateProhibitedWord({
        id: row.value.id!,
        message: row.value.message,
      });
      notifySuccess("编辑成功！");
      handleClose();
      await getProhibitedWordsList();
    } catch {
      /* ignored */
    }
  }
};

const getProhibitedWordsList = async (): Promise<void> => {
  try {
    const res: PaginationResponse<ProhibitedWord> = await api.getProhibitedWordsList(
      pagination.value
    );
    prohibitedWordsList.value = res.list || [];
    pagination.value.total = res.totalCount || 0;
  } catch {
    /* ignored */
  }
};

const searchForbiddenWord = (): void => {
  getProhibitedWordsList();
};

const deleteProWords = (id: number): void => {
  $confirm("确定要删除这个违禁词吗？", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
    center: true,
  }).then(async () => {
    try {
      await api.deleteProhibitedWord(id);
      notifySuccess("删除成功！");
      await getProhibitedWordsList();
    } catch {
      /* ignored */
    }
  });
};
</script>

<style lang="scss">
.custom-my-dialog {
  .prohibited-words-input {
    width: 100%;
    height: 30px;
  }
}
</style>
