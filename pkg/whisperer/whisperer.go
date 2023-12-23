package whisperer

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	whisper_cli    = "whisper"
	whisper_github = "https://github.com/openai/whisper"
	whisper_model  = "medium.en"
)

type Whisperer struct {
	log       zerolog.Logger
	wg        *sync.WaitGroup
	ctx       context.Context
	ctxCancel context.CancelFunc

	gRpcClient pb.WhispererClient
	httpClient *http.Client

	whisperPath string
	tempDir     string
}

func New(log zerolog.Logger, gRpcServer string) (*Whisperer, error) {
	w := &Whisperer{
		log:        log,
		wg:         &sync.WaitGroup{},
		httpClient: &http.Client{},
	}
	w.ctx, w.ctxCancel = context.WithCancel(context.Background())

	err := w.checkWhisperAvailabilty()
	if err != nil {
		return nil, err
	}

	w.tempDir, err = os.MkdirTemp("", "whisperer")
	if err != nil {
		return nil, err
	}

	err = w.dialGRpc(gRpcServer)
	if err != nil {
		return nil, err
	}

	return w, nil
}

func (w *Whisperer) Run() {
	w.wg.Add(1)
	go w.doRun()
}

func (w *Whisperer) Close() {
	w.ctxCancel()
	w.wg.Wait()
}

func (w *Whisperer) doRun() {
	defer w.wg.Done()
	for {
		if w.ctx.Err() != nil {
			w.log.Warn().Msg("Whisperer doRun() got signal to exit!")
			return
		}

		w.log.Info().Msg("Getting job from go-bif-examine")
		job, err := w.gRpcClient.GetJob(w.ctx, &pb.GetJobRequest{})
		if err != nil {
			w.log.Warn().Err(err).Msg("Failed to get job from go-bif-examine; will try again later")
			w.timeout(30 * time.Second)
			continue
		}
		w.log.Info().Uint32("resource_id", job.GetResourceId()).Msg("Got job from go-bif-examine")
		w.log.Debug().Str("name", job.GetName()).Uint32("resource_id", job.GetResourceId()).
			Str("presigned_url", job.GetPresignedUrl()).Uint32("offset", job.GetOffset()).
			Uint32("size", job.GetSize()).Msg("Job details")

		audioFile := filepath.Join(w.tempDir, fmt.Sprintf("%s.wav", job.GetName()))

		w.log.Info().Uint32("resource_id", job.GetResourceId()).Msg("Downloading audio file from S3")
		err = w.downloadTargetFile(audioFile, job.GetPresignedUrl(), job.GetOffset(), job.GetSize())
		if err != nil {
			w.log.Error().Err(err).Msg("Failed to create a download the audio file from s3")
			w.timeout(30 * time.Second)
			continue
		}

		w.log.Info().Uint32("resource_id", job.GetResourceId()).Msg("Running Whisper")
		start := time.Now()
		resultFile, err := w.runWhisperer(audioFile)
		if err != nil {
			w.log.Error().Err(err).Msg("Failed to run whisper on the downloaded file")
			w.timeout(30 * time.Second)
			continue
		}
		duration := time.Since(start).String()
		w.log.Info().Str("duration", duration).Uint32("resource_id", job.GetResourceId()).Msg("Finished running Whisper")

		output, rawOutput, err := w.extractOutput(resultFile)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to run parse whisper output")
		}

		// Send the results in a go routine so we can immediately go get another job
		w.wg.Add(1)
		go w.sendResults(duration, job.GetResourceId(), output, rawOutput)
	}
}
