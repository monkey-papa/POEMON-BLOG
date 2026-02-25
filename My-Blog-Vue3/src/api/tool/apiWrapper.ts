import { notifyError } from "../../utils/notify";
import common from "../../utils/common";
import type { ApiWrapperOptions } from "@/types";

interface EmptyResponseError extends Error {
  isEmptyResponse: boolean;
}

/**
 * 通用 API 请求包装器
 * 统一处理请求错误和空响应
 *
 * @param requestFn - 请求函数
 * @param errorMessage - 自定义错误消息
 * @param options - 配置选项
 * @returns API 响应数据（保证非空，除非 allowEmpty=true）
 *
 * @example
 * // 基础用法 - 响应为空会抛出错误
 * const res = await apiWrapper(() => http.post('/api/xxx'), '获取数据失败！');
 * // res 保证非空，可以直接使用
 * console.log(res.list);
 *
 * @example
 * // 允许空响应
 * const res = await apiWrapper(() => http.post('/api/xxx'), '获取数据失败！', { allowEmpty: true });
 * // res 可能为空，需要自行判断
 */
export default async function apiWrapper<T>(
  requestFn: () => Promise<T>,
  errorMessage = "请求失败！",
  options: ApiWrapperOptions = {}
): Promise<T> {
  const { allowEmpty = false, silent = false } = options;

  try {
    const response = await requestFn();

    // 检查空响应（除非明确允许）
    if (!allowEmpty && common.isEmpty(response)) {
      const emptyError = new Error("响应数据为空") as EmptyResponseError;
      emptyError.isEmptyResponse = true;
      throw emptyError;
    }

    return response;
  } catch (error) {
    // 空响应错误静默处理，不显示通知
    if ((error as EmptyResponseError).isEmptyResponse) {
      throw error;
    }

    // 其他错误显示通知（除非静默模式）
    // 优先使用后端返回的错误消息，其次使用自定义消息
    if (!silent) {
      const backendMsg = (error as Error).message;
      const displayMsg = backendMsg && backendMsg !== "请求失败" ? backendMsg : errorMessage;
      notifyError(displayMsg);
    }
    throw error;
  }
}

/**
 * 创建一个允许空响应的 API 包装器
 * @param requestFn - 请求函数
 * @param errorMessage - 自定义错误消息
 * @returns API 响应数据（可能为空）
 */
export async function apiWrapperAllowEmpty<T>(
  requestFn: () => Promise<T>,
  errorMessage: string
): Promise<T> {
  return apiWrapper(requestFn, errorMessage, { allowEmpty: true });
}
