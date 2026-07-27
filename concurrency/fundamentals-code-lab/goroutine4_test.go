//go:build goroutine4

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"testing"
	"time"
)

func TestRunViaGoroutine(t *testing.T) {
	runTest(t, 4.0)
}

func runTest(t *testing.T, maxTime float64) {
	// Redirect standard output for testing
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Capture the output
	var buf bytes.Buffer
	go func() {
		// Call the function to be tested periodically
		main()
		// Reset standard output after the test
		w.Close()
		os.Stdout = old
	}()

	// Wait for the goroutine to finish
	time.Sleep(8 * time.Second)

	// Copy the captured output
	io.Copy(&buf, r)

	// Compare actual and expected output
	output := buf.String()
	println(" Output : " + output)

	regforTime := regexp.MustCompile(`\d+\.\d+`)
	totalTimeStr := regforTime.FindString(output)

	totalTime, err := strconv.ParseFloat(totalTimeStr, 64)
	if err != nil {
		t.Errorf("Check if the goroutine is terminating prematurely due to main goroutine termination")
		return
	}

	fmt.Println("Total Time ", totalTime)
	if totalTime > maxTime {
		t.Errorf("Program Taking more time than %f seconds. It seems sum is not being calculated concurrency or you've added sleep more than %f seconds.", maxTime, maxTime)
		return
	}

	regforTotal := regexp.MustCompile(`\d+`)
	matches := regforTotal.FindAllString(output, -1)

	var totalSales int64
	for _, match := range matches {
		n, err := strconv.ParseInt(match, 10, 64)
		if err != nil {
			fmt.Println("Error parsing integer:", err)
		}
		totalSales = n
		break
	}

	if totalSales != 50000000 {
		t.Errorf("The total sales amount %d is not correct it should be 50000000 , check if the goroutine is terminating prematurely due to main goroutine termination", totalSales)
	}
}
