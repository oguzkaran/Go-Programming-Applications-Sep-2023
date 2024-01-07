package test

import (
	"SampleGoLand/csd/cmath"
	"math"
	"testing"
)

func TestCLog(t *testing.T) {
	var input float64 = 100
	var expected float64 = math.Log(input)
	actual, _ := cmath.CLog(input)

	if actual != expected {
		t.Errorf("Test FAIL:Input:%f, Expected:%f, Actual:%f\n", input, expected, actual)
	} else {
		t.Logf("Test SUCCESS:Input:%f, Expected:%f, Actual:%f\n", input, expected, actual)
	}
}
