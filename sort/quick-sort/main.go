package main

import "log"

func main() {
	arr := []int{0, 4, 2, 6, 5}
	//log.Println(Partition(arr, 4))
	QuickSort(arr)
	log.Println("arr :", arr)
}

func QuickSort(arr []int) {
	log.Println("arr :", arr)
	if len(arr) <= 1 {
		return
	}

	partIndex := Partition(arr, 0)
	log.Println("Index :", partIndex)
	QuickSort(arr[:partIndex])
	QuickSort(arr[partIndex+1:])
	return
}

func Partition(arr []int, index int) int {
	start := 0
	end := len(arr) - 1

	for {
		for arr[start] <= arr[index] && start < len(arr)-1 {
			start++
		}

		for arr[end] > arr[index] {
			end--
		}
		if start < end {
			arr[start], arr[end] = arr[end], arr[start]
		} else {
			break
		}

	}
	arr[end], arr[index] = arr[index], arr[end]
	return end
}
