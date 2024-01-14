package app

import (
	"fmt"
	"os"
	"time"
)

func checkArgs(message string) {
	if len(os.Args) != 2 {
		fmt.Println(message)
		os.Exit(1)
	}
}

func checkError(err error, message string) {
	if err != nil {
		fmt.Println(message)
		os.Exit(1)
	}
}

func getAge(birthDate, now time.Time) float64 {
	return float64(now.Unix()-birthDate.Unix()) / float64(60*60*24*365)
}

func Run() {
	checkArgs("Geçersiz komut satırı argümanları!...")
	str := os.Args[1]
	birthDate, err := time.Parse("02/01/2006", str)
	checkError(err, "Geçersiz tarih!...")
	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	birthDay := time.Date(now.Year(), birthDate.Month(), birthDate.Day(), 0, 0, 0, 0, time.UTC)
	age := getAge(birthDate, now)

	switch status := now.Compare(birthDay); status {
	case 1:
		fmt.Printf("Geçmiş doğum gününüz kutlu olsun. Yani yaşınız:%f\n", age)
	case -1:
		fmt.Printf("Doğum gününüz şimdiden kutlu olsun. Yani yaşınız:%f\n", age)
	default:
		fmt.Printf("Doğum gününüz kutlu olsun. Yani yaşınız:%f\n", age)
	}
}
