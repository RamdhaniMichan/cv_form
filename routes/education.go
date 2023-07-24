package routes

import (
	controllerEducation "template/controller/education"
	router "template/http"
	repoEducation "template/repository/education"
	serviceEducation "template/service/education"
)

var (
	educationRepository repoEducation.EducationRepository       = repoEducation.NewEducationRepository()
	educationService    serviceEducation.EducationService       = serviceEducation.NewEducationService(educationRepository)
	educationController controllerEducation.EducationController = controllerEducation.NewProfileController(educationService, profileController)
)

type EducationRoute struct{}

func (r *EducationRoute) Routing(httpRouter router.Router) {
	httpRouter.GET("/education/{profileID}", educationController.GetEducation)
	httpRouter.POST("/education/{profileID}", educationController.CreateEducation)
	httpRouter.DELETE("/education/{profileID}", educationController.DeleteEducation)
}
