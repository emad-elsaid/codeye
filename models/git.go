package models

import (
	"os"

	git "gopkg.in/src-d/go-git.v4"
	object "gopkg.in/src-d/go-git.v4/plumbing/object"
)

func NewContributors() Contributors {
	cwd, err := os.Getwd()
	checkError(err)

	repo, err := git.PlainOpen(cwd)
	checkError(err)

	logs, err := repo.Log(&git.LogOptions{})
	checkError(err)

	contributors := make(map[string]*Contributor)

	err = logs.ForEach(func(c *object.Commit) error {
		contributor, ok := contributors[c.Author.Name]

		if !ok {
			contributor = &Contributor{c.Author.Name, 1}
			contributors[c.Author.Name] = contributor
		} else {
			contributor.commits++
		}

		return nil
	})

	values := make(Contributors, 0, len(contributors))
	for _, c := range contributors {
		values = append(values, *c)
	}
	values.Sort()

	return values
}
