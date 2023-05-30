package internal

import (
	"net/http"
	"watchtv/models"
	"watchtv/requests"
	"watchtv/utils"

	"github.com/go-chi/chi/v5/middleware"
)

type Channels struct{}

func (o Channels) List(w http.ResponseWriter, r *http.Request) {
	reqID := middleware.GetReqID(r.Context())
	var request requests.Channel
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

	var model models.Channel
	model.SetDB(r.Context())
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
