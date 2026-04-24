package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"elian-blog/internal/logic/admin"
	"elian-blog/internal/middleware"
	"elian-blog/internal/model"
	"elian-blog/internal/svc"
	"elian-blog/internal/types"
	"elian-blog/internal/utils"

	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

// parseJSON 解析 JSON body
func parseJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

// getPathID 从 URL path 获取 ID
func getPathID(r *http.Request, key string) (uint, bool) {
	val := r.PathValue(key)
	if val == "" {
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

func getUserIDFromRequest(r *http.Request) uint {
	return middleware.GetUserID(r.Context())
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
		data, err := admin.NewArticleLogic(svcCtx).Create(r.Context(), &req, getUserIDFromRequest(r))
		if err != nil {
			veInternalError(w, "创建文章失败: "+err.Error())
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
		var req types.UpdateArticleTopReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		// 直接更新 is_top 字段，支持 0（取消置顶）
		if err := svcCtx.DB.Table("article").Where("id = ?", req.ID).Update("is_top", req.IsTop).Error; err != nil {
			veInternalError(w, "更新置顶失败")
			return
		}
		veOk(w, nil)
	}
}

func VeUpdateArticleDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateArticleDeleteReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		if req.IsDelete == 1 {
			err := admin.NewArticleLogic(svcCtx).Delete(r.Context(), req.ID)
			if err != nil {
				veInternalError(w, "操作失败")
				return
			}
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
		var req types.QueryCategoryReq
		_ = parseJSON(r, &req)
		if req.Page == 0 {
			req.Page = 1
		}
		if req.PageSize == 0 {
			req.PageSize = 100
		}
		list, total, err := admin.NewCategoryLogic(svcCtx).List(r.Context(), &req)
		if err != nil {
			veInternalError(w, "获取分类失败")
			return
		}
		veOkPage(w, list, total, req.Page, req.PageSize)
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
		var req types.QueryTagReq
		_ = parseJSON(r, &req)
		if req.Page == 0 {
			req.Page = 1
		}
		if req.PageSize == 0 {
			req.PageSize = 100
		}
		list, total, err := admin.NewTagLogic(svcCtx).List(r.Context(), &req)
		if err != nil {
			veInternalError(w, "获取标签失败")
			return
		}
		veOkPage(w, list, total, req.Page, req.PageSize)
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
		var req types.QueryFriendReq
		_ = parseJSON(r, &req)
		if req.Page == 0 {
			req.Page = 1
		}
		if req.PageSize == 0 {
			req.PageSize = 100
		}
		list, total, err := admin.NewFriendLogic(svcCtx).List(r.Context(), &req)
		if err != nil {
			veInternalError(w, "获取友链失败")
			return
		}
		veOkPage(w, list, total, req.Page, req.PageSize)
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
		var req types.QueryMessageReq
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
		var req types.UpdateMessageStatusReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		err := admin.NewMessageLogic(svcCtx).UpdateStatus(r.Context(), &req)
		if err != nil {
			veInternalError(w, "更新状态失败")
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
		var req types.QueryPageReq
		_ = parseJSON(r, &req)
		if req.Page == 0 {
			req.Page = 1
		}
		if req.PageSize == 0 {
			req.PageSize = 100
		}
		list, total, err := admin.NewPageLogic(svcCtx).List(r.Context(), &req)
		if err != nil {
			veInternalError(w, "获取页面失败")
			return
		}
		veOkPage(w, list, total, req.Page, req.PageSize)
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
		var req types.QueryRoleReq
		_ = parseJSON(r, &req)
		if req.Page == 0 {
			req.Page = 1
		}
		if req.PageSize == 0 {
			req.PageSize = 100
		}
		list, total, err := admin.NewRoleLogic(svcCtx).List(r.Context(), &req)
		if err != nil {
			veInternalError(w, "获取角色失败")
			return
		}
		veOkPage(w, list, total, req.Page, req.PageSize)
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
		var req types.UpdateRoleApisReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
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
		veOk(w, map[string]interface{}{
			"list": data,
		})
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

func VeCleanMenuListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
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
			UserID uint `json:"user_id"`
			Status int  `json:"status"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		err := svcCtx.DB.Exec("UPDATE user SET status = ? WHERE id = ?", req.Status, req.UserID).Error
		if err != nil {
			veInternalError(w, "更新状态失败")
			return
		}
		veOk(w, nil)
	}
}

func VeUpdateAccountRolesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			UserID   json.Number `json:"user_id"`
			RoleIDs  []uint      `json:"role_ids"`
			Nickname string      `json:"nickname"`
			Email    string      `json:"email"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		userID, err := req.UserID.Int64()
		if err != nil || userID <= 0 {
			veBadRequest(w, "参数错误")
			return
		}
		uid := uint(userID)

		// Update user profile fields
		updates := map[string]interface{}{}
		if req.Nickname != "" {
			updates["nickname"] = req.Nickname
		}
		if req.Email != "" {
			updates["email"] = req.Email
		}
		if len(updates) > 0 {
			svcCtx.DB.Table("user").Where("id = ?", uid).Updates(updates)
		}

		// Update roles
		svcCtx.DB.Exec("DELETE FROM user_roles WHERE user_id = ?", uid)
		for _, roleID := range req.RoleIDs {
			svcCtx.DB.Exec("INSERT INTO user_roles (user_id, role_id) VALUES (?, ?)", uid, roleID)
		}
		veOk(w, nil)
	}
}

func VeUpdateAccountPasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			UserID   uint   `json:"user_id"`
			Password string `json:"password"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		hashed, err := utils.HashPassword(req.Password)
		if err != nil {
			veInternalError(w, "加密失败")
			return
		}
		err = svcCtx.DB.Exec("UPDATE user SET password = ? WHERE id = ?", hashed, req.UserID).Error
		if err != nil {
			veInternalError(w, "更新密码失败")
			return
		}
		veOk(w, nil)
	}
}

func VeFindAccountOnlineListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		keys, err := svcCtx.RDB.Keys(ctx, "online:*").Result()
		if err != nil || len(keys) == 0 {
			veOkPage(w, []interface{}{}, 0, 1, 10)
			return
		}

		type OnlineUser struct {
			UserID    uint   `json:"user_id"`
			Username  string `json:"username"`
			Nickname  string `json:"nickname"`
			Avatar    string `json:"avatar"`
			LoginTime string `json:"login_time"`
		}

		list := make([]OnlineUser, 0, len(keys))
		for _, key := range keys {
			// Extract user ID from key "online:123"
			var uid uint
			fmt.Sscanf(key, "online:%d", &uid)
			if uid == 0 {
				continue
			}
			user, err := svcCtx.UserDao.GetByID(uid)
			if err != nil {
				continue
			}
			ts, _ := svcCtx.RDB.Get(ctx, key).Result()
			avatar := user.Avatar
			if avatar == "" || avatar == "https://example.com/avatar.png" || (!strings.HasPrefix(avatar, "http") && !strings.HasPrefix(avatar, "/")) {
				avatar = "https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
			}
			list = append(list, OnlineUser{
				UserID:    user.ID,
				Username:  user.Username,
				Nickname:  user.Nickname,
				Avatar:    avatar,
				LoginTime: ts,
			})
		}
		veOkPage(w, list, int64(len(list)), 1, 10)
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
		var req json.RawMessage
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		var data interface{}
		if err := json.Unmarshal(req, &data); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		err := admin.NewSiteLogic(svcCtx).SetFullConfig(r.Context(), data)
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
			"today_uv_count": 0,
			"total_uv_count": 0,
			"uv_growth_rate": 0,
			"today_pv_count": 0,
			"total_pv_count": 0,
			"pv_growth_rate": 0,
		})
	}
}

func VeGetVisitTrendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOk(w, map[string]interface{}{
			"visit_trend": []interface{}{},
		})
	}
}

func VeGetAboutMeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cfg, err := svcCtx.SiteDao.GetByKey("about_me")
		content := ""
		if err == nil && cfg != nil {
			content = cfg.Value
		}
		veOk(w, map[string]interface{}{"content": content})
	}
}

func VeUpdateAboutMeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Content string `json:"content"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		if err := svcCtx.SiteDao.Set("about_me", req.Content); err != nil {
			veInternalError(w, "保存失败")
			return
		}
		veOk(w, nil)
	}
}

func VeGetSystemStateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Disk stats
		diskStats := getDiskStats()

		// Memory stats
		memStats := getMemStats()

		veOk(w, map[string]interface{}{
			"os": map[string]interface{}{
				"goos":         runtime.GOOS,
				"goVersion":    runtime.Version(),
				"numCpu":       runtime.NumCPU(),
				"compiler":     "gc",
				"numGoroutine": runtime.NumGoroutine(),
			},
			"disk": diskStats,
			"mem":  memStats,
		})
	}
}

func VeGetUserAreaStatsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOk(w, map[string]interface{}{
			"user_areas":    []interface{}{},
			"tourist_areas": []interface{}{},
		})
	}
}

// --- 文件上传 ---

func VeUploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, header, err := r.FormFile("file")
		if err != nil {
			veBadRequest(w, "请选择文件")
			return
		}
		defer file.Close()

		// Determine subdirectory from form param
		subDir := r.FormValue("path")
		if subDir == "" {
			subDir = "misc"
		}

		// Create uploads/subDir/YYYY/MM directory structure
		now := time.Now()
		uploadDir := filepath.Join("uploads", subDir, now.Format("2006"), now.Format("01"))
		os.MkdirAll(uploadDir, 0755)

		// Generate unique filename with timestamp
		ext := filepath.Ext(header.Filename)
		filename := fmt.Sprintf("%d%s", now.UnixNano(), ext)
		filePath := filepath.Join(uploadDir, filename)

		dst, err := os.Create(filePath)
		if err != nil {
			veInternalError(w, "保存文件失败")
			return
		}
		defer dst.Close()
		io.Copy(dst, file)

		// Use forward slashes for URL
		urlPath := "/" + filepath.ToSlash(filePath)

		veOk(w, map[string]interface{}{
			"file_url":  urlPath,
			"file_path": filepath.ToSlash(filePath),
			"file_name": header.Filename,
		})
	}
}

func VeMultiUploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(32 << 20)
		subDir := r.FormValue("path")
		if subDir == "" {
			subDir = "misc"
		}

		files := r.MultipartForm.File["files"]
		results := make([]interface{}, 0, len(files))

		for _, fh := range files {
			f, err := fh.Open()
			if err != nil {
				continue
			}
			now := time.Now()
			uploadDir := filepath.Join("uploads", subDir, now.Format("2006"), now.Format("01"))
			os.MkdirAll(uploadDir, 0755)

			ext := filepath.Ext(fh.Filename)
			filename := fmt.Sprintf("%d%s", now.UnixNano(), ext)
			filePath := filepath.Join(uploadDir, filename)

			dst, err := os.Create(filePath)
			if err != nil {
				f.Close()
				continue
			}
			io.Copy(dst, f)
			dst.Close()
			f.Close()

			urlPath := "/" + filepath.ToSlash(filePath)
			results = append(results, map[string]interface{}{
				"file_url":  urlPath,
				"file_path": filepath.ToSlash(filePath),
				"file_name": fh.Filename,
			})
		}
		veOk(w, results)
	}
}

func VeListUploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListUploadFileReq
		if err := parseJSON(r, &req); err != nil {
			req.Page = 1
			req.Limit = 20
		}
		if req.Page <= 0 {
			req.Page = 1
		}
		if req.Limit <= 0 {
			req.Limit = 20
		}

		type FileInfo struct {
			ID        uint   `json:"id"`
			FilePath  string `json:"file_path"`
			FileName  string `json:"file_name"`
			FileType  string `json:"file_type"`
			FileUrl   string `json:"file_url"`
			FileSize  int64  `json:"file_size"`
			UpdatedAt int64  `json:"updated_at"`
		}

		var files []FileInfo
		id := uint(1)
		basePath := "uploads"
		if req.FilePath != "" {
			basePath = req.FilePath
		}

		filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
			if err != nil || path == basePath {
				return nil
			}
			relPath := filepath.ToSlash(path)
			ext := strings.ToLower(filepath.Ext(info.Name()))
			fileType := ""
			if ext != "" {
				switch ext {
				case ".jpg", ".jpeg", ".png", ".gif", ".webp", ".bmp":
					fileType = "image"
				case ".mp4", ".avi", ".mov", ".mkv":
					fileType = "video"
				case ".mp3", ".wav", ".flac":
					fileType = "audio"
				case ".pdf", ".doc", ".docx", ".xls", ".xlsx":
					fileType = "document"
				default:
					fileType = ext[1:]
				}
			}
			files = append(files, FileInfo{
				ID:        id,
				FilePath:  relPath,
				FileName:  info.Name(),
				FileType:  fileType,
				FileUrl:   "/" + relPath,
				FileSize:  info.Size(),
				UpdatedAt: info.ModTime().Unix(),
			})
			id++
			return nil
		})

		total := int64(len(files))
		start := (req.Page - 1) * req.Limit
		if start >= int(total) {
			veOkPage(w, []interface{}{}, total, req.Page, req.Limit)
			return
		}
		end := start + req.Limit
		if end > int(total) {
			end = int(total)
		}
		veOkPage(w, files[start:end], total, req.Page, req.Limit)
	}
}

func VeDeletesUploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			FilePaths []string `json:"file_paths"`
		}
		if err := parseJSON(r, &req); err != nil || len(req.FilePaths) == 0 {
			veBadRequest(w, "参数错误")
			return
		}
		for _, fp := range req.FilePaths {
			if !filepath.IsLocal(fp) {
				continue
			}
			fullPath := filepath.Join(".", fp)
			os.Remove(fullPath)
		}
		veOk(w, nil)
	}
}

// --- 通知 ---

func VeFindUserNoticeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOkPage(w, []interface{}{}, 0, 1, 10)
	}
}

// --- 用户个人中心 ---

func VeUpdateUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := getUserIDFromRequest(r)
		if userID == 0 {
			veUnauthorized(w, "未登录")
			return
		}
		var req struct {
			Nickname string `json:"nickname"`
			Email    string `json:"email"`
			Intro    string `json:"intro"`
			Website  string `json:"website"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		updates := map[string]interface{}{}
		if req.Nickname != "" {
			updates["nickname"] = req.Nickname
		}
		if req.Email != "" {
			updates["email"] = req.Email
		}
		if req.Intro != "" {
			updates["intro"] = req.Intro
		}
		if req.Website != "" {
			updates["website"] = req.Website
		}
		if len(updates) > 0 {
			svcCtx.DB.Table("user").Where("id = ?", userID).Updates(updates)
		}
		veOk(w, nil)
	}
}

func VeUpdateUserAvatarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := getUserIDFromRequest(r)
		if userID == 0 {
			veUnauthorized(w, "未登录")
			return
		}
		var req struct {
			Avatar string `json:"avatar"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		svcCtx.DB.Table("user").Where("id = ?", userID).Update("avatar", req.Avatar)
		veOk(w, nil)
	}
}

func VeUpdateUserPasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := getUserIDFromRequest(r)
		if userID == 0 {
			veUnauthorized(w, "未登录")
			return
		}
		var req struct {
			OldPassword string `json:"old_password"`
			NewPassword string `json:"new_password"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		// Verify old password
		user, err := svcCtx.UserDao.GetByID(userID)
		if err != nil {
			veInternalError(w, "用户不存在")
			return
		}
		if !utils.CheckPassword(req.OldPassword, user.Password) {
			veFail(w, 400, "旧密码错误")
			return
		}
		hashed, err := utils.HashPassword(req.NewPassword)
		if err != nil {
			veInternalError(w, "加密失败")
			return
		}
		svcCtx.DB.Table("user").Where("id = ?", userID).Update("password", hashed)
		veOk(w, nil)
	}
}

// --- 相册 ---

func VeFindAlbumListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryAlbumReq
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
		list, total, err := admin.NewAlbumLogic(svcCtx).List(r.Context(), &req)
		if err != nil {
			veInternalError(w, "获取相册失败")
			return
		}
		veOkPage(w, list, total, req.Page, req.PageSize)
	}
}

func VeGetAlbumHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ID uint `json:"id"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		data, err := admin.NewAlbumLogic(svcCtx).Get(r.Context(), req.ID)
		if err != nil {
			veNotFound(w, "相册不存在")
			return
		}
		veOk(w, data)
	}
}

func VeAddAlbumHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NewAlbumReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		data, err := admin.NewAlbumLogic(svcCtx).Create(r.Context(), &req)
		if err != nil {
			veInternalError(w, "创建相册失败")
			return
		}
		veOk(w, data)
	}
}

func VeUpdateAlbumHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NewAlbumReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		err := admin.NewAlbumLogic(svcCtx).Update(r.Context(), &req)
		if err != nil {
			veInternalError(w, "更新相册失败")
			return
		}
		veOk(w, nil)
	}
}

func VeDeletesAlbumHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			IDs []uint `json:"ids"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		admin.NewAlbumLogic(svcCtx).Delete(r.Context(), req.IDs)
		veOk(w, nil)
	}
}

func VeUpdateAlbumDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateAlbumDeleteReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		err := admin.NewAlbumLogic(svcCtx).UpdateDeleteStatus(r.Context(), &req)
		if err != nil {
			veInternalError(w, "操作失败")
			return
		}
		veOk(w, nil)
	}
}

// --- 照片 ---

func VeFindPhotoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryPhotoReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		if req.Page == 0 {
			req.Page = 1
		}
		if req.PageSize == 0 {
			req.PageSize = 20
		}
		list, total, err := admin.NewPhotoLogic(svcCtx).List(r.Context(), &req)
		if err != nil {
			veInternalError(w, "获取照片失败")
			return
		}
		veOkPage(w, list, total, req.Page, req.PageSize)
	}
}

func VeAddPhotoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NewPhotoReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		data, err := admin.NewPhotoLogic(svcCtx).Create(r.Context(), &req)
		if err != nil {
			veInternalError(w, "添加照片失败")
			return
		}
		veOk(w, data)
	}
}

func VeUpdatePhotoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NewPhotoReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		err := admin.NewPhotoLogic(svcCtx).Update(r.Context(), &req)
		if err != nil {
			veInternalError(w, "更新照片失败")
			return
		}
		veOk(w, nil)
	}
}

func VeDeletesPhotoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			IDs []uint `json:"ids"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		admin.NewPhotoLogic(svcCtx).Delete(r.Context(), req.IDs)
		veOk(w, nil)
	}
}

func VeUpdatePhotoDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdatePhotoDeleteReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		err := admin.NewPhotoLogic(svcCtx).UpdateDeleteStatus(r.Context(), &req)
		if err != nil {
			veInternalError(w, "操作失败")
			return
		}
		veOk(w, nil)
	}
}

// --- 产品管理 ---

func VeFindProductListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageQuery
		if err := parseJSON(r, &req); err != nil {
			req.Page = 1
			req.PageSize = 10
		}
		if req.Page <= 0 {
			req.Page = 1
		}
		if req.PageSize <= 0 {
			req.PageSize = 10
		}
		products, total, err := svcCtx.ProductDao.List(req.Page, req.PageSize)
		if err != nil {
			veInternalError(w, "获取产品列表失败")
			return
		}
		voList := make([]types.ProductVO, 0, len(products))
		for _, p := range products {
			cover := p.Cover
			if cover != "" && !strings.HasPrefix(cover, "http") {
				if !strings.HasPrefix(cover, "/") {
					cover = "/" + cover
				}
				cover = "http://localhost:8080" + cover
			}
			voList = append(voList, types.ProductVO{
				ID:          p.ID,
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
				Cover:       cover,
				Status:      p.Status,
				Sort:        p.Sort,
				Type:        p.Type,
				Link:        p.Link,
				CreatedAt:   fmt.Sprintf("%d", p.CreatedAt.Unix()),
				UpdatedAt:   fmt.Sprintf("%d", p.UpdatedAt.Unix()),
			})
		}
		veOkPage(w, voList, total, req.Page, req.PageSize)
	}
}

func VeAddProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateProductReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		product := &model.Product{
			Name:        req.Name,
			Description: req.Description,
			Price:       req.Price,
			Cover:       req.Cover,
			Status:      req.Status,
			Sort:        req.Sort,
			Type:        req.Type,
			Link:        req.Link,
		}
		if product.Status == 0 {
			product.Status = 1
		}
		if err := svcCtx.ProductDao.Create(product); err != nil {
			veInternalError(w, "创建产品失败")
			return
		}
		veOk(w, product)
	}
}

func VeUpdateProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateProductReq
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		product, err := svcCtx.ProductDao.GetByID(req.ID)
		if err != nil {
			veNotFound(w, "产品不存在")
			return
		}
		if req.Name != "" {
			product.Name = req.Name
		}
		if req.Description != "" {
			product.Description = req.Description
		}
		if req.Price != 0 {
			product.Price = req.Price
		}
		if req.Cover != "" {
			product.Cover = req.Cover
		}
		if req.Status != 0 {
			product.Status = req.Status
		}
		if req.Sort != 0 {
			product.Sort = req.Sort
		}
		if req.Type != 0 {
			product.Type = req.Type
		}
		if req.Link != "" {
			product.Link = req.Link
		}
		if err := svcCtx.ProductDao.Update(product); err != nil {
			veInternalError(w, "更新产品失败")
			return
		}
		veOk(w, nil)
	}
}

func VeDeletesProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Ids []int `json:"ids"`
		}
		if err := parseJSON(r, &req); err != nil {
			veBadRequest(w, "参数错误")
			return
		}
		for _, id := range req.Ids {
			svcCtx.ProductDao.Delete(uint(id))
		}
		veOk(w, nil)
	}
}

// ensure types import is used
var _ = types.IDReq{}
var _ strconv.NumError

func getDiskStats() map[string]interface{} {
	stat, err := disk.Usage(".")
	if err != nil {
		return map[string]interface{}{"totalMb": 0, "usedMb": 0, "freeMb": 0}
	}
	return map[string]interface{}{
		"totalMb":  stat.Total / 1024 / 1024,
		"usedMb":   stat.Used / 1024 / 1024,
		"freeMb":   stat.Free / 1024 / 1024,
		"usedPct":  stat.UsedPercent,
	}
}

func getMemStats() map[string]interface{} {
	stat, err := mem.VirtualMemory()
	if err != nil {
		return map[string]interface{}{"totalMb": 0, "usedMb": 0, "freeMb": 0}
	}
	return map[string]interface{}{
		"totalMb":  stat.Total / 1024 / 1024,
		"usedMb":   stat.Used / 1024 / 1024,
		"freeMb":   stat.Free / 1024 / 1024,
		"usedPct":  stat.UsedPercent,
	}
}
