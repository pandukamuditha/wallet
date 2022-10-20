package common

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Err string `json:"error"`
}

func WriteJsonResponse(rw http.ResponseWriter, status int, class interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)

	json.NewEncoder(rw).Encode(class)
}
