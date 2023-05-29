package internal

import (
	"net/http"
	"watchtv/models"
	"watchtv/requests"
	"watchtv/utils"

	"github.com/go-chi/chi/middleware"
	"gorm.io/gorm"
)

type Countries struct{}

func (o Countries) List(w http.ResponseWriter, r *http.Request) {
	reqID := middleware.GetReqID(r.Context())
	var request requests.Country
	request.Parse(r)
	err := request.Validate()
	if err != nil {
		utils.RespondWithJSON(w, map[string]interface{}{
			"request_id": reqID,
			"code":       400,
			"message":    err.Error(),
		})
		return
	}

	var model models.Country
	db, _ := r.Context().Value(utils.GormContextKey).(*gorm.DB)
	model.SetDB(db)
	total, data := model.GetList(request)

	utils.RespondWithJSON(w, map[string]interface{}{
		"request_id": reqID,
		"code":       200,
		"message":    "成功",
		"data": map[string]interface{}{
			"total": total,
			"list":  data,
		},
	})
}
