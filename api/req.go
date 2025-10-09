/*
 * Handle generic request data
 */
package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func getJson(r *http.Request, t any) error {
	bodyText, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyText, t)
	if err != nil {
		return err
	}

	return nil
}