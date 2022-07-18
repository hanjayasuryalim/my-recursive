package my_recursive

import "log"

func Factorial(n int) int {
	log.Printf("Current n :%d", n)
	if n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}
