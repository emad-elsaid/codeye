package handlers

import (
	"net/http"

	"github.com/emad-elsaid/codeye/models"
	"github.com/emad-elsaid/codeye/templates"
)

func Home(w http.ResponseWriter, r *http.Request) {
	page := templates.Page{
		Title:    "Home",
		Template: "table",
		Data:     models.Contributors(),
	}
	page.Render(w)
}
