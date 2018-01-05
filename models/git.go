package models

import (
	"os"

	git "gopkg.in/src-d/go-git.v4"
	object "gopkg.in/src-d/go-git.v4/plumbing/object"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func Contributors() [][]string {
	cwd, err := os.Getwd()
	checkError(err)

	repo, err := git.PlainOpen(cwd)
	checkError(err)

	head, err := repo.Head()
	checkError(err)

	logs, err := repo.Log(&git.LogOptions{From: head.Hash()})
	checkError(err)

	output := make([][]string, 0, 100)
	err = logs.ForEach(func(c *object.Commit) error {
		output = append(output, []string{c.Author.Name, c.Message})
		return nil
	})

	return output
}
