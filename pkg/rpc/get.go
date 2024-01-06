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

func (s *Server) GetProjectById(ctx context.Context, req *pb.GetProjectByIdRequest) (*pb.GetProjectByIdResponse, error) {
	s.log.Info().Msg("GetProjectById")
	resp := &pb.GetProjectByIdResponse{}

	project, err := s.examineRepository.GetProjectById(uint(req.GetId()))
	if err != nil {
		resp.ErrorDescription = err.Error()
		return resp, err
	}

	resp.Project = &pb.Project{
		Id:                  uint32(project.ID),
		Name:                project.Name,
		OriginalKeyFileName: project.OriginalKeyFileName,
	}

	return resp, nil
}

func (s *Server) GetBifsMissingContents(ctx context.Context, req *pb.GetBifsMissingContentsRequest) (*pb.GetBifsMissingContentsResponse, error) {
	s.log.Info().Msg("GetBifsMissingContents")
	resp := &pb.GetBifsMissingContentsResponse{}

	bifs, err := s.examineRepository.GetBifsMissingContent(uint(req.GetProjectId()))
	if err != nil {
		resp.ErrorDescription = err.Error()
		return resp, err
	}

	for _, bif := range bifs {
		resp.NameInKey = append(
			resp.NameInKey,
			bif.NameInKey,
		)
	}

	return resp, nil
}
