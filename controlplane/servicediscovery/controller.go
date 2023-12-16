package servicediscovery

import (
	"io"
	"log"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	log.Println("ServiceDiscovery: Registering new Worker..")
	io.WriteString(w, "Hello, HTTP!\n")
}

func Unregister(w http.ResponseWriter, r *http.Request) {
	log.Println("ServiceDiscovery: Unregistering Worker..")
	io.WriteString(w, "Hello, HTTP!\n")
}
