package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

func New() *sql.DB {
	err := godotenv.Load()

	if err != nil {
		return nil
	}

	dsn := os.Getenv("DB_DSN")

	if len(dsn) == 0 {
		panic("DSN invalid")
	}

	pool, err := sql.Open(os.Getenv("DB_CONNECTION"), dsn)

	if err != nil {
		fmt.Println("DB error connection")
		panic(err)
	}

	fmt.Println("DB connect with DSN: ", dsn)

	return pool
}
