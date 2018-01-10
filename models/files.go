package models

import (
	"os"
	"sort"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
	object "gopkg.in/src-d/go-git.v4/plumbing/object"
)

type File struct {
	object.File
}

type Files []File

func (this *File) Basename() string {
	parts := strings.Split(this.Name, "/")
	return parts[len(parts)-1]
}

func (this *File) Ext() string {
	basename := this.Basename()

	dot := strings.LastIndex(basename, ".")
	if dot > 0 && dot < len(basename)-1 {
		return basename[dot+1:]
	}

	return ""
}

func FilesList() Files {
	cwd, err := os.Getwd()
	checkError(err)

	repo, err := git.PlainOpen(cwd)
	checkError(err)

	ref, err := repo.Head()
	checkError(err)

	commit, err := repo.CommitObject(ref.Hash())
	checkError(err)

	tree, err := commit.Tree()
	checkError(err)

	files := Files{}
	tree.Files().ForEach(func(f *object.File) error {
		files = append(files, File{*f})
		return nil
	})

	return files
}

type ExtGroup struct {
	Ext   string
	Files Files
}

type ExtGroups []ExtGroup

func (this ExtGroups) Len() int           { return len(this) }
func (this ExtGroups) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }
func (this ExtGroups) Less(i, j int) bool { return len(this[i].Files) > len(this[j].Files) }
func (this ExtGroups) Sort()              { sort.Sort(this) }

func (this Files) ExtGroups() ExtGroups {
	extmap := make(map[string]ExtGroup)
	for _, f := range this {
		ext := f.Ext()
		if ext == "" {
			continue
		}

		stat, ok := extmap[ext]
		if !ok {
			stat = ExtGroup{
				Ext:   ext,
				Files: Files{},
			}
		}
		stat.Files = append(stat.Files, f)
		extmap[ext] = stat
	}

	groups := ExtGroups{}
	for _, v := range extmap {
		groups = append(groups, v)
	}
	groups.Sort()

	return groups
}
