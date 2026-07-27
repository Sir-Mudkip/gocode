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

	salesCh := make(chan int)

	for _, salesRegion := range salesData {
		// Invoke to calculate Total for Each Region
		go calculateRegionSales(salesRegion, salesCh)

	}

    // Allow goroutines to complete
	// Receive total sales from each goroutine asynchronously
	for i := 0; i < len(salesData); i++ {
		totalSales += <-salesCh
	}

	fmt.Printf("Total %d , Time taken to calculate %s \n", totalSales, time.Since(startTime))
}

// Function to Calculate Total Sales per Region
func calculateRegionSales(salesRegion []int, salesCh chan<- int) {
	regionTotal := 0
	for _, storeSales := range salesRegion {
		// Calculate region total
		regionTotal += storeSales
		time.Sleep(100 * time.Millisecond)
	}
	// Post Calculation of regionTotal
	salesCh <- regionTotal
}
