package rest

import (
	"log"
	"net/http"
)

// security middlewares and response anhancements

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		// You can log additional information here if needed
		next.ServeHTTP(w, r)
	})
}
