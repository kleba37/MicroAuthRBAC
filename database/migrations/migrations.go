package migrations

import "database/sql"

// Migration defines the interface for a single migration.
// Each migration file should implement this interface.
type Migration interface {
	// Name returns the name of the migration.
	Name() string
	// Up runs the migration.
	Up(*sql.DB) error
	// Down reverts the migration.
	Down(*sql.DB) error
}

// registry holds all the registered migrations.
var registry []Migration

// Register adds a migration to the registry.
// This function is called from the init() function of each migration file.
func Register(m Migration) {
	registry = append(registry, m)
}

// GetMigrations returns all registered migrations.
func GetMigrations() []Migration {
	return registry
}
