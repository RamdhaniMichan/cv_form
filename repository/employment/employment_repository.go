package employment

import (
	"github.com/jinzhu/gorm"
	"template/datasource"
	"template/entity"
)

type employmentRepository struct {
}

func NewEmploymentRepository() EmploymentRepository {
	return &employmentRepository{}
}

func (e employmentRepository) Get(id int) (*[]entity.Employment, error) {
	db := datasource.OpenDB()
	var employment []entity.Employment

	err := db.Where("profile_id = ?", id).Find(&employment).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, err
		default:
			return nil, err
		}
	}

	defer db.Close()

	return &employment, nil
}

func (e employmentRepository) Post(payload *entity.Employment) (*entity.Employment, error) {
	db := datasource.OpenDB()

	err := db.Create(&payload).Error

	if err != nil {
		return nil, err
	}

	defer db.Close()

	return payload, nil
}

func (e employmentRepository) Delete(id int) error {
	db := datasource.OpenDB()

	err := db.Delete(&entity.Employment{}, id).Error

	if err != nil {
		return err
	}

	defer db.Close()

	return nil
}
