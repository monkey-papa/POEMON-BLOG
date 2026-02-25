<template>
  <div class="about-container">
    <!-- 首页图片 -->
    <div
      style="animation: header-effect 2s"
      :style="{ background: `${changeBgState}` }"
      class="background-image background-image-changeBg"
    ></div>
    <!-- 标签区域 -->
    <div class="main">
      <div class="layout">
        <div class="content-main">
          <div class="author-main">
            <div class="author-tag-left">
              <span class="author-tag">💻 Like数码科技</span>
              <span class="author-tag">🥣 干饭人 干饭魂</span>
              <span class="author-tag">🕊 咕咕咕咕咕咕~</span>
              <span class="author-tag">🧱 CV工程师</span>
            </div>
            <div class="author-box myCenter">
              <span></span>
              <div class="author-img">
                <img :src="webInfo.avatar" alt="博主头像" />
              </div>
            </div>
            <div class="author-tag-right">
              <span class="author-tag">吃饭不如碎觉 💤</span>
              <span class="author-tag">乐观 积极 向上 🤝</span>
              <span class="author-tag">专攻各种困难 🔨</span>
              <span class="author-tag">人不狠话超多 💢</span>
            </div>
            <div class="image-dot"></div>
          </div>
          <div class="author-title">关于本站</div>
          <div class="author-content">
            <div class="author-content-item myInfoAndSayHello">
              <div class="title1">你好~♪，很高兴认识你👋</div>
              <div class="title2">
                网名
                <span class="inline-word">{{ webInfo.webName }}</span>
              </div>
              <div class="title1">
                是一只还在求学路上的
                <span class="inline-word">大学生、个人博主</span>
              </div>
            </div>
            <div class="aboutsiteTips author-content-item">
              <div class="author-content-item-tips">追求</div>
              <h2>
                源于<br />热爱而去感受
                <span class="inline-word"></span>
              </h2>
            </div>
          </div>
          <div class="author-content">
            <div class="about-statistic author-content-item">
              <div class="card-content">
                <div class="author-content-item-tips">数据</div>
                <span class="author-content-item-title">访问统计</span>
                <div id="statistic">
                  <div>
                    <span>今日人数:</span>
                    <span class="num">{{ user_today }}</span>
                  </div>
                  <div>
                    <span>今日访问:</span>
                    <span class="num">{{ today_sum }}</span>
                  </div>
                  <div>
                    <span>昨日人数:</span>
                    <span class="num">{{ last_user }}</span>
                  </div>
                  <div>
                    <span>昨日访问:</span>
                    <span class="num">{{ yesterday_sum }}</span>
                  </div>
                  <div>
                    <span>总访问量:</span>
                    <span class="num">{{ total_sum }}</span>
                  </div>
                </div>
                <div class="post-tips">统计信息来自 <a href="javascript:;">博客后台统计</a></div>
              </div>
            </div>
            <div class="author-content-item-group column mapAndInfo">
              <div class="author-content-item map single">
                <span class="map-title">我现在住在 <b>中国，包头市</b></span>
              </div>
              <div class="author-content-item selfInfo myBetween single">
                <div>
                  <span class="selfInfo-title">生于</span>
                  <span class="year selfInfo-content">2002</span>
                </div>
                <div>
                  <span class="selfInfo-title">内蒙古科技大学</span>
                  <span class="software selfInfo-content">软件工程</span>
                </div>
                <div>
                  <span class="selfInfo-title">职业（2023年）</span>
                  <span class="job selfInfo-content">大三学生 👨‍🎓</span>
                </div>
                <div>
                  <span class="selfInfo-title">工作</span>
                  <span class="work selfInfo-content">金山前端打工人 🧑‍💻</span>
                </div>
              </div>
            </div>
          </div>
          <div class="author-content">
            <div class="author-content-item maxim">
              <div class="author-content-item-tips">个性签名</div>
              <span class="maxim-title">
                <span class="maxim-title-text">天命，</span>
                <span>不可违。</span>
              </span>
            </div>
            <div class="author-content-item myphoto">
              <span class="map-title">我的生活 <b>吐槽</b><sup>闲话</sup></span>
            </div>
          </div>
          <div class="author-content accordion">
            <div class="author-content-item chuckle-img myCenter">
              <div class="container">
                <div
                  v-for="(item, i) in images"
                  :key="i"
                  class="item"
                  @click="btnClick(i)"
                  :class="{ active: current == i }"
                  :style="{ backgroundImage: `url(${item.img})` }"
                >
                  <div class="shadow"></div>
                  <div class="content">
                    <div class="icon myCenter">
                      <i class="fa fa-image"></i>
                    </div>
                    <div class="text">
                      <div class="tit">{{ item.tit }}</div>
                      <div class="sub">{{ item.sub }}</div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="author-content">
            <div class="author-content-item game-lol">
              <div class="card-content">
                <div class="content-bottom">
                  <div class="icon-group">
                    <i class="icon-pos-mid"></i>
                    <i class="icon-pos-dy"></i>
                    <i class="icon-pos-sup"></i>
                  </div>
                  <div class="tips">手Q111区 黑猫糖果</div>
                </div>
              </div>
            </div>
            <div class="author-content-item game-wolf">
              <div class="card-content">
                <div class="author-content-item-tips">爱好影视</div>
                <span class="author-content-item-title">The Walking Dead</span>
              </div>
            </div>
          </div>
          <div class="author-content">
            <div class="author-content-item like-technology">
              <div class="card-content">
                <div class="author-content-item-tips">关注偏好</div>
                <span class="author-content-item-title">花里胡哨产品</span>
                <div class="content-bottom myBetween">
                  <div class="tips">手机、电脑软硬件，退役搞机er</div>
                </div>
              </div>
            </div>
            <div class="author-content-item like-music">
              <div class="card-content">
                <div class="author-content-item-tips">音乐偏好</div>
                <span class="author-content-item-title">Rap、华语流行</span>
                <span class="author-content-item-changjian">关注 Jay、马思唯、法老……</span>
                <div class="content-bottom myBetween">
                  <div class="tips music-tips">和 {{ webInfo.webName }} 一起欣赏更多音乐</div>
                </div>
                <div class="banner-button-group">
                  <a :href="$constant.siteURL" class="banner-button myCenter">
                    <i class="fas fa-circle-plus"></i>
                    <span class="banner-button-text">我的音乐</span>
                  </a>
                </div>
              </div>
            </div>
          </div>
          <div class="about-content-item-group author-content">
            <div class="create-site-post author-content-item single">
              <div class="author-content-item-tips">全军复诵</div>
              <span class="author-content-item-title">天罡赋-唐·袁天罡</span>
              <br /><br />
              <p>
                世间万物，风云变幻，苍黄翻覆。纵使波谲云诡，但制心一处，便无事不办。天定胜人，人定兮胜天。
              </p>
              <p>
                天下尽做饵，唯本帅执竿。三百年光阴，京师长安到东都洛阳兴亡起伏，从鼓动杨广大兴师旅到造就贞观之治，渭水之盟到大败突厥、废王立武到二圣临朝，谋天宝之乱到纵黄巢造反警示昭宗。臣无不是为大唐，这其中多少苦心，多少苦难，本帅始终不曾停下脚步，却因为你，踌躇了数十年！
              </p>
              <p>
                乱世数十载，以致当下之势。前有朱贼父子篡权，中有沙陀李鸦儿，岐地宋文通野心颇大，尤其那李嗣源更甚，我大唐世道，何时竟成了如此模样。
              </p>
              <p>一直以来，本帅才是那个执棋者，才是那个造局人。</p>
              <p>
                玄宗当年是何等雄武之才，开元之际，励精政事，遂袭大统，几致太平，却沉迷于女人，宴安酖毒，府卫崩坏，穷尽天下之欲而不足为其乐。
              </p>
              <p>高祖，太宗，高总，玄宗，代，德，顺，武，宣，懿，僖，昭，诸君，臣，尽力了！</p>
            </div>
          </div>
          <div class="about-content-item-group author-content">
            <div class="create-site-post author-content-item single">
              <span class="disclaimer author-content-item-title">🤷‍♀️免责声明</span>
              <br /><br />
              <p>
                本站以分享互联网经验、学习知识为目的，所有文章所涉及使用的工具、资源等均来自互联网，
                仅供学习和研究使用，版权归作者所有，如果无意之中侵犯了您的版权，请来信告知。本站将在第一时间删除！另外，
                本站内的文章多为博主原创，仅供学习交流之用，不参与商业用途。
              </p>
              <p>
                遵守相关法律法规，由于本站资源部分来源于网络，开发也是使用开源框架，故无法核实资源侵权的真实性，无论出于何种目的要求本站删除内容，“您”均需要提供相关证明，否则不予处理。
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref, computed } from "vue";
import { useStore } from "@/stores";
import { useGlobalProperties } from "@/composables/useGlobalProperties";

const store = useStore();
const { $constant } = useGlobalProperties();

const changeBgState = computed(() => store.changeBg);
const webInfo = computed(() => store.webInfo);
const pageView = computed(() => store.pageView);

const current = ref<number>(0);
const images = ref($constant.about);
const total_sum = ref<number>(pageView.value?.total_sum || 0);
const yesterday_sum = ref<number>(pageView.value?.yesterday_sum || 0);
const today_sum = ref<number>(pageView.value?.today_sum || 0);
const user_today = ref<number>(pageView.value?.data?.today?.users?.length || 0);
const last_user = ref<number>(pageView.value?.data?.yesterday?.users?.length || 0);

const btnClick = (i: number): void => {
  current.value = i;
};
</script>
<style lang="scss" scoped>
.about-container {
  background: var(--background);
  .main {
    margin: 0 auto;
    padding: 60px 100px;
    min-height: 100vh;
    letter-spacing: 0.7px;
    overflow-wrap: break-word;

    .layout {
      display: flex;
      max-width: 1300px;
    }

    .content-main {
      width: 100%;
      height: 100%;
      font-weight: 700;
      padding: 10px 35px;
      display: flex;
      flex-direction: column;
      align-items: center;
    }

    .author-main {
      position: relative;
      display: flex;
      align-items: center;
      user-select: none;

      .author-tag-left {
        display: flex;
        flex-direction: column;
        align-items: flex-end;
        margin-right: 30px;

        .author-tag {
          &:first-child,
          &:last-child {
            margin-right: -16px;
          }
        }
      }

      .author-tag-right {
        display: flex;
        flex-direction: column;
        align-items: flex-start;
        margin-left: 30px;

        .author-tag {
          &:first-child,
          &:last-child {
            margin-left: -16px;
          }
        }
      }

      .author-tag {
        color: var(--darkBlue);
        text-align: center;
        width: 100%;
        height: 32px;
        line-height: 32px;
        transform: translate(0, -4px);
        padding: 1px 8px;
        background: var(--favoriteBg);
        border: 1px solid var(--gray1);
        border-radius: 40px;
        margin-top: 6px;
        font-size: 14px;
        font-weight: 400;
        animation: floating 6s ease-in-out 0s infinite normal none running;

        &:nth-child(1) {
          animation-delay: 0s;
        }
        &:nth-child(2) {
          animation-delay: 0.6s;
        }
        &:nth-child(3) {
          animation-delay: 1.2s;
        }
        &:nth-child(4) {
          animation-delay: 1.8s;
        }

        &:hover {
          color: var(--black8);
          border-color: var(--gray4);
        }
      }

      .author-box {
        position: relative;
        width: 189px;
        height: 189px;
        background: var(--favoriteBg);
        border-radius: 50%;
        overflow: hidden;

        &::before {
          content: "";
          position: absolute;
          width: 500px;
          height: 500px;
          background-image: conic-gradient(transparent, transparent, transparent, var(--purple1));
          animation: rotate 4s linear infinite;
          animation-delay: -2s;
        }

        &::after {
          content: "";
          position: absolute;
          width: 500px;
          height: 500px;
          background-image: conic-gradient(transparent, transparent, transparent, var(--blue2));
          animation: rotate 4s linear infinite;
        }

        span {
          position: absolute;
          inset: 5px;
          border-radius: 50%;
          background: var(--favoriteBg);
          z-index: 1;
        }
      }

      .author-img {
        margin: auto;
        border-radius: 50%;
        overflow: hidden;
        width: 180px;
        height: 180px;
        z-index: 10;
        background: var(--maxMaxWhiteMask);

        img {
          border-radius: 11px;
          margin-right: 4px;
          display: block;
          margin: 0 auto 20px;
          max-width: 100%;
        }
      }

      .image-dot {
        width: 30px;
        height: 30px;
        background: var(--green1);
        position: absolute;
        border-radius: 50%;
        border: 6px solid var(--favoriteBg);
        top: 50%;
        left: 50%;
        z-index: 20;
        transform: translate(51px, 54px);
      }
    }

    .author-title {
      width: 210px;
      margin-top: 10px;
      font-size: 2.7rem;
      font-weight: 700;
      letter-spacing: 6px;
      -webkit-background-clip: text;
      background-clip: text;
      background-image: linear-gradient(
        90deg,
        var(--blue2),
        var(--blue1) 24%,
        var(--purple5) 45%,
        var(--red4) 80%,
        var(--bigRed4)
      );
      display: inline-block;
      text-align: center;
      color: transparent;
      -webkit-text-stroke: 1px var(--blue10);
    }

    .author-content {
      display: flex;
      flex-wrap: wrap;
      justify-content: space-between;
      width: 100%;
      margin-top: 1rem;

      &.about-content-item-group {
        color: var(--darkBlue);
      }

      .author-content-item {
        width: 49%;
        border-radius: 24px;
        background: var(--maxMaxWhiteMask);
        border: 1px solid var(--maxMaxWhiteMask);
        box-shadow: 0 8px 16px -4px var(--mask);
        position: relative;
        padding: 1.5rem 2.5rem;
        overflow: hidden;

        .author-content-item-tips {
          opacity: 0.8;
          font-size: 0.8rem;
          margin-bottom: 0.5rem;
          font-weight: 400;
        }

        .author-content-item-title {
          margin-bottom: 0.5rem;
          font-size: 36px;
          font-weight: 700;
          line-height: 1;
          font-family: STZhongsong, "Microsoft YaHei";

          &.disclaimer {
            font-size: 25px;
          }
        }

        .card-content {
          position: absolute;
          width: 100%;
          height: 100%;
          top: 0;
          left: 0;
          z-index: 2;
          display: flex;
          flex-direction: column;
          padding: 1rem 2rem;

          .content-bottom {
            margin-top: auto;

            .icon-group {
              display: flex;

              i {
                &.icon-pos-mid {
                  background: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABYAAAASCAYAAABfJS4tAAAACXBIWXMAAAsTAAALEwEAmpwYAAAAkElEQVQ4ja3UzQ2AIAyG4RfjySVcxc3UzVjFJbziRfxBflrkS0hIkz4H2mCccwCwbwvy2PMAzMDdO4wGoFNgZfQRLSxCAfpEPdlQQFd/qXmKHHrVtbAI1cISdNLCUlQFq9EUbJWoh18JYUt5Tz+DiuUJN0NjcBM0hJuhIdwMhfdfUUInItNPJbfH1WgO/oUCHK48MCZZLMkWAAAAAElFTkSuQmCC);
                }

                &.icon-pos-dy {
                  background: url(data:image/webp;base64,UklGRswBAABXRUJQVlA4WAoAAAAQAAAAFQAAEQAAQUxQSPQAAAAB5+KgjSRHuuqecOGfP92+QxEReRjxKaTEqw/FxMEqUlbowVdnS4wRppSK4vZFo2B12L+nRUa8gFjbtmlnfTvOR2zbdr5t7P7reE5KiOh/Yvs3kzhy50G8Tu6ww41HPGebW1gPgxiuKFa+hDmvV9y3F5iX2sB40h7ubT1j7KtIZtBLY6SvcFzpgfRImZwaASLNGxxzfZ3S0gu/dYy3QRxnQd3Mlzqh5XuWdIfrWka5JOMYH5/awXUoxztGsjvM4loZy66F2dVwAVe4JLuzyJkqUdzTVVmTdT6aU3glK7LWOFrFO1WV0Z7h70Re3/wn9nOABwAAVlA4ILIAAACwBQCdASoWABIAAAAAJZjbAPwAXID+AA1AroH4AfqrTgP4B+IG9/5wB/QMoA/UCVAAwo6eH9AA/v8e3//yV///k4P/kq3//8Kn+dRaGe4yayRzyHOnBYsf//uqmBUxxKdq/lxND/+dyv//wsr5HXg/m3J5HXk/+n1f/+GE/+qX7/8Th/9Ln/Pn37ZD/lzRvwhP/oef/nMg9DkxTTJVX8O4uvyTcpduZNQv6DQwybUAAAAA);
                }

                &.icon-pos-sup {
                  background: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABYAAAASCAYAAABfJS4tAAAACXBIWXMAAAsTAAALEwEAmpwYAAABZUlEQVQ4ja3TPUsdYRAF4GejXYSA2NgatQhYCl4Uy0BSGFPbClqIEvz4DeIHgtaSKoVlmggRUqQTOyFBRNFKG5WAiiBetNi5ZN1s7l3BgZedOXPm7O68M8mfk10NrAVLSCK+xywu6xU1N1LFBMZy2DHm6xW9KCH8uiT2ZOFxbGTijcDqWnO8fV3ay7x9wQpGIk7Cr+JTBs/aDUaTuLxefEV7AXEOi2iKuCq9vIUC7ik+YrvWih30YBVXOfJMiFbjNAWWtSushcY2j3t8jim8wjBuA/8sbdN0nJbABGc4aiZDAyQFczyIb3gpndkudOB75N/iCPvSnl/jPX5mRYqm4gy/wv+BQ1Qy+QoOIie4Z3mRogX5jX7pb+8F1pfJ1/w1bGEZd3mRolb8w4kvao34Am3SNv3XyizIm4yo8LsbFZUR7ivABp5DuFISe7JwVwHW+RzCQ/4uhPA/NCoqMxU1exfPzTLkBzctS76SprNHAAAAAElFTkSuQmCC);
                }
              }
            }

            .tips {
              &.music-tips {
                margin-bottom: 27px;
              }
            }
          }

          .banner-button-group {
            position: absolute;
            bottom: 1.5rem;
            right: 2rem;

            .banner-button {
              height: 40px;
              width: 124px;
              border-radius: 20px;
              background: var(--favoriteBg);
              text-decoration: none;
              z-index: 1;
              transition: all 0.3s ease;

              &:hover {
                text-decoration: none;
                background: linear-gradient(to bottom right, var(--blue5), var(--blue2));
                .banner-button-text {
                  color: var(--white);
                }
              }

              .banner-button-text {
                color: var(--orange1);
              }
            }
          }
        }

        &.myInfoAndSayHello {
          display: flex;
          flex-direction: column;
          justify-content: center;
          color: var(--white);
          background: linear-gradient(130deg, var(--blue8) 0, var(--blue9) 80%);
          background-size: 200%;
          animation: gradient 15s ease infinite;
          width: 59%;

          .title1 {
            opacity: 0.8;
            line-height: 1.3;
          }

          .title2 {
            font-size: 36px;
            font-weight: 700;
            line-height: 1.1;
            margin: 0.5rem 0;
            font-family: STZhongsong, "Microsoft YaHei";
          }

          .title3 {
            opacity: 0.8;
            line-height: 2;
            font-size: 10px;
          }
        }

        &.aboutsiteTips {
          display: flex;
          justify-content: center;
          align-items: flex-start;
          flex-direction: column;
          width: 39%;

          h2 {
            margin-right: auto;
            font-size: 36px;
            font-family: Helvetica;
            line-height: 1.06;
            letter-spacing: -0.02em;
            color: var(--black5);
            margin-top: 0;
          }
        }

        &.about-statistic {
          min-height: 380px;
          width: 39%;
          background: url(https://p.zhheo.com/l7Bfmg24690881724048986403.webp!cover) no-repeat top;
          background-size: cover;
          color: var(--favoriteBg);
          overflow: hidden;
        }

        &.single {
          width: 100%;
        }

        &.map {
          background: no-repeat center; /* 请替换为你自己的地图背景图 */
          min-height: 200px;
          max-height: 400px;
          position: relative;
          overflow: hidden;
          margin-bottom: 0.5rem;
          height: 60%;
          border: 1px;
          background-size: 100%;
          transition: 1s ease-in-out;

          &:hover {
            background-size: 120%;
            transition: 3s ease-in-out;
            background-position-x: 80%;
            background-position-y: 36%;

            .map-title {
              bottom: -100%;
            }
          }

          .map-title {
            position: absolute;
            bottom: 0;
            left: 0;
            width: 100%;
            background: var(--maxMaxWhiteMask);
            padding: 0.5rem 2rem;
            backdrop-filter: saturate(180%) blur(2px);
            -webkit-backdrop-filter: blur(20px);
            transition: 1s ease-in-out;
            font-size: 20px;

            b {
              color: var(--black5);
              font-weight: bolder;
              -webkit-box-sizing: border-box;
              -moz-box-sizing: border-box;
              box-sizing: border-box;
            }
          }
        }

        &.selfInfo {
          min-height: 100px;
          max-height: 400px;
          flex-wrap: wrap;
          height: -webkit-fill-available;
          height: 40%;

          div {
            display: flex;
            flex-direction: column;
            margin: 0.5rem 2rem 0.5rem 0;
          }

          .selfInfo-title {
            opacity: 0.8;
            font-size: 0.6rem;
            line-height: 1;
            margin-bottom: 8px;
            font-weight: 400;
            color: var(--red);
          }

          .selfInfo-content {
            font-weight: 700;
            font-size: 34px;
            line-height: 1;

            &.year {
              color: var(--blue1);
              font-family: "Microsoft YaHei";
            }

            &.software {
              color: var(--orange1);
            }

            &.job {
              color: var(--purple1);
            }

            &.work {
              color: var(--red);
              font-size: 16px;
            }
          }
        }

        &.maxim {
          font-size: 36px;
          font-weight: 700;
          line-height: 1.1;
          display: flex;
          color: var(--favoriteBg);
          background: no-repeat; /* 请替换为你自己的图片 */
          background-size: cover;
          align-items: flex-start;
          flex-direction: column;
          justify-content: center;
          width: 53%;

          .maxim-title {
            display: flex;
            flex-direction: column;
            font-family: STZhongsong, "Microsoft YaHei";

            .maxim-title-text {
              opacity: 0.6;
              margin-bottom: 8px;
            }
          }
        }

        &.myphoto {
          background: no-repeat; /* 请替换为你自己的图片 */
          height: 60%;
          background-size: 100%;
          transition: 2s ease-in-out;
          position: relative;
          min-height: 200px;
          max-height: 400px;
          border: 1px;
          overflow: hidden;
          width: 45%;

          &:hover {
            background-size: 150%;
            transition: 1.5s ease-in-out;
            background-position-x: 60%;
            background-position-y: 20%;

            .map-title {
              bottom: -100%;
            }
          }

          .map-title {
            position: absolute;
            bottom: 0;
            left: 0;
            width: 100%;
            background: var(--maxMaxWhiteMask);
            padding: 0.5rem 2rem;
            backdrop-filter: saturate(180%) blur(2px);
            -webkit-backdrop-filter: blur(20px);
            transition: 1s ease-in-out;
            font-size: 20px;

            b {
              color: var(--black5);
              font-weight: bolder;
              -webkit-box-sizing: border-box;
              -moz-box-sizing: border-box;
              box-sizing: border-box;
            }

            sup {
              top: -0.5em;
              font-size: 75%;
              line-height: 0;
              position: relative;
              vertical-align: baseline;
            }
          }
        }

        &.game-lol {
          background: no-repeat top; /* 请替换为你自己的图片 */
          background-size: cover;
          min-height: 300px;
          overflow: hidden;
          color: var(--white);
          width: 42%;
          border: 1px;

          .card-content {
            .content-bottom {
              .icon-group {
                i {
                  display: inline-block;
                  width: 22px;
                  height: 18px;
                  margin-right: 0.5rem;
                }
              }
            }
          }
        }

        &.game-wolf {
          border: 1px;
          width: 56%;
          background: no-repeat top; /* 请替换为你自己的图片 */
          background-size: cover;
          min-height: 300px;
          overflow: hidden;
          color: var(--white);
        }

        &.like-technology {
          border: 1px;
          background: no-repeat; /* 请替换为你自己的图片 */
          background-size: cover;
          min-height: 230px;
          color: var(--red);
        }

        &.like-music {
          border: 1px;
          background: no-repeat top; /* 请替换为你自己的图片 */
          background-size: cover;
          min-height: 400px;
          color: var(--orange1);
          overflow: hidden;
        }

        .author-content-item-changjian {
          font-size: 16px;
          font-weight: 600;
          line-height: 1;
        }

        .inline-word {
          word-break: keep-all;
          white-space: nowrap;
        }

        sup {
          top: -0.5em;
          font-size: 75%;
          line-height: 0;
          position: relative;
          vertical-align: baseline;
        }

        p {
          line-height: 1.7;
          font-weight: 400;
          text-align: left;
          letter-spacing: 0.7px;
          margin: 0 0 16px;
        }
      }

      .author-content-item-group {
        &.column {
          display: flex;
          flex-direction: column;
          justify-content: space-between;

          &.mapAndInfo {
            width: 59%;
          }
        }
      }

      #statistic {
        font-size: 16px;
        border-radius: 15px;
        width: 100%;
        color: var(--white);
        display: flex;
        justify-content: space-between;
        flex-direction: row;
        flex-wrap: wrap;
        margin-top: 1rem;
        margin-bottom: 2rem;

        div {
          display: flex;
          justify-content: space-between;
          flex-direction: column;
          width: 50%;
          margin-bottom: 0.5rem;

          span {
            &:first-child {
              opacity: 0.8;
              font-size: 0.6rem;
            }

            &:last-child {
              font-weight: 700;
              font-size: 34px;
              line-height: 1;
              white-space: nowrap;
            }
          }
        }
      }

      .post-tips {
        color: var(--maxMaxWhiteMask);
        font-size: 16px;
        position: absolute;
        bottom: 1rem;
        left: 2rem;

        a {
          color: var(--red);
          padding: 0 0.2em;
          font-weight: 700;
          transition: all 0.3s;

          &:hover {
            background-color: var(--red);
            border-radius: 5px;
            color: var(--white);
            padding: 0.4em 0.2em;
            text-decoration: none;
            font-weight: 700;
          }
        }
      }

      &.accordion {
        .chuckle-img {
          height: auto;
          width: 100%;
          background: radial-gradient(circle at top center, var(--black6), var(--black7));

          .container {
            display: flex;
            width: 90vw;
            max-width: auto;
            height: 350px;
            overflow: hidden;
          }

          .item {
            position: relative;
            width: 60px;
            margin: 10px;
            border-radius: 30px;
            background-size: cover;
            background-position: center;
            transition: 0.5s cubic-bezier(0.05, 0.61, 0.41, 0.95);
            overflow: hidden;

            &.active {
              flex: 1;
              margin: 0;
              border-radius: 40px;

              .shadow {
                background: linear-gradient(to top, var(--mask) 65%, transparent);
              }

              .content {
                bottom: 20px;
                left: 20px;

                .text {
                  div {
                    opacity: 1;
                  }
                }
              }
            }

            .shadow {
              position: absolute;
              bottom: 0;
              left: 0;
              right: 0;
              height: 100px;
              transition: 0.5s cubic-bezier(0.05, 0.61, 0.41, 0.95);
            }

            .content {
              display: flex;
              position: absolute;
              left: 10px;
              right: 0;
              bottom: 10px;
              height: 40px;
              transition: 0.5s cubic-bezier(0.05, 0.61, 0.41, 0.95);

              .icon {
                min-width: 40px;
                width: 40px;
                height: 40px;
                border-radius: 50%;
                background-color: var(--white);
                font-size: 16px;
              }

              .text {
                display: flex;
                flex-direction: column;
                justify-content: center;
                margin-left: 10px;
                color: var(--white);
                width: 100%;

                .tit {
                  font-weight: bold;
                  font-size: 18px;
                }

                div {
                  width: calc(100% - 70px);
                  text-overflow: ellipsis;
                  overflow: hidden;
                  white-space: nowrap;
                  line-height: 20px;
                  opacity: 0;
                  transition: opacity 0.5s ease-out;
                }

                .sub {
                  transition-delay: 0.1s;
                }
              }
            }

            &:nth-child(1) .fa {
              color: var(--red4);
            }
            &:nth-child(2) .fa {
              color: var(--yellow7);
            }
            &:nth-child(3) .fa {
              color: var(--green5);
            }
            &:nth-child(4) .fa {
              color: var(--blue4);
            }
            &:nth-child(5) .fa {
              color: var(--purple5);
            }
            &:nth-child(6) .fa {
              color: var(--black);
            }
          }
        }
      }
    }
  }
}

@keyframes floating {
  0% {
    transform: translate(0, -4px);
  }
  50% {
    transform: translate(0, 4px);
  }
  100% {
    transform: translate(0, -4px);
  }
}

@media screen and (max-width: 1286px) {
  .about-container {
    .main {
      padding: 60px 5px;

      .content-main {
        padding: 10px 0;
      }

      .author-content {
        .author-content-item {
          &.myphoto {
            background-size: cover;
          }
        }
      }
    }
  }
}

@media screen and (max-width: 900px) {
  .about-container {
    .main {
      .author-content {
        &.accordion {
          margin: 0;

          .chuckle-img {
            display: none;
          }
        }
      }
    }
  }
}

@media screen and (max-width: 768px) {
  .about-container {
    .main {
      .author-content {
        .author-content-item {
          &.like-music,
          &.aboutsiteTips,
          &.myphoto,
          &.game-wolf {
            margin-top: 16px;
          }

          &,
          &.myInfoAndSayHello,
          &.aboutsiteTips,
          &.maxim,
          &.myphoto,
          &.game-lol,
          &.game-wolf,
          &.about-statistic {
            width: 100%;
          }

          &.single {
            margin-top: 16px;
            margin-bottom: 0;
            width: 170%;

            &.create-site-post {
              margin: 0;
            }
          }
        }
      }
    }
  }
}

@media screen and (max-width: 550px) {
  .about-container {
    .main {
      .author-main {
        .author-tag-left,
        .author-tag-right {
          display: none;
        }
      }
    }
  }
}
</style>
