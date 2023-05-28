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
)

func Gorm(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := utils.GetMariaDBInstance()
		timeoutContext, _ := context.WithTimeout(context.Background(), time.Second*30)
		ctx := context.WithValue(r.Context(), utils.GormContextKey("DB"), db.WithContext(timeoutContext))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
