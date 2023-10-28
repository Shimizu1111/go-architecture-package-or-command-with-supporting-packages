package subtract

import "testing"

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	if result != 2 {
		t.Errorf("Expected 2, but got %d", result)
	}
}
