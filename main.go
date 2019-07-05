package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/StevenACoffman/go-gql-todo/internal/server"
)

func main() {

	// Setting up the server
	port := 8081
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      server.Handler(),
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		IdleTimeout:  time.Second * 30,
	}

	log.Println(server.ListenAndServe())
}
