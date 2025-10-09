package api

import (
	"fmt"
	"net/http"

	"nixii.dev/zipp/requests"
)

func Get(w http.ResponseWriter, req *http.Request) error {
	var data requests.GetRequest
	err := getJson(req, &data)
	if err != nil {
		return err
	}

	fmt.Println(data)

	return nil
}