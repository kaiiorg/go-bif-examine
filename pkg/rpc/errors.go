package rpc

import (
	"errors"
)

var (
	ErrMustProvideFilenameOrNameInKey = errors.New("must provide either the normalized filename or the exact string listed in the key file")
	ErrBifNotYetUploaded              = errors.New("bif file for the resource has not yet been uploaded")
)
