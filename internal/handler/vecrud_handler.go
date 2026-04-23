package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"elian-blog/internal/logic/admin"
	"elian-blog/internal/svc"
	"elian-blog/internal/types"
)

// parseJSON 解析 JSON body
func parseJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

// getPathID 从 URL path 获取 ID
func getPathID(r *http.Request, key string) (uint, bool) {
	val := r.PathValue(key)
	if val == "" {
		// go-zero path param style
		val = r.URL.Query().Get(key)
	}
	if val == "" {
		return 0, false
	}
	n, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0, false
	}
	return uint(n), true
}

// --- 文章 ---

func VeFindArticleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryArticleHomeReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		if req.Page == 0 {
			req.Page = 1
		}
		if req.PageSize == 0 {
			req.PageSize = 10
		}
		list, total, err := admin.NewArticleLogic(svcCtx).List(r.Context(), &req)
		if err != nil {
			veInternalError(w, "获取文章失败")
			return
		}
		veOkPage(w, list, total, req.Page, req.PageSize)
	}
}

func VeGetArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ID uint `json:"id"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		data, err := admin.NewArticleLogic(svcCtx).Get(r.Context(), req.ID)
		if err != nil {
			veNotFound(w, "文章不存在")
			return
		}
		veOk(w, data)
	}
}

func VeAddArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateArticleReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		data, err := admin.NewArticleLogic(svcCtx).Create(r.Context(), &req)
		if err != nil {
			veInternalError(w, "创建文章失败")
			return
		}
		veOk(w, data)
	}
}

func VeUpdateArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateArticleReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		err := admin.NewArticleLogic(svcCtx).Update(r.Context(), &req)
		if err != nil {
			veInternalError(w, "更新文章失败")
			return
		}
		veOk(w, nil)
	}
}

func VeDeleteArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ID uint `json:"id"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		err := admin.NewArticleLogic(svcCtx).Delete(r.Context(), req.ID)
		if err != nil {
			veInternalError(w, "删除失败")
			return
		}
		veOk(w, nil)
	}
}

func VeUpdateArticleTopHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ID    uint `json:"id"`
			IsTop int  `json:"is_top"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		err := admin.NewArticleLogic(svcCtx).Update(r.Context(), &types.UpdateArticleReq{
			ID:    req.ID,
			IsTop: req.IsTop,
		})
		if err != nil {
			veInternalError(w, "更新置顶失败")
			return
		}
		veOk(w, nil)
	}
}

func VeUpdateArticleDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ID       uint `json:"id"`
			IsDelete int  `json:"is_delete"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		err := admin.NewArticleLogic(svcCtx).Delete(r.Context(), req.ID)
		if err != nil {
			veInternalError(w, "操作失败")
			return
		}
		veOk(w, nil)
	}
}

func VeExportArticleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOk(w, nil)
	}
}

// --- 分类 ---

func VeFindCategoryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := admin.NewCategoryLogic(svcCtx).List(r.Context())
		if err != nil {
			veInternalError(w, "获取分类失败")
			return
		}
		veOk(w, data)
	}
}

func VeAddCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateCategoryReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		data, err := admin.NewCategoryLogic(svcCtx).Create(r.Context(), &req)
		if err != nil {
			veInternalError(w, "创建分类失败")
			return
		}
		veOk(w, data)
	}
}

func VeUpdateCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateCategoryReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		err := admin.NewCategoryLogic(svcCtx).Update(r.Context(), &req)
		if err != nil {
			veInternalError(w, "更新分类失败")
			return
		}
		veOk(w, nil)
	}
}

func VeDeletesCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			IDs []uint `json:"ids"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		for _, id := range req.IDs {
			_ = admin.NewCategoryLogic(svcCtx).Delete(r.Context(), id)
		}
		veOk(w, nil)
	}
}

// --- 标签 ---

func VeFindTagListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := admin.NewTagLogic(svcCtx).List(r.Context())
		if err != nil {
			veInternalError(w, "获取标签失败")
			return
		}
		veOk(w, data)
	}
}

func VeAddTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateTagReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		data, err := admin.NewTagLogic(svcCtx).Create(r.Context(), &req)
		if err != nil {
			veInternalError(w, "创建标签失败")
			return
		}
		veOk(w, data)
	}
}

func VeUpdateTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateTagReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		err := admin.NewTagLogic(svcCtx).Update(r.Context(), &req)
		if err != nil {
			veInternalError(w, "更新标签失败")
			return
		}
		veOk(w, nil)
	}
}

func VeDeletesTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			IDs []uint `json:"ids"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		for _, id := range req.IDs {
			_ = admin.NewTagLogic(svcCtx).Delete(r.Context(), id)
		}
		veOk(w, nil)
	}
}

// --- 评论 ---

func VeFindCommentBackListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryCommentReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		if req.Page == 0 {
			req.Page = 1
		}
		if req.PageSize == 0 {
			req.PageSize = 10
		}
		data, err := admin.NewCommentLogic(svcCtx).List(r.Context(), &req)
		if err != nil {
			veInternalError(w, "获取评论失败")
			return
		}
		veOk(w, data)
	}
}

func VeUpdateCommentStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateCommentStatusReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		err := admin.NewCommentLogic(svcCtx).UpdateStatus(r.Context(), &req)
		if err != nil {
			veInternalError(w, "更新状态失败")
			return
		}
		veOk(w, nil)
	}
}

func VeDeletesCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			IDs []uint `json:"ids"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		for _, id := range req.IDs {
			_ = admin.NewCommentLogic(svcCtx).Delete(r.Context(), id)
		}
		veOk(w, nil)
	}
}

// --- 友链 ---

func VeFindFriendListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := admin.NewFriendLogic(svcCtx).List(r.Context())
		if err != nil {
			veInternalError(w, "获取友链失败")
			return
		}
		veOk(w, data)
	}
}

func VeAddFriendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateFriendLinkReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		data, err := admin.NewFriendLogic(svcCtx).Create(r.Context(), &req)
		if err != nil {
			veInternalError(w, "创建友链失败")
			return
		}
		veOk(w, data)
	}
}

func VeUpdateFriendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateFriendLinkReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		err := admin.NewFriendLogic(svcCtx).Update(r.Context(), &req)
		if err != nil {
			veInternalError(w, "更新友链失败")
			return
		}
		veOk(w, nil)
	}
}

func VeDeletesFriendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			IDs []uint `json:"ids"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		for _, id := range req.IDs {
			_ = admin.NewFriendLogic(svcCtx).Delete(r.Context(), id)
		}
		veOk(w, nil)
	}
}

// --- 留言 ---

func VeFindMessageListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageQuery
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		if req.Page == 0 {
			req.Page = 1
		}
		if req.PageSize == 0 {
			req.PageSize = 10
		}
		data, err := admin.NewMessageLogic(svcCtx).List(r.Context(), &req)
		if err != nil {
			veInternalError(w, "获取留言失败")
			return
		}
		veOk(w, data)
	}
}

func VeUpdateMessageStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ID     uint `json:"id"`
			Status int  `json:"status"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		veOk(w, nil)
	}
}

func VeDeletesMessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			IDs []uint `json:"ids"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		for _, id := range req.IDs {
			_ = admin.NewMessageLogic(svcCtx).Delete(r.Context(), id)
		}
		veOk(w, nil)
	}
}

// --- 页面 ---

func VeFindPageListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := admin.NewPageLogic(svcCtx).List(r.Context())
		if err != nil {
			veInternalError(w, "获取页面失败")
			return
		}
		veOk(w, data)
	}
}

func VeAddPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreatePageReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		data, err := admin.NewPageLogic(svcCtx).Create(r.Context(), &req)
		if err != nil {
			veInternalError(w, "创建页面失败")
			return
		}
		veOk(w, data)
	}
}

func VeUpdatePageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdatePageReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		err := admin.NewPageLogic(svcCtx).Update(r.Context(), &req)
		if err != nil {
			veInternalError(w, "更新页面失败")
			return
		}
		veOk(w, nil)
	}
}

func VeDeletePageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ID uint `json:"id"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		err := admin.NewPageLogic(svcCtx).Delete(r.Context(), req.ID)
		if err != nil {
			veInternalError(w, "删除失败")
			return
		}
		veOk(w, nil)
	}
}

// --- 角色 ---

func VeFindRoleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := admin.NewRoleLogic(svcCtx).List(r.Context())
		if err != nil {
			veInternalError(w, "获取角色失败")
			return
		}
		veOk(w, data)
	}
}

func VeAddRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateRoleReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		data, err := admin.NewRoleLogic(svcCtx).Create(r.Context(), &req)
		if err != nil {
			veInternalError(w, "创建角色失败")
			return
		}
		veOk(w, data)
	}
}

func VeUpdateRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateRoleReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		err := admin.NewRoleLogic(svcCtx).Update(r.Context(), &req)
		if err != nil {
			veInternalError(w, "更新角色失败")
			return
		}
		veOk(w, nil)
	}
}

func VeDeletesRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			IDs []uint `json:"ids"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		for _, id := range req.IDs {
			_ = admin.NewRoleLogic(svcCtx).Delete(r.Context(), id)
		}
		veOk(w, nil)
	}
}

func VeFindRoleResourcesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOk(w, map[string]interface{}{"menus": []interface{}{}, "apis": []interface{}{}})
	}
}

func VeUpdateRoleMenusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateRoleMenusReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		err := admin.NewRoleLogic(svcCtx).UpdateMenus(r.Context(), &req)
		if err != nil {
			veInternalError(w, "更新角色菜单失败")
			return
		}
		veOk(w, nil)
	}
}

func VeUpdateRoleApisHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOk(w, nil)
	}
}

// --- 菜单 ---

func VeFindMenuListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := admin.NewMenuLogic(svcCtx).List(r.Context())
		if err != nil {
			veInternalError(w, "获取菜单失败")
			return
		}
		veOk(w, data)
	}
}

func VeAddMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateMenuReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		data, err := admin.NewMenuLogic(svcCtx).Create(r.Context(), &req)
		if err != nil {
			veInternalError(w, "创建菜单失败")
			return
		}
		veOk(w, data)
	}
}

func VeUpdateMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateMenuReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		err := admin.NewMenuLogic(svcCtx).Update(r.Context(), &req)
		if err != nil {
			veInternalError(w, "更新菜单失败")
			return
		}
		veOk(w, nil)
	}
}

func VeDeletesMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			IDs []uint `json:"ids"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		for _, id := range req.IDs {
			_ = admin.NewMenuLogic(svcCtx).Delete(r.Context(), id)
		}
		veOk(w, nil)
	}
}

func VeSyncMenuListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOk(w, nil)
	}
}

// --- 账号管理 ---

func VeFindAccountListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageQuery
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		if req.Page == 0 {
			req.Page = 1
		}
		if req.PageSize == 0 {
			req.PageSize = 10
		}
		data, err := admin.NewUserLogic(svcCtx).List(r.Context(), &req)
		if err != nil {
			veInternalError(w, "获取用户列表失败")
			return
		}
		veOk(w, data)
	}
}

func VeUpdateAccountStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			UserID string `json:"user_id"`
			Status int    `json:"status"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		veOk(w, nil)
	}
}

func VeUpdateAccountRolesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOk(w, nil)
	}
}

func VeUpdateAccountPasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOk(w, nil)
	}
}

func VeFindAccountOnlineListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOk(w, []interface{}{})
	}
}

// --- 首页/网站 ---

func VeGetAdminHomeInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := admin.NewDashboardLogic(svcCtx).GetStats(r.Context())
		if err != nil {
			veInternalError(w, "获取首页信息失败")
			return
		}
		veOk(w, data)
	}
}

func VeGetWebsiteConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := admin.NewSiteLogic(svcCtx).GetConfig(r.Context())
		if err != nil {
			veInternalError(w, "获取配置失败")
			return
		}
		veOk(w, data)
	}
}

func VeUpdateWebsiteConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetSiteConfigReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		err := admin.NewSiteLogic(svcCtx).SetConfig(r.Context(), &req)
		if err != nil {
			veInternalError(w, "保存配置失败")
			return
		}
		veOk(w, nil)
	}
}

func VeGetVisitStatsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOk(w, map[string]interface{}{
			"pv_count": 0, "uv_count": 0, "ip_count": 0,
		})
	}
}

func VeGetVisitTrendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOk(w, []interface{}{})
	}
}

func VeGetAboutMeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOk(w, map[string]interface{}{"content": ""})
	}
}

func VeUpdateAboutMeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOk(w, nil)
	}
}

func VeGetSystemStateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOk(w, map[string]interface{}{})
	}
}

func VeGetUserAreaStatsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOk(w, []interface{}{})
	}
}

// --- 文件上传 ---

func VeUploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOk(w, map[string]interface{}{"file_url": ""})
	}
}

func VeMultiUploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOk(w, []interface{}{})
	}
}

func VeListUploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOk(w, []interface{}{})
	}
}

func VeDeletesUploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOk(w, nil)
	}
}

// --- 用户个人中心 ---

func VeUpdateUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOk(w, nil)
	}
}

func VeUpdateUserAvatarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOk(w, nil)
	}
}

func VeUpdateUserPasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOk(w, nil)
	}
}

// ensure types import is used
var _ = types.IDReq{}
var _ strconv.NumError
