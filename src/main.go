package main

import (
	"fmt"
	"alexjohnson/go-unit-testing/src/solution"
)

func main() {
	x := solution.Solution{A: 2, B: 3}

	fmt.Println("add: " + fmt.Sprint(x.Add()))
	fmt.Println("mul: " + fmt.Sprint(x.Mul()))
	fmt.Println("sub: " + fmt.Sprint(x.Sub()))
}
