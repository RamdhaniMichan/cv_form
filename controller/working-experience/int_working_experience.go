package working_experience

import "net/http"

type WorkingExperienceController interface {
	GetWorkingExperience(resp http.ResponseWriter, req *http.Request)
	CreateWorkingExperience(resp http.ResponseWriter, req *http.Request)
	UpdateWorkingExperience(resp http.ResponseWriter, req *http.Request)
}
