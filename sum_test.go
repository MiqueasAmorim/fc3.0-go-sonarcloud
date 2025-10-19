package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("sum(2, 3) = %d; want %d", result, expected)
	}
}
