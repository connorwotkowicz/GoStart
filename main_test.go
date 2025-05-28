package main

import (
	"os"
	"testing"
)

func TestSearchQuotes(t *testing.T) {
	quotes := []*Quote{
		{Content: "Life is short", Author: "Somebody"},
		{Content: "Just do it", Author: "Nike"},
	}

	results := searchQuotes(quotes, "life")
	if len(results) != 1 {
		t.Errorf("Expected 1 result, got %d", len(results))
	}
}

func TestSaveAndLoadQuotes(t *testing.T) {
	tmpFile := "test_quotes.json"
	defer os.Remove(tmpFile)

	original := []*Quote{
		{Content: "Code is poetry", Author: "Anonymous"},
	}

	err := saveQuotesToFile(original, tmpFile)
	if err != nil {
		t.Errorf("Failed to save: %v", err)
	}

	loaded, err := loadQuotesFromFile(tmpFile)
	if err != nil {
		t.Errorf("Failed to load: %v", err)
	}

	if len(loaded) != 1 || loaded[0].Content != "Code is poetry" {
		t.Errorf("Loaded data mismatch: %+v", loaded)
	}
}
