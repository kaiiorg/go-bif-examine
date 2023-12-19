package storage

type BifStorage interface {
	// UploadFileFromTempFile uploads the contents of the given file to S3 compatible storage
	UploadFileFromTempFile(objectKey, pathToTempBif string) error
	// GetSectionFromObject pulls a section of data from the given object instead of the full object
	GetSectionFromObject(objectKey string, offsetFromFileStart, size uint32) ([]byte, error)
}
