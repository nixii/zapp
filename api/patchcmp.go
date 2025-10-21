package api

import (
	"net/http"

	"nixii.dev/zipp/requests"
	"nixii.dev/zipp/save"
)

func PatchCmp(w http.ResponseWriter, r *http.Request) error {

	// Get the request
	var data requests.PatchCmpRequest
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

	// Change the master password
	file, err := save.ReadSaveFile(data.MasterPassword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not read"))
		return err
	}
	
	err = save.WriteSaveFile(file, data.NewMasterPassword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not write"))
		return err
	}

	// Yayy successs
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("changed master password"))
	return nil
}