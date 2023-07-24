package working_experience

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
	service "template/service/working_experience"
)

type workingExperienceController struct {
	service           service.WorkingExperienceService
	profileController profileController.ProfileController
}

func NewWorkingExperienceController(service service.WorkingExperienceService, profileController profileController.ProfileController) WorkingExperienceController {
	return &workingExperienceController{
		service:           service,
		profileController: profileController,
	}
}

func (w workingExperienceController) GetWorkingExperience(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	profileID := params["profileID"]

	p, err := w.profileController.CheckProfile(profileID)
	if err != nil {
		_ = function.SendResponse(resp, 404, fmt.Sprintf("tidak ada data dengan id %s", profileID), err)
		return
	}

	id, errConv := strconv.ParseInt(req.URL.Query().Get("id"), 10, 64)

	if errConv != nil {
		log.Println("terjadi kesalahan ketika mengkonversi id ke integer", errConv)
	}

	res, err := w.service.GetWorkingExperience(int(p.ID))

	if err != nil {
		_ = function.SendResponse(resp, 404, fmt.Sprintf("tidak ada data dengan id %d", id), err)
		return
	}

	_ = function.SendResponse(resp, 200, "success", res)
	return
}

func (w workingExperienceController) CreateWorkingExperience(resp http.ResponseWriter, req *http.Request) {
	var we entity.WorkingExperience

	params := mux.Vars(req)

	profileID := params["profileID"]

	profile, err := w.profileController.CheckProfile(profileID)
	if err != nil {
		_ = function.SendResponse(resp, 404, fmt.Sprintf("tidak ada data dengan id %d", profileID), err)
		return
	}

	we.ProfileID = int(profile.ID)

	err = json.NewDecoder(req.Body).Decode(&we)
	if err != nil {
		_ = function.SendResponse(resp, 500, "something when wrong", err)
		return
	}

	res, err := w.service.CreateWorkingExperience(&we)
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

func (w workingExperienceController) UpdateWorkingExperience(resp http.ResponseWriter, req *http.Request) {
	var we entity.WorkingExperience
	params := mux.Vars(req)

	profileID := params["profileID"]

	p, err := w.profileController.CheckProfile(profileID)
	if err != nil {
		_ = function.SendResponse(resp, 404, fmt.Sprintf("tidak ada data dengan id %d", profileID), err)
		return
	}

	err = json.NewDecoder(req.Body).Decode(&we)
	if err != nil {
		_ = function.SendResponse(resp, 500, "something when wrong", err)
		return
	}

	err = w.service.UpdateWorkingExperience(int(p.ID), &we)
	if err != nil {
		_ = function.SendResponse(resp, 500, "something when wrong", err)
		return
	}

	_ = function.SendResponse(resp, 200, "success", fmt.Sprintf("sukses mengupdate data dengan id %s", profileID))
	return
}
