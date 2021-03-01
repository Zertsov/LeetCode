package tree

// Inorder will return the nodes in their in-order traversal.
//
// Inorder traversals work as follows:
// 		1. Visit left subtree
//		2. Visit node itself
//		3. Visit right subtree
func Inorder(root *BinaryTree) []int {
	v := []int{}
	inorder(root, &v)
	return v
}

func inorder(root *BinaryTree, vals *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, vals)
	*vals = append(*vals, root.Val)
	inorder(root.Right, vals)
}

// Preorder will return the values of the tree in their preorder traversal.
//
// Preorder traversals work as follows:
//		1. Visit node itself
//		2. Visit left subtree
//		3. Visit right subtree
func Preorder(root *BinaryTree) []int {
	v := []int{}
	preorder(root, &v)
	return v
}

func preorder(root *BinaryTree, vals *[]int) {
	if root == nil {
		return
	}
	*vals = append(*vals, root.Val)
	preorder(root.Left, vals)
	preorder(root.Right, vals)
}

// Postorder will return the tree values in their postorder traversal.
//
// Postorder traversal works as follows:
// 		1. Visit left subtree
// 		2. Visit right subtree
// 		3. Visit node itself
func Postorder(root *BinaryTree) []int {
	v := []int{}
	postorder(root, &v)
	return v
}

func postorder(root *BinaryTree, vals *[]int) {
	if root == nil {
		return
	}
	postorder(root.Left, vals)
	postorder(root.Right, vals)
	*vals = append(*vals, root.Val)
}

// LevelOrder will return the tree values in their postorder traversal.
//
// LevelOrder traversal works as follows:
// 		1. Visit each node, from left to right, that are on the same level (equivalent depth)
func LevelOrder(root *BinaryTree) [][]int {
	v := [][]int{}
	levelOrder(root, &v)
	return v
}

func levelOrder(root *BinaryTree, vals *[][]int) {
	if root == nil {
		return
	}

	/*
		The idea here is to add every node to a processing queue to be processed.

		We'll get the number of nodes in the queue, which will be the total
		number of nodes on the level, and create an int array of that size to
		be appended to `vals`.

		Each time we process a node, we'll add it's left
		and right child to the queue if they're not nil. We'll also add the
		value of the node itself to the nodeVals array, and after we've gone
		through the number of nodes that are on the level, we exit the loop,
		append the values to our vals matrix, and start a new level process.
	*/

	// create a processing queue (pq)
	pq := queue([]*BinaryTree{root})

	// while we still have something to process
	for pq.len() > 0 {

		// create another array to hold just the values of the nodes
		nodesOnLevel := pq.len()
		nodeVals := make([]int, nodesOnLevel)

		// Only add the number of nodes that're on the level to vals.
		for i := 0; i < nodesOnLevel; i++ {
			node := pq.dequeue()
			if node.Left != nil {
				pq.enqueue(node.Left)
			}
			if node.Right != nil {
				pq.enqueue(node.Right)
			}
			nodeVals[i] = node.Val
		}

		*vals = append(*vals, nodeVals)
	}
}
