/* 先引入打包分析插件 */
const BundleAnalyzerPlugin =
  require("webpack-bundle-analyzer").BundleAnalyzerPlugin;
const cdn = {
  js: [
    "https://cdnjs.cloudflare.com/ajax/libs/vue/2.6.11/vue.min.js",
    // "https://lf26-cdn-tos.bytecdntp.com/cdn/expire-1-M/vuex/3.6.2/vuex.min.js",
    "https://cdnjs.cloudflare.com/ajax/libs/vue-router/3.2.0/vue-router.min.js",
    "https://cdnjs.cloudflare.com/ajax/libs/axios/0.21.1/axios.min.js",
  ],
};
module.exports = {
  devServer: {
    port: 81,
  },
  configureWebpack: {
    plugins: [
      // 依赖大小分析工具
      new BundleAnalyzerPlugin(),
    ],
    //这就是在告诉Webpack：请不要将这个模块注入编译后的JS文件里，对于我源代码里出现的任何import/require这个模块的语句，请将它保留。
    externals: {
      vue: "Vue",
      "vue-router": "VueRouter",
      // vuex: "Vuex",
      axios: "axios",
      echarts: "echarts", // 配置使用CDN
    },
  },
  lintOnSave: false,
  productionSourceMap: false, // 打包成 web页面 使用，一般建议 false，不然会生成很多map文件
  chainWebpack(config) {
    // 引入CDN
    // webpack需要排除的依赖名称和挂载在window上的对象属性名称
    config.set("externals", cdn.externals);
    config.plugin("html").tap((args) => {
      args[0].cdn = cdn;
      return args;
    });
  },
};
