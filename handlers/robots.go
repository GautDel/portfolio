package handlers

import (
	"net/http"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/template"
)

func RobotsHandler(registry *template.Registry, e *core.RequestEvent) error {

	html, err := registry.LoadFiles(
		"static/robots.txt",
	).Render(nil)

	if err != nil {
		return e.NotFoundError("test", err)
	}

	return e.HTML(http.StatusOK, html)
}
