package collections

import (
	. "github.com/tushardhole/go-collections/src/common/behaviours"
)

type ComparableFloat struct {
	value float64
}

func (i ComparableFloat) CompareTo(other Comparable) int {
	diff := i.value - other.(ComparableFloat).value
	if diff > 0 {
		return 1
	}
	if diff < 0 {
		return -1
	}
	return 0
}

func (i ComparableFloat) Value() float64 {
	return i.value
}

func NewComparableFloat(value float64) Comparable {
	return ComparableFloat{value: value}
}
