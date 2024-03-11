package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Serving static files from the frontend directory
	fs := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fs)

	// Endpoint /hello, which returns an HTML fragment
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<div>Hello, World!</div>`)
	})

	// Another endpoint /trigger_delay, which returns a string
	http.HandleFunc("/trigger_delay", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `sth`)
	})

	// Starting the server
	fmt.Println("Server is running on http://localhost:8888")
	http.ListenAndServe(":8888", nil)
}
