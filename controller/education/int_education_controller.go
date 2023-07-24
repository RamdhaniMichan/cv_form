package education

import "net/http"

type EducationController interface {
	GetEducation(resp http.ResponseWriter, req *http.Request)
	CreateEducation(resp http.ResponseWriter, req *http.Request)
	DeleteEducation(resp http.ResponseWriter, req *http.Request)
}
