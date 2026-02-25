<template>
  <div class="tags-container">
    <!-- 图片 -->
    <div
      style="animation: header-effect 2s"
      :style="{ background: `${changeBgState}` }"
      class="background-image background-image-changeBg"
    ></div>
    <!-- 标签区域 -->
    <div class="main">
      <div class="layout">
        <div id="tag">
          <!-- 标签栏 -->
          <div id="catalog-bar">
            <i class="fa fa-tags"></i>
            <div id="catalog-list">
              <div
                v-for="(label, i) in labelInfo"
                :key="i"
                @click="toggleLabel(label)"
                class="catalog-list-item"
                :class="{
                  selected: labelId && labelId === label.id,
                }"
              >
                <a href="javascript:;"
                  >{{ label.labelName }}<sup>{{ label.countOfLabel }}</sup></a
                >
              </div>
            </div>
          </div>
          <h1 class="article-sort-title">标签 - {{ curLabelName }}</h1>
          <!-- 内容 -->
          <div v-if="!$common.isEmpty(articles)" class="article-sort">
            <div class="article-sort-item shadow-box" v-for="(item, i) in articles" :key="i">
              <a
                @click="
                  router.push({
                    path: '/article',
                    query: { id: item.id },
                  })
                "
                href="javascript:;"
                class="article-sort-item-img"
              >
                <el-image class="img" lazy :src="item.articleCover">
                  <!-- 懒加载图片 -->
                  <template v-slot:placeholder>
                    <div>
                      <div>
                        <img class="img-loading" src="../assets/file/lazy.gif" />
                      </div>
                    </div>
                  </template>
                  <!-- 错误图片 -->
                  <template v-slot:error>
                    <div class="image-slot myCenter">
                      <img class="error-img" src="../assets/file/errorBG.png" />
                      <div class="error-text">
                        <div class="error-text-content">ya，图片跑丢了┮﹏┭</div>
                      </div>
                    </div>
                  </template>
                </el-image>
              </a>
              <div class="article-sort-item-info">
                <div class="article-sort-item-time">
                  <time class="post-meta-date-created">{{ formatTime(item.createTime) }}</time>
                </div>
                <a
                  href="javascript:;"
                  @click="
                    router.push({
                      path: '/article',
                      query: { id: item.id },
                    })
                  "
                  class="article-sort-item-title"
                  >{{ item.articleTitle }}<sup class="article-sort-item-index">{{ i + 1 }}</sup></a
                >
                <div class="article-meta-wrap">
                  <span
                    @click.stop="
                      router.push({
                        path: '/sort',
                        query: { sortId: item.sortId },
                      })
                    "
                    class="article-sort-item-categories"
                  >
                    <img src="../assets/svg/sort2.svg" />
                    {{ getArticleWithRelations(item).sort?.[0]?.sortName || "" }}</span
                  >
                  <span
                    @click.stop="
                      router.push({
                        path: '/sort',
                        query: { sortId: item.sortId, labelId: labelId },
                      })
                    "
                    class="article-sort-item-tags"
                  >
                    <img src="../assets/svg/tag2.svg" />
                    {{ getArticleWithRelations(item).label?.[0]?.labelName || "" }}</span
                  >
                </div>
              </div>
            </div>
          </div>
          <div v-else class="no-article-text myCenter">
            <h3>主人偷懒了，这个标签还没有文章噢~</h3>
          </div>
          <!-- 分页器 -->
          <div class="pagination">
            <el-pagination
              background
              layout="prev, pager, next"
              :current-page="pagination.current"
              :page-size="pagination.size"
              :total="pagination.total"
              @current-change="handlePageChange"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import api from "@/api";
import { ref, computed, onMounted, nextTick } from "vue";
import { useStore } from "@/stores";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
import { useRouter } from "vue-router";
import type { LabelInfo } from "@/types/sort";
import type { ArticleListItem } from "@/types/article";
import type { GetArticleListParams } from "@/types/article";

type ArticleListItemWithRelations = ArticleListItem;

const store = useStore();
const router = useRouter();
const { $common, $route } = useGlobalProperties();

const changeBgState = computed<string>(() => store.changeBg);
const labelInfo = computed<LabelInfo[]>(() => store.labelInfo);

const labelId = ref<number | null>(
  $route.query.labelId ? parseInt($route.query.labelId as string) : null
);
const pagination = ref<GetArticleListParams & { total: number }>({
  current: 1,
  size: 10,
  total: 0,
  searchKey: "",
  sortId: undefined,
  labelId: $route.query.labelId ? parseInt($route.query.labelId as string) : undefined,
});
const articles = ref<ArticleListItemWithRelations[]>([]);

const getArticleWithRelations = (article: ArticleListItem): ArticleListItemWithRelations => {
  return article as ArticleListItemWithRelations;
};

const curLabelName = computed<string>(() => {
  const labelList = labelInfo.value;
  const curLabel = labelList.filter((item) => {
    return item.id === labelId.value;
  });
  return curLabel[0]?.labelName || "";
});

onMounted(() => {
  scrollTo();
  getArticles();
});

const formatTime = (row: string): string => {
  const day = row.split(".")[0].split("T")[0];
  const time = row.split(".")[0].split("T")[1];
  return `${day} 日 ${time}`;
};

const handlePageChange = (val: number): void => {
  pagination.value.current = val;
  getArticles();
};

const scrollTo = (): void => {
  const content = document.querySelector("#catalog-list");
  if (content) {
    // 需要阻止默认行为，所以使用 passive: false
    content.addEventListener(
      "wheel",
      (event: Event) => {
        const wheelEvent = event as WheelEvent;
        wheelEvent.preventDefault();
        (content as HTMLElement).scrollLeft += wheelEvent.deltaY;
      },
      { passive: false }
    );
  }
};

const toggleLabel = (label: LabelInfo): void => {
  if (labelId.value === label.id) {
    return;
  }
  articles.value = [];
  router.push({ path: "/tags", query: { labelId: label.id } });
  labelId.value = label.id;
  pagination.value = {
    current: 1,
    size: 10,
    total: 0,
    searchKey: "",
    sortId: undefined,
    labelId: label.id,
  };
  nextTick(() => {
    getArticles();
  });
};

const getArticles = async (): Promise<void> => {
  try {
    const res = await api.getArticleList(pagination.value);
    articles.value = res.list || [];
    pagination.value.total = res.totalCount || 0;
  } catch {
    /* ignored */
  }
};
</script>
<style lang="scss" scoped>
.tags-container {
  .main {
    margin: 0 auto;
    padding: 45px 0 0;

    .layout {
      display: flex;

      #tag {
        padding: 10px 35px;
        display: flex;
        flex-direction: column;
        min-height: calc(100vh - 50px);
        width: 100%;
        border: 2px dashed var(--blue9);
        font-weight: 700;
        background: var(--background);

        #catalog-bar {
          padding: 0.4rem 0.8rem;
          border-radius: 0.5rem;
          display: flex;
          border: 2px dashed var(--blue);
          justify-content: space-between;

          .fa-tags {
            color: var(--purple1);
            margin-right: 4px;
          }

          #catalog-list {
            display: flex;
            white-space: nowrap;
            overflow-x: scroll;
            overflow-y: hidden;

            .catalog-list-item {
              a {
                font-size: 16px;
                margin: 0 0.3em;
                padding: 0.5em 0.4em 0.5em;
                font-weight: 400;
                border-radius: 0.5rem;
                color: var(--red1);
                line-height: 1.5rem;
              }

              &:hover a {
                transition: all 0.3s;
                border-radius: 0.5rem;
                background-color: var(--blue25);
                color: var(--white);
              }

              &.selected a {
                background: var(--green2);
                color: var(--white);
              }
            }
          }
        }

        .article-sort-title {
          text-indent: 1em;
          font-weight: 400;
          color: var(--darkBlue);
        }

        .article-sort {
          padding-left: 5px;
          margin-top: 7.5px;
          display: flex;
          flex-direction: row;
          flex-wrap: wrap;
          justify-content: space-between;

          &-item {
            position: relative;
            display: flex;
            align-items: center;
            margin: 12px 0;
            border-radius: 12px;
            padding: 8px;
            width: calc(50% - 7.5px);
            background: var(--myAsideColor);
            border: 1px solid var(--gray1);

            &:hover {
              border-color: var(--gray4);
            }

            &-img {
              overflow: hidden;
              width: 160px;
              height: 90px;
              border-radius: 7px;
              border: 1px solid var(--blue);

              .img {
                transition: all 0.6s;
                width: 100%;
                height: 100%;
                object-fit: cover;
                border-radius: 7px;

                .img-loading {
                  object-fit: cover;
                  width: 100%;
                  height: 100%;
                }

                &:hover {
                  transform: scale(1.2);
                }
              }

              .image-slot {
                .error-img {
                  position: relative;
                  z-index: 1;
                  object-fit: cover;
                  width: 100%;
                  height: 100%;
                }

                .error-text {
                  z-index: 2;
                  position: absolute;
                  font-size: 16px;
                  line-height: 1.8;
                  letter-spacing: 8px;

                  .error-text-content {
                    color: var(--wheat1);
                  }
                }
              }
            }

            &-info {
              -webkit-box-flex: 1;
              -moz-box-flex: 1;
              -o-box-flex: 1;
              box-flex: 1;
              -webkit-flex: 1;
              -ms-flex: 1;
              flex: 1;
              padding-left: 10px;

              .article-sort-item-time {
                .post-meta-date-created {
                  bottom: 5px;
                  right: 10px;
                  position: absolute;
                  font-size: 14px;
                  line-height: 14px;
                  color: var(--darkBlue);
                }
              }

              .article-meta-wrap {
                margin: 10px;
              }
            }

            &-title {
              width: 90%;
              color: var(--bigRed1);
              font-size: 17px;
              transition: all 0.3s;
              font-weight: 500;
              line-height: 1.4em;
              display: -webkit-box;
              -webkit-box-orient: vertical;
              white-space: wrap;
              -webkit-line-clamp: 2;
              overflow: hidden;
              text-overflow: ellipsis;

              &:hover {
                color: var(--green2);
                padding: 0px 15px;
              }

              .article-sort-item-index {
                opacity: 0.8;
                position: absolute;
                top: 0.2rem;
                right: 0.5rem;
                font-style: italic;
                font-size: 2rem;
                line-height: 1.5rem;
                color: var(--red);
              }
            }

            &-categories,
            &-tags {
              margin: 5px;
              color: var(--purple4);
              font-size: 14px;
              font-weight: 400;
              display: -webkit-box;
              -webkit-box-orient: vertical;
              white-space: wrap;
              -webkit-line-clamp: 1;
              overflow: hidden;
              text-overflow: ellipsis;

              img {
                vertical-align: -3px;
              }

              &:hover {
                transition: all 0.4s;
                color: var(--darkBlue1);
              }
            }
          }
        }

        .no-article-text {
          flex: 1;
          color: var(--darkBlue);
        }

        .pagination {
          margin: 20px 0;
          text-align: center;

          :deep(.el-pagination.is-background) {
            display: flex;
            justify-content: center;

            .el-pager {
              li {
                &:hover {
                  color: var(--white);
                  background: var(--blue25);
                }

                &:not(.disabled).is-active {
                  background: var(--green2);
                  color: var(--white);
                }
              }
            }
          }
        }
      }
    }
  }
}

@media screen and (max-width: 1286px) {
  .tags-container {
    .main {
      padding: 60px 3px;
    }
  }
}

@media screen and (max-width: 974px) {
  .tags-container {
    .main {
      .layout {
        #tag {
          .article-sort {
            &-item {
              width: 100%;
            }
          }
        }
      }
    }
  }
}

@media screen and (max-width: 523px) {
  .tags-container {
    .main {
      .layout {
        #tag {
          padding: 10px 8px;

          .article-sort {
            &-item {
              margin: 6px 0;

              &-img {
                width: 80px;
                height: 76px;
              }

              &-title {
                width: 218px;
                font-size: 16px;
              }
            }
          }
        }
      }
    }
  }
}
</style>
