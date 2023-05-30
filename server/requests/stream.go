package requests

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Stream struct {
	Channel string `json:"channel" validate:"omitempty,regexp=^[a-zA-Z0-9.]*$"`
}

func (o *Stream) Parse(r *http.Request) {
	query := r.URL.Query()
	o.Channel = query.Get("channel")
}

func (o *Stream) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("regexp", regexpValidation)

	return validate.Struct(o)
}
