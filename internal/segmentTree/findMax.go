package internal

func MaxOfRangeRecursive(root *SegmentNode, start int, end int) (int, error) {
	if root == nil {
		return 0, nil
	}

	rootStart := *root.Start()
	rootEnd := *root.End()

	if rootStart == start && rootEnd == end {
		return *root.Max(), nil
	}

	mid := *root.Start() + (*root.End()-*root.Start())/2

	if end <= mid {
		return MaxOfRangeRecursive(root.Left(), start, end)
	} else if start > mid {
		return MaxOfRangeRecursive(root.Right(), start, end)
	} else {
		leftmin, _ := MaxOfRangeRecursive(root.Left(), start, mid)
		rightmin, _ := MaxOfRangeRecursive(root.Right(), mid+1, end)
		return max(leftmin, rightmin), nil
	}
}
