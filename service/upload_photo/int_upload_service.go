package upload_photo

import (
	"context"
	"mime/multipart"
	"net/url"
	"template/entity"
)

type UploadService interface {
	UploadPhoto(ctx context.Context, fileData []byte, handler *multipart.FileHeader, profileID string) (*entity.Profile, error)
	DownloadPhoto(ctx context.Context, filename string) (*url.URL, error)
	RemovePhoto(ctx context.Context, filename, profileID string) error
}
