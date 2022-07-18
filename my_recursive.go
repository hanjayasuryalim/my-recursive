package my_recursive

func FactorialTailRecursive(n int, total int) int {
	if n == 1 {
		return total
	}
	total *= n
	return FactorialTailRecursive(n-1, total)
}
