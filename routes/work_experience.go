package routes

import (
	controllerWE "template/controller/working-experience"
	router "template/http"
	repoWE "template/repository/working_experience"
	serviceWE "template/service/working_experience"
)

var (
	weRepository repoWE.WorkExperienceRepository          = repoWE.NewWorkExperienceRepository()
	weService    serviceWE.WorkingExperienceService       = serviceWE.NewWorkingExperienceService(weRepository)
	weController controllerWE.WorkingExperienceController = controllerWE.NewWorkingExperienceController(weService, profileController)
)

type WERoute struct{}

func (r *WERoute) Routing(httpRouter router.Router) {
	httpRouter.GET("/working-experience/{profileID}", weController.GetWorkingExperience)
	httpRouter.POST("/working-experience/{profileID}", weController.CreateWorkingExperience)
	httpRouter.PUT("/working-experience/{profileID}", weController.UpdateWorkingExperience)
}
