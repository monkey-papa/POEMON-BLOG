<template>
  <div class="tools-container">
    <div
      style="animation: header-effect 2s"
      :style="{ background: `${changeBgState}` }"
      class="background-image background-image-changeBg"
    ></div>
    <!-- 封面 -->
    <div class="tools-header my-animation-slide-top">
      <!-- 顶部视频 -->
      <video
        class="index-video"
        :autoplay="true"
        :muted="true"
        :loop="true"
        src="https://www.wulihub.com.cn/gc/Q5ND2N/1920.js"
      ></video>
      <div class="tools-header-content">
        <!-- 标题 -->
        <div class="tools-header-title">
          <div class="tools-header-title-left">记录</div>
          <div class="tools-header-title-right">百宝箱</div>
        </div>
        <div class="card-container">
          <!-- 收藏夹 -->
          <div @click="changeFavorite(1)" class="card-item shadow-box card-item__contain">
            <div class="tools-image"></div>
            <div class="tools-header-content-right">
              <div class="card-name">生产力</div>
              <div class="card-desc">希望这些工具对你有帮助喔ヾ(≧▽≦*)o</div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- 内容 -->
    <div class="tools-content my-animation-slide-bottom">
      <!-- 收藏夹 -->
      <div v-show="card === 1 && !$common.isEmpty(collects)">
        <div class="collects-wrap" v-for="(value, key) in collects" :key="key">
          <div class="collect-classify">
            {{ key }}
          </div>
          <div class="tools-item-wrap">
            <div
              v-for="(item, index) in value"
              :key="index"
              @click="toUrl(item.url)"
              class="tools-item"
            >
              <div>
                <el-avatar class="tools-item-image" :size="60" :src="item.cover" />
              </div>
              <div class="tools-item-content">
                <div class="tools-item-title">
                  {{ item.title }}
                </div>
                <div class="tools-item-introduction">
                  {{ item.introduction }}
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import api from "@/api";
import { ref, onMounted, computed, type Ref, type ComputedRef } from "vue";
import { useStore } from "@/stores";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
import type { CollectItem } from "@/types";

const store = useStore();
const { $common } = useGlobalProperties();

const changeBgState: ComputedRef<string> = computed(() => store.changeBg);

const card: Ref<number | null> = ref(null);
const collects: Ref<Record<string, CollectItem[]>> = ref({});

onMounted(() => {
  card.value = 1;
  getCollect();
});

const toUrl = (url: string | null): void => {
  if (url) window.open(url);
};

const changeFavorite = (cardValue: number): void => {
  card.value = cardValue;
  if (cardValue === 1) {
    if ($common.isEmpty(collects.value)) {
      getCollect();
    }
  }
};

const getCollect = async (): Promise<void> => {
  try {
    const res = await api.getCollectList();
    collects.value = res;
  } catch {
    /* ignored */
  }
};
</script>
<style lang="scss" scoped>
.tools-container {
  background: var(--background);

  .tools-header {
    .tools-header-content {
      position: absolute;
      left: 0;
      top: 0;
      padding: 100px 20px 0px;

      .tools-header-title {
        color: var(--green1);
        margin: 10px;

        &-left {
          margin-left: 10px;
        }

        &-right {
          font-size: 36px;
          font-weight: bold;
          line-height: 2;
        }
      }

      .card-container {
        display: flex;
        flex-wrap: wrap;
        margin-top: 60px;

        .card-item {
          transition: all 0.3s;
          position: relative;
          width: 250px;
          height: 120px;
          border-radius: 20px;
          animation: hideToShow 1s ease-in-out;
          overflow: hidden;
          margin: 10px;
          color: var(--darkBlue);

          &:hover {
            transform: translateY(-6px);
          }

          &.card-item__contain {
            border: 1px solid var(--gray1);

            &:hover {
              color: var(--black8);
              border-color: var(--gray4);
            }
          }

          .tools-image {
            &::before {
              content: "";
              position: absolute;
              width: 100%;
              height: 100%;
              background-color: var(--pink);
            }
          }

          .tools-header-content-right {
            position: absolute;
            left: 0;
            top: 0;
            padding: 20px 25px 15px;

            .card-name {
              font-weight: 400;
              font-size: 25px;

              &:after {
                top: 50px;
                width: 22px;
                left: 26px;
                height: 2px;
                background: var(--white);
                content: "";
                border-radius: 1px;
                position: absolute;
              }
            }

            .card-desc {
              font-weight: bold;
              margin-top: 15px;
            }
          }
        }
      }
    }
  }

  .tools-content {
    margin: 0 auto;
    max-width: 1200px;

    .collects-wrap {
      margin-top: 20px;

      .collect-classify {
        color: var(--bigRed);
        font-size: 28px;
        margin-bottom: 10px;
        padding-left: 10px;
      }

      .tools-item-wrap {
        justify-content: center;
        display: flex;
        flex-wrap: wrap;

        .tools-item {
          color: var(--bigRed1);
          transition: all 0.3s;
          border-radius: 12px;
          box-shadow: 0 8px 16px -4px var(--mask);
          background: var(--pink);
          display: flex;
          width: calc(100% / 4 - 20px);
          max-width: 320px;
          height: 90px;
          overflow: hidden;
          padding: 15px;
          margin: 10px;
          border: 1px solid var(--gray1);
          transition: all 0.3s ease;

          &:hover {
            border-color: var(--gray4);
            background: var(--blue);
            color: var(--white1);

            .tools-item-image {
              transition: all 0.6s;
              width: 0;
              height: 0;
              opacity: 0;
              margin-right: 0;
            }

            div:nth-child(2) {
              width: 100%;
            }
          }

          .tools-item-image {
            margin-right: 20px;
            transition: all 0.3s;
            background: var(--transparent);
          }

          .tools-item-content {
            width: calc(100% - 80px);

            .tools-item-title {
              font-size: 19px;
              font-weight: bold;
              white-space: nowrap;
              text-overflow: ellipsis;
              overflow: hidden;
              margin-bottom: 5px;
            }

            .tools-item-introduction {
              opacity: 0.7;
              font-weight: bold;
              letter-spacing: 1px;
              font-size: 14px;
              line-height: 1.2;
              overflow: hidden;
              text-overflow: ellipsis;
              display: -webkit-box;
              -webkit-line-clamp: 2;
              -webkit-box-orient: vertical;
            }
          }
        }
      }
    }
  }
}

@media screen and (max-width: 1100px) {
  .tools-header {
    .index-video {
      width: 1100px;
    }
  }
}

@media screen and (max-width: 900px) {
  .tools-container {
    .tools-content {
      .collects-wrap {
        .tools-item-wrap {
          .tools-item {
            width: calc(100% / 3 - 20px);
          }
        }
      }
    }
  }
}

@media screen and (max-width: 700px) {
  .tools-container {
    .tools-content {
      .collects-wrap {
        .tools-item-wrap {
          .tools-item {
            width: calc(100% / 2 - 20px);
          }
        }
      }
    }
  }
}

@media screen and (max-width: 620px) {
  .tools-container {
    .tools-header {
      .tools-header-content {
        .card-container {
          margin-top: 0;
        }
      }

      .index-video {
        width: 620px;
      }
    }
  }
}

@media screen and (max-width: 400px) {
  .tools-container {
    .tools-content {
      .collects-wrap {
        .tools-item-wrap {
          .tools-item {
            width: calc(100% - 20px);
          }
        }
      }
    }

    .tools-header {
      .index-video {
        width: 400px;
      }

      .tools-header-content {
        .card-container {
          display: none;
        }
      }
    }
  }
}
</style>
