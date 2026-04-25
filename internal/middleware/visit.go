package middleware

import (
	"fmt"
	"net/http"
	"time"

	"elian-blog/internal/svc"
)

func VisitTrackMiddleware(svcCtx *svc.ServiceContext) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			now := time.Now()
			dateKey := now.Format("2006-01-02")

			// Increment PV
			svcCtx.RDB.Incr(ctx, "pv:"+dateKey)

			// Track UV by IP
			ip := r.RemoteAddr
			if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
				ip = forwarded
			}
			if ip != "" {
				uvKey := fmt.Sprintf("uv_set:%s", dateKey)
				svcCtx.RDB.SAdd(ctx, uvKey, ip)
				svcCtx.RDB.Expire(ctx, uvKey, 48*time.Hour)

				// Cache UV count for quick lookup
				count, _ := svcCtx.RDB.SCard(ctx, uvKey).Result()
				svcCtx.RDB.Set(ctx, "uv:"+dateKey, count, 48*time.Hour)
			}

			next(w, r)
		}
	}
}
