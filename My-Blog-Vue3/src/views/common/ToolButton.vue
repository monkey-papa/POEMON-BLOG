<template>
  <!-- 右下角按钮 -->
  <div class="toolButton">
    <!-- 火箭 -->
    <div class="backTop" v-if="isMobile && showBackTop" @click="$emit('toTop')">
      <img src="../../assets/svg/rocket.svg" />
    </div>
    <!-- 右下角切换按钮 -->
    <el-popover
      transition="el-zoom-in-top"
      placement="left"
      :hide-after="500"
      trigger="hover"
      popper-class="toolButton-popover"
    >
      <!-- 旋转齿轮 -->
      <template #reference>
        <div>
          <div class="button myCenter">
            <i class="fa fa-cog setting-color" aria-hidden="true"></i>
          </div>
          <div class="button myCenter" v-show="topPercentage < 98 && !topPercentageType">
            {{ topPercentage }}%
          </div>
        </div>
      </template>
      <div class="my-setting">
        <el-tooltip placement="top" effect="light">
          <!-- 雪花片 -->
          <template #content>想看雪花吗？(◕ᴗ◕✿)</template>
          <div>
            <i
              class="fa fa-snowflake-o"
              :class="{ active: mouseAnimation }"
              aria-hidden="true"
              @click="$emit('changeMouseAnimation')"
            ></i>
          </div>
        </el-tooltip>
      </div>
    </el-popover>
  </div>
</template>

<script setup lang="ts">
interface Props {
  isMobile?: boolean;
  showBackTop?: boolean;
  topPercentage?: number;
  topPercentageType?: boolean;
  mouseAnimation?: boolean;
}

withDefaults(defineProps<Props>(), {
  isMobile: false,
  showBackTop: false,
  topPercentage: 0,
  topPercentageType: false,
  mouseAnimation: false,
});

interface Emits {
  (e: "toTop"): void;
  (e: "changeMouseAnimation"): void;
}

defineEmits<Emits>();
</script>

<style lang="scss" scoped>
// 右下角按钮
.toolButton {
  position: fixed;
  right: 3vh;
  bottom: 3vh;
  animation: slide-bottom 0.5s ease-in-out both;
  z-index: 100;
  font-size: 25px;

  .backTop {
    transition: all 0.3s ease-in;
    position: relative;
    top: 0;
    left: -4px;

    &:hover {
      top: -10px;
    }
  }

  .button {
    opacity: 0.6;
    border-radius: 10px;
    width: 38px;
    height: 38px;
    font-size: 15px;
    z-index: 20;
    position: relative;
    background-color: var(--blue3);
    color: var(--bigRed);
    text-align: center;
    line-height: 35px;
    margin-top: 8px;

    .setting-color {
      color: var(--bigRed);
      animation: rotate 4s linear infinite;
      font-size: large;
    }
  }
}

// 媒体查询
@media screen and (max-width: 400px) {
  .toolButton {
    right: 0.5vh;
  }
}
</style>

<style lang="scss">
// 底部工具栏下拉内容
.toolButton-popover {
  .my-setting {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-around;
    font-size: 20px;

    i {
      padding: 5px;
      color: var(--white);

      &:hover {
        color: var(--bigRed2);
      }
    }

    .active {
      color: var(--orange3);
    }
  }
}
</style>
