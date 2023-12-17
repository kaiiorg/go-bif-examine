package bif

type KeyBifResourceEntry struct {
	Name            [8]byte
	Type            uint16
	LocatorBitfield uint32
}

const KeyBifResourceEntryLength = 14

func NewKeyBifResourceEntry() *KeyBifResourceEntry {
	return &KeyBifResourceEntry{
		Name: [8]byte{},
	}
}
