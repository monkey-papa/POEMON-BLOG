import hasPermi from "./permission/hasPermi";

const install = function (Vue) {
  Vue.directive("hasPermi", hasPermi);
};

if (window.Vue) {
  window.hasPermi = hasPermi;
  Vue.use(install);
}

export default install;
