package bif

import (
	"errors"
)

const (
	MaxFileNameLength = 1024

	KeyHeaderLength           = 24
	KeyBifResourceEntryLength = 14
	KeyBifEntryLength         = 12

	BifHeaderLength         = 20
	BifV1EntryLength        = 16
	BifV1TilesetEntryLength = 20
)

var (
	ErrUnexpectedSignature               = errors.New("unexpected signature")
	ErrUnexpectedVersion                 = errors.New("unexpected version")
	ErrUnexpectedFileLength              = errors.New("unexpected file length")
	ErrBifEntryOffsetExceedsFile         = errors.New("the offset for the start of the bif entries is beyond the length of the file")
	ErrBifEntryResourceOffsetExceedsFile = errors.New("the offset for the start of the bif resource entries is beyond the length of the file")
	ErrOffsetToBifDataExceedsFile        = errors.New("the offset for the start of the bif data is beyond the length of the file")
	ErrBifDataExceedsFile                = errors.New("the offset+size of the data described by the entry exceeds the length of the file")
)

var (
	supportedKeySignatures = map[string]interface{}{
		"KEY ": nil,
	}
	supportedKeyVersions = map[string]interface{}{
		"V1  ": nil,
	}
)

var (
	supportedBifSignatures = map[string]interface{}{
		"BIFF": nil,
	}
	supportedBifVersions = map[string]interface{}{
		"V1  ": nil,
	}
)
