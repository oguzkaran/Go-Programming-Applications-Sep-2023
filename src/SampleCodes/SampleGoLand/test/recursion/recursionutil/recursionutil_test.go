package recursionutil

import (
	"SampleGoLand/csd/recursion/recursionutil"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func readTestFile(path string) []string {
	f, err := os.OpenFile(path, os.O_RDONLY, 0777)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	defer func() { _ = f.Close() }()

	reader := bufio.NewReader(f)

	var lines []string

	for {
		line, err := reader.ReadString('\n')

		if err != nil && err != io.EOF {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		if err == io.EOF && len(line) == 0 {
			break
		}
		line = strings.TrimRight(line, "\n")
		lines = append(lines, line)
	}

	return lines
}

func TestReverseString(t *testing.T) {
	values := readTestFile("reverse_string_test_values")
	for _, val := range values {
		testVal := strings.Split(val, " ")
		if len(testVal) != 2 {
			fmt.Fprintf(os.Stderr, "Wrong test input: %s\n", testVal)
			continue
		}
		str := testVal[0]
		reverseStr := testVal[1]
		if actual := recursionutil.ReverseString(str); actual != reverseStr {
			t.Errorf("ReverseStr(%s) = %s, want %s", str, actual, reverseStr)
		}
	}
}

func TestFactorial(t *testing.T) {

	values := readTestFile("factorial_test_values")

	for _, val := range values {
		testVal := strings.Split(val, " ")
		input, err := strconv.Atoi(testVal[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s cannot be converted into integer\n", testVal[0])
			continue
		}

		result, err := strconv.Atoi(testVal[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s cannot be converted into integer\n", testVal[0])
			continue
		}

		if actual := recursionutil.Factorial(input); actual != result {
			t.Errorf("Factorial(%d) = %d, want %d", input, actual, result)
		}
	}
}

func TestGcd(t *testing.T) {
	tests := []struct {
		a, b int
		gcd  int
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
		if result := recursionutil.Gcd(test.a, test.b); result != test.gcd {
			t.Errorf("Gcd(%d, %d) = %d, want %d", test.a, test.b, result, test.gcd)
		}
	}
}
