<template>
  <div class="user-list-container">
    <div class="user-list-content">
      <div class="header-search">
        <el-select
          @blur="closeOptions('userType')"
          @clear="handleClear('userType')"
          ref="fuzzyUserType"
          v-model="pagination.userType"
          placeholder="用户类型"
          class="mr10"
          clearable
        >
          <el-option key="1" label="Boss" :value="0" />
          <el-option key="2" label="管理员" :value="1" />
          <el-option key="4" label="访客" :value="3" />
          <el-option key="3" label="普通用户" :value="2" />
        </el-select>
        <el-select
          @blur="closeOptions('userStatus')"
          @clear="handleClear('userStatus')"
          ref="fuzzyUserStatus"
          v-model="pagination.userStatus"
          placeholder="用户状态"
          class="mr10"
          clearable
        >
          <el-option key="1" label="启用" :value="true" />
          <el-option key="2" label="禁用" :value="false" />
        </el-select>
        <el-input
          v-model="pagination.searchKey"
          placeholder="手机号"
          type="number"
          class="mr10"
          clearable
        />
        <el-button type="primary" @click="searchUser()"
          ><el-icon color="white"><Search /></el-icon>搜索</el-button
        >
        <el-button type="danger" @click="clearSearch()">清除参数</el-button>
      </div>
      <div class="user-list-warning">警告：Boss用户只能有1个，其他用户可以任意个</div>
      <el-table :data="users" border class="admin-table" header-cell-class-name="table-header">
        <el-table-column prop="id" label="ID" width="55" align="center" />
        <el-table-column prop="username" label="用户名" align="center" />
        <el-table-column prop="phoneNumber" label="手机号" align="center" />
        <el-table-column prop="email" label="邮箱" align="center" />
        <el-table-column prop="province" label="注册地址" align="center" />
        <el-table-column label="赞赏" width="100" align="center">
          <template v-slot="scope">
            <el-input
              size="small"
              maxlength="30"
              v-model="scope.row.admire"
              @blur="changeUserAdmire(scope.row)"
            />
          </template>
        </el-table-column>
        <el-table-column label="用户状态" align="center">
          <template v-slot="scope">
            <el-tag
              :type="scope.row.userStatus === false ? 'danger' : 'success'"
              disable-transitions
            >
              {{ scope.row.userStatus === false ? "禁用" : "启用" }}
            </el-tag>
            <el-switch
              class="admin-switch"
              v-hasPermi="['user:visit:read']"
              @click="changeUserStatus(scope.row)"
              v-model="scope.row.userStatus"
            />
          </template>
        </el-table-column>
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
        <el-table-column label="性别" align="center">
          <template v-slot="scope">
            <el-tag
              type="success"
              v-if="scope.row.gender === 1"
              :class="{ boy: scope.row.gender === 1 }"
              disable-transitions
            >
              男生
            </el-tag>
            <el-tag
              type="success"
              v-else-if="scope.row.gender === 2"
              :class="{ girl: scope.row.gender === 2 }"
              disable-transitions
            >
              女生
            </el-tag>
            <el-tag type="success" v-else disable-transitions> 保密 </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="introduction" label="简介" align="center" />
        <el-table-column label="用户类型" width="100" align="center">
          <template v-slot="scope">
            <el-tag
              v-hasPermi="['user:visit:read']"
              type="success"
              v-if="scope.row.userType === 0"
              :class="{ boss: scope.row.userType === 0 }"
              @click="editUser(scope.row)"
              disable-transitions
            >
              Boss
            </el-tag>
            <el-tag
              v-hasPermi="['user:visit:read']"
              type="success"
              v-else-if="scope.row.userType === 1"
              :class="{ manager: scope.row.userType === 1 }"
              @click="editUser(scope.row)"
              disable-transitions
            >
              管理员
            </el-tag>
            <el-tag
              v-hasPermi="['user:visit:read']"
              type="success"
              v-else-if="scope.row.userType === 3"
              :class="{ visit: scope.row.userType === 3 }"
              @click="editUser(scope.row)"
              disable-transitions
            >
              访客
            </el-tag>
            <el-tag
              v-hasPermi="['user:visit:read']"
              type="success"
              v-else
              @click="editUser(scope.row)"
              disable-transitions
            >
              普通用户
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          :formatter="$common?.formatter"
          prop="createTime"
          label="注册时间"
          align="center"
        />
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

    <!-- 编辑弹出框 -->
    <el-dialog
      title="修改用户类型"
      v-model="editVisible"
      width="30%"
      :before-close="handleClose"
      :append-to-body="true"
      destroy-on-close
      center
    >
      <div class="myCenter">
        <el-radio-group v-model="changeUser.userType">
          <el-radio-button :label="0">Boss</el-radio-button>
          <el-radio-button :label="1">管理员</el-radio-button>
          <el-radio-button :label="3">访客</el-radio-button>
          <el-radio-button :label="2">普通用户</el-radio-button>
        </el-radio-group>
      </div>
      <template v-slot:footer>
        <span class="dialog-footer">
          <el-button @click="handleClose()">取 消</el-button>
          <el-button type="primary" @click="saveEdit()">确 定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { notifySuccess, notifyError } from "@/utils/notify";
import api from "@/api";
import { ref, onMounted, computed, type Ref, type ComputedRef } from "vue";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
import { useStore } from "@/stores";
import { UserType } from "@/types";
import type { AdminUserListItem } from "@/types";
import type { ElSelect } from "element-plus";

const store = useStore();
const { $common, $confirm } = useGlobalProperties();

const webInfo: ComputedRef = computed(() => store.webInfo);

interface Pagination {
  current: number;
  size: number;
  total: number;
  searchKey: string;
  userStatus: boolean | null;
  userType: UserType | null;
}

interface ChangeUser {
  userId: number;
  userType: UserType;
}

interface UserListItemWithStatus extends AdminUserListItem {}

const pagination: Ref<Pagination> = ref({
  current: 1,
  size: 10,
  total: 0,
  searchKey: "",
  userStatus: null,
  userType: null,
});
const users: Ref<UserListItemWithStatus[]> = ref([]);
const changeUser: Ref<ChangeUser> = ref({
  userId: 0,
  userType: UserType.User,
});
const editVisible: Ref<boolean> = ref(false);
const fuzzyUserType: Ref<InstanceType<typeof ElSelect> | null> = ref(null);
const fuzzyUserStatus: Ref<InstanceType<typeof ElSelect> | null> = ref(null);

onMounted(() => {
  getUsers();
});

const handleClear = (field: keyof Pagination): void => {
  pagination.value[field] = null as never;
  getUsers();
};

const clearSearch = (): void => {
  pagination.value = {
    current: 1,
    size: 10,
    total: 0,
    searchKey: "",
    userStatus: null,
    userType: null,
  };
  getUsers();
};

const getUsers = async (): Promise<void> => {
  try {
    const res = await api.getAdminUserList(pagination.value);
    users.value = (res.list || []) as UserListItemWithStatus[];
    pagination.value.total = res.totalCount || 0;
  } catch {
    /* ignored */
  }
};

const changeUserStatus = async (user: UserListItemWithStatus): Promise<void> => {
  try {
    await api.changeUserStatus({
      userId: user.id,
      status: user.userStatus,
    });
    notifySuccess("修改成功！");
  } catch {
    /* ignored */
  }
};

const changeUserAdmire = (user: UserListItemWithStatus): void => {
  if (!$common.isEmpty(user.admire)) {
    $confirm("确认保存？", "提示", {
      confirmButtonText: "确定",
      cancelButtonText: "取消",
      type: "success",
      center: true,
    })
      .then(async () => {
        try {
          await api.changeUserAdmire({
            userId: user.id,
            admire: user.admire || "",
          });
          notifySuccess("修改成功！");
        } catch {
          /* ignored */
        }
      })
      .catch(() => {
        notifySuccess("已取消保存！");
      });
  }
};

const editUser = (user: UserListItemWithStatus): void => {
  changeUser.value.userId = user.id;
  changeUser.value.userType = user.userType;
  if (changeUser.value.userType === 0) {
    notifyError("禁止修改Boss用户");
    return;
  }
  editVisible.value = true;
};

const handlePageChange = (val: number): void => {
  pagination.value.current = val;
  getUsers();
};

const searchUser = (): void => {
  pagination.value.total = 0;
  pagination.value.current = 1;
  getUsers();
};

const handleClose = (): void => {
  changeUser.value = {
    userId: 0,
    userType: UserType.User,
  };
  editVisible.value = false;
};

const saveEdit = async (): Promise<void> => {
  try {
    await api.changeUserType(changeUser.value);
    handleClose();
    await getUsers();
    notifySuccess("修改成功！");
  } catch {
    /* ignored */
  }
};

const closeOptions = (item: "userType" | "userStatus"): void => {
  if (item === "userType") {
    fuzzyUserType.value?.blur();
  } else {
    fuzzyUserStatus.value?.blur();
  }
};
</script>

<style lang="scss" scoped>
.user-list-container {
  .user-list-content {
    .user-list-warning {
      color: var(--red);
    }

    .admin-table {
      .boy,
      .manager {
        color: var(--blue25);
      }

      .visit {
        color: var(--orange);
      }

      .girl,
      .boss {
        color: var(--red);
      }
    }
  }
}
</style>
