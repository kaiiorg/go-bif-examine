package rpc

import (
	"context"
	"time"

	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"
)

func (s *Server) GetJob(context.Context, *pb.GetJobRequest) (*pb.GetJobResponse, error) {
	start := time.Now()
	s.log.Info().Msg("GetJob start")
	defer s.log.Info().Str("duration", time.Since(start).String()).Msg("GetJob end")

	resp := &pb.GetJobResponse{}
	return resp, nil
}

func (s *Server) JobResults(context.Context, *pb.JobResultsRequest) (*pb.JobResultsResponse, error) {
	start := time.Now()
	s.log.Info().Msg("JobResults start")
	defer s.log.Info().Str("duration", time.Since(start).String()).Msg("JobResults end")

	resp := &pb.JobResultsResponse{}
	return resp, nil
}
