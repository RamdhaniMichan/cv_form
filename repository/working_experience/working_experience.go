package working_experience

import (
	"github.com/jinzhu/gorm"
	"template/datasource"
	"template/entity"
)

type workExperienceRepository struct {
}

func NewWorkExperienceRepository() WorkExperienceRepository {
	return &workExperienceRepository{}
}

func (w workExperienceRepository) Get(id int) (*entity.WorkingExperience, error) {
	db := datasource.OpenDB()
	var we entity.WorkingExperience

	err := db.Where("profile_id = ?", id).Take(&we).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, err
		default:
			return nil, err
		}
	}

	defer db.Close()

	return &we, nil
}

func (w workExperienceRepository) Post(payload *entity.WorkingExperience) (*entity.WorkingExperience, error) {
	db := datasource.OpenDB()

	err := db.Create(&payload).Error

	if err != nil {
		return nil, err
	}

	defer db.Close()

	return payload, nil
}

func (w workExperienceRepository) Update(id int, payload *entity.WorkingExperience) error {
	db := datasource.OpenDB()

	err := db.Model(&entity.WorkingExperience{}).Where("profile_id = ?", id).Updates(&payload).Error

	if err != nil {
		return err
	}

	defer db.Close()

	return nil
}
