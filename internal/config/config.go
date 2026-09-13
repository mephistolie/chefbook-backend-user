package config

import (
	"context"
	"github.com/mephistolie/chefbook-backend-user/internal/logging"
)

const (
	EnvDev  = "develop"
	EnvProd = "production"
)

type Config struct {
	Environment *string
	Port        *int
	LogsPath    *string

	Firebase Firebase
	Database Database
	S3       S3
	Amqp     Amqp
}

type Firebase struct {
	Credentials *string
}

type Database struct {
	Host     *string
	Port     *int
	User     *string
	Password *string
	DBName   *string
}

type S3 struct {
	Host            *string
	AccessKeyId     *string
	SecretAccessKey *string
	Bucket          *string
	Region          *string
}

type Amqp struct {
	Host     *string
	Port     *int
	User     *string
	Password *string
	VHost    *string
}

func (c Config) Validate() error {
	if *c.Environment != EnvProd {
		*c.Environment = EnvDev
	}
	return nil
}

func (c Config) Print() {
	logging.Events{}.ConfigLoaded(context.Background())
}
