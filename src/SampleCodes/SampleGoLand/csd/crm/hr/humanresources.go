package hr

import (
	"SampleGoLand/csd/crm/employee"
	"fmt"
)

type HumanResources struct {
	//...
}

func (hr *HumanResources) PayInsurance(insurance employee.Insurance) {
	fmt.Printf("Payment:%f\n", insurance.CalculatePayment())
}
