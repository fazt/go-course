package calc

import "testing"

func TestAdd(t *testing.T) {
	var result int = Add(10, 20)
	if result != 30 {
		t.Error("Expected 30, got ", result)
	}
}

func TestSubtract(t *testing.T) {
	var result int = Subtract(10, 20)
	if result != -10 {
		t.Error("Expected -10, got ", result)
	}
}
