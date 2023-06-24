package handlers

import (
	"encoding/json"
	"net/http"
)

func SendJSONResponse(rw http.ResponseWriter, httpStatus int, jsonResponse any) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(httpStatus)
	json.NewEncoder(rw).Encode(jsonResponse)
}
