package skill

import (
	"github.com/jinzhu/gorm"
	"template/datasource"
	"template/entity"
)

type skillRepository struct {
}

func NewSkillRepository() SkillRepository {
	return &skillRepository{}
}

func (s skillRepository) Get(skillID int) (*[]entity.Skill, error) {
	db := datasource.OpenDB()
	var skill []entity.Skill

	err := db.Where("profile_id = ?", skillID).Find(&skill).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, err
		default:
			return nil, err
		}
	}

	defer db.Close()

	return &skill, nil
}

func (s skillRepository) Post(skill *entity.Skill) (*entity.Skill, error) {
	db := datasource.OpenDB()

	err := db.Create(&skill).Error

	if err != nil {
		return nil, err
	}

	defer db.Close()

	return skill, nil
}

func (s skillRepository) Delete(skillID int) error {
	db := datasource.OpenDB()

	err := db.Delete(&entity.Skill{}, skillID).Error

	if err != nil {
		return err
	}

	defer db.Close()

	return nil
}
