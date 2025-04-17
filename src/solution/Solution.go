package solution

type Solution struct {
  A int
  B int
}

func (s Solution) Add() int {
  return (s.A + s.B)
}

func (s Solution) Sub() int {
  return (s.A - s.B)
}
