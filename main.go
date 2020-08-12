package main

import (
	"flag"
	"fmt"
	"time"
)

var funcs []func()

var problem int

func init() {
	flag.IntVar(&problem, "problem", 1, "which ProjectEuler problem to run")
	flag.Parse()
	fmt.Println("Problem:", problem)
	fmt.Println()

	funcs = append(funcs, p1)
	funcs = append(funcs, p2)
}

func main() {
	if len(funcs) < problem {
		fmt.Println("That problem is not registered.")
		return
	}

	startTime := time.Now()
	funcs[problem-1]()
	delta := time.Since(startTime)
	fmt.Printf("\nDelta for problem %d: %s\n", problem, delta)
}
