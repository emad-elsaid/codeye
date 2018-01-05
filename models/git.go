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

	logs, err := repo.Log(&git.LogOptions{})
	checkError(err)

	commitsCountPerAuthor := newCounter()
	err = logs.ForEach(func(c *object.Commit) error {
		commitsCountPerAuthor.add(c.Author.Name)
		return nil
	})

	return commitsCountPerAuthor.toA()
}
