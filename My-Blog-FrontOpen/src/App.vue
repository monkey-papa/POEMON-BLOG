<template>
  <div id="app" @contextmenu.prevent="openMenu($event,item)">
    <router-view />
    <aplayer></aplayer>
    <div v-if="$store.state.isShowLoading" class="loading">
      <div class="author-box">
        <span></span>
        <div class="author-img">
          <img src="./assets/file/avatar.jpg" alt="">
        </div>
      </div>
    </div>
    <!-- 右键菜单部分 -->
    <ul v-show="visible" :style="{left:left+'px',top:top+'px'}" class="contextmenu">
      <div class="rightMenu-group">
        <div class="rightMenu-item">
          <i class="fa fa-arrow-left" @click="backAndForward('2')"></i>
        </div>
        <div class="rightMenu-item">
          <i class="fa fa-arrow-right" @click="backAndForward('1')"></i>
        </div>
        <div class="rightMenu-item">
          <i class="fa fa-repeat" @click="refresh"></i>
        </div>
        <div class="rightMenu-item">
          <i class="fa fa-music" @click="musicHandle"></i>
        </div>
      </div>
      <div class="rightMenu-group rightMenu-line">
        <p @click="$router.push('/')" class="rightMenu-item">
          <i class="fa fa-home"></i>
          <span>博客首页</span>
        </p>
        <p @click="dayAndNight" class="rightMenu-item">
          <i class="fa fa-adjust"></i>
          <span>昼夜切换</span>
        </p>
        <p @click="onCopy" class="rightMenu-item">
          <i class="fa fa-code-fork"></i>
          <span>加入我们</span>
        </p>
        <p @click="changeTheme" class="rightMenu-item">
          <i class="fa fa-image"></i>
          <span>美化设置</span>
        </p>
        <p @click="$router.push('/tags?labelId=15')" class="rightMenu-item">
          <i class="fa fa-bookmark"></i>
          <span>标签</span>
        </p>
        <p @click="$router.push('/sort?sortId=1')" class="rightMenu-item">
          <i class="fa fa-folder-open"></i>
          <span>分类</span>
        </p>
      </div>
    </ul>
  </div>
</template>
<script>
const aplayer = () => import('./views/common/aplayer.vue')
export default {
  name: 'App',
  components: {
    aplayer
  },
  data() {
    return {
      coorY: 0,
      rightClickItem: '',
      visible: false, // 是否展示右键菜单
      top: 0,
      left: 0,
      copyContent: 'z15523692545'
    }
  },
  mounted() {
    window.addEventListener('scroll', this.handleScroll) // 监听滚动条事件
    window.addEventListener('beforeunload', this.handleBeforeUnload)
    let styleTitle1 = `
font-size: 20px;
font-weight: 600;
color: rgb(244,167,89);
`
    let styleTitle2 = `
font-size:12px;
color: #425AEF;
`
    let styleContent = `
color: rgb(30,152,255);
`
    let title1 = 'ZJHの主页 被我发现了吧，既然你已经破解了，转发、拿东西记得标明出处喔~~'
    let title2 = `
    く__,.ヘヽ.        /  ,ー､ 〉
           ＼ ', !-─‐-i  /  /´
           ／｀ｰ'       L/／｀ヽ､
         /   ／,   /|   ,   ,       ',
       ｲ   / /-‐/  ｉ  L_ ﾊ ヽ!   i
        ﾚ ﾍ 7ｲ｀ﾄ   ﾚ'ｧ-ﾄ､!ハ|   |
          !,/7 '0'     ´0iソ|    |
          |.从"    _     ,,,, / |./    |
          ﾚ'| i＞.､,,__  _,.イ /   .i   |
            ﾚ'| | / k_７_/ﾚ'ヽ,  ﾊ.  |
              | |/i 〈|/   i  ,.ﾍ |  i  |
             .|/ /  ｉ：    ﾍ!    ＼  |
              kヽ>､ﾊ    _,.ﾍ､    /､!
              !'〈//｀Ｔ´', ＼ ｀'7'ｰr'
              ﾚ'ヽL__|___i,___,ンﾚ|ノ
                  ﾄ-,/  |___./
                  'ｰ'    !_,.:
██████╗ ██╗   ██╗║██████╗ 
██╔══██╗╚██╗ ██╔╝║██╔═══╗
██████╔╝ ╚████╔╝ ║██████║
██╔══██╗  ╚██╔╝  ║██╔═══╗
██████╔╝   ██║   ║██████║
╚═════╝    ╚═╝   ╚══════╝(wx:z15523692545)OVO
`
    let content = `
誰もが信じ崇めてる
まさに最強で無敵のアイドル
弱点なんて見当たらない
一番星を宿している
弱いとこなんて見せちゃダメダメ
知りたくないとこは見せずに
唯一無二じゃなくちゃイヤイヤ
それこそ本物のアイ
  ⌜IDOL⌟
`
    console.log(
      `%c${title1} %c${title2}
%c${content}`,
      styleTitle1,
      styleTitle2,
      styleContent
    )
    this.$nextTick(() => {
      //禁止右键
      document.oncontextmenu = new Function('event.returnValue=false')
    })
  },
  beforeDestroy() {
    window.removeEventListener('beforeunload', this.handleBeforeUnload)
    window.removeEventListener('scroll', this.handleScroll)
  },
  watch: {
    // 监听 visible，来触发关闭右键菜单，调用关闭菜单的方法
    visible(value) {
      if (value) {
        document.body.addEventListener('click', this.closeMenu)
      } else {
        document.body.removeEventListener('click', this.closeMenu)
      }
    }
  },
  methods: {
    // 打开右键菜单
    openMenu(e, item) {
      this.visible = true
      this.top = e.pageY
      this.left = e.pageX
      this.rightClickItem = item
    },
    // 关闭右键菜单
    closeMenu() {
      this.visible = false
    },
    // 音乐跳转
    musicHandle() {
      window.open('https://www.zjh2002.icu/#/discover/recommend')
    },
    // 路由跳转
    backAndForward(val) {
      if (val === '1') {
        window.history.forward()
      } else {
        window.history.back()
      }
    },
    // 刷新
    refresh() {
      location.reload()
    },
    // 昼夜切换
    dayAndNight() {
      document.getElementById('changeColorRef').click()
    },
    // 加入我们
    onCopy() {
      let input = document.createElement('input') // 直接构建input
      input.value = this.copyContent // 设置内容
      document.body.appendChild(input) // 添加临时实例
      input.select() // 选择实例内容
      document.execCommand('Copy') // 执行复制
      document.body.removeChild(input) // 删除临时实例
      this.$notify({
        title: '可以啦🍨',
        message: '本博主的微信已经到你的剪贴板啦，快加入我们吧~~🎉',
        type: 'success',
        offset: 50,
        duration: 0
      })
    },
    // 美化设置
    changeTheme() {
      document.getElementById('changeThemeRef').click()
    },
    handleScroll() {
      // 屏幕剩余的高度
      let surplus = document.documentElement.scrollHeight - document.documentElement.clientHeight
      // 当前滑动高度
      let scrollY = document.documentElement.scrollTop
      if (scrollY < 0) {
        scrollY = 0
      }
      if (scrollY > 0) {
        this.visible = false
      }
      // 当前位置百分比小数
      this.coorY = scrollY / surplus
      this.$store.commit('topPercentage', Math.floor(this.coorY * 100))
      // 设置导航栏，这里使用NProgress.set() 动态更改进度条
      NProgress.set(this.coorY)
    },
    handleBeforeUnload() {
      this.$store.commit('topPercentage', 0)
    }
  }
}
</script>
<style>
#nprogress .bar {
  background: linear-gradient(to right, #61c454, #ec695c) no-repeat !important;
  height: 5px !important;
}

#nprogress .peg {
  box-shadow: 0 0 10px #dd181800, 0 0 5px #c2282800 !important;
}
</style>

<style scoped>
.contextmenu {
  margin: 0;
  background: white;
  z-index: 3000;
  position: absolute;
  width: 9rem;
  height: fit-content;
  border-radius: 12px;
  border: 1px solid #e3e8f7;
  font-size: 12px;
  font-weight: 700;
  color: #353535;
  transition: 0.3s;
  padding: 0 0.25rem;
}
.rightMenu-group {
  padding: 0.35rem 0.3rem;
  display: flex;
  justify-content: space-between;
}
.rightMenu-group:not(:nth-last-child(1)) {
  border-bottom: 1px dashed #4259ef23;
}
.rightMenu-group .rightMenu-item {
  border-radius: 8px;
  transition: 0.3s;
}
.rightMenu-group .rightMenu-item i {
  font-size: 15px;
  display: inline-block;
  text-align: center;
  line-height: 1.5rem;
  font-weight: 900;
  width: 1.5rem;
  padding: 0 0.25rem;
}
.rightMenu-line {
  display: block;
}
.rightMenu-line .rightMenu-item {
  margin: 0.25rem 0;
  padding: 0.25rem 0;
  display: flex;
  font-size: 15px;
}
.rightMenu-line .rightMenu-item i {
  margin: 0 0.25rem;
}
.rightMenu-line .rightMenu-item span {
  line-height: 1.5rem;
  font-weight: 500;
}
a.rightMenu-item {
  color: #353535;
  text-decoration: none;
}
.rightMenu-group .rightMenu-item:hover {
  background-color: #1dbffff1;
  color: white;
  box-shadow: 0 8px 12px -3px #4259ef23;
}
.loading {
  width: 100%;
  height: 100%;
  background: linear-gradient(55deg, #0095c2 21%, #64e1c8 100%);
  position: absolute;
  top: 0px;
  left: 0px;
  z-index: 99;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 30px;
}

.author-box {
  position: relative;
  width: 159px;
  height: 159px;
  background: rgba(253, 253, 253, 0.8);
  border-radius: 50%;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
}

.author-box::before {
  content: '';
  position: absolute;
  width: 500px;
  height: 500px;
  background-image: conic-gradient(transparent, transparent, transparent, #8758ff);
  animation: animate 2s linear infinite;
  animation-delay: -1s;
}

.author-box::after {
  content: '';
  position: absolute;
  width: 500px;
  height: 500px;
  background-image: conic-gradient(transparent, transparent, transparent, #5cb8e4);
  animation: animate 2s linear infinite;
}

@keyframes animate {
  0% {
    transform: rotate(0);
  }

  100% {
    transform: rotate(360deg);
  }
}

.author-box span {
  position: absolute;
  inset: 5px;
  border-radius: 50%;
  background: rgba(253, 253, 253, 0.8);
  z-index: 1;
}

.author-img {
  margin: auto;
  border-radius: 50%;
  overflow: hidden;
  width: 150px;
  height: 150px;
  z-index: 10;
  background: rgba(255, 255, 255, 0.67);
}

.author-img img {
  border-radius: 11px;
  margin-right: 4px;
  display: block;
  margin: 0 auto 20px;
  max-width: 100%;
  animation: breath 700ms ease-in-out infinite;
}

@keyframes breath {
  from {
    opacity: 0.7;
  }

  25% {
    opacity: 0.9;
  }

  50% {
    opacity: 1;
  }

  75% {
    opacity: 0.9;
  }

  to {
    opacity: 0.7;
  }
}
</style>