package models

import (
	"os"
	"sort"

	git "gopkg.in/src-d/go-git.v4"
	object "gopkg.in/src-d/go-git.v4/plumbing/object"
)

type Contributor struct {
	name    string
	commits int
}

func (this *Contributor) Name() string {
	return this.name
}

func (this *Contributor) Commits() int {
	return this.commits
}

type Contributors []Contributor

func (this Contributors) TotalCommits() (i int) {
	for _, c := range this {
		i += c.commits
	}
	return
}

func (this Contributors) Len() int           { return len(this) }
func (this Contributors) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }
func (this Contributors) Less(i, j int) bool { return this[i].commits > this[j].commits }
func (this Contributors) Sort()              { sort.Sort(this) }

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
