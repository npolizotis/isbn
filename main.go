package main

import (
	"log"
	"os"
)

func main() {
	var errCode int
	if len(os.Args) != 2 {
		log.Printf("expecting an ISBN as arg")
		os.Exit(2)
	}
	if isbn := os.Args[1]; !IsValidISBN(isbn) {
		log.Fatalf("Invalid:%s", isbn)
	}

	os.Exit(errCode)
}
