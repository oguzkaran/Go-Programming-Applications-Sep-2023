/*
------------------------------------------------------------------------------------------------------------------------

	Mantıksal operatörler 3 tanedir: Logical AND (&&), Logical OR(||), Logical NOT(!)

------------------------------------------------------------------------------------------------------------------------
*/
package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Input two numbers:")
	fmt.Scan(&a, &b)

	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a == b)
	fmt.Println(a != b)
}
