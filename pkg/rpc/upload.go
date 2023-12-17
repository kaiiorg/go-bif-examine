package rpc

import (
	"context"

	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"
)

func (s *Server) UploadBifKey(ctx context.Context, req *pb.UploadBifKeyRequest) (*pb.UploadBifKeyResponse, error) {
	s.log.Info().
		Str("projectName", req.GetProjectName()).
		Int("contentsLength", len(req.GetContents())).
		Msg("UploadBifKey called")
	resp := &pb.UploadBifKeyResponse{Ok: true}
	return resp, nil
}
