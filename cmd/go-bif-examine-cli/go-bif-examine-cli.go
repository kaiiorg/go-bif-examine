package main

import (
	"github.com/rs/zerolog/log"
)

func main() {
	if err := NewCli().Execute(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
