package collections

import . "github.com/tushardhole/go-collections/src/common/behaviours"

type PriorityQueue interface {
	Poll() Comparable
	Add(Comparable) error
	Peek() Comparable
	IsEmpty() bool
	Size() int
}
