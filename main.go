package main

import (
	"fmt"

	"github.com/Zertsov/LeetCode/tree"
)

func main() {
	tests := map[string]struct {
		input string
		// add verification to this struct if you want to ensure output matches something
	}{
		"tree without nils": {
			"[1,2,3]",
		},
		"tree with nils": {"[1,2,3,nil,nil,8]"},
	}

	for name, test := range tests {
		fmt.Printf("running test \"%s\"\n", name)
		t := tree.CreateBinaryTree(test.input)
		tree.InorderPrint(t)
		fmt.Println()
		tree.PostorderPrint(t)
		fmt.Println()
		fmt.Printf("finished running test %s\n", name)
	}
}
