package servicediscovery

import (
	"encoding/json"
	"flotte/controlplane/model"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var RegToken string = GetToken()

func RegisterWorker(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Set Header
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		if r.Body != nil {
			defer r.Body.Close()
			// Parse JSON Body
			jsonBody := make(map[string]interface{})
			err := json.NewDecoder(r.Body).Decode(&jsonBody)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Check if token is correct
			if jsonBody["Token"] != RegToken {
				log.Println("Service Discovery: Wrong Token")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			} else {
				id := uuid.New().String()
				result := db.Create(&model.WorkerPlane{
					ID:      id,
					Name:    jsonBody["Name"].(string),
					IP:      jsonBody["IP"].(string),
					APIPort: jsonBody["APIPort"].(string),
				})

				if result.Error != nil {
					log.Println("Service Discovery: Worker could not be registered")
					http.Error(w, result.Error.Error(), http.StatusInternalServerError)
					return
				}

				log.Println("Service Discovery: Worker registered")
				var worker model.WorkerPlane
				db.First(&worker, "ID = ?", id)
				json.NewEncoder(w).Encode(worker)

				return
			}
		}

		io.WriteString(w, "Hello, HTTP!\n")
	} else {
		log.Println("Service Discovery: Wrong HTTP Method")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func Unregister(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}
