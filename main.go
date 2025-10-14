package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"nixii.dev/zipp/api"
	"nixii.dev/zipp/save"
)

func main() {
	
	// Attempt making the save dir and file
	if err := save.Init() ; err != nil {
		fmt.Printf("Error initiating save: %s\n", err)
		return
	}

	// Start the server
	mux := http.NewServeMux()
	
	// Connect the function
	mux.HandleFunc("/pwd/", api.HandleRequest)

	// Init the server
	err := http.ListenAndServe(":2327", mux)

	// Check the error
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
	} else if err != nil {
		fmt.Printf("Error in the server: %s", err)
		os.Exit(1)
	}
}