package handlers

import (
	"net/http"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/template"
)

type PostFull struct {
	Title string
	Slug string
	Content string
	Created string
}

func PostHandler(registry *template.Registry, e *core.RequestEvent, app *pocketbase.PocketBase) error {

	test := e.Request.PathValue("slug")

	record, err := app.FindFirstRecordByData("posts", "slug", test)

	post := PostFull{
		Title: record.GetString("title"),
		Slug: record.GetString("slug"),
		Content: record.GetString("content"),
		Created: record.GetString("created"),
	}

	html, err := registry.LoadFiles(
		"templates/post.html",
		"templates/partials/home_button.html",
		"templates/partials/footer.html",
	).Render(post)

	if err != nil {
		// or redirect to a dedicated 404 HTML page
		return e.NotFoundError("", err)
	}

	return e.HTML(http.StatusOK, html)

}



