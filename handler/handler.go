package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"url-shorter/service"
)

type URLRequest struct {
	URL string `json:"url"`
}

func CreateShortURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var urlRequest URLRequest

		err := json.NewDecoder(r.Body).Decode(&urlRequest)
		if err != nil {
			log.Printf("Error decoding JSON: %v", err)
			http.Error(w, "Invalid input data. Please provide a valid URL.", http.StatusBadRequest)
			return
		}

		if urlRequest.URL == "" {
			http.Error(w, "Invalid input data. Please provide a valid URL.", http.StatusBadRequest)
			return
		}

		shortURL, err := service.CreateShortURL(urlRequest.URL)
		if err != nil {
			log.Printf("Error creating short URL: %v", err)
			http.Error(w, "Failed to create short URL", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"short_url": shortURL}
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]

	longURL, err := service.GetLongURL(shortURL)
	if err != nil {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longURL, http.StatusMovedPermanently)
}
