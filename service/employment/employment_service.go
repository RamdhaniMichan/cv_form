package employment

import (
	"template/entity"
	"template/repository/employment"
)

type employmentService struct {
	repository employment.EmploymentRepository
}

func NewEmploymentService(repository employment.EmploymentRepository) EmploymentService {
	return &employmentService{
		repository: repository,
	}
}

func (e employmentService) GetEmployment(employmentID int) (*[]entity.Employment, error) {
	education, err := e.repository.Get(employmentID)
	if err != nil {
		return nil, err
	}

	return education, nil
}

func (e employmentService) CreateEmployment(employment *entity.Employment) (*entity.Employment, error) {
	employment, err := e.repository.Post(employment)
	if err != nil {
		return nil, err
	}

	return employment, nil
}

func (e employmentService) DeleteEmployment(employmentID int) error {
	err := e.repository.Delete(employmentID)
	if err != nil {
		return err
	}

	return nil
}
