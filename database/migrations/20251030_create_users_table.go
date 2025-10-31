package migrations

import (
	"database/sql"
	"fmt"
)

// Automatically register the migration when the application starts
func init() {
	Register(&CreateUsersTable{})
	fmt.Println("Registered migration: CreateUsersTable")
}

type CreateUsersTable struct{}

func (m *CreateUsersTable) Name() string {
	return "20251030_create_users_table"
}

func (m *CreateUsersTable) Up(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			token TEXT
		);
	`)
	return err
}

func (m *CreateUsersTable) Down(db *sql.DB) error {
	_, err := db.Exec(`DROP TABLE IF EXISTS users;`)
	return err
}
