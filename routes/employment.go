package routes

import (
	controllerEmployment "template/controller/employment"
	router "template/http"
	repoEmployment "template/repository/employment"
	serviceEmployment "template/service/employment"
)

var (
	employmentRepository repoEmployment.EmploymentRepository       = repoEmployment.NewEmploymentRepository()
	employmentService    serviceEmployment.EmploymentService       = serviceEmployment.NewEmploymentService(employmentRepository)
	employmentController controllerEmployment.EmploymentController = controllerEmployment.NewEmploymentController(employmentService, profileController)
)

type EmploymentRoute struct{}

func (r *EmploymentRoute) Routing(httpRouter router.Router) {
	httpRouter.GET("/employment/{profileID}", employmentController.GetEmployment)
	httpRouter.POST("/employment/{profileID}", employmentController.CreateEmployment)
	httpRouter.DELETE("/employment/{profileID}", employmentController.DeleteEmployment)
}
