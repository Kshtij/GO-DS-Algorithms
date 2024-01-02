package main

func main() {

}

func longestPalindromeSubseq(s string) int {
	length := len(s)
	matrix := make([][]int, length+1)
	for i := 0; i <= length; i++ {
		matrix[i] = make([]int, length+1)
	}

	// palindrome := ""
	// for i := length - 1; i >= 0; i-- {
	// 	palindrome = palindrome + string(s[i])
	// }

	for i := 1; i <= length; i++ {
		for j := 1; j <= length; j++ {
			if s[length-1] == s[j-1] {
				matrix[i][j] = 1 + matrix[i-1][j-1]
			} else {
				max := matrix[i][j-1]
				if max < matrix[i-1][j] {
					max = matrix[i-1][j]
				}
				matrix[i][j] = max
			}
		}
	}
	return matrix[length][length]
}
