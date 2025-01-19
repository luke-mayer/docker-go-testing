package main

import (
	"fmt"
	"log"
	"net/http"
)

// var PORT = os.Getenv("PORT")
var PORT = 8080

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world\n")
}

func main() {
	http.HandleFunc("/", helloHandler)
	err := http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil)
	if err != nil {
		log.Fatalf("Error calling ListenAndServe(): %v", err)
	}
	log.Printf("Server running on port: %v\n", PORT)
}
