package api

import (
	"net/http"

	"nixii.dev/zipp/requests"
	"nixii.dev/zipp/save"
)

func Put(w http.ResponseWriter, req *http.Request) error {

	// Get the request
	var data requests.PutRequest
	err := getJson(req, &data)
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

	// Load the file
	saves, err := save.ReadSaveFile(data.MasterPassword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	// Save
	saves.SetPassword(data.Website, data.Username, save.Password{
		Password:  *data.Password,
		Email: *data.Email,
	})

	// Set the password file
	err = save.WriteSaveFile(saves, data.MasterPassword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	// Success
	w.WriteHeader(http.StatusOK)
	return nil
}