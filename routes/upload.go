package routes

import (
	controllerupload "template/controller/upload_photo"
	router "template/http"
	repoupload "template/repository/upload_photo"
	serviceupload "template/service/upload_photo"
)

var (
	uploadRepository repoupload.UploadRepository       = repoupload.NewUploadPhoto()
	uploadService    serviceupload.UploadService       = serviceupload.NewUploadService(uploadRepository, profileRepository)
	uploadController controllerupload.UploadController = controllerupload.NewUploadPhoto(uploadService, profileController)
)

type UploadRoute struct{}

func (r *UploadRoute) Routing(httpRouter router.Router) {
	httpRouter.GET("/photo/{profileID}", uploadController.DownloadPhoto)
	httpRouter.POST("/photo/{profileID}", uploadController.UploadPhoto)
	httpRouter.DELETE("/photo/{profileID}", uploadController.RemovePhoto)
}
