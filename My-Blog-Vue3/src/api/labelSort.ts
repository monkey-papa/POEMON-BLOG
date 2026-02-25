import http from "../utils/request";
import constant from "../utils/constant";
import apiWrapper from "./tool/apiWrapper";
import { clearCache } from "../utils/requestCache";
import type { SaveSortParams, UpdateSortParams, SaveLabelParams, UpdateLabelParams } from "@/types";

// ==================== 分类标签相关 API ====================

/**
 * 清除分类标签相关缓存
 */
function clearSortLabelCache(): void {
  clearCache("/webInfo/getSortInfo");
  clearCache("/webInfo/listSortAndLabel");
  clearCache("/webInfo/getClassifyList");
}

/**
 * 删除分类（管理员）
 * @param id - 分类ID
 */
export async function deleteSort(id: number): Promise<void> {
  await apiWrapper(
    () => http.post<void>(`${constant.baseURL}/webInfo/deleteSort`, { id }),
    "删除分类失败！",
    { allowEmpty: true }
  );
  clearSortLabelCache();
}

/**
 * 删除标签（管理员）
 * @param id - 标签ID
 */
export async function deleteLabel(id: number): Promise<void> {
  await apiWrapper(
    () => http.post<void>(`${constant.baseURL}/webInfo/deleteLabel`, { id }),
    "删除标签失败！",
    { allowEmpty: true }
  );
  clearSortLabelCache();
}

/**
 * 保存分类（新增）
 * 注意：后端返回 SuccessMessage，不返回分类对象
 * @param sort - 分类数据
 */
export async function saveSort(sort: SaveSortParams): Promise<void> {
  await apiWrapper(
    () => http.post<void>(`${constant.baseURL}/webInfo/saveOrUpdateSort`, sort),
    "保存分类失败！",
    { allowEmpty: true }
  );
  clearSortLabelCache();
}

/**
 * 更新分类
 * @param sort - 分类数据（包含 id）
 */
export async function updateSort(sort: UpdateSortParams): Promise<void> {
  await apiWrapper(
    () => http.post<void>(`${constant.baseURL}/webInfo/saveOrUpdateSort`, sort),
    "更新分类失败！",
    { allowEmpty: true }
  );
  clearSortLabelCache();
}

/**
 * 保存标签（新增）
 * @param label - 标签数据
 */
export async function saveLabel(label: SaveLabelParams): Promise<void> {
  await apiWrapper(
    () => http.post<void>(`${constant.baseURL}/webInfo/saveOrUpdateLabel`, label),
    "保存标签失败！",
    { allowEmpty: true }
  );
  clearSortLabelCache();
}

/**
 * 更新标签
 * @param label - 标签数据（包含 id）
 */
export async function updateLabel(label: UpdateLabelParams): Promise<void> {
  await apiWrapper(
    () => http.post<void>(`${constant.baseURL}/webInfo/saveOrUpdateLabel`, label),
    "更新标签失败！",
    { allowEmpty: true }
  );
  clearSortLabelCache();
}

export default {
  deleteSort,
  deleteLabel,
  saveSort,
  updateSort,
  saveLabel,
  updateLabel,
};
