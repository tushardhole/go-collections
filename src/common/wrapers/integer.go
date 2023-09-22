package collections

import (
	"go-collections/src/common/behaviours"
)

type ComparableInteger struct {
	value int
}

func (i ComparableInteger) CompareTo(other behaviours.Comparable) int {
	return i.value - other.(ComparableInteger).value
}

func (i ComparableInteger) Value() int {
	return i.value
}

func NewComparableInteger(value int) behaviours.Comparable {
	return ComparableInteger{value: value}
}
