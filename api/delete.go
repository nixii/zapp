package api

import (
	"net/http"

	"nixii.dev/zipp/requests"
	"nixii.dev/zipp/save"
)

// TODO: finish

func Delete(w http.ResponseWriter, r *http.Request) error {

	// Get the request
	var data requests.DeleteRequest
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

	// Perform the reqeust
	saves, err := save.ReadSaveFile(data.MasterPassword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return err
	}

	// Get the password
	err = saves.RemovePassword(data.Website, data.Username)

	// code
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	// Save the file
	err = save.WriteSaveFile(saves, data.MasterPassword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)

	// return thing
	return nil
}