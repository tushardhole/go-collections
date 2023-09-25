package trees

type SegmentTree interface {
	minOfRange(start int, end int) (int, error)
}

type segmentTreeImpl struct {
	root *SegmentNode
}

func (st *segmentTreeImpl) build(array []int, start int, end int) *SegmentNode {
	current := NewSegmentNode(start, end, 0)

	if start == end {
		current.min = array[start]
		return current
	}

	mid := start + (end-start)/2
	current.left = st.build(array, start, mid)
	current.right = st.build(array, mid+1, end)
	current.min = min(current.left.min, current.right.min)

	return current
}

func (st *segmentTreeImpl) minOfRange(start int, end int) (int, error) {
	return minOfRangeRecursive(st.root, start, end)
}

func minOfRangeRecursive(root *SegmentNode, start int, end int) (int, error) {
	if root == nil {
		return 0, nil
	}

	rootStart := *root.Start()
	rootEnd := *root.End()

	if rootStart == start && rootEnd == end {
		return *root.Min(), nil
	}

	mid := root.start + (root.end-root.start)/2

	if end <= mid {
		return minOfRangeRecursive(root.left, start, end)
	} else if start > mid {
		return minOfRangeRecursive(root.right, start, end)
	} else {
		leftmin, _ := minOfRangeRecursive(root.left, start, mid)
		rightmin, _ := minOfRangeRecursive(root.right, mid+1, end)
		return min(leftmin, rightmin), nil
	}
}

func NewSegmentTree(array []int) SegmentTree {
	tree := segmentTreeImpl{}
	tree.root = tree.build(array, 0, len(array)-1)
	return &tree
}
