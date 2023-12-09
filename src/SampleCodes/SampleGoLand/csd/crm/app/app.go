package app

import (
	"SampleGoLand/csd/crm/employee"
	"math/rand"
)

func createWorker() *employee.Worker {
	return &employee.Worker{employee.Employee{"12345678112", "Ali", "mecidiyeköy"}, 8, 300}
}

func createManager() *employee.Manager {
	return &employee.Manager{employee.Employee{"12345678114", "Veli", "Şişli"}, "marketing", 100000.3}
}

func createEmployee() employee.Insurance {
	switch val := rand.Intn(2); val {
	case 0:
		return createWorker()
		//...
	default:
		return createManager()
	}
}

func RunApp() {

}
