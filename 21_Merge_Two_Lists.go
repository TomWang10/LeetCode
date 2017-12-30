package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	resultList := new(ListNode)
	resultList.Next = nil
	resultStart := resultList
	for {
		if l1 == nil {
			resultList.Next = l2
			break
		}
		if l2 == nil {
			resultList.Next = l1
			break
		}
		var tmp *ListNode = nil
		if l1.Val < l2.Val {
			tmp = l1
			l1 = l1.Next
		} else {
			tmp = l2
			l2 = l2.Next
		}
		resultList.Next = tmp
		resultList = resultList.Next
	}
	return resultStart.Next
}

func main() {
	List1Note1 := new(ListNode)
	List1Note1.Val = 1
	List1Note2 := new(ListNode)
	List1Note2.Val = 2
	List1Note3 := new(ListNode)
	List1Note3.Val = 4
	List2Note1 := new(ListNode)
	List2Note1.Val = 1
	List2Note2 := new(ListNode)
	List2Note2.Val = 3
	List2Note3 := new(ListNode)
	List2Note3.Val = 4

	List1Note1.Next = List1Note2
	List1Note2.Next = List1Note3

	List2Note1.Next = List2Note2
	List2Note2.Next = List2Note3

	var result *ListNode = nil
	result = mergeTwoLists(List1Note1, List2Note1)
	for result != nil {
		fmt.Print(result.Val)
		fmt.Println(",")
		result = result.Next
	}
}
