/*
------------------------------------------------------------------------------------------------------------------------

	strings.ToUpper ve strings.ToLower metotları

------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	for {
		var str string

		fmt.Print("Input text:")
		fmt.Scan(&str)

		if str == "quit" {
			break
		}

		fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, str))
		fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, str))
	}
}
