package controller

import "net/http"

type InterfaceBookController interface {
	GetBooks(resp http.ResponseWriter, req *http.Request)
	CreateBook(resp http.ResponseWriter, req *http.Request)
	GetDetailBookWithParams(resp http.ResponseWriter, req *http.Request)
	GetDetailBookWithQueryParams(resp http.ResponseWriter, req *http.Request)
	DeleteBook(resp http.ResponseWriter, req *http.Request)
	UpdateBook(resp http.ResponseWriter, req *http.Request)
}
