package solution

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
