package cmath

import (
	"SampleGoLand/csd/cmath"
	"SampleGoLand/test/assert"
	"fmt"
	"testing"
)

func Test_Mid_Float64(t *testing.T) {
	var a float64 = 10.5
	var b float64 = 10.6
	var c float64 = 10.7
	var expected float64 = b
	actual := cmath.Mid(a, b, c)

	msg := fmt.Sprintf("Test Mid FAIL:Input:%f, %f, %f, , Expected:%f, Actual:%f\n", a, b, c, expected, actual)

	assert.EqualsFloat64(t, msg, expected, actual, 0.0001)
}

func Test_Mid_Int(t *testing.T) {
	a := 10
	b := 20
	c := 30
	expected := b
	actual := cmath.Mid(a, b, c)

	msg := fmt.Sprintf("Test Mid FAIL:Input:%d, %d, %d, , Expected:%d, Actual:%d\n", a, b, c, expected, actual)

	assert.Equals(t, msg, expected, actual)
}
