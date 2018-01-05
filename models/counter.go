package models

type counter struct {
	count map[interface{}]int64
}

func newCounter() *counter {
	return &counter{
		count: make(map[interface{}]int64),
	}
}

func (this *counter) add(key interface{}) {
	v, ok := this.count[key]
	if !ok {
		this.count[key] = 1
	} else {
		this.count[key] = v + 1
	}
}
