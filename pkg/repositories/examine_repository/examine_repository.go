package examine_repository

import (
	"github.com/kaiiorg/go-bif-examine/pkg/config"
	"github.com/kaiiorg/go-bif-examine/pkg/repositories/migration"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type GormExamineRepository struct {
	config *config.Config
	db     *gorm.DB
	log    zerolog.Logger
}

func New(conf *config.Config, db *gorm.DB, log zerolog.Logger) (*GormExamineRepository, error) {
	r := &GormExamineRepository{
		config: conf,
		db:     db,
		log:    log,
	}
	migrator := migration.New("examine_migrations", r.getMigrationStages(), r.log)
	err := migrator.Migrate(db)
	return r, err
}
