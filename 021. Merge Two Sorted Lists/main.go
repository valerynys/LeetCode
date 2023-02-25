package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Insert(val int) {
	li := &ListNode{
		Val:  val,
		Next: l.Next,
	}
	l.Next = li
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	for l1 != nil || l2 != nil {
		if l1 == nil {
			cur.Next = l2
			l2 = nil
			continue
		}
		if l2 == nil {
			cur.Next = l1
			l1 = nil
			continue
		}

		if l1.Val > l2.Val {
			cur.Next = l2
			cur, l2 = cur.Next, l2.Next
		} else {
			cur.Next = l1
			cur, l1 = cur.Next, l1.Next
		}
	}
	return res.Next
}

func main() {
	list1 := new(ListNode)
	list2 := new(ListNode)
	listres := new(ListNode)
	list1.Insert(12)
	list1.Insert(24)
	list1.Insert(5)
	list2.Insert(3)
	list2.Insert(8)
	list2.Insert(6)
	list2.Insert(25)
	listres = mergeTwoLists(list1, list2)
	fmt.Println(listres.Val)
}
