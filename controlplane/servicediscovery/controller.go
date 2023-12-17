package servicediscovery

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var RegToken string = InitializedToken()

func Register(w http.ResponseWriter, r *http.Request) {
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
				log.Println("ServiceDiscovery: Wrong Token")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			} else {
				log.Println("ServiceDiscovery: Token is correct")
			}
		}

		log.Println("ServiceDiscovery: Registering Worker..")
		io.WriteString(w, "Hello, HTTP!\n")
	} else {
		log.Println("ServiceDiscovery: Wrong HTTP Method")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func Unregister(w http.ResponseWriter, r *http.Request) {
	log.Println("ServiceDiscovery: Unregistering Worker..")
	io.WriteString(w, "Hello, HTTP!\n")
}
