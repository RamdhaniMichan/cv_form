package education

import "template/entity"

type EducationService interface {
	GetEducation(educationID int) (*[]entity.Education, error)
	CreateEducation(education *entity.Education) (*entity.Education, error)
	DeleteEducation(educationID int) error
}
