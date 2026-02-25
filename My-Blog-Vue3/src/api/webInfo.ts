import http from "../utils/request";
import constant from "../utils/constant";
import apiWrapper from "./tool/apiWrapper";
import { cachedRequest, clearCache } from "../utils/requestCache";
import type { WebInfo, SortInfo, SortAndLabelResponse, ClassifyItem } from "@/types";

// ==================== 网站信息相关 API ====================

/**
 * 获取网站信息（前台，带缓存）
 * @returns 网站信息
 */
export function getWebInfo(): Promise<WebInfo> {
  const url = `${constant.baseURL}/webInfo/getWebInfo`;
  const params = { isAdmin: false };

  return apiWrapper(
    () => cachedRequest<WebInfo>(url, params, () => http.post<WebInfo>(url, params)),
    "获取网站信息失败！"
  );
}

/**
 * 获取分类列表（带缓存）
 * @param params - 查询参数
 * @returns 分类列表
 */
export function getClassifyList(params: Record<string, unknown> = {}): Promise<ClassifyItem[]> {
  const url = `${constant.baseURL}/webInfo/getClassifyList`;

  return apiWrapper(
    () => cachedRequest<ClassifyItem[]>(url, params, () => http.post<ClassifyItem[]>(url, params)),
    "获取分类列表失败！"
  );
}

/**
 * 获取排序信息（带缓存）
 * @returns 排序信息
 */
export function getSortInfo(): Promise<SortInfo[]> {
  const url = `${constant.baseURL}/webInfo/getSortInfo`;
  const params = {};

  return apiWrapper(
    () => cachedRequest<SortInfo[]>(url, params, () => http.post<SortInfo[]>(url, params)),
    "获取分类信息失败！"
  );
}

/**
 * 获取分类和标签列表（带缓存）
 * 后端返回 { sorts: [...], labels: [...] }
 * @returns 分类和标签数据
 */
export function getSortAndLabelList(): Promise<SortAndLabelResponse> {
  const url = `${constant.baseURL}/webInfo/listSortAndLabel`;
  const params = {};

  return apiWrapper(
    () =>
      cachedRequest<SortAndLabelResponse>(url, params, () =>
        http.post<SortAndLabelResponse>(url, params)
      ),
    "获取分类标签失败！"
  );
}

/**
 * 获取管理员网站信息（后台）
 * @returns 网站信息
 */
export function getAdminWebInfo(): Promise<WebInfo> {
  return apiWrapper(
    () =>
      http.post<WebInfo>(`${constant.baseURL}/webInfo/getWebInfo`, {
        isAdmin: true,
      }),
    "获取网站信息失败！"
  );
}

/**
 * 更新管理员网站信息
 * @param params - 网站信息
 */
export async function updateAdminWebInfo(params: Partial<WebInfo>): Promise<void> {
  await apiWrapper(
    () => http.post<void>(`${constant.baseURL}/admin/webInfo/updateAdminWebInfo`, params),
    "更新网站信息失败！",
    { allowEmpty: true }
  );

  // 更新成功后清除网站信息缓存
  clearCache("/webInfo/getWebInfo");
}

export default {
  getWebInfo,
  getClassifyList,
  getSortInfo,
  getSortAndLabelList,
  getAdminWebInfo,
  updateAdminWebInfo,
};
