/**
 * 评论相关类型定义
 */

/** 子评论分页结构（后端返回格式） */
export interface ChildCommentsPage {
  current: number;
  records: CommentListItem[];
  total: number;
}

/** 评论列表项 */
export interface CommentListItem {
  id: number;
  source: number;
  type: string;
  userId: number;
  commentContent: string;
  /** 父评论ID（顶级评论为 null） */
  parentCommentId: number | null;
  parentUserId: number | null;
  floorCommentId: number | null;
  likeCount: number;
  /** 附加信息 */
  commentInfo: string | null;
  createTime: string;
  /** 评论者用户名 */
  username: string;
  /** 评论者头像 */
  avatar: string;
  /** 用户类型：0-Boss，1-管理员，2-普通用户，3-访客 */
  userType: number;
  /** 父评论的用户名（仅回复时存在） */
  parentUsername?: string;
  /** 子评论分页列表（仅顶级评论且未指定 floorCommentId 时返回） */
  childComments?: ChildCommentsPage[];
}

/**
 * 前台评论列表响应（ClientCommentList 返回的特殊结构，非标准分页）
 * 注意：此接口不返回标准 PaginationResponse，而是自定义结构
 */
export interface ClientCommentListResponse {
  /** 总评论数（包含子评论） */
  total: number;
  /** 子评论总数 = total - parentTotal */
  parenttotal: number;
  /** 评论列表 */
  data: CommentListItem[];
}

/** 获取评论列表参数 */
export interface GetCommentListParams {
  current: number;
  size: number;
  source?: number;
  commentType?: string;
  /** 楼层评论ID（展开子评论时使用） */
  floorCommentId?: number | null;
  /** 子评论每页条数 */
  csize?: number;
}

/** 添加评论参数 */
export interface AddCommentParams {
  source: number;
  type?: string;
  commentContent: string;
  parentCommentId?: number;
  parentUserId?: number;
  floorCommentId?: number;
  /** 附加信息 */
  commentInfo?: string;
}

/** Boss添加评论参数（与 AddCommentParams 相同） */
export type BossAddCommentParams = AddCommentParams;
