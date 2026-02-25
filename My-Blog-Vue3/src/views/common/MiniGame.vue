<template>
  <!-- 小游戏 -->
  <el-dialog
    class="mini-game-dialog"
    title="随机小游戏"
    v-model="dialogVisible"
    :modal-append-to-body="false"
    width="auto"
    align="center"
  >
    <div class="mini-game-dialog-content">
      <iframe :src="game" style="width: 100%; height: 900px"></iframe>
    </div>
  </el-dialog>
</template>

<script setup lang="ts">
import { computed, type ComputedRef } from "vue";
import constant from "@/utils/constant";

interface Props {
  modelValue?: boolean;
  gameUrl?: string;
}

const props = withDefaults(defineProps<Props>(), {
  modelValue: false,
  gameUrl: "",
});

interface Emits {
  (e: "update:modelValue", value: boolean): void;
}

const emit = defineEmits<Emits>();

const dialogVisible: ComputedRef<boolean> = computed({
  get(): boolean {
    return props.modelValue;
  },
  set(val: boolean): void {
    emit("update:modelValue", val);
  },
});

const game: ComputedRef<string> = computed((): string => {
  return props.gameUrl || constant.gameURL;
});
</script>

<style lang="scss">
// 随机小游戏对话框
.mini-game-dialog {
  .mini-game-dialog-content {
    text-align: center;
  }
}
</style>
