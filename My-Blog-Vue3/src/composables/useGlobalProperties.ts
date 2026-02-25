import { useRoute } from "vue-router";
import type { RouteLocationNormalizedLoaded } from "vue-router";
import { useStore, type MainStoreType } from "@/stores";
import { getCurrentInstance } from "vue";
import type { ElMessageBox } from "element-plus";
import type common from "@/utils/common";
import type constant from "@/utils/constant";

interface GlobalProperties {
  $route: RouteLocationNormalizedLoaded;
  $store: MainStoreType;
  $common: typeof common;
  $constant: typeof constant;
  $confirm: typeof ElMessageBox.confirm;
}

/**
 * 统一获取全局属性和常用实例的 composable
 * 避免在每个组件中重复导入和获取
 * @returns 包含 router, route, store 和所有全局属性的对象
 */
export function useGlobalProperties(): GlobalProperties {
  const route = useRoute();
  const store = useStore();
  const instance = getCurrentInstance();

  // 获取全局属性
  const globalProperties = instance?.appContext.config.globalProperties as
    | (Record<string, unknown> & Partial<GlobalProperties>)
    | undefined;

  // 返回全局属性
  return {
    $route: route,
    $store: store,
    $common: globalProperties?.$common as typeof common,
    $constant: globalProperties?.$constant as typeof constant,
    $confirm: globalProperties?.$confirm as typeof ElMessageBox.confirm,
  };
}
