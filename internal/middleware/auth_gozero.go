package middleware

import (
	"context"
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

func JWTAuthMiddleware(svcCtx *svc.ServiceContext) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if auth == "" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"code":401,"message":"未登录"}`))
				return
			}

			tokenString := strings.TrimPrefix(auth, "Bearer ")
			claims, err := utils.ParseToken(tokenString, svcCtx.Config.JWT.Secret)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"code":401,"message":"token无效"}`))
				return
			}

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
				w.Write([]byte(`{"code":403,"message":"无权限"}`))
				return
			}

			roleStr, ok := role.(string)
			if !ok || (roleStr != "admin" && roleStr != "editor") {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"code":403,"message":"权限不足"}`))
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