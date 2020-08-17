package main

import (
	"flag"
	"fmt"
	"time"
)

var funcs map[string]func()

var problem string

func init() {
	flag.StringVar(&problem, "problem", "p1", "which ProjectEuler problem to run")
	flag.Parse()

	funcs = make(map[string]func())
	funcs["p1"] = p1
	funcs["p2"] = p2
	funcs["p3"] = p3
	funcs["p5"] = p5

	fmt.Println("Problem:", problem)
	fmt.Println()
}

func main() {
	p, ok := funcs[problem]
	if !ok {
		fmt.Println("That problem is not registered.")
		return
	}

	startTime := time.Now()
	p()
	delta := time.Since(startTime)
	fmt.Printf("\nDelta for problem %s: %s\n", problem, delta)
}
