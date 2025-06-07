package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		undefinedFunction() 
		io.WriteString(w, "5/28/2025 10:38 AM.\n")
	}

	http.HandleFunc("/", helloHandler)
    log.Println("Listing for requests at http://localhost:8080/hello")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
