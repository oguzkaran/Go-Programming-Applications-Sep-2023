package app

import (
	"PaymentServiceApp/csd/console"
	"PaymentServiceApp/service"
	"PaymentServiceApp/service/dto"
	"fmt"
	"os"
)

const (
	users int = iota + 1
	products
	payments
	exit
)

const (
	register int = iota + 1
	login
	find
	quit
)

func getUser() *dto.UserSaveDTO {
	username := console.ReadString("Input username:")

	if username == "quit" {
		return nil
	}

	return dto.NewUserSaveDTO(username, console.ReadString("Input password:"), console.ReadString("Input name:"),
		console.ReadString("Input phone:"))
}

func displayMainMenu() {
	fmt.Println("Main operations:")
	fmt.Println("1.User Operations")
	fmt.Println("2.Product Operations")
	fmt.Println("3.Payment Operations")
	fmt.Println("4.Exit")
	fmt.Print("Option:")
}

func displayUsersMenu() {
	fmt.Println("User Operations")
	fmt.Println("1.Register")
	fmt.Println("2.Login")
	fmt.Println("3.Find")
	fmt.Println("4.Quit")
	fmt.Print("Option:")
}

func usersProc(dataService *service.PaymentAppDataService) {
	for {
		displayUsersMenu()
		option := console.ReadInt("", "You must input option as a number!...")
		if !doForUsersOption(option, dataService) {
			break
		}
	}
}

func productsProc(dataService *service.PaymentAppDataService) {
	fmt.Println("Products selected!...")
}

func paymentsProc(dataService *service.PaymentAppDataService) {
	fmt.Println("Payments selected!...")
}

func exitProc(dataService *service.PaymentAppDataService) {
	dataService.Close()
	fmt.Println("C and System Programmers Association")
	fmt.Println("Thanks!...")
	os.Exit(0)
}

func doForMainOption(option int, dataService *service.PaymentAppDataService) {
	switch option {
	case users:
		usersProc(dataService)
	case products:
		productsProc(dataService)
	case payments:
		paymentsProc(dataService)
	case exit:
		exitProc(dataService)
	default:
		fmt.Println("Invalid option!...")
	}
}

func registerProc(dataService *service.PaymentAppDataService) {
	fmt.Println("Register selected!...")
	user := getUser()
	if dataService.ExistsUserByUsername(user.Username) {
		fmt.Printf("%s already registered!....", user.Username)
		return
	}

	err := dataService.SaveUser(user)

	if err == nil {
		fmt.Printf("%s registered successfully!..\n", user.Username)
	} else {
		fmt.Printf("Problem occurred:%s\n", err.Error())
	}
}

func loginProc(dataService *service.PaymentAppDataService) {
	fmt.Println("Login selected!...")
}

func findProc(dataService *service.PaymentAppDataService) {
	fmt.Println("Find selected!...")
	username := console.ReadString("Input username:")

	user, err := dataService.FindUserByUsername(username)

	if err != nil {
		fmt.Println("Problem occurred!...")
		return
	}

	if user != nil {
		fmt.Printf("%s found:\n", user.Username)
		fmt.Printf("%s, %s\n", user.Name, user.Phone)
	} else {
		fmt.Printf("%s not found:\n", username)
	}

}

func doForUsersOption(option int, dataService *service.PaymentAppDataService) bool {
	result := true

	switch option {
	case register:
		registerProc(dataService)
	case login:
		loginProc(dataService)
	case find:
		findProc(dataService)
	case quit:
		result = false
	default:
		fmt.Println("Invalid option!...")
	}

	return result
}

func createDataService() *service.PaymentAppDataService {
	connStr := "postgresql://postgres:csd1993@localhost/gs23_paymentdb?sslmode=disable"

	dataService, err := service.NewPaymentAppDataService(connStr)

	if err == nil {
		return dataService
	}
	return nil
}

func Run() {
	dataService := createDataService()
	if dataService == nil {
		fmt.Println("Connection problem. Try again later!...")
		os.Exit(1)
	}
	for {
		displayMainMenu()
		option := console.ReadInt("", "You must input option as a number!...")
		doForMainOption(option, dataService)
	}
}
