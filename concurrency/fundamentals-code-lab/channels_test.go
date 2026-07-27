//go:build channels

package main

import (
	"testing"
)

// TestCalculateRegionSales tests the calculateRegionSales function
func TestCalculateRegionSales(t *testing.T) {
	salesRegion := []int{100, 200, 300}
	expected := 600

	// Initialize channel to receive result
	resultCh := make(chan int)

	// Execute the function under test
	go calculateRegionSales(salesRegion, resultCh)

	// Receive the result from the channel
	result := <-resultCh

	// Check if result matches expected value
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
