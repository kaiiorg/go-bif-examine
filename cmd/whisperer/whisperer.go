package main

import (
	"context"
	"flag"
	"github.com/google/uuid"
	"os/exec"

	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"
	"github.com/kaiiorg/go-bif-examine/pkg/util"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	applicationName        = "Whisperer"
	applicationDescription = "A rather dumb way of calling https://github.com/openai/whisper"

	whisper = "whisper"
)

var (
	logLevel   = flag.String("log-level", "info", "Zerolog log level to use; trace, debug, info, warn, error, panic, etc")
	grpcServer = flag.String("grpc-server", "localhost:50051", "IP:Port of gRPC server")
)

func main() {
	flag.Parse()
	util.ConfigureLogging(*logLevel, applicationName, applicationDescription)

	checkWhisperAvailabilty()
	client := dial()

	job, err := client.GetJob(context.Background(), &pb.GetJobRequest{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get job from go-bif-examine")
	}
	log.Info().
		Str("name", job.GetName()).
		Uint32("resource_id", job.GetResourceId()).
		Str("presigned_url", job.GetPresignedUrl()).
		Uint32("offset", job.GetOffset()).
		Uint32("size", job.GetSize()).
		Msg("Got job from go-bif-examine")

	log.Info().Msg("This is the part where we'd run the job, but that isn't implemented yet")

	jobResults := &pb.JobResultsRequest{
		ResourceId: job.GetResourceId(),
		Text:       uuid.NewString(),
		RawOutput:  uuid.NewString(),
		Model:      "medium.en",
	}
	_, err = client.JobResults(context.Background(), jobResults)
	if err == nil {
		log.Info().Msg("Successfully submitted job results to go-bif-examine")
	} else {
		log.Fatal().Err(err).Msg("Failed to submit job results to go-bif-examine")
	}
}

func checkWhisperAvailabilty() {
	path, err := exec.LookPath(whisper)
	if err != nil {
		log.Fatal().Err(err).Str("github", "https://github.com/openai/whisper").Msg("Unable to validate that whisper is installed. See its github page for install instructions")
	}
	log.Info().Str("path", path).Msg("Found whisper install!")
}

func dial() pb.WhispererClient {
	conn, err := grpc.Dial(
		*grpcServer,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to gRPC server")
	}
	return pb.NewWhispererClient(conn)
}
