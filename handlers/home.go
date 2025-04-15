package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"portfolio/utils"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/template"
)


func getRecords(
	recType string,
	recAmount int,
	recOrderField string,
	app *pocketbase.PocketBase,
) ([]*core.Record, error) {
	records := []*core.Record{}

	err := app.RecordQuery(recType).
		OrderBy(recOrderField + " DESC").
		Limit(int64(recAmount)).
		All(&records)

	if err != nil {
		return nil, err
	}

	return records, nil
}

func getProjects(e *core.RequestEvent, app *pocketbase.PocketBase) (Projects, error) {

	records, err := getRecords("projects", 6, "coded_date", app)
	if err != nil {
		return nil, e.NotFoundError("", err)
	}

	var projects Projects

	for _, record := range records {

		errs := app.ExpandRecord(record, []string{"project_type"}, nil)
		if len(errs) > 0 {
			return nil, fmt.Errorf("Failed to expand: %v", errs)
		}

		projType := record.ExpandedOne("project_type")	

		project := Project{
			Name: record.GetString("name"),
			CodedDate: record.GetInt("coded_date"),
			ProjectType: projType.GetString("name"),
			Url: record.GetString("url"),
		}

		projects = append(projects, project)
	}

	return projects, nil
}

func getPostSummaries(e *core.RequestEvent, app *pocketbase.PocketBase) (PostSummaries, error) {
	records, err := getRecords("posts", 6, "created", app)
	if err != nil {
		return nil, e.NotFoundError("", err)
	}

	var postSummaries PostSummaries

	for _, record := range records {

		formattedDate, err := utils.DateFormat(record.GetString("created"), "Jan 2, 2006")
		if err != nil {
			log.Println("Failed to format date time", err)
		}

		postSummary := PostSummary{
			Title: record.GetString("title"),
			Created: formattedDate,
			Slug: record.GetString("slug"),
		}

		postSummaries = append(postSummaries, postSummary)
	}

	return postSummaries, nil
}

func getHomeImage(app *pocketbase.PocketBase) (Image, error) {
	record, err := app.FindFirstRecordByData("images", "position", 999)
	if err != nil {
		return Image{}, err
	}

	baseUrl := os.Getenv("BASE_URL")
	imgPath := record.BaseFilesPath() + "/" + record.GetString("image")
	fullPath := baseUrl + "/api/files/" + imgPath

	image := Image{
		Image: fullPath,
		Alt: record.GetString("alt"),
	}

	return image, nil
}

func getPreviewImages(app *pocketbase.PocketBase) (Images, error) {
	records, err := app.FindRecordsByFilter(
		"images",
		"position >= {:min} && position <= {:max}",
		"position",
		4,
		0,
		dbx.Params{"min": 0, "max": 3,})

	if err != nil {
		// or redirect to a dedicated 404 HTML page
		return nil, err
	}
	
	var images Images
	baseUrl := os.Getenv("BASE_URL")

	for _, record := range records {
		imgPath := record.BaseFilesPath() + "/" + record.GetString("image")
		fullPath := baseUrl + "/api/files/" + imgPath

		image := Image{
			Image: fullPath,
			Alt: record.GetString("alt"),
		}

		images = append(images, image)
	}

	return images, nil
}


func IndexHandler(
	app *pocketbase.PocketBase, 
	registry *template.Registry, 
	e *core.RequestEvent,
) error {

	image, err := getHomeImage(app)
	images, err := getPreviewImages(app)
	projects, err := getProjects(e, app) 
	posts, err := getPostSummaries(e, app) 


	html, err := registry.LoadFiles(
				"templates/index.html", 
				"templates/partials/projects.html",
				"templates/partials/preview.html",
				"templates/partials/posts.html",
				"templates/partials/footer.html",
			).Render(map[string]any{"Projects": projects, "Posts": posts, "Images": images, "Image": image})

	if err != nil {
		// or redirect to a dedicated 404 HTML page
		return e.NotFoundError("", err)
	}

	return e.HTML(http.StatusOK, html)
}
