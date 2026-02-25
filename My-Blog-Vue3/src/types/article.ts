/**
 * 文章相关类型定义
 */

import type { Prettify } from "./common";
import type { LabelInfo } from "./sort";

/** 文章状态字段 */
export interface ArticleStatus {
  viewStatus: boolean;
  commentStatus: boolean;
  recommendStatus: boolean;
}

/** 文章内嵌的分类信息 */
export interface ArticleSortInfo {
  id: number;
  sortName: string;
  sortDescription: string;
  sortType: number;
  priority: number | null;
  countOfSort: number;
  labels: number;
}

/** 文章内嵌的标签信息 */
export type ArticleLabelInfo = Prettify<
  LabelInfo & {
    /** 关联的文章ID */
    articleId?: number;
  }
>;

/** 文章列表项 */
export interface ArticleListItem extends ArticleStatus {
  id: number;
  userId: number;
  sortId: number;
  labelId: number;
  /** 文章封面 */
  articleCover: string | null;
  articleTitle: string;
  articleContent: string;
  viewCount: number;
  likeCount: number;
  createTime: string;
  updateTime: string;
  /** 更新者/文章作者 */
  updateBy: string | null;
  /** 访问密码 */
  password: string | null;
  /** 作者用户名 */
  username: string;
  /** 评论数 */
  commentCount: number;
  /** 是否已有AI摘要 */
  hasSummary: boolean;
  /** 分类信息 */
  sort: ArticleSortInfo[] | null;
  /** 标签信息 */
  label: ArticleLabelInfo[] | null;
}

/** 文章详情（前台 buildArticleDetail 返回结构） */
export type ArticleDetail = Prettify<
  ArticleListItem & {
    /** 文章摘要 */
    summary: string;
    /** 文章作者名（实际等于 updateBy） */
    articleAuthor: string | null;
    /** 当前用户的点赞状态 (0=未点赞, 1=已点赞) */
    articleLikeStatus: number;
  }
>;

/** 管理员文章详情（后端精简返回，用于编辑） */
export interface AdminArticleDetail {
  id: number;
  userId: number;
  sortId: number;
  labelId: number;
  articleCover: string | null;
  articleTitle: string;
  articleContent: string;
  articleAuthor: string | null;
  viewStatus: boolean;
  password: string | null;
  recommendStatus: boolean;
  commentStatus: boolean;
}

/** 获取文章列表参数 */
export interface GetArticleListParams {
  current: number;
  size: number;
  sortId?: number | null;
  labelId?: number | null;
  searchKey?: string;
  recommendStatus?: boolean | null;
  viewAll?: boolean;
}

/** 获取文章详情参数 */
export interface GetArticleDetailParams {
  id: number;
  isAdmin?: boolean;
}

/** 文章点赞参数 */
export interface ArticleLikeParams {
  articleId: number;
}

/** 保存文章参数（新增） */
export interface SaveArticleParams {
  articleTitle: string;
  articleAuthor?: string;
  articleContent?: string;
  sortId: number;
  labelId: number;
  articleCover?: string;
  viewStatus?: boolean;
  commentStatus?: boolean;
  recommendStatus?: boolean;
  password?: string;
}

/** 更新文章参数 */
export interface UpdateArticleParams {
  id: number;
  articleTitle?: string;
  articleAuthor?: string;
  articleContent?: string;
  sortId?: number;
  labelId?: number;
  articleCover?: string;
  viewStatus?: boolean;
  commentStatus?: boolean;
  recommendStatus?: boolean;
  password?: string;
}

/** 修改文章状态参数 */
export interface ChangeArticleStatusParams {
  articleId: number;
  viewStatus?: boolean;
  commentStatus?: boolean;
  recommendStatus?: boolean;
}

/** 生成摘要参数 */
export interface GenerateSummaryParams {
  message: string;
  articleId?: number;
}

/** 生成摘要响应 */
export interface GenerateSummaryResponse {
  summary: string;
}

/** 最新文章（用于侧边栏展示） */
export interface NewArticle {
  id: number;
  articleTitle: string;
  articleCover: string | null;
  createTime: string;
}
