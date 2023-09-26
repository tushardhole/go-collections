package internal

type SegmentNode struct {
	start int
	end   int
	min   int
	max   int
	sum   int
	left  *SegmentNode
	right *SegmentNode
}

func NewSegmentNode(start int, end int, min int) *SegmentNode {
	return &SegmentNode{start: start, end: end, min: min}
}

func (node *SegmentNode) Min() *int {
	return &node.min
}

func (node *SegmentNode) Max() *int {
	return &node.max
}

func (node *SegmentNode) Sum() *int {
	return &node.sum
}

func (node *SegmentNode) Start() *int {
	return &node.start
}

func (node *SegmentNode) End() *int {
	return &node.end
}

func (node *SegmentNode) SetLeft(left *SegmentNode) {
	node.left = left
}

func (node *SegmentNode) SetRight(right *SegmentNode) {
	node.right = right
}

func (node *SegmentNode) SetMin(min int) {
	node.min = min
}

func (node *SegmentNode) SetMax(max int) {
	node.max = max
}

func (node *SegmentNode) SetSum(sum int) {
	node.sum = sum
}

func (node *SegmentNode) Left() *SegmentNode {
	return node.left
}

func (node *SegmentNode) Right() *SegmentNode {
	return node.right
}
