import http from "../utils/request";
import constant from "../utils/constant";
import apiWrapper from "./tool/apiWrapper";
import type {
  CommentListItem,
  ClientCommentListResponse,
  GetCommentListParams,
  BossAddCommentParams,
  PaginationResponse,
} from "@/types";

// ==================== 评论相关 API ====================

/**
 * 获取评论列表（前台）
 * 注意：此接口返回自定义结构 { total, parenttotal, data }，非标准分页
 * @param params - 分页参数 { current, size, source, commentType, floorCommentId, csize }
 * @returns 评论列表（自定义结构）
 */
export function getCommentList(params: GetCommentListParams): Promise<ClientCommentListResponse> {
  return apiWrapper(
    () => http.post<ClientCommentListResponse>(`${constant.baseURL}/comment/listComment`, params),
    "获取评论列表失败！"
  );
}

/**
 * 获取管理员评论列表（后台，标准分页）
 * @param pagination - 分页参数
 * @returns 评论列表
 */
export function getBossCommentList(
  pagination: GetCommentListParams
): Promise<PaginationResponse<CommentListItem>> {
  return apiWrapper(
    () =>
      http.post<PaginationResponse<CommentListItem>>(
        `${constant.baseURL}/admin/comment/boss/list`,
        pagination
      ),
    "获取评论列表失败！"
  );
}

/**
 * 添加评论（需登录）
 * 注意：后端返回 SuccessMessage，不返回评论对象
 * @param comment - 评论数据
 */
export function bossAddComment(comment: BossAddCommentParams): Promise<void> {
  return apiWrapper(
    () => http.post<void>(`${constant.baseURL}/admin/comment/boss/addComment`, comment),
    "添加评论失败！",
    { allowEmpty: true }
  );
}

/**
 * 删除评论（管理员）
 * @param id - 评论ID
 */
export function bossDeleteComment(id: number): Promise<void> {
  return apiWrapper(
    () => http.post<void>(`${constant.baseURL}/admin/comment/boss/deleteComment`, { id }),
    "删除评论失败！",
    { allowEmpty: true }
  );
}

export default {
  getCommentList,
  getBossCommentList,
  bossAddComment,
  bossDeleteComment,
};
