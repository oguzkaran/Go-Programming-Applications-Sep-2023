package cmath

import (
	"SampleGoLand/csd/cmath"
	"SampleGoLand/test/assert"
	"fmt"
	"testing"
)

func Test_Signum_Float64(t *testing.T) {
	var a float64 = -10.5
	expected := -1
	actual := cmath.Signum(a)

	msg := fmt.Sprintf("Test Signum FAIL:Input:%f, Expected:%d, Actual:%d\n", a, expected, actual)

	assert.Equals(t, msg, expected, actual)
}

func Test_Signum_Int(t *testing.T) {
	a := -10
	expected := -1
	actual := cmath.Signum(a)

	msg := fmt.Sprintf("Test Signum FAIL:Input:%d, Expected:%d, Actual:%d\n", a, expected, actual)

	assert.Equals(t, msg, expected, actual)
}
