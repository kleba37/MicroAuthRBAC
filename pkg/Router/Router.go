package Router

import (
	"go-test/internal/Routing"
	//"go-test/internal/Routing"
	"net/http"

	"github.com/kleba37/GoServiceContainer/pkg/Container"
)

type Router struct {
	Container *Container.Container
	rw        *http.ResponseWriter
	req       *http.Request
	mapRoute  *map[string]http.HandlerFunc
}

func (r Router) Router(rw http.ResponseWriter, req *http.Request) {
	route := Routing.MakeRouteMap(r.Container)

	handler, ok := route[req.URL.Path]

	if !ok {
		http.Error(rw, "Not Found", http.StatusNotFound)
		return
	}

	handler(rw, req)
}
