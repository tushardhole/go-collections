package internal

func SumOfRangeRecursive(root *SegmentNode, start int, end int) (int, error) {
	if root == nil {
		return 0, nil
	}

	rootStart := *root.Start()
	rootEnd := *root.End()

	if rootStart == start && rootEnd == end {
		return *root.Sum(), nil
	}

	mid := *root.Start() + (*root.End()-*root.Start())/2

	if end <= mid {
		return SumOfRangeRecursive(root.Left(), start, end)
	} else if start > mid {
		return SumOfRangeRecursive(root.Right(), start, end)
	} else {
		leftSum, _ := SumOfRangeRecursive(root.Left(), start, mid)
		rightSum, _ := SumOfRangeRecursive(root.Right(), mid+1, end)
		return leftSum + rightSum, nil
	}
}
