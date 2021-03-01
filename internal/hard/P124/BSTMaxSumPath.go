package p124

import "github.com/Zertsov/LeetCode/tree"

func maxPathSum(root *tree.BinaryTree) int {
	maxVal := (2 << 31) * -1 // arbitrary gigantic negative value to be overwritten

	// To allow for this first-class function to be called recursively, Go needs
	// us to declare the variable before initialization, so using the walrus
	// syntax (the := declarator) won't work because we haven't declared the
	// first class function at that time. So, we declare the variable, which
	// is initially set to `nil`, and then assign the function body itself, which
	// will be able to reference itself because the variable would have already been
	// declared.
	//
	// This could all be avoided by making it a non-first-class function, but this
	// is a good exercise and reference for this particular situation.

	var maxAddition func(*tree.BinaryTree) int // set to nil, declare before function is defined for recursive reference
	maxAddition = func(node *tree.BinaryTree) int {
		// recursion guard clause
		if node == nil {
			return 0
		}

		// go down left and right subtree, and see if a path gives us a value higher than 0, meaning
		// that they are additive to our total sum. Since we can only choose one path, we take whichever
		// path (left or right subtree(s)) give us the largest sum.
		left, right := max(maxAddition(node.Left), 0), max(maxAddition(node.Right), 0)
		cur := node.Val + left + right
		maxVal = max(maxVal, cur)

		return node.Val + max(left, right)
	}

	maxAddition(root) // recursion chain go brrrr
	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
