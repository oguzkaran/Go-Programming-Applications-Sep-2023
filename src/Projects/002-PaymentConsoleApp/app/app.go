package app

import (
	"PaymentServiceApp/csd/console"
	"PaymentServiceApp/data/entity"
	"PaymentServiceApp/service"
	"fmt"
	"os"
)

func getUser() *entity.User {
	username := console.ReadString("Input username:")

	if username == "quit" {
		return nil
	}

	return entity.NewUser(username, console.ReadString("Input password:"), console.ReadString("Input name:"),
		console.ReadString("Input phone:"))
}

func Run() {
	connStr := "postgresql://postgres:csd1993@localhost/gs23_paymentdb?sslmode=disable"

	s, err := service.NewPaymentAppDataService(connStr)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Connected...")

	for {
		username := console.ReadString("Input username:")

		if username == "quit" {
			break
		}

		user, err := s.FindUserByUsername(username)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		if user != nil {
			fmt.Printf("User %s found:%s\n", username, user.Name)
		} else {
			fmt.Printf("User %s not found\n", username)
		}

	}

	err = s.Close()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
