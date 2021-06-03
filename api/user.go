package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rjNemo/go-api-template/error"
)

const userURI = "https://jsonplaceholder.typicode.com/users/1"

// GeoLocation encapsulates a location coordinates
type GeoLocation struct {
	Lat, Lng string
}

// Company describes a business
type Company struct {
	Name, CatchPhrase, Bs string
}

// Address serves to locate a business
type Address struct {
	Street, Suite, City, Zipcode string
	Geo                          GeoLocation
}

// User encapsulates an employee's information
type User struct {
	ID                                    int
	Name, UserName, Email, Phone, Website string
	Address                               Address
	Company                               Company
}

// GetUser fetches a User from API
func GetUser() *User {
	// perform request
	resp, err := http.Get(userURI)
	error.Handle(err)
	// read data
	data, err := ioutil.ReadAll(resp.Body)
	error.Handle(err)
	// create User struct
	var u User
	err = json.Unmarshal(data, &u)
	error.Handle(err)
	fmt.Println(u)
	return &u
}

// HandleGetUser returns User via http
func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	u := GetUser()
	fmt.Fprintf(w, "User: %v", u)
}
