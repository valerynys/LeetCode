package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var nodes []*TreeNode
	nodes = append(nodes, root)

	var i int
	for nodes[i] != nil {
		nodes = append(nodes, nodes[i].Left, nodes[i].Right)
		i++
	}

	for j := i; j < len(nodes); j++ {
		if nodes[j] != nil {
			return false
		}
	}

	return true
}

func main() {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, &TreeNode{6, nil, nil}, nil}}
	fmt.Println(isCompleteTree(root)) // true

	root2 := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}
	fmt.Println(isCompleteTree(root2)) // true
}
