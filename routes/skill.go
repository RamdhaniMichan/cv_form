package routes

import (
	controllerSkill "template/controller/skill"
	router "template/http"
	repoSkill "template/repository/skill"
	serviceSkill "template/service/skill"
)

var (
	skillRepository repoSkill.SkillRepository       = repoSkill.NewSkillRepository()
	skillService    serviceSkill.SkillService       = serviceSkill.NewSkillService(skillRepository)
	skillController controllerSkill.SkillController = controllerSkill.NewSkillController(skillService, profileService, profileController)
)

type SkillRoute struct{}

func (r *SkillRoute) Routing(httpRouter router.Router) {
	httpRouter.GET("/skill/{profileID}", skillController.GetSkill)
	httpRouter.POST("/skill/{profileID}", skillController.CreateSkill)
	httpRouter.DELETE("/skill/{profileID}", skillController.DeleteSkill)
}
