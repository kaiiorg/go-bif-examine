package storage

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"os"
	"time"

	"github.com/kaiiorg/go-bif-examine/pkg/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	s3Credentials "github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/zerolog"
)

type Storage struct {
	config *config.Config
	log    zerolog.Logger

	s3Client        *s3.Client
	s3PresignClient *s3.PresignClient
}

func New(conf *config.Config, log zerolog.Logger) (*Storage, error) {
	u := &url.URL{
		Scheme: conf.S3.Scheme,
		Host:   fmt.Sprintf("%s:%d", conf.S3.Host, conf.S3.Port),
	}

	s := &Storage{
		config: conf,
		log:    log,
		s3Client: s3.New(
			s3.Options{
				BaseEndpoint: aws.String(u.String()),
				Region:       conf.S3.Region,
				UsePathStyle: conf.S3.ForcePathStyle,
				Credentials: s3Credentials.NewStaticCredentialsProvider(
					conf.S3.AccessKey,
					conf.S3.SecretKey,
					"",
				),
			},
		),
	}
	s.s3PresignClient = s3.NewPresignClient(s.s3Client)

	return s, nil
}

// UploadObjectFromTempFile uploads the contents of the given file to S3 compatible storage
func (s *Storage) UploadObjectFromTempFile(objectKey, pathToTempBif string) error {
	file, err := os.Open(pathToTempBif)
	if err != nil {
		return err
	}
	defer file.Close()

	return s.UploadObject(objectKey, file)
}

// UploadObject uploads the contents of the given reader to S3 compatible storage
func (s *Storage) UploadObject(objectKey string, reader io.ReadSeeker) error {
	// Create s3 request to put the data in the configured bucket
	putObjectInput := &s3.PutObjectInput{
		Bucket: aws.String(s.config.S3.Bucket),
		Key:    aws.String(objectKey),
		Body:   reader,
	}

	// Make the request
	_, err := s.s3Client.PutObject(context.Background(), putObjectInput)
	if err != nil {
		s.log.Error().
			Err(err).
			Str("bucket", s.config.S3.Bucket).
			Str("objectKey", objectKey).
			Msg("Failed to upload object")
		return err
	}
	return nil
}

// GetSectionFromObject pulls a section of data from the given object instead of the full object
func (s *Storage) GetSectionFromObject(objectKey string, offsetFromFileStart, size uint32) ([]byte, error) {
	// Setup a GetObject request to pull from the key from the configured bucket and to only give us the contents we want
	input := &s3.GetObjectInput{
		Bucket: aws.String(s.config.S3.Bucket),
		Key:    aws.String(objectKey),
		Range:  aws.String(fmt.Sprintf("bytes=%d-%d", offsetFromFileStart, offsetFromFileStart+size)),
	}

	// Make the request
	result, err := s.s3Client.GetObject(context.Background(), input)
	if err != nil {
		s.log.Error().
			Err(err).
			Str("bucket", s.config.S3.Bucket).
			Str("objectKey", objectKey).
			Uint32("offsetFromFileStart", offsetFromFileStart).
			Uint32("size", size).
			Uint32("calculatedEndOfData", offsetFromFileStart+size).
			Msg("Failed to get slice of object from S3 storage")
		return nil, err
	}
	defer result.Body.Close()

	// Read the contents
	contents, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

func (s *Storage) PresignGetObject(objectKey string) (string, error) {
	request, err := s.s3PresignClient.PresignGetObject(
		context.Background(),
		&s3.GetObjectInput{
			Bucket: aws.String(s.config.S3.Bucket),
			Key:    aws.String(objectKey),
		},
		func(opts *s3.PresignOptions) {
			opts.Expires = time.Minute
		},
	)
	if err != nil {
		return "", err
	}

	return request.URL, nil
}
