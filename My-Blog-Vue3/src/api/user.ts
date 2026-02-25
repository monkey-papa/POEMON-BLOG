import http from "../utils/request";
import constant from "../utils/constant";
import apiWrapper from "./tool/apiWrapper";
import type {
  CurrentUser,
  CurrentAdmin,
  AdminUserListItem,
  PublicUserListItem,
  LoginParams,
  AdminLoginParams,
  RegisterParams,
  ForgetPasswordParams,
  SendCodeParams,
  UpdateUserParams,
  ChangeUserStatusParams,
  ChangeUserAdmireParams,
  ChangeUserTypeParams,
  PaginationResponse,
  PaginationParams,
} from "@/types";

// ==================== 用户相关 API ====================

/**
 * 用户登录
 * @param user - 用户信息 { account, password, province }
 * @returns 用户数据
 */
export function login(user: LoginParams): Promise<CurrentUser> {
  return apiWrapper(
    () => http.post<CurrentUser>(`${constant.baseURL}/user/login`, user),
    "登录失败！"
  );
}

/**
 * 管理员登录
 * @param user - 管理员信息
 * @returns 管理员数据
 */
export function adminLogin(user: AdminLoginParams): Promise<CurrentAdmin> {
  return apiWrapper(
    () => http.post<CurrentAdmin>(`${constant.baseURL}/admin/user/login`, user),
    "登录失败！"
  );
}

/**
 * 用户注册
 * @param user - 用户信息 { username, password, email, code, province }
 * @returns 用户信息（包含 accessToken）
 */
export function register(user: RegisterParams): Promise<CurrentUser> {
  return apiWrapper(
    () => http.post<CurrentUser>(`${constant.baseURL}/user/registration`, user),
    "注册失败！"
  );
}

/**
 * 发送验证码
 * @param params - 邮箱参数 { email }
 */
export function sendVerificationCode(params: SendCodeParams): Promise<void> {
  return apiWrapper(() => http.post<void>(`${constant.baseURL}/code`, params), "验证码发送失败！", {
    allowEmpty: true,
  });
}

/**
 * 更新用户信息
 * @param user - 用户信息
 * @returns 更新后的用户信息
 */
export function updateUserInfo(user: UpdateUserParams): Promise<CurrentUser> {
  return apiWrapper(
    () => http.post<CurrentUser>(`${constant.baseURL}/user/updateUserInfo`, user),
    "更新用户信息失败！"
  );
}

/**
 * 忘记密码重置
 * @param params - 重置密码参数（需包含 username）
 */
export function updateForForgetPassword(params: ForgetPasswordParams): Promise<void> {
  return apiWrapper(
    () => http.post<void>(`${constant.baseURL}/user/updateForForgetPassword`, params),
    "重置密码失败！",
    { allowEmpty: true }
  );
}

/**
 * 获取用户列表（前台，仅公开信息）
 * @returns 用户列表（分页响应）
 */
export function getUserListPaginated(): Promise<PaginationResponse<PublicUserListItem>> {
  return apiWrapper(
    () =>
      http.post<PaginationResponse<PublicUserListItem>>(`${constant.baseURL}/user/list`, {
        noPagination: true,
        isAdmin: false,
      }),
    "获取用户列表失败！"
  );
}

/**
 * 获取管理员用户列表（后台，完整信息）
 * @param pagination - 分页参数
 * @returns 用户列表
 */
export function getAdminUserList(
  pagination: PaginationParams
): Promise<PaginationResponse<AdminUserListItem>> {
  return apiWrapper(
    () =>
      http.post<PaginationResponse<AdminUserListItem>>(`${constant.baseURL}/user/list`, {
        ...pagination,
        isAdmin: true,
      }),
    "获取用户列表失败！"
  );
}

/**
 * 修改用户状态（管理员）
 * @param params - { userId, status }
 */
export function changeUserStatus(params: ChangeUserStatusParams): Promise<void> {
  return apiWrapper(
    () =>
      http.post<void>(`${constant.baseURL}/admin/user/updateAttribute`, {
        userId: params.userId,
        field: "status",
        boolFlag: params.status,
      }),
    "修改用户状态失败！",
    { allowEmpty: true }
  );
}

/**
 * 修改用户赞赏二维码（管理员）
 * @param params - { userId, admire }
 */
export function changeUserAdmire(params: ChangeUserAdmireParams): Promise<void> {
  return apiWrapper(
    () =>
      http.post<void>(`${constant.baseURL}/admin/user/updateAttribute`, {
        userId: params.userId,
        field: "admire",
        value: params.admire,
      }),
    "修改赞赏二维码失败！",
    { allowEmpty: true }
  );
}

/**
 * 修改用户类型（管理员）
 * @param params - { userId, userType }
 */
export function changeUserType(params: ChangeUserTypeParams): Promise<void> {
  return apiWrapper(
    () =>
      http.post<void>(`${constant.baseURL}/admin/user/updateAttribute`, {
        userId: params.userId,
        field: "type",
        intValue: params.userType,
      }),
    "修改用户类型失败！",
    { allowEmpty: true }
  );
}

export default {
  login,
  adminLogin,
  register,
  sendVerificationCode,
  updateUserInfo,
  updateForForgetPassword,
  getUserListPaginated,
  getAdminUserList,
  changeUserStatus,
  changeUserAdmire,
  changeUserType,
};
