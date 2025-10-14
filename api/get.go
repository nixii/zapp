package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"nixii.dev/zipp/requests"
	"nixii.dev/zipp/save"
)

func Get(w http.ResponseWriter, req *http.Request) error {

	// Get the request
	var data requests.GetRequest
	err := getJson(req, &data)
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

	// Perform the reqeust
	saves, err := save.ReadSaveFile(data.MasterPassword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return err
	}

	// Get the password
	pwd, err := saves.GetAccountInfo(data.Website, data.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return err
	}

	// Report the pwd
	fmt.Printf("[GET]: Password for W%s,U%s was GET\n", data.Website, data.Username)

	// Convert the password to json
	jsonstr, err := json.Marshal(pwd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	// Write and return
	w.WriteHeader(http.StatusOK)
	w.Write(jsonstr)
	return nil
}