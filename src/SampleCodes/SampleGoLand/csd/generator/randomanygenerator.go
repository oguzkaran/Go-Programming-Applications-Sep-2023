package generator

import (
	"SampleGoLand/csd/util/str"
	"math/rand"
)

func createRandomValue() any {
	var result any

	switch val := rand.Intn(5); val {
	case 0:
		result = rand.Intn(100)
	case 1:
		result = rand.Float32()
	case 2:
		result = rand.Float64()
	case 3:
		result = str.GenerateRandomTextEN(rand.Intn(5) + 5)
	case 4:
		result = complex(rand.Float64()*20+10, rand.Float64()*20+10)
	}

	return result
}
func Create(count int) []any {
	a := make([]any, count)

	for i := 0; i < count; i++ {
		a[i] = createRandomValue()
	}

	return a
}
