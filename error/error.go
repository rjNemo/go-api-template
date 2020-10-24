package error

import "log"

// Handle terminates the server in case of error
func Handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
