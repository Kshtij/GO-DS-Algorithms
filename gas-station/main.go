package main

import "log"

func main() {
	log.Println("starting point :", canCompleteCircuit([]int{5, 8, 2, 8}, []int{6, 5, 6, 6}))
}

func canCompleteCircuit(gas []int, cost []int) int {
	length := len(gas)
	aux := make([]int, length)

	startIndices := make([]int, 0, 1)
	s := 0
	for i := 0; i < length; i++ {
		aux[i] = gas[i] - cost[i]
		s = s + aux[i]
		if aux[i] > 0 {
			startIndices = append(startIndices, i)
		}
	}

	if s < 0 {
		return -1
	}

	if length == 1 && aux[0] >= 0 {
		return 0
	}

	for _, i := range startIndices {

		start := i
		destination := start
		fuel := 0
		for {
			fuel = fuel + aux[start]
			if fuel < 0 {
				break
			}

			start = (start + 1)
			if start == length {
				start = 0
			}
			if start == destination {
				return start
			}
		}

	}
	return -1

}
