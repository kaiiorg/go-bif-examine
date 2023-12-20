package models

import "gorm.io/gorm"

type Bif struct {
	gorm.Model
	// Name is the normalized filename (lower case, base name, .bif)
	Name string `gorm:"not null"`
	// NameInKey is exactly as the file is defined in the key file
	NameInKey string `gorm:"not null"`
	// ObjectKey is the S3 key needed to retrieve this file, if it has been uploaded
	ObjectKey *string
	// ObjectHash is the SHA256 hash of the file
	ObjectHash *string
}
