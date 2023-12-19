package bif

import (
	"encoding/binary"
	"os"

	"github.com/rs/zerolog"
)

type Bif struct {
	Header *BifHeader
	Files  map[uint32]*BifV1FileEntry

	log zerolog.Logger
}

func NewBifFromFile(path string, log zerolog.Logger) (*Bif, error) {
	b := &Bif{
		Header: NewBifHeader(),
		Files:  map[uint32]*BifV1FileEntry{},
		log:    log,
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

	err = b.readAndValidateHeader(file, fileInfo.Size())
	if err != nil {
		return nil, err
	}

	err = b.readAndValidateBifEntries(file, fileInfo.Size())
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (b *Bif) readAndValidateHeader(file *os.File, fileSize int64) error {
	// Make sure we're at the start of the file
	_, err := file.Seek(0, 0)
	if err != nil {
		return err
	}

	err = binary.Read(file, binary.LittleEndian, b.Header)
	if err != nil {
		return err
	}

	err = b.Header.Validate(fileSize)
	if err != nil {
		return err
	}

	b.log.Debug().Interface("header", b.Header).Msg("Read header from bif file")

	return nil
}

func (b *Bif) readAndValidateBifEntries(file *os.File, fileSize int64) error {
	for i := uint32(0); i < b.Header.FileEntryCount; i++ {
		// Make sure we're at the start of the bif entries + our current bif offset
		_, err := file.Seek(int64(b.Header.OffsetToFileEntries)+int64(BifV1EntryLength)*int64(i), 0)
		if err != nil {
			return err
		}
		// Read the entry
		entry := NewBifV1FileEntry()
		err = binary.Read(file, binary.LittleEndian, entry)
		if err != nil {
			return err
		}
		err = entry.Validate(fileSize)
		if err != nil {
			b.log.Warn().Err(err).Msg("Failed to validate a bif file entry; skipping")
			continue
		}

		// Log a warning if we've already stored this entry
		_, found := b.Files[entry.NonTileSetIndex]
		if found {
			b.log.Warn().Uint32("nonTileSetIndex", entry.NonTileSetIndex).Msg("Found a bif file entry that a previous entry already listed; skipping")
			continue
		}
		b.Files[entry.NonTileSetIndex] = entry
	}

	b.log.Debug().
		Int("entriesRead", len(b.Files)).
		Uint32("expectedToRead", b.Header.FileEntryCount).
		Msg("Read all bif file entries from the file")

	return nil
}
