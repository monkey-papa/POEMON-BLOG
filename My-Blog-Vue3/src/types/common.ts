/**
 * 通用类型定义
 */

/**
 * 工具类型：展开接口继承/交叉类型，使鼠标悬停时显示完整类型
 * @description 将 interface extends 或 Type & Type 展开为扁平的属性列表
 * @example
 * type ExpandedUser = Prettify<UserBase & { extra: string }>
 * // 悬停会显示所有属性，而不是 UserBase & { extra: string }
 */
export type Prettify<T> = {
  [K in keyof T]: T[K];
} & {};

/** 时间差结果 */
export interface TimeDiffResult {
  diffYear: number;
  diffMonth: number;
  diffDay: number;
  diffHour: number;
  diffMinute: number;
  diffSecond: number;
}

/** 倒计时结果 */
export interface CountdownResult {
  d: number;
  h: number;
  m: number;
  s: number;
}

/** 位置和天气结果 */
export interface IpCityResult {
  city: string;
  /** 省份 */
  province: string;
  weather: import("./webInfo").WeatherInfo[];
  tip: string;
}

/** 包含创建时间的行数据 */
export interface RowWithCreateTime {
  createTime?: string;
}

/** 主题配置项 */
export interface ThemeMapConfigItem {
  title: string;
  collapseTitle: string;
  handleVal: "pc" | "mobile" | "gradient" | "solid";
  class: string;
  dataList: string[];
  style: "img" | "gradient" | "solid";
}

/** 关于页配置项 */
export interface AboutConfigItem {
  img: string;
  tit: string;
  sub: string;
}

/** 天气图标映射 */
export interface IconWeatherMap {
  [iconName: string]: string[];
}

/** 通知选项 */
export interface NotifyOptions {
  position?: "top-left" | "top-right" | "bottom-left" | "bottom-right";
  offset?: number;
  duration?: number;
}

/** 分页状态基础类型 */
export interface PaginationState {
  current: number;
  size: number;
  total: number;
}

/** 资源类型选项（下拉菜单用） */
export interface ResourceTypeOption {
  label: string;
  value: string;
}
