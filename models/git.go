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

func Contributors() [][]interface{} {
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

	var s sortable
	for k, v := range commitsCountPerAuthor.count {
		s = append(s, sortableRow{key: v, value: k})
	}
	s.Sort()

	output := make([][]interface{}, 0, len(commitsCountPerAuthor.count))
	for _, row := range s {
		output = append(output, []interface{}{row.value.(string), row.key})
	}
	return output
}
