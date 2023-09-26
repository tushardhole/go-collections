package internal

func MinOfRangeRecursive(root *SegmentNode, start int, end int) (int, error) {
	if root == nil {
		return 0, nil
	}

	rootStart := *root.Start()
	rootEnd := *root.End()

	if rootStart == start && rootEnd == end {
		return *root.Min(), nil
	}

	mid := *root.Start() + (*root.End()-*root.Start())/2

	if end <= mid {
		return MinOfRangeRecursive(root.Left(), start, end)
	} else if start > mid {
		return MinOfRangeRecursive(root.Right(), start, end)
	} else {
		leftmin, _ := MinOfRangeRecursive(root.Left(), start, mid)
		rightmin, _ := MinOfRangeRecursive(root.Right(), mid+1, end)
		return min(leftmin, rightmin), nil
	}
}
