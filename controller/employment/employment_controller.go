package employment

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
	service "template/service/employment"
)

type employmentController struct {
	service           service.EmploymentService
	profileController profileController.ProfileController
}

func NewEmploymentController(service service.EmploymentService, profileController profileController.ProfileController) EmploymentController {
	return &employmentController{
		service:           service,
		profileController: profileController,
	}
}

func (e employmentController) GetEmployment(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	profileID := params["profileID"]

	p, err := e.profileController.CheckProfile(profileID)
	if err != nil {
		_ = function.SendResponse(resp, 404, fmt.Sprintf("tidak ada data dengan id %s", profileID), err)
		return
	}

	res, err := e.service.GetEmployment(int(p.ID))

	if err != nil {
		_ = function.SendResponse(resp, 404, fmt.Sprintf("tidak ada data dengan id %d", p.ID), err)
		return
	}

	_ = function.SendResponse(resp, 200, "success", res)
	return
}

func (e employmentController) CreateEmployment(resp http.ResponseWriter, req *http.Request) {
	var employment entity.Employment

	params := mux.Vars(req)

	profileID := params["profileID"]

	profile, err := e.profileController.CheckProfile(profileID)
	if err != nil {
		_ = function.SendResponse(resp, 404, fmt.Sprintf("tidak ada data dengan id %d", profileID), err)
		return
	}

	employment.ProfileID = int(profile.ID)

	err = json.NewDecoder(req.Body).Decode(&employment)
	if err != nil {
		_ = function.SendResponse(resp, 500, "something when wrong", err)
		return
	}

	res, err := e.service.CreateEmployment(&employment)
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

func (e employmentController) DeleteEmployment(resp http.ResponseWriter, req *http.Request) {
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

	err = e.service.DeleteEmployment(int(id))
	if err != nil {
		_ = function.SendResponse(resp, 500, "something when wrong", err)
		return
	}

	_ = function.SendResponse(resp, 200, "success", fmt.Sprintf("sukses mengupdate data dengan id %s", profileID))
	return
}
