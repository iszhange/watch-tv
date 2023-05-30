package requests

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type Channel struct {
	Country string `json:"country" validate:"omitempty,alpha"`
	IsNsfw  int    `json:"is_nsfw" validate:"omitempty,oneof=0 1"`
	IsShow  int    `json:"is_show" validate:"omitempty,oneof=0 1 2"`
}

func (o *Channel) Parse(r *http.Request) {
	query := r.URL.Query()
	o.Country = query.Get("country")
	o.IsNsfw, _ = strconv.Atoi(query.Get("is_nsfw"))
	o.IsShow, _ = strconv.Atoi(query.Get("is_show"))
}

func (o *Channel) Validate() error {
	return validator.New().Struct(o)
}
