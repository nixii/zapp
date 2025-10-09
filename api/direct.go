package api

import (
	"errors"
	"fmt"
	"net/http"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	
	// Get ready to handle errors
	var err error
	
	// Handle requests
	if r.Method == http.MethodGet {
		err = Get(w, r)
	} else {
		err = errors.ErrUnsupported
	}

	// If any errors occurred
	if err != nil {
		fmt.Printf("An error has occurred: %s\n", err)
	}
}