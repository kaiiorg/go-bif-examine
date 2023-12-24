package rpc

import (
	"context"
	"time"

	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"
)

func (s *Server) GetJob(ctx context.Context, req *pb.GetJobRequest) (*pb.GetJobResponse, error) {
	start := time.Now()
	s.log.Info().Msg("GetJob start")
	defer s.log.Info().Str("duration", time.Since(start).String()).Msg("GetJob end")
	resp := &pb.GetJobResponse{}

	resource, err := s.examineRepository.GetResourceForWhisper()
	if err != nil {
		s.log.Error().Err(err).Msg("Failed to get a resource for whisperer")
		return resp, err
	}

	resourceBif, err := s.examineRepository.GetBifById(resource.BifID)
	if err != nil {
		s.log.Error().
			Err(err).
			Uint("resourceId", resource.ID).
			Uint("bifId", resource.BifID).
			Msg("Failed to get the bif for the selected resource for whisperer")
		return resp, err
	}
	// Shouldn't need to worry about this check failing, but we're going to check anyway
	if resourceBif.ObjectKey == nil {
		s.log.Error().
			Uint("resourceId", resource.ID).
			Uint("bifId", resource.BifID).
			Msg("The bif for the resource hasn't been uploaded!")
		return resp, ErrBifNotYetUploaded
	}

	resp.PresignedUrl, err = s.storage.PresignGetObject(*resourceBif.ObjectKey)
	if err != nil {
		s.log.Error().Err(err).Msg("Failed to get a presigned URL to the resource for whisper")
		return resp, err
	}

	resp.Name = resource.Name
	resp.ResourceId = uint32(resource.ID)
	resp.Offset = resource.OffsetToData
	resp.Size = resource.Size

	return resp, nil
}

func (s *Server) JobResults(ctx context.Context, req *pb.JobResultsRequest) (*pb.JobResultsResponse, error) {
	start := time.Now()
	s.log.Info().Msg("JobResults start")
	defer s.log.Info().Str("duration", time.Since(start).String()).Msg("JobResults end")

	resource, err := s.examineRepository.GetResourceById(uint(req.GetResourceId()))
	if err != nil {
		return &pb.JobResultsResponse{}, nil
	}
	resource.Text = req.GetText()
	resource.RawOutput = string(req.GetRawOutput())
	resource.WhisperModel = req.GetModel()
	resource.JobDuration = req.GetDuration()

	err = s.examineRepository.UpdateResource(resource)
	if err != nil {
		return &pb.JobResultsResponse{}, err
	}

	s.log.Info().
		Uint32("resourceId", req.GetResourceId()).
		Str("text", req.GetText()).
		Str("model", req.GetModel()).
		Str("duration", req.GetDuration()).
		Msg("Got and saved result from an instance of whisperer")

	return &pb.JobResultsResponse{}, nil
}
