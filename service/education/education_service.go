package education

import (
	"template/entity"
	"template/repository/education"
)

type educationService struct {
	repository education.EducationRepository
}

func NewEducationService(repository education.EducationRepository) EducationService {
	return &educationService{
		repository: repository,
	}
}

func (e educationService) GetEducation(educationID int) (*[]entity.Education, error) {
	education, err := e.repository.Get(educationID)
	if err != nil {
		return nil, err
	}

	return education, nil
}

func (e educationService) CreateEducation(education *entity.Education) (*entity.Education, error) {
	education, err := e.repository.Post(education)
	if err != nil {
		return nil, err
	}

	return education, nil
}

func (e educationService) DeleteEducation(educationID int) error {
	err := e.repository.Delete(educationID)
	if err != nil {
		return err
	}

	return nil
}
