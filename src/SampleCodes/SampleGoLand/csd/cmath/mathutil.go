package cmath

import (
	"cmp"
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

func Add[T cmp.Ordered](a, b T) T {
	return a + b
}
