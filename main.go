package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "5/27/2025 4:14 PM.\n")
	}

	http.HandleFunc("/", helloHandler)
    log.Println("Listing for requests at http://localhost:8080/hello")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
