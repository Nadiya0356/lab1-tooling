package internal

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 5},
		{-1, 1, 0},
	}

	for _, tt := range tests {
		result := Add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, result)
		}
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		a, b int
		want int
		err  bool
	}{
		{10, 2, 5, false},
		{5, 0, 0, true},
	}

	for _, tt := range tests {
		result, e := Divide(tt.a, tt.b)

		if (e != nil) != tt.err {
			t.Errorf("error mismatch")
		}

		if result != tt.want {
			t.Errorf("expected %d, got %d", tt.want, result)
		}
	}
}
