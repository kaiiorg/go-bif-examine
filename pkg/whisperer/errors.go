package whisperer

import (
	"errors"
)

var (
	ErrFailedToDownloadFileFromS3Not200Or206 = errors.New("response from S3 was not a 200 OK or 206 Partial Content")
)
