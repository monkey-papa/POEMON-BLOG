/**
 * 表白墙相关类型定义
 */

/** 表白墙信息 */
export interface Family {
  id: number;
  userId: number;
  bgCover: string;
  manCover: string;
  womanCover: string;
  manName: string;
  womanName: string;
  timing: string;
  /** 倒计时标题 */
  countdownTitle: string | null;
  /** 倒计时时间 */
  countdownTime: string | null;
  /** 表白内容 */
  familyInfo: string | null;
  likeCount: number;
  status: boolean;
  createTime: string;
}

/** 添加表白墙参数 */
export interface AddFamilyParams {
  bgCover: string;
  manCover?: string;
  womanCover?: string;
  manName?: string;
  womanName?: string;
  timing?: string;
  countdownTitle?: string;
  countdownTime?: string;
  familyInfo?: string;
  status?: boolean;
  likeCount?: number;
}

/** 修改表白墙状态参数 */
export interface ChangeLoveStatusParams {
  id: number;
  flag: boolean;
}
