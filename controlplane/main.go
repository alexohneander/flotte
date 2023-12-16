package main

import (
	"flotte/controlplane/servicediscovery"
	"log"
	"net/http"
)

func main() {

	// Create Register Token if not exists
	servicediscovery.InitializedToken()

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
