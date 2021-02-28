package tree

import (
	"fmt"
	"strconv"
	"strings"
)

// treeString is a representation of a tree as a string in the form of
// 		[<integer or nil>, <integer or nil>, ...]
type treeString string

// TODO: add a verification step in here to make sure array format is proper for parsing

// deserialize will return the tree string as a treeAsArray
func (t treeString) deserialize() treeAsArray {
	nodes := treeAsArray([]treeVal{})
	stringNodes := t[1 : len(t)-1] // strip the [] from the string representation
	nodeTokens := strings.Split(string(stringNodes), ",")

	for _, node := range nodeTokens {
		node = strings.Trim(node, " ")

		// value we're gonna append to the tree array
		val := treeVal(nil)

		// don't bother overwriting, write nil and move along
		if node == "nil" {
			nodes.push(val)
			continue
		}

		// node string we're looking at isn't nil, cast to an int
		v, err := strconv.Atoi(node)
		if err != nil {
			fmt.Printf("couldn't convert string %s to int\n", node)
			continue
		}

		// make val into a treeVal with concrete type as a pointer to the converted string node
		// ex:
		// 		"1" becomes a pointer to a value 1
		val = treeVal(&v)
		nodes.push(val)
	}

	return nodes
}

// BinaryTree represents a typical binary tree.
type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}

// InorderPrint will print the values of the tree in an inorder traversal.
func InorderPrint(node *BinaryTree) {
	fmt.Println("starting inorder traversal")
	inorderPrint(node)
	fmt.Printf("\n----------------")
}

func inorderPrint(node *BinaryTree) {
	if node == nil {
		return
	}
	inorderPrint(node.Left)
	fmt.Printf("%d ", node.Val)
	inorderPrint(node.Right)
}

// PostorderPrint will print the values of the tree in postorder traversal
func PostorderPrint(node *BinaryTree) {
	fmt.Println("starting postorder traversal")
	postorderPrint(node)
	fmt.Printf("\n---------------")
}

func postorderPrint(node *BinaryTree) {
	if node == nil {
		return
	}
	postorderPrint(node.Left)
	postorderPrint(node.Right)
	fmt.Printf("%d ", node.Val)
}

// treeVal is the value a node in a tree can hold. Is a pointer
// to accomodate passing in `nil` in the array to be turned into a tree.
//
// This also allows us to more easily convert the tree to hold a different type with
// minimal code changes.
type treeVal *int

// TreeAsArray represents a tree as a slice (array)
type treeAsArray []treeVal

// push appends a treeVal to the array.
func (t *treeAsArray) push(v treeVal) {
	*t = append(*t, v)
}

// pop returns the last element in the list and removes it.
func (t *treeAsArray) pop() treeVal {
	if len(*t) == 0 {
		fmt.Println("array is empty")
		return nil
	}
	tmp := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return tmp
}

// reverse will reverse the tree array
func (t *treeAsArray) reverse() treeAsArray {
	ret := make([]treeVal, len(*t))
	for i, j := 0, len(*t)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = (*t)[j], (*t)[i]
	}
	return ret
}

// ToTree converts the tree val array to a binary tree
func (t *treeAsArray) ToTree() *BinaryTree {
	children := t.reverse()
	root := children.pop()
	tree := &BinaryTree{
		Val: *root,
	}

	for _, node := range *t {
		if node != nil {
			if len(children) <= 0 {
				continue
			}
			v := children.pop()
			tree.Left = nil
			if v != nil {
				tree.Left = &BinaryTree{
					Val: *v,
				}
			}

			if len(children) <= 0 {
				continue
			}
			v = children.pop()
			tree.Right = nil
			if v != nil {
				tree.Right = &BinaryTree{
					Val: *v,
				}
			}

		}
	}

	return tree
}

// CreateBinaryTree creates a binary tree from a string-array representation
func CreateBinaryTree(tree string) *BinaryTree {
	if tree == "" {
		return nil
	}

	t := treeString(tree)
	arr := t.deserialize()
	return arr.ToTree()
}
