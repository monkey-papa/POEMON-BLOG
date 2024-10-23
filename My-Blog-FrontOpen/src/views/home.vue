<template>
  <div>
    <!-- el过渡动画 -->
    <transition name="el-fade-in-linear">
      <!-- 导航栏 -->
      <div
        v-show="toolbar.visible"
        @mouseenter="hoverEnter = true"
        @mouseleave="hoverEnter = false"
        :class="[
          { enter: toolbar.enter },
          {
            hoverEnter: hoverEnter && !toolbar.enter,
          },
        ]"
        class="toolbar-content myBetween"
      >
        <!-- 网站名称 -->
        <div @click="$router.push({ path: '/' })" class="toolbar-title">
          <h2>
            {{ $store.state.webInfo.webName }}
          </h2>
        </div>
        <!-- 手机抽屉按钮 -->
        <div
          v-if="$common.mobile() || mobile"
          class="toolbar-mobile-menu"
          @click="toolbarDrawer = !toolbarDrawer"
          :class="{ enter: toolbar.enter }"
        >
          <i class="el-icon-s-operation"></i>
        </div>
        <!-- 常规头部导航列表 -->
        <template v-else>
          <div class="center-toolbar">
            <ul class="scroll-menu">
              <!-- 文章 -->
              <li>
                <el-popover
                  :visible-arrow="false"
                  transition="el-zoom-in-center"
                  close-delay="300"
                  popper-class="mk-popper"
                  placement="bottom"
                  trigger="hover"
                >
                  <div class="mk-popover_item">
                    <!-- 分类 -->
                    <li @click="$router.push({ path: '/sort' })">
                      <img
                        style="vertical-align: -3px"
                        src="../assets/svg/sort.svg"
                      />
                      分类
                    </li>
                    <!-- 标签 -->
                    <li
                      @click="
                        $router.push({ path: '/tags', query: { labelId: 25 } })
                      "
                    >
                      <img
                        style="vertical-align: -3px"
                        src="../assets/svg/tag.svg"
                      />
                      标签
                    </li>
                  </div>
                  <span slot="reference" class="el-dropdown-link">
                    <img
                      style="vertical-align: -3px"
                      src="../assets/svg/document.svg"
                    />
                    文章
                  </span>
                </el-popover>
              </li>
              <!-- 空间 -->
              <li>
                <el-popover
                  :visible-arrow="false"
                  transition="el-zoom-in-center"
                  close-delay="300"
                  popper-class="mk-popper"
                  placement="bottom"
                  trigger="hover"
                >
                  <div class="mk-popover_item">
                    <!-- 音乐 -->
                    <li @click="$router.push({ path: '/funny' })">
                      <img
                        style="vertical-align: -3px"
                        src="../assets/svg/music.svg"
                      />
                      幻音坊
                    </li>
                    <!-- 藏宝阁 -->
                    <li @click="$router.push({ path: '/tools' })">
                      <img
                        style="vertical-align: -3px"
                        src="../assets/svg/treasure.svg"
                      />
                      藏宝阁
                    </li>
                  </div>
                  <span slot="reference" class="el-dropdown-link">
                    <img
                      style="vertical-align: -3px"
                      src="../assets/svg/space.svg"
                    />
                    空间
                  </span>
                </el-popover>
              </li>
              <!-- 我的 -->
              <li>
                <el-popover
                  :visible-arrow="false"
                  transition="el-zoom-in-center"
                  close-delay="300"
                  popper-class="mk-popper"
                  placement="bottom"
                  trigger="hover"
                >
                  <div class="mk-popover_item">
                    <!-- 相册 -->
                    <li @click="$router.push({ path: '/travel' })">
                      📸 <span>相册</span>
                    </li>
                    <!-- 爱链 -->
                    <li @click="$router.push({ path: '/love' })">
                      <img
                        style="vertical-align: -3px"
                        src="../assets/svg/love.svg"
                      />
                      爱链
                    </li>
                    <!-- 小游戏 -->
                    <li @click="openPcGame">🎮 <span>小游戏</span></li>
                    <!-- 编辑怪 -->
                    <li @click="EDIT()">
                      <img
                        style="vertical-align: -5px"
                        src="../assets/svg/pencil.svg"
                      />
                      编辑怪
                    </li>
                    <!-- 关于 -->
                    <li @click="$router.push({ path: '/about' })">
                      🐷 <span>关于我</span>
                    </li>
                  </div>
                  <span slot="reference" class="el-dropdown-link">
                    <img
                      style="vertical-align: -3px"
                      src="../assets/svg/home.svg"
                    />
                    我的
                  </span>
                </el-popover>
              </li>
              <!-- 社交 -->
              <li>
                <el-popover
                  :visible-arrow="false"
                  transition="el-zoom-in-center"
                  close-delay="300"
                  popper-class="mk-popper"
                  placement="bottom"
                  trigger="hover"
                >
                  <div class="mk-popover_item">
                    <!-- 留言厅 -->
                    <li @click="$router.push({ path: '/message' })">
                      ✍🏻 <span>留言厅</span>
                    </li>
                    <!-- 友链 -->
                    <li @click="$router.push({ path: '/friend' })">
                      🎀 <span>友链</span>
                    </li>
                  </div>
                  <span slot="reference" class="el-dropdown-link">
                    <img
                      style="vertical-align: -3px"
                      src="../assets/svg/socialContact.svg"
                    />
                    社交
                  </span>
                </el-popover>
              </li>
            </ul>
          </div>
          <div>
            <ul class="scroll-menu">
              <!-- 随机文章 -->
              <li
                title="随便看看-随机访问一篇本站文章"
                style="color: var(--fontColor); font-size: 20px"
                @click="openRandomArticle"
              >
                <div class="my-menu"><i class="fa fa-grav"></i></div>
              </li>
              <!-- 切换背景 -->
              <li
                title="切换背景-换一种背景，换一种感觉。"
                id="changeThemeRef"
                style="color: var(--fontColor); font-size: 20px"
                @click="openChangeBg"
              >
                <div class="my-menu"><i class="fa fa-image"></i></div>
              </li>
              <!-- 关闭樱花 -->
              <li
                @click="handleSakura()"
                style="color: var(--fontColor); font-size: 21px"
              >
                <div class="my-menu"><i class="fa fa-pagelines"></i></div>
              </li>
              <!-- 关灯 -->
              <li
                style="transform: scale(0.16); width: 70px; height: 0"
                class="my-menu"
              >
                <switchBtn
                  id="changeColorRef"
                  @click.native="changeColor"
                ></switchBtn>
              </li>
              <!-- 个人中心 -->
              <li>
                <el-dropdown hide-timeout="300" placement="bottom">
                  <el-avatar
                    class="user-avatar"
                    :size="36"
                    style="margin-top: 4px"
                    :src="
                      $store.state.currentUser.avatar
                        ? $store.state.currentUser.avatar
                        : $store.state.webInfo.avatar
                    "
                  >
                  </el-avatar>
                  <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item
                      @click.native="$router.push({ path: '/user' })"
                      v-if="!$common.isEmpty($store.state.currentUser)"
                      class="user"
                    >
                      <i class="fa fa-user-circle" aria-hidden="true"></i>
                      <span>个人</span>
                    </el-dropdown-item>
                    <el-dropdown-item
                      @click.native="logout()"
                      v-if="!$common.isEmpty($store.state.currentUser)"
                      class="logout"
                    >
                      <i class="fa fa-sign-out" aria-hidden="true"></i>
                      <span>退出</span>
                    </el-dropdown-item>
                    <el-dropdown-item
                      @click.native="$router.push({ path: '/user' })"
                      v-if="$common.isEmpty($store.state.currentUser)"
                      class="login"
                    >
                      <i class="fa fa-sign-in" aria-hidden="true"></i>
                      <span>登录</span>
                    </el-dropdown-item>
                  </el-dropdown-menu>
                </el-dropdown>
              </li>
            </ul>
          </div>
        </template>
      </div>
    </transition>
    <div id="main-container">
      <router-view></router-view>
    </div>
    <!-- 猫 -->
    <div
      href="#"
      class="cd-top"
      v-if="!$common.mobile()"
      @click="toTop()"
    ></div>
    <!-- 右下角按钮 -->
    <div class="toolButton">
      <!-- 火箭 -->
      <div
        class="backTop"
        v-if="$common.mobile() && toolButton"
        @click="toTop()"
      >
        <img src="../assets/svg/rocket.svg" />
      </div>
      <!-- 右下角切换按钮 -->
      <el-popover
        transition="el-zoom-in-top"
        :visible-arrow="false"
        placement="left"
        :close-delay="500"
        trigger="hover"
      >
        <!-- 旋转齿轮 -->
        <div slot="reference">
          <div class="button">
            <i
              class="fa fa-cog setting-color"
              aria-hidden="true"
              style="font-size: large"
            ></i>
          </div>
        </div>
        <div slot="reference">
          <div class="button" v-show="topPercentage < 98 && !topPercentageType">
            {{ topPercentage }}%
          </div>
        </div>
        <!-- 关灯 -->
        <div class="my-setting">
          <el-tooltip placement="top" effect="light">
            <!-- 雪花片 -->
            <div slot="content">想看雪花吗？(◕ᴗ◕✿)</div>
            <div>
              <i
                class="fa fa-snowflake-o"
                :class="{ active: mouseAnimation }"
                aria-hidden="true"
                @click="changeMouseAnimation()"
              ></i>
            </div>
          </el-tooltip>
        </div>
      </el-popover>
    </div>
    <!-- 点击动画 -->
    <canvas
      v-if="mouseAnimation"
      id="mousedown"
      style="
        position: fixed;
        left: 0;
        top: 0;
        pointer-events: none;
        z-index: 1000;
      "
    >
    </canvas>
    <!-- 移动端抽屉 -->
    <el-drawer
      :visible.sync="toolbarDrawer"
      :show-close="false"
      size="300px"
      custom-class="toolbarDrawer"
      direction="ltr"
    >
      <div class="backdrop-color">
        <!-- 信息 -->
        <div class="sidebar is-center">
          <div class="avatar-img">
            <img :src="$store.state.webInfo.avatar" />
          </div>
          <div class="author-info_name">
            『{{ $store.state.webInfo.webName }}』
          </div>
          <div class="author-info__description">让我再享受一下</div>
        </div>
        <!-- 分类 -->
        <div class="site-data is-center">
          <div class="blog-info-box">
            <span>文章</span>
            <span class="blog-info-num">{{ $store.getters.articleTotal }}</span>
          </div>
          <div class="blog-info-box">
            <span>标签</span>
            <span class="blog-info-num">{{
              $store.getters.labelInfo.length
            }}</span>
          </div>
          <div class="blog-info-box">
            <span>分类</span>
            <span class="blog-info-num">{{
              $store.state.sortInfo.length
            }}</span>
          </div>
        </div>
        <hr />
        <!-- 菜单 -->
        <div>
          <div class="small-menu">
            <!-- 文章 -->
            <div>
              <span class="el-dropdown-link">
                <img
                  style="vertical-align: -3px"
                  src="../assets/svg/document.svg"
                />
                <span>文章</span>
              </span>
              <div class="menu_item">
                <!-- 分类 -->
                <span @click="$router.push({ path: '/sort' })">
                  <div class="my-menu">
                    <img
                      style="vertical-align: -3px"
                      src="../assets/svg/sort.svg"
                    />
                    分类
                  </div>
                </span>
                <!-- 标签 -->
                <span
                  @click="
                    $router.push({ path: '/tags', query: { labelId: 25 } })
                  "
                >
                  <div class="my-menu">
                    <img
                      style="vertical-align: -3px"
                      src="../assets/svg/tag.svg"
                    />
                    标签
                  </div>
                </span>
              </div>
            </div>
            <!-- 空间 -->
            <div>
              <span class="el-dropdown-link">
                <img
                  style="vertical-align: -3px"
                  src="../assets/svg/space.svg"
                />
                <span>空间</span>
              </span>
              <div class="menu_item">
                <!-- 音乐 -->
                <span @click="smallMenu({ path: '/funny' })">
                  <div class="my-menu">
                    <img
                      style="vertical-align: -3px"
                      src="../assets/svg/music.svg"
                    />
                    幻音坊
                  </div>
                </span>
                <!-- 藏宝阁 -->
                <span @click="smallMenu({ path: '/tools' })">
                  <div class="my-menu">
                    <img
                      style="vertical-align: -3px"
                      src="../assets/svg/treasure.svg"
                    />
                    藏宝阁
                  </div>
                </span>
              </div>
            </div>
            <!-- 我的 -->
            <div>
              <span class="el-dropdown-link">
                <img
                  style="vertical-align: -3px"
                  src="../assets/svg/home.svg"
                />
                <span>我的</span>
              </span>
              <div class="menu_item">
                <!-- 相册 -->
                <span @click="smallMenu({ path: '/travel' })">
                  <div class="my-menu">📸 <span>相册</span></div>
                </span>
                <!-- 爱链 -->
                <span @click="smallMenu({ path: '/love' })">
                  <div class="my-menu">
                    <img
                      style="vertical-align: -3px"
                      src="../assets/svg/love.svg"
                    />
                    爱链
                  </div>
                </span>
                <span @click="openPcGame">
                  <div class="my-menu">🎮 <span>小游戏</span></div>
                </span>
                <span @click="EDIT()">
                  <div class="my-menu">
                    <img
                      style="vertical-align: -5px"
                      src="../assets/svg/pencil.svg"
                    />
                    编辑怪
                  </div>
                </span>
                <!-- 关于 -->
                <span @click="smallMenu({ path: '/about' })">
                  <div class="my-menu">🐷 <span>关于我</span></div>
                </span>
              </div>
            </div>
            <!-- 社交 -->
            <div>
              <span class="el-dropdown-link">
                <img
                  style="vertical-align: -3px"
                  src="../assets/svg/socialContact.svg"
                />
                <span>社交</span>
              </span>
              <div class="menu_item">
                <!-- 留言厅 -->
                <span @click="smallMenu({ path: '/message' })">
                  <div class="my-menu">✍🏻 <span>留言厅</span></div>
                </span>
                <!-- 友链 -->
                <span @click="smallMenu({ path: '/friend' })">
                  <div class="my-menu">🎀 <span>友链</span></div>
                </span>
              </div>
            </div>
            <template v-if="$common.isEmpty($store.state.currentUser)">
              <li @click="smallMenu({ path: '/user' })">
                <div>
                  <i class="fa fa-sign-in" aria-hidden="true"></i>
                  <span>&nbsp;登录</span>
                </div>
              </li>
            </template>
            <template v-else>
              <li @click="smallMenu({ path: '/user' })">
                <div>
                  <i class="fa fa-user-circle" aria-hidden="true"></i>
                  <span>&nbsp;个人</span>
                </div>
              </li>
              <li @click="smallMenuLogout()">
                <div>
                  <i class="fa fa-sign-out" aria-hidden="true"></i>
                  <span>&nbsp;退出</span>
                </div>
              </li>
            </template>
            <!-- 随机文章 -->
            <li
              style="color: var(--black); font-size: 20px"
              @click="openRandomArticle"
            >
              <div class="my-menu"><i class="fa fa-grav"></i></div>
            </li>
            <!-- 切换背景 -->
            <li
              style="color: var(--black); font-size: 20px"
              @click="openChangeBg"
            >
              <div class="my-menu"><i class="fa fa-image"></i></div>
            </li>
            <!-- 关灯 -->
            <li
              @click="changeColor"
              style="color: var(--black); font-size: 23px"
            >
              <div class="my-menu"><i class="fa-adjust fa"></i></div>
            </li>
            <!-- 关闭樱花 -->
            <li
              @click="handleSakura()"
              style="color: var(--black); font-size: 21px"
            >
              <div class="my-menu"><i class="fa fa-pagelines"></i></div>
            </li>
          </div>
        </div>
      </div>
    </el-drawer>
    <!-- 小游戏 -->
    <el-dialog
      title="随机小游戏"
      :visible.sync="disGame"
      :modal-append-to-body="false"
      width="auto"
      align="center"
    >
      <div style="text-align: center">
        <iframe :src="game" style="width: 100%; height: 900px"></iframe>
      </div>
    </el-dialog>
    <!-- 修改背景 -->
    <el-dialog
      custom-class="changeBgBox"
      title="切换背景"
      :visible.sync="changeBgBox"
      :modal="false"
      width="60%"
      align="left"
    >
      <div style="text-align: center">
        <button
          @click="defaultBtn"
          style="
            background: var(--blue9);
            display: block;
            width: 100%;
            padding: 15px 0;
            border-radius: 6px;
            color: var(--white2);
            border: none;
            font-size: 18px;
          "
        >
          <i
            style="font-size: 18px; margin-right: 5px"
            class="fa fa-refresh"
          ></i
          >恢复默认主题
        </button>
        <div class="customImg">
          <div class="customImg-item">设置自定义背景</div>
          <el-input
            v-model="httpInput"
            type="text"
            id="pic-link"
            size="70%"
            maxlength="1000"
            placeholder="请输入有效的图片链接，如 https://source.fomal.cc/img/home_bg.webp"
          ></el-input>
          <button class="httpButton" @click="httpInputBtn">🌈切换背景🌈</button>
        </div>
        <div class="customImg">
          <div class="customImg-item" style="margin-bottom: 15px">
            设置主题色
          </div>
          <div class="color-box_contain">
            <div
              @click="setColor(item.color)"
              class="color-box"
              :class="{ active: themeActive === item.color }"
              :style="{ borderColor: item.varColor }"
              v-for="(item, i) in colorList"
              :key="i"
            >
              <div
                class="gun"
                :style="{ backgroundColor: item.varColor }"
              ></div>
              <div class="center">
                <div class="top" :style="{ color: item.varColor }">
                  {{ item.name }}
                </div>
                <div class="bottom" :style="{ color: item.varColor }">
                  {{ item.color }} | {{ item.rgb }}
                </div>
              </div>
            </div>
          </div>
        </div>
        <div v-for="(items, index) in themeMap" :key="index">
          <h2 style="color: var(--fontColor); text-align: left">
            {{ items.title }}
          </h2>
          <div style="display: flex; align-items: center">
            <span class="iconRotate showIcon">
              <i
                style="
                  display: flex;
                  height: 16px;
                  width: 16px;
                  color: var(--red);
                "
                class="iconfont icon-fengche"
              ></i>
            </span>
            <el-collapse @change="handleChangeBg(items.handleVal, index)">
              <el-collapse-item :title="items.collapseTitle" name="1">
                <div class="bgBox">
                  <a
                    @click="changeBg(item)"
                    v-for="(item, i) in items.dataList"
                    :key="i"
                    href="javascript:;"
                    :class="items.class"
                    :style="
                      items.style === 'img'
                        ? { backgroundImage: `url(${item})` }
                        : items.style === 'gradient'
                        ? { background: `linear-gradient(${item})` }
                        : { background: `${item}` }
                    "
                  ></a>
                </div>
              </el-collapse-item>
            </el-collapse>
          </div>
        </div>
      </div>
    </el-dialog>
  </div>
</template>
<script>
import mousedown from "../utils/mousedown";
const switchBtn = () => import("./common/switchBtn");
export default {
  components: {
    switchBtn,
  },
  data() {
    return {
      disGame: false,
      pcGameList: [1, 2, 3, 4, 5, 6, 7, 8],
      game: "http://game.eean.cn/pc/game",
      toolButton: false,
      hoverEnter: false,
      mouseAnimation: false,
      isDark: false, // false是黑夜,显示太阳按钮
      scrollTop: 0,
      toolbarDrawer: false,
      mobile: false,
      changeBgBox: false,
      editFlag: false,
      sakuraFlag: false,
      topPercentageType: false,
      themeMap: this.$constant.themeMapConfig,
      httpInput: "",
      themeActive: "",
      colorList: [
        {
          name: "兔子坦克形态",
          color: "#04597b",
          rgb: "rgb(4,89,123)",
          rgba: "var(--sDarkBlueRgb)",
          varColor: "var(--sDarkBlue)",
        },
        {
          name: "鳄鱼恶霸形态",
          color: "#b04fe6",
          rgb: "rgb(176,79,230)",
          rgba: "var(--sPurpleRgb)",
          varColor: "var(--sPurple)",
        },
        {
          name: "巨龙熔岩形态",
          color: "#ff7500",
          rgb: "rgb(255, 117, 0)",
          rgba: "var(--sOrangeRgb)",
          varColor: "var(--sOrange)",
        },
        {
          name: "向日癸形态",
          color: "#ffa500",
          rgb: "rgb(255,165,0)",
          rgba: "var(--sYellowRgb)",
          varColor: "var(--sYellow)",
        },
        {
          name: "自然精灵形态",
          color: "#6bdf8f",
          rgb: "rgb(107,223,143)",
          rgba: "var(--sGreenRgb)",
          varColor: "var(--sGreen)",
        },
        {
          name: "锦鲤粉形态",
          color: "#ec695c",
          rgb: "rgb(236,105,92)",
          rgba: "var(--sRedRgb)",
          varColor: "var(--sRed)",
        },
        {
          name: "中国红形态",
          color: "#d61010",
          rgb: "rgb(214, 16, 16)",
          rgba: "var(--sBigRedRgb)",
          varColor: "var(--sBigRed)",
        },
        {
          name: "至尊龙形态",
          color: "#425aef",
          rgb: "rgb(66, 90, 239)",
          rgba: "var(--sBlueRgb)",
          varColor: "var(--sBlue)",
        },
      ],
    };
  },
  mounted() {
    //默认关闭樱花
    this.handleSakura();
    if (localStorage.getItem("theme") == null) {
      localStorage.setItem("theme", true);
    }
    if (this.mouseAnimation) {
      mousedown();
    }
    window.addEventListener("scroll", this.onScrollPage);
    // 关灯
    if (localStorage.getItem("theme") == "false") {
      this.isDark = false;
    } else {
      this.isDark = true;
    }
    this.changeBg();
  },
  beforeDestroy() {
    window.removeEventListener("scroll", this.onScrollPage);
  },
  watch: {
    $route: {
      immediate: true,
      handler(to) {
        if (to.path == "/funny") {
          this.topPercentageType = true;
        } else {
          this.topPercentageType = false;
        }
      },
    },
    scrollTop(scrollTop, oldScrollTop) {
      //如果滑动距离超过屏幕高度三分之一视为进入页面，背景改为蓝色
      let enter = scrollTop > window.innerHeight / 2;
      const top = scrollTop - oldScrollTop < 0;
      let isShow = scrollTop - window.innerHeight > 30;
      this.toolButton = isShow;
      if (isShow && !this.$common.mobile()) {
        if (window.innerHeight > 950) {
          $(".cd-top").css("top", "0");
        } else {
          $(".cd-top").css("top", window.innerHeight - 950 + "px");
        }
      } else if (!isShow && !this.$common.mobile()) {
        $(".cd-top").css("top", "-900px");
      }
      //导航栏显示与颜色
      let toolbarStatus = {
        enter: enter,
        visible: top,
      };
      this.$store.commit("changeToolbarStatus", toolbarStatus);
    },
  },
  created() {
    let toolbarStatus = {
      enter: false,
      visible: true,
    };
    this.$store.commit("changeToolbarStatus", toolbarStatus);
    this.getWebInfo();
    this.getSortInfo();
    this.mobile = document.body.clientWidth < 600;
    window.addEventListener("resize", () => {
      let docWidth = document.body.clientWidth;
      if (docWidth < 600) {
        this.mobile = true;
      } else {
        this.mobile = false;
      }
    });
  },
  computed: {
    toolbar() {
      return this.$store.state.toolbar;
    },
    topPercentage() {
      return this.$store.state.top;
    },
  },
  methods: {
    setColor(color) {
      this.themeActive = color;
      const root = document.querySelector(":root");
      root.style.setProperty("--themeColor", color);
      this.$common.getThemeRgb();
      document.documentElement.dataset.theme =
        localStorage.getItem("theme") === "false" ? "theme2-dark" : "theme2";
      localStorage.setItem("themeColor", color);
    },
    openRandomArticle() {
      // 我的数据库是从id为12的文章开始的,所以随机数 + 12
      const articleTotal = this.$store.state.articleTotal;
      const random = Math.floor(Math.random() * articleTotal) + 12;
      this.$router.push({ path: "/article", query: { id: random } });
    },
    async httpInputBtn() {
      if (this.httpInput.length === 0) {
        return this.$notify({
          title: "链接不对🤣",
          message: "请输入有效的图片链接！",
          type: "warning",
          offset: 50,
          position: "top-left",
        });
      }
      const status = await this.$common.isValidHttpUrl(this.httpInput);
      if (status) {
        this.changeBg(this.httpInput);
        this.$notify({
          title: "可以啦🍨",
          message: "切换自定义背景成功",
          type: "success",
          offset: 50,
          position: "top-left",
        });
        this.httpInput = "";
      } else {
        this.$notify({
          title: "链接不对🤣",
          message: "请输入有效的图片链接！",
          type: "warning",
          offset: 50,
          position: "top-left",
        });
      }
    },
    handleSakura() {
      if (this.sakuraFlag) {
        window.startSakura();
        this.sakuraFlag = false;
      } else {
        const dom = document.querySelector("#canvas_sakura");
        this.sakuraFlag = true;
        if (dom) {
          dom.remove();
        }
      }
    },
    smallMenu(data) {
      this.$router.push(data);
      this.toolbarDrawer = false;
    },
    smallMenuLogout() {
      this.logout();
      this.toolbarDrawer = false;
    },
    goIm() {
      if (this.$common.isEmpty(this.$store.state.currentUser)) {
        this.$notify({
          type: "error",
          title: "可恶🤬",
          message: "请先登录！",
          position: "top-left",
          offset: 50,
        });
      } else {
        let userToken = this.$common.encrypt(localStorage.getItem("userToken"));
        window.open(this.$constant.imBaseURL + "?userToken=" + userToken);
      }
    },
    logout() {
      this.$notify({
        title: "可以啦🍨",
        message: "退出登录成功！",
        type: "success",
        offset: 50,
        position: "top-left",
      });
      this.$store.commit("loadCurrentUser", {});
      localStorage.removeItem("userToken");
      this.$router.push({ path: "/" });
    },
    getWebInfo() {
      this.$http
        .get(this.$constant.baseURL + "/webInfo/getWebInfo")
        .then((res) => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.$store.commit("loadWebInfo", res.result[0].data[0]);
          }
        })
        .catch((error) => {
          this.$notify({
            type: "error",
            title: "可恶🤬",
            message: error.message,
            position: "top-left",
            offset: 50,
          });
        });
    },
    getSortInfo() {
      this.$http
        .get(this.$constant.baseURL + "/webInfo/getSortInfo/")
        .then((res) => {
          if (!this.$common.isEmpty(res.result[0])) {
            const sortInfo = res.result[0].data.filter((item) => {
              return item.id !== 11;
            });
            this.$store.commit("loadSortInfo", sortInfo);
          }
        })
        .catch((error) => {
          this.$notify({
            type: "error",
            title: "可恶🤬",
            message: error.message,
            position: "top-left",
            offset: 50,
          });
        });
    },
    changeColor() {
      this.isDark = !this.isDark;
      localStorage.setItem("theme", this.isDark);
      //false是黑夜，显示太阳图标
      if (!this.isDark) {
        document.documentElement.dataset.theme = localStorage.getItem(
          "themeColor"
        )
          ? "theme2-dark"
          : "dark";
        if (document.documentElement.dataset.theme === "theme2-dark") {
          const root = document.querySelector(":root");
          root.style.setProperty("--themeColor", "var(--sYellow)");
          localStorage.setItem("themeColor", "#ffc848");
          this.$common.getThemeRgb();
        }
        this.$notify({
          title: "可以啦🍨",
          message: "是要关灯睡觉了吗~~~",
          type: "success",
          offset: 50,
          position: "top-left",
        });
      } else {
        document.documentElement.dataset.theme = localStorage.getItem(
          "themeColor"
        )
          ? "theme2"
          : "light";
        if (document.documentElement.dataset.theme !== "theme2-dark") {
          const root = document.querySelector(":root");
          root.style.setProperty("--themeColor", "var(--sBlue)");
          localStorage.setItem("themeColor", "#425aef");
          this.$common.getThemeRgb();
        }
      }
    },
    toTop() {
      window.scrollTo({
        top: 0,
        behavior: "smooth",
      });
    },
    onScrollPage() {
      this.scrollTop =
        document.documentElement.scrollTop || document.body.scrollTop;
    },
    changeMouseAnimation() {
      this.mouseAnimation = !this.mouseAnimation;
      if (this.mouseAnimation) {
        this.$nextTick(() => {
          mousedown();
        });
      }
    },
    openPcGame() {
      this.disGame = true;
      // 随机获取一个游戏
      let index = Math.floor(Math.random() * this.pcGameList.length);
      if (index === 0 || index === 1) {
        index = 6;
      }
      this.game = "http://game.eean.cn/pc/game" + index++;
    },
    openChangeBg() {
      this.changeBgBox = true;
    },
    changeBg(item) {
      // 刷新时触发并且没有本地缓存的背景，也没有点击切换背景
      if (!item && !JSON.parse(localStorage.getItem("vuex")).changeBg) {
        const dom = document.querySelector(".background-image-changeBg");
        const storeInfo = this.$store.state.changeBg;
        const string = storeInfo.split("#");
        if (string[0] === "") {
          dom.setAttribute(
            "style",
            `background-color: ${this.$store.state.changeBg}`
          );
        } else if (string.length === 1) {
          dom.setAttribute(
            "style",
            `background-image: ${this.$store.state.changeBg};background-size: cover;background-attachment: local;background-position: center;width: 100%;height: 100%;`
          );
        } else {
          dom.setAttribute(
            "style",
            `background-image: ${this.$store.state.changeBg}`
          );
        }
      } else if (item) {
        //只有点击切换时触发
        const dom = document.querySelector(".background-image-changeBg");
        const string = item.split("#");
        // 纯色
        if (string[0] === "") {
          this.$store.commit("changeBgBox", item);
          dom.setAttribute(
            "style",
            `background-color: ${this.$store.state.changeBg}`
          );
        } else if (string.length === 1) {
          // 图片
          const changeBgUrl = "url(" + item + ")";
          this.$store.commit("changeBgBox", changeBgUrl);
          dom.setAttribute(
            "style",
            `background-image: ${this.$store.state.changeBg};background-size: cover;background-attachment: local;background-position: center;width: 100%;height: 100%;`
          );
        } else {
          // 渐变
          const changeBgUrl = "linear-gradient(" + item + ")";
          this.$store.commit("changeBgBox", changeBgUrl);
          dom.setAttribute(
            "style",
            `background-image: ${this.$store.state.changeBg}`
          );
        }
      }
    },
    defaultBtn() {
      document.documentElement.dataset.theme = "light";
      localStorage.removeItem("themeColor");
      const dom = document.querySelector(".background-image-changeBg");
      dom.setAttribute(
        "style",
        "background-image: url(https://www.qiniuyun.zjh2002.icu/images/changeBg3);background-size: cover;background-attachment: local;background-position: center;width: 100%;height: 100%;"
      );
      this.$store.commit(
        "changeBgBox",
        "url(https://www.qiniuyun.zjh2002.icu/images/changeBg3)"
      );
    },
    handleChangeBg(val, i) {
      if (val === "pc" && this.themeMap[i].dataList.length === 0) {
        this.getAllBg(i);
      }
      if (val === "mobile" && this.themeMap[i].dataList.length === 0) {
        const pagination = {
          current: 1,
          size: 999,
          total: 0,
          resourceType: "mobilePhoto",
        };
        this.$http
          .post(
            this.$constant.baseURL + "/resource/listResource/",
            pagination,
            true
          )
          .then((res) => {
            if (!this.$common.isEmpty(res.result[0])) {
              res.result[0].records.forEach((item) => {
                this.themeMap[i].dataList.push(item.path);
              });
            }
          })
          .catch((error) => {
            this.$notify({
              type: "error",
              title: "可恶🤬",
              message: error.message,
              position: "top-left",
              offset: 50,
            });
          });
      }
      if (val === "gradient") {
        this.themeMap[i].dataList = this.$constant.gradient;
      }
      if (val === "solid") {
        this.themeMap[i].dataList = this.$constant.SolidColor;
      }
    },
    getAllBg(i) {
      const pagination = {
        current: 1,
        size: 999,
        total: 0,
        resourceType: "webBackgroundImage",
      };
      this.$http
        .post(
          this.$constant.baseURL + "/webInfo/allWebBackgroundImage/",
          pagination,
          true
        )
        .then((res) => {
          if (!this.$common.isEmpty(res.result[0])) {
            res.result[0].records.forEach((item) => {
              this.themeMap[i].dataList.push(item.path);
            });
          }
        })
        .catch((error) => {
          this.$notify({
            type: "error",
            title: "可恶🤬",
            message: error.message,
            position: "top-left",
            offset: 50,
          });
        });
    },
    EDIT() {
      if (this.editFlag) {
        this.$notify({
          title: "可以啦🍨",
          message: "已关闭编辑！！",
          type: "success",
          offset: 50,
          position: "top-left",
        });
        document.body.contentEditable = "false";
        this.editFlag = false;
        return;
      }
      this.$notify({
        title: "可以啦🍨",
        message: "可以随意修改本站的文字喔~~~不要干坏事！！",
        type: "success",
        offset: 50,
        position: "top-left",
      });
      document.body.contentEditable = "true";
      this.editFlag = true;
    },
  },
};
</script>
<style lang="scss">
.mk-popper {
  border-radius: 100px;
  padding: 6px 4px;
  background: var(--blue);
}
.popper__arrow {
  display: none !important;
}
</style>
<style lang="scss" scoped>
.small-menu {
  .el-dropdown-link {
    display: flex;
    align-items: center;
    font-size: 14px;
    img {
      width: 14px;
      height: 14px;
    }
    span {
      color: var(--red);
      margin-left: 6px;
    }
  }
  .menu_item {
    flex-wrap: wrap;
    margin: 0;
    width: 100%;
    display: flex;
    justify-content: space-between;
    flex-direction: row;
    margin-top: 8px;
    .my-menu {
      font-size: 15px;
      line-height: 27px;
      img {
        width: 15px;
        height: 15px;
      }
    }
    > span {
      width: calc(50% - 6px);
      margin: 0;
      border: 1px solid var(--blue);
      border-radius: 8px;
      padding: 4px 15px;
      margin-bottom: 8px;
      transition: all 0.3s;
      &:hover {
        border: 1px solid var(--red);
        background: var(--red);
        color: var(--white);
      }
    }
  }
}
.mk-popover_item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  li {
    list-style: none;
    padding: 8px;
    font-size: 16px;
    color: var(--red);
    &:hover {
      background: var(--red);
      color: var(--white);
      padding: 8px;
      border-radius: 100px;
      animation: slide-top 0.5s ease-in-out both;
    }
  }
}
.customImg {
  font-size: 18px;
  margin-top: 15px;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  &-item {
    color: var(--blue6);
  }
  .color-box_contain {
    display: flex;
    align-items: center;
    justify-content: space-between;
    flex-wrap: wrap;
    .color-box {
      font-size: 14px;
      width: 49%;
      padding: 9px 10px 9px 4px;
      display: flex;
      align-items: center;
      border: 2px dashed;
      border-radius: 0.5rem;
      margin-bottom: 10px;
      transition: all 0.3s ease;
      position: relative;
      &.active:before {
        position: absolute;
        right: 8px;
        bottom: 22px;
        font-size: 24px;
        content: "🌻";
        transform: scale(1);
      }
      &:hover {
        transform: scale(0.9);
      }
      .gun {
        width: 8px;
        height: 50px;
        border-radius: 0.375rem;
      }
      .center {
        padding-left: 10px;
        .top {
          font-weight: 400;
          line-height: 24px;
          text-align: left;
          margin-bottom: 4px;
        }
        .bottom {
          font-weight: 400;
          line-height: 24px;
        }
      }
    }
  }
}
.el-popper[x-placement^="bottom"] {
  margin-top: -6px;
}
.httpButton {
  margin: 10px auto 0;
  background: var(--blue5);
  display: block;
  width: 25%;
  padding: 15px 0;
  border-radius: 30px;
  color: var(--white2);
  border: none;
  font-size: 15px;
  &:hover {
    background-color: var(--orange4);
    transition: all 0.3s ease;
  }
}
.setting-color {
  color: var(--bigRed);
  animation: rotate 4s linear infinite;
}
.toolbar-content {
  padding: 0 40px;
  top: 0;
  width: 100%;
  height: 50px;
  color: var(--red);
  position: fixed;
  z-index: 99999;
  user-select: none;
  transition: all 0.3s ease;
  &.enter {
    background: var(--background);
    color: var(--red);
    box-shadow: 0 1px 3px 0 var(--miniMask);
  }
  &.hoverEnter {
    background: var(--background);
    box-shadow: 0 1px 3px 0 var(--miniMask);
  }
}
.toolbar-title {
  display: flex;
  justify-content: center;
  align-items: center;
  transition: 0.3s;
  height: 50px;
  &:hover {
    background-color: var(--red);
    color: var(--red);
    border-radius: 8px;
    &:after {
      opacity: 1;
      transform: translateY(0) scale(0.3);
      transition-timing-function: ease-in;
    }
  }
  &:after {
    background: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAMAAACdt4HsAAAABGdBTUEAALGPC/xhBQAAAAFzUkdCAK7OHOkAAABgUExURUxpcf///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////05o8T8AAAAfdFJOUwCqxd+3iU1VQL+Uz7IrG+mO5TN4EVrvRditZp8S9oBdLCSzAAABAUlEQVRYw+3WyRKCMAyA4bKjgLjv2vd/S5FBIF2T9OKBXMP/DTidsUI4J4/rXARMIbspAvsAYejZwtgzhVnPEkDPEJSeLGg9UTD0JMHYEwRLjxasPVJw9CjB2SMET+8VvL1HQPROAdU7BGRvFdC9RSD0RoHUGwRirwnkXhEYPRBY/Uxg9qPA7gchoO+FXMowQQaOiMP6WFx2IX1Sdr9i+R1622fTWWJ8PJwFWIAF+F+g2pbjHFI6kMH9gwpE6r9vRgT2KtAQgaN2gUhogH4DWS2AD4jgttWBGj6RKus7XK91QDlJJ3UPT/tTBzbukyrE+/pbntPGdJXd3NrhgVc1veEHPLUmePVAutQAAAAASUVORK5CYII=)
      no-repeat 50% 50%;
    opacity: 0;
    position: absolute;
    z-index: 1;
    width: 75px;
    height: 58px;
    content: "";
    transition: 0.3s;
    transition-timing-function: ease-in;
    transform: scale(0.4);
  }
}
.toolbar-mobile-menu {
  font-size: 30px;
  margin-right: 15px;
}
.scroll-menu {
  color: var(--red1);
  display: flex;
  justify-content: flex-end;
  padding: 0;
  > li {
    list-style: none;
    margin-right: 24px;
    font-size: 17px;
    height: 50px;
    line-height: 50px;
    position: relative;
    text-decoration: none;
    background: var(--gradientAnimation);
    background-size: 0px 3px;
    transition: background-size 800ms;
    &:hover {
      color: var(--orange2);
      background-position-x: left;
      background-size: 100% 3px;
    }
    &:last-child {
      margin-right: 0;
    }
  }
}
.el-dropdown {
  font-size: unset;
  color: unset;
  &-menu {
    align-items: center;
    border-radius: 13px;
    padding: 6px;
    border: 0;
    background: var(--blue);
    &__item {
      font-size: unset;
      line-height: 26px;
      padding: 10px;
      color: var(--red);
      &:hover {
        border-radius: 8px;
        animation: slide-top 0.5s ease-in-out both;
        background-color: var(--red);
        color: var(--white);
      }
    }
    &.active {
      display: flex;
    }
  }
}
.toolButton {
  position: fixed;
  right: 3vh;
  bottom: 3vh;
  animation: slide-bottom 0.5s ease-in-out both;
  z-index: 100;
  font-size: 25px;
}
.my-setting {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
  font-size: 20px;
  i {
    padding: 5px;
    color: var(--white);
    &:hover {
      color: var(--bigRed2);
    }
  }
  .active {
    color: var(--orange3);
  }
}
.user:hover {
  color: var(--white);
}
.login:hover {
  color: var(--white);
}
.logout:hover {
  color: var(--white);
}
.cd-top {
  background: var(--toTop) no-repeat center;
  position: fixed;
  right: 5vh;
  top: -900px;
  z-index: 99;
  width: 70px;
  height: 900px;
  background-size: contain;
  transition: all 0.5s ease-in-out;
}
.backTop {
  transition: all 0.3s ease-in;
  position: relative;
  top: 0;
  left: -4px;
  &:hover {
    top: -10px;
  }
}
::v-deep .el-drawer__body {
  background: linear-gradient(60deg, var(--pink1) 20%, var(--yellow2) 93%);
  z-index: 999;
}
.is-center {
  text-align: center;
}
.sidebar {
  background-image: var(--toolbar);
  background-position: top;
  background-size: 120%;
  background-repeat: no-repeat;
}
.avatar-img {
  overflow: hidden;
  margin: 0 auto;
  width: 110px;
  height: 110px;
  border-radius: 25px;
  box-shadow: 2.2px 2.2px 2.2px var(--toolbarBackground);
  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
}
.author-info_name {
  margin-top: 10px;
  color: var(--blue);
  font-weight: 800;
  font-size: 25.2px;
}
.author-info__description {
  color: var(--blue);
  margin-top: 20px;
  font-size: 18px;
  padding-bottom: 8px;
}
.site-data {
  margin-top: 20px;
  color: var(--blue);
  width: 100%;
  display: flex;
  flex-direction: row;
  justify-content: space-around;
}
.blog-info-box {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-around;
}
.blog-info-num {
  margin-top: 12px;
}
hr {
  position: relative;
  margin: 20px auto 10px;
  border: 2px dashed var(--blue);
  overflow: visible;
  &:before {
    position: absolute;
    top: -21px;
    left: 5%;
    color: var(--red);
    content: "\e673";
    font-size: 40px;
    line-height: 1;
    transition: all 1s ease-in-out;
    font-family: iconfont;
  }
  &:hover:before {
    left: calc(95% - 20px);
  }
}
.bgBox {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  .box {
    width: 166px;
    margin: 10px;
    height: 100px;
    border-radius: 6px;
    border-bottom: none;
    background-size: cover;
    transition: all 1s;
    &:hover {
      transform: scale(1.1);
    }
  }
  .mobileBox {
    height: 240px;
  }
}
.button {
  opacity: 0.6;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 38px;
  height: 38px;
  font-size: 15px;
  z-index: 20;
  position: relative;
  background-color: var(--blue3);
  color: var(--bigRed);
  text-align: center;
  line-height: 35px;
  margin-top: 8px;
}
::v-deep #pic-link {
  margin: 15px auto 0;
  width: 80%;
  border-radius: 30px;
  border: 1px solid var(--blue6);
  padding: 5px 10px 5px 10px;
  line-height: 2;
  outline: 1px solid var(--blue6);
  &:hover {
    outline-color: var(--blue5);
    border-color: var(--blue5);
  }
  &:focus {
    outline-color: var(--blue5);
    border-color: var(--blue5);
  }
}
::v-deep .el-drawer__header {
  padding: 0;
  margin: 0;
}
::v-deep .changeBgBox {
  background: var(--favoriteBg);
  border-radius: 14px;
  overflow: scroll;
  box-shadow: 0 14px 28px var(--mask), 0 10px 10px var(--miniMask);
  height: 500px;
  &::-webkit-scrollbar {
    width: 0px;
  }
  .el-dialog__header {
    top: 0;
    z-index: 999;
    position: sticky;
    overflow: hidden;
    padding: 5px 20px 10px;
    background: var(--blue9);
    color: var(--favoriteBg);
    .el-dialog__headerbtn i {
      font-size: 24px;
      color: var(--favoriteBg);
      position: relative;
      bottom: 10px;
      right: 0;
    }
    .el-dialog__title {
      color: var(--favoriteBg);
    }
  }
  .el-dialog__body {
    padding: 10px 15px;
  }
}
::v-deep .el-collapse {
  flex: 1;
  margin-left: 10px;
  border: none;
  &-item__header {
    background: var(--gray5);
    padding: 16px;
    border-radius: 6px;
    color: var(--fontColor);
    font-size: 0.875rem;
    font-weight: 400;
    position: relative;
    border-bottom-color: var(--toolbarBackground1);
    line-height: normal;
  }
  &-item__wrap {
    border-radius: 6px;
  }
}
@media screen and (max-width: 400px) {
  .toolButton {
    right: 0.5vh;
  }
}
@media screen and (max-width: 510px) {
  .bgBox .box {
    height: 73px;
    width: 135px;
  }
  .bgBox .mobileBox {
    height: 240px;
    width: 135px;
  }
  .iconRotate.showIcon {
    display: none;
  }
}
@media screen and (max-width: 950px) {
  .center-toolbar {
    display: none;
  }
}
</style>
