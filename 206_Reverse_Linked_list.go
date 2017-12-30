package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	prev := (*ListNode)(nil)
	curr := head
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}
	return prev
}

func main() {
	note1 := new(ListNode)
	note2 := new(ListNode)
	note1.Val = 1
	note2.Val = 2
	note1.Next = note2
	result := reverseList(note1)
	for result != nil {
		fmt.Print(result.Val)
		fmt.Print(",")
		result = result.Next
	}
}
