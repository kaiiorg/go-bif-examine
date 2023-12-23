package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

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
	model   = "medium.en"
)

var (
	logLevel   = flag.String("log-level", "info", "Zerolog log level to use; trace, debug, info, warn, error, panic, etc")
	grpcServer = flag.String("grpc-server", "localhost:50051", "IP:Port of gRPC server")
)

func main() {
	flag.Parse()
	util.ConfigureLogging(*logLevel, applicationName, applicationDescription)

	whichWhisper := checkWhisperAvailabilty()
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

	outputDir, err := os.MkdirTemp("", "whisperer")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create a temp dir")
	}

	audioFile, err := downloadTargetFile(outputDir, job.GetName(), job.GetPresignedUrl(), job.GetOffset(), job.GetSize())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create a download the audio file from s3")
	}

	start := time.Now()
	outputFile, err := runWhisperer(outputDir, whichWhisper, audioFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to run whisper on the file")
	}
	duration := time.Since(start).String()
	log.Info().
		Str("outputFile", outputFile).
		Str("duration", duration).
		Msg("Whisper ran and should have written something to the given file")

	output, rawOutput, err := extractWhisperOutput(outputFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to run parse whisper output")
	}

	jobResults := &pb.JobResultsRequest{
		ResourceId: job.GetResourceId(),
		Text:       output.Text,
		RawOutput:  rawOutput,
		Model:      "medium.en",
		Duration:   duration,
	}
	_, err = client.JobResults(context.Background(), jobResults)
	if err == nil {
		log.Info().Msg("Successfully submitted job results to go-bif-examine")
	} else {
		log.Fatal().Err(err).Msg("Failed to submit job results to go-bif-examine")
	}
}

func checkWhisperAvailabilty() string {
	path, err := exec.LookPath(whisper)
	if err != nil {
		log.Fatal().Err(err).Str("github", "https://github.com/openai/whisper").Msg("Unable to validate that whisper is installed. See its github page for install instructions")
	}
	log.Info().Str("path", path).Msg("Found whisper install!")
	return path
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

func downloadTargetFile(outputDir, outputFilename, presignedUrl string, offsetToData uint32, size uint32) (string, error) {
	saveTo := filepath.Join(outputDir, fmt.Sprintf("%s.wav", outputFilename))
	file, err := os.Create(saveTo)
	if err != nil {
		return "", err
	}
	defer file.Close()
	log.Info().Str("tempFile", saveTo).Msg("Created temporary file")

	req, err := http.NewRequest(http.MethodGet, presignedUrl, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", offsetToData, offsetToData+size))

	// TODO don't recreate the http client every time
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 && resp.StatusCode != 206 {
		log.Fatal().Str("status", resp.Status).Msg("Response from S3 was not a 200 OK or 206 Partial Content")
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}

	log.Info().Str("path", saveTo).Msg("Downloaded audio file")

	return saveTo, nil
}

func runWhisperer(outputDir, whichWhisper, inputFilepath string) (string, error) {
	args := []string{
		whichWhisper,
		inputFilepath,
		"--model", model,
		"--output_format", "json",
		"--output_dir", outputDir,
		"--language", "en",
	}
	log.Info().Strs("args", args).Msg("Attempting to run whisper")

	cmd := &exec.Cmd{
		Path: whichWhisper,
		Args: args,
		// Stdout: os.Stdout,
		// Stderr: os.Stderr,
	}

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	wavName := filepath.Base(inputFilepath)
	name := strings.TrimSuffix(wavName, filepath.Ext(wavName))
	jsonName := fmt.Sprintf("%s.json", name)
	log.Info().Str("wavName", wavName).Str("name", name).Str("jsonName", jsonName).Send()

	return filepath.Join(outputDir, jsonName), nil
}

func extractWhisperOutput(outputFile string) (*WhisperOutput, []byte, error) {
	contents, err := os.ReadFile(outputFile)
	if err != nil {
		return nil, nil, err
	}

	whisperOutput := &WhisperOutput{Segments: []WhisperOutputSegments{}}

	err = json.Unmarshal(contents, whisperOutput)
	if err != nil {
		return nil, nil, err
	}

	return whisperOutput, contents, nil
}