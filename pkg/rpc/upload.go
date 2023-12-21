package rpc

import (
	"bytes"
	"context"
	"errors"
	"time"

	"github.com/kaiiorg/go-bif-examine/pkg/bif"
	"github.com/kaiiorg/go-bif-examine/pkg/models"
	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"
	"github.com/kaiiorg/go-bif-examine/pkg/util"
)

var (
	ErrMustProvideFilenameOrNameInKey = errors.New("must provide either the normalized filename or the exact string listed in the key file")
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

	// Make sure that either the file name or the name in key were set
	if req.GetFileName() == "" && req.GetNameInKey() == "" {
		resp.ErrorDescription = ErrMustProvideFilenameOrNameInKey.Error()
		return resp, ErrMustProvideFilenameOrNameInKey
	}

	// Find the project related to this file
	project, err := s.examineRepository.GetProjectById(uint(req.GetProjectId()))
	if err != nil {
		s.log.Warn().Err(err).Msg("Failed to find project record related to this bif")
		resp.ErrorDescription = err.Error()
		return resp, err
	}

	// Find the existing bif record related to this file
	modelBif, err := s.examineRepository.GetBifByNormalizedNameOrNameInKey(req.GetFileName(), req.GetNameInKey())
	if err != nil {
		s.log.Warn().Err(err).Msg("Failed to find bif record")
		resp.ErrorDescription = err.Error()
		return resp, err
	}

	// Find the resources that are expected to live in this file
	relatedResources, err := s.examineRepository.FindProjectResourcesForBif(project.ID, modelBif.ID)
	if err != nil {
		resp.ErrorDescription = err.Error()
		return resp, err
	}

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

	// Update contents of the existing bif model
	modelBif.ObjectKey = &bifHash
	modelBif.ObjectHash = &bifHash

	// Upload the object
	err = s.storage.UploadObject(*modelBif.ObjectKey, bytes.NewReader(contents))
	if err != nil {
		s.log.Warn().Err(err).Msg("Failed to upload bif file to S3 storage")
		resp.ErrorDescription = err.Error()
		return resp, err
	}

	// Update the existing record
	err = s.examineRepository.UpdateBif(modelBif)
	if err != nil {
		s.log.Warn().Err(err).Msg("Failed to update bif record")
		resp.ErrorDescription = err.Error()
		return resp, err
	}

	for _, resource := range relatedResources {
		// Check if the bif contains the resource the key claims that is there
		bifEntry, found := bif.Files[resource.NonTileSetIndex]
		if !found {
			resp.ResourcesNotFound++
			continue
		}
		resp.ResourcesFound++

		// Update our resource record
		resource.OffsetToData = bifEntry.OffsetToData
		resource.Size = bifEntry.Size
		err = s.examineRepository.UpdateResource(resource)
		if err != nil {
			s.log.Warn().Err(err).Msg("Failed to update resource record")
			resp.ErrorDescription = err.Error()
			return resp, err
		}
	}

	return resp, nil
}
