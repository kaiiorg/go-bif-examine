package main

import (
	"flag"
	"github.com/kaiiorg/go-bif-examine/pkg/util"
	"github.com/kaiiorg/go-bif-examine/pkg/whisperer"

	"github.com/rs/zerolog/log"
)

const (
	applicationName        = "Whisperer"
	applicationDescription = "A rather dumb way of calling https://github.com/openai/whisper"
)

var (
	logLevel   = flag.String("log-level", "info", "Zerolog log level to use; trace, debug, info, warn, error, panic, etc")
	grpcServer = flag.String("grpc-server", "localhost:50051", "IP:Port of gRPC server")
)

func main() {
	flag.Parse()
	util.ConfigureLogging(*logLevel, applicationName, applicationDescription)

	w, err := whisperer.New(
		log.With().Str("component", "whisperer").Logger(),
		*grpcServer,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize Whisperer")
	}

	w.Run()
	util.WaitForInterrupt()
	w.Close()
}
