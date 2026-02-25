<template>
  <div class="myTag" :style="bg">
    <span class="name">🔖 {{ info }}</span>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue";

interface Props {
  info?: string;
  color?: string;
}

const props = withDefaults(defineProps<Props>(), {
  info: "tag",
});

const bg = computed(() => ({ background: props.color }));
</script>

<style lang="scss" scoped>
.myTag {
  color: var(--white);
  border-radius: 5px;
  padding: 8px;
  margin: 12px;

  &:hover {
    background: var(--gradualRed) !important;
  }
}
</style>
