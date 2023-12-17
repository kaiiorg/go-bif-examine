package bif_examine

import (
	"context"

	"github.com/kaiiorg/go-bif-examine/pkg/config"
	"github.com/kaiiorg/go-bif-examine/pkg/rpc"
	"github.com/kaiiorg/go-bif-examine/pkg/storage"
	"github.com/kaiiorg/go-bif-examine/pkg/web"
	"github.com/kaiiorg/go-bif-examine/pkg/bif"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type BifExamine struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	config    *config.Config

	log zerolog.Logger

	storage *storage.Storage
	rpc     *rpc.Server
	web     *web.Web
}

func New(conf *config.Config) (*BifExamine, error) {
	s3Storage, err := storage.New(conf, log.With().Str("component", "storage").Logger())
	if err != nil {
		return nil, err
	}
	be := &BifExamine{
		config:  conf,
		log:     log.With().Str("component", "general").Logger(),
		storage: s3Storage,
		rpc:     rpc.New(conf, log.With().Str("component", "rpc").Logger()),
		web:     web.New(conf, log.With().Str("component", "web").Logger()),
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

	key, err := bif.NewKeyFromFile("./test_bifs/chitin.key", log.With().Str("component", "bif-key").Logger())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read KEY")
	}
	log.Info().
		Interface("key", key).
		Str("version", key.Header.VersionToString()).
		Str("signature", key.Header.SignatureToString()).
		Msg("read KEY")

	return nil
}

func (be *BifExamine) Close() error {
	be.log.Info().Msg("Closing!")
	be.rpc.Close()
	be.web.Close()
	be.ctxCancel()
	return nil
}
