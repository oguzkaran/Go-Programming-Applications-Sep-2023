package matrix

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float64 | ~float32 | ~complex64 | ~complex128
}

type Matrix[T Numeric] struct {
	matrix [][]T
}

func New[T Numeric](row, col int) *Matrix[T] {
	return nil
}

func (m *Matrix[T]) GetRow() int { return 0 }

func (m *Matrix[T]) GetColumn() int { return 0 }

func (m *Matrix[T]) Get(i, j int) *T { return nil }

func (m *Matrix[T]) Set(i, j int, t *T) {

}

func (m *Matrix[T]) Add(other *Matrix[T]) *Matrix[T] {
	return nil
}

func (m *Matrix[T]) AddBy(val T) {

}

func (m *Matrix[T]) Subtract(other *Matrix[T]) *Matrix[T] {
	return nil
}

func (m *Matrix[T]) SubtractBy(val T) {

}

func (m *Matrix[T]) Multiply(other *Matrix[T]) *Matrix[T] {
	return nil
}

func (m *Matrix[T]) MultiplyBy(val T) {

}

func (m *Matrix[T]) DivideBy(val T) {

}

func (m *Matrix[T]) Transpose() {

}

func (m *Matrix[T]) Fill(f func() T) bool {
	return true
}

func (m *Matrix[T]) Equals(other *Matrix[T]) bool {
	return true
}
