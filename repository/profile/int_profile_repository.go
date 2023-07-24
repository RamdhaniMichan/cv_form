package profile

import (
	"template/entity"
)

type ProfileRepository interface {
	Get(profileID string) (*entity.Profile, error)
	Post(profile *entity.Profile) error
	Update(profileID string, payload *entity.Profile) error
}
