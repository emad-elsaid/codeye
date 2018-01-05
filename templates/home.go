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

var table *template.Template

func init() {
	var err error
	table, err = template.New("home").Parse("<table>{{ range .}}<tr> {{range .}} <td>{{.}}</td> {{end}} </tr>{{end}}</table>")
	check(err)
}

func Table(w io.Writer, data interface{}) {
	table.Execute(w, data)
}
