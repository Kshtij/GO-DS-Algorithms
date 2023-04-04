package main

import (
	"fmt"
)

func main() {
	list1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	head := createList(list1)
	printList(head)
	head = reverseKGroup(head, 3)
	printList(head)
}

func createList(dataArray []int) *ListNode {
	head := &ListNode{}
	temp := head
	for _, i := range dataArray {
		node := &ListNode{
			Val:  i,
			Next: nil,
		}
		temp.Next = node
		temp = node
	}
	return head
}

func printList(head *ListNode) {
	temp := head
	for temp.Next != nil {
		fmt.Print(temp.Val, "->")
		temp = temp.Next
	}
	fmt.Print(temp.Val)
	fmt.Println("->", temp.Next)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var start, finish, temp, newHead, nextItrStart, prevFinish *ListNode
	start = head

	newHead = start

	for i := 1; i < k; i++ {
		newHead = newHead.Next
	}

	breakOut := false
	for start != nil {
		finish = start
		for i := 1; i < k; i++ {
			finish = finish.Next
			if finish == nil {
				breakOut = true
				break
			}
		}
		if breakOut {
			break
		}
		if prevFinish != nil {
			prevFinish.Next = finish
		}

		nextItrStart = finish.Next

		//reverse
		prevFinish = start
		for start != finish && start != nil {
			temp = start
			start = start.Next

			temp.Next = finish.Next
			finish.Next = temp
		}

		start = nextItrStart

	}

	return newHead
}
