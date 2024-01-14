/*
------------------------------------------------------------------------------------------------------------------------

	Date fonksiyonu parametresiyle tarih-zaman bilgilerinin geçerlilik kontrolünü yapmaz. Verilen değer sınırlar dışında
	bile olsa uygun tarih-zaman bilgisini hesaplar. Bu durumda örneğin bir tarihin en son günü basit olarak Date
	fonksiyonu gün ve geri kalan bilgiler sıfır girilerek ay bilgisi de tarihe ilişkin ayın bir sonraki ayı olarak
	girilerek bulunabilir

	Aşağıdaki demo örneği ve ilgili fonksiyonları inceleyiniz

------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"SampleGoLand/csd/console"
	"SampleGoLand/csd/util/ctime"
	"fmt"
	"time"
)

func main() {

	for {
		month := console.ReadInt("Input expiry month:", "Invalid value!...")
		if month <= 0 {
			break
		}

		year := console.ReadInt("Input expiry year:", "Invalid value!...")
		expiryDate := ctime.EndOfMonthYear(time.Month(month), year, time.Local)
		fmt.Println(expiryDate.Format("02/01/2006"))
	}
}
