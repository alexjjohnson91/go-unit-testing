package main

import (
	"fmt"
	"alexjohnson/go-unit-testing/src/solution"
)

func main() {
	x := solution.Solution{2, 3}

	fmt.Println("solution: " + fmt.Sprint(x.Add()))
}
