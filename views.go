package codeye

import (
	"bytes"
	"html/template"
	"io"
)

var tmpl *template.Template

func init() {
	tmpl = template.New("views")
	loadHTML(tmpl)
}

func loadHTML(tmpl *template.Template) {
	const HTMLPath string = "assets/html"
	files, err := AssetDir(HTMLPath)
	checkError(err)

	for _, file := range files {
		path := HTMLPath + "/" + file
		asset := string(MustAsset(path))

		_, err := tmpl.Parse(asset)
		checkError(err)
	}
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

type Navbar struct {
	Name string
}
