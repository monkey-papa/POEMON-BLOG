import http from "../utils/request";
import constant from "../utils/constant";
import apiWrapper from "./tool/apiWrapper";
import type {
  Family,
  AddFamilyParams,
  ChangeLoveStatusParams,
  PaginationParams,
  PaginationResponse,
} from "@/types";

// ==================== 表白墙相关 API ====================

/**
 * 获取管理员表白墙信息
 * @returns 表白墙信息
 */
export function getAdminFamily(): Promise<Family> {
  return apiWrapper(
    () => http.post<Family>(`${constant.baseURL}/family/getAdminFamily`),
    "获取表白墙信息失败！"
  );
}

/**
 * 获取随机表白墙列表（前台）
 * @returns 表白墙列表
 */
export function getRandomFamilyList(): Promise<Family[]> {
  return apiWrapper(
    () =>
      http.post<Family[]>(`${constant.baseURL}/family/list`, {
        isAdmin: false,
      }),
    "获取表白墙列表失败！"
  );
}

/**
 * 添加表白墙
 * 注意：后端返回 SuccessMessage，不返回表白墙对象
 * @param family - 表白墙数据
 */
export function addFamily(family: AddFamilyParams): Promise<void> {
  return apiWrapper(
    () => http.post<void>(`${constant.baseURL}/family/addFamily`, family),
    "提交失败！",
    { allowEmpty: true }
  );
}

/**
 * 获取表白墙列表（管理员）
 * @param pagination - 分页参数
 * @returns 表白墙列表
 */
export function listFamily(pagination: PaginationParams): Promise<PaginationResponse<Family>> {
  return apiWrapper(
    () =>
      http.post<PaginationResponse<Family>>(`${constant.baseURL}/family/list`, {
        ...pagination,
        isAdmin: true,
      }),
    "获取表白墙列表失败！"
  );
}

/**
 * 删除表白墙（管理员）
 * @param id - 表白墙ID
 */
export function deleteFamily(id: number): Promise<void> {
  return apiWrapper(
    () => http.post<void>(`${constant.baseURL}/family/deleteFamily`, { id }),
    "删除表白墙失败！",
    { allowEmpty: true }
  );
}

/**
 * 修改表白墙状态（管理员审核）
 * @param params - { id, flag }
 */
export function changeLoveStatus(params: ChangeLoveStatusParams): Promise<void> {
  return apiWrapper(
    () => http.post<void>(`${constant.baseURL}/family/changeLoveStatus`, params),
    "修改状态失败！",
    { allowEmpty: true }
  );
}

export default {
  getAdminFamily,
  getRandomFamilyList,
  addFamily,
  listFamily,
  deleteFamily,
  changeLoveStatus,
};
