package main

import (
	"log"
	"os"
	"portfolio/handlers"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/template"
)

func main() {

  err := godotenv.Load()
  if err != nil {
	  log.Fatal("Error loading .env file")
  }

	app := pocketbase.New()

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {

		registry := template.NewRegistry()
		se.Router.GET("/static/{path...}", apis.Static(os.DirFS("./static"), false))

		se.Router.GET("/", func(e *core.RequestEvent) error {
			return handlers.IndexHandler(app, registry, e)
		})

		se.Router.GET("/robots.txt", func(e *core.RequestEvent) error {
			return handlers.RobotsHandler(registry, e)
		})

		se.Router.GET("/about", func(e *core.RequestEvent) error {
			return handlers.AboutHandler(registry, e, app)
		})

		se.Router.GET("/tech", func(e *core.RequestEvent) error {
			return handlers.TechHandler(registry, e, app)
		})

		se.Router.GET("/posts/{slug}", func(e *core.RequestEvent) error {
			return handlers.PostHandler(registry, e, app)
		})

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
