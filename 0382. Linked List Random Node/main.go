package main

import (
	"math/rand"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	rand.Seed(time.Now().UnixNano())
	return Solution{head}
}

func (this *Solution) GetRandom() int {
	result := this.head.Val
	node := this.head.Next
	i := 2

	for node != nil {
		if rand.Intn(i) == 0 {
			result = node.Val
		}

		node = node.Next
		i++
	}

	return result
}
