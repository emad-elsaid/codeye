package models

import "sort"

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

func (this *Contributors) TotalCommits() (i int) {
	for _, c := range *this {
		i += c.commits
	}
	return
}

func (this Contributors) Len() int           { return len(this) }
func (this Contributors) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }
func (this Contributors) Less(i, j int) bool { return this[i].commits > this[j].commits }
func (this Contributors) Sort()              { sort.Sort(this) }
