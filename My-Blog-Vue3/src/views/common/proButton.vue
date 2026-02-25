<template>
  <div class="myButton">
    <div :style="beforeColor">{{ info }}</div>
    <div :style="afterColor">{{ info }}</div>
    <div :style="afterColor">{{ info }}</div>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue";

interface Props {
  info?: string;
  before?: string;
  after?: string;
}

const props = withDefaults(defineProps<Props>(), {
  info: "确定",
});

const beforeColor = computed(() => ({ background: props.before }));
const afterColor = computed(() => ({ background: props.after }));
</script>

<style lang="scss" scoped>
.myButton {
  user-select: none;
  position: relative;
  width: 66px;
  height: 33px;
  border-radius: 4px;
  color: var(--white);
  font-size: 14px;
  overflow: hidden;

  div {
    width: 66px;
    height: 33px;
    line-height: 33px;
    border-radius: 4px;
    text-align: center;
    position: absolute;

    &:nth-child(2) {
      width: 100px;
      transition: all 0.3s ease;
      transform: translateX(-120px) skewX(-30deg);
    }

    &:nth-child(3) {
      transition: all 0.3s ease;
      transform: translateX(-120px);
    }
  }

  &:hover div:nth-child(2) {
    transform: translateX(20px) skewX(-30deg);
  }

  &:hover div:nth-child(3) {
    transform: translateX(0px);
  }
}
</style>
