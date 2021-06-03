package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	middleware "github.com/rjNemo/go-api-template/middlewares"
	"github.com/rjNemo/go-api-template/router"
)

const (
	port = ":3000"
)

func main() {
	// serve a static folder (here a react app)
	// fs := http.FileServer(http.Dir("./build"))
	// http.Handle("/", fs)

	// router
	r := router.Register()
	// middlewares
	r.Use(middleware.Logging)
	// configure server
	http.Handle("/", &Server{r})
	// start server using defaultMux
	log.Printf("Start Go server on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// Server configures Server behavior
type Server struct {
	r *mux.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}
	// set common Headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-DNS-Prefetch-Control", "off")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("Referrer-Policy", "no-referrer")
	// Lets Gorilla work
	s.r.ServeHTTP(w, r)
}
