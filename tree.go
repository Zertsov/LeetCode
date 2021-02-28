package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// TreeString is a representation of a tree as a string in the form of
// 		[<integer or nil>, <integer or nil>, ...]
type TreeString string

// Verify will determine if the TreeString is a valid array containing only
// int and nil types.
//
// Ex:
// 		["1","2","nil","4"] -> true
//		["1","2","blah","5"] -> false
// TODO: Doesn't work. Regex is wrong
func (t TreeString) Verify() bool {
	matcher := regexp.MustCompile(`^\[([0-9][nil]|.+(,[0-9])+)\]$`)
	return matcher.MatchString(string(t))
}

// Deserialize will return the tree string as a TreeAsArray
func (t TreeString) Deserialize() TreeAsArray {
	nodes := TreeAsArray([]TreeVal{})
	stringNodes := t[1 : len(t)-2] // strip the [] from the string representation
	nodeTokens := strings.Split(string(stringNodes), ",")

	for _, node := range nodeTokens {
		node = strings.Trim(node, " ")

		// value we're gonna append to the tree array
		val := TreeVal(nil)

		// don't bother overwriting, write nil and move along
		if node == "nil" {
			nodes = append(nodes, val)
			continue
		}

		// node string we're looking at isn't nil, cast to an int
		v, err := strconv.Atoi(node)
		if err != nil {
			fmt.Printf("couldn't convert string %s to int\n", node)
			continue
		}

		// make val into a TreeVal with concrete type as a pointer to the converted string node
		// ex:
		// 		"1" becomes a pointer to a value 1
		val = TreeVal(&v)
		nodes = append(nodes, val)
	}

	return nodes
}

// BinaryTree represents a typical binary tree.
type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}

// TreeVal is the value a node in a tree can hold. Is a pointer
// to accomodate passing in `nil` in the array to be turned into a tree.
type TreeVal *int

// TreeAsArray represents a tree as a slice (array)
type TreeAsArray []TreeVal

// Push appends a TreeVal to the array.
func (t *TreeAsArray) Push(v TreeVal) {
	*t = append(*t, v)
}

// Pop returns the last element in the list and removes it.
func (t *TreeAsArray) Pop() TreeVal {
	tmp := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return tmp
}

// Reverse will reverse the tree array
func (t *TreeAsArray) Reverse() TreeAsArray {
	ret := make([]TreeVal, len(*t))
	for i, j := 0, len(*t)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = (*t)[j], (*t)[i]
	}
	return ret
}

// ToTree converts the tree val array to a binary tree
func (t *TreeAsArray) ToTree() *BinaryTree {
	children := t.Reverse()
	root := children.Pop()
	tree := &BinaryTree{
		Val: *root,
	}

	for _, node := range *t {
		if node != nil {
			if len(children) > 0 {
				tree.Left = &BinaryTree{
					Val: *children.Pop(),
				}
			}
			if len(children) > 0 {
				tree.Right = &BinaryTree{
					Val: *children.Pop(),
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

	t := TreeString(tree)
	arr := t.Deserialize()
	return arr.ToTree()
}
