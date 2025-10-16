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
	switch r.Method {
	case http.MethodGet:
		err = Get(w, r)
	case http.MethodPut:
		err = Put(w, r)
	default:
		err = errors.ErrUnsupported
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	// If any errors occurred
	if err != nil {
		fmt.Printf("An error has occurred: %s\n", err)
	}
}