package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}
	dfs(root, 0, &ans)
	for i := 0; i < len(ans); i++ {
		if i%2 == 0 {
			continue
		}
		reverse(&ans[i])
	}
	return ans
}

func reverse(arr *[]int) {
	n := len(*arr)
	for i := 0; i < n/2; i++ {
		(*arr)[i], (*arr)[n-1-i] = (*arr)[n-1-i], (*arr)[i]
	}
}

func dfs(node *TreeNode, level int, ans *[][]int) {
	if level >= len(*ans) {
		*ans = append(*ans, []int{node.Val})
	} else {
		(*ans)[level] = append((*ans)[level], node.Val)
	}
	if node.Left != nil {
		dfs(node.Left, level+1, ans)
	}
	if node.Right != nil {
		dfs(node.Right, level+1, ans)
	}
}

func main() {

}
