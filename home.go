package codeye

import (
	"net/http"

	"github.com/emad-elsaid/codeye/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Contributors models.Contributors
		FirstCommit  *models.Commit
		Files        models.Files
	}{
		models.NewContributors(),
		models.FirstCommit(),
		models.FilesList(),
	}
	page := Page{
		Title:    "General Overview",
		Template: "home",
		Navbar: Navbar{
			Name: models.CurrentProject().Name,
		},
		Data: data,
	}
	page.Render(w)
}
