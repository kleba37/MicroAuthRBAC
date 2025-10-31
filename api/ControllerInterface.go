package Controller

import (
	"net/http"
)

type Controller interface {
	Handler(rw http.ResponseWriter, req *http.Request)
}
