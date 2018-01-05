package models

import "sort"

type sortableRow struct {
	key   int64
	value string
}
type sortable []sortableRow

func (this sortable) Len() int           { return len(this) }
func (this sortable) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }
func (this sortable) Less(i, j int) bool { return this[i].key > this[j].key }
func (this sortable) Sort()              { sort.Sort(this) }
