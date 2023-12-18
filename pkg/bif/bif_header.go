package bif

import (
	"errors"
	"fmt"
)

// BifHeader is the header in the bif file as defined by IESDP. The order and size of each
// field is very important because the encoding/binary package is used to read and parse the file
// https://gibberlings3.github.io/iesdp/file_formats/ie_formats/key_v1.htm#keyv1_Header
type BifHeader struct {
	Signature           [4]byte
	Version             [4]byte
	FileEntryCount      uint32
	TilesetCount        uint32
	OffsetToFileEntries uint32
}

func NewBifHeader() *BifHeader {
	return &BifHeader{
		Signature: [4]byte{},
		Version:   [4]byte{},
	}
}

func (bh *BifHeader) Validate(fileSize int64) error {
	_, found := supportedBifSignatures[bh.SignatureToString()]
	if !found {
		return errors.Join(ErrUnexpectedSignature, fmt.Errorf("got %s", bh.SignatureToString()))
	}

	_, found = supportedBifVersions[bh.VersionToString()]
	if !found {
		return errors.Join(ErrUnexpectedVersion, fmt.Errorf("got %s", bh.VersionToString()))
	}

	// Make sure the file tells us a length that is not more than what we're expecting. The mismatch accounts for
	// the actual data contained in the file
	// Header size + bif file entry count * bif file entry size + tileset entry count * tileset entry size
	expectedFileSize := int64(BifHeaderLength) + int64(bh.FileEntryCount)*int64(BifV1EntryLength) + int64(bh.TilesetCount)*int64(BifV1TilesetEntryLength)
	if expectedFileSize > fileSize {
		return errors.Join(ErrUnexpectedFileLength, fmt.Errorf("os reports %d, calculated %d from file; should be os >= calculated", fileSize, expectedFileSize))
	}

	return nil
}

func (bh *BifHeader) SignatureToString() string {
	return string(bh.Signature[:])
}

func (bh *BifHeader) VersionToString() string {
	return string(bh.Version[:])
}
