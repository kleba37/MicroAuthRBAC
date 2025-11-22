package main

import (
	"database/sql"
	"fmt"
	"go-test/database/migrations"
	DB "go-test/service/db"
	"log"

	_ "modernc.org/sqlite"

	DI "github.com/kleba37/GoServiceContainer"
)

type Migrated struct{}

func main() {
	di := DI.New()

	ser := DB.New()

	di.Register(ser)

	s, err := di.Get(&sql.DB{})

	if err != nil {
		log.Fatalf(err.Error())
	}

	db := (*s).(*sql.DB)

	if !existTableMigrations(db) {
		err := addSystemTable(db)

		if err != nil {
			log.Fatalf("Could not add system table")
		}
	}

	Migrate(db)
}

func Migrate(db *sql.DB) {
	registeredMigrations := migrations.GetMigrations()

	fmt.Printf("Found %d migrations to run\n", len(registeredMigrations))

	for _, m := range registeredMigrations {
		var exist bool

		err := db.QueryRow("SELECT COUNT(id) FROM `migrations` WHERE `name` = ?", m.Name()).Scan(&exist)

		if err != nil {
			log.Fatalf(err.Error())
		}

		if !exist {
			t, err := db.Begin()

			fmt.Printf("Running migration: %s...\n", m.Name())
			err = m.Up(db)

			if err != nil {
				err = t.Rollback()
				panic(fmt.Sprintf("Failed to run migration %s: %v", m.Name(), err))
			}

			_, err = db.Exec("INSERT INTO migrations (name) VALUES (?)", m.Name())

			if err != nil {
				err = t.Rollback()

				if err != nil {
					log.Fatalf(err.Error())
				}
			}

			err = t.Commit()

			if err != nil {
				log.Fatalf(err.Error())
			}

			fmt.Printf("migrate %s finished successfully\n", m.Name())
		}

	}

	fmt.Println("All migrations completed.")
}

func addSystemTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			Name TEXT NOT NULL
		);
	`)

	if err != nil {
		return err
	}

	return nil
}

func existTableMigrations(db *sql.DB) bool {
	var exists bool

	err := db.QueryRow("SELECT COUNT(*) FROM sqlite_master WHERE tbl_name = 'migrations' AND `type` = 'table'").Scan(&exists)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return exists
}
