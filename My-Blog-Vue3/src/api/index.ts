import article from "./article";
import comment from "./comment";
import collect from "./collect";
import family from "./family";
import labelSort from "./labelSort";
import other from "./other";
import prohibitedWord from "./prohibitedWord";
import resource from "./resource";
import user from "./user";
import webInfo from "./webInfo";
import weiYan from "./weiYan";
import apiWrapper from "./tool/apiWrapper";

const api = {
  ...article,
  ...comment,
  ...collect,
  ...family,
  ...labelSort,
  ...other,
  ...prohibitedWord,
  ...resource,
  ...user,
  ...webInfo,
  ...weiYan,
  apiWrapper,
};

export default api;

// 导出所有 API 模块
export {
  article,
  comment,
  collect,
  family,
  labelSort,
  other,
  prohibitedWord,
  resource,
  user,
  webInfo,
  weiYan,
  apiWrapper,
};
