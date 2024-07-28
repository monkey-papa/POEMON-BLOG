<template>
  <div class="myTag" :style="bg">
    <span class="name">🔖 {{ info }}</span>
  </div>
</template>

<script>
export default {
  props: {
    info: {
      type: String,
      default: "monkey-papa",
    },
    color: {
      type: String,
    },
  },
  data() {
    return {
      bg: { background: this.color },
    };
  },
};
</script>

<style lang="scss" scoped>
.myTag {
  color: var(--white);
  border-radius: 5px;
  padding: 8px;
  &:hover {
    background: var(--gradualRed) !important;
  }
}
</style>
