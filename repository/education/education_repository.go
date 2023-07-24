package education

import (
	"github.com/jinzhu/gorm"
	"template/datasource"
	"template/entity"
)

type educationRepository struct {
}

func NewEducationRepository() EducationRepository {
	return &educationRepository{}
}

func (e educationRepository) Get(educationID int) (*[]entity.Education, error) {
	db := datasource.OpenDB()
	var education []entity.Education

	err := db.Where("profile_id = ?", educationID).Find(&education).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, err
		default:
			return nil, err
		}
	}

	defer db.Close()

	return &education, nil
}

func (e educationRepository) Post(education *entity.Education) (*entity.Education, error) {
	db := datasource.OpenDB()

	err := db.Create(&education).Error

	if err != nil {
		return nil, err
	}

	defer db.Close()

	return education, nil
}

func (e educationRepository) Delete(educationID int) error {
	db := datasource.OpenDB()

	err := db.Delete(&entity.Education{}, educationID).Error

	if err != nil {
		return err
	}

	defer db.Close()

	return nil
}
