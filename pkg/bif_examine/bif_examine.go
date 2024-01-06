package bif_examine

import (
	"context"
	"github.com/kaiiorg/go-bif-examine/pkg/config"
	"github.com/kaiiorg/go-bif-examine/pkg/repositories/examine_repository"
	"github.com/kaiiorg/go-bif-examine/pkg/repositories/gorm_logger"
	"github.com/kaiiorg/go-bif-examine/pkg/rpc"
	"github.com/kaiiorg/go-bif-examine/pkg/storage"
	"github.com/kaiiorg/go-bif-examine/pkg/web"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type BifExamine struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	config    *config.Config

	log zerolog.Logger

	storage           storage.BifStorage
	rpc               *rpc.Server
	web               *web.Web
	examineRepository examine_repository.ExamineRepository
}

func New(conf *config.Config) (*BifExamine, error) {
	// Setup S3 storage
	s3Storage, err := storage.New(conf, log.With().Str("component", "storage").Logger())
	if err != nil {
		return nil, err
	}
	// Setup DB and repository
	db, err := gorm.Open(
		postgres.Open(conf.Db.ConnectionString()),
		&gorm.Config{
			Logger: gorm_logger.NewGormLogger(log.With().Str("component", "db-gorm").Logger(), zerolog.GlobalLevel()),
		},
	)
	if err != nil {
		return nil, err
	}
	examineRepository, err := examine_repository.New(conf, db, log.With().Str("component", "db").Logger())
	if err != nil {
		return nil, err
	}

	// Setup main struct along with the cli.go and web components
	be := &BifExamine{
		config:            conf,
		log:               log.With().Str("component", "general").Logger(),
		storage:           s3Storage,
		rpc:               rpc.New(examineRepository, s3Storage, conf, log.With().Str("component", "cli.go").Logger()),
		web:               web.New(conf, log.With().Str("component", "web").Logger()),
		examineRepository: examineRepository,
	}
	be.ctx, be.ctxCancel = context.WithCancel(context.Background())

	return be, nil
}

func (be *BifExamine) Run() error {
	err := be.rpc.Run()
	if err != nil {
		return err
	}

	be.web.Run()

	be.log.Info().Msg("Running!")

	return nil
}

func (be *BifExamine) Close() error {
	be.log.Info().Msg("Closing!")
	be.rpc.Close()
	be.web.Close()
	be.ctxCancel()
	return nil
}
