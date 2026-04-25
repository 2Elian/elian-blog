package handler

import (
	"net/http"

	"elian-blog/internal/svc"
)

// veStub returns a generic stub handler for unimplemented endpoints
func veStub(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOk(w, nil)
	}
}

func veStubList(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veOkPage(w, []interface{}{}, 0, 1, 10)
	}
}

// Get all stub route registrations
func VeStubRoutes(svcCtx *svc.ServiceContext) map[string]http.HandlerFunc {
	s := map[string]http.HandlerFunc{
		// Notice management (stub)
		"POST /admin-api/v1/notice/find_notice_list":     veStubList(svcCtx),
		"POST /admin-api/v1/notice/add_notice":           veStub(svcCtx),
		"GET /admin-api/v1/notice/get_notice":            veStub(svcCtx),
		"PUT /admin-api/v1/notice/update_notice":         veStub(svcCtx),
		"PUT /admin-api/v1/notice/update_notice_status":  veStub(svcCtx),

		// API management (stub)
		"POST /admin-api/v1/api/find_api_list":    veStubList(svcCtx),
		"POST /admin-api/v1/api/add_api":          veStub(svcCtx),
		"PUT /admin-api/v1/api/update_api":        veStub(svcCtx),
		"DELETE /admin-api/v1/api/deletes_api":    veStub(svcCtx),
		"POST /admin-api/v1/api/sync_api_list":    veStub(svcCtx),
		"POST /admin-api/v1/api/clean_api_list":   veStub(svcCtx),

		// Album (stub)
		"POST /admin-api/v1/album/find_album_list":    veStubList(svcCtx),
		"POST /admin-api/v1/album/add_album":           veStub(svcCtx),
		"POST /admin-api/v1/album/get_album":           veStub(svcCtx),
		"PUT /admin-api/v1/album/update_album":         veStub(svcCtx),
		"DELETE /admin-api/v1/album/deletes_album":     veStub(svcCtx),
		"PUT /admin-api/v1/album/update_album_delete":  veStub(svcCtx),

		// Photo (stub)
		"POST /admin-api/v1/photo/find_photo_list":    veStubList(svcCtx),
		"POST /admin-api/v1/photo/add_photo":           veStub(svcCtx),
		"PUT /admin-api/v1/photo/update_photo":         veStub(svcCtx),
		"DELETE /admin-api/v1/photo/deletes_photo":     veStub(svcCtx),
		"PUT /admin-api/v1/photo/update_photo_delete":  veStub(svcCtx),

		// Talk (stub)
		"POST /admin-api/v1/talk/find_talk_list":  veStubList(svcCtx),
		"POST /admin-api/v1/talk/get_talk":         veStub(svcCtx),
		"POST /admin-api/v1/talk/add_talk":         veStub(svcCtx),
		"PUT /admin-api/v1/talk/update_talk":       veStub(svcCtx),
		"DELETE /admin-api/v1/talk/delete_talk":    veStub(svcCtx),

		// Logs (stub)
		"POST /admin-api/v1/operation_log/find_operation_log_list":  veStubList(svcCtx),
		"DELETE /admin-api/v1/operation_log/deletes_operation_log":  veStub(svcCtx),
		"POST /admin-api/v1/file_log/find_file_log_list":            veStubList(svcCtx),
		"DELETE /admin-api/v1/file_log/deletes_file_log":            veStub(svcCtx),
		"POST /admin-api/v1/login_log/find_login_log_list":          veStubList(svcCtx),
		"DELETE /admin-api/v1/login_log/deletes_login_log":          veStub(svcCtx),
		"POST /admin-api/v1/visit_log/find_visit_log_list":          veStubList(svcCtx),
		"DELETE /admin-api/v1/visit_log/deletes_visit_log":          veStub(svcCtx),
		"POST /admin-api/v1/visitor/find_visitor_list":              veStubList(svcCtx),

		// Auth extras (stub)
		"GET /admin-api/v1/get_client_info":            veStub(svcCtx),
		"POST /admin-api/v1/refresh_token":             veStub(svcCtx),
		// User login history (stub)
		"POST /admin-api/v1/user/get_user_login_history_list": veStubList(svcCtx),
	}
	return s
}
