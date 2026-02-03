package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"nixii.dev/zipp/api"
	"nixii.dev/zipp/crypt"
	"nixii.dev/zipp/save"
)

func main() {
	
	// Attempt making the save dir and file
	if err := save.Init() ; err != nil {
		log.Fatalf("Error initiating save: %s", err)
	}

	// Attempt getting the secure stuff
	if err := crypt.Init() ; err != nil {
		log.Fatalf("failed to load crypt: %s", err)
	}

	// Start the server
	mux := http.NewServeMux()

	// Create the self-signed certificate
	// tlsConf := &tls.Config{
	// 	Certificates: []tls.Certificate{crypt.MyCertificate},
	// }
	
	// Connect the function
	mux.HandleFunc("/pwd/", api.HandlePwdRequest)
	mux.HandleFunc("/cmp/", api.HandleCmpRequest)
	mux.HandleFunc("/allpwds/", api.HandleAllRequest)
	mux.HandleFunc("/transf/", api.HandleTransfRequest)
	mux.HandleFunc("/recv/", api.HandleRecvRequest)

	// Make a server
	server := &http.Server{
		Addr: ":2327",
		Handler: mux,
	}

	// Init the server
	defer server.Close()
	err := server.ListenAndServe()

	// Check the error
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
	} else if err != nil {
		fmt.Printf("Error in the server: %s", err)
		os.Exit(1)
	}
}
