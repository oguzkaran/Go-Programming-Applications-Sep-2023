package legacy

func ForEach(s []interface{}, f func(int, interface{})) {
	for i, val := range s {
		f(i, val)
	}
}

func ReduceSlice(identity interface{}, a []interface{}, f func(interface{}, interface{}) interface{}) interface{} {
	result := identity

	for _, val := range a {
		result = f(result, val)
	}

	return result
}
