package p108

import "github.com/Zertsov/LeetCode/tree"

func sortedArrayToBST(nums []int) *tree.BinaryTree {
	// recursive guard clause
	if len(nums) == 0 {
		return nil
	}

	// pick a pivot point and create a node at said pivot
	pivot := len(nums) / 2
	root := &tree.BinaryTree{Val: nums[pivot]}

	// recursively create left and right children by slicing
	// nums value around the pivot
	root.Left = sortedArrayToBST(nums[:pivot])
	root.Right = sortedArrayToBST(nums[pivot+1:])
	return root
}
