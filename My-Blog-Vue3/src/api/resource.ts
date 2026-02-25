import http from "../utils/request";
import constant from "../utils/constant";
import apiWrapper from "./tool/apiWrapper";
import type {
  Resource,
  ResourcePath,
  GetResourceListParams,
  ChangeResourceStatusParams,
  SaveResourceParams,
  GetResourcePathListParams,
  SaveResourcePathParams,
  UpdateResourcePathParams,
  PaginationResponse,
} from "@/types";

// ==================== 资源相关 API ====================

/**
 * 获取资源列表（后台，返回全部）
 * @param pagination - 分页参数 { current, size, searchKey, type }
 * @returns 资源列表
 */
export function getResourceList(
  pagination: GetResourceListParams
): Promise<PaginationResponse<Resource>> {
  return apiWrapper(
    () =>
      http.post<PaginationResponse<Resource>>(`${constant.baseURL}/resource/listResource`, {
        ...pagination,
        isAdmin: true,
      }),
    "获取资源列表失败！"
  );
}

/**
 * 获取资源列表（前台，只返回启用的）
 * @param pagination - 分页参数 { current, size, resourceType }
 * @returns 资源列表
 */
export function getClientResourceList(
  pagination: GetResourceListParams
): Promise<PaginationResponse<Resource>> {
  return apiWrapper(
    () =>
      http.post<PaginationResponse<Resource>>(`${constant.baseURL}/resource/listResource`, {
        ...pagination,
        isAdmin: false,
      }),
    "获取资源列表失败！"
  );
}

/**
 * 修改资源状态
 * @param params - 状态参数 { id, flag }
 */
export function changeResourceStatus(params: ChangeResourceStatusParams): Promise<void> {
  return apiWrapper(
    () => http.post<void>(`${constant.baseURL}/resource/changeResourceStatus`, params),
    "修改状态失败！",
    { allowEmpty: true }
  );
}

/**
 * 删除文章图片（从七牛云和数据库中删除）
 * @param params - 删除参数 { url, id }
 */
export function deleteArticleImage(params: { url: string; id?: number }): Promise<void> {
  return apiWrapper(
    () => http.post<void>(`${constant.baseURL}/delArticleImage`, params),
    "删除图片失败！",
    { allowEmpty: true }
  );
}

/**
 * 获取资源路径列表（后台）
 * @param params - 查询参数 { current, size, type }
 * @returns 资源列表
 */
export function getResourcePathList(
  params: GetResourcePathListParams
): Promise<PaginationResponse<ResourcePath>> {
  return apiWrapper(
    () =>
      http.post<PaginationResponse<ResourcePath>>(`${constant.baseURL}/webInfo/listResourcePath`, {
        ...params,
        isAdmin: true,
      }),
    "获取资源列表失败！"
  );
}

/**
 * 获取资源路径列表（前台/娱乐列表）
 * @param params - 查询参数 { type }
 * @returns 娱乐列表（分页响应）
 */
export function getClientResourcePathList(
  params: GetResourcePathListParams
): Promise<PaginationResponse<ResourcePath>> {
  return apiWrapper(
    () =>
      http.post<PaginationResponse<ResourcePath>>(`${constant.baseURL}/webInfo/listResourcePath`, {
        ...params,
        isAdmin: false,
      }),
    "获取娱乐列表失败！"
  );
}

/**
 * 保存资源路径（新增）
 * @param resource - 资源数据
 */
export function saveResourcePath(resource: SaveResourcePathParams): Promise<void> {
  return apiWrapper(
    () => http.post<void>(`${constant.baseURL}/webInfo/saveResourcePath`, resource),
    "保存资源失败！",
    { allowEmpty: true }
  );
}

/**
 * 更新资源路径（管理员）
 * @param resource - 更新数据（包含 id）
 */
export function updateResourcePath(resource: UpdateResourcePathParams): Promise<void> {
  return apiWrapper(
    () => http.post<void>(`${constant.baseURL}/webInfo/updateResourcePath`, resource),
    "更新资源失败！",
    { allowEmpty: true }
  );
}

/**
 * 删除资源路径
 * @param id - 资源ID
 */
export function deleteResourcePath(id: number): Promise<void> {
  return apiWrapper(
    () => http.post<void>(`${constant.baseURL}/webInfo/deleteResourcePath`, { id }),
    "删除资源失败！",
    { allowEmpty: true }
  );
}

/**
 * 保存资源记录
 * @param resource - 资源数据 { type, path, size, mimeType }
 */
export function saveResource(resource: SaveResourceParams): Promise<void> {
  return apiWrapper(
    () => http.post<void>(`${constant.baseURL}/resource/saveResource`, resource),
    "保存资源失败！",
    { allowEmpty: true }
  );
}

export default {
  getResourceList,
  getClientResourceList,
  changeResourceStatus,
  deleteArticleImage,
  getResourcePathList,
  getClientResourcePathList,
  saveResourcePath,
  updateResourcePath,
  deleteResourcePath,
  saveResource,
};
