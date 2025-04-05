package api

import (
	"encoding/json"
	"net/http"
)

type ErrResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type SuccessResponse struct {
	Code int    `json:"code"`
	Info string `json:"info"`
}

func SendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
