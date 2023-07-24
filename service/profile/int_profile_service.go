package profile

import (
	"template/entity"
)

type ProfileService interface {
	GetProfile(profileID string) (*entity.Profile, error)
	CreateProfile(profile *entity.Profile) (string, error)
	UpdateProfile(profileID string, payload *entity.Profile) error
}
