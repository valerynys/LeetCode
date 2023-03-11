package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	bnode := &TreeNode{}
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		bnode.Val = nums[0]
		return bnode
	}
	mid := len(nums) / 2
	bnode.Val = nums[mid]
	bnode.Left = sortedArrayToBST(nums[:mid])
	bnode.Right = sortedArrayToBST(nums[mid+1:])
	return bnode

}

func sortedListToBST(head *ListNode) *TreeNode {
	numbers := []int{}

	for head != nil {
		numbers = append(numbers, head.Val)
		head = head.Next
	}
	return sortedArrayToBST(numbers)
}
