package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Tekton Pipeline")
}

func main() {
	host := ""
	port := "8080"

	server := &http.Server{
		Addr: host + ":" + port,
	}

	http.HandleFunc("/", hello)

	fmt.Printf("Start Server. host:%s port:%s\n", host, port)
	log.Fatal(server.ListenAndServe())
}
