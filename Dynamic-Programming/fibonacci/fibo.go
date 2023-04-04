package main

func main() {

}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	prev := 0
	next := 1
	sum := 0
	for i := 2; i <= n; i++ {
		sum = prev + next
		prev = next
		next = sum
	}
	return sum
}
