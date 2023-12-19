package rpc

import (
	"bytes"
	"context"
	"github.com/kaiiorg/go-bif-examine/pkg/bif"
	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"
)

func (s *Server) UploadKey(ctx context.Context, req *pb.UploadKeyRequest) (*pb.UploadKeyResponse, error) {
	contents := req.GetContents()
	s.log.Info().Int("contentsLength", len(contents)).Msg("UploadKey")
	resp := &pb.UploadKeyResponse{}

	key, err := bif.NewKey(bytes.NewReader(contents), int64(len(contents)), s.log.With().Str("component", "bif-key").Logger())
	if err != nil {
		s.log.Warn().Err(err).Msg("Failed to parse key")
		resp.ErrorDescription = err.Error()
		return resp, err
	}

	audioResources := key.AudioEntriesToModel()

	resp.Key = &pb.Key{
		ParsedVersion:           key.Header.VersionToString(),
		ParsedSignature:         key.Header.SignatureToString(),
		ResourceEntryCount:      uint32(len(key.ResourceEntries)),
		BifFilesContainingAudio: []string{},
	}
	for bifFile := range audioResources {
		resp.Key.BifFilesContainingAudio = append(resp.Key.GetBifFilesContainingAudio(), bifFile)
	}

	// TODO create project record
	// TODO create record mapping the project to all expected bif files and the S3 object key for that file
	// TODO   the object will will be empty at first

	return resp, nil
}

func (s *Server) UploadBif(ctx context.Context, req *pb.UploadBifRequest) (*pb.UploadBifResponse, error) {
	s.log.Info().Msg("UploadBif")
	resp := &pb.UploadBifResponse{}
	return resp, nil
}
