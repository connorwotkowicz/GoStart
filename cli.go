package main

import "strings"

func searchQuotes(quotes []*Quote, keyword string) []*Quote {
	var results []*Quote
	for _, q := range quotes {
		if containsIgnoreCase(q.Content, keyword) || containsIgnoreCase(q.Author, keyword) {
			results = append(results, q)
		}
	}
	return results
}

func containsIgnoreCase(text, keyword string) bool {
	return len(keyword) > 0 && len(text) > 0 && contains(strings.ToLower(text), strings.ToLower(keyword))
}

func contains(haystack, needle string) bool {
	return len(needle) > 0 && strings.Contains(haystack, needle)
}
