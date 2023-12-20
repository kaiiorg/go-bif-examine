package models

import (
	"gorm.io/gorm"
)

type Resource struct {
	gorm.Model

	// Bif is which file the key claims this bif exists in
	Bif   *Bif `gorm:"-"`
	BifID uint `gorm:"joinForeignKey:bif"`
	// Project is which project this is resource is part of
	Project   *Project `gorm:"-"`
	ProjectID uint     `gorm:"joinForeignKey:project"`

	// Name is the name of the resource according to the key
	Name            string
	TileSetIndex    uint32
	NonTileSetIndex uint32
	BifIndex        uint32

	// OffsetToData is where the data for this resource in the bif file start
	OffsetToData uint32
	// Size is how large this resource is
	Size uint32
}
