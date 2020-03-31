// Exercise 2.2: Write a general-purpos e unit-conversion program analogous to cf that reads numbers from its command-line arguments or
// from the standard input if there are no arguments, and converts each number into units like temperature in Celsius and Fahrenheit,
// length in feet and meters, weight in pounds and kilograms, and the like.

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/slamprou/tgopl-exercises-solutions/ch2/Exercise_2.2/lengthconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "mf: %v\n", err)
			os.Exit(1)
		}
		convert(t)
	}
}

func convert(t float64) {
	f := lengthconv.Feet(t)
	m := lengthconv.Meters(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, lengthconv.FToM(f), m, lengthconv.MToF(m))
}
