package main

import (
	"fmt"
	"os"
)

const (
	ob = '('
	cb = ')'
)

func main() {
	bracketStr := os.Args[1]
	fmt.Println(longestValidParentheses(bracketStr))
}

func longestValidParentheses(s string) int {

	length := len(s)
	stack := make([]int, length+1)
	stack[0] = -1
	stackIndex := 0
	max := 0

	for i := 0; i < length; i++ {

		switch s[i] {
		case ob:
			stackIndex++
			stack[stackIndex] = i

		case cb:
			if stackIndex == 0 || s[stack[stackIndex]] == cb {
				stackIndex++
				stack[stackIndex] = i

			} else {
				if s[stack[stackIndex]] == ob {
					stackIndex--
				}

				len := i - stack[stackIndex]
				if len > max {
					max = len
				}
			}
		}
	}
	return max
}
