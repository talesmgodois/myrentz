package rest

import (
	"backend/adapters/delivery/rest/people"
	"github.com/gorilla/mux"
	"net/http"
)

func BuildServer() *http.Server {
	ro := mux.NewRouter()
	ro.Use(LoggingMiddleware)

	buildSimpleCrud(ro, "/people", &people.PeopleHandler{})

	return &http.Server{
		Handler: ro,
	}
}
