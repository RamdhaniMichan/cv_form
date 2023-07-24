package profile

import (
	"net/http"
	"template/entity"
)

type ProfileController interface {
	GetProfile(resp http.ResponseWriter, req *http.Request)
	CreateProfile(resp http.ResponseWriter, req *http.Request)
	UpdateProfile(resp http.ResponseWriter, req *http.Request)
	CheckProfile(profileID string) (*entity.Profile, error)
}
