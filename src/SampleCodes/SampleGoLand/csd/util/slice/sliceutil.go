package slice

import "math/rand"

func ForEach[T any](s []T, f func(int, T)) {
	ForEachCount(s, len(s), f)
}

func ForEachCount[T any](s []T, count int, f func(int, T)) {
	for i := 0; i < count; i++ {
		f(i, s[i])
	}
}

func ReduceSlice[T any](identity T, a []T, f func(T, T) T) T {
	result := identity

	for _, val := range a {
		result = f(result, val)
	}

	return result
}

func CopyIf[T any](src, dest []T, pred func(T) bool) int {
	index := 0

	for _, v := range src {
		if pred(v) {
			dest[index] = v
			index++
		}
	}

	return index
}

func Transform[T, R any](src []T, dest []R, f func(T) R) {
	for i, v := range src {
		dest[i] = f(v)
	}
}

func Equals[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func CreateSlice[T any](count int, f func() T) []T {
	result := make([]T, count)

	for i := 0; i < count; i++ {
		result[i] = f()
	}

	return result
}

func CreateRandomIntSlice(count, origin, bound int) []int {
	return CreateSlice(count, func() int { return rand.Intn(bound-origin) + origin })
}

func CreateRandomFloat64Slice(count int, origin, bound float64) []float64 {
	return CreateSlice(count, func() float64 { return rand.Float64()*(bound-origin) + origin })
}

func MaxFunc[T any](s []T, comp func(T, T) int) T {
	size := len(s)
	result := s[0]

	for i := 0; i < size; i++ {
		if comp(result, s[i]) < 0 {
			result = s[i]
		}
	}

	return result
}

func MinFunc[T any](s []T, comp func(T, T) int) T {
	size := len(s)
	result := s[0]

	for i := 0; i < size; i++ {
		if comp(result, s[i]) > 0 {
			result = s[i]
		}
	}

	return result
}
