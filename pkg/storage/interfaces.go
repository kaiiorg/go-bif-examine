package storage

import "io"

type BifStorage interface {
	// UploadFileFromTempFile uploads the contents of the given file to S3 compatible storage
	UploadObjectFromTempFile(objectKey, pathToTempBif string) error
	// UploadObject uploads the contents of the given reader to S3 compatible storage
	UploadObject(objectKey string, reader io.ReadSeeker) error
	// GetSectionFromObject pulls a section of data from the given object instead of the full object
	GetSectionFromObject(objectKey string, offsetFromFileStart, size uint32) ([]byte, error)
	// PresignGetObject created a presigned URL to use to get the requested object without needing to share the creds
	PresignGetObject(objectKey string) (string, error)
}
