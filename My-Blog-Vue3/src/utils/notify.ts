import { ElNotification } from "element-plus";
import type { NotifyOptions } from "@/types";

/**
 * 通知提示工具函数
 * 统一管理项目中的通知消息，提供简洁的 API
 */

interface _NotificationOptions extends NotifyOptions {
  type?: "success" | "error" | "warning" | "info";
  title?: string;
  message?: string;
}
// 保留供扩展使用
export type { _NotificationOptions as NotificationOptions };

const defaultOptions: NotifyOptions = {
  position: "top-left",
  offset: 50,
  duration: 3000,
};

/**
 * 成功提示
 * @param message - 提示消息
 * @param title - 提示标题，默认为 "成功 🍨"
 * @param options - 其他配置选项
 */
export function notifySuccess(
  message: string,
  title = "可以啦🍨",
  options: NotifyOptions = {}
): ReturnType<typeof ElNotification> {
  return ElNotification({
    type: "success",
    title,
    message,
    ...defaultOptions,
    ...options,
  });
}

/**
 * 错误提示
 * @param message - 提示消息
 * @param title - 提示标题，默认为 "可恶🤬"
 * @param options - 其他配置选项
 */
export function notifyError(
  message: string,
  title = "可恶🤬",
  options: NotifyOptions = {}
): ReturnType<typeof ElNotification> {
  return ElNotification({
    type: "error",
    title,
    message,
    ...defaultOptions,
    ...options,
  });
}

/**
 * 警告提示
 * @param message - 提示消息
 * @param title - 提示标题，默认为 "淘气👻"
 * @param options - 其他配置选项
 */
export function notifyWarning(
  message: string,
  title = "淘气👻",
  options: NotifyOptions = {}
): ReturnType<typeof ElNotification> {
  return ElNotification({
    type: "warning",
    title,
    message,
    ...defaultOptions,
    ...options,
  });
}

/**
 * 信息提示
 * @param message - 提示消息
 * @param title - 提示标题
 * @param options - 其他配置选项
 */
export function notifyInfo(
  message: string,
  title = "提示",
  options: NotifyOptions = {}
): ReturnType<typeof ElNotification> {
  return ElNotification({
    type: "info",
    title,
    message,
    ...defaultOptions,
    ...options,
  });
}

// 默认导出所有通知函数
export default {
  success: notifySuccess,
  error: notifyError,
  warning: notifyWarning,
  info: notifyInfo,
};
