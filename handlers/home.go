package handlers

import (
	"net/http"

	"github.com/emad-elsaid/codeye/templates"
)

func Home(w http.ResponseWriter, r *http.Request) {
	templates.Table(w, [][]string{{"hello", "1"}, {"world", "2"}})
}
