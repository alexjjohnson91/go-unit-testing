package main

import (
	"fmt"
	"alexjohnson/go-unit-testing/src/solution"
)

func main() {
	fmt.Println("add: " + fmt.Sprint(solution.Add(3, 4)))
	fmt.Println("mul: " + fmt.Sprint(solution.Mul(3, 4)))
	fmt.Println("sub: " + fmt.Sprint(solution.Sub(3, 4)))
	fmt.Println("div: " + fmt.Sprint(solution.Div(4, 3.00)))
}
