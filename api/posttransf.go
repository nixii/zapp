package api

import (
	"net/http"
	"fmt"

	"nixii.dev/zipp/requests"
)

func PostTransf(w http.ResponseWriter, r *http.Request) error {

	// Get the request
	var data requests.TransferRequest
	err := getJson(r, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return err
	}

	// Is this request getting it from somewhere 
	if (data.From != "") {
		fmt.Println("Wooo!")
	}

	// Verify the data
	err = data.VerifyRequest()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return err
	}

	// All good! :D
	w.WriteHeader(http.StatusOK)
	return nil
}

