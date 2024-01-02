package main

import "log"

func main() {
	log.Println(bestClosingTime("YYNY"))
}

type Count struct {
	Ncount int
	Ycount int
}

func bestClosingTime(customers string) int {
	arr := make([]Count, len(customers))

	if customers[len(customers)-1] == 'Y' {
		arr[len(customers)-1].Ycount = 1
	} else {
		arr[len(customers)-1].Ncount = 1
	}

	for i := len(customers) - 2; i >= 0; i-- {
		if customers[i] == 'Y' {
			arr[i].Ycount = arr[i+1].Ycount + 1
			arr[i].Ncount = arr[i+1].Ncount
		} else {
			arr[i].Ncount = arr[i+1].Ncount + 1
			arr[i].Ycount = arr[i+1].Ycount
		}
	}

	for i := 0; i < len(customers); i++ {
		if arr[i].Ncount >= arr[i].Ycount {
			return i
		}
	}
	return len(customers) - 1
}
