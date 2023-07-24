package working_experience

import (
	"template/entity"
	"template/repository/working_experience"
)

type workingExperienceService struct {
	repository working_experience.WorkExperienceRepository
}

func NewWorkingExperienceService(repository working_experience.WorkExperienceRepository) WorkingExperienceService {
	return &workingExperienceService{
		repository: repository,
	}
}

func (w workingExperienceService) GetWorkingExperience(id int) (*entity.WorkingExperience, error) {
	skill, err := w.repository.Get(id)
	if err != nil {
		return nil, err
	}

	return skill, nil
}

func (w workingExperienceService) CreateWorkingExperience(payload *entity.WorkingExperience) (*entity.WorkingExperience, error) {
	skill, err := w.repository.Post(payload)
	if err != nil {
		return nil, err
	}

	return skill, nil
}

func (w workingExperienceService) UpdateWorkingExperience(id int, payload *entity.WorkingExperience) error {
	err := w.repository.Update(id, payload)
	if err != nil {
		return err
	}

	return nil
}
