/**
 * 资源相关类型定义
 */

/** 资源信息 */
export interface Resource {
  id: number;
  userId: number;
  /** 上传者用户名 */
  username: string;
  /** 资源类型：userAvatar, articleCover, articlePicture 等 */
  type: string;
  path: string;
  /** 文件大小 */
  size: number | null;
  /** MIME 类型 */
  mimeType: string | null;
  status: boolean;
  createTime: string;
}

/** 获取资源列表参数 */
export interface GetResourceListParams {
  current?: number;
  size?: number;
  /** 资源类型筛选（按 mime_type） */
  resourceType?: string;
  isAdmin?: boolean;
  noPagination?: boolean;
}

/** 修改资源状态参数 */
export interface ChangeResourceStatusParams {
  id: number;
  flag: boolean;
}

/** 保存资源参数 */
export interface SaveResourceParams {
  type: string;
  path: string;
  size: number;
  mimeType: string;
}

/** 资源路径信息 */
export interface ResourcePath {
  id: number;
  title: string;
  /** 分类 */
  classify: string | null;
  /** 封面 */
  cover: string | null;
  /** 链接 */
  url: string | null;
  /** 简介 */
  introduction: string | null;
  /** 资源类型：friendUrl, movie, music, wallpaper, favorites 等 */
  type: string;
  /** 友链头像 */
  friendAvatar: string | null;
  status: boolean;
  createTime: string;
}

/** 获取资源路径列表参数 */
export interface GetResourcePathListParams {
  current?: number;
  size?: number;
  /** 资源类型筛选 */
  resourceType?: string;
  /** 分类筛选 */
  classify?: string;
  isAdmin?: boolean;
  noPagination?: boolean;
}

/** 保存资源路径参数 */
export interface SaveResourcePathParams {
  title: string;
  classify?: string;
  cover?: string;
  url?: string;
  introduction?: string;
  type: string;
  friendAvatar?: string;
}

/** 更新资源路径参数（管理员） */
export interface UpdateResourcePathParams {
  id: number;
  title?: string;
  classify?: string | null;
  cover?: string | null;
  url?: string | null;
  introduction?: string | null;
  type?: string;
  remark?: string | null;
  friendAvatar?: string | null;
  status?: boolean;
}
