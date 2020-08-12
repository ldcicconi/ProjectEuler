package main

import (
	"fmt"
	"time"
)

/*

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.

*/

func sumOfMultiples3And5(upperLimit int) (sum int) {
	// brute force -- nothing smart about this
	for i := 1; i < upperLimit; i++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
	}
	return
}

func p1() {
	// Timing how long this takes just out of curiosity
	startTime := time.Now()
	sumBelowOneThou := sumOfMultiples3And5(1000)
	delta := time.Since(startTime)

	fmt.Printf("Sum of natural numbers below 10 that are multiples of 3 or 5: %d\nTime to calculate: %s\n", sumBelowOneThou, delta)
}
