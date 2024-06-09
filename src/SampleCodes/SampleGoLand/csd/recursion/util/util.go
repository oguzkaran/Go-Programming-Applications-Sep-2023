package util

import (
	"fmt"
	"math"
)

func Factorial(n int) int {
	if n <= 0 {
		return 1
	}

	result := 1

	for i := 2; i <= n; i++ {
		result *= i
	}

	return result
}

func Fibonacci(n int) int {
	if n <= 0 {
		return -1
	}

	if n <= 2 {
		return n - 1
	}

	prev1 := 1
	prev2 := 0
	var val int

	for i := 2; i < n; i++ {
		val = prev1 + prev2
		prev2 = prev1
		prev1 = val
	}

	return val
}

func Gcd(a, b int) int {
	minValue := int(min(math.Abs(float64(a)), math.Abs(float64(b))))

	for i := minValue; i >= 2; i-- {
		if a%i == 0 && b%i == 0 {
			return i
		}
	}

	return 1
}

func PrintNumber(a int) {
	if a == 0 {
		fmt.Printf("%c", '0')
		return
	}
	s := make([]rune, 11)
	i := 0
	isNegative := false

	if a < 0 {
		isNegative = true
		a = -a
	}

	for a != 0 {
		s[i] = rune(a%10 + '0')
		i++
		a /= 10
	}

	if isNegative {
		s[i] = '-'
		i++
	}

	for i = i - 1; i >= 0; i-- {
		fmt.Printf("%c", s[i])
	}
}

func PrintCollatz(a int) {
	panic("not yet implemented!...")
}

func PrintReverseASCII(s string) {
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Printf("%c", s[i])
	}
}
func ReverseString(s string) string {
	c := []rune(s)
	left := 0
	right := len(c) - 1

	for left < right {
		temp := c[left]
		c[left] = c[right]
		c[right] = temp
		left++
		right--
	}

	return string(c)
}
