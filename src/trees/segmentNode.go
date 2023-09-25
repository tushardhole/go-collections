package trees

type SegmentNode struct {
	start int
	end   int
	min   int
	left  *SegmentNode
	right *SegmentNode
}

func NewSegmentNode(start int, end int, min int) *SegmentNode {
	return &SegmentNode{start: start, end: end, min: min}
}

func (node *SegmentNode) Min() *int {
	return &node.min
}

func (node *SegmentNode) Start() *int {
	return &node.start
}

func (node *SegmentNode) End() *int {
	return &node.end
}

func (node *SegmentNode) SetLeft(left *SegmentNode) {
	node.left = node
}

func (node *SegmentNode) SetRight(left *SegmentNode) {
	node.right = node
}

func (node *SegmentNode) Left() *SegmentNode {
	return node.left
}

func (node *SegmentNode) Right() *SegmentNode {
	return node.right
}
