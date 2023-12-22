package main

import (
	"context"
	"flag"
	"os"

	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"
	"github.com/kaiiorg/go-bif-examine/pkg/util"

	"github.com/rs/zerolog/log"
)

const (
	applicationName        = "go-bif-examine-cli"
	applicationDescription = "Quick and dirty CLI tool, for now."
)

var (
	logLevel   = flag.String("log-level", "info", "Zerolog log level to use; trace, debug, info, warn, error, panic, etc")
	grpcServer = flag.String("grpc-server", "localhost:50051", "IP:Port of gRPC server")
)

func main() {
	flag.Parse()
	util.ConfigureLogging(*logLevel, applicationName, applicationDescription)
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
