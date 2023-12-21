package bif_examine

import (
	"context"
	"github.com/kaiiorg/go-bif-examine/pkg/repositories/gorm_logger"
	"os"
	"path/filepath"
	"strings"

	"github.com/kaiiorg/go-bif-examine/pkg/config"
	"github.com/kaiiorg/go-bif-examine/pkg/models"
	"github.com/kaiiorg/go-bif-examine/pkg/repositories/examine_repository"
	"github.com/kaiiorg/go-bif-examine/pkg/rpc"
	"github.com/kaiiorg/go-bif-examine/pkg/storage"
	"github.com/kaiiorg/go-bif-examine/pkg/web"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type BifExamine struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	config    *config.Config

	log zerolog.Logger

	storage           storage.BifStorage
	rpc               *rpc.Server
	web               *web.Web
	examineRepository examine_repository.ExamineRepository
}

func New(conf *config.Config) (*BifExamine, error) {
	// Setup S3 storage
	s3Storage, err := storage.New(conf, log.With().Str("component", "storage").Logger())
	if err != nil {
		return nil, err
	}
	// Setup DB and repository
	db, err := gorm.Open(
		postgres.Open(conf.Db.ConnectionString()),
		&gorm.Config{
			Logger: gorm_logger.NewGormLogger(log.With().Str("component", "db-gorm").Logger(), zerolog.GlobalLevel()),
		},
	)
	if err != nil {
		return nil, err
	}
	examineRepository, err := examine_repository.New(conf, db, log.With().Str("component", "db").Logger())
	if err != nil {
		return nil, err
	}

	// Setup main struct along with the rpc and web components
	be := &BifExamine{
		config:            conf,
		log:               log.With().Str("component", "general").Logger(),
		storage:           s3Storage,
		rpc:               rpc.New(examineRepository, s3Storage, conf, log.With().Str("component", "rpc").Logger()),
		web:               web.New(conf, log.With().Str("component", "web").Logger()),
		examineRepository: examineRepository,
	}
	be.ctx, be.ctxCancel = context.WithCancel(context.Background())

	return be, nil
}

func (be *BifExamine) Run() error {
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

func (be *BifExamine) Dev_UploadBifFiles() {
	/*
		bifDir := "./test_bifs/data"
		filesInDir, err := os.ReadDir(bifDir)
		if err != nil {
			be.log.Fatal().Err(err).Msg("Failed to read contents of bif dir")
		}

		for _, dirEntry := range filesInDir {
			// Skip anything that isn't a .bif
			if filepath.Ext(dirEntry.Name()) != ".bif" {
				be.log.Trace().Str("found", filepath.Ext(dirEntry.Name())).Msg("Found a file, but it isn't a .bif file")
				continue
			}
			bifPath := filepath.Join(bifDir, dirEntry.Name())

			// Calculate the SHA256 hash of the file
			hash, err := sha256OfFile(bifPath)
			if err != nil {
				be.log.Fatal().Err(err).Msg("Failed to calculate hash of bif")
			}
			// Upload to s3
			err = be.storage.UploadFileFromTempFile(hash, bifPath)
			if err != nil {
				be.log.Fatal().Err(err).Msg("Failed to upload bif")
			}
		}
	*/
}

func (be *BifExamine) Dev_GetSectionOfFile() {
	hash := "6dcba28d08fac8c8a4a9c6bc8c3cd0721dbe5c55da0a37b016396f30d9cbb16a"
	offsetToData := uint32(454392)
	size := uint32(224648)
	data, err := be.storage.GetSectionFromObject(hash, offsetToData, size)
	if err != nil {
		be.log.Fatal().Err(err).Msg("Failed to get section from object")
	}
	err = os.WriteFile("./output.wav", data, 0666)
	if err != nil {
		be.log.Fatal().Err(err).Msg("Failed to get write data to file")
	}
	be.log.Info().Msg("Wrote section from object to file")
}

/*
func (be *BifExamine) Dev_ProcessKeyAndBifFiles() {
	// Read the key; this will tell us what is in each bif file
	key, err := bif.NewKeyFromFile("./test_bifs/chitin.key", be.log.With().Str("component", "bif-key").Logger())
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

	// Find the relevant bif files
	relevantBifs := be.Dev_FindRelevantBifs(audioResources)
	be.log.Info().Int("relevantBifsCount", len(relevantBifs)).Msg("Found relevant bif files")

	// Open each relevant bif, calculate its sha265 hash, and scan for the file entries.
	// We're not extractings the audio files; we just need to know the offset from the
	// start of the file and the size of the resource
	for bifName := range relevantBifs {
		bifPath := filepath.Join("./test_bifs/data", bifName)
		// Calculating the sha256 will later be done while uploading the file to S3, not when we get to this processing step
		hash, err := sha256OfFile(bifPath)
		if err != nil {
			be.log.Fatal().Err(err).Msg("Failed to calculate the sha256 hash of the bif")
		}

		// Get the resources that are relevant to this bif
		_, found := audioResources[strings.ToLower(bifName)]
		if !found {
			be.log.Fatal().Str("bifName", bifName).Msg("Did not find any audio resources for a bif file that was previously determined to be relevant")
		}

		// Parse the bif
		bif, err := bif.NewBifFromFile(bifPath, be.log.With().Str("component", "bif").Logger())
		be.log.Info().
			Str("filepath", bifPath).
			Str("version", bif.Header.VersionToString()).
			Str("signature", bif.Header.SignatureToString()).
			Str("hash", hash).
			Interface("bif", bif).
			Msg("Read bif")
		os.Exit(0)
	}

	// Exit now; we're just testing stuff for later use
	os.Exit(0)
}
*/

func (be *BifExamine) Dev_FindRelevantBifs(audioResources map[string]map[uint32]*models.Resource) map[string]interface{} {
	// Get the list of bif files and determine which ones we need to keep for asset extraction and
	// which ones are ok to delete
	bifDir := "./test_bifs/data"
	filesInDir, err := os.ReadDir(bifDir)
	if err != nil {
		be.log.Fatal().Err(err).Msg("Failed to read contents of bif dir")
	}
	bifsOkToDelete := []string{}
	bifsFoundNeededMap := map[string]interface{}{}
	for _, dirEntry := range filesInDir {
		entryName := strings.ToLower(dirEntry.Name())
		// Skip anything that isn't a .bif
		if filepath.Ext(entryName) != ".bif" {
			be.log.Trace().Str("found", filepath.Ext(entryName)).Msg("Found a file, but it isn't a .bif file")
			continue
		}
		// Determine if we have an audio resource that relies on this file
		_, found := audioResources[entryName]
		if found {
			bifsFoundNeededMap[dirEntry.Name()] = nil
		} else {
			bifsOkToDelete = append(bifsOkToDelete, dirEntry.Name())
		}
	}
	return bifsFoundNeededMap
}
