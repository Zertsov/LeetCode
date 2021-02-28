package p101

import "github.com/Zertsov/LeetCode/tree"

func isSymmetric(root *tree.BinaryTree) bool {
	return mirrored(root, root)
}

func mirrored(node1, node2 *tree.BinaryTree) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}

	return node1.Val == node2.Val && mirrored(node1.Left, node2.Right) && mirrored(node2.Left, node1.Right)
}
