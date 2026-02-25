<template>
  <div class="funny-container">
    <div
      style="animation: header-effect 2s"
      :style="{ background: `${changeBgState}` }"
      class="background-image background-image-changeBg"
    ></div>
    <!-- 两句诗 -->
    <div class="my-animation-slide-top">
      <twoPoem :isShehui="true" />
    </div>
    <div class="funny-content my-animation-slide-bottom">
      <h1>♪(^∇^*) 欢迎来到{{ webInfo.webName }}音乐平台，请尽情欣赏~~</h1>
      <h3>还想听更多音乐吗？请移步<a class="my-music" :href="$constant.siteURL">我的云音乐</a></h3>
      <!-- 音乐 -->
      <div class="funny-wrap" v-if="!$common.isEmpty(funnys)">
        <div class="music-wrap" v-for="(item, index) in funnys" :key="index">
          <div class="music-title">
            <span class="rotate">
              <i class="iconfont icon-fengche"></i>
            </span>
            <span class="funny-title">{{ index + 1 }}号厅：{{ item.classify }}</span>
          </div>
          <div class="process-wrap">
            <el-collapse
              class="funny-collapse shadow-box"
              v-model="activeName"
              accordion
              @change="changeFunny(item.classify)"
            >
              <el-collapse-item
                :title="`🎻 要来${index + 1}号厅吗，为您播放<${item.classify}>歌单`"
                :name="index"
              >
                <div
                  class="funny-item-wrap my-animation-slide-bottom"
                  v-if="!$common.isEmpty(item.data)"
                >
                  <div class="funny-item-wrap-item" v-for="(funny, i) in item.data" :key="i">
                    <el-image
                      class="funny-avatar myCenter"
                      lazy
                      :size="110"
                      @click="playSound(funny.url, item.data, i)"
                      :src="funny.cover"
                    />
                    <div class="funny-item-title">{{ funny.title }}</div>
                  </div>
                </div>
              </el-collapse-item>
            </el-collapse>
            <hr />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import api from "@/api";
import { ref, onMounted, onBeforeUnmount, computed, type Ref, type ComputedRef } from "vue";
import { defineAsyncComponent } from "vue";
import { useStore } from "@/stores";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
import type { ResourcePath, GetResourcePathListParams, ClassifyItem } from "@/types";

interface FunnyItem {
  classify: string;
  count: number | null;
  data: Partial<ResourcePath>[];
}

interface PaginationParams extends GetResourcePathListParams {
  resourceType: string;
  classify: string;
  noPagination: boolean;
}

const store = useStore();
const { $common, $constant } = useGlobalProperties();

const changeBgState: ComputedRef<string> = computed(() => store.changeBg);
const webInfo = computed(() => store.webInfo);

const twoPoem = defineAsyncComponent(() => import("./common/twoPoem.vue"));

// 获取全部资源参数
const pagination: Ref<PaginationParams> = ref({
  resourceType: "funny",
  classify: "",
  noPagination: true,
});

const activeName: Ref<number> = ref(0);
const audio: Ref<HTMLAudioElement | null> = ref(null);
const playList: Ref<ResourcePath[] | null> = ref(null);
const index: Ref<number | null> = ref(null);
const funnys: Ref<FunnyItem[]> = ref([
  {
    classify: "",
    count: null,
    data: [
      {
        id: 0,
        title: "",
        url: "",
        classify: "",
        cover: "",
        type: "",
        createTime: "",
      },
    ],
  },
]);

onMounted(() => {
  getFunny();
});

onBeforeUnmount(() => {
  if (audio.value != null && !audio.value.paused) {
    audio.value.pause();
  }
});

const getFunny = async (): Promise<void> => {
  try {
    const res: ClassifyItem[] = await api.getClassifyList({ type: "funny" });
    funnys.value = (res || []).map((item) => ({
      classify: item.classify,
      count: item.count,
      data: [],
    }));
    if (funnys.value.length > 0) {
      changeFunny(funnys.value[0]?.classify);
    }
  } catch {
    /* ignored */
  }
};

const listFunny = async (): Promise<void> => {
  try {
    const res = await api.getClientResourcePathList(pagination.value);
    const list = res.list || [];
    funnys.value.forEach((funny) => {
      if (funny.classify === pagination.value.classify) {
        funny.data = list;
      }
    });
    pagination.value.classify = "";
  } catch {
    /* ignored */
  }
};

const changeFunny = (classify: string): void => {
  funnys.value.forEach((funny) => {
    if (funny.classify === classify && $common.isEmpty(funny.data)) {
      pagination.value.classify = classify;
      listFunny();
    }
  });
};

const playSound = (
  src: string | null | undefined,
  playListData: Partial<ResourcePath>[],
  indexValue: number
): void => {
  if (!src) return;
  // audio.value.volume = 0.1
  playList.value = playListData as ResourcePath[];
  index.value = indexValue;
  if (audio.value != null) {
    if (audio.value.src === src) {
      if (audio.value.paused) {
        audio.value.play();
      } else {
        audio.value.pause();
      }
    } else {
      audio.value.pause();
      audio.value.src = src;
      audio.value.load();
      audio.value.play();
    }
  } else {
    audio.value = new Audio(src);
    audio.value.play();
    audio.value.onended = () => {
      if (index.value !== null && playList.value !== null) {
        index.value = index.value + 1;
        if (index.value < playList.value.length && audio.value !== null) {
          audio.value.src = playList.value[index.value].url || "";
          audio.value.load();
          setTimeout(() => {
            if (audio.value !== null) {
              audio.value.play();
            }
          }, 3000);
        }
      }
    };
  }
};
</script>

<style lang="scss">
.funny-collapse {
  border: 1px solid var(--gray1);
  border-radius: 10px;
  overflow: hidden;
  transition: all 0.3s ease;

  &:hover {
    border-color: var(--gray4);
  }

  .el-collapse-item__header {
    border-bottom: unset;
    font-size: 16px;
    font-weight: 400;
    background-color: var(--gray5);
    color: var(--fontColor);
    padding: 30px;
    line-height: 20px;
    border-radius: 6px;

    &:hover {
      color: var(--black8);
    }
  }

  .el-collapse-item__wrap {
    background-color: var(--maxMaxWhiteMask);
    border-bottom: unset;
    border-radius: 6px;

    .funny-item-wrap {
      display: flex;
      flex-flow: wrap;
      margin-left: 20px;

      .funny-item-wrap-item {
        width: 150px;

        .funny-avatar {
          margin: 20px;
          transition: all 0.5s;
          user-select: none;
          border-radius: 50%;

          &:hover {
            transform: rotate(360deg);
          }
        }

        .funny-item-title {
          text-align: center;
          margin: 0 10px;
          font-size: 18px;
          font-weight: 400;
          white-space: nowrap;
          text-overflow: ellipsis;
          overflow: hidden;
          color: var(--bigRed1);

          &:hover {
            text-overflow: unset;
            overflow: unset;
            color: var(--bigRed);
          }
        }
      }
    }
  }
}
</style>
<style lang="scss" scoped>
.funny-container {
  .funny-content {
    padding-top: 40px;
    background: var(--background);

    h1 {
      text-align: center;
      color: var(--green3);
    }

    h3 {
      text-align: center;
      color: var(--green3);

      .my-music {
        background: var(--gradientAnimation);
        background-size: 0px 3px;
        transition: background-size 800ms;
        color: var(--red);

        &:hover {
          background-position-x: left;
          background-size: 100% 3px;
        }
      }
    }

    .funny-wrap {
      width: 99%;
      border-radius: 10px;
      max-width: 1600px;
      margin: 0 auto;
      padding: 40px 20px 80px;

      .music-wrap {
        .music-title {
          display: flex;
          align-items: center;

          .rotate {
            height: 16px;
            width: 16px;
            animation: rotate 1s linear infinite;

            .icon-fengche {
              display: flex;
              height: 16px;
              width: 16px;
              color: var(--red);

              &:before {
                height: 16px;
                width: 16px;
              }
            }
          }

          .funny-title {
            color: var(--yellow3);
            font-size: 28px;
            font-weight: 400;
            margin-left: 12px;

            &:hover {
              color: var(--yellow6);
            }
          }
        }

        .process-wrap {
          margin: 20px 0 40px;

          hr {
            position: relative;
            margin: 30px auto 100px;
            border: 2px dashed var(--blue2);
            overflow: visible;

            &:before {
              position: absolute;
              top: -20px;
              left: 5%;
              color: var(--red);
              content: "\e673";
              font-size: 40px;
              line-height: 1;
              transition: all 1s ease-in-out;
              font-family: iconfont;
            }

            &:hover:before {
              left: calc(95% - 20px);
            }
          }
        }
      }
    }
  }
}
</style>
