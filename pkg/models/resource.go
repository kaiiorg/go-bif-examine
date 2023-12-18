package models

import (
	"gorm.io/gorm"
)

type Resource struct {
	gorm.Model
	Name    string
	BifPath string
	Type    uint16

	TileSetIndex uint32
	NonTileSetIndex uint32
	BifIndex uint32
}
