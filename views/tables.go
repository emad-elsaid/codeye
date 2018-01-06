package views

import (
	"fmt"

	"github.com/emad-elsaid/codeye/models"
)

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
