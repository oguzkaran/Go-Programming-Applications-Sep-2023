package app

import (
	"SampleGoLand/csd/console"
	"SampleGoLand/csd/parser/source/str"
	"SampleGoLand/deniz/parser"
)

func RunDigitCountParserApp() {
	dp := parser.New()

	for {
		s := console.ReadString("Input text:")
		if s == "quit" {
			break
		}

		ss := str.New(s)

		dp.Parse(ss)

		console.WriteLine("Number of digits Count:%d", dp.Count())
	}
}
