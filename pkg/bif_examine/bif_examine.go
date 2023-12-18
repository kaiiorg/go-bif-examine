package bif_examine

import (
	"context"
	"os"
	"strings"
	"path/filepath"

	"github.com/kaiiorg/go-bif-examine/pkg/bif"
	"github.com/kaiiorg/go-bif-examine/pkg/config"
	"github.com/kaiiorg/go-bif-examine/pkg/rpc"
	"github.com/kaiiorg/go-bif-examine/pkg/storage"
	"github.com/kaiiorg/go-bif-examine/pkg/web"
	"github.com/kaiiorg/go-bif-examine/pkg/models"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type BifExamine struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	config    *config.Config

	log zerolog.Logger

	storage *storage.Storage
	rpc     *rpc.Server
	web     *web.Web
}

func New(conf *config.Config) (*BifExamine, error) {
	s3Storage, err := storage.New(conf, log.With().Str("component", "storage").Logger())
	if err != nil {
		return nil, err
	}
	be := &BifExamine{
		config:  conf,
		log:     log.With().Str("component", "general").Logger(),
		storage: s3Storage,
		rpc:     rpc.New(conf, log.With().Str("component", "rpc").Logger()),
		web:     web.New(conf, log.With().Str("component", "web").Logger()),
	}
	be.ctx, be.ctxCancel = context.WithCancel(context.Background())

	return be, nil
}

func (be *BifExamine) Run() error {
	be.Dev_ProcessKeyAndBifFiles()
	
	err := be.rpc.Run()
	if err != nil {
		return err
	}

	be.web.Run()

	be.log.Info().Msg("Running!")

	return nil
}

func (be *BifExamine) Close() error {
	be.log.Info().Msg("Closing!")
	be.rpc.Close()
	be.web.Close()
	be.ctxCancel()
	return nil
}

func (be *BifExamine) Dev_ProcessKeyAndBifFiles() {
	// Read the key; this will tell us what is in each bif file
	key, err := bif.NewKeyFromFile("./test_bifs/chitin.key", log.With().Str("component", "bif-key").Logger())
	if err != nil {
		be.log.Fatal().Err(err).Msg("Failed to read KEY")
	}
	be.log.Info().
		Str("version", key.Header.VersionToString()).
		Str("signature", key.Header.SignatureToString()).
		Msg("Read KEY")
	
	// Filter out the audio related resources listed in the key and convert that to a model seperate from the in-file model
	audioResources := key.AudioEntriesToModel()
	be.log.Info().Msg("Determined location of audio assets")
	
	relevantBifs := be.Dev_FindRelevantBifs(audioResources)
	be.log.Info().Int("relevantBifsCount", len(relevantBifs)).Msg("Found relevant bif files")

	// Exit now; we're just testing stuff for later use
	os.Exit(0)
}

func (be *BifExamine) Dev_FindRelevantBifs(audioResources map[string]map[uint32]*models.Resource) []string {
	// Get the list of bif files and determine which ones we need to keep for asset extraction and 
	// which ones are ok to delete
	filesInDir, err := os.ReadDir("./test_bifs/data")
	if err != nil {
		be.log.Fatal().Err(err).Msg("Failed to read contents of bif dir")
	}
	bifsOkToDelete := []string{}
	bifsFoundNeeded := []string{}
	bifsFoundNeededMap := map[string]interface{}{}
	for _, dirEntry := range filesInDir{
		entryName := strings.ToLower(dirEntry.Name())
		// Skip anything that isn't a .bif
		if filepath.Ext(entryName) != ".bif" {
			be.log.Trace().Str("found", filepath.Ext(entryName)).Msg("Found a file, but it isn't a .bif file")
			continue
		}
		// Determine if we have an audio resource that relies on this file
		_, found := audioResources[entryName]
		if found {
			bifsFoundNeeded = append(bifsFoundNeeded, dirEntry.Name())
			bifsFoundNeededMap[entryName] = nil
		} else {
			bifsOkToDelete = append(bifsOkToDelete, entryName)
		}
	}
	// Determine if we're missing any bif files
	bifsMissing := []string{}
	resourcesMissing := 0
	for needed, resources := range audioResources {
		_, found := bifsFoundNeededMap[needed]
		if !found {
			resourcesMissing += len(resources)
			bifsMissing = append(bifsMissing, needed)
		}
	}
	be.log.Info().
		Int("okToDeleteCount", len(bifsOkToDelete)).
		Int("neededFound", len(bifsFoundNeeded)).
		Int("expected", len(audioResources)).
		Int("bifsMissingCount", len(bifsMissing)).
		Int("resourcesMissing", resourcesMissing).
		Msg("Scanned bif dir and determined what's needed and what's not")
	
	return bifsFoundNeeded
}