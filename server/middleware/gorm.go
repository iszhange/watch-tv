/*
设置gorm持续会话
https://gorm.io/zh_CN/docs/context.html
*/
package middleware

import (
	"context"
	"net/http"
	"time"
	"watchtv/utils"

	"github.com/go-chi/chi/v5/middleware"
)

func Gorm(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqID := middleware.GetReqID(ctx)
		db := utils.GetMariaDBInstance()
		timeoutCtx, _ := context.WithTimeout(context.Background(), time.Second*30)
		ctx = context.WithValue(ctx, middleware.RequestIDKey, reqID)
		ctx = context.WithValue(ctx, utils.GormContextKey, db.WithContext(timeoutCtx))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
