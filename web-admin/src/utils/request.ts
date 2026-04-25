import axios, { type AxiosError, type AxiosResponse, type InternalAxiosRequestConfig } from "axios";
import { AuthStorage, redirectToLogin } from "./auth";

const axiosInstance = axios.create({
  baseURL: "",
  timeout: 15000,
  withCredentials: false,
  headers: {
    "Content-Type": "application/json;charset=UTF-8",
  },
});

// 请求拦截器 — 注入 JWT
axiosInstance.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    const accessToken = AuthStorage.getAccessToken();
    console.log("[REQUEST]", config.method?.toUpperCase(), config.url, "token:", accessToken ? `Bearer ${accessToken.substring(0, 20)}...` : "NONE");
    if (accessToken) {
      config.headers.set("Authorization", `Bearer ${accessToken}`);
    }
    return config;
  },
  (error: AxiosError) => {
    return Promise.reject(error);
  }
);

// 响应拦截器
axiosInstance.interceptors.response.use(
  (response: AxiosResponse) => {
    if (response.config.responseType === "blob" || response.config.responseType === "arraybuffer") {
      return response;
    }

    const { code, data, msg } = response.data;

    switch (code) {
      case 200:
        break;
      case 400:
        return Promise.reject(new Error(msg || "请求参数错误"));
      case 401:
        redirectToLogin("登录已过期，请重新登录");
        return Promise.reject(new Error(msg || "用户未登录"));
      case 403:
        return Promise.reject(new Error(msg || "无权限访问"));
      default:
        return Promise.reject(new Error(msg || "系统错误"));
    }
    return response.data;
  },
  (error: AxiosError) => {
    console.error("request error", error);
    let { message } = error;
    if (message == "Network Error") {
      message = "后端接口连接异常";
    } else if (message.includes("timeout")) {
      message = "系统接口请求超时";
    } else if (message.includes("Request failed with status code")) {
      message = "系统接口" + message.substring(message.length - 3) + "异常";
    }
    ElMessage.error(message);
    return Promise.reject(error);
  }
);

export default axiosInstance;
