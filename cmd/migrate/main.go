package main

import (
	"database/sql"
	"fmt"
	"go-test/database/migrations"
	"go-test/internal/DI"
)

type Migrated struct {
}

func main() {
	di := (&DI.DI{}).New()

	ser := sql.DB{}
	s := di.Get(&ser)
	db := (*s).(*sql.DB)

	Migrate(db)
}

func Migrate(db *sql.DB) {
	registeredMigrations := migrations.GetMigrations()

	fmt.Printf("Found %d migrations to run\n", len(registeredMigrations))

	for _, m := range registeredMigrations {
		fmt.Printf("Running migration: %s...\n", m.Name())
		err := m.Up(db)

		if err != nil {
			panic(fmt.Sprintf("Failed to run migration %s: %v", m.Name(), err))
		}

		fmt.Printf("migrate %s finished successfully\n", m.Name())
	}

	fmt.Println("All migrations completed.")
}
