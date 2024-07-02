<template>
  <div>
    <!-- 首页图片 -->
    <div style="animation: header-effect 2s" :style="{ background: `${$store.state.changeBg}` }" class="background-image background-image-changeBg"></div>
    <!-- 标签区域 -->
    <div class="main">
      <div class="layout hide-aside">
        <div id="tag" style="max-width: 100%; min-height: 100vh" ref="geo"></div>
      </div>
    </div>
    <!-- 页脚 -->
    <div>
      <myFooter></myFooter>
    </div>
  </div>
</template>
<script>
import * as echarts from 'echarts'
const myFooter = () => import('./common/myFooter')
export default {
  mounted() {
    this.map()
  },
  components: {
    myFooter
  },
  methods: {
    async map() {
      const resp = await this.getMap()
      const users = await this.getUsers()
      const myChart = echarts.init(this.$refs.geo)
      myChart.showLoading() //初始化，获得echarts实例
      //注册地图
      echarts.registerMap('China', resp)
      myChart.setOption({
        title: {
          text: '注册用户分布图',
          textStyle: {
            color: '#ec695c'
          }
        },
        //配置了该项，鼠标指上去有提示
        tooltip: {
          formatter: '{b}注册用户{c}人',
          textStyle: {
            color: '#ec695c'
          },
          backgroundColor: '#74bdf0',
          borderColor: '#74bdf0'
        },
        //地图条
        visualMap: {
          //可视地图，一般用户设置不同颜色来展示数据差异
          left: 'left', //可视地图显示的位置
          top: 'center', //可视地图显示的位置
          min: 0, //区间的最小值
          max: 10000, //区间数据的最大值
          text: ['高', '低'],
          calculable: true //是否允许控制区间
        },
        series: [
          {
            type: 'map', //图标类型：地图
            map: 'China', //使用注册的地图
            center: [100, 40], //当前视角的中心点，用经纬度表示
            roam: true, //是否开启鼠标缩放和平移
            zoom: 2.3, //当前视角的缩放比例
            scaleLimit: {
              min: 1, //最小缩放2.3倍
              max: 5 //最大放大3倍
            },
            data: users,
            label: {
              //图形上的文本标签，可用于说明图形的一些数据信息，比如值，名称等
              normal: {
                show: true, //是否在普通状态下显示标签。
                textStyle: {
                  color: '#ec695c', //文字颜色
                  fontStyle: 'normal', //italic斜体  oblique倾斜
                  fontWeight: 'normal', //文字粗细bold   bolder   lighter  100 | 200 | 300 | 400...
                  fontFamily: 'sans-serif', //字体系列
                  fontSize: 12 //字体大小
                }
              }
            },
            itemStyle: {
              //地图区域的多边形 图形样式
              normal: {
                color: 'red', //颜色
                borderColor: '#000', //边框颜色
                borderWidth: 0, //柱条的描边宽度，默认不描边。
                borderType: 'solid', //柱条的描边类型，默认为实线，支持 'dashed', 'dotted'。
                barBorderRadius: 0, //柱形边框圆角半径，单位px，支持传入数组分别指定柱形4个圆角半径。
                shadowBlur: 5, //图形阴影的模糊大小。
                shadowColor: '#000', //阴影颜色
                shadowOffsetX: 5, //阴影水平方向上的偏移距离。
                shadowOffsetY: 5, //阴影垂直方向上的偏移距离。
                opacity: 1 //图形透明度。支持从 0 到 1 的数字，为 0 时不绘制该图形。
              }
            }
          }
        ]
      })
      myChart.hideLoading()
      var myChartDom = echarts.init(document.getElementById('tag'))
      window.addEventListener('resize', function () {
        myChartDom.resize()
      })
    },
    getMap() {
      return new Promise(resolve => {
        this.$http
          .get(this.$constant.baseURL + '/map')
          .then(res => {
            if (!this.$common.isEmpty(res)) {
              resolve(res)
            }
          })
          .catch(error => {
            this.$message({
              message: error.message,
              type: 'error'
            })
          })
      })
    },
    getUsers() {
      return new Promise(resolve => {
        this.$http
          .get(this.$constant.baseURL + '/user')
          .then(res => {
            if (!this.$common.isEmpty(res)) {
              resolve(res)
            }
          })
          .catch(error => {
            this.$message({
              message: error.message,
              type: 'error'
            })
          })
      })
    }
  }
}
</script>
<style scoped>
.main {
  margin: 0 auto;
  padding: 60px 40px 10px;
  min-height: 100vh;
}
.layout {
  display: flex;
}
.layout .hide-aside {
  max-width: 1408px;
}
div #tag {
  height: 100%;
  width: 100%;
  border: 2px dashed rgba(0, 255, 255, 0.6);
  font-weight: 700;
  background: var(--background);
}
.layout > div {
  border-radius: 18px;
  padding: 10px 5px;
}
@media screen and (max-width: 1286px) {
  .main {
    padding: 60px 3px;
  }
}
@media screen and (max-width: 523px) {
  #tag {
    padding: 10px 8px;
  }
}
</style>
