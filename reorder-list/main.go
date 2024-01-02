package main

import "log"

func foo(param int) int {
	return param
}

func main() {
	hd := CreateList([]int{1, 2, 3, 4, 5, 6, 7})
	reorderList(hd)
}

type node struct {
	data int
	next *node
}

func CreateList(arr []int) *node {
	dummyNode := node{}
	temp := &dummyNode
	for i := 0; i < len(arr); i++ {
		newNode := node{
			data: arr[i],
			next: nil,
		}
		temp.next = &newNode
		temp = temp.next
	}

	return dummyNode.next

}

func reorderList(hd *node) {

	log.Println("entire List")
	printList(hd)
	currentPointer := hd
	fastPointer := hd

	for fastPointer != nil && currentPointer != nil {
		if fastPointer.next != nil {
			fastPointer = fastPointer.next.next
			currentPointer = currentPointer.next
		} else {
			fastPointer = fastPointer.next
		}
	}

	//reversing start
	secondList := currentPointer.next

	log.Println("before reversing ")
	printList(secondList)

	refNode := secondList

	for refNode != nil {
		nextNode := refNode.next
		subsequentNode := refNode.next.next

		refNode.next = subsequentNode
		refNode = refNode.next

		nextNode.next = secondList
		secondList = nextNode

	}

	//reversing end
	log.Println("after reversing")
	printList(secondList)

}

func printList(nd *node) {
	log.Println("----------")
	for nd != nil {
		log.Println(nd.data)
		nd = nd.next
	}
	log.Println("----------")

}
