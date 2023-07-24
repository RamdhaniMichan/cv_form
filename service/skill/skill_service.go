package skill

import (
	"template/entity"
	"template/repository/skill"
)

type skillService struct {
	repository skill.SkillRepository
}

func NewSkillService(repository skill.SkillRepository) SkillService {
	return &skillService{
		repository: repository,
	}
}

func (s skillService) GetSkill(skillID int) (*[]entity.Skill, error) {
	skill, err := s.repository.Get(skillID)
	if err != nil {
		return nil, err
	}

	return skill, nil
}

func (s skillService) CreateSkill(skill *entity.Skill) (*entity.Skill, error) {

	skill, err := s.repository.Post(skill)
	if err != nil {
		return nil, err
	}

	return skill, nil
}

func (s skillService) DeleteSkill(skillID int) error {
	err := s.repository.Delete(skillID)
	if err != nil {
		return err
	}

	return nil
}
