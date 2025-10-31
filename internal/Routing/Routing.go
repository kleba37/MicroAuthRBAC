package Routing

import (
	"go-test/api/Controllers"
	"go-test/pkg/Container"
	"net/http"
)

func MakeRouteMap(container *Container.Container) map[string]http.HandlerFunc {
	return map[string]http.HandlerFunc{
		"/":       Controllers.MainController{Container: container}.Handler,
		"/access": Controllers.AuthOperationController{Container: container}.Handler,
	}
}
