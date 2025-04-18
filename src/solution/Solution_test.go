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
			got := Add(tt.a, tt.b)
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
			got := Sub(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Sub() = %d; want %d", got, tt.want)
			}
		})
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		name          string
		a, b          float32
		expected      float32
		expectedError bool
	}{
		{"testing positives", 10, 2, 5, false},
		{"testing positives", 9, 3, 3, false},
		{"testing positives", 0, 3, 0, false},
		{"testing positives", 1, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, actualError := Div(tt.a, tt.b)
			if (actualError != nil) != tt.expectedError {
				t.Errorf("Div(%f, %f): unexpected error = %v", tt.a, tt.b, actualError)
        return
			}
			if actual != tt.expected {
				t.Errorf("Div() = %f; expected %f", actual, tt.expected)
			}
		})
	}
}

func TestMul(t *testing.T) {
  tests := []struct {
    name string
    a, b int
    expected int
  }{
		{"testing positives", 5, 2, 10},
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      actual := Mul(tt.a, tt.b)
      if actual != tt.expected {
        t.Errorf("Mul() = %d; expected %d", actual, tt.expected)
      }
    })
  }
}
