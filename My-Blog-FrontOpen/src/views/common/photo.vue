<template>
  <div
    class="card-container photoRef"
    v-if="!$common.isEmpty(resourcePathList)"
  >
    <div
      v-for="(resourcePath, index) in resourcePathList"
      :key="index"
      class="card-item wow shadow-box-mini"
    >
      <div class="card-image container">
        <el-image
          :style="{
            opacity: hoverIndex === -1 ? 1 : index === hoverIndex ? 1 : 0.2,
          }"
          crossorigin="anonymous"
          @mouseenter="handleMouseEnter($event.target, index)"
          @mouseleave="handleMouseLeave"
          class="my-el-image"
          v-once
          lazy
          :preview-src-list="[resourcePath.cover]"
          :src="resourcePath.cover"
          fit="cover"
        >
          <div slot="error" class="image-slot">
            <div class="error-aside-image">
              <div class="error-text">肥肠抱歉，图片跑掉了┮﹏┭</div>
            </div>
          </div>
        </el-image>
      </div>
      <div class="card-body">
        <el-tooltip placement="bottom-start" effect="light">
          <div slot="content">
            {{ resourcePath.title }}
          </div>
          <div class="card-desc">
            <img style="vertical-align: -2px" src="../../assets/svg/leaf.svg" />
            {{ resourcePath.title }}
          </div>
        </el-tooltip>
        <div class="card-time">
          Date: {{ $common.getDateDiff(resourcePath.createTime) }}
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import ColorThief from "colorthief";
export default {
  props: {
    resourcePathList: {
      type: Array,
    },
  },
  data() {
    return {
      hoverIndex: -1,
    };
  },
  mounted() {
    this.handleMouseLeave();
  },
  methods: {
    async handleMouseEnter(img, i) {
      const colorThief = new ColorThief();
      const html = document.documentElement;
      this.hoverIndex = i;
      // 得到这张图片的调色盘（前三种主要颜色）
      const colors = await colorThief.getPalette(img, 3);
      const [c1, c2, c3] = colors.map((c) => `rgb(${c[0]},${c[1]},${c[2]})`);
      if (this.$route.path == "/travel") {
        const photoRef = document.querySelector(".photoRef");
        photoRef.style.setProperty("--background", "var(--transparent)");
        photoRef.style.setProperty("--c1", c1);
        photoRef.style.setProperty("--c2", c2);
        photoRef.style.setProperty("--c3", c3);
      } else {
        const root = document.querySelector(":root");
        root.style.setProperty("--background", "var(--transparent)");
        html.style.setProperty("--c1", c1);
        html.style.setProperty("--c2", c2);
        html.style.setProperty("--c3", c3);
      }
    },
    handleMouseLeave() {
      this.hoverIndex = -1;
    },
  },
};
</script>
<style lang="scss" scoped>
.card-container {
  display: flex;
  flex-wrap: wrap;
  border-radius: 1.5rem;
}
.card-item {
  position: relative;
  overflow: hidden;
  margin: 15px;
  height: 510px;
  flex-shrink: 0;
  width: calc(100% / 3 - 30px);
  animation: zoomIn 0.8s ease-in-out;
  padding: 1.3rem 1rem 1.5rem;
  background: var(--background);
  border-radius: 1.5rem;
  transition: all 0.2s;
}
.card-image {
  width: 100%;
  height: 400px;
  border-radius: 1rem;
  overflow: hidden;
  box-shadow: 0 -2px 10px var(--whiteMask);
  margin-bottom: 1rem;
  ::v-deep .el-image__inner {
    transition: all 1s;
    &:hover {
      transform: scale(1.2);
    }
  }
}
.card-body {
  padding: 10px 5px;
}
.card-desc {
  font-weight: 400;
  font-size: 1.05rem;
  color: var(--red1);
  letter-spacing: 1px;
  line-height: 1.5;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  &:hover {
    color: var(--red);
  }
}
.card-time {
  position: absolute;
  bottom: 15px;
  color: var(--green4);
  font-weight: 500;
  &:hover {
    color: var(--green1);
  }
}
.error-aside-image {
  background: var(--gradientAnimation);
  color: var(--white);
  text-align: center;
  height: 100%;
}
.error-text {
  color: var(--wheat1);
  position: relative;
  top: 50%;
  transform: translate(0, -50%);
}
@media screen and (max-width: 1300px) {
  .card-item {
    width: calc(100% / 2 - 30px);
  }
}
@media screen and (max-width: 1000px) {
  .card-item {
    height: 410px;
  }
  .card-image {
    height: 300px;
  }
}
@media screen and (max-width: 750px) {
  .card-item {
    width: 100%;
    margin: 10px 0;
  }
}
@media screen and (max-width: 450px) {
  .card-item {
    height: 360px;
  }
  .card-image {
    height: 250px;
  }
}
</style>
