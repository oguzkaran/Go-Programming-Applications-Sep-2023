package employee

type Worker struct {
	Emp        Employee
	HourPerDay int
	FeePerHour float64
}

func (w *Worker) CalculatePayment() float64 {
	return float64(w.HourPerDay) * w.FeePerHour * 30
}
