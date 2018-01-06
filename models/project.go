package models

import (
	"os"
	"strings"
)

type Project struct {
	name string
}

func (this Project) Name() string {
	return this.name
}

func CurrentProject() Project {
	dir, _ := os.Getwd()
	parts := strings.Split(dir, "/")
	name := parts[len(parts)-1]

	return Project{
		name: name,
	}
}
