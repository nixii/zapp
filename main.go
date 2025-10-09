package main

import (
	"fmt"

	"nixii.dev/zipp/save"
)

func main() {
	
	// Attempt making the save dir and file
	if err := save.Init() ; err != nil {
		fmt.Printf("Error initiating save: %s\n", err)
		return
	}
}