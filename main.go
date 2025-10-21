package main

import (
	"crypto/tls"
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

	// TLS conf
	tlsconf := &tls.Config{
		Certificates: []tls.Certificate{crypt.MyCertificate},
		MinVersion: tls.VersionTLS13,
	}

	// Start the server
	mux := http.NewServeMux()
	
	// Connect the function
	mux.HandleFunc("/pwd/", api.HandlePwdRequest)
	mux.HandleFunc("/cmp/", api.HandleCmpRequest)
	mux.HandleFunc("/allpwds/", api.HandleAllRequest)

	// Make a server
	server := &http.Server{
		Addr: ":2327",
		Handler: mux,
		TLSConfig: tlsconf,
	}

	// Init the server
	err := server.ListenAndServeTLS("", "")

	// Check the error
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
	} else if err != nil {
		fmt.Printf("Error in the server: %s", err)
		os.Exit(1)
	}
}