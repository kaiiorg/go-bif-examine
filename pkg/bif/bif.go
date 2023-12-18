package bif

import (
	"os"
	"encoding/binary"

	"github.com/rs/zerolog"
)

type Bif struct {
	Header *BifHeader

	log zerolog.Logger
}

func NewBifFromFile(path string, log zerolog.Logger) (*Bif, error) {
	b := &Bif{
		Header: NewBifHeader(),
		log: log,
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
