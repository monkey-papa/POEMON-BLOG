<template>
  <div class="card-container" v-if="!$common.isEmpty(props.resourcePathList)">
    <div
      v-for="(resourcePath, index) in props.resourcePathList"
      :key="index"
      class="card-item shadow-box wow"
      @click="clickResourcePath(resourcePath)"
    >
      <div class="card-image">
        <el-image class="custom-el-image" lazy :src="resourcePath.cover" fit="cover">
          <template v-slot:error>
            <div class="image-slot myCenter">
              <div class="error-text">
                <div class="error-text-content">肥肠抱歉，图片跑掉了</div>
              </div>
            </div>
          </template>
        </el-image>
      </div>
      <div class="card-body">
        <div class="card-title">
          <el-image class="custom-el-image" lazy :src="resourcePath.friendAvatar" fit="cover">
            <template v-slot:error>
              <div class="image-slot myCenter"></div>
            </template>
          </el-image>
          {{ resourcePath.title }}
        </div>
        <div class="card-desc">
          {{ resourcePath.introduction }}
        </div>
        <div class="card-time">
          <el-icon size="14" class="icon-loading"><Loading /></el-icon>
          添加于 {{ $common.getDateDiff(resourcePath.createTime) }}
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import type { ResourcePath } from "@/types";

interface Props {
  resourcePathList?: ResourcePath[];
}

const props = defineProps<Props>();

const emit = defineEmits<{
  clickResourcePath: [url: string];
}>();

const clickResourcePath = (resourcePath: ResourcePath): void => {
  emit("clickResourcePath", resourcePath.url || "");
};
</script>
<style lang="scss" scoped>
.card-container {
  display: flex;
  flex-wrap: wrap;

  .card-item {
    position: relative;
    border-radius: 10px;
    background: var(--gray5);
    overflow: hidden;
    margin: 10px;
    height: 300px;
    flex-shrink: 0;
    width: calc(100% / 4 - 20px);
    animation: zoomIn 0.8s ease-in-out;
    border: 1px solid var(--gray1);
    transition: all 0.3s ease;

    .card-image {
      width: 100%;
      height: 180px;

      .custom-el-image {
        .image-slot {
          background: var(--gradualRed);

          .error-text-content {
            color: var(--wheat1);
          }
        }
      }

      :deep(.el-image__inner) {
        transition: all 1s;

        &:hover {
          transform: scale(1.2);
        }
      }
    }

    .card-body {
      padding: 10px 20px;

      .card-title {
        color: var(--bigRed1);
        font-size: 18px;
        font-weight: 400;
        margin-bottom: 10px;
        white-space: nowrap;
        text-overflow: ellipsis;
        overflow: hidden;
        transition: all 0.3s ease;

        .custom-el-image {
          width: 32px;
          height: 32px;
          margin-right: 4px;
          border-radius: 11px;

          .image-slot {
            height: 32px;
            width: 32px;
            background-image: var(--defaultAvatar);
            background-size: cover;
          }
        }

        &:hover {
          color: var(--black8);
        }
      }

      .card-desc {
        color: var(--darkBlue);
        font-size: 14px;
        line-height: 1.5;
        overflow: hidden;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-line-clamp: 2;
        -webkit-box-orient: vertical;
        letter-spacing: 1px;
        transition: all 0.3s ease;

        &:hover {
          color: var(--black8);
        }
      }

      .card-time {
        position: absolute;
        bottom: 4px;
        font-size: 12px;
        color: var(--bigRed1);
        transition: all 0.3s ease;

        &:hover {
          color: var(--black8);
        }

        .icon-loading {
          margin-right: 5px;
          color: var(--maxLightRed);
          animation: rotate 2s linear infinite;
        }
      }
    }

    &:hover {
      border-color: var(--gray4);
    }
  }
}

@media screen and (max-width: 1100px) {
  .card-container {
    .card-item {
      width: calc(100% / 3 - 20px);
    }
  }
}

@media screen and (max-width: 900px) {
  .card-container {
    .card-item {
      width: calc(100% / 2 - 20px);
    }
  }
}

@media screen and (max-width: 700px) {
  .card-container {
    .card-item {
      width: calc(100% - 20px);
    }
  }
}
</style>
