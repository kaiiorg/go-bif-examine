package bif_examine

import (
	"context"

	"github.com/kaiiorg/go-bif-examine/pkg/config"
	"github.com/kaiiorg/go-bif-examine/pkg/rpc"
	"github.com/kaiiorg/go-bif-examine/pkg/web"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type BifExamine struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	config    *config.Config

	log zerolog.Logger

	rpc *rpc.Server
	web *web.Web
}

func New(conf *config.Config) *BifExamine {
	be := &BifExamine{
		config: conf,
		log:    log.With().Str("component", "general").Logger(),
		rpc:    rpc.New(conf, log.With().Str("component", "rpc").Logger()),
		web:    web.New(conf, log.With().Str("component", "web").Logger()),
	}
	be.ctx, be.ctxCancel = context.WithCancel(context.Background())

	return be
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
