<template>
  <div>
    <Transition name="body">
      <div v-show="showEmoji">
        <span
          class="emoji-item"
          v-for="(value, key, index) in emojiListURL"
          :key="index"
          @click="addEmoji(key)"
        >
          <img class="emoji" :src="value" :title="key" />
        </span>
      </div>
    </Transition>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useGlobalProperties } from "@/composables/useGlobalProperties";

interface Props {
  showEmoji?: boolean;
}

defineProps<Props>();

const emit = defineEmits<{
  addEmoji: [key: string];
}>();

const { $constant } = useGlobalProperties();

const emojiList = ref<string[]>($constant.emojiList);
const emojiListURL = ref<Record<string, string>>({});

onMounted(() => {
  emojiListURL.value = getEmojiList(emojiList.value);
});

const getEmojiList = (emojiListData: string[]): Record<string, string> => {
  const result: Record<string, string> = {};
  for (let i = 0; i < emojiListData.length; i++) {
    const emojiName = "[" + emojiListData[i] + "]";
    const j = i + 1;
    const url = $constant.qiniuUploadEntrance + "emoji/q" + j + ".gif";
    result[emojiName] = url;
  }
  return result;
};

const addEmoji = (key: string): void => {
  emit("addEmoji", key);
};
</script>

<style lang="scss" scoped>
.emoji-item {
  display: inline-block;

  &:hover {
    transition: all 0.2s;
    border-radius: 0.25rem;
    background: var(--gray15);
  }

  .emoji {
    margin: 0.25rem;
    vertical-align: middle;
    height: 24px;
    width: 24px;
  }
}

.body-enter-active,
.body-leave-active {
  transition: all 0.3s;
}

.body-enter-form,
.body-leave-to {
  opacity: 0;
  transform: scale(0.5);
}
</style>
