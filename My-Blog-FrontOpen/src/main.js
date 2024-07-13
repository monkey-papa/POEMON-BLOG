import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import ElementUI from "element-ui";
import http from "./utils/request";
import common from "./utils/common";
import constant from "./utils/constant";
import mavonEditor from "mavon-editor";
NProgress.configure({ showSpinner: false });
NProgress.configure({ minimum: 0 });
NProgress.configure({ ease: "linear", speed: 100 });
import "./utils/title";
//引入css
import "./assets/css/index.css";
import "./assets/css/tocbot.css";
import "./assets/css/color.css";
import "./assets/css/markdown-highlight.css";
import "mavon-editor/dist/css/index.css";
import "./assets/css/icon.css";
//图标引入
import "./assets/css/font-awesome.min.css";
//图片跳动
import "./assets/css/animation.css";
//树洞
import { vueBaberrage } from "vue-baberrage";
//樱花
import "./utils/sakura.js";
Vue.use(ElementUI);
Vue.use(vueBaberrage);
Vue.use(mavonEditor);
Vue.prototype.$http = http;
Vue.prototype.$common = common;
Vue.prototype.$constant = constant;
Vue.config.productionTip = false;
// 修改 el-dialog 默认点击遮照不关闭
ElementUI.Dialog.props.closeOnClickModal.default = false;
new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount("#app");
