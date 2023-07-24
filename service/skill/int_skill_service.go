package skill

import "template/entity"

type SkillService interface {
	GetSkill(skillID int) (*[]entity.Skill, error)
	CreateSkill(skill *entity.Skill) (*entity.Skill, error)
	DeleteSkill(skillID int) error
}
