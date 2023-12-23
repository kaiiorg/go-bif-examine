package bif

import (
	"errors"
	"fmt"
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

func NewKeyHeader() *KeyHeader {
	return &KeyHeader{
		Signature: [4]byte{},
		Version:   [4]byte{},
	}
}

func (kh *KeyHeader) Validate(fileSize int64) error {
	_, found := supportedKeySignatures[kh.SignatureToString()]
	if !found {
		return errors.Join(ErrUnexpectedSignature, fmt.Errorf("got %s", kh.SignatureToString()))
	}

	_, found = supportedKeyVersions[kh.VersionToString()]
	if !found {
		return errors.Join(ErrUnexpectedVersion, fmt.Errorf("got %s", kh.VersionToString()))
	}

	// Make sure the file tells us a length that is not more than what we're expecting. The mismatch accounts for
	// the path to each bif file, which isn't stored with each entry
	// Header size + bif entry count * bif entry size + resource entry count * resource entry size
	expectedFileSize := int64(KeyHeaderLength) + int64(kh.BifEntryCount)*int64(KeyBifEntryLength) + int64(kh.ResourceEntryCount)*int64(KeyBifResourceEntryLength)
	if expectedFileSize > fileSize {
		return errors.Join(ErrUnexpectedFileLength, fmt.Errorf("os reports %d, calculated %d from file; should be os >= calculated", fileSize, expectedFileSize))
	}

	if int64(kh.OffsetToBifEntries) > fileSize {
		return errors.Join(ErrBifEntryOffsetExceedsFile, fmt.Errorf("file offset = %d, which is beyond the file size of %d reported by the os", kh.OffsetToBifEntries, fileSize))
	}

	if int64(kh.OffsetToResourceEntries) > fileSize {
		return errors.Join(ErrBifEntryResourceOffsetExceedsFile, fmt.Errorf("file offset = %d, which is beyond the file size of %d reported by the os", kh.OffsetToResourceEntries, fileSize))
	}

	return nil
}

func (kh *KeyHeader) SignatureToString() string {
	return string(kh.Signature[:])
}

func (kh *KeyHeader) VersionToString() string {
	return string(kh.Version[:])
}
