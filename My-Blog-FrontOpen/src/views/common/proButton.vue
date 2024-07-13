<template>
  <div class="myButton">
    <div :style="beforeColor">{{ info }}</div>
    <div :style="afterColor">{{ info }}</div>
    <div :style="afterColor">{{ info }}</div>
  </div>
</template>

<script>
export default {
  props: {
    info: {
      type: String,
      default: '确定'
    },
    before: {
      type: String
    },
    after: {
      type: String
    }
  },
  data() {
    return {
      beforeColor: { background: this.before },
      afterColor: { background: this.after }
    }
  }
}
</script>

<style lang="scss" scoped>
.myButton {
  user-select: none;
  position: relative;
  width: 66px;
  height: 33px;
  border-radius: 4px;
  color: var(--white);
  font-size: 14px;
  overflow: hidden;
  div {
    width: 66px;
    height: 33px;
    line-height: 33px;
    border-radius: 4px;
    text-align: center;
    position: absolute;
    &:nth-child(2) {
      width: 100px;
      transition: all 0.3s ease;
      transform: translateX(-120px) skewX(-30deg);
    }
    &:nth-child(3) {
      transition: all 0.3s ease;
      transform: translateX(-120px);
    }
  }
  &:hover div:nth-child(2) {
    transform: translateX(20px) skewX(-30deg);
  }
  &:hover div:nth-child(3) {
    transform: translateX(0px);
  }
}
</style>
