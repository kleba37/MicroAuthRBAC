package Controllers

import (
	"encoding/json"
	MessageResponses "go-test/api/Responses"
	"net/http"

	"github.com/kleba37/GoServiceContainer/pkg/Container"
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
