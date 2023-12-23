package whisperer

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"
)

func (w *Whisperer) downloadTargetFile(downloadTo, presignedUrl string, offsetToData uint32, size uint32) error {
	// Create the file to save the audio bytes to
	file, err := os.Create(downloadTo)
	if err != nil {
		return err
	}
	defer file.Close()
	w.log.Trace().Str("downloadTo", downloadTo).Msg("created a temporary file to download the audio data to")

	// Build the request to the presigned URL with the range header to select only the bytes we've been told are relevant
	req, err := http.NewRequest(http.MethodGet, presignedUrl, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", offsetToData, offsetToData+size))

	// Fire the request
	resp, err := w.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Return an error if we didn't get 200 or 206
	if resp.StatusCode != 200 && resp.StatusCode != 206 {
		w.log.Error().Str("status", resp.Status).Msg("Response from S3 was not a 200 OK or 206 Partial Content")
		return ErrFailedToDownloadFileFromS3Not200Or206
	}

	// Copy the bytes from the body to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	w.log.Trace().Str("downloadTo", downloadTo).Msg("Downloaded audio data to temporary file")

	return nil
}

func (w *Whisperer) sendResults(duration string, resourceId uint32, output *WhisperOutput, rawOutput []byte) {
	defer w.wg.Done()
	w.log.Info().Uint32("resource_id", resourceId).Msg("Attempting to send results to go-bif-examine")

	jobResults := &pb.JobResultsRequest{
		ResourceId: resourceId,
		Text:       output.Text,
		RawOutput:  rawOutput,
		Model:      whisper_model,
		Duration:   duration,
	}

	for attempt := 1; attempt < 5; attempt++ {
		if w.ctx.Err() != nil {
			w.log.Warn().Uint32("resourceId", resourceId).Msg("Not sending results for job due to shutdown request")
			return
		}

		_, err := w.gRpcClient.JobResults(w.ctx, jobResults)
		if err == nil {
			w.log.Info().Uint32("resourceId", resourceId).Msg("Successfully submitted job results to go-bif-examine")
			return
		} else {
			w.log.Warn().Err(err).Uint32("resourceId", resourceId).Int("attempt", attempt).
				Msg("Failed to submit job results to go-bif-examine; may attempt to retry")
			w.timeout(5 * time.Second)
		}
		w.log.Error().Uint32("resourceId", resourceId).
			Msg("Giving up on submitting results to go-bif-examine")
	}
}
