import type { ThemeMapConfigItem, AboutConfigItem, IconWeatherMap } from "@/types";

interface ConstantConfig {
  baseURL: string;
  webURL: string;
  siteURL: string;
  hitokoto: string;
  shehui: string;
  jinrishici: string;
  qiniuUploadImages: string;
  qiniuUploadEntrance: string;
  favoriteVideo: string;
  before_color_1: string;
  after_color_1: string;
  commentPageColor: string;
  userId: number;
  // 站点 & 联系方式配置
  siteName: string;
  bossEmail: string;
  qqNumber: string;
  wechatId: string;
  // 默认图片配置
  defaultAvatar: string;
  defaultBackground: string;
  gameURL: string;
  metingApiURL: string;
  emojiList: string[];
  SolidColor: string[];
  gradient: string[];
  about: AboutConfigItem[];
  themeMapConfig: ThemeMapConfigItem[];
  tree_hole_color: string[];
  iconWeatherMap: IconWeatherMap;
  cryptojs_key: string;
}

const constant: ConstantConfig = {
  baseURL: import.meta.env.VITE_API_BASE_URL || "http://localhost:8000/api",
  webURL: import.meta.env.VITE_WEB_URL || "http://localhost:8000",
  siteURL: import.meta.env.VITE_SITE_URL || "http://localhost:81",
  // 诗词语录
  hitokoto: "https://v1.hitokoto.cn",
  shehui: "https://api.oick.cn/yulu/api.php",
  jinrishici: "https://v1.jinrishici.com/all.json",
  // 上传图片文件地址 elementUI:action
  qiniuUploadImages: `${
    import.meta.env.VITE_API_BASE_URL || "http://localhost:8000/api"
  }/resource/updateImage`,
  // 表情地址,md图片地址
  qiniuUploadEntrance: "", // 图床地址，请替换为你自己的
  // 视频地址
  favoriteVideo: "", // 视频地址，请替换为你自己的
  // 按钮颜色
  before_color_1: "var(--blue13)",
  after_color_1: "linear-gradient(45deg, var(--red), var(--purple1))",
  // 评论分页颜色 微言分页颜色
  commentPageColor: "var(--green6)",
  userId: 1, // 博客主人的用户id（注册的第一个用户）
  // 站点 & 联系方式配置（请替换为你自己的信息）
  siteName: "My-Blog", // 站点名称
  bossEmail: "", // 博主邮箱（用于审核通知等）
  qqNumber: "", // 博主QQ号
  wechatId: "", // 博主微信号
  // 加密密钥
  cryptojs_key: "undefined",
  // 默认图片配置
  defaultAvatar: "", // 默认头像地址，请替换为你自己的
  defaultBackground: "", // 默认背景图地址，请替换为你自己的
  // 外部服务URL
  gameURL: "https://game.eean.cn/pc/game",
  metingApiURL: "https://api.i-meto.com/meting/api",
  // emoji含义
  emojiList: [
    "衰",
    "鄙视",
    "再见",
    "捂嘴",
    "摸鱼",
    "奋斗",
    "白眼",
    "可怜",
    "皱眉",
    "鼓掌",
    "烦恼",
    "吐舌",
    "挖鼻",
    "委屈",
    "滑稽",
    "啊这",
    "生气",
    "害羞",
    "晕",
    "好色",
    "流泪",
    "吐血",
    "微笑",
    "酷",
    "坏笑",
    "吓",
    "大兵",
    "哭笑",
    "困",
    "呲牙",
  ],
  // 纯色
  SolidColor: [
    "#f7f9fe",
    "#30303c",
    "#7a7374",
    "#eea6b7",
    "#918072",
    "#fbecde",
    "#a4aca7",
    "#a4cab6",
    "#83a78d",
    "#70887d",
    "#57c3c2",
    "#10aec2",
    "#93d5dc",
    "#3b818c",
    "#5cb3cc",
    "#93b5cf",
  ],
  // 渐变
  gradient: [
    "55deg, #0095c2 21%, #64E1C8 100%",
    "90deg, #ffd7e4 0%, #c8f1ff 100%",
    "45deg, #e5737b, #c6999e, #96b9c2, #00d6e8",
    "25deg, #31354b, #38536c, #3b738e, #3995b2",
    "26deg, #0e6183, #387ea6, #599dcb, #7abdf1",
    "25deg, #0583bf, #159bc5, #16b4cb, #0aced0",
    "25deg, #3e47d1, #8b5fb8, #ba7b9d, #df9980",
    "25deg, #0e5c71, #15828f, #19a9ae, #1ad3ce",
  ],
  // 关于页图片配置（请替换为你自己的图片）
  about: [
    {
      img: "",
      tit: "示例标题",
      sub: "示例描述",
    },
  ],
  themeMapConfig: [
    {
      title: "1. 图片（电脑）",
      collapseTitle: "查看适配电脑端背景",
      handleVal: "pc",
      class: "box",
      dataList: [],
      style: "img",
    },
    {
      title: "2. 图片（移动端）",
      collapseTitle: "查看适配移动端背景",
      handleVal: "mobile",
      class: "box mobileBox",
      dataList: [],
      style: "img",
    },
    {
      title: "3. 渐变色",
      collapseTitle: "查看渐变色背景",
      handleVal: "gradient",
      class: "box",
      dataList: [],
      style: "gradient",
    },
    {
      title: "4. 纯色",
      collapseTitle: "查看纯色背景",
      handleVal: "solid",
      class: "box",
      dataList: [],
      style: "solid",
    },
  ],
  // 随机 微言颜色 标签颜色
  tree_hole_color: (function (): string[] {
    function getRandomColor(): string {
      return `rgb(${Math.random() * 255}, ${Math.random() * 255}, ${Math.random() * 255})`;
    }
    const colors: string[] = [];
    for (let i = 0; i < 6; i++) {
      colors.push(getRandomColor());
    }
    return colors;
  })(),
  // 天气图标映射
  iconWeatherMap: {
    "icon-dafeng": [
      "有风",
      "平静",
      "微风",
      "和风",
      "清风",
      "强风/劲风",
      "强风",
      "劲风",
      "疾风",
      "大风",
      "烈风",
      "风暴",
      "狂爆风",
      "飓风",
      "热带风暴",
      "龙卷风",
      // 风向（带风字的都归到风图标）
      "北风",
      "东北风",
      "东风",
      "东南风",
      "南风",
      "西南风",
      "西风",
      "西北风",
      "东北偏北风",
      "东北偏东风",
      "东南偏东风",
      "东南偏南风",
      "西南偏南风",
      "西南偏西风",
      "西北偏西风",
      "西北偏北风",
    ],
    "icon-duoyunzhuanqing": [
      "少云",
      "晴间多云",
      "多云",
      "多云转晴",
      "晴转多云",
      "局部多云",
      "大部多云",
    ],
    "icon-jiangxue": [
      "雪",
      "阵雪",
      "小雪",
      "中雪",
      "大雪",
      "暴雪",
      "小雪-中雪",
      "中雪-大雪",
      "大雪-暴雪",
      "冷",
      "小到中雪",
      "中到大雪",
      "大到暴雪",
      "阵雨夹雪转小雪",
      "暴风雪",
    ],
    "icon-zhongwu": [
      "浮尘",
      "扬沙",
      "沙尘暴",
      "强沙尘暴",
      "雾",
      "浓雾",
      "强浓雾",
      "轻雾",
      "大雾",
      "特强浓雾",
      "霾",
      "薄雾",
    ],
    "icon-qingtian": ["晴", "热", "晴朗", "天晴"],
    "icon-yujiaxue": [
      "雨雪天气",
      "雨夹雪",
      "阵雨夹雪",
      "雪转雨夹雪",
      "雨夹雪转雪",
      "小雨夹雪",
      "中雨夹雪",
      "大雨夹雪",
    ],
    "icon-xiaoyu": [
      "阵雨",
      "雷阵雨",
      "雷阵雨并伴有冰雹",
      "小雨",
      "中雨",
      "大雨",
      "暴雨",
      "大暴雨",
      "特大暴雨",
      "强阵雨",
      "强雷阵雨",
      "极端降雨",
      "毛毛雨/细雨",
      "毛毛雨",
      "细雨",
      "雨",
      "小雨-中雨",
      "中雨-大雨",
      "大雨-暴雨",
      "暴雨-大暴雨",
      "大暴雨-特大暴雨",
      "冻雨",
      "小到中雨",
      "中到大雨",
      "大到暴雨",
      "暴雨到大暴雨",
      "大暴雨到特大暴雨",
      "雷雨",
      "雷电",
      "局部阵雨",
      "零星小雨",
      "阵雨转晴",
    ],
    "icon-yintian": ["阴", "阴天", "中度霾", "重度霾", "严重霾", "未知", "阴转多云", "多云转阴"],
  },
};

export default constant;
