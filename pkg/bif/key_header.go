package bif

import (
	"errors"
	"fmt"
)

var (
	validSignatures = map[string]interface{}{
		"KEY ": nil,
	}
	validVersions = map[string]interface{}{
		"V1  ": nil,
	}
)

var (
	ErrUnexpectedSignature               = errors.New("unexpected signature")
	ErrUnexpectedVersion                 = errors.New("unexpected version")
	ErrUnexpectedFileLength              = errors.New("unexpected file length")
	ErrBifEntryOffsetExceedsFile         = errors.New("the offset for the start of the bif entries is beyond the length of the file")
	ErrBifEntryResourceOffsetExceedsFile = errors.New("the offset for the start of the bif resource entries is beyond the length of the file")
)

// KeyHeader is the header in the key file as defined by IESDP. The order and size of each
// field is very important because the encoding/binary package is used to read and parse the file
// https://gibberlings3.github.io/iesdp/file_formats/ie_formats/key_v1.htm#keyv1_Header
type KeyHeader struct {
	Signature               [4]byte
	Version                 [4]byte
	BifEntryCount           uint32
	ResourceEntryCount      uint32
	OffsetToBifEntries      uint32
	OffsetToResourceEntries uint32
}

const KeyHeaderLength = 24

func NewKeyHeader() *KeyHeader {
	return &KeyHeader{
		Signature: [4]byte{},
		Version:   [4]byte{},
	}
}

func (kh *KeyHeader) Validate(fileSize int64) error {
	_, found := validSignatures[kh.SignatureToString()]
	if !found {
		return errors.Join(ErrUnexpectedSignature, fmt.Errorf("got %s", kh.SignatureToString()))
	}

	_, found = validVersions[kh.VersionToString()]
	if !found {
		return errors.Join(ErrUnexpectedVersion, fmt.Errorf("got %s", kh.VersionToString()))
	}

	// Make sure the file tells us the length we're actually expected
	// Header size + bif entry count * bif entry size + resource entry count * resource entry size
	expectedFileSize := int64(KeyHeaderLength) + int64(kh.BifEntryCount)*int64(KeyBifEntryLength) + int64(kh.ResourceEntryCount)*int64(KeyBifResourceEntryLength)
	if expectedFileSize > fileSize {
		return errors.Join(ErrUnexpectedFileLength, fmt.Errorf("os reports %d, calculated %d from file; should be os >= calculated", fileSize, expectedFileSize))
	}

	if int64(kh.OffsetToBifEntries) > fileSize {
		return errors.Join(ErrBifEntryOffsetExceedsFile, fmt.Errorf("file offset = %s, which is beyond the file size of %d reported by the os", kh.OffsetToBifEntries, fileSize))
	}

	if int64(kh.OffsetToResourceEntries) > fileSize {
		return errors.Join(ErrBifEntryResourceOffsetExceedsFile, fmt.Errorf("file offset = %s, which is beyond the file size of %d reported by the os", kh.OffsetToResourceEntries, fileSize))
	}

	return nil
}

func (kh *KeyHeader) SignatureToString() string {
	return string(kh.Signature[:])
}

func (kh *KeyHeader) VersionToString() string {
	return string(kh.Version[:])
}
