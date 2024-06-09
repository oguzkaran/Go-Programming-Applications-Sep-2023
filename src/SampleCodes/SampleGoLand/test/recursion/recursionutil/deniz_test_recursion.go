package Go_Programming_Applications_Sep_2023_main

import (
	"SampleGoLand/csd/recursion/recursionutil"
	"bytes"
	"os"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 1},
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
		if actual := recursionutil.Fibonacci(test.input); actual != test.expected {
			t.Errorf("Fibonacci(%d) = %d, want %d", test.input, actual, test.expected)
		}
	}
}

func getOutputOfPrintNumberV2(f func(int), a int) string {
	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f(a)

	w.Close()
	os.Stdout = old
	buf.ReadFrom(r)
	return buf.String()
}

func TestPrintNumberV2(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1234, "1234"},
		{0, ""},
		{56, "56"},
		{789, "789"},
		{-1234, "-1234"},
		{-56, "-56"},
		{-789, "-789"},
	}

	for _, test := range tests {
		output := getOutputOfPrintNumberV2(recursionutil.PrintNumberV2, test.input)
		if output != test.expected {
			t.Errorf("PrintNumberV2(%d) = %q; want %q", test.input, output, test.expected)
		}
	}
}
