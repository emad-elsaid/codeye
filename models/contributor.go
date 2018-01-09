package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"sort"

	git "gopkg.in/src-d/go-git.v4"
	object "gopkg.in/src-d/go-git.v4/plumbing/object"
)

type Contributor struct {
	object.Signature
	commits int
}

func (this *Contributor) Commits() int {
	return this.commits
}

func (this *Contributor) Image() string {
	hash := md5.New()
	hash.Write([]byte(this.Email))
	segment := hex.EncodeToString(hash.Sum(nil))
	return fmt.Sprintf("http://www.gravatar.com/avatar/%s", segment)
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
	defer logs.Close()

	contributors := make(map[string]*Contributor)

	err = logs.ForEach(func(c *object.Commit) error {
		contributor, ok := contributors[c.Author.Name]

		if !ok {
			contributor = &Contributor{c.Author, 1}
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
