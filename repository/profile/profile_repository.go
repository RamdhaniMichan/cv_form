package profile

import (
	"github.com/jinzhu/gorm"
	"template/datasource"
	"template/entity"
)

type profileRepository struct {
}

func NewProfileRepository() ProfileRepository {
	return &profileRepository{}
}

func (p profileRepository) Get(profileID string) (*entity.Profile, error) {
	db := datasource.OpenDB()
	var profile entity.Profile

	err := db.Where("profile_code = ?", profileID).Take(&profile).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, err
		default:
			return nil, err
		}
	}

	defer db.Close()

	return &profile, nil
}

func (p profileRepository) Post(profile *entity.Profile) error {
	db := datasource.OpenDB()

	err := db.Create(&profile).Error

	if err != nil {
		return err
	}

	defer db.Close()

	return nil
}

func (p profileRepository) Update(profileID string, payload *entity.Profile) error {
	db := datasource.OpenDB()

	err := db.Model(&entity.Profile{}).Where("profile_code = ?", profileID).Updates(&payload).Error

	if err != nil {
		return err
	}

	defer db.Close()

	return nil
}
