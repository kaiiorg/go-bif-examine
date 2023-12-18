package bif

import (
	"strings"
)

type KeyBifResourceEntry struct {
	Name            [8]byte
	Type            uint16
	LocatorBitfield uint32
}

func NewKeyBifResourceEntry() *KeyBifResourceEntry {
	return &KeyBifResourceEntry{
		Name: [8]byte{},
	}
}

func (kbre *KeyBifResourceEntry) NameToString() string {
	return strings.Trim(string(kbre.Name[:]), "\u0000")
}

func (kbre *KeyBifResourceEntry) NonTileSetIndex() uint32 {
	return (0x3FFFF & kbre.LocatorBitfield)
}

func (kbre *KeyBifResourceEntry) TileSetIndex() uint32 {
	return (0xFC000 & kbre.LocatorBitfield) >> 14
}

func (kbre *KeyBifResourceEntry) BifIndex() uint32 {
	return (0xFFF00000 & kbre.LocatorBitfield) >> 20
}
