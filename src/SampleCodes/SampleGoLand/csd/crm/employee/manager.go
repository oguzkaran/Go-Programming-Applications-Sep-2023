package employee

type Manager struct {
	Emp        Employee
	Department string
	Salary     float64
}

//...

func (m *Manager) CalculatePayment() float64 {
	return m.Salary * 1.5
}
