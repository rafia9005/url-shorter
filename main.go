package main

import (
	"fmt"
	"log"
	"net/http"
	"url-shorter/handler"
)

func main() {
	http.HandleFunc("/", handler.RedirectHandler)

	http.HandleFunc("/create", handler.CreateShortURLHandler)

	fmt.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
