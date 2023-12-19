package migration

import (
	"errors"
	"fmt"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Stage struct {
	Name     string
	Function func(db *gorm.DB, name string) error
}

type Migration struct {
	gorm.Model
	Name string
}

type Migrator struct {
	migrationsTable string
	log             zerolog.Logger
	migrations      []Stage
}

func New(migrationsTable string, migrations []Stage, log zerolog.Logger) *Migrator {
	return &Migrator{
		migrationsTable: migrationsTable,
		log:             log,
		migrations:      migrations,
	}
}

func (m *Migrator) Migrate(db *gorm.DB) error {
	if err := m.bootstrap(db); err != nil {
		return err
	}
	for _, migration := range m.migrations {
		// Check if this migration has already been applied
		if m.migrationExists(db, migration.Name) {
			continue
		}
		// Apply the migration
		if err := migration.Function(db, migration.Name); err != nil {
			return err
		}
		// Save that we've applied this migration
		if err := m.saveMigration(db, migration.Name); err != nil {
			return err
		}
	}
	return nil
}

func (m *Migrator) bootstrap(db *gorm.DB) error {
	migration := new(Migration)
	// Using Error instead of RecordNotFound because we want to check
	// if the migrations table exists. This is different from later migrations
	// where we query the already create migrations table.
	exists := nil == db.Where("name = ?", m.migrationsTable).First(migration).Error

	if exists {
		m.log.Debug().Str("migration", m.migrationsTable).Msg("Skipping migration; migrations table already applied")
		return nil
	}

	m.log.Info().Str("migration", m.migrationsTable).Msg("Creating migrations table")

	migrator := db.Migrator()

	// Create migrations table
	if err := migrator.CreateTable(new(Migration)); err != nil {
		return fmt.Errorf("error creating migrations table: %s", db.Error)
	}

	// Save a record to migrations table,
	// so we don't rerun this migration again
	migration.Name = m.migrationsTable
	if err := db.Create(migration).Error; err != nil {
		return fmt.Errorf("error saving record to migrations table: %s", err)
	}

	return nil
}

// migrationExists checks if the migration called migrationName has been run already
func (m *Migrator) migrationExists(db *gorm.DB, migrationName string) bool {
	migration := new(Migration)
	err := db.Where("name = ?", migrationName).First(migration).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		m.log.Info().Str("migration", migrationName).Msg("Skipping migration; already applied")
		return true
	}

	m.log.Info().Str("migration", migrationName).Msg("Running migration")
	return false
}

// saveMigration saves a migration to the migration table
func (m *Migrator) saveMigration(db *gorm.DB, migrationName string) error {
	migration := new(Migration)
	migration.Name = migrationName

	if err := db.Create(migration).Error; err != nil {
		m.log.Fatal().Err(err).Str("migration", migrationName).Msg("Error saving record to migrations table")
		return fmt.Errorf("error saving record to migrations table: %s", err)
	}

	return nil
}
