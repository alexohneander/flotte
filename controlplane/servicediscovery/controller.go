package servicediscovery

import (
	"io"
	"log"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	log.Println("Registering new Worker..")
	io.WriteString(w, "Hello, HTTP!\n")
}
