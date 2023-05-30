package router

import (
	"net/http"
	"watchtv/internal"
	mymiddleware "watchtv/middleware"
	"watchtv/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	// 启用中间件
	r.Use(middleware.RequestID)
	r.Use(mymiddleware.Gorm)
	r.Use(middleware.Recoverer)
	r.Use(mymiddleware.Tracing)
	r.Use(middleware.Compress(5, "text/html", "application/json"))
	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		reqID := middleware.GetReqID(r.Context())
		utils.RespondWithJSON(w, map[string]interface{}{
			"request_id": reqID,
			"code":       0,
			"message":    "WatchTV Server.",
		})
	})

	var countries internal.Countries
	r.Get("/countries", countries.List)

	var channels internal.Channels
	r.Get("/channels", channels.List)

	var streams internal.Streams
	r.Get("/streams", streams.List)

	return r
}
