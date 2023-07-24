package upload_photo

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
	"net/url"
	"os"
	"template/datasource"
	"template/entity"
	"time"
)

type uploadPhoto struct {
}

func NewUploadPhoto() UploadRepository {
	return &uploadPhoto{}
}

func (u uploadPhoto) Upload(ctx context.Context, fileData []byte, handler *multipart.FileHeader) error {

	storage := datasource.Minio()
	_, err := storage.PutObject(ctx, os.Getenv("bucket"), handler.Filename, bytes.NewReader(fileData), handler.Size, minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	return nil

}

func (u uploadPhoto) Save(payload *entity.Profile) (*entity.Profile, error) {
	db := datasource.OpenDB()
	err := db.Save(&payload).Error
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return payload, nil
}

func (u uploadPhoto) Remove(ctx context.Context, filename string) error {
	storage := datasource.Minio()
	err := storage.RemoveObject(ctx, os.Getenv("bucket"), filename, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (u uploadPhoto) Download(ctx context.Context, filename string) (*url.URL, error) {
	storage := datasource.Minio()
	url, err := storage.PresignedGetObject(ctx, os.Getenv("bucket"), filename, 1*time.Hour, url.Values{})
	if err != nil {
		return nil, err
	}

	return url, nil
}
