package routes

import (
	controllerProfile "template/controller/profile"
	router "template/http"
	repoProfile "template/repository/profile"
	serviceProfile "template/service/profile"
)

var (
	profileRepository repoProfile.ProfileRepository       = repoProfile.NewProfileRepository()
	profileService    serviceProfile.ProfileService       = serviceProfile.NewProfileService(profileRepository)
	profileController controllerProfile.ProfileController = controllerProfile.NewProfileController(profileService)
)

type ProfileRoute struct{}

func (r *ProfileRoute) Routing(httpRouter router.Router) {
	httpRouter.GET("/profile/{profileID}", profileController.GetProfile)
	httpRouter.POST("/profile", profileController.CreateProfile)
	httpRouter.PUT("/profile/{profileID}", profileController.UpdateProfile)
}
