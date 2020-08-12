package main

import (
	"fmt"
)

/*

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

*/

func primeFactorize(x int64) (primeFactors []int64) {
	var f int64 = 2
	inc := func(i *int64) {
		if *i == 2 {
			f++
		} else {
			*i += 2
		}
	}
	for f <= x {
		if f*f > x {
			break
		}
		if x%f != 0 {
			inc(&f)
			continue
		}

		// Check if it is prime
		var i int64 = 1
		factorFactors := []int64{}
		for i <= f {
			if f%i == 0 {
				factorFactors = append(factorFactors, i)
				if len(factorFactors) > 2 {
					break
				}
			}
			inc(&i)
		}
		if len(factorFactors) == 2 {
			primeFactors = append(primeFactors, f)
		}
		inc(&f)
	}
	return
}

func largestPrimeFactorLessThan(x int64) int64 {
	primeFactors := primeFactorize(x)
	return primeFactors[len(primeFactors)-1]
}

func p3() {
	fmt.Printf("The largest prime factor of the number 600851475143: %d", largestPrimeFactorLessThan(600851475143))
}
