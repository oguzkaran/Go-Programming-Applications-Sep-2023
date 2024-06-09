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

// Author: Sinan Kesen
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

// Author: Sinan Kesen
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

// Author: Sinan Kesen
func TestGcd(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{12, 16, 4},
		{10, 20, 10},
		{6, 3, 3},
		{81, 6, 3},
		{8148, 3, 3},
		{49, 7, 7},
		{64, 48, 16},
	}

	for _, test := range tests {
		actual := util.Gcd(test.a, test.b)
		message := fmt.Sprintf("Gcd(%d, %d) = %d, want %d", test.a, test.b, actual, test.expected)
		assert.Equals(t, message, test.expected, actual)
	}
}

// Author: Deniz Öğüt
func TestFibonacci(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, -1},
		{1, 0},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 3},
		{6, 5},
		{7, 8},
		{8, 13},
		{9, 21},
		{10, 34},
	}

	for _, test := range tests {
		actual := util.Fibonacci(test.input)
		message := fmt.Sprintf("Fibonacci(%d) = %d, want %d", test.input, actual, test.expected)
		assert.Equals(t, message, test.expected, actual)
	}
}

// Author: Deniz Öğüt
func TestPrintNumber(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1234, "1234"},
		{0, "0"},
		{56, "56"},
		{789, "789"},
		{-1234, "-1234"},
		{-56, "-56"},
		{-789, "-789"},
	}

	for _, test := range tests {
		output := file.GetOutputOfPrintNumber(util.PrintNumber, test.input)
		message := fmt.Sprintf("PrintNumber(%d) = %q; want %q", test.input, output, test.expected)
		assert.Equals(t, message, test.expected, output)
	}
}
