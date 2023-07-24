package skill

import "template/entity"

type SkillRepository interface {
	Get(skillID int) (*[]entity.Skill, error)
	Post(skill *entity.Skill) (*entity.Skill, error)
	Delete(skillID int) error
}
