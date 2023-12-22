package rpc

import (
	"context"

	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"
)

func (s *Server) GetAllProjects(ctx context.Context, req *pb.GetAllProjectsRequest) (*pb.GetAllProjectsResponse, error) {
	s.log.Info().Msg("GetAllProjects")
	resp := &pb.GetAllProjectsResponse{}
	allProjects, err := s.examineRepository.GetAllProjects()

	if err != nil {
		s.log.Warn().Err(err).Msg("GetAllProjects failed")
		resp.ErrorDescription = err.Error()
		return resp, err
	}

	for _, project := range allProjects {
		resp.Projects = append(
			resp.Projects,
			&pb.Project{
				Id:                  uint32(project.ID),
				Name:                project.Name,
				OriginalKeyFileName: project.OriginalKeyFileName,
			},
		)
	}

	return resp, nil
}

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
