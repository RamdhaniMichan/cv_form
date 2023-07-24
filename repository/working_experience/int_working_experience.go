package working_experience

import "template/entity"

type WorkExperienceRepository interface {
	Get(id int) (*entity.WorkingExperience, error)
	Post(payload *entity.WorkingExperience) (*entity.WorkingExperience, error)
	Update(id int, payload *entity.WorkingExperience) error
}
