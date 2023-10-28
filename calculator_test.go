package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	a, b := 10, 5
	sum, difference := calculate(a, b)

	if sum != 15 {
		t.Errorf("Expected sum 15, but got %d", sum)
	}
	if difference != 5 {
		t.Errorf("Expected difference 5, but got %d", difference)
	}
}
