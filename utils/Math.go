package utils

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func Modulo(a, b int) int {
	return (a%b + b) % b
}
