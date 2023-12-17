package config

import (
	"fmt"
)

type Database struct {
	DbName   string `hcl:"db_name"`
	Host     string `hcl:"host"`
	Port     int    `hcl:"port"`
	Username string `hcl:"username"`
	Password string `hcl:"password"`
	Ssl      bool   `hcl:"ssl"`
}

func (db *Database) ConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=UTC",
		db.Host,
		db.Port,
		db.Username,
		db.Password,
		db.DbName,
		db.Ssl,
	)
}
