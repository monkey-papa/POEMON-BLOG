import { notifyError, notifyWarning } from "../utils/notify";

/**
 * 表单验证工具 composable
 * 提供常用的表单验证逻辑，减少重复代码
 */

type FieldsObject = Record<string, string>;

/**
 * 验证必填字段
 * @param fields - 字段对象，格式为 { fieldValue: 'fieldName' }
 * @param message - 自定义错误消息
 * @returns 验证是否通过
 */
export function validateRequired(fields: FieldsObject, message?: string): boolean {
  for (const [value, name] of Object.entries(fields)) {
    if (!value || (typeof value === "string" && value.trim() === "")) {
      notifyError(message || `${name}不能为空！`);
      return false;
    }
  }
  return true;
}

/**
 * 验证必填字段
 * @param fields - 字段对象，格式为 { fieldValue: 'fieldName' }
 * @param customMessage - 自定义警告消息
 * @returns 验证是否通过
 */
export function validateRequiredWarn(fields: FieldsObject, customMessage?: string): boolean {
  for (const [value, name] of Object.entries(fields)) {
    if (!value || (typeof value === "string" && value.trim() === "")) {
      notifyWarning(customMessage || `你还没${name}呢~`);
      return false;
    }
  }
  return true;
}

/**
 * 验证字段不包含空格
 * @param fields - 字段对象，格式为 { fieldValue: 'fieldName' }
 * @returns 验证是否通过
 */
export function validateNoSpaces(fields: FieldsObject): boolean {
  for (const [value, name] of Object.entries(fields)) {
    if (typeof value === "string" && value.indexOf(" ") !== -1) {
      notifyError(`${name}不能包含空格！`);
      return false;
    }
  }
  return true;
}

/**
 * 验证邮箱格式
 * @param email - 邮箱地址
 * @returns 验证是否通过
 */
export function validateEmail(email: string): boolean {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  if (!emailRegex.test(email)) {
    notifyError("邮箱格式不正确！");
    return false;
  }
  return true;
}

/**
 * 验证手机号格式
 * @param phone - 手机号
 * @returns 验证是否通过
 */
export function validatePhone(phone: string): boolean {
  if (!phone) return true; // 允许为空，由 validateRequired 处理
  if (!/^1[345789]\d{9}$/.test(phone)) {
    notifyError("手机号格式有误！");
    return false;
  }
  return true;
}

/**
 * 验证 URL 格式（必须以 http 或 https 开头）
 * @param url - URL 地址
 * @param message - 自定义错误消息
 * @returns 验证是否通过
 */
export function validateUrl(url: string, message?: string): boolean {
  if (!url || !url.includes("http")) {
    notifyError(message || "请填写以http或https开头的有效地址！");
    return false;
  }
  return true;
}

interface FormValidationMethods {
  validateRequired: typeof validateRequired;
  validateRequiredWarn: typeof validateRequiredWarn;
  validateNoSpaces: typeof validateNoSpaces;
  validateEmail: typeof validateEmail;
  validatePhone: typeof validatePhone;
  validateUrl: typeof validateUrl;
}

/**
 * 使用表单验证 composable
 * 提供统一的验证方法和状态管理
 * 注意：登录验证为后端处理，前端通过 401 响应自动跳转登录页
 */
export function useFormValidation(): FormValidationMethods {
  return {
    validateRequired,
    validateRequiredWarn,
    validateNoSpaces,
    validateEmail,
    validatePhone,
    validateUrl,
  };
}

export default useFormValidation;
