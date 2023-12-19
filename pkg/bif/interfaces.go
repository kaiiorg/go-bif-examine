package bif

import "io"

type ReaderSeekerReaderAt interface {
	io.Reader
	io.Seeker
	io.ReaderAt
}
