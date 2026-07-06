package step3

import (
	"testing"
)

func TestCalculateArea(t *testing.T) {
	width, length := -4, 5

	area, err := CalculateArea(width, length)

	if area != 0 {
		t.Fail()
	}
	if err == nil {
		t.Fail()
	}
}

func TestCalculateNegativeLength(t *testing.T) {
	width, length := 5, -7

	area, err := CalculateArea(width, length)

	if area != 0 {
		t.Errorf("Area cannot be negative, expected %v, returned %v", 0, area)
	}
	if err == nil {
		t.Error("Invalid maths provided")
	}
}
