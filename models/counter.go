package models

import "strconv"

type counter struct {
	count map[string]int64
}

func newCounter() *counter {
	return &counter{
		count: make(map[string]int64, 100),
	}
}

func (this *counter) add(key string) {
	v, ok := this.count[key]
	if !ok {
		this.count[key] = 1
	} else {
		this.count[key] = v + 1
	}
}

func (this *counter) toA() [][]string {
	var s sortable
	for k, v := range this.count {
		s = append(s, sortableRow{key: v, value: k})
	}
	s.Sort()

	output := make([][]string, 0, 100)
	for _, row := range s {
		output = append(output, []string{row.value, strconv.FormatInt(row.key, 10)})
	}
	return output
}
