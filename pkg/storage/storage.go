package storage

import (
	"fmt"
	"io"
	"net/url"
	"os"

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
	s3Client  *s3.S3
}

func New(conf *config.Config, log zerolog.Logger) (*Storage, error) {
	u := &url.URL{
		Scheme: conf.S3.Scheme,
		Host:   fmt.Sprintf("%s:%d", conf.S3.Host, conf.S3.Port),
	}

	s3Session, err := session.NewSession(&aws.Config{
		Endpoint:         aws.String(u.String()),
		S3ForcePathStyle: aws.Bool(conf.S3.ForcePathStyle),
		Region:           aws.String(conf.S3.Region),
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
	_, err := s.s3Client.PutObject(putObjectInput)
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
	result, err := s.s3Client.GetObject(input)
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
