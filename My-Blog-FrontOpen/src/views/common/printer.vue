<template>
  <div>
    <slot name="paper" :content="content"></slot>
  </div>
</template>
<script>
export default {
  props: {
    //内容
    printerInfo: {
      type: String,
      default: ""
    },
    //速度
    duration: {
      type: Number,
      default: 100
    },
    //延迟
    delay: {
      type: Number,
      default: 3000
    },
    working: {
      type: Boolean,
      default: true
    },
    once: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      content: "",
      cursor: 0,
      timer: null,
      timeout: null,
      print: true
    };
  },
  created() {
    if (this.working) {
      //this.work：引用函数
      //this.work()：执行函数
      this.start(this.work);
    } else {
      this.content = this.printerInfo
    }
  },
  watch: {
    working(newVal) {
      this.toBegin();
    },
    printerInfo(newVal) {
      this.toBegin();
    },
    cursor(cursor) {
      //slice(start,end)：不包含end
      this.content = this.printerInfo.slice(0, cursor)
    }
  },
  beforeDestroy() {
    clearInterval(this.timer);
  },
  methods: {
    /**
     * 定时
     */
    start(work) {
      //延迟
      this.timeout = setTimeout(() => {
        //速度
        this.timer = setInterval(() => {
          work();
          if (this.cursor === 0 || (this.cursor === this.printerInfo.length && !this.once)) {
            //此处为了延迟
            clearInterval(this.timer);
            this.start(this.work);
          } else if (this.cursor === this.printerInfo.length && this.once) {
            clearInterval(this.timer);
          }
        }, this.duration);
      }, this.delay)
    },
    /**
     * 逻辑
     */
    work() {
      let cursor = this.cursor;
      cursor += this.print ? 1 : -1;
      if (this.print) {
        if (cursor === this.printerInfo.length + 1) {
          cursor -= 2;
          this.print = !this.print;
        }
      } else {
        if (cursor === -1) {
          cursor += 2;
          this.print = !this.print;
        }
      }
      this.cursor = cursor;
    },
    toBegin() {
      this.cursor = 0;
      if (this.timeout !== null) {
        clearTimeout(this.timeout);
        if (this.timer !== null) {
          clearInterval(this.timer);
        }
      }
      if (this.working) {
        this.start(this.work);
      } else {
        this.content = this.printerInfo;
      }
    }
  }
};
</script>
