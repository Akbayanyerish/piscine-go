package piscine

func Fibonacci(index int) int {
	result := 1
	if index < 0 {
		result = -1
	} else {
		result = Fibonacci(index-2) + Fibonacci(index-1)
	}
	return result
}
