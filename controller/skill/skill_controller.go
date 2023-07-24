package skill

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
	"template/service/profile"
	"template/service/skill"
)

type skillController struct {
	serviceSkill      skill.SkillService
	serviceProfile    profile.ProfileService
	profileController profileController.ProfileController
}

func NewSkillController(serviceSkill skill.SkillService, serviceProfile profile.ProfileService, profileController profileController.ProfileController) SkillController {
	return &skillController{
		serviceSkill:      serviceSkill,
		serviceProfile:    serviceProfile,
		profileController: profileController,
	}
}

func (s skillController) GetSkill(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	profileID := params["profileID"]

	p, err := s.profileController.CheckProfile(profileID)
	if err != nil {
		_ = function.SendResponse(resp, 404, fmt.Sprintf("tidak ada data dengan id %s", profileID), err)
		return
	}

	res, err := s.serviceSkill.GetSkill(int(p.ID))

	if err != nil {
		_ = function.SendResponse(resp, 404, fmt.Sprintf("tidak ada data dengan id %d", p.ID), err)
		return
	}

	_ = function.SendResponse(resp, 200, "success", res)
	return
}

func (s skillController) CreateSkill(resp http.ResponseWriter, req *http.Request) {
	var skill entity.Skill

	params := mux.Vars(req)

	profileID := params["profileID"]

	profile, err := s.profileController.CheckProfile(profileID)
	if err != nil {
		_ = function.SendResponse(resp, 404, fmt.Sprintf("tidak ada data dengan id %d", profileID), err)
		return
	}

	skill.ProfileID = int(profile.ID)

	err = json.NewDecoder(req.Body).Decode(&skill)
	if err != nil {
		_ = function.SendResponse(resp, 500, "something when wrong", err)
		return
	}

	res, err := s.serviceSkill.CreateSkill(&skill)
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

func (s skillController) DeleteSkill(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	profileID := params["profileID"]

	_, err := s.profileController.CheckProfile(profileID)
	if err != nil {
		_ = function.SendResponse(resp, 404, fmt.Sprintf("tidak ada data dengan id %d", profileID), err)
		return
	}

	id, errConv := strconv.ParseInt(req.URL.Query().Get("id"), 10, 64)

	if errConv != nil {
		log.Println("terjadi kesalahan ketika mengkonversi id ke integer", errConv)
	}

	err = s.serviceSkill.DeleteSkill(int(id))
	if err != nil {
		_ = function.SendResponse(resp, 500, "something when wrong", err)
		return
	}

	_ = function.SendResponse(resp, 200, "success", fmt.Sprintf("sukses delete data dengan id %s", profileID))
	return
}
