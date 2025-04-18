package solution

import "fmt"

func Add(a int, b int) int {
	return (a + b)
}

func Sub(a int, b int) int {
	return (a - b)
}

func Mul(a int, b int) int {
	return (a * b)
}

func Div(a float32, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
