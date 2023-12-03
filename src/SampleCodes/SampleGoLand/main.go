/*
------------------------------------------------------------------------------------------------------------------------

	Interface

------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"SampleGoLand/csd/cmath"
	"SampleGoLand/csd/console"
)

func main() {
	a := console.ReadFloat64("Input first value:", "Invalid value!...")
	result := cmath.CLog2(a)

	console.WriteLine("log2(%f) = %f", a, result)
}
