package Controllers

import (
	"encoding/json"
	MessageResponses "go-test/api/Responses"
	"go-test/pkg/Container"
	"net/http"
)

type MainController struct {
	Container *Container.Container
}

func (h MainController) Handler(rw http.ResponseWriter, req *http.Request) {
	data := MessageResponses.Message{
		Status:  "ok",
		Message: "Hello World!",
	}

	rw.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(rw).Encode(data)

	if err != nil {
		return
	}
}
