package main

import (
	"fmt"
	"alexjohnson/go-unit-testing/src/solution"
)

func main() {
	x := solution.Solution{A: 2, B: 3}

	fmt.Println("solution: " + fmt.Sprint(x.Mul()))
}
