package main

import (
	"flotte/controlplane/servicediscovery"
	"log"
	"net/http"
)

func main() {

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register Handler
	mux = servicediscovery.RegisterHandler(mux)

	log.Println("Starting flotte, control plane initialized")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
