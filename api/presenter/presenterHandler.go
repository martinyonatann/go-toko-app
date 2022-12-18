package presenter

import (
	"encoding/json"
	"net/http"
)

const (
	SuccessText = "success"
	FailedText  = "failed"
)

type response struct {
	Body       *responseBody
	StatusCode int
}

type responseBody struct {
	StatusCode int         `json:"rc"`
	Message    string      `json:"message,omitempty"`
	Error      string      `json:"error,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func (r response) ToJSON(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	return json.NewEncoder(w).Encode(r.Body)
}

func OK(data interface{}) *response {
	return &response{&responseBody{Message: SuccessText, Data: data, StatusCode: http.StatusOK}, http.StatusOK}
}

func Fail(message string, statusCode int) *response {
	return &response{&responseBody{StatusCode: statusCode, Message: FailedText, Error: message}, statusCode}
}
