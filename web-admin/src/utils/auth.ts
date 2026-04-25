import { useUserStoreHook } from "@/store/modules/user";
import router from "@/router";

export const APP_NAME = "admin-web";

export const STORAGE_KEYS = {
  UID: `${APP_NAME}:auth:uid`,
  ACCESS_TOKEN: `${APP_NAME}:auth:access_token`,
  REMEMBER_ME: `${APP_NAME}:auth:remember_me`,
} as const;

export const AuthStorage = {
  getUid() {
    const isRememberMe = localStorage.getItem(STORAGE_KEYS.REMEMBER_ME);
    const storage = isRememberMe ? localStorage : sessionStorage;
    return storage.getItem(STORAGE_KEYS.UID);
  },
  getAccessToken() {
    const isRememberMe = localStorage.getItem(STORAGE_KEYS.REMEMBER_ME);
    const storage = isRememberMe ? localStorage : sessionStorage;
    return storage.getItem(STORAGE_KEYS.ACCESS_TOKEN);
  },
  setTokens(uid: string, accessToken: string) {
    const isRememberMe = localStorage.getItem(STORAGE_KEYS.REMEMBER_ME);
    const storage = isRememberMe ? localStorage : sessionStorage;
    storage.setItem(STORAGE_KEYS.UID, uid);
    storage.setItem(STORAGE_KEYS.ACCESS_TOKEN, accessToken);
  },
  getRememberMe() {
    return localStorage.getItem(STORAGE_KEYS.REMEMBER_ME) === "true";
  },
  setRememberMe(rememberMe: boolean) {
    if (rememberMe) {
      localStorage.setItem(STORAGE_KEYS.REMEMBER_ME, "true");
    } else {
      localStorage.removeItem(STORAGE_KEYS.REMEMBER_ME);
    }
  },
  clearAuth() {
    const keys = [STORAGE_KEYS.UID, STORAGE_KEYS.ACCESS_TOKEN];
    keys.forEach((k) => {
      localStorage.removeItem(k);
      sessionStorage.removeItem(k);
    });
  },
};

/**
 * 权限判断
 */
export function hasPerm(value: string | string[], type: "button" | "role" = "button"): boolean {
  return true;
  // const { roles, perms } = useUserStoreHook().userInfo;
  //
  // if (!roles || !perms) {
  //   return false;
  // }
  //
  // // 超级管理员拥有所有权限
  // if (type === "button" && roles.includes(ROLE_ROOT)) {
  //   return true;
  // }
  //
  // const auths = type === "button" ? perms : roles;
  // return typeof value === "string"
  //   ? auths.includes(value)
  //   : value.some((perm) => auths.includes(perm));
}

/**
 * 重定向到登录页面
 */
export async function redirectToLogin(message: string = "请重新登录"): Promise<void> {
  ElNotification({
    title: "提示",
    message,
    type: "warning",
    duration: 3000,
  });

  await useUserStoreHook().resetAllState();

  try {
    // 跳转到登录页，保留当前路由用于登录后跳转
    const currentPath = router.currentRoute.value.fullPath;
    await router.push(`/login?redirect=${encodeURIComponent(currentPath)}`);
  } catch (error) {
    console.error("Redirect to login error:", error);
    // 强制跳转，即使路由重定向失败
    window.location.href = "/login";
  }
}
