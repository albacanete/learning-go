package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my home page")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
