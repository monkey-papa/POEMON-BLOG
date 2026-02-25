/**
 * 微言相关类型定义
 */

import type { Prettify } from "./common";

/** 微言信息 */
export interface WeiYan {
  id: number;
  userId: number;
  content: string;
  /** 关联文章ID */
  source: number | null;
  /** 微言类型 */
  type: string;
  isPublic: boolean;
  likeCount: number;
  createTime: string;
  username: string;
  /** 头像 */
  avatar: string | null;
}

/** 获取微言列表参数 */
export interface GetWeiYanListParams {
  current: number;
  size: number;
}

/** 保存微言参数 */
export interface SaveWeiYanParams {
  content: string;
  isPublic?: boolean;
  source?: number;
}

/** 进展（关联文章的微言，source 必须存在） */
export type News = Prettify<WeiYan & { source: number }>;

/** 获取进展列表参数 */
export interface GetNewsListParams {
  current?: number;
  size?: number;
  source?: number;
  all?: boolean;
  noPagination?: boolean;
}
