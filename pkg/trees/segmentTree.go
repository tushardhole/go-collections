package trees

import (
	. "github.com/tushardhole/go-collections/internal"
)

type SegmentTree interface {
	MinOfRange(start int, end int) (int, error)
	MaxOfRange(start int, end int) (int, error)
	SumOfRange(start int, end int) (int, error)
}

type segmentTreeImpl struct {
	root *SegmentNode
}

func (st *segmentTreeImpl) build(array []int, start int, end int) *SegmentNode {
	current := NewSegmentNode(start, end, 0)

	if start == end {
		current.SetMin(array[start])
		current.SetMax(array[start])
		current.SetSum(array[start])
		return current
	}

	mid := start + (end-start)/2
	current.SetLeft(st.build(array, start, mid))
	current.SetRight(st.build(array, mid+1, end))
	current.SetMin(min(*current.Left().Min(), *current.Right().Min()))
	current.SetMax(max(*current.Left().Max(), *current.Right().Max()))
	current.SetSum(*current.Left().Sum() + *current.Right().Sum())

	return current
}

func (st *segmentTreeImpl) MinOfRange(start int, end int) (int, error) {
	return MinOfRangeRecursive(st.root, start, end)
}

func (st *segmentTreeImpl) MaxOfRange(start int, end int) (int, error) {
	return MinOfRangeRecursive(st.root, start, end)
}

func (st *segmentTreeImpl) SumOfRange(start int, end int) (int, error) {
	return MinOfRangeRecursive(st.root, start, end)
}

func NewSegmentTree(array []int) SegmentTree {
	tree := segmentTreeImpl{}
	tree.root = tree.build(array, 0, len(array)-1)
	return &tree
}
