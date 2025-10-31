package DI

import (
	"database/sql"
	"fmt"
	"go-test/pkg/Container"
	"os"

	_ "modernc.org/sqlite"

	"github.com/joho/godotenv"
)

var autorun = map[string]Container.Service{
	//"DB": db(),
}

type DI struct{}

func (DI *DI) New() *Container.Container {
	container := Container.NewContainer()

	for _, ser := range autorun {
		container.Register(ser)
	}

	return container
}

// Для autoload
func db() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		return nil
	}

	dsn := os.Getenv("DB_DSN")

	fmt.Println("DSN: ", dsn)
	if len(dsn) == 0 {
		panic("DSN invalid")
	}

	pool, err := sql.Open(os.Getenv("DB_CONNECTION"), dsn)

	if err != nil {
		fmt.Println("DB error connection")
		panic(err)
	}

	return pool
}
