<template>
  <div class="myFooter-wrap" v-show="showFooter">
    <div class="myFooter">
      <div class="footer-title">{{ $store.state.webInfo.footer }}</div>
      <div class="footer-title">
        自 2023-8-13 建站以来，小站已苟活 {{ countdownChange }} ( =•ω•= )m
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    showFooter: {
      type: Boolean,
      default: true,
    },
  },
  data() {
    return {
      countdownTime: "2023-08-13 00:00:00",
      countdownChange: "",
    };
  },
  mounted() {
    setInterval(() => {
      this.countdown();
    }, 1000);
  },
  methods: {
    countdown() {
      if (this.$common.isEmpty(this.countdownTime)) {
        return;
      }
      let countdown = this.$common.countdown(this.countdownTime);
      this.countdownChange =
        "-" +
        countdown.d +
        "天" +
        "-" +
        countdown.h +
        "时" +
        "-" +
        countdown.m +
        "分" +
        "-" +
        countdown.s +
        "秒";
    },
  },
};
</script>

<style scoped>
.myFooter-wrap {
  user-select: none;
  animation: hideToShow 2s;
}

.myFooter {
  border-radius: 1.5rem 1.5rem 0 0;
  background: var(--gradientBG);
  text-align: center;
  color: white;
  background-size: 300% 300%;
  animation: gradientBG 10s ease infinite;
}

.footer-title {
  padding-top: 5px;
  padding-bottom: 5px;
  font-size: 16px;
}
</style>
