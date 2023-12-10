package algorithm

func ForEach[T any](s []T, f func(int, T)) {
	for i, val := range s {
		f(i, val)
	}
}

func ReduceSlice[T any](identity T, a []T, f func(T, T) T) T {
	result := identity

	for _, val := range a {
		result = f(result, val)
	}

	return result
}
