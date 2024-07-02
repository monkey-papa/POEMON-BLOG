<template>
  <div>
    <transition name="body">
      <div v-show="showEmoji">
        <span class="emoji-item" v-for="(value, key, index) in emojiListURL" :key="index" @click="addEmoji(key)">
          <img class="emoji" :src="value" :title="key" width="24px" height="24px" />
        </span>
      </div>
    </transition>
  </div>
</template>

<script>
export default {
  props: {
    showEmoji: {
      type: Boolean
    }
  },
  data() {
    return {
      emojiList: this.$constant.emojiList,
      emojiListURL: {}
    };
  },
  created() {
    this.emojiListURL = this.getEmojiList(this.emojiList);
  },
  methods: {
    addEmoji(key) {
      this.$emit("addEmoji", key);
    },
    getEmojiList(emojiList) {
      let emojiName;
      let url;
      let result = {}
      for (let i = 0; i < emojiList.length; i++) {
        emojiName = "[" + emojiList[i] + "]";
        let j = i + 1;
        url = this.$constant.qiniuUploadEntrance + "emoji/q" + j + ".gif";
        result[emojiName] = url;
      }
      return result;
    }
  }
}
</script>

<style scoped>
.emoji-item {
  display: inline-block;
}

.emoji-item:hover {
  transition: all 0.2s;
  border-radius: 0.25rem;
  background: var(--lightGray);
}

.emoji {
  margin: 0.25rem;
  /* 把此元素放置在父元素的中部 */
  vertical-align: middle;
}

.body-enter-active,
.body-leave-active {
  transition: all 0.3s;
}

.body-enter,
.body-leave-to {
  opacity: 0;
  transform: scale(0.5);
}
</style>
