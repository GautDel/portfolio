package handlers

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)


func getRichText(app *pocketbase.PocketBase, id string) (*core.Record, error) {

	record, err := app.FindRecordById("pages", id)
	if err != nil {
		return nil, err
	}

	return record, nil
}
