package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Quote struct {
	Content string `json:"content"`
	Author  string `json:"author"`
}

func fetchQuote() (*Quote, error) {
	resp, err := http.Get("https://zenquotes.io/api/random")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []struct {
		Q string `json:"q"`
		A string `json:"a"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no quotes found")
	}

	return &Quote{
		Content: result[0].Q,
		Author:  result[0].A,
	}, nil
}
