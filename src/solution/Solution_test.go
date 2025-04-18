package solution

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"testing positives", 2, 3, 5},
		{"testing positives", 4, 3, 7},
		{"testing negatives", -1, -1, -2},
		{"testing negatives", -5, 5, 0},
		{"testing zeroes", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Solution{}
			got := s.Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add() = %d; want %d", got, tt.want)
			}
		})
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"testing positives", 3, 2, 1},
		{"testing positives", 10, 3, 7},
		{"testing negatives", -1, 1, -2},
		{"testing negatives", -5, -5, 0},
		{"testing zeroes", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Solution{}
			got := s.Sub(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Sub() = %d; want %d", got, tt.want)
			}
		})
	}
}
