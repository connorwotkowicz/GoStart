package main

import (
	"flag"
	"fmt"

)

func main() {
	n := flag.Int("n", 1, "Number of quotes to fetch")
	save := flag.Bool("save", false, "Save quotes to quotes.json")
	load := flag.Bool("load", false, "Load and display saved quotes from quotes.json")
	search := flag.String("search", "", "Search saved quotes by keyword")
	flag.Parse()

	switch {
	case *search != "":
		quotes, err := loadQuotesFromFile("quotes.json")
		if err != nil {
			fmt.Println("Error loading quotes:", err)
			return
		}
		results := searchQuotes(quotes, *search)
		if len(results) == 0 {
			fmt.Println("No matching quotes found.")
			return
		}
		fmt.Printf(" Results for \"%s\":\n", *search)
		for i, q := range results {
			fmt.Printf("\n%d. \"%s\"\n   — %s\n", i+1, q.Content, q.Author)
		}
	case *load:
		quotes, err := loadQuotesFromFile("quotes.json")
		if err != nil {
			fmt.Println("Error loading quotes:", err)
			return
		}
		if len(quotes) == 0 {
			fmt.Println("No saved quotes found.")
			return
		}
		fmt.Println(" Saved Quotes:")
		for i, q := range quotes {
			fmt.Printf("\n%d. \"%s\"\n   — %s\n", i+1, q.Content, q.Author)
		}
	default:
		var collected []*Quote
		for i := 0; i < *n; i++ {
			quote, err := fetchQuote()
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Printf("\n \"%s\"\n   — %s\n", quote.Content, quote.Author)
			collected = append(collected, quote)
		}
		if *save {
			if err := saveQuotesToFile(collected, "quotes.json"); err != nil {
				fmt.Println("Error saving quotes:", err)
			} else {
				fmt.Println(" Quotes saved to quotes.json")
			}
		}
	}
}
