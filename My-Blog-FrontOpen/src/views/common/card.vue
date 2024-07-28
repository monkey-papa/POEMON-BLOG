<template>
  <div class="card-container" v-if="!$common.isEmpty(resourcePathList)">
    <div v-for="(resourcePath, index) in resourcePathList" :key="index" class="card-item shadow-box wow" @click="clickResourcePath(resourcePath)">
      <div class="card-image">
        <el-image class="my-el-image" v-once lazy :src="resourcePath.cover" fit="cover">
          <div slot="error" class="image-slot myCenter" style="background: var(--gradualRed)">
            <div class="error-text">
              <div style="color:var(--wheat1);">肥肠抱歉，图片跑掉了</div>
            </div>
          </div>
        </el-image>
      </div>
      <div class="card-body">
        <div class="card-title">
          <el-image style="width: 32px;height: 32px;margin-right: 4px;border-radius: 11px;" v-once lazy :src="resourcePath.friendAvatar" fit="cover">
            <div slot="error" class="image-slot myCenter" style="height:32px;width:32px;background-image: var(--defaultAvatar);background-size: cover;">
            </div>
          </el-image>
          {{ resourcePath.title }}
        </div>
        <div class="card-desc">
          {{ resourcePath.introduction }}
        </div>
        <div class="card-time">
          <i class="el-icon-loading" style="
            margin-right: 5px;
            color: var(--maxLightRed);
            font-size: 14px;
          "></i>
          添加于 {{ $common.getDateDiff(resourcePath.createTime) }}
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    resourcePathList: {
      type: Array
    }
  },
  methods: {
    clickResourcePath(resourcePath) {
      this.$emit('clickResourcePath', resourcePath.url)
    }
  }
}
</script>
<style lang="scss" scoped>
.card-container {
  display: flex;
  flex-wrap: wrap;
}
.card-item {
  position: relative;
  border-radius: 10px;
  background: var(--gray5);
  overflow: hidden;
  margin: 10px;
  height: 300px;
  flex-shrink: 0;
  width: calc(100% / 4 - 20px);
  animation: zoomIn 0.8s ease-in-out;
  border: 1px solid var(--gray1);
  transition: all 0.3s ease;
  &:hover {
    border-color: var(--gray4);
  }
}
.card-image {
  width: 100%;
  height: 180px;
  ::v-deep .el-image__inner {
    transition: all 1s;
    &:hover {
      transform: scale(1.2);
    }
  }
}
.card-body {
  padding: 10px 20px;
}
.card-title {
  color: var(--bigRed1);
  font-size: 18px;
  font-weight: 400;
  margin-bottom: 10px;
  white-space: nowrap;
  text-overflow: ellipsis;
  overflow: hidden;
  transition: all 0.3s ease;
  &:hover {
    color: var(--black8);
  }
}
.card-desc {
  color: var(--darkBlue);
  font-size: 14px;
  line-height: 1.5;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  letter-spacing: 1px;
  transition: all 0.3s ease;
  &:hover {
    color: var(--black8);
  }
}
.card-time {
  position: absolute;
  bottom: 4px;
  font-size: 12px;
  color: var(--bigRed1);
  transition: all 0.3s ease;
  &:hover {
    color: var(--black8);
  }
}
@media screen and (max-width: 1100px) {
  .card-item {
    width: calc(100% / 3 - 20px);
  }
}
@media screen and (max-width: 900px) {
  .card-item {
    width: calc(100% / 2 - 20px);
  }
}
@media screen and (max-width: 700px) {
  .card-item {
    width: calc(100% - 20px);
  }
}
</style>