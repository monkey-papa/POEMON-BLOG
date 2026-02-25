/**
 * 分类标签相关类型定义
 */

import type { Prettify } from "./common";

/** 分类类型枚举 */
export enum SortType {
  /** 导航栏分类 */
  Navigation = 0,
  /** 普通分类 */
  Normal = 1,
}

/** 标签信息 */
export interface LabelInfo {
  id: number;
  labelName: string;
  labelDescription: string;
  sortId: number;
  /** 该标签下的文章数量 */
  countOfLabel: number;
}

/** 分类信息 */
export interface SortInfo {
  id: number;
  sortName: string;
  sortDescription: string;
  sortType: SortType;
  /** 优先级 */
  priority: number | null;
  /** 该分类下的文章数量 */
  countOfSort: number;
  /** 该分类下的标签数量 */
  lengthOfLabel: number;
  /** 下属标签列表 */
  labels: LabelInfo[];
}

/** 保存分类参数 */
export interface SaveSortParams {
  sortName: string;
  sortDescription: string;
  sortType: SortType;
  priority: number;
}

/** 更新分类参数 */
export type UpdateSortParams = Prettify<SaveSortParams & { id: number }>;

/** 保存标签参数 */
export interface SaveLabelParams {
  labelName: string;
  labelDescription: string;
  sortId: number;
}

/** 更新标签参数 */
export type UpdateLabelParams = Prettify<SaveLabelParams & { id: number }>;

/** 分类统计项 */
export interface ClassifyItem {
  classify: string;
  count: number;
}

/** ListSortAndLabel 响应结构 */
export interface SortAndLabelResponse {
  sorts: SortInfo[];
  labels: LabelInfo[];
}
