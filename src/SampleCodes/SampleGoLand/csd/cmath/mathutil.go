package cmath

import "math"

type Numeric interface {
	int | float32 | rune | float64
}

type UnsignedInteger interface {
	uint | uint8 | uint16 | uint64
}

type Additive interface {
	int | float64 | float32 | string
}

func Sqrt[T Numeric](a T) any {
	return math.Sqrt(float64(a))
}

func SqrtUnsigned[T UnsignedInteger](a T) any {
	return math.Sqrt(float64(a))
}

func Add[T Additive](a, b T) T {
	return a + b
}
