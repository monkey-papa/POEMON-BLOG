import http from "../utils/request";
import constant from "../utils/constant";
import apiWrapper from "./tool/apiWrapper";
import type {
  WeiYan,
  GetWeiYanListParams,
  SaveWeiYanParams,
  GetNewsListParams,
  TreeHole,
  SaveTreeHoleParams,
  PaginationResponse,
  PaginationParams,
} from "@/types";

// ==================== 微言/树洞相关 API ====================

/**
 * 获取微言列表
 * @param pagination - 分页参数
 * @returns 微言列表（标准分页）
 */
export function getWeiYanList(
  pagination: GetWeiYanListParams
): Promise<PaginationResponse<WeiYan>> {
  return apiWrapper(
    () =>
      http.post<PaginationResponse<WeiYan>>(`${constant.baseURL}/weiYan/listWeiYan`, pagination),
    "获取微言列表失败！"
  );
}

/**
 * 修改微言公开状态（管理员）
 * @param id - 微言ID
 * @param isPublic - 是否公开
 */
export function changeWeiYanIsPublic(id: number, isPublic: boolean): Promise<void> {
  return apiWrapper(
    () =>
      http.post<void>(`${constant.baseURL}/weiYan/changePublic`, {
        id,
        isPublic,
      }),
    "修改微言状态失败！",
    { allowEmpty: true }
  );
}

/**
 * 保存进展（关联文章的微言）
 * 注意：后端返回 SuccessMessage，不返回微言对象
 * @param weiYan - 进展数据 { content, source }
 */
export function saveNews(weiYan: SaveWeiYanParams & { source: number }): Promise<void> {
  return apiWrapper(
    () =>
      http.post<void>(`${constant.baseURL}/weiYan/save`, {
        content: weiYan.content,
        source: weiYan.source,
        isPublic: weiYan.isPublic !== undefined ? weiYan.isPublic : true,
      }),
    "添加进展失败！",
    { allowEmpty: true }
  );
}

/**
 * 删除微言（作者或管理员）
 * @param id - 微言ID
 */
export function deleteWeiYan(id: number): Promise<void> {
  return apiWrapper(
    () => http.post<void>(`${constant.baseURL}/weiYan/deleteWeiYan`, { id }),
    "删除失败！",
    { allowEmpty: true }
  );
}

/**
 * 保存微言（不关联文章）
 * 注意：后端返回 SuccessMessage，不返回微言对象
 * @param weiYan - 微言数据 { content, isPublic }
 */
export function saveWeiYan(weiYan: SaveWeiYanParams): Promise<void> {
  return apiWrapper(
    () =>
      http.post<void>(`${constant.baseURL}/weiYan/save`, {
        content: weiYan.content,
        isPublic: weiYan.isPublic !== undefined ? weiYan.isPublic : true,
      }),
    "保存微言失败！",
    { allowEmpty: true }
  );
}

/**
 * 获取进展列表（关联文章的微言）
 * @param params - 查询参数 { source, current, size, all, noPagination }
 * @returns 进展列表（标准分页）
 */
export function getNewsList(params: GetNewsListParams): Promise<PaginationResponse<WeiYan>> {
  return apiWrapper(
    () => http.post<PaginationResponse<WeiYan>>(`${constant.baseURL}/weiYan/listNews`, params),
    "获取进展列表失败！"
  );
}

/**
 * 保存树洞留言
 * 后端返回创建的树洞对象
 * @param treeHole - 树洞数据 { message, avatar }
 * @returns 创建的树洞
 */
export function saveTreeHole(treeHole: SaveTreeHoleParams): Promise<TreeHole> {
  return apiWrapper(
    () => http.post<TreeHole>(`${constant.baseURL}/webInfo/saveTreeHole`, treeHole),
    "保存留言失败！"
  );
}

/**
 * 删除树洞留言（管理员）
 * @param id - 树洞ID
 */
export function deleteTreeHole(id: number): Promise<void> {
  return apiWrapper(
    () => http.post<void>(`${constant.baseURL}/webInfo/deleteTreeHole`, { id }),
    "删除留言失败！",
    { allowEmpty: true }
  );
}

/**
 * 获取管理员树洞列表
 * @param pagination - 分页参数
 * @returns 树洞列表（标准分页）
 */
export function getBossTreeHoleList(
  pagination: PaginationParams
): Promise<PaginationResponse<TreeHole>> {
  return apiWrapper(
    () =>
      http.post<PaginationResponse<TreeHole>>(
        `${constant.baseURL}/admin/treeHole/boss/list`,
        pagination
      ),
    "获取树洞列表失败！"
  );
}

export default {
  getWeiYanList,
  changeWeiYanIsPublic,
  saveNews,
  deleteWeiYan,
  saveWeiYan,
  getNewsList,
  saveTreeHole,
  deleteTreeHole,
  getBossTreeHoleList,
};
