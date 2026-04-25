package handler

import (
	"encoding/json"
	"net/http"

	"elian-blog/internal/types"
)

// VeResponse ve-admin-element 期望的响应格式
type VeResponse struct {
	Flag    int         `json:"flag"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Msg     string      `json:"msg"`
	TraceID string      `json:"trace_id"`
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(v)
}

// veOk 成功响应 (code: 200)
func veOk(w http.ResponseWriter, data interface{}) {
	writeJSON(w, VeResponse{
		Flag:    0,
		Code:    200,
		Data:    data,
		Msg:     "success",
		TraceID: "",
	})
}

// veOkPage 分页成功响应
func veOkPage(w http.ResponseWriter, list interface{}, total int64, page, pageSize int) {
	veOk(w, types.PageResp{List: list, Total: total, Page: page, PageSize: pageSize})
}

// veOkMsg 成功响应带消息
func veOkMsg(w http.ResponseWriter, msg string, data interface{}) {
	writeJSON(w, VeResponse{
		Flag:    0,
		Code:    200,
		Data:    data,
		Msg:     msg,
		TraceID: "",
	})
}

// veFail 错误响应
func veFail(w http.ResponseWriter, code int, msg string) {
	writeJSON(w, VeResponse{
		Flag:    0,
		Code:    code,
		Data:    nil,
		Msg:     msg,
		TraceID: "",
	})
}

// veBadRequest 400错误
func veBadRequest(w http.ResponseWriter, msg string) {
	veFail(w, 400, msg)
}

// veUnauthorized 401错误
func veUnauthorized(w http.ResponseWriter, msg string) {
	veFail(w, 401, msg)
}

// veForbidden 403错误
func veForbidden(w http.ResponseWriter, msg string) {
	veFail(w, 403, msg)
}

// veNotFound 404错误
func veNotFound(w http.ResponseWriter, msg string) {
	veFail(w, 404, msg)
}

// veInternalError 500错误
func veInternalError(w http.ResponseWriter, msg string) {
	veFail(w, 500, msg)
}

// --- 旧响应格式兼容 (blog API) ---
func ok(w http.ResponseWriter, data interface{}) {
	writeJSON(w, types.Body{Code: 0, Message: "success", Data: data})
}

func okPage(w http.ResponseWriter, list interface{}, total int64, page, pageSize int) {
	ok(w, types.PageResp{List: list, Total: total, Page: page, PageSize: pageSize})
}

func fail(w http.ResponseWriter, code int, msg string) {
	writeJSON(w, types.Body{Code: code, Message: msg})
}
