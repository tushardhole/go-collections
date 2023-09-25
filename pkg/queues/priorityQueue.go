package queues

import . "github.com/tushardhole/go-collections/pkg/behaviours"

type PriorityQueue interface {
	Poll() Comparable
	Add(Comparable) error
	Peek() Comparable
	IsEmpty() bool
	Size() int
}
