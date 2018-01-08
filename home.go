package codeye

import (
	"net/http"

	"github.com/emad-elsaid/codeye/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	page := Page{
		Title:    "Home",
		Template: "table",
		Navbar: Navbar{
			Name: models.CurrentProject().Name,
		},
		Data: Contributors(models.NewContributors()),
	}
	page.Render(w)
}
