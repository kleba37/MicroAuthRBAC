package Controllers

import (
	"database/sql"
	"encoding/json"
	MessageResponses "go-test/api/Responses"
	"go-test/internal/DTO"
	"go-test/internal/model"
	"net/http"

	"github.com/kleba37/GoServiceContainer/pkg/Container"

	_ "modernc.org/sqlite"
)

type AuthOperationController struct {
	Container *Container.Container
}

func (h AuthOperationController) Handler(rw http.ResponseWriter, req *http.Request) {
	dto := new(DTO.OperationDTO)
	err := json.NewDecoder(req.Body).Decode(dto)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	ser := &sql.DB{}
	s, err := h.Container.Get(ser)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}

	db, ok := (*s).(*sql.DB)

	if !ok {
		rw.WriteHeader(http.StatusInternalServerError)
	}

	row := db.QueryRow("SELECT id, name, email, token FROM users WHERE token = ?", dto.Token)

	var user model.User

	err = row.Scan(&user.Id, &user.Name, &user.Email, &user.Token)

	var data MessageResponses.AuthOperationResponse

	if err != nil {
		rw.WriteHeader(http.StatusForbidden)
		data = MessageResponses.AuthOperationResponse{
			Status:  http.StatusForbidden,
			Message: "Not Authorized",
		}
	} else {
		rw.WriteHeader(http.StatusOK)
		data = MessageResponses.AuthOperationResponse{
			Status:  http.StatusOK,
			Access:  true,
			Message: "Success",
		}
	}

	rw.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(data)

	if err != nil {
		panic(err)
	}
}
