package service

import (
	"math/rand"
	"time"
	"url-shorter/storage"
)

const baseURL = "http://localhost:8080/"

func GenerateShortURL() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	length := 6
	var shortURL []rune
	for i := 0; i < length; i++ {
		shortURL = append(shortURL, chars[rand.Intn(len(chars))])
	}
	return string(shortURL)
}

func CreateShortURL(longURL string) (string, string, error) {
	var shortCode string
	var err error

	for {
		shortCode = GenerateShortURL()
		err = storage.AddURL(shortCode, longURL)
		if err == nil {
			break
		}
	}

	fullShortURL := baseURL + shortCode
	return fullShortURL, shortCode, nil
}

func GetLongURL(shortCode string) (string, error) {
	return storage.FindLongURL(shortCode)
}
