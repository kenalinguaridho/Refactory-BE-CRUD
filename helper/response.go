package helper

import (
	"encoding/json"
	"net/http"
)

type ResponseWithData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseWithoutData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Response(rw http.ResponseWriter, code int, message string, payload interface{}) {

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(code)

	var response any

	status := "Success"

	if code >= 400 {
		status = "Failed"
	}

	if payload != nil {
		response = &ResponseWithData{
			Status: status,
			Message: message,
			Data: payload,
		}
	} else {
		response = &ResponseWithoutData {
			Status: status,
			Message: message,
		}
	}

	resp, _ := json.Marshal(response)

	rw.Write(resp)

}
