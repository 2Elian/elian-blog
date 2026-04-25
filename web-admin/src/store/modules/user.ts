import { store } from "@/store";
import { AuthStorage } from "@/utils/auth";
import { usePermissionStoreHook } from "@/store/modules/permission";
import { useTagsViewStore } from "@/store";
import type {
  LoginReq,
  LoginResp,
  PhoneLoginReq,
  ThirdLoginReq,
  UserInfoResp,
} from "@/api/types";
import { AuthAPI } from "@/api/auth";
import { UserAPI } from "@/api/user";

export const useUserStore = defineStore("user", () => {
  const userInfo = ref<UserInfoResp>({} as UserInfoResp);
  const rememberMe = ref(AuthStorage.getRememberMe());

  function login(loginData: LoginReq) {
    return new Promise<LoginResp>((resolve, reject) => {
      AuthAPI.loginApi(loginData)
        .then((res) => {
          console.log("[LOGIN] response:", JSON.stringify(res));
          const uid = res.data?.user_id;
          const token = res.data?.token?.access_token;
          console.log("[LOGIN] uid:", uid, "token length:", token?.length);
          AuthStorage.setTokens(uid, token);
          // verify storage
          const stored = AuthStorage.getAccessToken();
          console.log("[LOGIN] stored token length:", stored?.length);
          resolve(res.data);
        })
        .catch((error) => {
          reject(error);
        });
    });
  }

  function phoneLogin(_loginData: PhoneLoginReq): Promise<LoginResp> {
    return Promise.reject(new Error("手机登录暂未开放"));
  }

  function thirdLogin(_loginData: ThirdLoginReq): Promise<LoginResp> {
    return Promise.reject(new Error("第三方登录暂未开放"));
  }

  function getUserInfo() {
    return new Promise<UserInfoResp>((resolve, reject) => {
      UserAPI.getUserInfoApi()
        .then((res) => {
          if (!res) {
            reject("Verification failed, please Login again.");
            return;
          }
          Object.assign(userInfo.value, { ...res.data });
          resolve(res.data);
        })
        .catch((error) => {
          reject(error);
        });
    });
  }

  function logout() {
    return new Promise<void>((resolve, reject) => {
      AuthAPI.logoutApi()
        .then(() => {
          resetAllState();
          resolve();
        })
        .catch((error) => {
          reject(error);
        });
    });
  }

  function resetAllState() {
    resetUserState();
    usePermissionStoreHook().resetRouter();
    useTagsViewStore().delAllViews();
    return Promise.resolve();
  }

  function resetUserState() {
    AuthStorage.clearAuth();
    userInfo.value = {} as any;
  }

  return {
    userInfo,
    rememberMe,
    isLoggedIn: () => !!AuthStorage.getAccessToken(),
    getUserInfo,
    login,
    phoneLogin,
    thirdLogin,
    logout,
    resetAllState,
    resetUserState,
  };
});

export function useUserStoreHook() {
  return useUserStore(store);
}
