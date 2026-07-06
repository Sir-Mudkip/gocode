package step2

import (
	"testing"
)

func TestCalculateArea(t *testing.T) {
	width, height := 3, 5
	expected := 15
	area, err := CalculateArea(width, height)

	if area != expected {
		t.Errorf("Function recieved wrong data, expected %v, got %v", expected, area)
	}
	if err != nil {
		t.Error("There was an error something, check code")
	}
}
