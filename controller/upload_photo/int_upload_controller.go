package upload_photo

import "net/http"

type UploadController interface {
	UploadPhoto(resp http.ResponseWriter, req *http.Request)
	DownloadPhoto(resp http.ResponseWriter, req *http.Request)
	RemovePhoto(resp http.ResponseWriter, req *http.Request)
}
