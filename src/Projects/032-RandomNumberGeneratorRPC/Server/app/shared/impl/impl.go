package impl

import (
	"Server/app/shared"
	"Server/csd/console"
	"errors"
	"fmt"
	"math/rand/v2"
)

type RandomNumberGenerator int

func (rg *RandomNumberGenerator) GenerateNumberInt(info *shared.RandomNumberInfo[int], result *int) error {
	if info.Min >= info.Bound {
		return errors.New(fmt.Sprintf("%d must be less than %d\n", info.Min, info.Bound))
	}

	*result = rand.IntN(info.Bound-info.Min) + info.Min

	_ = console.Write("%d ", *result)

	return nil
}

func (rg *RandomNumberGenerator) GenerateNumberFloat64(info *shared.RandomNumberInfo[float64], result *float64) error {
	if info.Min >= info.Bound {
		return errors.New(fmt.Sprintf("%f must be less than %f\n", info.Min, info.Bound))
	}

	*result = rand.Float64()*(info.Bound-info.Min) + info.Min

	_ = console.Write("%f ", *result)

	return nil
}
