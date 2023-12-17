package storage

import (
	"fmt"
	"net/url"

	"github.com/kaiiorg/go-bif-examine/pkg/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/rs/zerolog"
)

type Storage struct {
	config *config.Config
	log    zerolog.Logger

	s3Session *session.Session
	s3Client *s3.S3
}

func New(conf *config.Config, log zerolog.Logger) (*Storage, error) {
	u := &url.URL{
		Scheme: conf.S3.Scheme,
		Host:   fmt.Sprintf("%s:%d", conf.S3.Host, conf.S3.Port),
	}

	s3Session, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(u.String()),
		Region:   aws.String(conf.S3.Region),
		Credentials: credentials.NewStaticCredentials(
			conf.S3.AccessKey,
			conf.S3.SecretKey,
			"",
		),
	})
	if err != nil {
		return nil, err
	}

	s := &Storage{
		config:    conf,
		log:       log,
		s3Session: s3Session,
		s3Client:  s3.New(s3Session),
	}

	return s, nil
}
