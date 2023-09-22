package collections

import (
	. "github.com/tushardhole/go-collections/src/common/behaviours"
)

type ComparableInteger struct {
	value int
}

func (i ComparableInteger) CompareTo(other Comparable) int {
	return i.value - other.(ComparableInteger).value
}

func (i ComparableInteger) Value() int {
	return i.value
}

func NewComparableInteger(value int) Comparable {
	return ComparableInteger{value: value}
}
