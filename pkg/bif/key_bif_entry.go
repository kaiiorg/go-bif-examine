package bif

import (
	"errors"
	"fmt"
)

var (
	ErrOffsetToBifFileNameExceedsFile = errors.New("bif entry gives offset to filename beyond the file")
	ErrUnexpectedNameLength           = errors.New("bif entry gives unexpected length of filename")
)

// KeyBifEntry describes a bif entry in the key file as defined by IESDP. The order and size of each
// field is very important because the encoding/binary package is used to read and parse the file
// https://gibberlings3.github.io/iesdp/file_formats/ie_formats/key_v1.htm#keyv1_BifIndices
type KeyBifEntry struct {
	FileLength          uint32
	OffsetToBifFileName uint32
	FileNameLength      uint16
	DetailsBitfield     uint16
}

func NewKeyBifEntry() *KeyBifEntry {
	return &KeyBifEntry{}
}

func (kbe *KeyBifEntry) Validate(fileSize int64) error {
	if int64(kbe.OffsetToBifFileName) > fileSize {
		return errors.Join(ErrOffsetToBifFileNameExceedsFile, fmt.Errorf("%d > %d", kbe.OffsetToBifFileName, fileSize))
	}
	if int(kbe.FileNameLength) > MaxFileNameLength {
		return errors.Join(ErrUnexpectedNameLength, fmt.Errorf("%d > %d", kbe.FileNameLength, MaxFileNameLength))
	}
	return nil
}
