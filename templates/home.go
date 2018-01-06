package templates

import (
	"bytes"
	"html/template"
	"io"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func htmlTemplates(tmpl *template.Template) {
	files, err := AssetDir("assets/html")
	check(err)

	for _, file := range files {
		path := "assets/html/" + file
		asset := string(MustAsset(path))

		_, err := tmpl.Parse(asset)
		check(err)
	}
}

var tmpl *template.Template

func init() {
	tmpl = template.New("templates")
	htmlTemplates(tmpl)
}

func Page(w io.Writer, title string, templateName string, data interface{}) {

	body := bytes.NewBufferString("")
	tmpl.ExecuteTemplate(body, templateName, data)

	page := struct {
		Title string
		Body  template.HTML
	}{
		title,
		template.HTML(body.String()),
	}
	tmpl.ExecuteTemplate(w, "layout", page)
}
