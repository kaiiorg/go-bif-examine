package examine_repository

import (
	"github.com/kaiiorg/go-bif-examine/pkg/models"
	"github.com/kaiiorg/go-bif-examine/pkg/repositories/migration"
	"gorm.io/gorm"
)

func (r *GormExamineRepository) getMigrationStages() []migration.Stage {
	return []migration.Stage{
		{
			Name:     "init",
			Function: r.migrationInit,
		},
	}
}

func (r *GormExamineRepository) migrationInit(db *gorm.DB, name string) error {
	migrator := db.Migrator()

	if err := migrator.AutoMigrate(models.Project{}); err != nil {
		r.log.Error().Err(err).Str("migration", name).Msg("Failed to create project table")
		return err
	}

	if err := migrator.AutoMigrate(models.Bif{}); err != nil {
		r.log.Error().Err(err).Str("migration", name).Msg("Failed to create bif table")
		return err
	}

	if err := migrator.AutoMigrate(models.Resource{}); err != nil {
		r.log.Error().Err(err).Str("migration", name).Msg("Failed to create resource table")
		return err
	}

	return nil
}
