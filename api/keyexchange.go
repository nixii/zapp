package api

import (
	"net/http"

	"nixii.dev/zipp/requests"
)

func GetKeys(w http.ResponseWriter, r *http.Request) error {

	// Get the request
	var data requests.GetKey
	err := getJson(r, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	// Check
	if err := data.VerifyRequest(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	// Marshal the key
	// key, err := x509.ParsePKC // TODO: parse the key

	return nil
}