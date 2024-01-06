package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"
)

func (s *Server) DownloadResource(ctx context.Context, req *pb.DownloadResourceRequest) (*pb.DownloadResourceResponse, error) {
	start := time.Now()
	s.log.Info().Msg("DownloadResource start")
	defer s.log.Info().Str("duration", time.Since(start).String()).Msg("DownloadResource end")
	resp := &pb.DownloadResourceResponse{}

	// Get the details about this resource
	resource, err := s.examineRepository.GetResourceById(uint(req.GetResourceId()))
	if err != nil {
		s.log.Warn().Err(err).Uint32("resource", req.GetResourceId()).Msg("Failed to find resource record for given resource id")
		resp.ErrorDescription = err.Error()
		return resp, err
	}

	// TODO validate that the offset to data and size aren't 0!
	resp.Name = fmt.Sprintf("%s.wav", resource.Name)
	resp.Size = resource.Size

	// Get the details about the bif related to this resource
	bif, err := s.examineRepository.GetBifById(resource.BifID)
	if err != nil {
		s.log.Warn().Err(err).Uint32("resource", req.GetResourceId()).Uint("bif", resource.BifID).Msg("Failed to find bif record related to this resource")
		resp.ErrorDescription = err.Error()
		return resp, err
	}

	// TODO validate that the S3 key is set!

	// Get the object from S3
	resp.Content, err = s.storage.GetSectionFromObject(*bif.ObjectHash, resource.OffsetToData, resource.Size)

	return resp, nil
}
