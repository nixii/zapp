package api

import (
	"net/http"
	"io"
	"nixii.dev/zapp/save"
)

func RecvRequest(w http.ResponseWriter, req *http.Request) error {

	// Get the data
	bodyText, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	// write the file
	err = save.WriteSaveFileBytes(bodyText)

	// success
	return nil
}
