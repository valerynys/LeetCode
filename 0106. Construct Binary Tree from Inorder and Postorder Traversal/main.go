package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	// find the root node from postorder traversal
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}

	// find the index of root node in inorder traversal
	var idx int
	for i, v := range inorder {
		if v == rootVal {
			idx = i
			break
		}
	}

	// recursively build left and right subtrees
	root.Left = buildTree(inorder[:idx], postorder[:idx])
	root.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])

	return root
}

func main() {

}
