package employment

import "template/entity"

type EmploymentRepository interface {
	Get(id int) (*[]entity.Employment, error)
	Post(payload *entity.Employment) (*entity.Employment, error)
	Delete(id int) error
}
