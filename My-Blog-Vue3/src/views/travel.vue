<template>
  <div class="travel-container">
    <!-- 背景图片 -->
    <div
      style="animation: header-effect 2s"
      :style="{ background: `${changeBgState}` }"
      class="background-image background-image-changeBg"
    ></div>
    <div class="travel-top">
      <div class="travel-header shadow-box my-animation-slide-top">
        <!-- 顶部视频 -->
        <video
          class="index-video"
          autoplay
          muted
          loop
          playsinline
          webkit-playsinline
          :src="$constant.favoriteVideo"
        ></video>
        <div class="travel-header-title">
          <!-- 标题 -->
          <div class="travel-header-title-content">
            <div class="travel-header-title-content-left">旅拍集</div>
            <div class="travel-introduce">这里是相册分类集</div>
          </div>
        </div>
        <div class="travel-header-title-bottom">生活中的小确幸。</div>
      </div>
      <div v-if="!$common.mobile() && !mobile" class="shadow-box travel-header-right">
        <div class="card-content">
          <div class="author-content-item-tips">┗|｀O′|┛ 嗷～～</div>
          <span class="author-content-item-title">cat</span>
          <div class="banner-button-group">
            <a :href="$constant.siteURL" class="banner-button myCenter">
              <span class="banner-button-text">🎻 秘密基地</span>
            </a>
          </div>
        </div>
      </div>
    </div>
    <div class="travel-content my-animation-slide-bottom">
      <!-- 标签 -->
      <div class="photo-title-warp" v-if="!$common.isEmpty(photoTitleList)">
        <div
          v-for="(item, index) in photoTitleList"
          :key="index"
          :class="{ isActive: photoPagination.classify === item.classify }"
          @click="changePhotoTitle(item.classify || '')"
        >
          <proTag
            :info="(item.classify || '') + ' ' + (item.count || 0)"
            :color="$constant.tree_hole_color[Math.floor(Math.random() * 6)]"
          />
        </div>
      </div>
      <div class="photo-title">
        {{ photoPagination.classify }}
      </div>
      <photo :resourcePathList="photoList" />
      <div class="pagination-wrap">
        <div
          @click="pagePhotos()"
          class="pagination"
          v-if="photoPagination.total !== photoList.length"
        >
          下一页
        </div>
        <div class="pagination-end" v-else>~~( •̀ ω •́ )y 到底啦~~</div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import api from "@/api";
import {
  ref,
  onMounted,
  onBeforeUnmount,
  defineAsyncComponent,
  computed,
  type Ref,
  type ComputedRef,
} from "vue";
import { useStore } from "@/stores";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
import type { ResourcePath, ClassifyItem, PaginationState } from "@/types";

interface PhotoPagination extends PaginationState {
  resourceType: string;
  classify: string;
}

const store = useStore();
const { $common, $constant } = useGlobalProperties();

const changeBgState: ComputedRef<string> = computed(() => store.changeBg);

const photo = defineAsyncComponent(() => import("./common/photo.vue"));
const proTag = defineAsyncComponent(() => import("./common/proTag.vue"));

const photoPagination: Ref<PhotoPagination> = ref({
  current: 1,
  size: 10,
  total: 0,
  resourceType: "lovePhoto",
  classify: "",
});
const photoTitleList: Ref<ClassifyItem[]> = ref([]);
const photoList: Ref<ResourcePath[]> = ref([]);
const mobile: Ref<boolean> = ref(false);
let resizeHandler: (() => void) | null = null;

onMounted(() => {
  getPhotoTitles();
  mobile.value = document.body.clientWidth < 780;
  resizeHandler = (): void => {
    const docWidth = document.body.clientWidth;
    if (docWidth < 780) {
      mobile.value = true;
    } else {
      mobile.value = false;
    }
  };
  window.addEventListener("resize", resizeHandler);
});

onBeforeUnmount(() => {
  if (resizeHandler) {
    window.removeEventListener("resize", resizeHandler);
    resizeHandler = null;
  }
});

const getPhotoTitles = async (): Promise<void> => {
  try {
    const res = await api.getClassifyList({ type: "lovePhoto" });
    const items = (res || []).map((item: ClassifyItem) => ({
      ...item,
      classify: item.classify || "",
      count: item.count || 0,
    }));
    photoTitleList.value = items;
    if (photoTitleList.value.length > 0) {
      photoPagination.value = {
        current: 1,
        size: 10,
        total: 0,
        resourceType: "lovePhoto",
        classify: photoTitleList.value[0]?.classify || "",
      };
      changePhoto();
    }
  } catch {
    /* ignored */
  }
};

const changePhotoTitle = (classify: string): void => {
  if (classify !== photoPagination.value.classify) {
    photoPagination.value = {
      current: 1,
      size: 10,
      total: 0,
      resourceType: "lovePhoto",
      classify,
    };
    photoList.value = [];
    changePhoto();
  }
};

const pagePhotos = (): void => {
  photoPagination.value.current = photoPagination.value.current + 1;
  changePhoto();
};

const changePhoto = async (): Promise<void> => {
  try {
    const res = await api.getClientResourcePathList(photoPagination.value);
    photoList.value = photoList.value.concat(res.list || []);
    photoPagination.value.total = res.totalCount || photoList.value.length;
  } catch {
    /* ignored */
  }
};
</script>
<style lang="scss" scoped>
.travel-container {
  background: var(--background);

  .travel-top {
    display: flex;
    justify-content: center;

    .travel-header {
      margin: 60px 10px 30px;
      height: 300px;
      overflow: hidden;
      border-radius: 20px;
      width: 1130px;
      color: var(--red1);
      user-select: none;
      transition: all 0.3s ease;
      border: 1px solid var(--gray1);

      &:hover {
        color: var(--red);
        border-color: var(--gray4);
      }

      .index-video {
        width: 100%;
        height: 180%;
        object-fit: cover;
      }

      .travel-header-title {
        position: absolute;
        left: 20px;
        top: 20px;

        &-content {
          &-left {
            text-indent: 2em;
          }

          .travel-introduce {
            font-size: 36px;
            font-weight: 400;
            line-height: 1.5;
          }
        }
      }

      .travel-header-title-bottom {
        position: absolute;
        left: 20px;
        bottom: 40px;
        margin: 10px;
      }
    }

    .travel-header-right {
      right: 0;
      margin: 60px 10px 30px;
      height: 300px;
      overflow: hidden;
      border-radius: 20px;
      width: 270px;
      color: var(--red);
      user-select: none;
      position: relative;
      background-color: var(--brown);
      background-repeat: no-repeat;
      background-position: center;
      background-size: contain;
      color: var(--red1);
      transition: all 0.3s ease;
      border: 1px solid var(--gray1);
      background-image: url(); /* 请替换为你自己的背景图 */

      &:hover {
        color: var(--red);
        border-color: var(--gray4);
      }

      .card-content {
        position: absolute;
        width: 100%;
        height: 100%;
        top: 0;
        left: 0;
        z-index: 2;
        display: flex;
        flex-direction: column;
        padding: 1rem 2rem;

        .author-content-item-tips {
          font-size: 1.2em;
          margin-bottom: 0.5rem;
        }

        .banner-button-group {
          position: absolute;
          bottom: 0.3rem;
          right: 0.5rem;

          .banner-button {
            color: var(--red1);
            height: 2rem;
            width: 6rem;
            border-radius: 1.2rem;
            background: var(--favoriteBg);
            text-decoration: none;
            outline: none;
            transition: all 0.3s ease;

            &:hover {
              color: var(--red3);
              background: var(--themeColor);
            }
          }
        }
      }
    }
  }

  .travel-content {
    margin: 0 auto;
    max-width: 1550px;

    .photo-title-warp {
      max-width: 1250px;
      margin: 15px auto;
      padding: 20px;
      border-radius: 10px;
      display: flex;
      flex-wrap: wrap;

      > div {
        &.isActive {
          animation: scale 2.5s ease-in-out infinite;
        }
      }
    }

    .photo-title {
      color: var(--red);
      text-align: center;
      font-size: 30px;
      font-weight: 700;
      line-height: 80px;
      letter-spacing: 2px;
    }

    .pagination-wrap {
      display: flex;
      justify-content: center;
      margin-top: 40px;

      .pagination {
        margin-bottom: 8px;
        padding: 13px 15px;
        border: 1px solid var(--red);
        border-radius: 3rem;
        color: var(--red);
        width: 100px;
        user-select: none;
        text-align: center;

        &:hover {
          border: 1px solid var(--blue);
          color: var(--orange);
          box-shadow: 0 0 5px var(--blue);
        }
      }

      .pagination-end {
        user-select: none;
        color: var(--red);
      }
    }
  }
}

@media screen and (max-width: 1150px) {
  .travel-container {
    .travel-content {
      .photo-title-warp {
        max-width: 780px;
      }
    }
  }
}
</style>
