<template>
  <div
    style="animation: header-effect 2s"
    :style="{
      background: `${changeBgState && [''].includes($route.path) ? changeBgState : ''}`,
    }"
    class="poem-container myCenter my-animation-hideToShow background-image background-image-changeBg"
    v-if="!$common.isEmpty(guShi.origin) || !$common.isEmpty(hitokoto.hitokoto)"
  >
    <div class="poem-wrap">
      <div v-if="isShehui">
        <span>{{ webInfo.webName }}</span>
      </div>
      <div v-else>
        <span>{{ props.isHitokoto ? hitokoto.from : guShi.origin }}</span>
      </div>
      <p class="poem">
        {{ props.isHitokoto ? hitokoto.hitokoto : guShi.content }}
      </p>
      <p
        class="info"
        v-if="
          !props.isShehui &&
          (!props.isHitokoto || (props.isHitokoto && !$common.isEmpty(hitokoto.from_who)))
        "
      >
        {{ props.isHitokoto ? hitokoto.from_who : guShi.author }}
      </p>
    </div>
    <slot></slot>
  </div>
</template>
<script setup lang="ts">
import { ref, onMounted, computed } from "vue";
import { useStore } from "@/stores";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
import type { WebInfo } from "@/types/webInfo";

const store = useStore();
const { $route, $common, $constant } = useGlobalProperties();

const changeBgState = computed<string | undefined>(() => store.changeBg);
const webInfo = computed<WebInfo>(() => store.webInfo);

interface Props {
  isHitokoto?: boolean;
  isShehui?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  isHitokoto: true,
  isShehui: false,
});

interface GuShi {
  content: string;
  origin: string;
  author: string;
  category: string;
}

interface Hitokoto {
  hitokoto: string;
  from: string;
  from_who: string;
}

const guShi = ref<GuShi>({
  content: "...",
  origin: "...",
  author: "...",
  category: "...",
});

const hitokoto = ref<Hitokoto>({
  hitokoto: "...",
  from: "...",
  from_who: "...",
});

onMounted(() => {
  if (!props.isShehui) {
    if (props.isHitokoto) {
      getHitokoto();
    } else {
      getGuShi();
    }
  } else {
    hitokoto.value.from = "";
    hitokoto.value.from_who = "";
    sendShehui();
  }
});

const sendShehui = (): void => {
  const xhr = new XMLHttpRequest();
  xhr.open("get", $constant.shehui);
  xhr.onreadystatechange = function () {
    if (xhr.readyState === 4) {
      const shehui = xhr.responseText;
      hitokoto.value.hitokoto = shehui.substring(1, shehui.length - 1);
    }
  };
  xhr.send();
};

const getGuShi = (): void => {
  const xhr = new XMLHttpRequest();
  xhr.open("get", $constant.jinrishici);
  xhr.onreadystatechange = function () {
    if (xhr.readyState === 4) {
      guShi.value = JSON.parse(xhr.responseText) as GuShi;
    }
  };
  xhr.send();
};

const getHitokoto = (): void => {
  const xhr = new XMLHttpRequest();
  xhr.open("get", $constant.hitokoto);
  xhr.onreadystatechange = function () {
    if (xhr.readyState === 4) {
      hitokoto.value = JSON.parse(xhr.responseText) as Hitokoto;
    }
  };
  xhr.send();
};
</script>
<style lang="scss" scoped>
.poem-container {
  padding: 90px 0 40px;
  position: relative;

  .poem-wrap {
    border-radius: 10px;
    z-index: 10;
    text-align: center;
    letter-spacing: 4px;
    font-weight: 300;
    width: 100%;
    max-width: 800px;

    div span {
      padding: 5px 10px;
      color: var(--red);
      font-size: 2em;
      border-radius: 5px;
    }

    p {
      width: 100%;
      max-width: 800px;
      color: var(--darkBlue1);

      &.poem {
        margin: 40px auto;
        font-size: 1.5em;
      }

      &.info {
        margin: 20px auto 40px;
        font-size: 1.1em;
      }
    }
  }

  &.background-image-changeBg {
    height: 100vh !important;
  }
}
</style>
