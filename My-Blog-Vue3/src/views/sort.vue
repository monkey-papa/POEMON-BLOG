<template>
  <div class="sort-container">
    <div
      style="animation: header-effect 2s"
      :style="{ background: `${changeBgState}` }"
      class="background-image background-image-changeBg"
    ></div>
    <!-- 两句诗 -->
    <div class="my-animation-slide-top">
      <twoPoem />
    </div>
    <div class="sort-content my-animation-slide-bottom">
      <!-- 标签 -->
      <div
        class="sort-warp shadow-box"
        v-if="sort && !$common.isEmpty(sort) && !$common.isEmpty(sort.labels)"
      >
        <div
          v-for="(label, index) in sort?.labels || []"
          :key="index"
          :class="{
            isActive: !$common.isEmpty(labelId) && labelId !== null && labelId === label.id,
          }"
          @click="toggleLabel(label)"
        >
          <proTag
            :info="label.labelName + ' ' + label.countOfLabel"
            :color="$constant.tree_hole_color[Math.floor(Math.random() * 6)]"
          />
        </div>
      </div>

      <!-- 文章 -->
      <div class="article-wrap">
        <articleList :articleList="articles" />
        <div class="pagination-wrap">
          <div
            @click="pageArticles()"
            class="pagination"
            v-if="pagination.total !== articles.length"
          >
            下一页
          </div>
          <div class="pagination-text" v-else>~~( •̀ ω •́ )y 到底啦~~</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import api from "@/api";
import { ref, onMounted, nextTick, defineAsyncComponent, computed } from "vue";
import { useStore } from "@/stores";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
import type { SortInfo } from "@/types/sort";
import type { ArticleListItem } from "@/types/article";
import type { GetArticleListParams } from "@/types/article";
import { useRouter } from "vue-router";

const store = useStore();
const { $common, $constant, $route } = useGlobalProperties();
const router = useRouter();

const changeBgState = computed<string>(() => store.changeBg);
const sortInfo = computed<SortInfo[]>(() => store.sortInfo);

const twoPoem = defineAsyncComponent(() => import("./common/twoPoem.vue"));
const proTag = defineAsyncComponent(() => import("./common/proTag.vue"));
const articleList = defineAsyncComponent(() => import("./articleList.vue"));

const sortId = ref<number | null>(
  $route.query.sortId ? parseInt($route.query.sortId as string) : null
);
const labelId = ref<number | null>(
  $route.query.labelId ? parseInt($route.query.labelId as string) : null
);
const sort = ref<SortInfo | null>(null);
const pagination = ref<GetArticleListParams & { total: number }>({
  current: 1,
  size: 10,
  total: 0,
  searchKey: "",
  sortId: $route.query.sortId ? parseInt($route.query.sortId as string) : undefined,
  labelId: $route.query.labelId ? parseInt($route.query.labelId as string) : undefined,
});
const articles = ref<ArticleListItem[]>([]);

onMounted(() => {
  getSort();
  getArticles();
});

const pageArticles = (): void => {
  pagination.value.current = pagination.value.current + 1;
  getArticles();
};

const getSort = (): void => {
  const sortInfoData = sortInfo.value;
  if (!$common.isEmpty(sortInfoData)) {
    const sortArray = sortInfoData.filter((f) => {
      return f.id === sortId.value;
    });
    if (!$common.isEmpty(sortArray)) {
      sort.value = sortArray[0];
    }
  }
};

const toggleLabel = (label: { id: number }): void => {
  if (labelId.value === label.id) {
    return;
  }
  articles.value = [];
  router.push({
    path: "/sort",
    query: { sortId: sortId.value, labelId: label.id },
  });
  labelId.value = label.id;
  pagination.value = {
    current: 1,
    size: 10,
    total: 0,
    searchKey: "",
    sortId: $route.query.sortId ? parseInt($route.query.sortId as string) : undefined,
    labelId: label.id,
  };
  nextTick(() => {
    getArticles();
  });
};

const getArticles = async (): Promise<void> => {
  try {
    const res = await api.getArticleList(pagination.value);
    articles.value = articles.value.concat(res.list || []);
    pagination.value.total = res.totalCount || 0;
  } catch {
    /* ignored */
  }
};
</script>

<style lang="scss" scoped>
.sort-container {
  .sort-content {
    background: var(--background);
    padding-top: 40px;

    .sort-warp {
      width: 70%;
      max-width: 780px;
      margin: 0 auto;
      padding: 20px;
      border-radius: 10px;
      display: flex;
      flex-wrap: wrap;

      > div {
        &.isActive {
          animation: scale 1.5s ease-in-out infinite;
        }
      }
    }

    .article-wrap {
      width: 70%;
      margin: 40px auto 0;
      min-height: 600px;

      .pagination-wrap {
        display: flex;
        justify-content: center;
        margin-top: 40px;

        .pagination {
          padding: 13px 15px;
          border: 1px solid var(--red2);
          border-radius: 3rem;
          color: var(--red1);
          width: 100px;
          user-select: none;
          text-align: center;

          .pagination-text {
            user-select: none;
            color: var(--red);
          }

          &:hover {
            border: 1px solid var(--blue2);
            color: var(--orange2);
            box-shadow: 0 0 5px var(--blue2);
          }
        }
      }
    }
  }
}

@media screen and (max-width: 900px) {
  .sort-container {
    .sort-content {
      .sort-warp {
        width: 90%;
      }

      .article-wrap {
        width: 90%;
      }
    }
  }
}
</style>
