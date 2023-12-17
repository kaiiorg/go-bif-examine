package config

type S3 struct {
	Host string `hcl:"host"`
	Port int    `hcl:"port"`

	// It'd be better to use something like Hashicorp's Vault for these secrets
	// but that's beyond the scope of this hobby project
	AccessKey string `hcl:"access_key"`
	SecretKey string `hcl:"secret_key"`
}
