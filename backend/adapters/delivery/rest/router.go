package rest

import (
	"github.com/gorilla/mux"
)

func buildSimpleCrud(ro *mux.Router, path string, h CrudHandler) *mux.Router {
	r := ro.PathPrefix(path).Subrouter()
	r.HandleFunc("/", h.Read).Methods("GET")
	r.HandleFunc("/", h.Create).Methods("POST")

	r.HandleFunc("/{id}", h.Read).Methods("GET")
	r.HandleFunc("/{id}", h.Update).Methods("PUT")
	r.HandleFunc("/{id}", h.Delete).Methods("DELETE")
	return r
}
