package models

import (
	"gorm.io/gorm"
)

type Resource struct {
	gorm.Model
	Name    string `gorm:"not null"`
	BifPath string `gorm:"not null"`
	Type    uint16 `gorm:"not null"`
}

func NewResource(Type uint16, name, bifPath string) *Resource {
	return &Resource{
		Name:    name,
		BifPath: bifPath,
		Type:    Type,
	}
}
