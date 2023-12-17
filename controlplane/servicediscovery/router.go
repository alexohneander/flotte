package servicediscovery

import (
	"log"
	"net/http"

	"gorm.io/gorm"
)

func RegisterHandler(mux *http.ServeMux, db *gorm.DB) *http.ServeMux {
	log.Println("Registering Service Discovery Handler")

	// Register Handler
	mux.HandleFunc("/servicediscovery/register", func(w http.ResponseWriter, r *http.Request) {
		Register(w, r, db)
	})
	mux.HandleFunc("/servicediscovery/unregister", func(w http.ResponseWriter, r *http.Request) {
		Unregister(w, r, db)
	})

	return mux
}
