<template>
  <div>
    <slot name="paper" :content="content"></slot>
  </div>
</template>
<script setup lang="ts">
import { ref, watch, onMounted, onBeforeUnmount } from "vue";

interface Props {
  printerInfo?: string;
  duration?: number;
  delay?: number;
  working?: boolean;
  once?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  printerInfo: "",
  duration: 100,
  delay: 3000,
  working: true,
  once: false,
});

const content = ref<string>("");
const cursor = ref<number>(0);
const timer = ref<ReturnType<typeof setInterval> | null>(null);
const timeout = ref<ReturnType<typeof setTimeout> | null>(null);
const print = ref<boolean>(true);

/**
 * 逻辑
 */
const work = (): void => {
  let currentCursor = cursor.value;
  currentCursor += print.value ? 1 : -1;
  if (print.value) {
    if (currentCursor === props.printerInfo.length + 1) {
      currentCursor -= 2;
      print.value = !print.value;
    }
  } else {
    if (currentCursor === -1) {
      currentCursor += 2;
      print.value = !print.value;
    }
  }
  cursor.value = currentCursor;
};

/**
 * 定时
 */
const start = (workFn: () => void): void => {
  // 延迟
  timeout.value = setTimeout(() => {
    // 速度
    const intervalId = setInterval(() => {
      workFn();
      if (cursor.value === 0 || (cursor.value === props.printerInfo.length && !props.once)) {
        // 此处为了延迟
        clearInterval(intervalId);
        start(work);
      } else if (cursor.value === props.printerInfo.length && props.once) {
        clearInterval(intervalId);
      }
    }, props.duration);
    timer.value = intervalId;
  }, props.delay);
};

const toBegin = (): void => {
  cursor.value = 0;
  if (timeout.value !== null) {
    clearTimeout(timeout.value);
    if (timer.value !== null) {
      clearInterval(timer.value);
    }
  }
  if (props.working) {
    start(work);
  } else {
    content.value = props.printerInfo;
  }
};

watch(
  () => props.working,
  (newVal) => {
    if (newVal) {
      toBegin();
    } else {
      if (timeout.value !== null) {
        clearTimeout(timeout.value);
        if (timer.value !== null) {
          clearInterval(timer.value);
        }
      }
      content.value = props.printerInfo;
    }
  },
  { immediate: true }
);

watch(
  () => props.printerInfo,
  () => {
    toBegin();
  }
);

watch(
  () => props.delay,
  () => {
    toBegin();
  }
);

watch(
  () => props.duration,
  () => {
    toBegin();
  }
);

onMounted(() => {
  toBegin();
});

onBeforeUnmount(() => {
  if (timeout.value !== null) {
    clearTimeout(timeout.value);
  }
  if (timer.value !== null) {
    clearInterval(timer.value);
  }
});

watch(
  () => cursor.value,
  (cursorVal) => {
    // slice(start,end)：不包含end
    content.value = props.printerInfo.slice(0, cursorVal);
  }
);
</script>
