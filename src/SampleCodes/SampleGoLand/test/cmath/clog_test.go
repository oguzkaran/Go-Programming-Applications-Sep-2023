package cmath

import (
	"SampleGoLand/csd/cmath"
	"SampleGoLand/test/assert"
	"fmt"
	"math"
	"testing"
)

func Test_CLog(t *testing.T) {
	var input float64 = 100
	var expected float64 = math.Log(input)
	actual, _ := cmath.CLog(input)

	assert.AssertEqualsFloat64(t, fmt.Sprintf("Test CLog FAIL:Input:%f, Expected:%f, Actual:%f\n", input, expected, actual), expected, actual)
}

func Test_CLog_Error(t *testing.T) {
	var input float64 = -2

	_, err := cmath.CLog(input)

	if err == nil {
		t.Errorf("Function must give error for the input:%f", input)
	} else {
		t.Logf("SUCCESS: Message:%s", err.Error())
	}
}
