package main

import (
	"fmt"
	"log"
	"net/http"
)

// var port = os.Getenv("PORT")
var port = 8080

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world\n")
}

func main() {
	http.HandleFunc("/", helloHandler)
	log.Printf("Server running on port: %v\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Fatalf("Error calling ListenAndServe(): %v", err)
	}
}
