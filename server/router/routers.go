package router

import (
	"net/http"
	"watchtv/internal"
	mymiddleware "watchtv/middleware"
	"watchtv/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	// 启用中间件
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
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
