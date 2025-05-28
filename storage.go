package main

import (
	"encoding/json"
	"os"
)

func saveQuotesToFile(quotes []*Quote, filename string) error {
	var existing []*Quote
	if file, err := os.ReadFile(filename); err == nil {
		_ = json.Unmarshal(file, &existing)
	}
	allQuotes := append(existing, quotes...)

	data, err := json.MarshalIndent(allQuotes, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func loadQuotesFromFile(filename string) ([]*Quote, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var quotes []*Quote
	if err := json.Unmarshal(data, &quotes); err != nil {
		return nil, err
	}

	return quotes, nil
}
