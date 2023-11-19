/*
------------------------------------------------------------------------------------------------------------------------

	Sınıf Çalışması: SampleGoLand/datetime paketi içerisine tarih ve zamanı temsil eden CSDDateTime isimli yapıyı
	tasarlayınız ve yazınız

	Açıklamalar: Standart date-time fonksiyonları kullanılmayacaktır

------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"SampleGoLand/datetime"
	"fmt"
)

func main() {
	runApp()
}

func runApp() {
	for {
		var d, m, y int
		fmt.Print("Tarih bilgisini giriniz:")
		fmt.Scan(&d, &m, &y)

		if d == 0 && m == 0 && y == 0 {
			break
		}

		printDateTR(d, m, y)
	}
}

func printDateTR(day, month, year int) {
	date := datetime.NewDate(day, month, year)

	if date == nil {
		fmt.Println("Geçersiz tarih!...")
		return
	}

	fmt.Printf("%02d.%02d.%04d %s\n", day, month, year, date.GetDayOfWeekEN())
}
