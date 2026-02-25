import http from "../utils/request";
import constant from "../utils/constant";
import apiWrapper from "./tool/apiWrapper";
import type {
  ProhibitedWord,
  GetProhibitedWordsListParams,
  AddProhibitedWordParams,
  UpdateProhibitedWordParams,
  PaginationResponse,
} from "@/types";

// ==================== 违禁词相关 API ====================

/**
 * 获取违禁词列表
 * @param pagination - 分页参数 { current, size, searchKey }
 * @returns 违禁词列表
 */
export function getProhibitedWordsList(
  pagination: GetProhibitedWordsListParams
): Promise<PaginationResponse<ProhibitedWord>> {
  return apiWrapper(
    () =>
      http.post<PaginationResponse<ProhibitedWord>>(
        `${constant.baseURL}/prohibitedWords/list`,
        pagination
      ),
    "获取违禁词列表失败！"
  );
}

/**
 * 添加违禁词
 * 注意：后端返回 SuccessMessage，不返回违禁词对象
 * @param word - 违禁词数据 { message, username, avatar }
 */
export function addProhibitedWord(word: AddProhibitedWordParams): Promise<void> {
  return apiWrapper(
    () => http.post<void>(`${constant.baseURL}/prohibitedWords/add`, word),
    "添加违禁词失败！",
    { allowEmpty: true }
  );
}

/**
 * 更新违禁词
 * @param word - 违禁词数据 { id, message }
 */
export function updateProhibitedWord(word: UpdateProhibitedWordParams): Promise<void> {
  return apiWrapper(
    () => http.post<void>(`${constant.baseURL}/prohibitedWords/update`, word),
    "更新违禁词失败！",
    { allowEmpty: true }
  );
}

/**
 * 删除违禁词
 * @param id - 违禁词ID
 */
export function deleteProhibitedWord(id: number): Promise<void> {
  return apiWrapper(
    () => http.post<void>(`${constant.baseURL}/prohibitedWords/delete`, { id }),
    "删除违禁词失败！",
    { allowEmpty: true }
  );
}

export default {
  getProhibitedWordsList,
  addProhibitedWord,
  updateProhibitedWord,
  deleteProhibitedWord,
};
