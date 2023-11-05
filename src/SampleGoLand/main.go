/*
------------------------------------------------------------------------------------------------------------------------

	Sınıf Çalışması: Aşağıdaki açıklamalara göre ilgili yazınız ve test ediniz
	Açıklamalar:
		- Yapı gün ay ve yıl bilgisini tutacaktır:
			type CSDDate struct {
				day, month, year int
			}
		- Yeni Date nesnesinin adresine geri dönen New fonksiyonu yazılacak ve fonksiyon aldığı parametrelere ilişkin
		tarihin geçerliliğini kontrol edecektir, geçersiz bir tarih durumunda nil değerine geri dönecektir

		- İlgili tarihe ilişkin yılın artık yıl olup olmadığına geri dönen IsLeap metodu yazılacaktır

		- İlgili tarihin yılın kaçıncı günü olduğuna geri dönen DayOfYear metodu yazılacaktır

		- Tarihi değiştiren Set isimli metot yazılacaktır

		- Tarihin herhangi bir bileşenini değiştiren Day, Month ve Year metotları yazılacaktır

		- Set işlemi yapan metotların geri dönüş değerleri bool türden olacaktır. İşlemin başarılı olup olmadığında göre
		true ya da false değerine geri döneceklerdir

------------------------------------------------------------------------------------------------------------------------
*/

package main

import "fmt"

type CSDDate struct {
	day, month, year int
}

func New(day, month, year int) *CSDDate {
	if isValidDate(day, month, year) {
		return &CSDDate{day, month, year}
	}

	return nil
}

func (d *CSDDate) IsLeap() bool {
	return isLeapYear(d.year)
}

func (d *CSDDate) DayOfYear() int {
	return getDayOfYear(d.day, d.month, d.year)
}

func (d *CSDDate) DayOfWeek() int {
	return dayOfWeek(d.day, d.month, d.year)
}

func (d *CSDDate) Set(day, month, year int) bool {
	if !isValidDate(day, month, year) {
		return false
	}

	d.day = day
	d.month = month
	d.year = year

	return true
}

func (d *CSDDate) Day(day int) bool {
	if isValidDate(day, d.month, d.year) {
		d.day = day
		return true
	}

	return false
}

func (d *CSDDate) Month(month int) bool {
	if isValidDate(d.day, month, d.year) {
		d.month = month
		return true
	}

	return false
}

func (d *CSDDate) Year(year int) bool {
	if isValidDate(d.day, d.month, year) {
		d.year = year
		return true
	}

	return false
}

func dayOfWeek(day, month, year int) int {
	return totalDays(getDayOfYear(day, month, year), year) % 7
}

func totalDays(dof, year int) int {
	result := dof

	for y := 1900; y < year; y++ {
		result += 365
		if isLeapYear(y) {
			result++
		}
	}

	return result
}

func getDayOfYear(day, month, year int) int {
	result := 0
	switch month - 1 {
	case 11:
		result += 30
		fallthrough
	case 10:
		result += 31
		fallthrough
	case 9:
		result += 30
		fallthrough
	case 8:
		result += 31
		fallthrough
	case 7:
		result += 31
		fallthrough
	case 6:
		result += 30
		fallthrough
	case 5:
		result += 31
		fallthrough
	case 4:
		result += 30
		fallthrough
	case 3:
		result += 31
		fallthrough
	case 2:
		result += 28
		if isLeapYear(year) {
			result++
		}
		fallthrough
	case 1:
		result += 31
	}

	return result
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

func main() {
	runApp()
}

func runApp() {
	for {
		var d, m, y int
		fmt.Print("Tarih bilgisini giriniz:")
		fmt.Scan(&d, &m, &y)
		
		printDateTR(d, m, y)

		if d == 0 && m == 0 && y == 0 {
			break
		}
	}
}

func printDateTR(day, month, year int) {
	switch dow := dayOfWeek(day, month, year); dow {
	case 0:
		fmt.Printf("%02d.%02d.%04d Pazar\n", day, month, year)
	case 1:
		fmt.Printf("%02d.%02d.%04d Pazartesi\n", day, month, year)
	case 2:
		fmt.Printf("%02d.%02d.%04d Salı\n", day, month, year)
	case 3:
		fmt.Printf("%02d.%02d.%04d Çarşamba\n", day, month, year)
	case 4:
		fmt.Printf("%02d.%02d.%04d Perşembe\n", day, month, year)
	case 5:
		fmt.Printf("%02d.%02d.%04d Cuma\n", day, month, year)
	case 6:
		fmt.Printf("%02d.%02d.%04d Cumartesi\n", day, month, year)
	default:
		fmt.Println("Geçersiz tarih!...")
	}
}
