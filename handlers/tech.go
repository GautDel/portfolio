package handlers

import (
	"net/http"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/template"
)

func TechHandler(registry *template.Registry, e *core.RequestEvent, app *pocketbase.PocketBase) error {

	record, err := getRichText(app, "dp8a14u8sz15a70")

	page := Page{
		Page: record.GetString("page"),
		RichText: record.GetString("richtext"),
		Description: record.GetString("description"),
	}

	html, err := registry.LoadFiles(
		"templates/tech.html",
		"templates/partials/home_button.html",
		"templates/partials/footer.html",
	).Render(page)

	if err != nil {
		// or redirect to a dedicated 404 HTML page
		return e.NotFoundError("", err)
	}

	return e.HTML(http.StatusOK, html)

}
