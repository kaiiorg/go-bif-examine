package main

import (
	"context"
	"flag"
	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	applicationName        = "go-bif-examine-li"
	applicationDescription = "Quick and dirty CLI tool, for now."
)

var (
	logLevel   = flag.String("log-level", "info", "Zerolog log level to use; trace, debug, info, warn, error, panic, etc")
	grpcServer = flag.String("grpc-server", "localhost:50051", "IP:Port of gRPC server")
)

func main() {
	flag.Parse()
	configureLogging()
	client, err := NewRpcClient(*grpcServer)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to gRPC server")
	}

	req := &pb.DownloadResourceRequest{
		ResourceId: 17751,
	}
	resp, err := client.grpcClient.DownloadResource(context.Background(), req)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to download the requested resource")
	}

	err = os.WriteFile(resp.GetName(), resp.GetContent(), 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to save the downloaded resource")
	}
	log.Info().Str("filename", resp.GetName()).Msg("Downloaded resource to file")
}

func configureLogging() {
	// Configure pretty logs
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	zerologLevel, err := zerolog.ParseLevel(*logLevel)
	if err != nil || zerologLevel == zerolog.NoLevel {
		zerologLevel = zerolog.InfoLevel
		log.Warn().Str("givenLogLevel", *logLevel).Msg("Given an unexpected log level; defaulting to info level")
	}
	// Log application name and description just before changing the log level. This makes sure it always get printed
	log.Info().
		Str("applicationName", applicationName).
		Str("applicationDescription", applicationDescription).
		Send()

	zerolog.SetGlobalLevel(zerologLevel)
}
