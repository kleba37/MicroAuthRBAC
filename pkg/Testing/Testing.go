package TestingTools

import (
	"database/sql"
	"go-test/database/migrations"
	"go-test/internal/DI"
	"go-test/pkg/Router"
	Middleware "go-test/pkg/middleware"
	"go-test/pkg/middleware/CustomMiddleware"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	// Пустой импорт нужен, чтобы сработали init() функции в файлах миграций
	_ "go-test/database/migrations"
)

type TestingTools struct {
	DB     *sql.DB
	Tx     *sql.Tx
	Router http.Handler
}

func New() (*TestingTools, error) {
	err := godotenv.Load("../../../.env.testing")
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite", os.Getenv("DB_DSN"))
	if err != nil {
		return nil, err
	}

	for _, migration := range migrations.GetMigrations() {
		err := migration.Up(db)
		if err != nil {
			return nil, err
		}
	}

	di := (&DI.DI{}).New()
	di.Register(db)

	httpHandler := http.HandlerFunc(Router.Router{Container: di}.Router)
	middlewares := Middleware.New(CustomMiddleware.AuthMiddleware{}.Handler).Apply(httpHandler)

	return &TestingTools{
		DB:     db,
		Router: middlewares,
	}, nil
}

func (t *TestingTools) Serve(rr http.ResponseWriter, req *http.Request) {
	t.Router.ServeHTTP(rr, req)
}

func (t *TestingTools) StartTest() error {
	tx, err := t.DB.Begin()
	if err != nil {
		return err
	}
	t.Tx = tx
	return nil
}

func (t *TestingTools) EndTest() error {
	if t.Tx != nil {
		err := t.Tx.Rollback()
		t.Tx = nil
		return err
	}
	return nil
}
