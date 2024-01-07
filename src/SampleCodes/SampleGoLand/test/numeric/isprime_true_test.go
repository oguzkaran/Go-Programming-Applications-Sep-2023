package numeric

import (
	"SampleGoLand/csd/numeric"
	"SampleGoLand/test/assert"
	"fmt"
	"testing"
)

func Test_IsPrimeTrue(t *testing.T) {
	input := 19

	assert.AssertTrue(t, fmt.Sprintf("Test_IsPrime_True FAIL: %d is not prime", input), numeric.IsPrime(input))
}
