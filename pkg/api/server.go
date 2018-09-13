package api

import (
	"log"
	"net/http"
	"os"
)

func Run() {
	router := NewRouter()
	var port string
	if port = os.Getenv("SOTERIA_API_PORT"); port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}
