package rpc

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/kaiiorg/go-bif-examine/pkg/bif"
	"github.com/kaiiorg/go-bif-examine/pkg/models"
	"io"
	"os"

	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"
)

func (s *Server) saveBifStreamToFile(stream pb.BifExamine_UploadBifServer) (*pb.UploadBifRequest, string, string, error) {
	// Create a temp file to dump the bytes to
	file, err := os.CreateTemp("", "go-bif-examine")
	if err != nil {
		return nil, "", "", err
	}
	defer file.Close()

	var returnReq *pb.UploadBifRequest
	hash := sha256.New()

	// Receive and dump the bytes to file
	for {
		req, err := stream.Recv()
		if err != nil {
			// Break out of the loop if the error result is EOF
			if errors.Is(err, io.EOF) {
				break
			}
			// Return an error if it was something else
			os.Remove(file.Name())
			return nil, "", "", err
		}
		returnReq = req

		// This will break if we ever receive chunks out of order
		hash.Write(req.GetContents())

		// Write this chunk at the location specified
		_, err = file.WriteAt(req.GetContents(), req.GetOffset())
		if err != nil {
			os.Remove(file.Name())
			return nil, "", "", err
		}
	}

	return returnReq, file.Name(), hex.EncodeToString(hash.Sum(nil)), nil
}

func (s *Server) processBif(tempBifPath, bifHash string, req *pb.UploadBifRequest, resp *pb.UploadBifResponse) error {
	// Make sure that either the file name or the name in key were set
	if req.GetFileName() == "" && req.GetNameInKey() == "" {
		resp.ErrorDescription = ErrMustProvideFilenameOrNameInKey.Error()
		return ErrMustProvideFilenameOrNameInKey
	}

	// Find the project related to this file
	project, err := s.examineRepository.GetProjectById(uint(req.GetProjectId()))
	if err != nil {
		s.log.Warn().Err(err).Msg("Failed to find project record related to this bif")
		resp.ErrorDescription = err.Error()
		return err
	}

	// Find the existing bif record related to this file
	modelBif, err := s.examineRepository.GetBifByNormalizedNameOrNameInKey(req.GetFileName(), req.GetNameInKey())
	if err != nil {
		s.log.Warn().Err(err).Msg("Failed to find bif record")
		resp.ErrorDescription = err.Error()
		return err
	}

	// Find the resources that are expected to live in this file
	relatedResources, err := s.examineRepository.FindProjectResourcesForBif(project.ID, modelBif.ID)
	if err != nil {
		resp.ErrorDescription = err.Error()
		return err
	}

	// Parse the contents as a bif file
	bif, err := bif.NewBifFromFile(tempBifPath, s.log.With().Str("component", "bif").Logger())
	if err != nil {
		s.log.Warn().Err(err).Msg("Failed to parse bif")
		resp.ErrorDescription = err.Error()
		return err
	}

	// Update contents of the existing bif model
	modelBif.ObjectKey = &bifHash
	modelBif.ObjectHash = &bifHash

	// Upload the object
	err = s.storage.UploadObjectFromTempFile(*modelBif.ObjectKey, tempBifPath)
	if err != nil {
		s.log.Warn().Err(err).Msg("Failed to upload bif file to S3 storage")
		resp.ErrorDescription = err.Error()
		return err
	}

	// Update the existing record
	err = s.examineRepository.UpdateBif(modelBif)
	if err != nil {
		s.log.Warn().Err(err).Msg("Failed to update bif record")
		resp.ErrorDescription = err.Error()
		return err
	}

	err = s.updateRelatedResources(resp, bif, relatedResources)
	if err != nil {
		s.log.Warn().Err(err).Msg("Failed to update related resource records")
		resp.ErrorDescription = err.Error()
		return err
	}

	return nil
}

func (s *Server) updateRelatedResources(resp *pb.UploadBifResponse, bif *bif.Bif, relatedResources []*models.Resource) error {
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
		err := s.examineRepository.UpdateResource(resource)
		if err != nil {
			s.log.Warn().Err(err).Msg("Failed to update resource record")
			resp.ErrorDescription = err.Error()
			return err
		}
	}
	return nil
}
