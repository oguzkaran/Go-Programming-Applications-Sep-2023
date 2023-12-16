/*
------------------------------------------------------------------------------------------------------------------------

	Aşağıdaki demo örnekte shallow copy yapıldığına dikkat ediniz

------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"SampleGoLand/csd/tuple"
)

type Point[T any] tuple.Pair[T, T]

func (p *Point[int]) Clone() Point[int] {
	return *p
}

func main() {

}
