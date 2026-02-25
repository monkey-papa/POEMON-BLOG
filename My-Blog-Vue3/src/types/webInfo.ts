/**
 * 网站信息相关类型定义
 */

/** 网站信息 */
export interface WebInfo {
  id: number;
  webName: string;
  webTitle: string[] | string;
  notices: string[] | string;
  footer: string;
  backgroundImage: string;
  avatar: string;
  randomCover: string[] | string;
  /** 看板娘配置 JSON */
  waifuJson?: string;
  status: boolean;
  /** 随机头像配置 */
  randomAvatar?: string[] | string;
  /** 随机名称配置 */
  randomName?: string[] | string;
}

/** 位置信息 */
export interface LocationInfo {
  city: string;
  /** 省份 */
  province: string;
  weather: WeatherInfo[];
  tip: string;
  isLoading: boolean;
  isLoaded: boolean;
}

/** 天气信息 */
export interface WeatherInfo {
  /** 日期 YYYY-MM-DD */
  date: string;
  /** 星期几 */
  week: string;
  /** 天气类型（如"晴"、"多云"） */
  type: string;
  /** 最高温度（如"10°C"） */
  high: string;
  /** 最低温度（如"0°C"） */
  low: string;
  /** 风力等级 */
  fengli: string;
  /** 风向 */
  fengxiang: string;
}

/** IP定位响应 */
export interface IpLocationResponse {
  ip: string;
  city: string;
  province: string;
  weather: WeatherInfo[];
  tip: string;
}

/** 页面浏览量/IP统计 */
export interface PageView {
  total_sum?: number;
  yesterday_sum?: number;
  today_sum?: number;
  data?: {
    today?: { users?: unknown[] };
    yesterday?: { users?: unknown[] };
  };
  [key: string]: unknown;
}

/** 树洞留言 */
export interface TreeHole {
  id: number;
  /** 发布者ID */
  userId: number;
  /** 发布者用户名 */
  username: string;
  message: string;
  /** 头像 */
  avatar: string | null;
  createTime: string;
}

/** 保存树洞参数 */
export interface SaveTreeHoleParams {
  message: string;
  avatar: string;
}

/** 收藏项 */
export interface CollectItem {
  id: number;
  title: string;
  url: string | null;
  cover?: string;
  introduction?: string;
  classify?: string;
  type: string;
  remark?: string;
  status: boolean;
  createTime: string;
}

/** 收藏列表响应（按 classify 分组的 Map） */
export interface CollectListResponse {
  [classify: string]: CollectItem[];
}
