package servicediscovery

import (
	"log"
	"net/http"
)

func RegisterHandler(mux *http.ServeMux) *http.ServeMux {
	log.Println("Registering Service Discovery Handler")

	// Register Handler
	mux.HandleFunc("/servicediscovery/register", Register)

	return mux
}
