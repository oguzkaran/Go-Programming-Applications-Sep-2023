package cmath

import (
	"SampleGoLand/csd/cmath"
	"SampleGoLand/test/assert"
	"fmt"
	"testing"
)

func Test_CSqrt(t *testing.T) {
	var input float64 = 100
	var expected float64 = 10
	actual, _ := cmath.CSqrt(input)

	assert.AssertEqualsFloat64(t, fmt.Sprintf("Test Sqrt FAIL:Input:%f, Expected:%f, Actual:%f\n", input, expected, actual), expected, actual)
}
