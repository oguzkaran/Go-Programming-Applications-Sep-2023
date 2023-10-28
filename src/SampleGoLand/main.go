/*
------------------------------------------------------------------------------------------------------------------------

	Sınıf Çalışması: Parametresi ile aldığı gün, ay ve yıl değerlerine ilişkin tarihin kaçıncı günü olduğunu döndüren
	dayOfYear isimli fonksiyonu yazınız ve aşağıdaki kod ile test ediniz
	Açıklamalar:
		- Go'da standart olarak bulunan veya başka bir paketteki tarih zaman işlemi yapan fonksiyonlar kullanılmayacaktır
		- Fonksiyon geçersiz bir tarih değeri içn -1 değerine geri dönecektir

	Not: İleride daha iyisi yazılacaktır

------------------------------------------------------------------------------------------------------------------------
*/

package main

import "fmt"

func main() {
	runDayOfYearTest()
}

func runDayOfYearTest() {
	for {
		var d, m, y int
		fmt.Print("Input date values:")
		fmt.Scan(&d, &m, &y)

		dof := dayOfYear(d, m, y)
		if dof != -1 {
			fmt.Printf("%02d/%02d/%04d is the %d. day of the year\n", d, m, y, dof)
		} else {
			fmt.Println("Invalid date values!...")
		}
		if d == 0 && m == 0 && y == 0 {
			break
		}
	}
}

func dayOfYear(day, month, year int) int {
	//TODO:
}

func isValidDate(day, month, year int) bool {
	return (1 <= day && day <= 31) && (1 <= month && month <= 12) && (1900 <= year) && (day <= getDays(month, year))
}

func getDays(month, year int) int {
	days := 31

	switch month {
	case 4, 6, 9, 11:
		days = 30
	case 2:
		days = 28
		if isLeapYear(year) {
			days++
		}
	}

	return days
}

func isLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}
