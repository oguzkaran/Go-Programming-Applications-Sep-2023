package cqueue

type Queue[E any] struct {
	//...
}

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (stack *Queue[T]) Enqueue(value *T) {
	//...
}

func (stack *Queue[T]) Dequeue() *T {
	return nil
}

func (stack *Queue[T]) Peek() *T {
	return nil
}

func (stack *Queue[T]) Len() int {
	return 0
}
