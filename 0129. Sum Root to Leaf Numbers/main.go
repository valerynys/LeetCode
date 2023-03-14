package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}

	// calculate the sum of the current path
	sum = sum*10 + node.Val

	// if it's a leaf node, return the sum
	if node.Left == nil && node.Right == nil {
		return sum
	}

	// recursively calculate the sum of left and right subtrees
	leftSum := dfs(node.Left, sum)
	rightSum := dfs(node.Right, sum)

	// return the sum of both subtrees
	return leftSum + rightSum
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	// calculate the sum of all root-to-leaf paths
	totalSum := sumNumbers(root)

	// print the result
	fmt.Println(totalSum) // Output: 262
}
