package main

import (
	"fmt"
	"github.com/lucasfpascoali/go-web-application/pkg/config"
	"github.com/lucasfpascoali/go-web-application/pkg/handlers"
	"github.com/lucasfpascoali/go-web-application/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("error creating template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	err = http.ListenAndServe(portNumber, nil)
	if err != nil {
		return
	}
}
