package handler

import (
	"net/http"

	"elian-blog/internal/logic/admin"
	"elian-blog/internal/svc"
	"elian-blog/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// --- ve-admin-element 认证接口 ---

// VeLoginHandler 登录 (POST /admin-api/v1/login)
func VeLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			veFail(w, 400, "参数错误")
			return
		}
		data, err := admin.NewVeAuthLogic(svcCtx).VeLogin(r.Context(), &req)
		if err != nil {
			veFail(w, 401, err.Error())
			return
		}
		veOk(w, data)
	}
}

// VeLogoutHandler 登出 (GET /admin-api/v1/logout)
func VeLogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOk(w, nil)
	}
}

// VeGetUserInfoHandler 获取用户信息 (GET /admin-api/v1/user/get_user_info)
func VeGetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := admin.NewVeAuthLogic(svcCtx).VeGetUserInfo(r.Context())
		if err != nil {
			veFail(w, 401, err.Error())
			return
		}
		veOk(w, data)
	}
}

// VeGetUserMenusHandler 获取用户菜单 (GET /admin-api/v1/user/get_user_menus)
func VeGetUserMenusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := admin.NewVeAuthLogic(svcCtx).VeGetUserMenus(r.Context())
		if err != nil {
			veFail(w, 500, err.Error())
			return
		}
		veOk(w, map[string]interface{}{"list": data})
	}
}

// VeGetUserRolesHandler 获取用户角色 (GET /admin-api/v1/user/get_user_roles)
func VeGetUserRolesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := admin.NewVeAuthLogic(svcCtx).VeGetUserRoles(r.Context())
		if err != nil {
			veFail(w, 500, err.Error())
			return
		}
		veOk(w, data)
	}
}

// VeGetUserApisHandler 获取用户API权限 (GET /admin-api/v1/user/get_user_apis)
func VeGetUserApisHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := admin.NewVeAuthLogic(svcCtx).VeGetUserApis(r.Context())
		if err != nil {
			veFail(w, 500, err.Error())
			return
		}
		veOk(w, data)
	}
}
