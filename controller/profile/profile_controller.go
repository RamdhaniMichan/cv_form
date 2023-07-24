package profile

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"template/entity"
	"template/function"
	"template/service/profile"
)

type profileController struct {
	service profile.ProfileService
}

func NewProfileController(service profile.ProfileService) ProfileController {
	return &profileController{
		service: service,
	}
}

func (p profileController) GetProfile(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	profileID := params["profileID"]

	profile, err := p.service.GetProfile(profileID)

	if err != nil {
		_ = function.SendResponse(resp, 404, fmt.Sprintf("tidak ada data dengan id %d", profileID), err)
		return
	}

	_ = function.SendResponse(resp, 200, "success", profile)
	return
}

func (p profileController) CreateProfile(resp http.ResponseWriter, req *http.Request) {
	var profile entity.Profile

	err := json.NewDecoder(req.Body).Decode(&profile)
	if err != nil {
		_ = function.SendResponse(resp, 500, "something when wrong", err)
		return
	}

	profileID, err := p.service.CreateProfile(&profile)
	if err != nil {
		_ = function.SendResponse(resp, 400, "something when wrong", err)
		return
	}

	_ = function.SendResponse(resp, 201, "success", profileID)
	return
}

func (p profileController) UpdateProfile(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	var profile entity.Profile

	err := json.NewDecoder(req.Body).Decode(&profile)
	if err != nil {
		_ = function.SendResponse(resp, 500, "something when wrong", err)
		return
	}

	profileID := params["profileID"]

	_, err = p.service.GetProfile(profileID)

	if err != nil {
		_ = function.SendResponse(resp, 404, fmt.Sprintf("tidak ada data dengan id %s", profileID), err)
		return
	}

	err = p.service.UpdateProfile(profileID, &profile)
	if err != nil {
		_ = function.SendResponse(resp, 500, "something when wrong", err)
		return
	}

	_ = function.SendResponse(resp, 200, "success", fmt.Sprintf("sukses mengupdate data dengan id %s", profileID))
	return
}

func (p profileController) CheckProfile(profileID string) (*entity.Profile, error) {
	profile, err := p.service.GetProfile(profileID)
	if err != nil {
		return nil, err
	}

	return profile, nil
}
