package clone

type Cloneable[T any] interface {
	Clone() T
}
