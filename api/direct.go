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
	} else if r.Method == http.MethodPut {
		err = Put(w, r)
	} else {
		err = errors.ErrUnsupported
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	// If any errors occurred
	if err != nil {
		fmt.Printf("An error has occurred: %s\n", err)
	}
}

func HandleKeyExchangeRequest(w http.ResponseWriter, r *http.Request) {
	
	var err error

	if r.Method == http.MethodGet {
		err = GetKeys(w, r)
	} else {
		err = errors.ErrUnsupported
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	if err != nil {
		fmt.Printf("An error has occurred: %s\n", err)
	}
}