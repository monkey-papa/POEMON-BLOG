/**
 * API 相关类型定义
 */

/** API 响应基础结构 */
export interface ApiResponse<T = unknown> {
  code: number;
  msg: string;
  data: T;
  error: boolean;
  success: boolean;
}

/** 分页请求参数 */
export interface PaginationParams {
  /** 当前页（noPagination 为 true 时可选） */
  current?: number;
  /** 每页条数（noPagination 为 true 时可选） */
  size?: number;
  /** 是否不分页（获取全部数据） */
  noPagination?: boolean;
}

/** 分页响应结构 */
export interface PaginationResponse<T> {
  list: T[];
  totalCount: number;
  pageSize: number;
  totalPage: number;
  /** 当前页码 */
  page: number;
  empty: boolean;
}

/** API 包装器选项 */
export interface ApiWrapperOptions {
  /** 是否允许空响应 */
  allowEmpty?: boolean;
  /** 是否静默处理错误 */
  silent?: boolean;
}

/** 缓存配置 */
export interface CacheConfig {
  [key: string]: number;
}

/** 缓存条目 */
export interface CacheEntry<T = unknown> {
  data: T;
  timestamp: number;
  expireAt: number;
}

/** 缓存统计信息 */
export interface CacheStats {
  totalEntries: number;
  entries: CacheStatsEntry[];
}

/** 缓存统计条目 */
export interface CacheStatsEntry {
  key: string;
  expireAt: string;
  remainingTime: number;
}
