package education

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	profileController "template/controller/profile"
	"template/entity"
	"template/function"
	service "template/service/education"
)

type educationController struct {
	service           service.EducationService
	profileController profileController.ProfileController
}

func NewProfileController(service service.EducationService, profileController profileController.ProfileController) EducationController {
	return &educationController{
		service:           service,
		profileController: profileController,
	}
}

func (e educationController) GetEducation(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	profileID := params["profileID"]

	p, err := e.profileController.CheckProfile(profileID)
	if err != nil {
		_ = function.SendResponse(resp, 404, fmt.Sprintf("tidak ada data dengan id %s", profileID), err)
		return
	}

	res, err := e.service.GetEducation(int(p.ID))

	if err != nil {
		_ = function.SendResponse(resp, 404, fmt.Sprintf("tidak ada data dengan id %d", p.ID), err)
		return
	}

	_ = function.SendResponse(resp, 200, "success", res)
	return
}

func (e educationController) CreateEducation(resp http.ResponseWriter, req *http.Request) {
	var education entity.Education

	params := mux.Vars(req)

	profileID := params["profileID"]

	profile, err := e.profileController.CheckProfile(profileID)
	if err != nil {
		_ = function.SendResponse(resp, 404, fmt.Sprintf("tidak ada data dengan id %d", profileID), err)
		return
	}

	education.ProfileID = int(profile.ID)

	err = json.NewDecoder(req.Body).Decode(&education)
	if err != nil {
		_ = function.SendResponse(resp, 500, "something when wrong", err)
		return
	}

	res, err := e.service.CreateEducation(&education)
	if err != nil {
		_ = function.SendResponse(resp, 400, "something when wrong", err)
		return
	}

	_ = function.SendResponse(resp, 201, "success", entity.Response{
		ProfileCode: profileID,
		ID:          int(res.ID),
	})
	return
}

func (e educationController) DeleteEducation(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	profileID := params["profileID"]

	_, err := e.profileController.CheckProfile(profileID)
	if err != nil {
		_ = function.SendResponse(resp, 404, fmt.Sprintf("tidak ada data dengan id %d", profileID), err)
		return
	}

	id, errConv := strconv.ParseInt(req.URL.Query().Get("id"), 10, 64)

	if errConv != nil {
		log.Println("terjadi kesalahan ketika mengkonversi id ke integer", errConv)
	}

	err = e.service.DeleteEducation(int(id))
	if err != nil {
		_ = function.SendResponse(resp, 500, "something when wrong", err)
		return
	}

	_ = function.SendResponse(resp, 200, "success", fmt.Sprintf("sukses delete data dengan id %s", profileID))
	return
}
