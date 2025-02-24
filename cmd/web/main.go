package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/lucasfpascoali/go-web-application/pkg/config"
	"github.com/lucasfpascoali/go-web-application/pkg/handlers"
	"github.com/lucasfpascoali/go-web-application/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var sessionManager *scs.SessionManager

// main is the main function
func main() {
	app.InProduction = false

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.InProduction

	app.SessionManager = sessionManager

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("error creating template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Printf("Starting application on port %s\n", portNumber)
	err = srv.ListenAndServe()
	log.Fatal(err)
}
