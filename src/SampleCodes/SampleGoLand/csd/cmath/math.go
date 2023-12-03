package cmath

import (
	"errors"
	"fmt"
	"math"
)

func CLog(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("indeterminate")
	}

	if a == 0 {
		return 0, errors.New("undefined")
	}

	return math.Log(a), nil
}

func CSqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("undefined -> value:%f", a)
	}

	return math.Sqrt(a), nil
}

func CLog2(a float64) float64 {
	if a < 0 {
		panic("indeterminate:")
	}

	if a == 0 {
		panic("indeterminate:")
	}

	return math.Log2(a)
}
