package recursionutil

import (
	"SampleGoLand/csd/recursion/util"
	"SampleGoLand/test/assert"
	"SampleGoLand/test/recursion/file"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestReverseString(t *testing.T) {
	values := file.ReadTestFile("../testfiles/reverse_string_test_values")
	for _, val := range values {
		testVal := strings.Split(val, " ")
		str := testVal[0]
		expected := testVal[1]
		actual := util.ReverseString(str)
		message := fmt.Sprintf("ReverseStr(%s) = %s, want %s", str, actual, expected)

		assert.Equals(t, message, expected, actual)
	}
}

func TestFactorial(t *testing.T) {
	values := file.ReadTestFile("../testfiles/factorial_test_values")

	for _, val := range values {
		testVal := strings.Split(val, " ")
		input, _ := strconv.Atoi(testVal[0])
		expected, _ := strconv.Atoi(testVal[1])
		actual := util.Factorial(input)
		message := fmt.Sprintf("Factorial(%d) = %d, want %d", input, actual, expected)
		assert.Equals(t, message, expected, actual)
	}
}
