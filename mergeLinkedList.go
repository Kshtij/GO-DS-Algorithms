package main

import "fmt"

type singlyNode struct {
	Data int
	Next *singlyNode
}

func main() {
	list1 := []int{1, 5, 6, 7, 8, 15, 16, 17}
	list2 := []int{2, 3, 4, 9, 10, 11, 12, 13, 14, 18, 19, 20}

	head1 := createList(list1)
	head2 := createList(list2)
	printList(head1)
	fmt.Println()
	printList(head2)
	fmt.Println()

	head3 := mergeLinkedList(head1, head2)
	printList(head3)

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
	fmt.Println(temp.Next)
}

func mergeLinkedList(head1, head2 *singlyNode) *singlyNode {
	var temp1, temp2 *singlyNode
	temp1 = head1.Next
	temp2 = head2.Next
	mergedListPtr := head1
	for temp1 != nil && temp2 != nil {
		if temp1.Data <= temp2.Data {
			mergedListPtr.Next = temp1
			temp1 = temp1.Next
		} else {
			mergedListPtr.Next = temp2
			temp2 = temp2.Next
		}
		mergedListPtr = mergedListPtr.Next

	}
	if temp1 == nil {
		mergedListPtr.Next = temp2
	} else {
		mergedListPtr.Next = temp1
	}

	return head1
}
