package shared

type RandomNumberGenerator interface {
	GenerateNumberInt(info *RandomNumberInfo[int], result *int) error
	GenerateNumberFloat64(info *RandomNumberInfo[float64], result *float64) error
	//...
}
