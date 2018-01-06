package controllers

import (
	"net/http"

	"github.com/emad-elsaid/codeye/models"
	"github.com/emad-elsaid/codeye/views"
)

func Home(w http.ResponseWriter, r *http.Request) {
	page := views.Page{
		Title:    "Home",
		Template: "table",
		Navbar:   models.CurrentProject(),
		Data:     views.Contributors(models.NewContributors()),
	}
	page.Render(w)
}
