package main

import "fmt"

var heap = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {
	buildMaxHeap(heap)
	fmt.Println("Max Heap = ", heap)

	fmt.Println("Extracting max")
	fmt.Println("max = ", extractMax())
	fmt.Println("Max Heap = ", heap)

	fmt.Println("Inserting element 9")
	insert(9)
	fmt.Println("Max Heap = ", heap)
}

func maxHeapify(arr []int, index int) {
	leftChildIndex := index*2 + 1  //Getting left child index
	rightChildIndex := index*2 + 2 //Getting right child index
	var childIndex int
	//Check if both children are within array and select the largest child which is greater than current node
	if leftChildIndex < len(arr) && rightChildIndex < len(arr) {
		if arr[leftChildIndex] > arr[rightChildIndex] && arr[index] < arr[leftChildIndex] {
			childIndex = leftChildIndex
		} else if arr[leftChildIndex] < arr[rightChildIndex] && arr[index] < arr[rightChildIndex] {
			childIndex = rightChildIndex
		} else {
			return
		}
	} else if leftChildIndex < len(arr) && arr[index] < arr[leftChildIndex] {
		childIndex = leftChildIndex
	} else {
		return
	}

	//swap
	temp := arr[index]
	arr[index] = arr[childIndex]
	arr[childIndex] = temp

	maxHeapify(arr, childIndex)
	return
}

func buildMaxHeap(arr []int) {
	//start from middle as, leaves can be considered as heap
	for i := len(arr) / 2; i >= 0; i-- {
		maxHeapify(arr, i)
	}
}

func extractMax() int {
	max := heap[0]
	heap[0] = heap[len(heap)-1]
	heap = heap[:len(heap)-1]
	maxHeapify(heap, 0)
	return max
}

func insert(element int) {
	heap = append(heap, element)
	index := len(heap) - 1
	var parentIndex int
	if index%2 == 0 {
		parentIndex = (index - 2) / 2
	} else {
		parentIndex = (index - 1) / 2
	}

	for parentIndex >= 0 && heap[parentIndex] < heap[index] {
		temp := heap[index]
		heap[index] = heap[parentIndex]
		heap[parentIndex] = temp
		index = parentIndex
		if index%2 == 0 {
			parentIndex = (index - 2) / 2
		} else {
			parentIndex = (index - 1) / 2
		}
	}

}
