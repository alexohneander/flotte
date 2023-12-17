package main

import (
	"flotte/controlplane/database"
	"flotte/controlplane/servicediscovery"
	"log"
	"net/http"
)

var RegToken string = servicediscovery.InitializedToken()

func main() {
	// Initialize Database
	db := database.InitializeDatabase()

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register Handler
	mux = servicediscovery.RegisterHandler(mux, db)

	log.Println("Starting flotte, control plane initialized")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
