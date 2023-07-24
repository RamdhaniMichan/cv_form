package employment

import "net/http"

type EmploymentController interface {
	GetEmployment(resp http.ResponseWriter, req *http.Request)
	CreateEmployment(resp http.ResponseWriter, req *http.Request)
	DeleteEmployment(resp http.ResponseWriter, req *http.Request)
}
