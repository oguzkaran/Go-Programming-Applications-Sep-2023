package numeric

import (
	"SampleGoLand/csd/numeric"
	"SampleGoLand/test/assert"
	"fmt"
	"testing"
)

func Test_IsPrime(t *testing.T) {
	inputs := []struct { //Şüphesiz veriler herhangi bir kaynaktan (source) da elde edilebilir
		val      int
		expected bool
	}{
		{val: 19, expected: true},
		{val: 2, expected: true},
		{val: 1000003, expected: true},
		{val: -1, expected: false},
		{val: 1, expected: false},
	}

	for _, data := range inputs {
		actual := numeric.IsPrime(data.val)
		assert.EqualsBool(t, fmt.Sprintf("Test_IsPrime FAIL: Value:%d, Expected:%t, Actual:%t", data.val, data.expected, actual), data.expected, actual)
	}
}
