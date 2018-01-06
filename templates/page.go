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

type Page struct {
	Title    string
	Template string
	Data     interface{}
	Body     template.HTML
}

func (this *Page) Render(w io.Writer) {

	body := bytes.NewBufferString("")
	tmpl.ExecuteTemplate(body, this.Template, this.Data)
	this.Body = template.HTML(body.String())

	tmpl.ExecuteTemplate(w, "layout", this)
}
