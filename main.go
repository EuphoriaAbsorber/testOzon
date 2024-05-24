package main

import (
	"log"
	"net/http"
)

func main() {

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic("Error when starting the http server", err)
	}
	log.Print("Server is started")
}
