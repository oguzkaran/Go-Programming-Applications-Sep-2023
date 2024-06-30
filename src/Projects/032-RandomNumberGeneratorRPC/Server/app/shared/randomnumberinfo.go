package shared

type RandomNumberInfo struct {
	Min, Bound int
}

type RandomNumberGenerator interface {
	GenerateNumber(info *RandomNumberInfo, result *int) error
	//...
}
