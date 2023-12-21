package rpc

import (
	"bytes"
	"context"
	"github.com/kaiiorg/go-bif-examine/pkg/bif"
	"github.com/kaiiorg/go-bif-examine/pkg/models"
	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"
	"github.com/kaiiorg/go-bif-examine/pkg/util"
	"path/filepath"
	"strings"
	"time"
)

func (s *Server) UploadKey(ctx context.Context, req *pb.UploadKeyRequest) (*pb.UploadKeyResponse, error) {
	contents := req.GetContents()
	start := time.Now()
	s.log.Info().Int("contentsLength", len(contents)).Msg("UploadKey start")
	defer s.log.Info().Str("duration", time.Since(start).String()).Msg("UploadKey end")
	resp := &pb.UploadKeyResponse{}

	// Parse the contents as a bif key
	key, err := bif.NewKey(bytes.NewReader(contents), int64(len(contents)), s.log.With().Str("component", "bif-key").Logger())
	if err != nil {
		s.log.Warn().Err(err).Msg("Failed to parse key")
		resp.ErrorDescription = err.Error()
		return resp, err
	}

	// Create project record and store it in the DB
	project := &models.Project{
		Name:                req.GetProjectName(),
		OriginalKeyFileName: req.GetFileName(),
	}
	project.ID, err = s.examineRepository.CreateProject(project)
	if err != nil {
		s.log.Warn().Err(err).Msg("Failed to save project to db")
		resp.ErrorDescription = err.Error()
		return resp, err
	}
	resp.ProjectId = uint32(project.ID)

	// Determine what bif files the key claims contains audio resources
	audioResources, bifsWithAudio := key.AudioEntriesToModel(project)

	// Add each bif to the database
	err = s.examineRepository.CreateManyBifs(bifsWithAudio)
	if err != nil {
		s.log.Error().Err(err).Int("count", len(bifsWithAudio)).Msg("Failed to add bifs to DB")
		resp.ErrorDescription = err.Error()
		return resp, err
	}

	// Add each audio resource to the database
	err = s.examineRepository.CreateManyResources(audioResources)
	if err != nil {
		s.log.Error().Err(err).Int("count", len(bifsWithAudio)).Msg("Failed to add resources to DB")
		resp.ErrorDescription = err.Error()
		return resp, err
	}

	// Build the response key information
	resp.Key = &pb.Key{
		ParsedVersion:           key.Header.VersionToString(),
		ParsedSignature:         key.Header.SignatureToString(),
		ResourceEntryCount:      uint32(len(key.ResourceEntries)),
		ResourcesWithAudio:      uint32(len(audioResources)),
		BifFilesContainingAudio: []string{},
	}
	for _, bifWithAudio := range bifsWithAudio {
		resp.Key.BifFilesContainingAudio = append(resp.Key.BifFilesContainingAudio, bifWithAudio.NameInKey)
	}

	return resp, nil
}

func (s *Server) UploadBif(ctx context.Context, req *pb.UploadBifRequest) (*pb.UploadBifResponse, error) {
	contents := req.GetContents()
	start := time.Now()
	s.log.Info().Int("contentsLength", len(contents)).Msg("UploadBif start")
	defer s.log.Info().Str("duration", time.Since(start).String()).Msg("UploadBif end")
	resp := &pb.UploadBifResponse{}

	// Parse the contents as a bif file
	bif, err := bif.NewBif(bytes.NewReader(contents), int64(len(contents)), s.log.With().Str("component", "bif").Logger())
	if err != nil {
		s.log.Warn().Err(err).Msg("Failed to parse bif")
		resp.ErrorDescription = err.Error()
		return resp, err
	}

	// Calculate sha256 hash
	bifHash, err := util.CalculateSha256(bytes.NewReader(contents))
	if err != nil {
		s.log.Warn().Err(err).Msg("Failed to calculate hash of bif file contents")
		resp.ErrorDescription = err.Error()
		return resp, err
	}

	// Build the model that'll go into the DB
	modelBif := &models.Bif{
		Name:       strings.ToLower(filepath.Base(req.GetFileName())),
		NameInKey:  req.GetFileName(),
		ObjectKey:  &bifHash,
		ObjectHash: &bifHash,
	}

	// Upload the object
	err = s.storage.UploadObject(*modelBif.ObjectKey, bytes.NewReader(contents))
	if err != nil {
		s.log.Warn().Err(err).Msg("Failed to upload bif file to S3 storage")
		resp.ErrorDescription = err.Error()
		return resp, err
	}

	// TODO Add the bif file to the DB

	// TODO Update existing records with contents of bif entries
	for bifIndex, bifEntry := range bif.Files {
		_ = bifIndex
		_ = bifEntry
	}

	return resp, nil
}
