package solution

import "log"

type Solution struct {
  A int
  B int
}

func (s Solution) Add(a int, b int) int {
  return (a + b)
}

func (s Solution) Sub(a int, b int) int {
  return (a - b)
}

func (s Solution) Mul(a int, b int) int {
  return (a * b)
}

func (s Solution) Div(a float32, b float32) float32 {
  if b == 0 {
    log.Fatal("ERROR: cannot divide by 0")
  }
  return (a / b)
}
