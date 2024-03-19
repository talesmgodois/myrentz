package rest

import (
	"encoding/json"
	"net/http"
)

type Response[T any] struct {
	Error          error  `json:"error"`
	DisplayMessage string `json:"display_message"`
	Success        bool   `json:"success"`
	Data           T      `json:"data"`
}

func doSuccess[T any](w http.ResponseWriter, r *http.Request, data T, m string, status int) {
	response := Response[T]{
		Data:           data,
		Success:        true,
		Error:          nil,
		DisplayMessage: m,
	}
	doResponse(w, r, response, status)
}

func doResponse[T any](w http.ResponseWriter, r *http.Request, response Response[T], status int) {
	responseStr, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(responseStr)
}

func DoSuccess[T any](w http.ResponseWriter, r *http.Request, data T, m string) {
	doSuccess(w, r, data, m, http.StatusOK)
}

func DoSuccessCreate[T any](w http.ResponseWriter, r *http.Request, data T) {
	m := "Created Successfully"
	doSuccess(w, r, data, m, http.StatusCreated)
}
