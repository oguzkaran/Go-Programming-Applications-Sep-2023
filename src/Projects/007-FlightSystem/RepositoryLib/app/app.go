package app

import (
	"SampleTimeServiceApp/app/data/dal"
	"fmt"
	"os"
)

func Run() {
	_, err := dal.NewFlightSearchHelper()

	if err != nil {
		fmt.Printf("Problem occurred:%s\n", err.Error())
		os.Exit(1)
	}
}
