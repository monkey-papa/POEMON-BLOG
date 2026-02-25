<template>
  <div class="card-container photoRef" v-if="!$common.isEmpty(props.resourcePathList)">
    <div
      v-for="(resourcePath, index) in props.resourcePathList"
      :key="index"
      class="card-item wow shadow-box-mini"
    >
      <div class="card-image container">
        <el-image
          :style="{
            opacity: hoverIndex === -1 ? 1 : index === hoverIndex ? 1 : 0.2,
          }"
          crossorigin="anonymous"
          @mouseenter="handleMouseEnter($event, index)"
          @mouseleave="handleMouseLeave"
          class="custom-el-image"
          lazy
          :preview-src-list="[resourcePath.cover]"
          :src="resourcePath.cover"
          fit="cover"
        >
          <template v-slot:error>
            <div class="image-slot">
              <div class="error-aside-image">
                <div class="error-text">肥肠抱歉，图片跑掉了┮﹏┭</div>
              </div>
            </div>
          </template>
        </el-image>
      </div>
      <div class="card-body">
        <el-tooltip placement="bottom-start" effect="light">
          <template #content>
            {{ resourcePath.title }}
          </template>
          <div class="card-desc">
            <img src="../../assets/svg/leaf.svg" />
            {{ resourcePath.title }}
          </div>
        </el-tooltip>
        <div class="card-time">Date: {{ $common.getDateDiff(resourcePath.createTime) }}</div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref, onMounted, type Ref } from "vue";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
import type { ResourcePath } from "@/types";

// ColorThief 通过 CDN 引入，全局变量为 ColorThief
// 使用全局类型定义 (env.d.ts 中已定义 ColorThiefStatic)
const ColorThief = window.ColorThief;

interface Props {
  resourcePathList?: ResourcePath[];
}

const props = defineProps<Props>();

const { $route } = useGlobalProperties();

const hoverIndex: Ref<number> = ref(-1);

onMounted(() => {
  // 组件挂载时不需要特殊处理
});

const handleMouseEnter = async (event: MouseEvent, i: number): Promise<void> => {
  try {
    const colorThief = new ColorThief();
    const html = document.documentElement;
    hoverIndex.value = i;

    // 找到实际的 img 元素
    // el-image 组件内部的实际图片元素在 .el-image__inner 中
    let imgElement: HTMLImageElement | null = event.target as HTMLImageElement | null;
    if (!imgElement || imgElement.tagName !== "IMG") {
      // 如果不是 img 元素，尝试在组件内部查找
      const elImage =
        (event.currentTarget as HTMLElement) || (event.target as HTMLElement)?.closest(".el-image");
      if (elImage) {
        imgElement =
          (elImage.querySelector("img.el-image__inner") as HTMLImageElement) ||
          (elImage.querySelector("img") as HTMLImageElement) ||
          (elImage.querySelector(".el-image__inner") as HTMLImageElement);
      }
    }

    // 如果还是找不到，尝试通过图片 URL 创建新的图片元素
    if (!imgElement || imgElement.tagName !== "IMG") {
      const imgUrl = props.resourcePathList?.[i]?.cover;
      if (imgUrl) {
        imgElement = new Image();
        imgElement.crossOrigin = "anonymous";
        imgElement.src = imgUrl;
        // 等待图片加载完成
        await new Promise<void>((resolve, reject) => {
          if (!imgElement) {
            reject(new Error("Image element not found"));
            return;
          }
          imgElement.onload = () => resolve();
          imgElement.onerror = () => reject(new Error("Image load failed"));
          // 如果图片已经加载，立即 resolve
          if (imgElement.complete) {
            resolve();
          }
        });
      }
    }

    if (imgElement && imgElement.tagName === "IMG") {
      // 确保图片已加载
      if (!imgElement.complete) {
        await new Promise<void>((resolve) => {
          imgElement!.onload = () => resolve();
          if (imgElement!.complete) resolve();
        });
      }

      // 得到这张图片的调色盘（前三种主要颜色）
      const colors = await colorThief.getPalette(imgElement, 3);
      if (!colors || colors.length < 3) {
        return;
      }
      const [c1, c2, c3] = colors.map((c) => `rgb(${c[0]},${c[1]},${c[2]})`);

      if ($route.path == "/travel") {
        const photoRef = document.querySelector(".photoRef") as HTMLElement | null;
        if (photoRef) {
          photoRef.style.setProperty("--background", "var(--transparent)");
          photoRef.style.setProperty("--c1", c1);
          photoRef.style.setProperty("--c2", c2);
          photoRef.style.setProperty("--c3", c3);
        }
      } else {
        const root = document.querySelector(":root") as HTMLElement | null;
        if (root) {
          root.style.setProperty("--background", "var(--transparent)");
          html.style.setProperty("--c1", c1);
          html.style.setProperty("--c2", c2);
          html.style.setProperty("--c3", c3);
        }
      }
    }
  } catch {
    /* ignored */
  }
};

const handleMouseLeave = (): void => {
  hoverIndex.value = -1;
};
</script>
<style lang="scss" scoped>
.card-container {
  display: flex;
  flex-wrap: wrap;
  border-radius: 1.5rem;

  .card-item {
    position: relative;
    overflow: hidden;
    margin: 15px;
    height: 510px;
    flex-shrink: 0;
    width: calc(100% / 3 - 30px);
    animation: zoomIn 0.8s ease-in-out;
    padding: 1.3rem 1rem 1.5rem;
    background: var(--background);
    border-radius: 1.5rem;
    transition: all 0.2s;

    .card-image {
      width: 100%;
      height: 400px;
      border-radius: 1rem;
      overflow: hidden;
      box-shadow: 0 -2px 10px var(--whiteMask);
      margin-bottom: 1rem;

      :deep(.el-image__inner) {
        transition: all 1s;

        &:hover {
          transform: scale(1.2);
        }
      }

      .error-aside-image {
        background: var(--gradientAnimation);
        color: var(--white);
        text-align: center;
        height: 100%;

        .error-text {
          color: var(--wheat1);
          position: relative;
          top: 50%;
          transform: translate(0, -50%);
        }
      }
    }

    .card-body {
      padding: 10px 5px;

      .card-desc {
        font-weight: 400;
        font-size: 1.05rem;
        color: var(--red1);
        letter-spacing: 1px;
        line-height: 1.5;
        overflow: hidden;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-line-clamp: 2;
        line-clamp: 2;
        -webkit-box-orient: vertical;

        img {
          vertical-align: -2px;
        }

        &:hover {
          color: var(--red);
        }
      }

      .card-time {
        position: absolute;
        bottom: 15px;
        color: var(--green4);
        font-weight: 500;

        &:hover {
          color: var(--green1);
        }
      }
    }
  }
}

@media screen and (max-width: 1300px) {
  .card-container {
    .card-item {
      width: calc(100% / 2 - 30px);
    }
  }
}

@media screen and (max-width: 1000px) {
  .card-container {
    .card-item {
      height: 410px;

      .card-image {
        height: 300px;
      }
    }
  }
}

@media screen and (max-width: 750px) {
  .card-container {
    .card-item {
      width: 100%;
      margin: 10px 0;
    }
  }
}

@media screen and (max-width: 450px) {
  .card-container {
    .card-item {
      height: 360px;

      .card-image {
        height: 250px;
      }
    }
  }
}
</style>
