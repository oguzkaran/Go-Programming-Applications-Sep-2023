/*
------------------------------------------------------------------------------------------------------------------------

    Format ve Parse fonksiyonları için önceden tanımlanmış bazı format string'leri vardır. Bunlar resmi dokumanlara
	göre şu şekildedir:

	const (
		Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
		ANSIC       = "Mon Jan _2 15:04:05 2006"
		UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
		RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
		RFC822      = "02 Jan 06 15:04 MST"
		RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
		RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
		RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
		RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
		RFC3339     = "2006-01-02T15:04:05Z07:00"
		RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
		Kitchen     = "3:04PM"
		// Handy time stamps.
		Stamp      = "Jan _2 15:04:05"
		StampMilli = "Jan _2 15:04:05.000"
		StampMicro = "Jan _2 15:04:05.000000"
		StampNano  = "Jan _2 15:04:05.000000000"
		DateTime   = "2006-01-02 15:04:05"
		DateOnly   = "2006-01-02"
		TimeOnly   = "15:04:05"
	)
	Buradaki sabitler gerektiğinde kullanılabilir
------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Date(2020, time.March, 0, 0, 0, 0, 0, time.Local)

	fmt.Println(date)
}
