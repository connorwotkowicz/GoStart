# Quotify

A lightweight Go CLI tool to fetch, search, and save inspirational quotes from the ZenQuotes API.

<p align="left">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
  <img src="https://img.shields.io/badge/CLI-555555?style=for-the-badge&logo=gnubash&logoColor=white" />
  <img src="https://img.shields.io/badge/JSON-000000?style=for-the-badge&logo=json&logoColor=white" />
  <img src="https://img.shields.io/badge/ZenQuotes.io-003049?style=for-the-badge&logo=data&logoColor=white" />
  <img src="https://img.shields.io/badge/VSCode-007ACC?style=for-the-badge&logo=visual%20studio%20code&logoColor=white" />
  <img src="https://img.shields.io/badge/Tested%20with-Go%20testing-00ADD8?style=for-the-badge" />
</p>


## Features

- Fetch one or more random quotes from ZenQuotes.io
- Save quotes to a local `quotes.json` file
- Load and view saved quotes
- Search saved quotes by keyword (author or text)
- Includes basic unit tests
- Modular code structure across multiple files

## Usage

### Fetch one or more quotes
```bash
go run main.go -n=3
```

### Load and display all saved quotes 
```bash
go run main.go -n=2 -save
```

### Load and display all saved quotes 
```bash
go run main.go -load
```

### Search saved quotes by keyword
```bash
go run main.go -search="life"
```

### Run tests
```bash
go test
```

### Build binary 
```bash
go build -o quotify
./quotify -n=1 -save
```

## Structure 
```pgsql
main.go        # CLI entry point and flag routing
quote.go       # Fetching quotes from the API
storage.go     # Save/load quotes to/from local JSON
cli.go         # Search + helper CLI logic
main_test.go   # Unit tests
quotes.json    # Local saved quotes
```

## Stack
- Language: Go
- Data source: ZenQuotes.io
- File I/O: os, encoding/json
- CLI: flag package
- Tests: testing package

