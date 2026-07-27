//go:build errorchannels

package main

import (
	"testing"
	"time"
)

func TestCalculateRegionSales(t *testing.T) {
	// Test case 1: Valid input, no errors
	salesRegion := []int{100, 200, 300}
	expectedTotal := 600
	salesCh := make(chan int)
	errorCh := make(chan error)

	// Execute calculateRegionSales
	go calculateRegionSales(salesRegion, salesCh, errorCh)

	// Wait for some time to let the goroutine finish
	time.Sleep(1 * time.Second)

	// Validate output
	select {
	case total := <-salesCh:
		if total != expectedTotal {
			t.Errorf("Unexpected total, got: %d, want: %d", total, expectedTotal)
		}
	case err := <-errorCh:
		t.Errorf("Unexpected error: %v", err)
	case <-time.After(2 * time.Second):
		t.Error("Timeout occurred while waiting for sales data")
	}

	// Test case 2: Invalid input with negative values
	salesRegionNegative := []int{100, -200, 300}
	expectedError := "SKIPPED REGION! Found store with sale value [-200] less than 0. Please recheck data for region"

	// Execute calculateRegionSales with negative values
	go calculateRegionSales(salesRegionNegative, salesCh, errorCh)

	// Wait for some time to let the goroutine finish
	time.Sleep(1 * time.Second)

	// Validate output
	select {
	case total := <-salesCh:
		t.Errorf("Unexpected total: %d", total)
	case err := <-errorCh:
		if err == nil {
			t.Errorf("Unexpected error, got: %v, want: %s", err, expectedError)
		}
	case <-time.After(2 * time.Second):
		t.Error("Timeout occurred while waiting for sales data ( Check if error is posted on errorCh for storeSales <0 )")
	}
}
