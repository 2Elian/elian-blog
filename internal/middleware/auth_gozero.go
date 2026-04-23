package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"elian-blog/internal/svc"
	"elian-blog/internal/utils"
)

type contextKey string

const (
	UserIDKey   contextKey = "user_id"
	UsernameKey contextKey = "username"
	RoleKey     contextKey = "role"
)

// VeResponse ve-admin-element 响应格式
type VeResponse struct {
	Flag    int         `json:"flag"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
	TraceID string      `json:"trace_id"`
}

func veError(code int, msg string) VeResponse {
	return VeResponse{Flag: 0, Code: code, Data: nil, Msg: msg, TraceID: ""}
}

func JWTAuthMiddleware(svcCtx *svc.ServiceContext) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if auth == "" {
				fmt.Printf("[JWT] No Authorization header for %s\n", r.URL.Path)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"flag":0,"code":401,"data":null,"msg":"未登录","trace_id":""}`))
				return
			}

			tokenString := strings.TrimPrefix(auth, "Bearer ")
			claims, err := utils.ParseToken(tokenString, svcCtx.Config.JWT.Secret)
			if err != nil {
				fmt.Printf("[JWT] Token parse failed for %s: %v (token length=%d)\n", r.URL.Path, err, len(tokenString))
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"flag":0,"code":401,"data":null,"msg":"token无效","trace_id":""}`))
				return
			}

			fmt.Printf("[JWT] OK: user=%s role=%s path=%s\n", claims.Username, claims.Role, r.URL.Path)

			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
			ctx = context.WithValue(ctx, UsernameKey, claims.Username)
			ctx = context.WithValue(ctx, RoleKey, claims.Role)

			next(w, r.WithContext(ctx))
		}
	}
}

func RBACAuthMiddleware(svcCtx *svc.ServiceContext) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			role := r.Context().Value(RoleKey)
			if role == nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"flag":0,"code":403,"data":null,"msg":"无权限","trace_id":""}`))
				return
			}

			roleStr, ok := role.(string)
			if !ok || (roleStr != "admin" && roleStr != "editor") {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"flag":0,"code":403,"data":null,"msg":"权限不足","trace_id":""}`))
				return
			}

			next(w, r)
		}
	}
}

func GetUserID(ctx context.Context) uint {
	id := ctx.Value(UserIDKey)
	if id == nil {
		return 0
	}
	switch v := id.(type) {
	case uint:
		return v
	case int:
		return uint(v)
	case float64:
		return uint(v)
	}
	return 0
}

func GetUsername(ctx context.Context) string {
	name := ctx.Value(UsernameKey)
	if name == nil {
		return ""
	}
	return name.(string)
}

func GetRole(ctx context.Context) string {
	role := ctx.Value(RoleKey)
	if role == nil {
		return ""
	}
	return role.(string)
}
