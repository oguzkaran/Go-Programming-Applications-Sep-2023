/*
------------------------------------------------------------------------------------------------------------------------

	Aşağıdaki demo menü uygulamasını inceleyiniz

	Not: İleride daha iyisi yazılacaktır

------------------------------------------------------------------------------------------------------------------------
*/

package main

import "fmt"

func main() {
	runDemoApp()
}

func runDemoApp() {
	for {
		printMenu()
		option := readOption()
		if option < 1 || option > 5 {
			fmt.Println("Geçersiz seçenek")
			continue
		}

		if option == 5 {
			break
		}

		doForOption(option)
	}

	fmt.Println("Teşekkürler!...")
}

func doForOption(option int) {
	switch option {
	case 1:
		insertProc()
	case 2:
		searchProc()
	case 3:
		deleteProc()
	default:
		updateProc()
	}
}

func insertProc() {
	fmt.Println("------------------------------------------")
	fmt.Println("Ekle")
	fmt.Println("------------------------------------------")
}

func searchProc() {
	fmt.Println("------------------------------------------")
	fmt.Println("Ara")
	fmt.Println("------------------------------------------")
}

func deleteProc() {
	fmt.Println("------------------------------------------")
	fmt.Println("Sil")
	fmt.Println("------------------------------------------")
}

func updateProc() {
	fmt.Println("------------------------------------------")
	fmt.Println("Güncelle")
	fmt.Println("------------------------------------------")
}

func printMenu() {
	fmt.Println("1.Ekle")
	fmt.Println("2.Ara")
	fmt.Println("3.Sil")
	fmt.Println("4.Güncelle")
	fmt.Println("5.Çıkış")
	fmt.Print("Seçenek:")
}

func readOption() int {
	var option int

	fmt.Scan(&option)

	return option
}
