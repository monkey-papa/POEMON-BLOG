import constant from "./constant";
import CryptoJS from "crypto-js";
import type {
  TimeDiffResult,
  CountdownResult,
  IpCityResult,
  RowWithCreateTime,
  IpLocationResponse,
} from "@/types";

interface CommonUtils {
  getThemeRgb(): void;
  isValidHttpUrl(string: string): Promise<boolean>;
  formatter(row: RowWithCreateTime): string;
  mobile(): boolean;
  isEmpty(value: unknown): boolean;
  encrypt(plaintText: string): string;
  decrypt(encryptedBase64Str: string): string;
  faceReg(content: string): string;
  pictureReg(content: string): string;
  processContent(content: string): string;
  getDateTimeStamp(dateStr: string): number;
  getDateDiff(dateStr: string): string;
  getIpAndCity(): Promise<IpCityResult>;
  timeDiff(oldTime: string, newTime?: string | Date): TimeDiffResult;
  countdown(time: string): CountdownResult;
}

const common: CommonUtils = {
  getThemeRgb() {
    const createRgb = (r: number, g: number, b: number, opacity: number): string => {
      const newR = Math.min(Math.round(r * opacity), 255);
      const newG = Math.min(Math.round(g * opacity), 255);
      const newB = Math.min(Math.round(b * opacity), 255);
      return `${newR}, ${newG}, ${newB}`;
    };
    const rgbToHex = (rgb: string): string => {
      const parts = rgb.split(",").map((num) => parseInt(num, 10));
      const r = parts[0] ?? 0;
      const g = parts[1] ?? 0;
      const b = parts[2] ?? 0;
      const toHex = (c: number): string => {
        const hex = c.toString(16);
        return hex.length == 1 ? "0" + hex : hex;
      };
      return `#${toHex(r)}${toHex(g)}${toHex(b)}`;
    };
    // 转化为rgb值
    const root = document.querySelector(":root") as HTMLElement | null;
    if (!root) return;
    const style = getComputedStyle(root);
    const theme = style.getPropertyValue("--themeColor");
    const str = theme.replace("#", "");
    const count = str.length / 3;
    const power = 6 / str.length;
    const r = parseInt("0x" + str.substring(0 * count, 1 * count)) ** power;
    const g = parseInt("0x" + str.substring(1 * count, 2 * count)) ** power;
    const b = parseInt("0x" + str.substring(2 * count)) ** power;
    const rgb = `${r}, ${g}, ${b}`;
    const rgbDark = createRgb(r, g, b, 0.8);
    const rgbLight = createRgb(r, g, b, 1.2);
    const rgbLight1 = createRgb(r, g, b, 2);
    const hexRgbaLight = rgbToHex(rgbLight1);
    root.style.setProperty("--themeColorRgb", rgb);
    root.style.setProperty("--themeColorRgbDark", rgbDark);
    root.style.setProperty("--themeColorRgbLight", rgbLight);
    root.style.setProperty("--themeColorHexLight", hexRgbaLight);
  },
  // 判断url是否有效
  async isValidHttpUrl(string: string): Promise<boolean> {
    try {
      const newUrl = await fetch(string);
      return !!newUrl.ok;
    } catch {
      return false;
    }
  },
  /* 时间格式化 */
  formatter(row: RowWithCreateTime): string {
    if (!row?.createTime) return "";
    const createTime = String(row.createTime);
    const day = createTime.split(".")[0]?.split("T")[0] ?? "";
    const time = createTime.split(".")[0]?.split("T")[1] ?? "";
    return `${day} 日 ${time}`;
  },
  mobile(): boolean {
    const flag = navigator.userAgent.match(
      /(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i
    );
    return !!(flag && flag.length && flag.length > 0);
  },

  /**
   * 判断是否为空
   */
  isEmpty(value: unknown): boolean {
    if (
      typeof value === "undefined" ||
      value === null ||
      (typeof value === "string" && value.trim() === "") ||
      (Array.isArray(value) && value.length === 0) ||
      (typeof value === "object" &&
        value !== null &&
        !Array.isArray(value) &&
        Object.keys(value).length === 0)
    ) {
      return true;
    } else {
      return false;
    }
  },

  /**
   * 加密（AES-ECB，用于登录密码传输）
   */
  encrypt(plaintText: string): string {
    const options = {
      mode: CryptoJS.mode.ECB,
      padding: CryptoJS.pad.Pkcs7,
    };
    const key = CryptoJS.enc.Utf8.parse(constant.cryptojs_key);
    const encryptedData = CryptoJS.AES.encrypt(plaintText, key, options);
    return encryptedData.toString().replace(/\//g, "_").replace(/\+/g, "-");
  },

  /**
   * 解密
   */
  decrypt(encryptedBase64Str: string): string {
    const val = encryptedBase64Str.replace(/-/g, "+").replace(/_/g, "/");
    const options = {
      mode: CryptoJS.mode.ECB,
      padding: CryptoJS.pad.Pkcs7,
    };
    const key = CryptoJS.enc.Utf8.parse(constant.cryptojs_key);
    const decryptedData = CryptoJS.AES.decrypt(val, key, options);
    return CryptoJS.enc.Utf8.stringify(decryptedData);
  },

  /**
   * 表情包转换
   */
  faceReg(content: string): string {
    content = content.replace(/\[[^\[^\]]+\]/g, (word: string): string => {
      const index = constant.emojiList.indexOf(word.replace("[", "").replace("]", ""));
      if (index > -1) {
        const url = constant.qiniuUploadEntrance + "emoji/q" + (index + 1) + ".gif";
        return (
          '<img style="vertical-align: middle;width: 32px;height: 32px" src="' +
          url +
          '" title="' +
          word +
          '"/>'
        );
      } else {
        return word;
      }
    });
    return content;
  },

  /**
   * 图片转换
   */
  pictureReg(content: string): string {
    content = content.replace(/<[^<^>]+>/g, (word: string): string => {
      const index = word.indexOf(",");
      if (index > -1) {
        const arr = word.replace("<", "").replace(">", "").split(",");
        return (
          '<img style="border-radius: 5px;width: 100%;max-width: 250px;display: block" src="' +
          (arr[1] ?? "") +
          '" title="' +
          (arr[0] ?? "") +
          '"/>'
        );
      } else {
        return word;
      }
    });
    return content;
  },

  /**
   * 统一处理内容：换行转HTML + 表情替换 + 图片替换
   * 用于微言、树洞、评论等包含用户内容的场景
   */
  processContent(content: string): string {
    let result = content.replace(/\n{2,}/g, '<div style="height: 12px"></div>');
    result = result.replace(/\n/g, "<br/>");
    result = this.faceReg(result);
    result = this.pictureReg(result);
    return result;
  },

  /**
   * 字符串转换为时间戳
   */
  getDateTimeStamp(dateStr: string): number {
    return Date.parse(dateStr.replace(/-/gi, "/"));
  },

  getDateDiff(dateStr: string): string {
    const publishTime = isNaN(Date.parse(dateStr.replace(/-/gi, "/")) / 1000)
      ? Date.parse(dateStr) / 1000
      : Date.parse(dateStr.replace(/-/gi, "/")) / 1000;
    const timeNow = Math.floor(new Date().getTime() / 1000);
    const date = new Date(publishTime * 1000);
    const Y = date.getFullYear();
    let M: string | number = date.getMonth() + 1;
    let D: string | number = date.getDate();
    let H: string | number = date.getHours();
    let m: string | number = date.getMinutes();
    let s: string | number = date.getSeconds();
    // 小于10的在前面补0
    if (M < 10) {
      M = "0" + M;
    }
    if (D < 10) {
      D = "0" + D;
    }
    if (H < 10) {
      H = "0" + H;
    }
    if (m < 10) {
      m = "0" + m;
    }
    if (s < 10) {
      s = "0" + s;
    }
    const d = timeNow - publishTime;
    const d_days = Math.floor(d / 86400);
    const d_hours = Math.floor(d / 3600);
    const d_minutes = Math.floor(d / 60);
    const d_seconds = Math.floor(d);
    if (d_days > 0 && d_days < 3) {
      return d_days + "天前";
    } else if (d_days <= 0 && d_hours > 0) {
      return d_hours + "小时前";
    } else if (d_hours <= 0 && d_minutes > 0) {
      return d_minutes + "分钟前";
    } else if (d_seconds < 60) {
      if (d_seconds <= 0) {
        return "刚刚发表";
      } else {
        return d_seconds + "秒前";
      }
    } else if (d_days >= 3) {
      return Y + "-" + M + "-" + D + " " + H + ":" + m;
    }
    return "";
  },
  // 获取ip和城市
  getIpAndCity(): Promise<IpCityResult> {
    return new Promise((resolve) => {
      // 动态导入以避免循环依赖
      import("@/api").then(({ default: api }) => {
        api
          .getIpLocation()
          .then((res: IpLocationResponse | null) => {
            if (!this.isEmpty(res) && res?.city) {
              resolve({
                city: res.city || "",
                province: res.province || "",
                weather: res.weather || [],
                tip: res.tip || "",
              });
            } else {
              // 本地网络或无法获取位置信息时静默处理
              resolve({ city: "", province: "", weather: [], tip: "" });
            }
          })
          .catch(() => {
            resolve({ city: "", province: "", weather: [], tip: "" });
          });
      });
    });
  },
  /**
   * 计算两个时间相差的年、月、日、时、分、秒
   */
  timeDiff(oldTime: string, newTime?: string | Date): TimeDiffResult {
    oldTime = oldTime.replace(new RegExp("-", "gm"), "/");

    // 将 Date 对象转为 "YYYY/MM/DD HH:mm:ss" 格式（兼容所有浏览器）
    const formatDate = (d: Date): string =>
      d.getFullYear() +
      "/" +
      (d.getMonth() + 1) +
      "/" +
      d.getDate() +
      " " +
      d.getHours() +
      ":" +
      d.getMinutes() +
      ":" +
      d.getSeconds();

    let newTimeStr: string;
    if (newTime) {
      if (typeof newTime === "string") {
        newTimeStr = newTime.replace(new RegExp("-", "gm"), "/");
      } else {
        newTimeStr = formatDate(newTime);
      }
    } else {
      newTimeStr = formatDate(new Date());
    }

    // 计算比较日期
    const getMaxMinDate = (
      time: string,
      twoTime: string,
      type: "month" | "day"
    ): { minTime: string; maxTime: string; maxMinTong: string } => {
      const minTime = new Date(time).getTime() - new Date(twoTime).getTime() > 0 ? twoTime : time;
      const maxTime = new Date(time).getTime() - new Date(twoTime).getTime() > 0 ? time : twoTime;
      const maxDateDay = new Date(
        new Date(maxTime).getFullYear(),
        new Date(maxTime).getMonth() + 1,
        0
      ).getDate();
      const maxMinDate =
        new Date(minTime).getDate() > maxDateDay ? maxDateDay : new Date(minTime).getDate();
      let maxMinTong: string;
      if (type === "month") {
        maxMinTong =
          new Date(maxTime).getFullYear() +
          "/" +
          (new Date(minTime).getMonth() + 1) +
          "/" +
          maxMinDate +
          " " +
          new Date(minTime).toLocaleTimeString("chinese", { hour12: false });
      } else {
        maxMinTong =
          new Date(maxTime).getFullYear() +
          "/" +
          (new Date(maxTime).getMonth() + 1) +
          "/" +
          maxMinDate +
          " " +
          new Date(minTime).toLocaleTimeString("chinese", { hour12: false });
      }
      return {
        minTime,
        maxTime,
        maxMinTong,
      };
    };

    // 相差年份
    const getYear = (time: string, twoTime: string): number => {
      const oneYear = new Date(time).getFullYear();
      const twoYear = new Date(twoTime).getFullYear();
      const { maxTime, maxMinTong } = getMaxMinDate(time, twoTime, "month");
      let chaYear = Math.abs(oneYear - twoYear);
      if (new Date(maxMinTong).getTime() > new Date(maxTime).getTime()) {
        chaYear--;
      }
      return chaYear;
    };

    // 相差月份
    const getMonth = (time: string, twoTime: string, value?: number): number => {
      const oneMonth = new Date(time).getFullYear() * 12 + (new Date(time).getMonth() + 1);
      const twoMonth = new Date(twoTime).getFullYear() * 12 + (new Date(twoTime).getMonth() + 1);
      const { maxTime, maxMinTong } = getMaxMinDate(time, twoTime, "day");
      let chaMonth = Math.abs(oneMonth - twoMonth);
      if (new Date(maxMinTong).getTime() > new Date(maxTime).getTime()) {
        chaMonth--;
      }
      if (value) {
        return chaMonth - value;
      } else {
        return chaMonth;
      }
    };

    // 相差天数
    const getDay = (time: string, twoTime: string, value?: number): number => {
      const chaTime = Math.abs(new Date(time).getTime() - new Date(twoTime).getTime());
      if (value) {
        return Math.floor(chaTime / 86400000) - value;
      } else {
        return Math.floor(chaTime / 86400000);
      }
    };

    // 相差小时
    const getHour = (time: string, twoTime: string, value?: number): number => {
      const chaTime = Math.abs(new Date(time).getTime() - new Date(twoTime).getTime());
      if (value) {
        return Math.floor(chaTime / 3600000) - value;
      } else {
        return Math.floor(chaTime / 3600000);
      }
    };

    // 相差分钟
    const getMinute = (time: string, twoTime: string, value?: number): number => {
      const chaTime = Math.abs(new Date(time).getTime() - new Date(twoTime).getTime());
      if (value) {
        return Math.floor(chaTime / 60000) - value;
      } else {
        return Math.floor(chaTime / 60000);
      }
    };

    // 相差秒
    const getSecond = (time: string, twoTime: string, value?: number): number => {
      const chaTime = Math.abs(new Date(time).getTime() - new Date(twoTime).getTime());
      if (value) {
        return Math.floor(chaTime / 1000) - value;
      } else {
        return Math.floor(chaTime / 1000);
      }
    };

    // 相差年月日时分秒
    const getDiffYMDHMS = (time: string, twoTime: string): TimeDiffResult => {
      const { minTime, maxTime, maxMinTong } = getMaxMinDate(time, twoTime, "day");
      let diffDay1 = getDay(minTime, maxMinTong);
      if (new Date(maxMinTong).getTime() > new Date(maxTime).getTime()) {
        const prevMonth = new Date(maxMinTong).getMonth() - 1;
        const lastTime = new Date(maxMinTong).setMonth(prevMonth);
        diffDay1 =
          diffDay1 -
          getDay(
            new Date(lastTime).getFullYear() +
              "/" +
              (new Date(lastTime).getMonth() + 1) +
              "/" +
              new Date(lastTime).getDate(),
            maxMinTong
          );
      }
      const diffYear = getYear(time, twoTime);
      const diffMonth = getMonth(time, twoTime, diffYear * 12);
      const diffDay = getDay(time, twoTime, diffDay1);
      const diffHour = getHour(time, twoTime, getDay(time, twoTime) * 24);
      const diffMinute = getMinute(time, twoTime, getDay(time, twoTime) * 24 * 60 + diffHour * 60);
      const diffSecond = getSecond(
        time,
        twoTime,
        getDay(time, twoTime) * 24 * 60 * 60 + diffHour * 60 * 60 + diffMinute * 60
      );
      return {
        diffYear,
        diffMonth,
        diffDay,
        diffHour,
        diffMinute,
        diffSecond,
      };
    };

    return getDiffYMDHMS(oldTime, newTimeStr);
  },

  countdown(time: string): CountdownResult {
    const timeDate = new Date(time.replace(new RegExp("-", "gm"), "/"));
    const nowTime = new Date();
    // 两个时间点的时间差(秒)
    const seconds = Math.floor((timeDate.getTime() - nowTime.getTime()) / 1000);
    const d = Math.floor(seconds / 3600 / 24);
    const h = Math.floor((seconds / 3600) % 24);
    const m = Math.floor((seconds / 60) % 60);
    const s = Math.floor(seconds % 60);
    return {
      d,
      h,
      m,
      s,
    };
  },
};

export default common;
