package recursionutil

import "fmt"

func fibonacciRecur(n int) int {
	if n <= 2 {
		return n - 1
	}
	return fibonacciRecur(n-1) + fibonacciRecur(n-2)
}

func printReverseASCIIRecur(s string, i int) {
	if i == len(s) {
		return
	}
	printReverseASCIIRecur(s, i+1)
	fmt.Printf("%c", s[i])
}

func reverseStringRecur(c []rune, left, right int) {
	if left >= right {
		return
	}

	temp := c[left]
	c[left] = c[right]
	c[right] = temp
	reverseStringRecur(c, left+1, right-1)
}

func Factorial(n int) int {
	if n <= 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func Fibonacci(n int) int {
	if n <= 0 {
		return 1
	}

	return fibonacciRecur(n)
}

func Gcd(a, b int) int {
	panic("not yet implemented!...")
}

func PrintNumber(a int) {
	panic("not yet implemented!...")
}

func PrintReverseASCII(s string) {
	printReverseASCIIRecur(s, 0)
}

func ReverseString(s string) string {
	c := []rune(s)
	reverseStringRecur(c, 0, len(c)-1)
	return string(c)
}
