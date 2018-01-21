package codeye

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/emad-elsaid/codeye/models"
)

func FileStructure(w http.ResponseWriter, r *http.Request) {
	// files := models.FilesList()
	// pathsList := make([]string, len(files))

	// for i, f := range files {
	// 	pathsList[i] = f.Name
	// }

	// paths := strings.Join(pathsList, "\n")
	// pathsReader := strings.NewReader(paths)

	// dirtree := exec.Command("/usr/bin/env", "sh", "-c", "dirtree", "--template=flame")
	// var stdout bytes.Buffer
	// var stderr bytes.Buffer
	// dirtree.Stdout = &stdout
	// dirtree.Stderr = &stderr
	// dirtree.Stdin = pathsReader
	// err := dirtree.Run()

	// if err != nil {
	// 	log.Printf("Error while running dirtree: %s", err)
	// }

	// page := Page{
	// 	Title:    "File Structure",
	// 	Template: "file_structure",
	// 	Navbar: Navbar{
	// 		Name: models.CurrentProject().Name,
	// 	},
	// 	Data: struct {
	// 		Output template.HTML
	// 		Error  string
	// 	}{
	// 		template.HTML(stdout.String()),
	// 		stderr.String(),
	// 	},
	// }

	page := Page{
		Title:    "File Structure",
		Template: "file_structure",
		Navbar: Navbar{
			Name: models.CurrentProject().Name,
		},
		Data: "/flame_graph",
	}
	page.Render(w)
}

func FlameGraph(w http.ResponseWriter, r *http.Request) {
	files := models.FilesList()
	pathsList := make([]string, len(files))

	for i, f := range files {
		pathsList[i] = f.Name
	}

	paths := strings.Join(pathsList, "\n")
	pathsReader := strings.NewReader(paths)

	dirtree := exec.Command("/usr/bin/env", "sh", "-c", "dirtree -t flame")
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	dirtree.Stdout = &stdout
	dirtree.Stderr = &stderr
	dirtree.Stdin = pathsReader
	err := dirtree.Run()

	if err != nil {
		log.Printf("Error while running dirtree: %s", err)
	}

	fmt.Fprint(w, dirtree.Stdout)
}
