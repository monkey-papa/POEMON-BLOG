import constant from "./constant";
import CryptoJS from "crypto-js";
import axios from "axios";

export default {
  // 判断url是否有效
  async isValidHttpUrl(string) {
    try {
      const newUrl = await fetch(string);
      return newUrl.ok ? true : false;
    } catch (err) {
      return false;
    }
  },
  /*时间格式化 */
  formatter(row) {
    const day = row.createTime.split(".")[0].split("T")[0];
    const time = row.createTime.split(".")[0].split("T")[1];
    return `${day} 日 ${time}`;
  },
  mobile() {
    let flag = navigator.userAgent.match(
      /(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i
    );
    return flag && flag.length && flag.length > 0;
  },

  /**
   * 判断是否为空
   */
  isEmpty(value) {
    if (
      typeof value === "undefined" ||
      value === null ||
      (typeof value === "string" && value.trim() === "") ||
      (Array.prototype.isPrototypeOf(value) && value.length === 0) ||
      (Object.prototype.isPrototypeOf(value) && Object.keys(value).length === 0)
    ) {
      return true;
    } else {
      return false;
    }
  },

  /**
   * 加密
   */
  encrypt(plaintText) {
    let options = {
      mode: CryptoJS.mode.ECB,
      padding: CryptoJS.pad.Pkcs7,
    };
    let key = CryptoJS.enc.Utf8.parse(constant.cryptojs_key);
    let encryptedData = CryptoJS.AES.encrypt(plaintText, key, options);
    return encryptedData.toString().replace(/\//g, "_").replace(/\+/g, "-");
  },

  /**
   * 解密
   */
  decrypt(encryptedBase64Str) {
    let val = encryptedBase64Str.replace(/\-/g, "+").replace(/_/g, "/");
    let options = {
      mode: CryptoJS.mode.ECB,
      padding: CryptoJS.pad.Pkcs7,
    };
    let key = CryptoJS.enc.Utf8.parse(constant.cryptojs_key);
    let decryptedData = CryptoJS.AES.decrypt(val, key, options);
    return CryptoJS.enc.Utf8.stringify(decryptedData);
  },

  /**
   * 表情包转换
   */
  faceReg(content) {
    content = content.replace(/\[[^\[^\]]+\]/g, (word) => {
      let index = constant.emojiList.indexOf(
        word.replace("[", "").replace("]", "")
      );
      if (index > -1) {
        let url =
          constant.qiniuUploadEntrance + "emoji/q" + (index + 1) + ".gif";
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
  pictureReg(content) {
    content = content.replace(/\<[^\<^\>]+\>/g, (word) => {
      let index = word.indexOf(",");
      if (index > -1) {
        let arr = word.replace("<", "").replace(">", "").split(",");
        return (
          '<img style="border-radius: 5px;width: 100%;max-width: 250px;display: block" src="' +
          arr[1] +
          '" title="' +
          arr[0] +
          '"/>'
        );
      } else {
        return word;
      }
    });
    return content;
  },

  /**
   * 字符串转换为时间戳
   */
  getDateTimeStamp(dateStr) {
    return Date.parse(dateStr.replace(/-/gi, "/"));
  },

  getDateDiff(dateStr) {
    let publishTime = isNaN(Date.parse(dateStr.replace(/-/gi, "/")) / 1000)
      ? Date.parse(dateStr) / 1000
      : Date.parse(dateStr.replace(/-/gi, "/")) / 1000;
    let d_seconds,
      d_minutes,
      d_hours,
      d_days,
      timeNow = Math.floor(new Date().getTime() / 1000),
      d,
      date = new Date(publishTime * 1000),
      Y = date.getFullYear(),
      M = date.getMonth() + 1,
      D = date.getDate(),
      H = date.getHours(),
      m = date.getMinutes(),
      s = date.getSeconds();
    //小于10的在前面补0
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
    d = timeNow - publishTime;
    d_days = Math.floor(d / 86400);
    d_hours = Math.floor(d / 3600);
    d_minutes = Math.floor(d / 60);
    d_seconds = Math.floor(d);
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
  },

  /**
   * 保存资源
   */
  saveResource(that, type, path, size, mimeType, id, isAdmin = false) {
    let resource = {
      type: type,
      path: path,
      size: size,
      mimeType: mimeType,
      id: id,
    };
    that.$http
      .post(
        that.$constant.baseURL + "/resource/saveResource/",
        resource,
        isAdmin
      )
      .catch((error) => {
        that.$message({
          message: error.message,
          type: "error",
        });
      });
  },
  // 获取ip和城市
  getIpAndCity(that) {
    return new Promise((resolve) => {
      let city = "";
      let address = "";
      let weather = "";
      that.$http
        .get(that.$constant.baseURL + "/ip/")
        .then(async (res) => {
          if (!that.$common.isEmpty(res.result[0].data[0].ip)) {
            const ip = res.result[0].data[0].ip;
            //高德
            const result1 = await axios.get(
              `https://restapi.amap.com/v3/ip?ip=${ip}&key=407cd13370e3e36bcb96759e9b08d958`
            );
            let u = navigator.userAgent;
            let isiOS = !!u.match(/\(i[^;]+;( U;)? CPU.+Mac OS X/); //ios终端
            if (
              !typeof result1.data.rectangle.split == "function" ||
              !result1.data.rectangle ||
              isiOS
            ) {
              that.$message({
                message: "由于设备隐私问题，无法获取您当前地点的天气信息",
                type: "error",
              });
              return;
            }
            const pos0 = Number(
              result1.data.rectangle.split(";")[0].split(",")[0]
            );
            const pos1 = Number(
              result1.data.rectangle.split(";")[0].split(",")[1]
            );
            const pos2 = Number(
              result1.data.rectangle.split(";")[1].split(",")[0]
            );
            const pos3 = Number(
              result1.data.rectangle.split(";")[1].split(",")[1]
            );
            const pos4 = (pos0 + pos2) / 2;
            const pos5 = (pos1 + pos3) / 2;
            const pos = pos4 + "," + pos5;
            const result2 = await axios.get(
              `https://restapi.amap.com/v3/geocode/regeo?output=JSON&location=${pos}&key=407cd13370e3e36bcb96759e9b08d958&radius=0&extensions=all`
            );
            city = result2.data.regeocode.addressComponent.city;
            address = result2.data.regeocode.formatted_address;
            const result3 = await axios.get(
              `https://restapi.amap.com/v3/weather/weatherInfo?key=407cd13370e3e36bcb96759e9b08d958&city=${city}&output=json&extensions=all`
            );
            weather = result3.data.forecasts[0];
            resolve({ city, address, weather });
          } else {
            that.$message({
              message: "由于设备隐私问题，无法获取您当前地点的天气信息",
              type: "error",
            });
          }
        })
        .catch((error) => {
          that.$message({
            message: "由于设备隐私问题，无法获取您当前地点的天气信息",
            type: "error",
          });
        });
    });
  },
  /**
   * 计算两个时间相差的年、月、日、时、分、秒
   *
   *
   */
  timeDiff(oldTime, newTime) {
    oldTime = oldTime.replace(new RegExp("-", "gm"), "/");
    if (newTime) {
      newTime = newTime.replace(new RegExp("-", "gm"), "/");
    } else {
      newTime = new Date();
    }

    // 计算比较日期
    const getMaxMinDate = (time, twoTime, type) => {
      let minTime =
        new Date(time).getTime() - new Date(twoTime).getTime() > 0
          ? twoTime
          : time;
      let maxTime =
        new Date(time).getTime() - new Date(twoTime).getTime() > 0
          ? time
          : twoTime;
      let maxDateDay = new Date(
        new Date(maxTime).getFullYear(),
        new Date(maxTime).getMonth() + 1,
        0
      ).getDate();
      let maxMinDate =
        new Date(minTime).getDate() > maxDateDay
          ? maxDateDay
          : new Date(minTime).getDate();
      let maxMinTong;
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
    const getYear = (time, twoTime) => {
      let oneYear = new Date(time).getFullYear();
      let twoYear = new Date(twoTime).getFullYear();
      const { minTime, maxTime, maxMinTong } = getMaxMinDate(
        time,
        twoTime,
        "month"
      );
      let chaYear = Math.abs(oneYear - twoYear);
      if (new Date(maxMinTong).getTime() > new Date(maxTime).getTime()) {
        chaYear--;
      }
      return chaYear;
    };

    // 相差月份
    const getMonth = (time, twoTime, value) => {
      let oneMonth =
        new Date(time).getFullYear() * 12 + (new Date(time).getMonth() + 1);
      let twoMonth =
        new Date(twoTime).getFullYear() * 12 +
        (new Date(twoTime).getMonth() + 1);
      const { minTime, maxTime, maxMinTong } = getMaxMinDate(
        time,
        twoTime,
        "day"
      );
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
    const getDay = (time, twoTime, value) => {
      let chaTime = Math.abs(
        new Date(time).getTime() - new Date(twoTime).getTime()
      );
      if (value) {
        return parseInt(chaTime / 86400000) - value;
      } else {
        return parseInt(chaTime / 86400000);
      }
    };

    // 相差小时
    const getHour = (time, twoTime, value) => {
      let chaTime = Math.abs(
        new Date(time).getTime() - new Date(twoTime).getTime()
      );
      if (value) {
        return parseInt(chaTime / 3600000) - value;
      } else {
        return parseInt(chaTime / 3600000);
      }
    };

    // 相差分钟
    const getMinute = (time, twoTime, value) => {
      let chaTime = Math.abs(
        new Date(time).getTime() - new Date(twoTime).getTime()
      );
      if (value) {
        return parseInt(chaTime / 60000) - value;
      } else {
        return parseInt(chaTime / 60000);
      }
    };

    // 相差秒
    const getSecond = (time, twoTime, value) => {
      let chaTime = Math.abs(
        new Date(time).getTime() - new Date(twoTime).getTime()
      );
      if (value) {
        return parseInt(chaTime / 1000) - value;
      } else {
        return parseInt(chaTime / 1000);
      }
    };

    // 相差年月日时分秒
    const getDiffYMDHMS = (time, twoTime) => {
      const { minTime, maxTime, maxMinTong } = getMaxMinDate(
        time,
        twoTime,
        "day"
      );
      let diffDay1 = getDay(minTime, maxMinTong);
      if (new Date(maxMinTong).getTime() > new Date(maxTime).getTime()) {
        let prevMonth = new Date(maxMinTong).getMonth() - 1;
        let lastTime = new Date(maxMinTong).setMonth(prevMonth);
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
      let diffYear = getYear(time, twoTime);
      let diffMonth = getMonth(time, twoTime, diffYear * 12);
      let diffDay = getDay(time, twoTime, diffDay1);
      let diffHour = getHour(time, twoTime, getDay(time, twoTime) * 24);
      let diffMinute = getMinute(
        time,
        twoTime,
        getDay(time, twoTime) * 24 * 60 + diffHour * 60
      );
      let diffSecond = getSecond(
        time,
        twoTime,
        getDay(time, twoTime) * 24 * 60 * 60 +
          diffHour * 60 * 60 +
          diffMinute * 60
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

    return getDiffYMDHMS(oldTime, newTime);
  },

  countdown(time) {
    time = new Date(time.replace(new RegExp("-", "gm"), "/"));
    let nowTime = new Date();
    //两个时间点的时间差(秒)
    let seconds = parseInt((time.getTime() - nowTime.getTime()) / 1000);
    let d = parseInt(seconds / 3600 / 24);
    let h = parseInt((seconds / 3600) % 24);
    let m = parseInt((seconds / 60) % 60);
    let s = parseInt(seconds % 60);
    return {
      d,
      h,
      m,
      s,
    };
  },
};
