package tuple

import "fmt"

type Triple[F, S, T any] struct {
	First  F
	Second S
	Third  T
}

func NewTriple[F, S, T any](first F, second S, third T) *Triple[F, S, T] {
	return &Triple[F, S, T]{first, second, third}
}

func (t *Triple[F, S, T]) String() string {
	return fmt.Sprintf("(%v, %v, %v)", t.First, t.Second, t.Third)
}
