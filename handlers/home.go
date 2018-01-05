package handlers

import (
	"net/http"

	"github.com/emad-elsaid/codeye/models"
	"github.com/emad-elsaid/codeye/templates"
)

func Home(w http.ResponseWriter, r *http.Request) {
	contributors := models.Contributors()
	templates.Table(w, contributors)
}
