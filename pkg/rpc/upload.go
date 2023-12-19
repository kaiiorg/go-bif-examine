package rpc

import (
	"bytes"
	"context"
	"github.com/kaiiorg/go-bif-examine/pkg/bif"
	"github.com/kaiiorg/go-bif-examine/pkg/models"
	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"
)

func (s *Server) UploadKey(ctx context.Context, req *pb.UploadKeyRequest) (*pb.UploadKeyResponse, error) {
	contents := req.GetContents()
	s.log.Info().Int("contentsLength", len(contents)).Msg("UploadKey")
	resp := &pb.UploadKeyResponse{}

	// Parse the contents as a bif key
	key, err := bif.NewKey(bytes.NewReader(contents), int64(len(contents)), s.log.With().Str("component", "bif-key").Logger())
	if err != nil {
		s.log.Warn().Err(err).Msg("Failed to parse key")
		resp.ErrorDescription = err.Error()
		return resp, err
	}
	// Determine what bif files the key claims contains audio resources
	audioResources := key.AudioEntriesToModel()

	// Create project record and store it in the DB
	project := &models.Project{
		Name:                req.GetProjectName(),
		OriginalKeyFileName: req.GetFileName(),
	}
	projectId, err := s.examineRepository.CreateProject(project)
	if err != nil {
		s.log.Warn().Err(err).Msg("Failed to save project to db")
		resp.ErrorDescription = err.Error()
		return resp, err
	}
	resp.ProjectId = uint32(projectId)

	// Build the response key information
	resp.Key = &pb.Key{
		ParsedVersion:           key.Header.VersionToString(),
		ParsedSignature:         key.Header.SignatureToString(),
		ResourceEntryCount:      uint32(len(key.ResourceEntries)),
		BifFilesContainingAudio: map[string]uint32{},
	}
	for bifFile, resources := range audioResources {
		resp.Key.BifFilesContainingAudio[bifFile] = uint32(len(resources))
	}

	// TODO create record mapping the project to all expected bif files and the S3 object key for that file
	// TODO   the object will will be empty at first

	return resp, nil
}

func (s *Server) UploadBif(ctx context.Context, req *pb.UploadBifRequest) (*pb.UploadBifResponse, error) {
	s.log.Info().Msg("UploadBif")
	resp := &pb.UploadBifResponse{}
	return resp, nil
}
