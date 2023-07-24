package skill

import "net/http"

type SkillController interface {
	GetSkill(resp http.ResponseWriter, req *http.Request)
	CreateSkill(resp http.ResponseWriter, req *http.Request)
	DeleteSkill(resp http.ResponseWriter, req *http.Request)
}
