package main

import "fmt"

type singlyNode struct {
	Data int
	Next *singlyNode
}

func main() {
	list1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	head := createList(list1)
	printList(head)
}

func createList(dataArray []int) *singlyNode {
	head := &singlyNode{}
	temp := head
	for _, i := range dataArray {
		node := &singlyNode{
			Data: i,
			Next: nil,
		}
		temp.Next = node
		temp = node
	}
	return head
}

func printList(head *singlyNode) {
	temp := head
	for temp.Next != nil {
		fmt.Print(temp.Data, "->")
		temp = temp.Next
	}
	fmt.Print(temp.Data)
	fmt.Println("->", temp.Next)
}
