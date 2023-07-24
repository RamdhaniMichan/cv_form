package upload_photo

import (
	"context"
	"mime/multipart"
	"net/url"
	"template/entity"
)

type UploadRepository interface {
	Upload(ctx context.Context, fileData []byte, handler *multipart.FileHeader) error
	Save(payload *entity.Profile) (*entity.Profile, error)
	Remove(ctx context.Context, filename string) error
	Download(ctx context.Context, filename string) (*url.URL, error)
}
