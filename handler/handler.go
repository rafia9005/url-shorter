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

type URLResponse struct {
	Status    string `json:"status"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	ShortURL  string `json:"short_url,omitempty"`
	ShortCode string `json:"shortCode"`
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

		shortURL, shortCode, err := service.CreateShortURL(urlRequest.URL)
		if err != nil {
			log.Printf("Error creating short URL for %s: %v", urlRequest.URL, err)
			http.Error(w, "Failed to create short URL", http.StatusInternalServerError)
			return
		}

		response := URLResponse{
			Status:    "success",
			Code:      200,
			Message:   "Short URL created successfully",
			ShortURL:  shortURL,
			ShortCode: shortCode,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[1:]

	longURL, err := service.GetLongURL(shortCode)
	if err != nil {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	response := map[string]interface{}{
		"status":    "success",
		"code":      200,
		"message":   "URL found",
		"url":       longURL,
		"shortCode": shortCode,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
