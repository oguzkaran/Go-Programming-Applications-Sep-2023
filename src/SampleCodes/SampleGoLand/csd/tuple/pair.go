package tuple

import "fmt"

type Pair[F, S any] struct {
	First  F
	Second S
}

func NewPair[F, S any](first F, second S) *Pair[F, S] {
	return &Pair[F, S]{first, second}
}

func (p *Pair[F, S]) String() string {
	return fmt.Sprintf("(%v, %v)", p.First, p.Second)
}
