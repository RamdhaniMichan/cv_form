package employment

import "template/entity"

type EmploymentService interface {
	GetEmployment(employmentID int) (*[]entity.Employment, error)
	CreateEmployment(employment *entity.Employment) (*entity.Employment, error)
	DeleteEmployment(employmentID int) error
}
