package cstack

type Stack[T any] struct {
	//...
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (stack *Stack[T]) Push(value *T) {
	//...
}

func (stack *Stack[T]) Peek() *T {
	return nil
}

func (stack *Stack[T]) Pop() *T {
	return nil
}

func (stack *Stack[T]) Len() int {
	return 0
}
