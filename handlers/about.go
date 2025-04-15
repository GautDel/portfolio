package handlers

import (
	"net/http"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/template"
)

func AboutHandler(registry *template.Registry, e *core.RequestEvent, app *pocketbase.PocketBase) error {

	record, err := getRichText(app, "887p6l7g1s2z71k")


	richText := RichText{
		RichText: record.GetString("richtext"),
	}

	html, err := registry.LoadFiles(
		"templates/about.html",
		"templates/partials/home_button.html",
		"templates/partials/footer.html",
	).Render(richText)

	if err != nil {
		// or redirect to a dedicated 404 HTML page
		return e.NotFoundError("test", err)
	}

	return e.HTML(http.StatusOK, html)

}
