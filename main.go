package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"url-shorter/middleware"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	logging := middleware.LoggingMiddleware(http.DefaultServeMux)

	fmt.Println("SERVER IS RUNNING ON PORT 8080")
	log.Fatal(http.ListenAndServe(":8080", logging))
}
