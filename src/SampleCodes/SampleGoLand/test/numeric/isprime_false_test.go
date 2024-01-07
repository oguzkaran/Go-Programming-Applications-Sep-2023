package numeric

import (
	"SampleGoLand/csd/numeric"
	"SampleGoLand/test/assert"
	"fmt"
	"testing"
)

func Test_IsPrimeFalse(t *testing.T) {
	input := 1

	assert.AssertFalse(t, fmt.Sprintf("Test_IsPrime_True FAIL: %d is prime", input), numeric.IsPrime(input))
}
