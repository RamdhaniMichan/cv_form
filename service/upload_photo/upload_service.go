package upload_photo

import (
	"context"
	"mime/multipart"
	"net/url"
	"template/entity"
	"template/repository/profile"
	"template/repository/upload_photo"
)

type uploadPhoto struct {
	repository  upload_photo.UploadRepository
	repoProfile profile.ProfileRepository
}

func NewUploadService(repository upload_photo.UploadRepository, repoProfile profile.ProfileRepository) UploadService {
	return &uploadPhoto{
		repository:  repository,
		repoProfile: repoProfile,
	}
}

func (u uploadPhoto) UploadPhoto(ctx context.Context, fileData []byte, handler *multipart.FileHeader, profileID string) (*entity.Profile, error) {
	err := u.repository.Upload(ctx, fileData, handler)
	if err != nil {
		return nil, err
	}

	profile, err := u.repoProfile.Get(profileID)
	if err != nil {
		return nil, err
	}

	profile.Photo = handler.Filename
	profile.ProfileCode = profileID

	p, err := u.repository.Save(profile)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (u uploadPhoto) DownloadPhoto(ctx context.Context, filename string) (*url.URL, error) {
	url, err := u.repository.Download(ctx, filename)
	if err != nil {
		return nil, err
	}

	return url, err
}

func (u uploadPhoto) RemovePhoto(ctx context.Context, filename, profileID string) error {
	err := u.repository.Remove(ctx, filename)
	if err != nil {
		return err
	}

	profile, err := u.repoProfile.Get(profileID)
	if err != nil {
		return err
	}

	profile.Photo = ""
	profile.ProfileCode = profileID

	_, err = u.repository.Save(profile)
	if err != nil {
		return err
	}

	return nil
}
