import axios from "axios";
import constant from "./constant";
//处理url参数
import qs from "qs";
import store from "../store";
let loadingCount = 0;
const addLoading = () => {
  // 路由跳转时loading都是-2
  if (loadingCount < 0) {
    loadingCount = 0;
    return;
  }
  // 开始请求时，添加数量，并且设置为显示loading
  loadingCount++;
  // 如果已有loading就取消添加
  if (store.state.isShowLoading) {
    return;
  }
  store.commit("SET_SHOWLOADING", true);
};
const isCloseLoading = () => {
  // 请求完成或者请求失败时,减少正在请求的数量,并且判断是否还有未完成的请求,如果没有了,则设置为隐藏loading
  loadingCount--;
  loadingCount = Math.max(loadingCount, 0); //做个保护
  if (loadingCount === 0) {
    loadingCount = -2;
    store.commit("SET_SHOWLOADING", false);
  }
};
axios.defaults.baseURL = constant.baseURL;
// 添加请求拦截器
axios.interceptors.request.use(
  function (config) {
    addLoading();
    // 在发送请求之前做些什么
    return config;
  },
  function (error) {
    isCloseLoading();
    // 对请求错误做些什么
    return Promise.reject(error);
  }
);
// 添加响应拦截器
axios.interceptors.response.use(
  function (response) {
    isCloseLoading();
    if (
      response.data !== null &&
      response.data.hasOwnProperty("code") &&
      response.data.code !== 200
    ) {
      if (response.data.code === 300) {
        store.commit("loadCurrentUser", {});
        localStorage.removeItem("userToken");
        store.commit("loadCurrentAdmin", {});
        localStorage.removeItem("adminToken");
        //跳转到用户资料界面
        window.location.href = constant.webURL + "/user";
      }
      return Promise.reject(new Error(response.data.message));
    } else {
      return response;
    }
  },
  function (error) {
    isCloseLoading();
    // 对响应错误做点什么
    return Promise.reject(error);
  }
);
// 当data为URLSearchParams对象时设置为application/x-www-form-urlencoded;charset=utf-8
// 当data为普通对象时，会被设置为application/json;charset=utf-8
export default {
  post(url, params = {}, isAdmin = false, json = true, authorization = false) {
    let config;
    if (isAdmin) {
      config = {
        headers: {
          Authorization: authorization
            ? "Token " + localStorage.getItem("adminToken")
            : localStorage.getItem("adminToken"),
        },
      };
    } else {
      config = {
        headers: {
          Authorization: !Boolean(localStorage.getItem("userToken"))
            ? localStorage.getItem("userToken")
            : "Token " + localStorage.getItem("userToken"),
        },
      };
    }
    return new Promise((resolve, reject) => {
      axios
        .post(url, json ? params : qs.stringify(params), config)
        .then((res) => {
          resolve(res.data);
        })
        .catch((err) => {
          reject(err);
        });
    });
  },
  get(url, params = {}, isAdmin = false, authorization = false) {
    let headers;
    if (isAdmin) {
      headers = {
        Authorization: authorization
          ? "Token " + localStorage.getItem("adminToken")
          : localStorage.getItem("adminToken"),
      };
    } else {
      headers = {
        Authorization: authorization
          ? "Token " + localStorage.getItem("userToken")
          : localStorage.getItem("userToken"),
      };
    }
    return new Promise((resolve, reject) => {
      axios
        .get(url, {
          params: params,
          headers: headers,
        })
        .then((res) => {
          resolve(res.data);
        })
        .catch((err) => {
          reject(err);
        });
    });
  },
  uploadQiniu(url, param) {
    let config = {
      headers: { "Content-Type": "multipart/form-data" },
      timeout: 60000,
    };
    return new Promise((resolve, reject) => {
      axios
        .post(url, param, config)
        .then((res) => {
          resolve(res.data);
        })
        .catch((err) => {
          reject(err);
        });
    });
  },
};
