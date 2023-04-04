package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {

	if len(lists) == 0 {
		return nil
	}

	var retListHdNode, temp *ListNode
	//init
	heap := Heap{
		arr: make([]*ListNode, 0, 1),
	}

	for i := 0; i < len(lists) && lists[i] != nil; i++ {
		heap.Push(lists[i])
	}

	//temp := retListHdNode
	node := heap.Pop()
	if node != nil {
		retListHdNode = &ListNode{
			Val: node.Val,
		}
		if node.Next != nil {
			heap.Push(node.Next)
		}
	}
	temp = retListHdNode
	node = heap.Pop()
	for node != nil {
		temp.Next = &ListNode{
			Val:  node.Val,
			Next: nil,
		}
		temp = temp.Next

		if node.Next != nil {
			heap.Push(node.Next)
		}
		node = heap.Pop()

	}

	return retListHdNode
}

type Heap struct {
	arr []*ListNode
}

func (h *Heap) Pop() *ListNode {
	if len(h.arr) == 0 {
		return nil
	}

	if len(h.arr) == 1 {
		node := h.arr[0]
		h.arr = h.arr[:0]
		return node
	}

	retNode := h.arr[0]
	lastNode := h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.arr[0] = lastNode
	index := 0
	for index < len(h.arr) {
		childIndex1 := index*2 + 1
		childIndex2 := index*2 + 2
		childIndex := -1
		if childIndex1 < len(h.arr) {
			if childIndex2 < len(h.arr) {
				if h.arr[childIndex1].Val < h.arr[childIndex2].Val {
					childIndex = childIndex1
				} else {
					childIndex = childIndex2
				}
			} else {
				childIndex = childIndex1
			}
		} else {
			break
		}

		if h.arr[childIndex].Val < h.arr[index].Val {
			temp := h.arr[childIndex]
			h.arr[childIndex] = h.arr[0]
			h.arr[0] = temp
		}

		index = childIndex
	}
	return retNode
}

func (h *Heap) Push(node *ListNode) {
	h.arr = append(h.arr, node)
	index := len(h.arr) - 1
	parentIndex := (index - 1) / 2
	for parentIndex > 0 {
		if h.arr[index].Val < h.arr[parentIndex].Val {
			h.arr[index], h.arr[parentIndex] = h.arr[parentIndex], h.arr[index]
		}
		index = parentIndex
		parentIndex = (index - 1) / 2
	}
	if parentIndex == 0 {
		if h.arr[index].Val < h.arr[parentIndex].Val {

			h.arr[index], h.arr[parentIndex] = h.arr[parentIndex], h.arr[index]
		}
	}
	return
}
