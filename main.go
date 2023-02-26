package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	var sentinel *ListNode = &ListNode{Next: head}
	for curr, currNext := sentinel, head; currNext != nil; {
		if currNext.Val == val {
			curr.Next, currNext = currNext.Next, currNext.Next
		} else {
			curr, currNext = curr.Next, currNext.Next
		}
	}
	return sentinel.Next
}