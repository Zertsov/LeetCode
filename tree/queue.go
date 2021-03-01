package tree

type queue []*BinaryTree

func (q queue) len() int { return len(q) }
func (q *queue) enqueue(b *BinaryTree) {
	*q = append(*q, b)
}
func (q *queue) dequeue() *BinaryTree {
	if q.len() == 0 {
		return nil
	}
	tmp := (*q)[0]
	*q = (*q)[:q.len()-1]
	return tmp
}
