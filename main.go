package main

import (
	"fmt"
	"net/http"
	"url-shorter/handler"

	"github.com/gorilla/handlers"
)

func main() {
	http.HandleFunc("/create", handler.CreateShortURLHandler)
	http.HandleFunc("/", handler.RedirectHandler)

	corsOptions := handlers.AllowedOrigins([]string{"*"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	corsHeaders := handlers.AllowedHeaders([]string{"Content-Type"})
	fmt.Println("SERVER  IS RUNNING CIHUY ASELOLE")
	http.ListenAndServe(":8080", handlers.CORS(corsOptions, corsMethods, corsHeaders)(http.DefaultServeMux))
}
