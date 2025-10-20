package main

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
		{-5, -3, -8},
		{100, 200, 300},
	}

	for _, test := range tests {
		result := sum(test.a, test.b)
		if result != test.expected {
			t.Errorf("sum(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{5, 3, 2},
		{0, 0, 0},
		{1, 1, 0},
		{-5, -3, -2},
		{100, 50, 50},
		{3, 5, -2},
	}

	for _, test := range tests {
		result := subtract(test.a, test.b)
		if result != test.expected {
			t.Errorf("subtract(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 6},
		{0, 5, 0},
		{1, 1, 1},
		{-2, 3, -6},
		{-2, -3, 6},
		{10, 10, 100},
	}

	for _, test := range tests {
		result := multiply(test.a, test.b)
		if result != test.expected {
			t.Errorf("multiply(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		a, b, expected int
		shouldError    bool
	}{
		{6, 2, 3, false},
		{0, 5, 0, false},
		{1, 1, 1, false},
		{-6, 2, -3, false},
		{-6, -2, 3, false},
		{5, 2, 2, false}, // divisão inteira
		{10, 0, 0, true}, // divisão por zero
	}

	for _, test := range tests {
		result := divide(test.a, test.b)

		if test.shouldError {
			// Para divisão por zero, esperamos 0 como retorno
			if result != 0 {
				t.Errorf("divide(%d, %d) = %d; want 0 (division by zero)", test.a, test.b, result)
			}
		} else {
			if result != test.expected {
				t.Errorf("divide(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
			}
		}
	}
}
