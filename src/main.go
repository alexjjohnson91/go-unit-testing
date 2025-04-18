package main

import (
	"fmt"
	"alexjohnson/go-unit-testing/src/solution"
)

func main() {
	x := solution.Solution{}

	fmt.Println("add: " + fmt.Sprint(x.Add(3, 4)))
	fmt.Println("mul: " + fmt.Sprint(x.Mul(3, 4)))
	fmt.Println("sub: " + fmt.Sprint(x.Sub(3, 4)))
	fmt.Println("div: " + fmt.Sprint(x.Div(4, 3.00)))
}
