package main

import (
	"database/sql"
	"fmt"
	"go-test/internal/DI"
	"go-test/pkg/Router"
	"go-test/pkg/middleware"
	"go-test/pkg/middleware/CustomMiddleware"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}

	db := &sql.DB{}

	di := (&DI.DI{}).New()
	di.Register(db)

	httpHandler := http.HandlerFunc(Router.Router{Container: di}.Router)
	middlewares := Middleware.New(CustomMiddleware.AuthMiddleware{}.Handler).Apply(httpHandler)

	fmt.Println("Server started")
	fmt.Println("Listening on port " + os.Getenv("SERVER_PORT"))

	err = http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), middlewares)

	if err != nil {
		fmt.Println(err)
		return
	}
}
