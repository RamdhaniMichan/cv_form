package working_experience

import "template/entity"

type WorkingExperienceService interface {
	GetWorkingExperience(id int) (*entity.WorkingExperience, error)
	CreateWorkingExperience(payload *entity.WorkingExperience) (*entity.WorkingExperience, error)
	UpdateWorkingExperience(id int, payload *entity.WorkingExperience) error
}
