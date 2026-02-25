<template>
  <div class="sort-list-container">
    <div class="sort-list-header-button-wrap">
      <el-button v-hasPermi="['user:visit:read']" type="primary" @click="openSortDialog"
        >新增分类</el-button
      >
    </div>
    <!-- 表格 -->
    <div class="sort-list-header-title">分类</div>
    <el-table :data="sortInfo" border class="admin-table" header-cell-class-name="table-header">
      <el-table-column prop="id" label="ID" width="55" align="center" />
      <el-table-column prop="sortName" label="分类名称" align="center" />
      <el-table-column prop="sortDescription" label="分类描述" align="center" />
      <el-table-column label="分类类型" align="center">
        <template v-slot="scope">
          <span v-if="scope.row.sortType === 0">导航栏分类</span>
          <span v-else-if="scope.row.sortType === 1">普通分类</span>
        </template>
      </el-table-column>
      <el-table-column prop="priority" label="分类优先级" align="center" />
      <el-table-column prop="lengthOfLabel" label="标签总数" align="center" />
      <el-table-column prop="countOfSort" label="文章总数" align="center" />
      <el-table-column label="操作" width="380" align="center">
        <template v-slot="scope">
          <el-button
            v-hasPermi="['user:visit:read']"
            type="text"
            @click="editSortAndLabel(scope.row, '编辑分类')"
          >
            <el-icon color="#409eff"><EditPen /></el-icon>
            编辑分类
          </el-button>
          <el-button v-hasPermi="['user:visit:read']" type="text" @click="seeLabel(scope.row)">
            <el-icon color="#409eff"><View /></el-icon>
            查看标签
          </el-button>
          <el-button v-hasPermi="['user:visit:read']" type="text" @click="insertLabel(scope.row)">
            <el-icon color="#409eff"><CirclePlusFilled /></el-icon>
            新增标签
          </el-button>
          <el-button
            v-hasPermi="['user:visit:read']"
            type="text"
            style="color: var(--red)"
            @click="deleteHandle(scope.row.id, 1)"
          >
            <el-icon color="var(--red)"><Delete /></el-icon>
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <!-- 标签弹窗 -->
    <el-dialog width="80%" title="标签" v-model="dialogTableVisible">
      <el-table
        v-if="!$common.isEmpty(sort)"
        :data="sort.labels"
        border
        header-cell-class-name="table-header"
      >
        <el-table-column prop="id" label="ID" width="55" align="center" />
        <el-table-column label="分类名称" align="center">
          <span>{{ sort.sortName }}</span>
        </el-table-column>
        <el-table-column prop="labelName" label="标签名称" align="center" />
        <el-table-column prop="labelDescription" label="标签描述" align="center" />
        <el-table-column prop="countOfLabel" label="文章总数" align="center" />
        <el-table-column label="操作" width="320" align="center">
          <template v-slot="scope">
            <el-button type="text" @click="editSortAndLabel(scope.row, '编辑标签')">
              <el-icon color="#409eff"><EditPen /></el-icon>
              编辑标签
            </el-button>
            <el-button type="text" style="color: var(--red)" @click="deleteHandle(scope.row.id, 2)">
              <el-icon color="var(--red)"><Delete /></el-icon>
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>
    <!-- 分类弹窗 -->
    <el-dialog
      :title="title"
      v-model="Dialog"
      width="30%"
      :before-close="handleClose"
      :append-to-body="true"
      destroy-on-close
      center
      class="sort-dialog"
    >
      <div class="sort-dialog-content">
        <template v-if="flag === '分类'">
          <div class="myCenter">
            <el-radio-group v-model="sortForHttp.sortType">
              <el-radio-button :label="0">导航栏分类</el-radio-button>
              <el-radio-button :label="1">普通分类</el-radio-button>
            </el-radio-group>
          </div>
          <el-input placeholder="请输入分类名称" v-model="sortForHttp.sortName">
            <template v-slot:prepend>分类名称</template>
          </el-input>
          <el-input placeholder="请输入分类描述" v-model="sortForHttp.sortDescription">
            <template v-slot:prepend>分类描述</template>
          </el-input>
          <el-input
            type="number"
            placeholder="请输入整数，数字小的在前面"
            v-model="sortForHttp.priority"
          >
            <template v-slot:prepend>分类优先级</template>
          </el-input>
        </template>
        <template v-else>
          <el-input placeholder="请输入标签名称" v-model="labelForHttp.labelName">
            <template v-slot:prepend>标签名称</template>
          </el-input>
          <el-input placeholder="请输入标签描述" v-model="labelForHttp.labelDescription">
            <template v-slot:prepend>标签描述</template>
          </el-input>
        </template>
      </div>

      <template v-slot:footer>
        <span class="dialog-footer">
          <el-button @click="handleClose()">取 消</el-button>
          <el-button type="primary" @click="saveEdit(flag)">确 定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { notifySuccess, notifyError } from "@/utils/notify";
import api from "@/api";
import { ref, onMounted } from "vue";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
import type {
  SortInfo,
  LabelInfo,
  SaveSortParams,
  UpdateSortParams,
  SaveLabelParams,
  UpdateLabelParams,
} from "@/types/sort";
import { SortType } from "@/types/sort";

const { $common, $confirm } = useGlobalProperties();

interface SortForHttp {
  id: number | null;
  sortName: string;
  sortDescription: string;
  sortType: SortType | null;
  priority: number | null;
}

interface LabelForHttp {
  id: number | null;
  sortId: number | null;
  labelName: string;
  labelDescription: string;
}

const Dialog = ref<boolean>(false);
const sortInfo = ref<SortInfo[]>([]);
const sort = ref<SortInfo>({} as SortInfo);
const sortForHttp = ref<SortForHttp>({
  id: null,
  sortName: "",
  sortDescription: "",
  sortType: null,
  priority: null,
});
const labelForHttp = ref<LabelForHttp>({
  id: null,
  sortId: null,
  labelName: "",
  labelDescription: "",
});
// 控制打开对话框类型
const flag = ref<string>("");
// 对话框标题
const title = ref<string>("");
// 控制标签弹窗
const dialogTableVisible = ref<boolean>(false);

onMounted(() => {
  getSortInfo();
});

const deleteHandle = (id: number, flagValue: number): void => {
  if (flagValue === 1) {
    title.value = "分类";
  } else if (flagValue === 2) {
    title.value = "标签";
  } else {
    return;
  }
  $confirm(
    `确认要删除这个${title.value}吗？删除${title.value}后请编辑文章，修改文章的${title.value}ID！！`,
    "提示",
    {
      confirmButtonText: "确定",
      cancelButtonText: "取消",
      type: "warning",
      center: true,
    }
  )
    .then(async () => {
      try {
        if (flagValue === 1) {
          await api.deleteSort(id);
        } else if (flagValue === 2) {
          await api.deleteLabel(id);
        }
        notifySuccess("删除成功！");
        await getSortInfo();
        sort.value = {} as SortInfo;
        dialogTableVisible.value = false;
      } catch {
        dialogTableVisible.value = false;
      }
    })
    .catch(() => {
      notifySuccess("已取消删除！");
      dialogTableVisible.value = false;
    });
};
// 保存编辑
const saveEdit = async (flagValue: string): Promise<void> => {
  if (flagValue === "分类") {
    if (
      $common.isEmpty(sortForHttp.value.sortType) ||
      $common.isEmpty(sortForHttp.value.priority) ||
      $common.isEmpty(sortForHttp.value.sortName) ||
      $common.isEmpty(sortForHttp.value.sortDescription)
    ) {
      notifyError("请完善所有分类信息！");
      return;
    }
    try {
      if ($common.isEmpty(sortForHttp.value.id)) {
        await api.saveSort(sortForHttp.value as SaveSortParams);
      } else {
        await api.updateSort(sortForHttp.value as UpdateSortParams);
      }
      notifySuccess("保存成功！");
      await getSortInfo();
      handleClose();
    } catch {
      /* ignored */
    }
  } else {
    if (
      $common.isEmpty(labelForHttp.value.labelName) ||
      $common.isEmpty(labelForHttp.value.labelDescription)
    ) {
      notifyError("请完善所有标签信息！");
      return;
    }
    try {
      if ($common.isEmpty(labelForHttp.value.id)) {
        await api.saveLabel(labelForHttp.value as SaveLabelParams);
      } else {
        await api.updateLabel(labelForHttp.value as UpdateLabelParams);
      }
      notifySuccess("保存成功！");
      await getSortInfo();
      handleClose();
      sort.value = {} as SortInfo;
    } catch {
      /* ignored */
    }
  }
  dialogTableVisible.value = false;
};

// 打开新增分类对话框
const openSortDialog = (): void => {
  Dialog.value = true;
  // 默认选中导航栏分类
  sortForHttp.value.sortType = 0;
  flag.value = "分类";
  title.value = "新增分类";
};

// 表格中编辑分类和标签事件
const editSortAndLabel = (SortOrLabel: SortInfo | LabelInfo, flagValue: string): void => {
  if (flagValue === "编辑分类") {
    title.value = "编辑分类";
    flag.value = "分类";
    const sortItem = SortOrLabel as SortInfo;
    sortForHttp.value.id = sortItem.id;
    sortForHttp.value.sortName = sortItem.sortName;
    sortForHttp.value.sortDescription = sortItem.sortDescription || "";
    sortForHttp.value.sortType = sortItem.sortType;
    sortForHttp.value.priority = sortItem.priority;
  } else {
    title.value = "编辑标签";
    flag.value = "编辑标签";
    const labelItem = SortOrLabel as LabelInfo;
    labelForHttp.value.id = labelItem.id;
    labelForHttp.value.sortId = labelItem.sortId;
    labelForHttp.value.labelName = labelItem.labelName;
    labelForHttp.value.labelDescription = labelItem.labelDescription || "";
  }
  Dialog.value = true;
};

// 新增标签
const insertLabel = (sortItem: SortInfo): void => {
  flag.value = "新增标签";
  title.value = "新增标签";
  labelForHttp.value.sortId = sortItem.id;
  Dialog.value = true;
};

// 关闭对话框
const handleClose = (): void => {
  labelForHttp.value = {
    id: null,
    sortId: null,
    labelName: "",
    labelDescription: "",
  };
  sortForHttp.value = {
    id: null,
    sortName: "",
    sortDescription: "",
    sortType: null,
    priority: null,
  };
  Dialog.value = false;
};

// 查看标签
const seeLabel = (sortItem: SortInfo): void => {
  dialogTableVisible.value = true;
  flag.value = "查看标签";
  sort.value = sortItem;
};

// 得到分类信息
const getSortInfo = async (): Promise<void> => {
  try {
    const res = await api.getSortInfo();
    sortInfo.value = (res || []) as SortInfo[];
  } catch {
    /* ignored */
  }
};
</script>

<style lang="scss">
.sort-dialog {
  .sort-dialog-content {
    > div {
      margin: 12px 0;
    }
  }
}
</style>
<style lang="scss" scoped>
.sort-list-container {
  .sort-list-header-button-wrap {
    margin-bottom: 20px;
  }

  .sort-list-header-title {
    margin-bottom: 10px;
    text-align: center;
  }
}
</style>
