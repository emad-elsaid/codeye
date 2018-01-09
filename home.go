package codeye

import (
	"net/http"

	"github.com/emad-elsaid/codeye/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	page := Page{
		Title:    "Home",
		Template: "home",
		Navbar: Navbar{
			Name: models.CurrentProject().Name,
		},
		Data: models.NewContributors(),
	}
	page.Render(w)
}
