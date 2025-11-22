package Routing

import (
	"go-test/api/Controllers"
	"net/http"

	"github.com/kleba37/GoServiceContainer/pkg/Container"
)

func MakeRouteMap(container *Container.Container) map[string]http.HandlerFunc {
	return map[string]http.HandlerFunc{
		"/":       Controllers.MainController{Container: container}.Handler,
		"/access": Controllers.AuthOperationController{Container: container}.Handler,
	}
}
