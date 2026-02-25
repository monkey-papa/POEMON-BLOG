import http from "../utils/request";
import constant from "../utils/constant";
import apiWrapper from "./tool/apiWrapper";
import { cachedRequest } from "../utils/requestCache";
import type { CollectListResponse } from "@/types";

// ==================== 收藏夹相关 API ====================

/**
 * 获取收藏夹列表（带缓存）
 * 后端返回按 classify 分组的 Map: { "分类名": [items...] }
 * @returns 收藏夹分组数据
 */
export function getCollectList(): Promise<CollectListResponse> {
  const url = `${constant.baseURL}/webInfo/listCollect`;
  const params = {};

  return apiWrapper(
    () =>
      cachedRequest<CollectListResponse>(url, params, () =>
        http.post<CollectListResponse>(url, params)
      ),
    "获取收藏夹失败！"
  );
}

export default {
  getCollectList,
};
