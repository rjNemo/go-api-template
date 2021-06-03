package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rjNemo/go-api-template/api"
)

const (
	apiPath    = "/v1"
	eventsPath = "/events"
)

// Register creates a router and makes all routes available
func Register() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	// create a subRouter for api
	apiRouter := router.PathPrefix(apiPath).Subrouter()

	// create a sub router for events routes
	eventRouter := apiRouter.PathPrefix(eventsPath).Subrouter()
	eventRouter.HandleFunc("/", api.CreateEvent).Methods(http.MethodPost)
	eventRouter.HandleFunc("/{id}", api.GetOneEvent).Methods(http.MethodGet)
	eventRouter.HandleFunc("/", api.GetAllEvents).Methods(http.MethodGet)
	eventRouter.HandleFunc("/{id}", api.UpdateEvent).Methods(http.MethodPatch)
	eventRouter.HandleFunc("/{id}", api.DeleteEvent).Methods(http.MethodDelete)

	// router.HandleFunc("/todo", api.HandleGetTodo)
	// router.HandleFunc("/user", api.HandleGetUser)
	return router
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}
