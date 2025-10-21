package api

import (
	"fmt"
	"net/http"

	"nixii.dev/zipp/requests"
	"nixii.dev/zipp/save"
)

func Patch(w http.ResponseWriter, r *http.Request) error {
	
	// Get the request
	var data requests.PatchRequest
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
	pwd, err := saves.GetAccountInfo(data.From.Website, data.From.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return err
	}

	// Set stuff
	if data.To.Email != "" {
		pwd.Email = data.To.Email
	} else if data.To.Password != "" {
		pwd.Password = data.To.Password
	}

	// To loc
	toSite := data.From.Website
	toUser := data.From.Username
	if data.To.Website != "" {
		toSite = data.To.Website
	}
	if data.To.Username != "" {
		toUser = data.To.Username
	}

	// Just remove and replace the new password lol
	_ = saves.RemovePassword(data.From.Website, data.From.Username)
	fmt.Println(data.To)
	err = saves.SetPassword(toSite, toUser, *pwd)
	if err != nil {
		return err
	}
	fmt.Println(saves)

	// Write the file
	save.WriteSaveFile(saves, data.MasterPassword)

	// Woohoooo
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))

	// Return
	return nil
}