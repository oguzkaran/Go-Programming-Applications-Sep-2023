/*
------------------------------------------------------------------------------------------------------------------------

	Sınıf Çalışması: Parametresi ile bir yazının içerisindeki içerisindeki en uzun ilk palindrome geri dönen
	GetShortestFirsPalindrome fonksiyonunu ve son en uzun palindroma geri dönen GetShortestFirsPalindrome fonksiyonlarını
	yazınız ve test ediniz. Yazıda palindrom yoksa boş string'e geri dönecektir

------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	kb := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Input text:")
		str, _ := kb.ReadString('\n')

		str = strings.TrimRight(str, "\r\n")

		for _, c := range str {
			fmt.Printf("%c ", c)
		}

		fmt.Println()

	}
}
