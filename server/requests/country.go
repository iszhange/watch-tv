package requests

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type Country struct {
	Code   string `json:"code" validate:"omitempty,alpha"`
	IsShow int    `json:"is_show" validate:"omitempty,oneof=0 1 2"`
}

func (o *Country) Parse(r *http.Request) {
	query := r.URL.Query()
	o.Code = query.Get("code")
	o.IsShow, _ = strconv.Atoi(query.Get("is_show"))
}

func (o *Country) Validate() error {
	return validator.New().Struct(o)
}
