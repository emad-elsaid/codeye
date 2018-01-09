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

	iter, err := repo.Log(&git.LogOptions{})
	checkError(err)

	commit, err := iter.Next()
	checkError(err)

	first := firstParent(commit)

	return &Commit{*first}
}

func firstParent(commit *object.Commit) *object.Commit {
	parents := commit.Parents()
	parent_commit, err := parents.Next()
	if err != nil {
		return commit
	}
	return firstParent(parent_commit)
}
