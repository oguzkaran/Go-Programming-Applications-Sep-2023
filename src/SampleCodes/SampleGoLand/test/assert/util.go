package assert

import (
	"cmp"
	"math"
	"testing"
)

func Equals[Type cmp.Ordered](t *testing.T, errorMessage string, expected, actual Type) {
	if expected != actual {
		t.Error(errorMessage)
	}
}

func NotEquals[Type cmp.Ordered](t *testing.T, errorMessage string, expected, actual Type) {
	if expected == actual {
		t.Error(errorMessage)
	}
}

func EqualsFloat64(t *testing.T, errorMessage string, expected, actual, delta float64) {
	if math.Abs(actual-expected) >= delta {
		t.Error(errorMessage)
	}
}

func NotEqualsFloat64(t *testing.T, errorMessage string, expected, actual, delta float64) {
	if math.Abs(actual-expected) < delta {
		t.Error(errorMessage)
	}
}

func EqualsFloat32(t *testing.T, errorMessage string, expected, actual, delta float32) {
	if float32(math.Abs(float64(actual-expected))) >= delta {
		t.Error(errorMessage)
	}
}

func NotEqualsFloat32(t *testing.T, errorMessage string, expected, actual, delta float32) {
	if float32(math.Abs(float64(actual-expected))) < delta {
		t.Error(errorMessage)
	}
}

func EqualsComplex64(t *testing.T, errorMessage string, expected, actual complex64) {
	if expected != actual {
		t.Error(errorMessage)
	}
}

func NotEqualsComplex64(t *testing.T, errorMessage string, expected, actual complex64) {
	if expected == actual {
		t.Error(errorMessage)
	}
}

func EqualsComplex128(t *testing.T, errorMessage string, expected, actual complex128) {
	if expected != actual {
		t.Error(errorMessage)
	}
}

func NotEqualsComplex128(t *testing.T, errorMessage string, expected, actual complex128) {
	if expected == actual {
		t.Error(errorMessage)
	}
}

func EqualsBool(t *testing.T, errorMessage string, expected, actual bool) {
	if expected != actual {
		t.Error(errorMessage)
	}
}

func True(t *testing.T, errorMessage string, result bool) {
	if !result {
		t.Error(errorMessage)
	}
}

func False(t *testing.T, errorMessage string, result bool) {
	if result {
		t.Error(errorMessage)
	}
}

func Nil[Type any](t *testing.T, errorMessage string, value any) {
	if value != nil {
		t.Error(errorMessage)
	}
}

func NotNil[Type any](t *testing.T, errorMessage string, value any) {
	if value == nil {
		t.Error(errorMessage)
	}
}
