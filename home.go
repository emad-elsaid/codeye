package codeye

import (
	"net/http"

	"github.com/emad-elsaid/codeye/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Contributors models.Contributors
		FirstCommit  *models.Commit
	}{
		models.NewContributors(),
		models.FirstCommit(),
	}
	page := Page{
		Title:    "Home",
		Template: "home",
		Navbar: Navbar{
			Name: models.CurrentProject().Name,
		},
		Data: data,
	}
	page.Render(w)
}
