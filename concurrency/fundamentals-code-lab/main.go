package main

import (
	"fmt"
	// Step 3,4 - Import
	//"sync"
	"time"
)

// Calculate Total Sales
func main() {
	startTime := time.Now() // Record the current time
	var totalSales int

	for _, salesRegion := range salesData {
		// Invoke to calculate Total for Each Region
		calculateRegionSales(salesRegion, &totalSales)
	}

    // Allow goroutines to complete


	fmt.Printf("Total %d , Time taken to calculate %s \n", totalSales, time.Since(startTime))
}

// Function to Calculate Total Sales per Region
func calculateRegionSales(salesRegion []int, totalSales *int) {
	regionTotal := 0
	for _, storeSales := range salesRegion {
		// Calculate region total
		regionTotal += storeSales
		time.Sleep(100 * time.Millisecond)
	}

	// Post Calculation of regionTotal
	*totalSales += regionTotal
}
