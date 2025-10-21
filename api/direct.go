package api

import (
	"errors"
	"fmt"
	"net/http"
)

func HandlePwdRequest(w http.ResponseWriter, r *http.Request) {
	
	// Get ready to handle errors
	var err error
	
	// Handle requests
	switch r.Method {
	case http.MethodGet:
		err = Get(w, r)
	case http.MethodPut:
		err = Put(w, r)
	case http.MethodPatch:
		err = Patch(w, r)
	default:
		err = errors.ErrUnsupported
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	// If any errors occurred
	if err != nil {
		fmt.Printf("An error has occurred: %s\n", err)
	}
}

func HandleCmpRequest(w http.ResponseWriter, r *http.Request) {

	var err error

	switch r.Method {
	case http.MethodPatch:
		err = PatchCmp(w, r)
	default:
		err = errors.ErrUnsupported
	}

	if err != nil {
		fmt.Printf("An error has occurred: %s\n", err)
	}
}