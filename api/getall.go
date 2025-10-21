package api

import (
	"encoding/json"
	"net/http"

	"nixii.dev/zipp/requests"
	"nixii.dev/zipp/save"
)

func GetAll(w http.ResponseWriter, r *http.Request) error {
	
	// Get the request
	var data requests.GetAllRequest
	err := getJson(r, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	// Verify the data
	err = data.VerifyRequest()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	// Get the stuff
	save, err := save.ReadSaveFile(data.MasterPassword)
	if err != nil {
		return err
	}

	// JSON-ify the save
	marshaled, err := json.Marshal(save.Headers())
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshaled)

	return nil
}