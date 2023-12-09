package employee

type SalesManager struct {
	manager  Manager
	extraFee float64
}

func (sm *SalesManager) CalculatePayment() float64 {
	return sm.manager.CalculatePayment() + sm.extraFee
}
