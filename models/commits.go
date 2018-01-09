package models

import (
	"os"

	git "gopkg.in/src-d/go-git.v4"
	object "gopkg.in/src-d/go-git.v4/plumbing/object"
)

type Commit struct {
	object.Commit
}

func (this *Commit) Author() *Contributor {
	return &Contributor{Signature: this.Commit.Author}
}

func FirstCommit() *Commit {
	cwd, err := os.Getwd()
	checkError(err)

	repo, err := git.PlainOpen(cwd)
	checkError(err)

	commits, err := repo.CommitObjects()
	checkError(err)
	var last_c *object.Commit

	commits.ForEach(func(c *object.Commit) error {
		if last_c == nil || last_c.Committer.When.Unix() > c.Committer.When.Unix() {
			last_c = c
		}

		return nil
	})

	return &Commit{*last_c}
}
