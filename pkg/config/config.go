package config

import (
	"github.com/hashicorp/hcl/v2/hclsimple"
)

type Config struct {
	Db *Database `hcl:"db,block"`
	S3 *S3       `hcl:"s3,block"`
}

func LoadFromFile(filepath string) (*Config, error) {
	config := &Config{}

	err := hclsimple.DecodeFile(filepath, nil, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
