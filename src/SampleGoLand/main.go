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

------------------------------------------------------------------------------------------------------------------------
*/

package main

type CSDDate struct {
	day, month, year int
}

func New(day, month, year int) *CSDDate {
	//...
}

func (d *CSDDate) IsLeap() bool {

}

func (d *CSDDate) DayOfYear() int {

}

func (d *CSDDate) Set(day, month, year int) {

}

func (d *CSDDate) Day(day int) {

}

func (d *CSDDate) Month(month int) {

}

func (d *CSDDate) Year(year int) {

}

func dayOfYear(day, month, year int) int {
	if isValidDate(day, month, year) {
		return getDayOfYear(day, month, year) + day
	}

	return -1
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

}
