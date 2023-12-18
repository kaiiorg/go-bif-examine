package bif

import (
	"encoding/binary"
	"path/filepath"
	"strings"
	"os"

	"github.com/kaiiorg/go-bif-examine/pkg/models"

	"github.com/rs/zerolog"
)

type Key struct {
	Header               *KeyHeader
	BifFileNameToEntries map[string]*KeyBifEntry
	BifIndexToFileName   map[uint32]string
	ResourceEntries      []*KeyBifResourceEntry

	log zerolog.Logger
}

func NewKeyFromFile(path string, log zerolog.Logger) (*Key, error) {
	k := &Key{
		Header:               NewKeyHeader(),
		BifFileNameToEntries: map[string]*KeyBifEntry{},
		BifIndexToFileName:   map[uint32]string{},
		ResourceEntries:      []*KeyBifResourceEntry{},
		log:                  log,
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Ask the OS for the stats for this file; we need the file size
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	err = k.readAndValidateHeader(file, fileInfo.Size())
	if err != nil {
		return nil, err
	}

	err = k.readBifFileNameToEntries(file, fileInfo.Size())
	if err != nil {
		return nil, err
	}

	err = k.readResourceEntries(file, fileInfo.Size())
	if err != nil {
		return nil, err
	}

	return k, nil
}

func (k *Key) readAndValidateHeader(file *os.File, fileSize int64) error {
	// Make sure we're at the start of the file
	_, err := file.Seek(0, 0)
	if err != nil {
		return err
	}

	err = binary.Read(file, binary.LittleEndian, k.Header)
	if err != nil {
		return err
	}

	err = k.Header.Validate(fileSize)
	if err != nil {
		return err
	}

	k.log.Debug().Interface("header", k.Header).Msg("Read header from key file")

	return nil
}

func (k *Key) readBifFileNameToEntries(file *os.File, fileSize int64) error {
	// Read each bif entry until we've read all of them or have an error
	for i := uint32(0); i < k.Header.BifEntryCount; i++ {
		// Make sure we're at the start of the bif entries + our current bif offset
		_, err := file.Seek(int64(k.Header.OffsetToBifEntries)+int64(KeyBifEntryLength)*int64(i), 0)
		if err != nil {
			return err
		}
		// Read the entry
		entry := NewKeyBifEntry()
		err = binary.Read(file, binary.LittleEndian, entry)
		if err != nil {
			return err
		}
		err = entry.Validate(fileSize)
		if err != nil {
			k.log.Warn().Err(err).Msg("Failed to validate a bif entry; skipping")
			continue
		}

		// Read the bif file's name from the offset
		filename := make([]byte, entry.FileNameLength-1)
		_, err = file.ReadAt(filename, int64(entry.OffsetToBifFileName))
		if err != nil {
			return err
		}

		// Log a warning if we've already stored this entry
		_, found := k.BifFileNameToEntries[string(filename)]
		if found {
			k.log.Warn().Str("filename", string(filename)).Msg("Found a bif entry that a preview entry already listed; skipping")
			continue
		}
		k.BifFileNameToEntries[string(filename)] = entry
		k.BifIndexToFileName[i] = string(filename)
	}

	k.log.Debug().
		Int("entriesRead", len(k.BifFileNameToEntries)).
		Uint32("expectedToRead", k.Header.BifEntryCount).
		Msg("Read all bif entries from the file")
	return nil
}

func (k *Key) readResourceEntries(file *os.File, fileSize int64) error {
	// Read each resource entry until we've read all of them or have an error
	for i := uint32(0); i < k.Header.ResourceEntryCount; i++ {
		// Make sure we're at the start of the bif entries + our current bif offset
		_, err := file.Seek(int64(k.Header.OffsetToResourceEntries)+int64(KeyBifResourceEntryLength)*int64(i), 0)
		if err != nil {
			return err
		}
		// Read the entry
		entry := NewKeyBifResourceEntry()
		err = binary.Read(file, binary.LittleEndian, entry)
		if err != nil {
			return err
		}
		k.log.Trace().Interface("entry", entry).Msg("Read resource entry from key")

		k.ResourceEntries = append(k.ResourceEntries, entry)
	}

	k.log.Debug().
		Int("entriesRead", len(k.ResourceEntries)).
		Uint32("expectedToRead", k.Header.ResourceEntryCount).
		Msg("Read all resource entries from the file")
	return nil
}

func (k *Key) AudioEntriesToModel() map[string]map[uint32]*models.Resource {
	// bifFileName -> NonTileSetIndex -> details about resource
	modelResources := map[string]map[uint32]*models.Resource{}
	totalAudio := 0
	for _, resource := range k.ResourceEntries {
		// Skip if resource does not describe an audio file
		if models.ResourceType(resource.Type) != models.TYPE_WAV {
			k.log.Trace().
				Str("resourceName", resource.NameToString()).
				Msg("Resource is not WAV type; skipping")
			continue
		}

		// Make sure we know which bif file it lives in
		bifPath, found := k.BifIndexToFileName[resource.BifIndex()]
		if !found {
			k.log.Warn().
				Uint32("calculatedIndex", resource.BifIndex()).
				Str("resourceName", resource.NameToString()).
				Msg("Did not find a bif file by the given index for the given resource")
			continue
		}

		totalAudio++

		// Store it for later
		modelResource := &models.Resource{
			Name: strings.ToLower(resource.NameToString()),
			BifPath: strings.ToLower(filepath.Base(bifPath)),
			Type: resource.Type,
			TileSetIndex: resource.TileSetIndex(),
			NonTileSetIndex: resource.NonTileSetIndex(),
			BifIndex: resource.BifIndex(),
		}

		byIndex, found := modelResources[modelResource.BifPath]
		if !found {
			byIndex = map[uint32]*models.Resource{} 
		}
		_, found = modelResources[modelResource.BifPath][modelResource.NonTileSetIndex]
		if found {
			k.log.Warn().
				Uint32("index", modelResource.NonTileSetIndex).
				Str("bifPath", modelResource.BifPath).
				Str("resource", modelResource.Name).
				Msg("Already found and stored a resource with the given NonTileSetIndex in the given bif file; skipping")
				continue
		}
		byIndex[modelResource.NonTileSetIndex] = modelResource
		modelResources[modelResource.BifPath] = byIndex
	}

	totalMapped := 0
	for _, byFile := range modelResources{
		totalMapped += len(byFile)
	}

	k.log.Info().
		Int("totalResources", len(k.ResourceEntries)).
		Int("totalAudioResources", totalAudio).
		Int("totalMappedByFile", totalMapped).
		Int("filesMapped", len(modelResources)).
		Msg("Determined audio resources")

	return modelResources
}
