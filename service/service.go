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

func CreateShortURL(longURL string) (string, error) {
	shortURL := GenerateShortURL()

	err := storage.AddURL(shortURL, longURL)
	if err != nil {
		return "", err
	}
	return baseURL + shortURL, nil
}

func GetLongURL(shortURL string) (string, error) {
	return storage.FindLongURL(shortURL)
}
