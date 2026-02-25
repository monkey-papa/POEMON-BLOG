/**
 * 违禁词相关类型定义
 */

/** 违禁词信息 */
export interface ProhibitedWord {
  id: number;
  message: string;
  username: string;
  /** 头像 */
  avatar: string | null;
}

/** 获取违禁词列表参数 */
export interface GetProhibitedWordsListParams {
  current: number;
  size: number;
  searchKey?: string;
}

/** 添加违禁词参数 */
export interface AddProhibitedWordParams {
  message: string;
  username?: string;
  avatar?: string | null;
}

/** 更新违禁词参数 */
export interface UpdateProhibitedWordParams {
  id: number;
  message: string;
}
