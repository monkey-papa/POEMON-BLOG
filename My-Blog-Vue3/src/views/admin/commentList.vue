<template>
  <div>
    <div class="header-search">
      <el-select
        @blur="closeOptions"
        @clear="handleClear('commentType')"
        ref="fuzzyCommentType"
        v-model="pagination.commentType"
        placeholder="评论来源类型"
        class="mr10"
        clearable
      >
        <el-option key="1" label="文章评论" value="article" />
        <el-option key="2" label="树洞评论" value="message" />
        <el-option key="3" label="恋爱留言" value="love" />
      </el-select>
      <el-input
        type="number"
        class="mr10"
        v-model="pagination.source"
        placeholder="评论来源标识"
        clearable
      />
      <el-button type="primary" @click="searchComments()"
        ><el-icon color="white"><Search /></el-icon>搜索
      </el-button>
      <el-button type="danger" @click="clearSearch()">清除参数</el-button>
    </div>
    <el-table :data="comments" border class="admin-table" header-cell-class-name="table-header">
      <el-table-column prop="id" label="ID" width="55" align="center" />
      <el-table-column prop="username" label="用户名称" align="center" />
      <el-table-column label="头像" align="center">
        <template v-slot="scope">
          <el-image
            lazy
            class="admin-image-thumb"
            :src="scope.row.avatar || webInfo.avatar"
            fit="cover"
          />
        </template>
      </el-table-column>
      <el-table-column prop="source" label="评论来源标识" align="center" />
      <el-table-column prop="type" label="评论来源类型" align="center" />
      <el-table-column prop="likeCount" label="点赞数" align="center" />
      <el-table-column
        width="200"
        prop="commentContent"
        label="评论内容"
        align="center"
        show-overflow-tooltip
      >
        <template v-slot="scope">
          <span v-html="scope.row.commentContent"></span>
        </template>
      </el-table-column>
      <el-table-column prop="commentInfo" label="评论额外信息" align="center" />
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
        layout="total,sizes, prev, pager, next"
        :current-page="pagination.current"
        :page-size="pagination.size"
        :total="pagination.total"
        :page-sizes="[10, 20, 50, 100]"
        @current-change="handlePageChange"
        @size-change="handleSizeChange"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { notifySuccess } from "@/utils/notify";
import api from "@/api";
import { ref, onMounted, computed } from "vue";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
import { useStore } from "@/stores";
import type { CommentListItem } from "@/types/comment";
import type { SelectInstance } from "element-plus";

const store = useStore();
const { $common, $confirm } = useGlobalProperties();

interface Pagination {
  current: number;
  size: number;
  total: number;
  source: number | null;
  commentType: string;
  type?: string;
}

const webInfo = computed(() => store.webInfo);

const pagination = ref<Pagination>({
  current: 1,
  size: 10,
  total: 0,
  source: null,
  commentType: "",
});

const comments = ref<CommentListItem[]>([]);
const fuzzyCommentType = ref<SelectInstance | null>(null);

onMounted(() => {
  getComments();
});

const handleSizeChange = (val: number): void => {
  pagination.value.size = val;
  getComments();
};

const handleClear = (field: keyof Pagination): void => {
  pagination.value[field] = "" as never;
  getComments();
};

const emoji = (commentsData: CommentListItem[]): void => {
  commentsData.forEach((c) => {
    c.commentContent = $common.processContent(c.commentContent);
  });
};

const clearSearch = (): void => {
  pagination.value = {
    current: 1,
    size: 10,
    total: 0,
    source: null,
    commentType: "",
  };
  getComments();
};

const getComments = async (): Promise<void> => {
  try {
    const params = {
      current: pagination.value.current,
      size: pagination.value.size,
      source: pagination.value.source || 0,
      commentType: pagination.value.commentType || "",
    };
    const res = await api.getBossCommentList(params);
    comments.value = res.list || [];
    emoji(comments.value);
    pagination.value.total = res.totalCount || 0;
  } catch {
    /* ignored */
  }
};

const handlePageChange = (val: number): void => {
  pagination.value.current = val;
  getComments();
};

const searchComments = (): void => {
  pagination.value.total = 0;
  pagination.value.current = 1;
  getComments();
};

const doDelete = async (item: CommentListItem): Promise<void> => {
  try {
    await api.bossDeleteComment(item.id);
    pagination.value.current = 1;
    await getComments();
    notifySuccess("删除成功！");
  } catch {
    /* ignored */
  }
};

const handleDelete = (item: CommentListItem): void => {
  $confirm("删除评论后，所有该评论的回复均不可见。确认删除？", "提示", {
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

const addProWords = (item: CommentListItem): void => {
  $confirm("确定要添加为违禁词吗？", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
    center: true,
  }).then(async () => {
    try {
      await api.addProhibitedWord({
        message: item.commentContent,
        username: item.username,
        avatar: item.avatar,
      });
      // 添加成功后删除评论
      await doDelete(item);
    } catch {
      // apiWrapper 已统一处理错误提示
    }
  });
};

const closeOptions = (): void => {
  fuzzyCommentType.value?.blur();
};
</script>
