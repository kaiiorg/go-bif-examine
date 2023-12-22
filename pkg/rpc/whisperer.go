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

	resp.Name = resource.Name
	resp.ResourceId = uint32(resource.ID)
	resp.PresignedUrl = "tbd"
	resp.Offset = resource.OffsetToData
	resp.Size = resource.Size

	return resp, nil
}

func (s *Server) JobResults(ctx context.Context, req *pb.JobResultsRequest) (*pb.JobResultsResponse, error) {
	start := time.Now()
	s.log.Info().Msg("JobResults start")
	defer s.log.Info().Str("duration", time.Since(start).String()).Msg("JobResults end")

	s.log.Info().
		Uint32("resourceId", req.GetResourceId()).
		// Str("text", req.GetText()).
		// Bytes("rawOutput", req.GetRawOutput()).
		Str("model", req.GetModel()).
		Str("duration", req.GetDuration()).
		Msg("Got result from an instance of whisperer")

	resp := &pb.JobResultsResponse{}
	return resp, nil
}
