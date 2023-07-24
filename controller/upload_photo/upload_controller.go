package upload_photo

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	profileController "template/controller/profile"
	"template/entity"
	"template/function"
	"template/service/upload_photo"
)

type uploadPhoto struct {
	service           upload_photo.UploadService
	profileController profileController.ProfileController
}

func NewUploadPhoto(service upload_photo.UploadService, profileController profileController.ProfileController) UploadController {
	return &uploadPhoto{
		service:           service,
		profileController: profileController,
	}
}

func (u uploadPhoto) UploadPhoto(resp http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	params := mux.Vars(req)
	data, handler, err := req.FormFile("file")
	if err != nil {
		_ = function.SendResponse(resp, 500, "something when wrong", err)
		return
	}

	fileData, err := ioutil.ReadAll(data)
	if err != nil {
		_ = function.SendResponse(resp, 500, "something when wrong", err)
		return
	}

	profileID := params["profileID"]

	_, err = u.profileController.CheckProfile(profileID)
	if err != nil {
		_ = function.SendResponse(resp, 404, fmt.Sprintf("tidak ada data dengan id %d", profileID), err)
		return
	}

	p, err := u.service.UploadPhoto(ctx, fileData, handler, profileID)
	if err != nil {
		_ = function.SendResponse(resp, 500, "something when wrong", err)
		return
	}

	_ = function.SendResponse(resp, 200, "success", entity.ResponsePhoto{
		ProfileCode: profileID,
		Photo:       p.Photo,
	})
	return
}

func (u uploadPhoto) DownloadPhoto(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	profileID := params["profileID"]
	ctx := context.Background()

	p, err := u.profileController.CheckProfile(profileID)
	if err != nil {
		_ = function.SendResponse(resp, 404, fmt.Sprintf("tidak ada data dengan id %d", profileID), err)
		return
	}

	filename := p.Photo

	url, err := u.service.DownloadPhoto(ctx, filename)
	if err != nil {
		_ = function.SendResponse(resp, 500, "something when wrong", err)
		return
	}

	_ = function.SendResponse(resp, 200, "success", url.Redacted())
	return
}

func (u uploadPhoto) RemovePhoto(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	profileID := params["profileID"]
	ctx := context.Background()

	p, err := u.profileController.CheckProfile(profileID)
	if err != nil {
		_ = function.SendResponse(resp, 404, fmt.Sprintf("tidak ada data dengan id %d", profileID), err)
		return
	}

	filename := p.Photo

	err = u.service.RemovePhoto(ctx, filename, profileID)

	if err != nil {
		_ = function.SendResponse(resp, 500, "something when wrong", err)
		return
	}

	_ = function.SendResponse(resp, 200, "success", entity.ResponsePhoto{
		ProfileCode: profileID,
		Photo:       "",
	})
	return
}
