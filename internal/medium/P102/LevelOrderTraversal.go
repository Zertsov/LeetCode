package p102

import "github.com/Zertsov/LeetCode/tree"

type queue []*tree.BinaryTree

func (q *queue) dequeue() *tree.BinaryTree {
	tmp := (*q)[0]
	*q = (*q)[1:]
	return tmp
}
func (q *queue) enqueue(t *tree.BinaryTree) {
	*q = append(*q, t)
}
func (q queue) Len() int {
	return len(q)
}

func levelOrder(root *tree.BinaryTree) [][]int {
	if root == nil {
		return nil
	}

	answer := [][]int{} // this is what we'll return

	// create a processing queue (pq) with the root as our starting point
	pq := queue([]*tree.BinaryTree{root})

	// while we have nodes in the processing queue
	for pq.Len() > 0 {
		/*
		 Idea here is that we will dequeue every node in the pq, because
		 doing that will give us all the nodes at a given level. Once
		 they've all been dequeue'd, we iterate through each node that
		 was dequeue'd and add it's left child, then the right (if they exist)
		 and add the node value itself to an array that we'll append to the
		 answer matrix.
		*/
		nodes := make([]*tree.BinaryTree, len(pq))

		// declare i before loop because .dequeue will change the length of
		// the pq
		i := 0
		for pq.Len() > 0 {
			nodes[i] = pq.dequeue()
			i++
		}

		// vals is what we'll append to the answer matrix
		vals := make([]int, len(nodes))
		for i, node := range nodes {
			if node.Left != nil {
				pq.enqueue(node.Left)
			}
			if node.Right != nil {
				pq.enqueue(node.Right)
			}
			vals[i] = node.Val
		}
		answer = append(answer, vals)
	}

	return answer
}
