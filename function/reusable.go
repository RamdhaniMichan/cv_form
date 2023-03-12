package function

import (
	"encoding/json"
	"net/http"
)

// BaseResponse base response
type BaseResponse struct {
	StatusCode int         `json:"status_code" example:"200"`
	Message    string      `json:"message" example:"success"`
	Data       interface{} `json:"data"`
}

// SendResponse ..
func SendResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) string {
	res := new(BaseResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	res.StatusCode = statusCode
	res.Message = message
	res.Data = data
	json.NewEncoder(w).Encode(res)
	return message
}
