package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>My name is xoxo</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About xoxo</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server is about to start...")
	http.ListenAndServe(":4000", nil)
}
