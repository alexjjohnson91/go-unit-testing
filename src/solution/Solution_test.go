package solution

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"add positives", 2, 3, 5},
		{"add more positives", 4, 3, 7},
		{"add negatives", -1, -1, -2},
		{"add mix", -5, 5, 0},
		{"zero", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Solution{A: tt.a, B: tt.b}
			got := s.Add()
			if got != tt.want {
				t.Errorf("Add() = %d; want %d", got, tt.want)
			}
		})
	}
}

