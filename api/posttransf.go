package api

import (
	"net/http"
	"fmt"

	"nixii.dev/zapp/requests"
	"nixii.dev/zapp/save"
)

var receiving bool

func PostTransf(w http.ResponseWriter, r *http.Request) error {

	// Get the request
	var data requests.TransferRequest
	err := getJson(r, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return err
	}

	// Verify the data
	err = data.VerifyRequest()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return err
	}

	if (data.Receiving) {
	     fmt.Println("receiving!")
	} else {
	     fmt.Println("not receiving..")
	}

	// enable the system
	if (data.From == "") {
		receiving = data.Receiving
	} else {
		// send the data back
		data, err := save.ReadSaveFileBytes()
		if (err != nil) {
			return err
		}
		w.Write(data)
		fmt.Println(string(data))
	}

	// All good! :D
	w.WriteHeader(http.StatusOK)
	return nil
}
