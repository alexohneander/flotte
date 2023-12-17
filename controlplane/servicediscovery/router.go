package servicediscovery

import (
	"log"
	"net/http"

	"gorm.io/gorm"
)

func RegisterHandler(mux *http.ServeMux, db *gorm.DB) *http.ServeMux {
	log.Println("HTTP Router: Registering Service Discovery Handler")

	// Register Handler
	mux.HandleFunc("/servicediscovery/worker/register", func(w http.ResponseWriter, r *http.Request) {
		RegisterWorker(w, r, db)
	})
	mux.HandleFunc("/servicediscovery/unregister", func(w http.ResponseWriter, r *http.Request) {
		Unregister(w, r, db)
	})

	return mux
}
