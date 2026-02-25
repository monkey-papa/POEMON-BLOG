import http from "../utils/request";
import constant from "../utils/constant";
import apiWrapper from "./tool/apiWrapper";
import type {
  IpLocationResponse,
  IpStatisticsResponse,
  IpResponse,
  CodeCommentParams,
  SubmitLocationParams,
  UserMapData,
} from "@/types";

// ==================== 其他 API ====================

/**
 * 发送评论/审核邮件通知
 * @param params - 通知参数
 */
export function sendCodeComment(params: CodeCommentParams): Promise<void> {
  return apiWrapper(
    () => http.post<void>(`${constant.baseURL}/codeComment`, params),
    "发送通知失败！",
    { allowEmpty: true }
  );
}

/**
 * 获取IP统计信息（访问量、省份/IP/用户统计）
 * @returns IP统计响应
 */
export function getIpStatistics(): Promise<IpStatisticsResponse> {
  return apiWrapper(
    () => http.post<IpStatisticsResponse>(`${constant.baseURL}/list/ip`),
    "获取IP统计失败！"
  );
}

/**
 * 提交地区信息（记录IP访问）
 * @param params - 地区参数
 */
export function submitLocation(params: SubmitLocationParams): Promise<void> {
  return apiWrapper(
    () => http.post<void>(`${constant.baseURL}/submit`, params),
    "提交地区信息失败！",
    { allowEmpty: true, silent: true }
  );
}

/**
 * 获取地图 GeoJSON 数据
 * @returns 地图数据（静态 JSON 文件内容）
 */
export function getMap(): Promise<unknown> {
  return apiWrapper(() => http.post<unknown>(`${constant.baseURL}/map`), "获取地图数据失败！");
}

/**
 * 获取客户端 IP 地址
 * @returns { ip: string }
 */
export function getIpAddress(): Promise<IpResponse> {
  return apiWrapper(() => http.post<IpResponse>(`${constant.baseURL}/ip`), "获取 IP 信息失败！");
}

/**
 * 获取 IP 定位和天气信息
 * @returns 定位和天气数据
 */
export function getIpLocation(): Promise<IpLocationResponse> {
  return apiWrapper(
    () => http.post<IpLocationResponse>(`${constant.baseURL}/ip/location`),
    "获取定位信息失败！",
    { silent: true }
  );
}

/**
 * 获取用户地图数据（用户按省份分布）
 * @returns 用户分布数据（注意 value 为 string 类型）
 */
export function getUserMapData(): Promise<UserMapData[]> {
  return apiWrapper(
    () => http.post<UserMapData[]>(`${constant.baseURL}/user/map`),
    "获取用户分布失败！"
  );
}

export default {
  sendCodeComment,
  getIpStatistics,
  submitLocation,
  getMap,
  getIpAddress,
  getIpLocation,
  getUserMapData,
};
