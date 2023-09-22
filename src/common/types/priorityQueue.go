package collections

import . "go-collections/src/common/behaviours"

type PriorityQueue interface {
	Poll() Comparable
	Add(Comparable) error
	Peek() Comparable
	IsEmpty() bool
	Size() int
}
