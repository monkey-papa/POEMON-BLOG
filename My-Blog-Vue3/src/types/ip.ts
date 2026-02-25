/**
 * IP 统计与地图相关类型定义
 */

/** 省份访问统计项 */
export interface ProvinceStatItem {
  province: string;
  city: string;
  count: number;
}

/** IP 访问统计项 */
export interface IpStatItem {
  ip: string;
  count: number;
}

/** 用户访问统计项 */
export interface UserStatItem {
  userId: number;
  avatar: string | null;
  userName: string;
}

/** 单日访问统计 */
export interface DailyStats {
  province: ProvinceStatItem[];
  ip: IpStatItem[];
  users: UserStatItem[];
}

/** IP 统计响应（来自 ListIP 接口） */
export interface IpStatisticsResponse {
  total_sum: number;
  today_sum: number;
  yesterday_sum: number;
  data: {
    province_all_top: ProvinceStatItem[];
    ip_all_top: IpStatItem[];
    today: DailyStats;
    yesterday: DailyStats;
  };
}

/** GetIP 接口响应 */
export interface IpResponse {
  ip: string;
}

/** 评论/审核邮件通知参数 */
export interface CodeCommentParams {
  comment: string;
  name: string;
  /** 接收者邮箱 */
  email?: string;
  /** 接收者用户ID（可选，后端查邮箱） */
  userId?: number;
  /** 通知类型：comment/approve */
  type?: string;
}

/** 提交地区信息参数 */
export interface SubmitLocationParams {
  province: string;
  city: string;
}

/** 用户地图数据项 */
export interface UserMapData {
  name: string;
  value: string;
}
