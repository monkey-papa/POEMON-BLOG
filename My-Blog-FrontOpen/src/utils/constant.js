export default {
  baseURL: "http://localhost:8000/api",
  webURL: "http://localhost:8000",
  // baseURL: "http://www.monkey-papa.com/api", // 线上地址加上/api
  // webURL: "http://www.monkey-papa.com",  // 线上地址
  //诗词语录
  hitokoto: "https://v1.hitokoto.cn",
  shehui: "https://api.oick.cn/yulu/api.php",
  jinrishici: "https://v1.jinrishici.com/all.json",
  //目录
  tocbot: "https://cdn.bootcdn.net/ajax/libs/tocbot/4.21.0/tocbot.min.js",
  //上传图片文件地址 elementUI:action
  qiniuUploadImages: "http://www.monkey-papa.com/api/resource/updateImage/", // 你的后端线上地址加上/api/resource/updateImage/
  //表情地址,md图片地址
  qiniuUploadEntrance: "http://www.monkey-papa.com/", // 你的七牛云存储空间的外链域名
  //视频地址
  favoriteVideo: "https://www.monkey-papa.mp4", // 自己替换一个视频外链
  //标签颜色
  before_color_list: [
    "#ff4b2b",
    "#EF794F",
    "#67C23A",
    "orange",
    "rgb(131, 123, 199)",
    "#23d5ab",
  ],
  //微言颜色
  tree_hole_color: [
    "#ee7752",
    "#e73c7e",
    "#23a6d5",
    "#23d5ab",
    "rgb(131, 123, 199)",
    "#23d5ab",
  ],
  //按钮颜色
  before_color_1: "#74bdf0",
  after_color_1: "linear-gradient(45deg, #f43f3b, #ec008c)",
  before_color_2: "rgb(131, 123, 199)",
  after_color_2: "linear-gradient(45deg, #f43f3b, #ec008c)",
  //分页颜色
  pageColor: "#ee7752",
  commentPageColor: "#23d5ab",
  userId: 9, // 博客主人的用户id也就是boss的id，我这里是9，你们可以替换成自己的
  //emoji含义
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
  //纯色
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
  //渐变
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
  // 图片都替换为自己的
  about: [
    {
      img: "https://www.qiniuyun.jpg",
      tit: "女帝",
      sub: "巾帼女帝，新生红衣😍",
    },
    {
      img: "https://www.qiniuyun.jpg",
      tit: "新不良帅",
      sub: "弃天道，择霸道🥱",
    },
    {
      img: "https://www.qiniuyun.jpg",
      tit: "四大尸祖",
      sub: "搞笑女+搞笑男😅",
    },
    {
      img: "https://www.qiniuyun.png",
      tit: "这俩货",
      sub: "导致我每集片头都不跳过😋",
    },
    {
      img: "https://www.qiniuyun.jpg",
      tit: "奥姑",
      sub: "好腿，只能说被打的真惨🤣",
    },
    {
      img: "https://www.qiniuyun.jpg",
      tit: "随心",
      sub: "(*^▽^*)",
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
};
