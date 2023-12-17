package slice

import "math/rand"

func ForEach[S ~[]T, T any](s S, f func(int, T)) {
	ForEachCount(s, len(s), f)
}

func ForEachCount[S ~[]T, T any](s S, count int, f func(int, T)) {
	for i := 0; i < count; i++ {
		f(i, s[i])
	}
}

func ReduceSlice[S ~[]T, T any](identity T, s S, f func(T, T) T) T {
	result := identity

	for _, val := range s {
		result = f(result, val)
	}

	return result
}

func CopyIf[S ~[]T, T any](src, dest S, pred func(T) bool) int {
	index := 0

	for _, v := range src {
		if pred(v) {
			dest[index] = v
			index++
		}
	}

	return index
}

func Transform[S ~[]T, D ~[]R, T, R any](src S, dest D, f func(T) R) {
	for i, v := range src {
		dest[i] = f(v)
	}
}

func Equals[S ~[]T, T comparable](s1, s2 S) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, _ := range s1 {
		if s1[i] != s2[i] {
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
	return CreateSlice[int](count, func() int { return rand.Intn(bound-origin) + origin })
}

func CreateRandomFloat64Slice(count int, origin, bound float64) []float64 {
	return CreateSlice(count, func() float64 { return rand.Float64()*(bound-origin) + origin })
}

func MaxFunc[S ~[]T, T any](s S, comp func(T, T) int) T {
	size := len(s)
	result := s[0]

	for i := 0; i < size; i++ {
		if comp(result, s[i]) < 0 {
			result = s[i]
		}
	}

	return result
}

func MinFunc[S ~[]T, T any](s S, comp func(T, T) int) T {
	size := len(s)
	result := s[0]

	for i := 0; i < size; i++ {
		if comp(result, s[i]) > 0 {
			result = s[i]
		}
	}

	return result
}
