package rpc

import (
	"context"

	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"
)

func (s *Server) DeleteProject(ctx context.Context, req *pb.DeleteProjectRequest) (*pb.DeleteProjectResponse, error) {
	s.log.Info().Msg("DeleteProject")
	resp := &pb.DeleteProjectResponse{}

	err := s.examineRepository.DeleteProject(uint(req.GetProjectId()))
	if err != nil {
		resp.ErrorDescription = err.Error()
		return resp, err
	}
	return resp, nil
}
