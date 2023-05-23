package main

import (
	"net/http"
	"time"
	"watchtv/router"
	"watchtv/utils"
)

func main() {
	// 设置时区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	time.Local = loc
	cfg := utils.Config.WatchTV

	r := router.NewRouter()
	http.ListenAndServe(cfg.HTTP, r)
}
