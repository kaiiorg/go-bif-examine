package bif_examine

import (
	"context"

	"github.com/kaiiorg/go-bif-examine/pkg/config"
	"github.com/kaiiorg/go-bif-examine/pkg/rpc"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type BifExamine struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	config    *config.Config

	log zerolog.Logger

	rpc *rpc.Server
}

func New(conf *config.Config) *BifExamine {
	be := &BifExamine{
		config: conf,
		log:    log.With().Str("component", "bif-examine-general").Logger(),
		rpc:    rpc.New(conf, log.With().Str("component", "bif-examine-rpc").Logger()),
	}
	be.ctx, be.ctxCancel = context.WithCancel(context.Background())

	return be
}

func (be *BifExamine) Run() error {
	err := be.rpc.Run()
	if err != nil {
		return err
	}

	be.log.Info().Msg("Running!")
	return nil
}

func (be *BifExamine) Close() error {
	be.log.Info().Msg("Closing!")
	be.rpc.Close()
	be.ctxCancel()
	return nil
}
