package app

import (
	"SampleGoLand/csd/console"
	"SampleGoLand/csd/parser/source/str"
	"SampleGoLand/deniz/parser"
	"fmt"
)

func RunDigitCountParserApp() {
	fmt.Println("Letter counter")
	dp := parser.New()

	for {
		s := console.ReadString("Input text:")
		if s == "quit" {
			break
		}

		ss := str.New(s)

		dp.Parse(ss)

		fmt.Printf("Number of digits Count:%d\n", dp.Count())
	}
}
