package codeye

import (
	"bytes"
	"fmt"
	"html/template"
	"io"

	"github.com/emad-elsaid/codeye/models"
)

var tmpl *template.Template

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

type Table struct {
	Title  string
	Header []interface{}
	Rows   [][]interface{}
	Footer []interface{}
}

func Contributors(contributors models.Contributors) *Table {
	t := Table{
		Title:  "Contributors",
		Header: []interface{}{"Name", "Commits"},
		Rows:   make([][]interface{}, len(contributors)),
	}

	for i, c := range contributors {
		t.Rows[i] = []interface{}{c.Name(), c.Commits()}
	}

	t.Footer = []interface{}{
		fmt.Sprintf("%d Contributor", len(contributors)),
		contributors.TotalCommits(),
	}
	return &t
}

type Navbar struct {
	Name string
}
