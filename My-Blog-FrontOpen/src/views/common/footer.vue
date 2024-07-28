<template>
  <footer id="footer">
    <div id="footer_icons">
      <div>
        <a class="icon_link" target="_blank" rel="noopener" href="http://www.blog.zjh2002.icu" title="个人主页">
          <i class="iconfont icon-zhinanzhen"></i>
        </a>
        <a @click="copy(1)" class="icon_link contact-qq" href="javascript:;" title="联系QQ" data-pjax-state="">
          <i class="iconfont icon-qq"></i>
        </a>
        <a @click="copy(2)" class="icon_link contact-wechat" href="javascript:;" title="联系微信" data-pjax-state="">
          <i class="iconfont icon-weixin"></i>
        </a>
        <a class="icon_link" href="https://blog.csdn.net/qq_53461589?spm=1010.2135.3001.5421" title="CSDN主页" data-pjax-state="">
          <i class="iconfont icon-csdn">
          </i>
        </a>
      </div>
      <img class="footer_logo entered loaded" :src="require('../../../public/title.jpg')" data-lazy-src="../../../public/title.jpg" data-ll-status="loaded" />
      <div>
        <a class="icon_link" target="_blank" rel="noopener" href="https://github.com/monkey-papa/POEMON-BLOG" title="Github主页">
          <i class="iconfont icon-github"></i>
        </a>
        <a class="icon_link" target="_blank" rel="noopener" href="https://www.bilibili.com/video/BV1nh4y1m7CV/?spm_id_from=333.999.0.0" title="哔哩哔哩主页">
          <i class="iconfont icon-wf_Bzhan"></i>
        </a>
        <a class="icon_link" target="_blank" rel="noopener" href="https://travellings.link" title="随机跳转到另一个加入开往的网页">
          <i class="iconfont icon-huochepiao"></i>
        </a>
        <a class="icon_link" target="_blank" rel="noopener" href="https://www.foreverblog.cn/go.html" title="随机访问十年之约友链博客">
          <i class="iconfont icon-AppIcon"></i>
        </a>
      </div>
    </div>
    <div id="footer_content">
      <div class="footer-group">
        <h3 class="footer-title">直达</h3>
        <div class="footer-links">
          <a class="footer-item" href="http://www.blog.zjh2002.icu/#/tools" data-pjax-state="">藏宝阁</a>
          <a class="footer-item" href="http://www.blog.zjh2002.icu/#/friend" data-pjax-state="">友链</a>
          <a class="footer-item" href="http://www.blog.zjh2002.icu/#/message" data-pjax-state="">留言厅</a>
          <a class="footer-item" href="http://www.blog.zjh2002.icu/#/travel " data-pjax-state="">个人相册</a>
          <a class="footer-item" href="http://www.blog.zjh2002.icu/#/funny" data-pjax-state="">幻音坊</a>
        </div>
      </div>
      <div class="footer-group">
        <h3 class="footer-title">分类</h3>
        <div class="footer-links">
          <a class="footer-item" href="http://www.blog.zjh2002.icu/#/sort?sortId=1" data-pjax-state="">学习笔记</a>
          <a class="footer-item" href="http://www.blog.zjh2002.icu/#/sort?sortId=6" data-pjax-state="">我的项目</a>
          <a class="footer-item" href="http://www.blog.zjh2002.icu/#/sort?sortId=3" data-pjax-state="">游戏</a>
          <a class="footer-item" href="http://www.blog.zjh2002.icu/#/sort?sortId=4" data-pjax-state="">动漫</a>
          <a class="footer-item" href="http://www.blog.zjh2002.icu/#/sort?sortId=7" data-pjax-state="">随想</a>
          <a class="footer-item" href="http://www.blog.zjh2002.icu/#/sort" data-pjax-state="">查看全部</a>
        </div>
      </div>
      <div class="footer-group" id="friend-links-in-footer">
        <h3 class="footer-title">友链</h3>
        <div class="footer-links">
          <a class="footer-item" target="_blank" href="https://poetize.cn/">POETIZE</a>
          <a class="footer-item" target="_blank" href="https://www.qcqx.cn/">轻笑Chuckle</a>
          <a class="footer-item" target="_blank" href="https://blog.ezuxoay.top/home">Ezuxoay</a>
          <a class="footer-item" target="_blank" href="https://blog.zhengxian.top/">380AM-0204</a>
          <a class="footer-item" href="http://www.blog.zjh2002.icu/#/friend" data-pjax-state="">查看更多</a>
        </div>
      </div>
    </div>
    <div id="footer-bottom">
      <div class="footer-bottom-content">
        <div class="footer-bottom-left">
          <span>Copyright © 2023 By
            <a target="_blank" rel="noopener" href="http://www.blog.zjh2002.icu" title="点击访问Monkey-PaPa的主页" style="margin-left: 5px">Monkey-PaPa
            </a>
          </span>
          <div>
            <a class="footer-bottom-link" target="_blank" href="https://cloud.tencent.com/" title="腾讯云提供图片存储和网站托管">腾讯云</a>
            <a class="footer-bottom-link" target="_blank" href="https://tongji.baidu.com/" title="百度统计提供用户访问统计">百度统计</a>
            <a class="footer-bottom-link" target="_blank" href="https://vercel.com/" title="Vercel提供镜像站托管">Vercel</a>
            <a class="footer-bottom-link" target="_blank" href="https://github.com/" title="Github提供源码托管">Github</a>
          </div>
        </div>
        <div class="footer-bottom-right">
          <div id="runtime" title="本站已运行 1.10 坤年🐔🏀">
            本站已运行：{{ countdownChange }}
          </div>
          <div>
            <a class="footer-bottom-link" target="_blank" href="https://v2.cn.vuejs.org/v2/guide/" title="框架">Vue</a>
            <a class="footer-bottom-link" target="_blank" href="https://docs.djangoproject.com/zh-hans/5.0/" title="框架">Django</a>
            <a class="footer-bottom-link" href="javascript:;" title="备案号">京cpu备5201314号</a>
          </div>
        </div>
      </div>
    </div>
  </footer>
</template>

<script>
export default {
  data() {
    return {
      countdownTime: '2023-08-13 00:00:00',
      countdownChange: ''
    }
  },
  mounted() {
    setInterval(() => {
      this.countdown()
    }, 1000)
  },
  methods: {
    copy(val) {
      const number = val === 1 ? '1816298537' : 'z15523692545'
      const text = val === 1 ? 'QQ' : '微信'
      const aux = document.createElement('input')
      aux.setAttribute('value', number)
      document.body.appendChild(aux)
      aux.select()
      document.execCommand('copy')
      document.body.removeChild(aux)
      this.$notify({
        title: '可以啦🍨',
        message: `本博主的${text}已经到你的剪贴板啦，快加入我们吧~~🎉`,
        type: 'success',
        offset: 50,
        position: 'top-left'
      })
    },
    countdown() {
      if (this.$common.isEmpty(this.countdownTime)) {
        return
      }
      let countdown = this.$common.countdown(this.countdownTime)
      this.countdownChange = countdown.d + '天' + countdown.h + '时' + countdown.m + '分' + countdown.s + '秒'
      this.countdownChange = this.countdownChange.replace(/-/g, ' ')
    }
  }
}
</script>

<style lang="scss" scoped>
#footer_icons,
#footer_content {
  margin: auto;
  background: var(--background);
}
#footer_icons {
  padding-top: 2rem;
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  justify-content: center;
  > div {
    display: flex;
  }
  img.footer_logo {
    border-style: none;
    border-radius: 50%;
    width: 60px;
    height: 60px;
    margin: 0 1rem;
    filter: drop-shadow(0 0 12px var(--blue22)) !important;
    transition: all 0.25s;
    &:hover {
      transform: scale(1.1);
      transition-duration: 0.4s;
    }
  }
  a {
    text-decoration: none;
    word-wrap: break-word;
    transition: all 0.2s;
    overflow-wrap: break-word;
  }
  .icon_link {
    height: 40px;
    width: 40px;
    display: flex;
    align-items: center;
    justify-content: center;
    margin: 1rem;
    background-color: var(--fontColor);
    border-radius: 50%;
    i {
      font-size: 20px;
      color: var(--favoriteBg);
    }
    &:hover {
      background: var(--red);
    }
  }
}
#footer_content {
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
  padding: 0 10rem 2rem;
  .footer-group {
    min-width: 120px;
    text-align: center;
  }
  .footer-title {
    color: var(--fontColor);
  }
  .footer-links {
    display: flex;
    flex-direction: column;
  }
  .footer-item {
    font-size: 1rem;
    line-height: 1.7;
    color: var(--fontColor);
    &:hover {
      color: var(--red);
    }
  }
}
#footer-bottom {
  padding: 1rem;
  background: var(--background);
  z-index: 2;
  .footer-bottom-content {
    display: flex;
    justify-content: space-between;
    max-width: 1100px;
    width: 100%;
    margin: 0 auto;
    flex-wrap: wrap;
  }
  .footer-bottom-left,
  .footer-bottom-right {
    display: flex;
    flex-direction: column;
  }
  .footer-bottom-content * {
    font-size: 17px;
    white-space: nowrap;
    color: var(--fontColor);
  }
  .footer-bottom-left div,
  .footer-bottom-right div,
  .footer-bottom-left span,
  .footer-bottom-right span {
    display: flex;
    align-items: center;
  }
  a {
    display: flex;
    align-items: center;
    &:not(:first-child)::before {
      content: '|';
      font-size: 14px;
      line-height: 2.45;
      margin: 0 10px;
    }
    &:hover {
      color: var(--red);
      &:not(:first-child)::before {
        color: var(--fontColor);
      }
    }
  }
  .footer-bottom-right div {
    justify-content: flex-end;
  }
}
@media screen and (max-width: 768px) {
  #footer-bottom {
    padding: 1rem 0 !important;
    .footer-bottom-content * {
      font-size: 1rem;
    }
    .footer-bottom-left,
    .footer-bottom-right {
      width: 100%;
      align-items: center;
    }
  }
  #footer_icons img.footer_logo {
    display: none;
  }
  #footer_content {
    padding: 1rem;
    .footer-group {
      text-align: center;
    }
    #friend-links-in-footer h3 {
      padding-left: 10px;
    }
  }
}
</style>
