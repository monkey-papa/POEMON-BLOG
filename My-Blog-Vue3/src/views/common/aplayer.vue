<template>
  <div class="aplayer" @mouseenter="mouseenter" @mouseleave="mouseleave">
    <!-- 音乐播放器 -->
    <div id="aplayer"></div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from "vue";
import constant from "@/utils/constant";

const ap = ref<InstanceType<typeof window.APlayer> | null>(null);
const playIcon = ref<string>("fa-solid fa-play");
// const aplayerWidth = ref(null);
const lyricsTimer = ref<NodeJS.Timeout | null>(null);
const keydownHandler = ref<((e: KeyboardEvent) => void) | null>(null);

onMounted(() => {
  const server = "netease"; // netease: 网易云音乐; tencent: QQ音乐; kugou: 酷狗音乐; xiami: 虾米; kuwo: 酷我
  const type = "playlist"; // song: 单曲; playlist: 歌单; album: 唱片
  const id = "13875729534"; // 封面 ID / 单曲 ID / 歌单 ID
  const apiUrl = `${constant.metingApiURL}?server=${server}&type=${type}&id=${id}`;
  fetch(apiUrl)
    .then((response) => response.json())
    .then((data) => {
      const container = document.getElementById("aplayer");
      if (container) {
        ap.value = new window.APlayer({
          container,
          order: "random",
          preload: "none",
          listMaxHeight: "260px",
          fixed: "true",
          volume: 0.1,
          mutex: true,
          lrcType: 3,
          audio: data,
        });
        /* 底栏歌词 */
        lyricsTimer.value = setInterval(() => {
          $("#music-name").html($(".aplayer-lrc-current").text());
        }, 500);
        /* 音乐通知及控制 */
        ap.value.on("play", () => {
          playIcon.value = "fa-solid fa-pause";
        });
        ap.value.on("pause", () => {
          playIcon.value = "fa-solid fa-play";
        });
        keydownHandler.value = (e: KeyboardEvent) => {
          if (e.keyCode === 32) {
            ap.value!.toggle();
          }
        };
        window.addEventListener("keydown", keydownHandler.value);
        ap.value.list.show();
      }
    })
    .catch(() => {
      /* ignored */
    });
});

onBeforeUnmount(() => {
  if (lyricsTimer.value) {
    clearInterval(lyricsTimer.value);
    lyricsTimer.value = null;
  }
  if (keydownHandler.value) {
    window.removeEventListener("keydown", keydownHandler.value);
    keydownHandler.value = null;
  }
  if (ap.value) {
    ap.value.destroy();
    ap.value = null;
  }
});

// 以下函数保留供外部调用使用
const _playLast = (): void => {
  ap.value!.skipBack();
  ap.value!.on("play", () => {
    const currentMusic = ap.value!.list.current;
    const music = `${currentMusic.title} - ${currentMusic.author}`;
    window.iziToast.info({
      timeout: 4000,
      icon: "fa-solid fa-circle-play",
      displayMode: "replace",
      message: music,
    });
    $("#music-name").text(`${currentMusic.title} - ${currentMusic.author}`);
    playIcon.value = "fa-solid fa-pause";
  });
  ap.value!.on("pause", () => {
    playIcon.value = "fa-solid fa-play";
  });
};

const _handlePlay = (): void => {
  const aplayerEl = document.getElementById("aplayer");
  const music =
    aplayerEl!.querySelector(".aplayer-title")!.textContent +
    aplayerEl!.querySelector(".aplayer-author")!.textContent;
  window.iziToast.info({
    timeout: 4000,
    icon: "fa-solid fa-circle-play",
    displayMode: "replace",
    message: music,
  });
  playIcon.value = "fa-solid fa-pause";
  if (window.innerWidth >= 990) {
    // showPower = false;
    // showLrc = true;
  }
};

const _togglePlay = (): void => {
  ap.value!.toggle();
};

const _playNext = (): void => {
  ap.value!.skipForward();
};

const mouseenter = (): void => {
  // $(".aplayer-body").css({
  //   transition: "all 0.3s linear",
  //   transform: "translateX(68px)",
  // });
};

const mouseleave = (): void => {
  ap.value!.list.hide();
  // aplayerWidth.value = $(".aplayer-body").width();
  // if (aplayerWidth.value < 400) {
  //   $(".aplayer-body").css({
  //     transition: "all 0.3s linear",
  //     transform: "translateX(0px)",
  //   });
  // }
};
</script>

<style lang="scss" scoped>
.aplayer {
  margin: 0;
}
:deep(#aplayer) {
  margin: 0;

  &.aplayer-fixed {
    bottom: 330px;

    .aplayer-body {
      bottom: 330px;
      /* left: -68px; */
    }

    .aplayer-lrc {
      text-shadow: none;

      p {
        font-size: 16px;
        color: var(--bigRed);
      }
    }
  }

  .aplayer-miniswitcher {
    .aplayer-icon {
      path {
        fill: var(--darkBlue);
      }
    }
  }
}
</style>
