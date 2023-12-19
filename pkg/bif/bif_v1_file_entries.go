package bif

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

// BifV1FileEntry describes a file entry in the bif file as defined by IESDP. The order and size of each
// field is very important because the encoding/binary package is used to read and parse the file
// https://gibberlings3.github.io/iesdp/file_formats/ie_formats/bif_v1.htm#bif_v1_FileEntry
type BifV1FileEntry struct {
	NonTileSetIndex uint32
	OffsetToData    uint32
	Size            uint32
	Type            uint16
	UnknownField    uint16
}

func NewBifV1FileEntry() *BifV1FileEntry {
	return &BifV1FileEntry{}
}

func (entry *BifV1FileEntry) Validate(fileSize int64) error {
	if int64(entry.OffsetToData) > fileSize {
		return errors.Join(ErrOffsetToBifDataExceedsFile, fmt.Errorf("%d > %d", entry.OffsetToData, fileSize))
	}

	if int64(entry.OffsetToData)+int64(entry.Size) > fileSize {
		return errors.Join(ErrBifDataExceedsFile, fmt.Errorf("%d + %d > %d", entry.OffsetToData, entry.Size, fileSize))
	}

	return nil
}

func (entry *BifV1FileEntry) ExtractFile(file *os.File) ([]byte, error) {
	_, err := file.Seek(int64(entry.OffsetToData), 0)
	if err != nil {
		return nil, err
	}

	data := make([]byte, entry.Size)
	err = binary.Read(file, binary.LittleEndian, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
