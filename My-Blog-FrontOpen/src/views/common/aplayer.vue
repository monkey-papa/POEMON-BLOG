<template>
  <div class="aplayer" @mouseenter="mouseenter" @mouseleave="mouseleave">
    <!-- 音乐播放器 -->
    <div id="aplayer"></div>
  </div>
</template>

<script>
import '../../utils/APlayer.min.js'
export default {
  data() {
    return {
      ap: null,
      playIcon: 'fa-solid fa-play',
      aplayerWidth: null
    }
  },
  mounted() {
    let server = 'netease' //netease: 网易云音乐; tencent: QQ音乐; kugou: 酷狗音乐; xiami: 虾米; kuwo: 酷我
    let type = 'playlist' //song: 单曲; playlist: 歌单; album: 唱片
    let id = '7452421335' //封面 ID / 单曲 ID / 歌单 ID
    const apiUrl = `https://api-meting.imsyy.top/api?server=${server}&type=${type}&id=${id}`
    fetch(apiUrl)
      .then(response => response.json())
      .then(data => {
        this.ap = new APlayer({
          container: document.getElementById('aplayer'),
          order: 'random',
          preload: 'none',
          listMaxHeight: '260px',
          fixed: 'true',
          volume: 0.1,
          mutex: true,
          lrcType: 3,
          audio: data
        })
        /* 底栏歌词 */
        setInterval(() => {
          $('#music-name').html($('.aplayer-lrc-current').text())
        }, 500)
        /* 音乐通知及控制 */
        this.ap.on('play', () => {
          const currentMusic = this.ap.list.current
          this.playIcon = 'fa-solid fa-pause'
        })
        this.ap.on('pause', () => {
          this.playIcon = 'fa-solid fa-play'
        })
        window.onkeydown = e => {
          if (e.keyCode === 32) {
            this.ap.toggle()
          }
        }
        this.ap.list.show()
      })
  },
  methods: {
    playLast() {
      this.ap.skipBack()
      this.ap.on('play', () => {
        const music = `${currentMusic.title} - ${currentMusic.author}`
        iziToast.info({
          timeout: 4000,
          icon: 'fa-solid fa-circle-play',
          displayMode: 'replace',
          message: music
        })
        $('#music-name').text(`${currentMusic.title} - ${currentMusic.author}`)
        this.playIcon = 'fa-solid fa-pause'
      })
      this.ap.on('pause', () => {
        this.playIcon = 'fa-solid fa-play'
      })
    },
    handlePlay() {
      const music = this.$refs.aplayer.$el.querySelector('.aplayer-title').textContent + this.$refs.aplayer.$el.querySelector('.aplayer-author').textContent
      iziToast.info({
        timeout: 4000,
        icon: 'fa-solid fa-circle-play',
        displayMode: 'replace',
        message: music
      })
      this.playIcon = 'fa-solid fa-pause'
      this.musicName = music
      if (window.innerWidth >= 990) {
        this.showPower = false
        this.showLrc = true
      }
    },
    togglePlay() {
      this.ap.toggle()
    },
    playNext() {
      this.ap.skipForward()
    },
    mouseenter() {
      // $(".aplayer-body").css({
      //   transition: "all 0.3s linear",
      //   transform: "translateX(68px)",
      // });
    },
    mouseleave() {
      this.ap.list.hide()
      // this.aplayerWidth = $(".aplayer-body").width();
      // if (this.aplayerWidth < 400) {
      //   $(".aplayer-body").css({
      //     transition: "all 0.3s linear",
      //     transform: "translateX(0px)",
      //   });
      // }
    }
  }
}
</script>

<style lang="scss" scoped>
.aplayer {
  margin: 0;
}
::v-deep #aplayer {
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
  .aplayer-miniswitcher .aplayer-icon path {
    fill: var(--darkBlue);
  }
}
</style>
