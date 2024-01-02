package main

func main() {

}

func kthGrammar(n int, k int) int {
	byteArr := make([]bool, n)

	byteArr[0] = false
	byteArr[1] = true

	if k < 2 {
		if !byteArr[k-1] {
			return 0
		} else {
			return 1
		}
	}

	elementsFilled := 2

	for row := 3; row <= n && elementsFilled <= k; row++ {
		for i := 0; i < elementsFilled; i++ {
			byteArr[elementsFilled+i] = !byteArr[i]
		}
		elementsFilled = elementsFilled * 2
	}

	if !byteArr[k-1] {
		return 0
	}
	return 1

}

// func kthGrammar(n int, k int) int {
// 	tempStr := "0"

// 	for row := 1; row < n; row++ {
// 		newTempStr := []byte(tempStr)
// 		for i := 0; i < len(newTempStr); i++ {
// 			if newTempStr[i] == '0' {
// 				newTempStr[i] = '1'
// 			} else {
// 				newTempStr[i] = '0'
// 			}
// 		}

// 		tempStr = tempStr + string(newTempStr)
// 	}
// 	if tempStr[k-1] == '0' {
// 		return 0
// 	}
// 	return 1
// }
