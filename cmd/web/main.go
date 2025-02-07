package main

import (
	"fmt"
	"github.com/lucasfpascoali/go-web-application/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		return
	}
}
