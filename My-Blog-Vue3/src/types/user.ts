/**
 * 用户相关类型定义
 */

import type { Prettify } from "./common";

/** 用户类型枚举 */
export enum UserType {
  /** Boss 管理员 */
  Boss = 0,
  /** 普通管理员 */
  Admin = 1,
  /** 普通用户 */
  User = 2,
  /** 访客 */
  Visitor = 3,
}

/**
 * 登录/注册响应
 * 前台和后台登录返回相同结构
 */
export interface CurrentUser {
  id: number;
  username: string;
  /** 头像 */
  avatar: string;
  /** 邮箱 */
  email: string | null;
  /** 简介 */
  introduction: string | null;
  userType: UserType;
  accessToken: string;
  /** 赞赏二维码 */
  admire: string | null;
  userStatus: boolean;
  /** 手机号 */
  phoneNumber: string | null;
  /** 性别：0-保密，1-男，2-女 */
  gender: number | null;
  createTime: string;
  /** 七牛云域名 */
  qiniuDomain: string | null;
  /** 七牛云 AccessKey */
  qiniuAccessKey: string | null;
  /** 七牛云 SecretKey */
  qiniuSecretKey: string | null;
  /** 七牛云 Bucket 名称 */
  qiniuBucketName: string | null;
}

/** 当前管理员信息 */
export type CurrentAdmin = CurrentUser;

/** 用户列表项-管理员视图 */
export interface AdminUserListItem {
  id: number;
  username: string;
  phoneNumber: string | null;
  email: string | null;
  admire: string | null;
  userStatus: boolean;
  avatar: string | null;
  gender: number | null;
  introduction: string | null;
  userType: UserType;
  createTime: string;
  province: string | null;
}

/** 用户列表项-前台视图 */
export interface PublicUserListItem {
  id: number;
  username: string;
  avatar: string | null;
  gender: number | null;
}

/** 通用用户列表项（兼容前台和后台） */
export type UserListItem = Prettify<AdminUserListItem>;

/** 登录请求参数 */
export interface LoginParams {
  account: string;
  password: string;
  province?: string;
}

/** 管理员登录请求参数 */
export interface AdminLoginParams {
  account: string;
  password: string;
}

/** 注册请求参数 */
export interface RegisterParams {
  username: string;
  password: string;
  email: string;
  code: string;
  province?: string;
}

/** 更新用户信息参数 */
export interface UpdateUserParams {
  userId?: number;
  username?: string;
  avatar?: string;
  email?: string;
  introduction?: string;
  phoneNumber?: string;
  gender?: number;
  code?: string;
  qiniuDomain?: string;
  qiniuAccessKey?: string;
  qiniuSecretKey?: string;
  qiniuBucketName?: string;
}

/** 忘记密码参数 */
export interface ForgetPasswordParams {
  username: string;
  password: string;
  email: string;
  code: string;
}

/** 发送验证码参数 */
export interface SendCodeParams {
  email: string;
}

/** 修改用户状态参数 */
export interface ChangeUserStatusParams {
  userId: number;
  status: boolean;
}

/** 修改用户赞赏参数 */
export interface ChangeUserAdmireParams {
  userId: number;
  admire: string;
}

/** 修改用户类型参数 */
export interface ChangeUserTypeParams {
  userId: number;
  userType: UserType;
}
