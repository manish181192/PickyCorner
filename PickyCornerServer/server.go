package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
