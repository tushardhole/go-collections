package internal

import "errors"

func UpdateRecursive(root *SegmentNode, index int, value int) (bool, error) {
	if index > *root.End() || index < *root.Start() {
		return false, errors.New("Update element Index out of bounds")
	}

	if *root.Start() == index && *root.End() == index {
		root.SetMin(value)
		root.SetMax(value)
		root.SetSum(value)
		return true, nil
	}

	mid := *root.Start() + (*root.End()-*root.Start())/2

	if index <= mid {
		leftUpdted, error := UpdateRecursive(root.Left(), index, value)
		if error != nil {
			return leftUpdted, error
		}
	}

	if index > mid {
		rightUpdated, error := UpdateRecursive(root.Right(), index, value)
		if error != nil {
			return rightUpdated, error
		}
	}

	root.SetMin(min(*root.Left().Min(), *root.Right().Min()))
	root.SetMax(max(*root.Left().Min(), *root.Right().Max()))
	root.SetSum(*root.Left().Sum() + *root.Right().Sum())

	return true, nil
}
