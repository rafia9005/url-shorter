package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type URLMapping struct {
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
}

func ReadData() ([]URLMapping, error) {
	file, err := os.Open("storage/db/data.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var mappings []URLMapping
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteValue, &mappings)
	if err != nil {
		return nil, err
	}
	return mappings, nil
}

func WriteData(mappings []URLMapping) error {
	data, err := json.MarshalIndent(mappings, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("storage/db/data.json", data, 0644)
}

func AddURL(shortURL, longURL string) error {
	mappings, err := ReadData()
	if err != nil {
		return err
	}

	mappings = append(mappings, URLMapping{ShortURL: shortURL, LongURL: longURL})

	return WriteData(mappings)
}

func FindLongURL(shortURL string) (string, error) {
	mappings, err := ReadData()
	if err != nil {
		return "", err
	}

	for _, mapping := range mappings {
		if mapping.ShortURL == shortURL {
			return mapping.LongURL, nil
		}
	}
	return "", fmt.Errorf("URL not found")
}
