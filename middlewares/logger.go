package middleware

import (
	"log"
	"net/http"
)

// Logging middleware records accessed URIs
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s%s %s", r.Host, r.RequestURI, r.UserAgent())
		next.ServeHTTP(w, r)
	})
}
