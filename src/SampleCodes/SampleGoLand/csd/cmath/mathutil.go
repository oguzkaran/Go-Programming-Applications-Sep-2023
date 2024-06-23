package cmath

import (
	"cmp"
	"golang.org/x/exp/constraints"
	"math"
)

type Numeric interface {
	~int | ~float32 | ~rune | ~float64
}

type UnsignedInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint64
}

func Sqrt[T Numeric](a T) any {
	return math.Sqrt(float64(a))
}

func SqrtUnsigned[T UnsignedInteger](a T) any {
	return math.Sqrt(float64(a))
}

func AbsFloat32(a float32) float32 {
	return float32(math.Abs(float64(a)))
}

func Abs[T constraints.Signed](a T) T {
	if a < T(0) {
		return -a
	}

	return a
}

func Add[T cmp.Ordered](a, b T) T {
	return a + b
}

func Mid[T cmp.Ordered](a, b, c T) T {
	if a <= b && b <= c || c <= b && b <= a {
		return b
	}

	if b <= a && a <= c || c <= a && a <= b {
		return a
	}

	return c
}

func Signum[T Numeric](a T) int {
	result := -1

	if a > T(0) {
		result = 1
	} else if a == T(0) {
		result = 0
	}

	return result
}
