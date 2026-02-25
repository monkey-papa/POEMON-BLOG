/**
 * Store 相关类型定义
 */

import type { CurrentUser, CurrentAdmin, PublicUserListItem } from "./user";
import type { SortInfo } from "./sort";
import type { WebInfo, LocationInfo, PageView } from "./webInfo";
import type { NewArticle } from "./article";

/** 工具栏状态 */
export interface ToolbarState {
  visible: boolean;
  enter: boolean;
}

/** Store State 类型 */
export interface MainState {
  toolbar: ToolbarState;
  sortInfo: SortInfo[];
  /** null 表示未登录 */
  currentUser: CurrentUser | null;
  /** null 表示未登录 */
  currentAdmin: CurrentAdmin | null;
  webInfo: WebInfo;
  locationInfo: LocationInfo;
  changeBg: string;
  isShowLoading: boolean;
  top: number;
  pageView: PageView;
  newArticles: NewArticle[];
  articleTotal: number;
  userList: PublicUserListItem[];
}

/** Store Getters 类型 */
export interface MainGetters {
  permissions: string[];
  totalArticles: number;
  navigationBar: SortInfo[];
  labelInfo: import("./sort").LabelInfo[];
}

/** 持久化配置 */
export interface PersistConfig {
  key: string;
  storage: Storage;
  paths: (keyof MainState)[];
}
