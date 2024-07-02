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
        <div v-else>
          <ul class="scroll-menu">
            <!-- 文章 -->
            <li>
              <el-popover
                close-delay="100"
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
                close-delay="100"
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
                  <!-- 藏宝 -->
                  <li @click="$router.push({ path: '/tools' })">
                    <img
                      style="vertical-align: -3px"
                      src="../assets/svg/treasure.svg"
                    />
                    藏宝
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
                close-delay="100"
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
                close-delay="100"
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
            <!-- 随机文章 -->
            <li
              title="随便看看-随机访问一篇本站文章"
              style="color: black; font-size: 20px"
              @click="openRandomArticle"
            >
              <div class="my-menu"><i class="fa fa-grav"></i></div>
            </li>
            <!-- 切换背景 -->
            <li
              title="切换背景-换一种背景，换一种感觉。"
              id="changeThemeRef"
              style="color: black; font-size: 20px"
              @click="openChangeBg"
            >
              <div class="my-menu"><i class="fa fa-image"></i></div>
            </li>
            <!-- 关闭樱花 -->
            <li @click="handleSakura()" style="color: black; font-size: 21px">
              <div class="my-menu"><i class="fa fa-pagelines"></i></div>
            </li>
            <!-- 关灯 -->
            <li
              style="transform: scale(0.16); width: 70px; height: 0"
              class="my-menu"
            >
              <switchBtn @click.native="changeColor"></switchBtn>
            </li>
            <!-- 个人中心 -->
            <li>
              <el-dropdown placement="bottom">
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
      <!-- 主题切换按钮 -->
      <el-popover placement="left" :close-delay="500" trigger="hover">
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
        <!-- 设置主题 -->
        <div class="my-setting">
          <el-tooltip placement="top" effect="light">
            <div slot="content">(*^▽^*)切换主题咯！</div>
            <div>
              <!-- 太阳按钮 -->
              <i
                v-if="!isDark"
                class="el-icon-sunny iconRotate"
                @click="changeColor"
              ></i>
              <!-- 月亮按钮 -->
              <i
                v-else
                class="fa fa-moon-o"
                aria-hidden="true"
                @click="changeColor"
              ></i>
            </div>
          </el-tooltip>
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
          <ul class="small-menu">
            <!-- 文章 -->
            <li>
              <el-dropdown placement="bottom-end">
                <span class="el-dropdown-link">
                  <img
                    style="vertical-align: -3px"
                    src="../assets/svg/document.svg"
                  />
                  文章<i class="el-icon-arrow-down el-icon--right"></i>
                </span>
                <el-dropdown-menu slot="dropdown">
                  <el-dropdown-item>
                    <!-- 分类 -->
                    <li @click="$router.push({ path: '/sort' })">
                      <div class="my-menu">
                        <img
                          style="vertical-align: -3px"
                          src="../assets/svg/sort.svg"
                        />
                        分类
                      </div>
                    </li>
                  </el-dropdown-item>
                  <el-dropdown-item>
                    <!-- 标签 -->
                    <li
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
                    </li>
                  </el-dropdown-item>
                </el-dropdown-menu>
              </el-dropdown>
            </li>
            <!-- 空间 -->
            <li>
              <el-dropdown placement="bottom-end">
                <span class="el-dropdown-link">
                  <img
                    style="vertical-align: -3px"
                    src="../assets/svg/space.svg"
                  />
                  空间<i class="el-icon-arrow-down el-icon--right"></i>
                </span>
                <el-dropdown-menu slot="dropdown">
                  <el-dropdown-item>
                    <!-- 音乐 -->
                    <li @click="smallMenu({ path: '/funny' })">
                      <div>
                        <img
                          style="vertical-align: -3px"
                          src="../assets/svg/music.svg"
                        />
                        幻音坊
                      </div>
                    </li>
                  </el-dropdown-item>
                  <el-dropdown-item>
                    <!-- 藏宝 -->
                    <li @click="smallMenu({ path: '/tools' })">
                      <div>
                        <img
                          style="vertical-align: -3px"
                          src="../assets/svg/treasure.svg"
                        />
                        藏宝
                      </div>
                    </li>
                  </el-dropdown-item>
                </el-dropdown-menu>
              </el-dropdown>
            </li>
            <!-- 我的 -->
            <li>
              <el-dropdown placement="bottom-end">
                <span class="el-dropdown-link">
                  <img
                    style="vertical-align: -3px"
                    src="../assets/svg/home.svg"
                  />
                  我的<i class="el-icon-arrow-down el-icon--right"></i>
                </span>
                <el-dropdown-menu slot="dropdown">
                  <el-dropdown-item>
                    <!-- 相册 -->
                    <li @click="smallMenu({ path: '/travel' })">
                      <div>📸 <span>相册</span></div>
                    </li>
                  </el-dropdown-item>
                  <el-dropdown-item>
                    <!-- 爱链 -->
                    <li @click="smallMenu({ path: '/love' })">
                      <div>
                        <img
                          style="vertical-align: -3px"
                          src="../assets/svg/love.svg"
                        />
                        爱链
                      </div>
                    </li>
                  </el-dropdown-item>
                  <el-dropdown-item>
                    <li @click="openPcGame">
                      <div class="my-menu">🎮 <span>小游戏</span></div>
                    </li>
                  </el-dropdown-item>
                  <el-dropdown-item>
                    <li @click="EDIT()">
                      <div class="my-menu">
                        <img
                          style="vertical-align: -5px"
                          src="../assets/svg/pencil.svg"
                        />
                        编辑怪
                      </div>
                    </li>
                  </el-dropdown-item>
                  <el-dropdown-item>
                    <!-- 关于 -->
                    <li @click="smallMenu({ path: '/about' })">
                      <div>🐷 <span>关于我</span></div>
                    </li>
                  </el-dropdown-item>
                </el-dropdown-menu>
              </el-dropdown>
            </li>
            <!-- 社交 -->
            <li>
              <el-dropdown placement="bottom-end">
                <span class="el-dropdown-link">
                  <img
                    style="vertical-align: -3px"
                    src="../assets/svg/socialContact.svg"
                  />
                  社交<i class="el-icon-arrow-down el-icon--right"></i>
                </span>
                <el-dropdown-menu slot="dropdown">
                  <el-dropdown-item>
                    <!-- 留言厅 -->
                    <li @click="smallMenu({ path: '/message' })">
                      <div>✍🏻 <span>留言厅</span></div>
                    </li>
                  </el-dropdown-item>
                  <el-dropdown-item>
                    <!-- 友链 -->
                    <li @click="smallMenu({ path: '/friend' })">
                      <div>🎀 <span>友链</span></div>
                    </li>
                  </el-dropdown-item>
                </el-dropdown-menu>
              </el-dropdown>
            </li>
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
              style="color: black; font-size: 20px"
              @click="openRandomArticle"
            >
              <div class="my-menu"><i class="fa fa-grav"></i></div>
            </li>
            <!-- 切换背景 -->
            <li style="color: black; font-size: 20px" @click="openChangeBg">
              <div class="my-menu"><i class="fa fa-image"></i></div>
            </li>
            <!-- 关灯 -->
            <li @click="changeColor" style="color: black; font-size: 23px">
              <div class="my-menu"><i class="fa-adjust fa"></i></div>
            </li>
            <!-- 关闭樱花 -->
            <li @click="handleSakura()" style="color: black; font-size: 21px">
              <div class="my-menu"><i class="fa fa-pagelines"></i></div>
            </li>
          </ul>
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
            background: #5fcdffc8;
            display: block;
            width: 100%;
            padding: 15px 0;
            border-radius: 6px;
            color: white;
            border: none;
            font-size: 18px;
          "
        >
          <i
            style="font-size: 18px; margin-right: 5px"
            class="fa fa-refresh"
          ></i
          >点我恢复默认背景
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
        <div v-for="(items, index) in themeMap" :key="index">
          <h2 style="text-align: left">{{ items.title }}</h2>
          <div style="display: flex; align-items: center">
            <span class="iconRotate showIcon">
              <img src="../assets/svg/windmill.svg" />
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
      //黑夜，所以显示太阳按钮
      isDark: false,
      scrollTop: 0,
      toolbarDrawer: false,
      mobile: false,
      changeBgBox: false,
      editFlag: false,
      sakuraFlag: false,
      topPercentageType: false,
      themeMap: this.$constant.themeMapConfig,
      httpInput: "",
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
    const root = document.querySelector(":root");
    // 关灯
    if (localStorage.getItem("theme") == "false") {
      this.isDark = false;
      root.style.setProperty("--background", "#272727");
      root.style.setProperty("--fontColor", "white");
      root.style.setProperty("--borderColor", "#4F4F4F");
      root.style.setProperty("--borderHoverColor", "black");
      root.style.setProperty("--articleFontColor", "#E4E4E4");
      root.style.setProperty("--articleGreyFontColor", "#D4D4D4");
      root.style.setProperty("--commentContent", "#D4D4D4");
      root.style.setProperty("--favoriteBg", "#1e1e1e");
    } else {
      this.isDark = true;
      root.style.setProperty("--background", "rgba(255,255,255,0.4)");
      root.style.setProperty("--fontColor", "#1f2d3d");
      root.style.setProperty("--borderColor", "rgba(0, 0, 0, 0.5)");
      root.style.setProperty("--borderHoverColor", "rgba(110, 110, 110, 0.4)");
      root.style.setProperty("--articleFontColor", "#1F1F1F");
      root.style.setProperty("--articleGreyFontColor", "#616161");
      root.style.setProperty("--commentContent", "#F7F9FE");
      root.style.setProperty("--favoriteBg", "#f7f9fe");
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
    this.mobile = document.body.clientWidth < 1286;
    window.addEventListener("resize", () => {
      let docWidth = document.body.clientWidth;
      if (docWidth < 1286) {
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
        window.test();
        this.sakuraFlag = false;
      } else {
        const dom = document.querySelector("#canvas_sakura");
        if (dom) {
          dom.remove();
          this.sakuraFlag = true;
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
        this.$message({
          message: "请先登录！",
          type: "error",
        });
      } else {
        let userToken = this.$common.encrypt(localStorage.getItem("userToken"));
        window.open(this.$constant.imBaseURL + "?userToken=" + userToken);
      }
    },
    logout() {
      this.$message({
        message: "退出登录成功！",
        type: "success",
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
          this.$message({
            message: error.message,
            type: "error",
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
          this.$message({
            message: error.message,
            type: "error",
          });
        });
    },
    changeColor() {
      if (
        this.$router.currentRoute.path === "/tools" ||
        this.$router.currentRoute.path === "/about"
      ) {
        this.$message({
          message: "遭了个糕，这个页面主人不让切换主题wo~~~",
          type: "error",
        });
        return;
      }
      this.isDark = !this.isDark;
      localStorage.setItem("theme", this.isDark);
      const roots = document.querySelector(":root");
      //false是黑夜，显示太阳图标
      if (!this.isDark) {
        roots.style.setProperty("--background", "#272727");
        roots.style.setProperty("--fontColor", "white");
        roots.style.setProperty("--borderColor", "#4F4F4F");
        roots.style.setProperty("--borderHoverColor", "black");
        roots.style.setProperty("--articleFontColor", "#E4E4E4");
        roots.style.setProperty("--articleGreyFontColor", "#D4D4D4");
        roots.style.setProperty("--commentContent", "#D4D4D4");
        roots.style.setProperty("--favoriteBg", "#1e1e1e");
      } else {
        roots.style.setProperty("--background", "rgba(255,255,255,0.4)");
        roots.style.setProperty("--fontColor", "#1f2d3d");
        roots.style.setProperty("--borderColor", "rgba(0, 0, 0, 0.5)");
        roots.style.setProperty(
          "--borderHoverColor",
          "rgba(110, 110, 110, 0.4)"
        );
        roots.style.setProperty("--articleFontColor", "#1F1F1F");
        roots.style.setProperty("--articleGreyFontColor", "#616161");
        roots.style.setProperty("--commentContent", "#F7F9FE");
        roots.style.setProperty("--favoriteBg", "#f7f9fe");
      }
      if (!this.isDark) {
        this.$message({
          message: "是要关灯睡觉了吗~~~",
          type: "success",
        });
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
      const dom = document.querySelector(".background-image-changeBg");
      dom.setAttribute(
        "style",
        "background-image: url(https://www.qiniuyun);background-size: cover;background-attachment: local;background-position: center;width: 100%;height: 100%;"
      );
      this.$store.commit(
        "changeBgBox",
        "url(https://www.qiniuyun)"
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
            this.$message({
              message: error.message,
              type: "error",
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
          this.$message({
            message: error.message,
            type: "error",
          });
        });
    },
    EDIT() {
      if (this.editFlag) {
        this.$message({
          message: "已关闭编辑！！",
          type: "success",
        });
        document.body.contentEditable = "false";
        this.editFlag = false;
        return;
      }
      this.$message({
        message: "可以随意修改本站的文字喔~~~不要干坏事！！",
        type: "success",
      });
      document.body.contentEditable = "true";
      this.editFlag = true;
    },
  },
};
</script>
<style>
.mk-popper {
  border-radius: 8px;
  padding: 4px;
  background: rgba(116, 189, 240, 0.7);
}
</style>
<style scoped>
.mk-popover_item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.mk-popover_item li {
  list-style: none;
  padding: 8px;
  font-size: 16px;
  color: var(--red);
}
.mk-popover_item li:hover {
  background: var(--red);
  color: white;
  padding: 8px;
  border-radius: 8px;
  animation: slide-top 0.5s ease-in-out both;
}
.customImg {
  font-size: 18px;
  margin-top: 15px;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}
.customImg-item {
  color: #00aeffc8;
}
.httpButton {
  margin: 10px auto 0;
  background: rgb(66, 90, 239);
  display: block;
  width: 25%;
  padding: 15px 0;
  border-radius: 30px;
  color: white;
  border: none;
  font-size: 15px;
}
.httpButton:hover {
  background-color: #ff7242;
  transition: all 0.2s ease-in-out;
}
.setting-color {
  color: var(--bigRed);
  animation: rotate 4s linear infinite;
}
.toolbar-content {
  width: 100%;
  height: 50px;
  color: var(--red);
  /* 固定位置，不随滚动条滚动 */
  position: fixed;
  z-index: 100;
  /* 禁止选中文字 */
  user-select: none;
  transition: all 0.3s ease-in-out;
}
.toolbar-content.enter {
  background: var(--toolbarBackground);
  color: var(--red);
  box-shadow: 0 1px 3px 0 rgba(0, 34, 77, 0.05);
}
.toolbar-content.hoverEnter {
  background: var(--translucent);
  box-shadow: 0 1px 3px 0 rgba(0, 34, 77, 0.05);
}
.toolbar-title {
  display: flex;
  justify-content: center;
  align-items: center;
  transition: 0.3s;
  height: 50px;
  margin-left: 30px;
}
.toolbar-title:hover {
  background-color: var(--red);
  color: var(--red);
  border-radius: 8px;
}
.toolbar-title:hover:after {
  opacity: 1;
  transform: translateY(0) scale(0.3);
  transition-timing-function: ease-in;
}
.toolbar-title:after {
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
.toolbar-mobile-menu {
  font-size: 30px;
  margin-right: 15px;
}
.scroll-menu {
  display: flex;
  justify-content: flex-end;
  padding: 0;
}
.scroll-menu > li {
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
}
.scroll-menu > li:hover {
  color: orange;
  background-position-x: left;
  background-size: 100% 3px;
}
.el-dropdown {
  font-size: unset;
  color: unset;
}
.el-popper[x-placement^="bottom"] {
  margin-top: -6px;
}
.el-dropdown-menu {
  align-items: center;
  border-radius: 20px;
  padding: 0px;
  border: 0;
  background: rgba(116, 189, 240, 0.7);
}
.el-dropdown-menu__item {
  font-size: unset;
  line-height: 26px;
  padding: 10px;
  color: var(--red);
}
.el-dropdown-menu__item:hover {
  border-radius: 20px;
  animation: slide-top 0.5s ease-in-out both;
  background-color: var(--red);
  color: white;
}
.el-dropdown-menu.active {
  display: flex;
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
}
.my-setting i {
  padding: 5px;
  color: white;
}
.my-setting i:hover {
  color: var(--bigRed);
}
.my-setting .active {
  color: orange;
}
.user:hover {
  color: white;
}
.login:hover {
  color: white;
}
.logout:hover {
  color: white;
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
}
.backTop:hover {
  top: -10px;
}
@media screen and (max-width: 400px) {
  .toolButton {
    right: 0.5vh;
  }
}
::v-deep .el-drawer__body {
  background: linear-gradient(60deg, #ffd7e4 0, #c8f1ff 93%);
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
  box-shadow: 2.2px 2.2px 2.2px rgba(10, 207, 233, 0.3);
}
.avatar-img img {
  width: 100%;
  height: 100%;
  object-fit: cover;
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
}
hr:before {
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
hr:hover:before {
  left: calc(95% - 20px);
}
.bgBox {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
}
.bgBox .box {
  width: 166px;
  margin: 10px;
  height: 100px;
  border-radius: 6px;
  border-bottom: none;
  background-size: cover;
}
.bgBox .mobileBox {
  height: 240px;
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
  background-color: #49b1f5;
  color: var(--bigRed);
  text-align: center;
  line-height: 35px;
  margin-top: 8px;
}
::v-deep #pic-link {
  margin: 15px auto 0;
  width: 80%;
  border-radius: 30px;
  border: 1px solid rgb(66, 90, 239);
  padding: 5px 10px 5px 10px;
  line-height: 2;
  outline: 1px solid rgb(66, 90, 239);
}
::v-deep .el-drawer__header {
  padding: 0;
  margin: 0;
}
::v-deep .changeBgBox {
  background: rgba(255, 255, 255, 0.85);
  border-radius: 14px;
  overflow: scroll;
  box-shadow: 0 14px 28px rgba(0, 0, 0, 0.25), 0 10px 10px rgba(0, 0, 0, 0.22);
  height: 500px;
}
::v-deep .changeBgBox::-webkit-scrollbar {
  width: 0px;
}
::v-deep .changeBgBox .el-dialog__header {
  top: 0;
  position: sticky;
  overflow: hidden;
  padding: 5px 20px 10px;
  background: rgba(27, 188, 216, 0.68);
  color: white;
}
::v-deep .changeBgBox .el-dialog__header .el-dialog__headerbtn i {
  font-size: 24px;
  color: white;
  position: relative;
  bottom: 10px;
  right: 0;
}
::v-deep .changeBgBox .el-dialog__header .el-dialog__title {
  color: white;
}
::v-deep .changeBgBox .el-dialog__body {
  padding: 10px 15px;
}
::v-deep .el-collapse {
  flex: 1;
  margin-left: 10px;
  border: none;
}
::v-deep .el-collapse-item__header {
  background: #e8fafe;
  padding: 16px;
  border-radius: 6px;
  color: rgba(68, 68, 68, 0.7);
  font-size: 0.875rem;
  font-weight: 400;
  position: relative;
  border-bottom-color: rgba(27, 205, 252, 0.3);
  line-height: normal;
}
::v-deep .el-collapse-item__wrap {
  border-radius: 6px;
}
::v-deep .el-collapse-item__header:hover {
  color: #444;
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
</style>
