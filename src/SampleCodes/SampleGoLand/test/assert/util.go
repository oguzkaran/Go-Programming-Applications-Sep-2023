package assert

import "testing"

func AssertEqualsFloat64(t *testing.T, errorMessage string, expected, actual float64) {
	if expected != actual {
		t.Error(errorMessage)
	}
}

func AssertEqualsBool(t *testing.T, errorMessage string, expected, actual bool) {
	if expected != actual {
		t.Error(errorMessage)
	}
}

func AssertTrue(t *testing.T, errorMessage string, result bool) {
	if !result {
		t.Error(errorMessage)
	}
}
func AssertFalse(t *testing.T, errorMessage string, result bool) {
	if result {
		t.Error(errorMessage)
	}
}
