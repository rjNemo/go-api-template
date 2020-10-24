package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rjNemo/sandbox/data"
	"github.com/rjNemo/sandbox/error"
	model "github.com/rjNemo/sandbox/models"
)

// CreateEvent add an event to data
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	// create new event data placeholder
	var newEvent model.Event
	// read request body
	resp, err := ioutil.ReadAll(r.Body)
	error.Handle(err)
	// decode json to event
	json.Unmarshal(resp, &newEvent)
	// add new event to events
	data.AllEvents = append(data.AllEvents, newEvent)
	// return succes status code and created json object
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEvent)
}

// GetOneEvent fetch one event by ID
func GetOneEvent(w http.ResponseWriter, r *http.Request) {
	// get url params
	eventID := mux.Vars(r)["id"]
	// find event by ID
	for _, event := range data.AllEvents {
		if event.ID == eventID {
			json.NewEncoder(w).Encode(event)
		}
	}
}

// GetAllEvents returns all events from data
func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(data.AllEvents)
}

// UpdateEvent edits a Event by ID
func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	// get url param
	eventID := mux.Vars(r)["id"]
	// get new values from request
	var updatedEvent model.Event
	resp, err := ioutil.ReadAll(r.Body)
	error.Handle(err)
	json.Unmarshal(resp, &updatedEvent)
	// find event to update
	for i, event := range data.AllEvents {
		if event.ID == eventID {
			event.Title = updatedEvent.Title
			event.Description = updatedEvent.Description
			data.AllEvents = append(data.AllEvents[:i], event)
			// return new json object
			json.NewEncoder(w).Encode(event)
		}
	}
}

// DeleteEvent removes an event from data
func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for i, singleEvent := range data.AllEvents {
		if singleEvent.ID == eventID {
			data.AllEvents = append(data.AllEvents[:i], data.AllEvents[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)
		}
	}
}
