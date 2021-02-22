package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	message := os.Getenv("APP_RESPONSE_MESSAGE")
	if message == "" {
		message = "Hello, World!"
	}
	fmt.Fprint(w, message)
}

func main() {
	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr: host + ":" + port,
	}

	http.HandleFunc("/", hello)

	fmt.Printf("Start Server. host:%s port:%s\n", host, port)
	log.Fatal(server.ListenAndServe())
}
