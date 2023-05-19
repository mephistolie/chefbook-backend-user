package config

import (
	"github.com/mephistolie/chefbook-backend-common/log"
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
	log.Infof("USER SERVICE CONFIGURATION\n"+
		"Environment: %v\n"+
		"Port: %v\n"+
		"Logs path: %v\n\n"+
		"Database host: %v\n"+
		"Database port: %v\n"+
		"Database name: %v\n\n"+
		"S3 host: %v\n"+
		"S3 bucket: %v\n\n"+
		"MQ host: %v\n"+
		"MQ port: %v\n"+
		"MQ vhost: %v\n\n",
		*c.Environment, *c.Port, *c.LogsPath,
		*c.Database.Host, *c.Database.Port, *c.Database.DBName,
		*c.S3.Host, *c.S3.Bucket,
		*c.Amqp.Host, *c.Amqp.Port, *c.Amqp.VHost,
	)
}
