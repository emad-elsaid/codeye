package models

import (
	"os"
	"strings"
)

type Project struct {
	Name string
}

func CurrentProject() Project {
	dir, _ := os.Getwd()
	parts := strings.Split(dir, "/")
	name := parts[len(parts)-1]

	return Project{
		Name: strings.Title(name),
	}
}
