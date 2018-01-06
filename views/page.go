package views

import (
	"bytes"
	"html/template"
	"io"
)

const HTMLPath string = "assets/html"

var tmpl *template.Template

func loadHTML(tmpl *template.Template) {
	files, err := AssetDir(HTMLPath)
	checkError(err)

	for _, file := range files {
		path := HTMLPath + "/" + file
		asset := string(MustAsset(path))

		_, err := tmpl.Parse(asset)
		checkError(err)
	}
}

func init() {
	tmpl = template.New("views")
	loadHTML(tmpl)
}

type Page struct {
	Title    string
	Navbar   Navbar
	Template string
	Data     interface{}
	Body     template.HTML
}

func (this *Page) Render(w io.Writer) {

	body := bytes.NewBufferString("")
	err := tmpl.ExecuteTemplate(body, this.Template, this.Data)
	checkError(err)

	this.Body = template.HTML(body.String())

	tmpl.ExecuteTemplate(w, "layout", this)
}
