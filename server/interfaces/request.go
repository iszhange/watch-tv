package interfaces

import "net/http"

type Request interface {
	Parse(r *http.Request)
	Validate() error
}
