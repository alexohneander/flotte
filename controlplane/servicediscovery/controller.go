package servicediscovery

import (
	"encoding/json"
	"flotte/controlplane/model"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var RegToken string = GetToken()

func RegisterWorker(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
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
				log.Println("Service Discovery: Token is correct")
				id := uuid.New().String()
				db.Create(&model.WorkerPlane{
					ID:      id,
					Name:    jsonBody["Name"].(string),
					IP:      jsonBody["IP"].(string),
					APIPort: jsonBody["APIPort"].(string),
				})

				log.Println("Service Discovery: Worker registered")
				var worker model.WorkerPlane
				db.First(&worker, "ID = ?", id)

				io.WriteString(w, fmt.Sprintf("%v", worker))
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
