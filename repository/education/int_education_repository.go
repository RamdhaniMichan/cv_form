package education

import "template/entity"

type EducationRepository interface {
	Get(educationID int) (*[]entity.Education, error)
	Post(education *entity.Education) (*entity.Education, error)
	Delete(educationID int) error
}
