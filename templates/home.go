package templates

import (
	"html/template"
	"io"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func htmlTemplate(file string) (t *template.Template) {
	path := "assets/html/" + file
	asset := string(MustAsset(path))

	t, err := template.New(path).Parse(asset)
	check(err)

	return
}

var table *template.Template

func init() {
	table = htmlTemplate("table.html")
}

func Table(w io.Writer, data interface{}) {
	table.Execute(w, data)
}
