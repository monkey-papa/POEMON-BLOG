<template>
  <div class="aside-container">
    <!-- 网站信息 -->
    <div class="aside-card-content shadow-box">
      <div class="author-box myCenter">
        <span></span>
        <div class="author-img">
          <img src="/title.jpg" alt="站点封面" />
        </div>
      </div>
      <div class="image-dot"></div>
      <div class="web-name">{{ webInfo.webName }}</div>
      <div class="web-info">
        <div class="blog-info-box mySpaceAround">
          <span>文章</span>
          <span class="blog-info-num">{{ articleTotal }}</span>
        </div>
        <div class="blog-info-box mySpaceAround">
          <span>分类</span>
          <span class="blog-info-num">{{ sortInfo.length }}</span>
        </div>
        <div class="blog-info-box mySpaceAround">
          <span>访问量</span>
          <span class="blog-info-num">{{ total_sum }}</span>
        </div>
      </div>
      <a class="collection-btn myCenter" @click="showTip()">
        <el-icon size="20" class="icon-loading"><Loading /></el-icon>微言圈
      </a>
      <div class="aside-link-box a-svg myBetween">
        <a href="https://blog.csdn.net/qq_53461589?spm=1038.2274.3001.5343" target="_blank">
          <i class="iconfont icon-csdn myCenter"></i>
        </a>
        <a
          href="https://www.bilibili.com/video/BV1nh4y1m7CV/?spm_id_from=333.999.0.0"
          target="_blank"
        >
          <i class="iconfont icon-wf_Bzhan myCenter"></i>
        </a>
        <a href="https://github.com/your-username/POEMON-BLOG" target="_blank">
          <i class="iconfont icon-github myCenter"></i>
        </a>
        <a :href="$constant.siteURL">
          <i class="iconfont icon-wangyiyunyinle myCenter"></i>
        </a>
      </div>
    </div>
    <!-- 天气时间 -->
    <div v-if="route.path !== '/article'" class="wow card-widget card-widget-weather">
      <div class="card-content-title">
        <i class="fa fa-thermometer-3 card-content-icon"></i>
        <span>小窝天气</span>
      </div>
      <div class="weather-content">
        <div class="weather-content-inner">
          <div class="mk-address">
            <div class="city">{{ city }}</div>
          </div>
          <div class="mk-Today_weather myBetween">
            <div class="left">
              <div class="weather-icon">
                <img
                  v-if="weather?.type"
                  class="iconfont"
                  :src="getWeatherIconUrl(weather?.type)"
                />
              </div>
              <div class="temp">
                <div class="number">
                  {{ (weather?.low || "") + "~" + (weather?.high || "") }}
                </div>
                <div class="weather">
                  <span>{{ weather?.type || "" }}</span>
                  <span class="weather-divider" v-if="weather?.type">|</span>
                  <span v-if="weather?.fengli">风力：{{ weather?.fengli }}</span>
                </div>
              </div>
            </div>
            <div class="right">{{ weather?.fengxiang || "" }}</div>
          </div>
          <div class="mk-week_weather">
            <div class="tomorrow">
              <div class="text">明天</div>
              <div class="icon">
                <img
                  v-if="tomWeather?.type"
                  class="iconfont"
                  :src="getWeatherIconUrl(tomWeather?.type)"
                />
              </div>
              <div class="temp">
                {{ (tomWeather?.low || "") + "~" + (tomWeather?.high || "") }}
              </div>
            </div>
            <div class="afterTomorrow">
              <div class="text">后天</div>
              <div class="icon">
                <img
                  v-if="afterTomWeather?.type"
                  class="iconfont"
                  :src="getWeatherIconUrl(afterTomWeather?.type)"
                />
              </div>
              <div class="temp">
                {{ (afterTomWeather?.low || "") + "~" + (afterTomWeather?.high || "") }}
              </div>
            </div>
            <div class="afterAfterTomWeather">
              <div class="text">{{ afterAfterTomWeather?.week || "" }}</div>
              <div class="icon">
                <img
                  v-if="afterAfterTomWeather?.type"
                  class="iconfont"
                  :src="getWeatherIconUrl(afterAfterTomWeather?.type)"
                />
              </div>
              <div class="temp">
                {{ (afterAfterTomWeather?.low || "") + "~" + (afterAfterTomWeather?.high || "") }}
              </div>
            </div>
          </div>
          <div class="weather-tip">Tip：{{ tip }}</div>
        </div>
      </div>
    </div>
    <!-- 最近更新 -->
    <div v-if="!$common.isEmpty(latelyArticles)" class="card-widget wow card-widget-recent">
      <div class="card-content-title">
        <i class="fa fa-history card-content-icon"></i>
        <span>最近更新</span>
      </div>
      <div
        v-for="(article, index) in latelyArticles"
        :key="index"
        @click="router.push({ path: '/article', query: { id: article.id } })"
      >
        <div class="aside-post-detail">
          <div class="aside-post-image">
            <el-image lazy class="custom-el-image" :src="article.articleCover" fit="cover">
              <!-- 懒加载颜色 -->
              <template v-slot:placeholder>
                <div>
                  <div
                    class="dot"
                    :class="{
                      leftImage: index % 2 !== 0,
                      rightImage: index % 2 === 0,
                    }"
                  ></div>
                </div>
              </template>
              <template v-slot:error>
                <div class="image-slot">
                  <div class="error-aside-image">
                    {{ article.username }}
                  </div>
                </div>
              </template>
            </el-image>
          </div>
          <div class="aside-post-title">
            <el-tooltip placement="top" effect="light" :content="article.articleTitle">
              <div>{{ article.articleTitle }}</div>
            </el-tooltip>
          </div>
        </div>
        <div class="aside-post-date">
          <el-icon :size="12" color="var(--red1)"><Calendar /></el-icon>
          {{ formatTime(article.createTime) }}
        </div>
      </div>
    </div>
    <!-- 流浪者 -->
    <div
      class="card-widget wow card-widget-wanderer"
      v-if="!$common.isEmpty(users) && route.path !== '/article'"
    >
      <div class="card-content-title">
        <i class="fa iconfont icon-pashan card-content-icon"></i>
        <span>流浪者</span>
      </div>
      <div class="wanderer-scroll">
        <vue3-seamless-scroll
          v-if="users && users.length > 0"
          :key="users.length"
          :list="users"
          :step="1"
          :hover="true"
        >
          <ul class="wanderer-list">
            <div v-for="(item, i) in users" class="wanderer-item" :key="i">
              <div class="wanderer-item-left">
                <el-avatar
                  class="wanderer-avatar"
                  :size="36"
                  :src="item.avatar ? item.avatar : webInfo.avatar"
                />
                <div class="wanderer-username">
                  {{ item.username }}
                </div>
              </div>
              <div class="wanderer-gender">
                {{ item.gender === 0 ? "保密" : item.gender === 1 ? "男" : "女" }}
              </div>
            </div>
          </ul>
        </vue3-seamless-scroll>
      </div>
      <div class="join-btn" @click="showTissue()">加入我们</div>
    </div>
    <!-- 分类 -->
    <div class="wow card-widget card-widget-sort">
      <div class="card-content-title">
        <el-icon class="card-content-icon"><Menu /></el-icon>
        <span>文章分类</span>
      </div>
      <div
        v-for="(sort, index) in sortInfo"
        :key="index"
        class="post-sort"
        @click="router.push({ path: '/sort', query: { sortId: sort.id } })"
      >
        <div>
          <span v-for="(s, i) in sort.sortName.split('')" :key="i"> {{ s }}</span>
        </div>
      </div>
    </div>
    <!-- 书虫 -->
    <div v-if="!$common.isEmpty(labelInfo)" class="wow card-widget card-widget-tag">
      <div class="card-content-title">
        <i class="fa iconfont icon-shuchongicon card-content-icon"></i>
        <span>书虫</span>
      </div>
      <div class="card-tag-cloud">
        <a
          class="item"
          v-for="(label, index) in labelInfo"
          :key="index"
          @click="router.push({ path: '/tags', query: { labelId: label.id } })"
        >
          {{ label.labelName }}
          <sup>{{ label.countOfLabel }}</sup>
        </a>
      </div>
    </div>
    <!-- 速览 -->
    <template v-if="route.path !== '/article'">
      <div
        v-for="(sort, index) in fastSeeInfo"
        @click="selectSort(sort)"
        :key="index"
        class="wow card-widget card-widget-preview"
      >
        <div class="card-content-title">
          <el-icon class="card-content-icon"><Histogram /></el-icon>
          速览
        </div>
        <div class="sort-name">
          {{ sort.sortName }}
        </div>
        <div class="sort-description">
          {{ sort.sortDescription }}
        </div>
      </div>
    </template>
    <!-- 加入我们 -->
    <el-dialog
      class="custom-my-dialog mk-join__dialog"
      title="加入组织"
      v-model="showTissueDialog"
      width="25%"
      :append-to-body="true"
      destroy-on-close
      center
    >
      <div>
        <div class="join-image"></div>
        <div>
          <div class="join-content">1. 欢迎大家一起交流学习</div>
          <div class="join-content">2. 申请通过后会加入博客交流群</div>
        </div>
      </div>
    </el-dialog>
  </div>
</template>
<script setup lang="ts">
import api from "@/api";
import { ref, computed, onMounted, type Ref, type ComputedRef } from "vue";
import { useStore } from "@/stores";
import { useRoute, useRouter } from "vue-router";
import { Vue3SeamlessScroll } from "vue3-seamless-scroll";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
import type {
  NewArticle,
  PublicUserListItem,
  SortInfo,
  LocationInfo,
  WebInfo,
  LabelInfo,
  PageView,
} from "@/types";

interface ArticleWithUsername extends NewArticle {
  username?: string;
}

type UserWithGender = PublicUserListItem;

interface Weather {
  type?: string;
  low?: string;
  high?: string;
  fengli?: string;
  fengxiang?: string;
  week?: string;
}

interface Emits {
  (e: "selectSort", sort: SortInfo): void;
}

const store = useStore();
const route = useRoute();
const router = useRouter();
const { $common, $constant } = useGlobalProperties();

const emit = defineEmits<Emits>();

const latelyArticles: Ref<ArticleWithUsername[]> = ref([]);
const showTissueDialog: Ref<boolean> = ref(false);
const total_sum: Ref<number> = ref(0);

// 从 store 获取位置和天气信息
const locationInfo: ComputedRef<LocationInfo> = computed(() => store.locationInfo);
const city: ComputedRef<string> = computed(() => locationInfo.value.city || "");
const tip: ComputedRef<string> = computed(() => locationInfo.value.tip || "");
// 天气数据（数组形式）
const weatherData: ComputedRef<Weather[]> = computed(() => {
  const w = locationInfo.value.weather;
  return Array.isArray(w) ? w : [];
});
const weather: ComputedRef<Weather> = computed(() => weatherData.value[0] || {});
const tomWeather: ComputedRef<Weather> = computed(() => weatherData.value[1] || {});
const afterTomWeather: ComputedRef<Weather> = computed(() => weatherData.value[2] || {});
const afterAfterTomWeather: ComputedRef<Weather> = computed(() => weatherData.value[3] || {});

const users: Ref<UserWithGender[]> = ref([]);
const weekDay: Ref<number> = ref(0);

const webInfo: ComputedRef<WebInfo> = computed(() => store.webInfo);
const sortInfo: ComputedRef<SortInfo[]> = computed(() => store.sortInfo);
const fastSeeInfo: ComputedRef<SortInfo[]> = computed(() => store.navigationBar);
const articleTotal: ComputedRef<number> = computed(() => store.totalArticles);
const labelInfo: ComputedRef<LabelInfo[]> = computed(() => store.labelInfo);
const newArticles: ComputedRef<ArticleWithUsername[]> = computed(() => store.newArticles);

const pageView = (value: PageView): void => store.setPageView(value);
const userList = (value: PublicUserListItem[]): void => {
  store.setUserList(value);
};

onMounted(() => {
  latelyArticles.value = newArticles.value;
  randomColor();
  getHistoryInfo();
  getUsers();
  if (!weather.value || !weather.value.week) {
    getWeek();
  }
});

const formatTime = (row: string): string => {
  const day = row.split(".")[0].split("T")[0];
  const time = row.split(".")[0].split("T")[1];
  return `${day} 日 ${time}`;
};

const weatherIcon = (dayWeather: string): string => {
  // 找到包含天气的图标
  if (!dayWeather) return "";
  const iconMap = $constant.iconWeatherMap;
  for (const key in iconMap) {
    if (iconMap[key].includes(dayWeather)) {
      return key;
    }
  }
  // 未匹配到时，返回默认图标
  return "icon-yintian";
};

const getWeatherIconUrl = (dayWeather: string): string => {
  const iconName = weatherIcon(dayWeather);
  if (!iconName) return "";
  return new URL(`../assets/svg/${iconName}.svg`, import.meta.url).href;
};

const getWeek = (): void => {
  const weekArr = ["日", "一", "二", "三", "四", "五", "六"];
  weekDay.value = new Date().getDay();
  if (weather.value && typeof weather.value === "object") {
    (weather.value as Weather).week = weekArr[weekDay.value];
  }
  if (tomWeather.value && typeof tomWeather.value === "object") {
    (tomWeather.value as Weather).week = weekArr[weekDay.value + 1];
    if (weekDay.value == 6) {
      (tomWeather.value as Weather).week = "日";
    }
  }
};

const selectSort = (sort: SortInfo): void => {
  emit("selectSort", sort);
};

const showTip = (): void => {
  router.push({ path: "/weiYan" });
};

const randomColor = (): void => {
  function getRandomColor(): string {
    return `rgb(${Math.random() * 255}, ${Math.random() * 255}, ${Math.random() * 255})`;
  }
  const itemEls = document.querySelectorAll(".item");
  for (const item of itemEls) {
    (item as HTMLElement).style.color = getRandomColor();
  }
};

const getHistoryInfo = async (): Promise<void> => {
  try {
    const res = await api.getIpStatistics();
    total_sum.value = res.total_sum || 0;
    pageView(res as unknown as PageView);
  } catch {
    /* ignored */
  }
};

const getUsers = async (): Promise<void> => {
  try {
    const res = await api.getUserListPaginated();
    users.value = (res.list || []) as UserWithGender[];
    userList(users.value);
  } catch {
    /* ignored */
  }
};

const showTissue = (): void => {
  showTissueDialog.value = true;
};
</script>
<style lang="scss">
.mk-join__dialog {
  background: var(--pink);

  .el-dialog__title {
    color: var(--fontColor);
  }

  .join-image {
    margin: 0 auto 10px;
    border-radius: 10px;
    height: 150px;
    width: 150px;
    background: var(--joinImage) center center / cover no-repeat;
  }

  .join-content {
    font-size: 14px;
    color: var(--bigRed1);
    line-height: 1.5;
    margin: 5px;
  }
}
</style>
<style lang="scss" scoped>
.aside-container {
  // 网站信息卡片
  .aside-card-content {
    background: linear-gradient(
      -45deg,
      var(--wheat),
      var(--red6),
      var(--blue7),
      var(--red),
      var(--pink1)
    );
    background-size: 400% 400%;
    animation: gradientBG 16s ease infinite;
    display: flex;
    flex-direction: column;
    align-items: center;
    border-radius: 18px;
    position: relative;
    overflow: hidden;

    .image-dot {
      width: 25px;
      height: 25px;
      background: var(--green1);
      position: absolute;
      border-radius: 50%;
      border: 4px solid var(--favoriteBg);
      z-index: 20;
      transform: translate(38px, 85px);
    }

    .author-box {
      position: relative;
      width: 120px;
      height: 120px;
      border-radius: 50%;
      overflow: hidden;

      span {
        position: absolute;
        inset: 5px;
        border-radius: 50%;
        background: var(--favoriteBg);
        z-index: 1;
      }

      .author-img {
        margin: auto;
        border-radius: 50%;
        overflow: hidden;
        width: 102px;
        height: 102px;
        z-index: 10;
        background: var(--maxMaxWhiteMask);

        img {
          border-radius: 11px;
          display: block;
          margin: 0 auto 20px;
          max-width: 100%;
          transition: all 0.3s ease;

          &:hover {
            transform: rotate(360deg);
          }
        }
      }
    }

    .web-name {
      font-size: 30px;
      font-weight: bold;
      margin: 20px 0;
      color: var(--noneWhite);
    }

    .web-info {
      margin: 15px 0;
      width: 80%;
      display: flex;
      flex-direction: row;
      justify-content: space-around;

      .blog-info-box {
        flex-direction: column;
        color: var(--noneWhite);

        .blog-info-num {
          margin-top: 12px;
        }
      }
    }

    .collection-btn {
      position: relative;
      margin-top: 12px;
      background: var(--purple3);
      width: 65%;
      height: 35px;
      border-radius: 1rem;
      text-align: center;
      line-height: 35px;
      color: var(--bigRed3);
      overflow: hidden;
      z-index: 1;
      margin-bottom: 10px;

      &::before {
        background: var(--gradualRed);
        position: absolute;
        top: 0;
        right: 0;
        bottom: 0;
        left: 0;
        content: "";
        transform: scaleX(0);
        transform-origin: 0;
        transition: transform 0.5s ease-out;
        transition-timing-function: cubic-bezier(0.45, 1.64, 0.47, 0.66);
        border-radius: 1rem;
        z-index: -1;
      }

      &:hover::before {
        transform: scaleX(1);
      }

      &:hover {
        color: var(--favoriteBg);
      }

      .icon-loading {
        margin-right: 5px;
        color: var(--maxLightRed1);
        animation: rotate 2s linear infinite;
      }
    }

    .aside-link-box {
      margin: 8px;
      width: 195px;

      a {
        width: 30px;
        height: 30px;

        i {
          width: 30px;
          border-radius: 50%;
          color: var(--black9);
          background: var(--mask);
          height: 30px;
          transition: all 0.3s ease;

          &:hover {
            background: var(--favoriteBg);
            color: var(--black8);
            transform: scale(1.2);
          }

          &:before {
            height: 16px;
            width: 16px;
            display: flex;
            align-items: center;
            justify-content: center;
          }
        }
      }
    }
  }

  // 卡片组件通用样式
  .card-widget {
    clip-path: polygon(
      100% 0,
      100% 100%,
      0 100%,
      0 calc(100% - 52.5px),
      12.5px calc(100% - 40px),
      12.5px calc(100% - 50px),
      0 calc(100% - 62.5px),
      0 calc(100% - 82.5px),
      12.5px calc(100% - 70px),
      12.5px calc(100% - 80px),
      0 calc(100% - 92.5px),
      0 calc(100% - 112.5px),
      12.5px calc(100% - 100px),
      12.5px calc(100% - 112.5px),
      12.5px calc(100% - 110px),
      0 calc(100% - 122.5px),
      0 calc(100% - 142.5px),
      12.5px calc(100% - 130px),
      12.5px calc(100% - 141.5px),
      0 calc(100% - 154px),
      0 0
    );
    border-left: none;
    font-size: 105%;
    border-radius: 18px;
    border: 2px solid var(--myAsideBorderColor);
    background: var(--myAsideColor);
    transition: all 1s;
    position: relative;
    margin-top: 10px;
    animation: hideToShow 1s ease-in-out;

    &::before {
      content: "";
      width: 12.5px;
      background: linear-gradient(to top, transparent, var(--red));
      display: block;
      position: absolute;
      left: 0;
      height: 113px;
      bottom: 27px;
    }

    &:hover {
      transform: scale(0.95);
      background-color: var(--blue);
      transition-duration: 0.4s;
      border-color: var(--myAsideBorderColor1);
    }

    .card-content-title {
      font-size: 18px;
      margin-bottom: 15px;
      display: flex;
      align-items: center;

      .card-content-icon {
        color: var(--red);
        margin-right: 5px;
        animation: scale 1s ease-in-out infinite;
      }
    }

    // 天气卡片
    &.card-widget-weather {
      padding: 25px 10px 5px 17px;

      .weather-content {
        line-height: 25px;
        word-break: break-all;
        color: var(--darkBlue);

        .weather-content-inner {
          padding: 0 8px;

          .mk-address {
            .city {
              font-size: 14px;
              color: var(--fontColor);
            }
          }

          .mk-Today_weather {
            .left {
              display: flex;

              .weather-icon {
                .iconfont {
                  width: 50px;
                  height: 50px;
                }
              }

              .temp {
                margin-left: 5px;

                .number {
                  font-size: 20px;
                  color: var(--fontColor);
                }

                .weather {
                  color: var(--fontColor);
                  display: flex;
                  justify-content: space-between;
                  font-size: 14px;

                  .weather-divider {
                    margin: 0 4px;
                    color: var(--red);
                  }
                }
              }
            }

            .right {
              background: var(--red);
              color: var(--favoriteBg);
              font-size: 12px;
              border-radius: 25px;
              padding: 1px 12px;
              text-align: center;
            }

            .iconfont {
              width: 50px;
              height: 50px;
            }
          }

          .mk-week_weather {
            display: flex;
            justify-content: space-between;

            .tomorrow,
            .afterTomorrow,
            .afterAfterTomWeather {
              color: var(--fontColor);
              display: flex;
              flex-direction: column;
              align-items: center;

              .text {
                font-size: 13px;
              }

              .iconfont {
                width: 30px;
                height: 30px;
              }

              .temp {
                font-size: 12px;
              }
            }
          }

          .weather-tip {
            font-size: 12px;
            color: var(--fontColor);
          }
        }
      }
    }

    // 最近更新卡片
    &.card-widget-recent {
      padding: 25px;

      .aside-post-detail {
        display: flex;

        .aside-post-image {
          width: 60%;
          border-radius: 0.2rem;
          margin-right: 8px;
          overflow: hidden;
          transition: all 1s;

          &:hover {
            transform: scale(1.1);
          }

          .dot {
            position: absolute;
            height: 100%;
            width: 100%;
            background: var(--gradientAnimation);
          }

          .error-aside-image {
            background: var(--gradientAnimation);
            color: var(--white);
            padding: 10px;
            text-align: center;
            width: 100%;
            height: 100%;
          }
        }

        .aside-post-title {
          height: 40px;
          line-height: 20px;
          width: 50%;
          text-overflow: ellipsis;
          display: -webkit-box;
          -webkit-line-clamp: 2;
          line-clamp: 2;
          -webkit-box-orient: vertical;
          overflow: hidden;
        }
      }

      .aside-post-date {
        margin: 8px 0;
        color: var(--red1);
        font-size: 12px;
        display: flex;
        align-items: center;
      }
    }

    // 流浪者卡片
    &.card-widget-wanderer {
      padding: 25px 10px 5px 17px;

      .wanderer-scroll {
        height: 200px;
        overflow: hidden;

        .wanderer-list {
          padding-left: 0;

          .wanderer-item {
            display: flex;
            justify-content: space-between;

            .wanderer-item-left {
              display: flex;

              .wanderer-avatar {
                margin-bottom: 10px;
              }

              .wanderer-username {
                margin-left: 10px;
                height: 36px;
                line-height: 36px;
                overflow: hidden;
                max-width: 110px;
              }
            }

            .wanderer-gender {
              height: 36px;
              line-height: 36px;
            }
          }
        }
      }

      .join-btn {
        padding: 13px 15px;
        background: var(--maxLightRed);
        border-radius: 3rem;
        color: var(--white);
        width: 100px;
        user-select: none;
        text-align: center;
        margin: 20px auto 0;
        transition: all 1s;

        &:hover {
          transform: scale(1.2);
        }
      }
    }

    // 分类卡片
    &.card-widget-sort {
      padding: 25px 10px 5px 17px;

      .post-sort {
        border-radius: 1rem;
        margin-bottom: 15px;
        line-height: 30px;
        transition: all 0.3s;

        &:hover {
          background: var(--gradientAnimation);
          padding: 0px 15px;
          color: var(--white);
        }
      }
    }

    // 书虫卡片
    &.card-widget-tag {
      padding: 10px 10px 10px 15px;

      .card-tag-cloud {
        width: inherit;
        display: flex;
        flex-wrap: wrap;
        justify-content: flex-start;

        a {
          height: 30px;
          line-height: 15px;
          padding: 4px;
          margin: 2px 2px 2px 0;
          border-radius: 3px;
          font-size: 18px;

          &:hover {
            color: var(--white) !important;
            background: var(--red);
            animation: hideToShow 0.3s ease-in-out;
          }

          sup {
            font-size: 10px;
            opacity: 0.6;
          }
        }
      }
    }

    // 速览卡片
    &.card-widget-preview {
      background-size: 100% 100%;
      padding: 20px 40px 40px;
      font-size: 18px;
      color: var(--fontColor);

      .sort-name {
        font-weight: bold;
        font-size: 22px;
        margin-top: 30px;
        white-space: nowrap;
        text-overflow: ellipsis;
        overflow: hidden;

        &:after {
          top: 104px;
          width: 70px;
          left: 40px;
          height: 3px;
          background: var(--gradientAnimation);
          content: "";
          border-radius: 1px;
          position: absolute;
        }
      }

      .sort-description {
        font-weight: 400;
        font-size: 14px;
        margin-top: 20px;
        white-space: nowrap;
        text-overflow: ellipsis;
        overflow: hidden;
      }
    }
  }
}

// 媒体查询
@media screen and (max-width: 1100px) {
  .aside-container {
    .card-widget {
      &.card-widget-recent {
        .aside-post-detail {
          .aside-post-image {
            width: 30%;
          }
        }
      }
    }
  }
}
</style>
