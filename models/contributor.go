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
	Commits int
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
		i += c.Commits
	}
	return
}

func (this Contributors) Len() int           { return len(this) }
func (this Contributors) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }
func (this Contributors) Less(i, j int) bool { return this[i].Commits > this[j].Commits }
func (this Contributors) Sort()              { sort.Sort(this) }

func NewContributors() Contributors {
	cwd, err := os.Getwd()
	checkError(err)

	repo, err := git.PlainOpen(cwd)
	checkError(err)

	logs, err := repo.Log(&git.LogOptions{})
	checkError(err)
	defer logs.Close()

	by_email := make(map[string]*Contributor)

	err = logs.ForEach(func(c *object.Commit) error {
		contributor, ok := by_email[c.Author.Email]

		if !ok {
			contributor = &Contributor{c.Author, 1}
			by_email[c.Author.Email] = contributor
		} else {
			contributor.Commits++
		}

		return nil
	})

	by_name := make(map[string]*Contributor, len(by_email))
	for _, v := range by_email {
		contributor, ok := by_name[v.Name]

		if ok {
			v.Commits += contributor.Commits
		}

		by_name[v.Name] = v
	}

	values := make(Contributors, 0, len(by_name))
	for _, c := range by_name {
		values = append(values, *c)
	}
	values.Sort()

	return values
}
