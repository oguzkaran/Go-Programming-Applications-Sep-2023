package algorithm

func Swap[T any](p1, p2 *T) {
	temp := *p1
	*p1 = *p2
	*p2 = temp
}
